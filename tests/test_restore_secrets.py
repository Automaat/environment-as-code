import os
from pathlib import Path
import subprocess
import tempfile
import unittest


REPO = Path(__file__).resolve().parents[1]
SCRIPT = REPO / "dotfiles/bin/restore-secrets"

FAKE_BW = """#!/bin/sh
echo "$*" >> "$HOME/bw.log"
case "$1" in
    status) cat "$HOME/bw-status" ;;
    login | unlock) echo fake-session ;;
    get)
        note="$HOME/notes/$3"
        [ -f "$note" ] || exit 1
        cat "$note"
        ;;
esac
"""

FAKE_AGE_KEYGEN = """#!/bin/sh
grep -q '^BROKEN' "$2" && exit 1
sed -n 's/^AGE-SECRET-KEY-/age1/p' "$2"
"""


class RestoreSecretsTest(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory(prefix="restore-secrets-test-")
        self.addCleanup(self.temp.cleanup)
        self.home = Path(self.temp.name)
        tools = self.home / "tools"
        tools.mkdir()
        for name, body in (("bw", FAKE_BW), ("age-keygen", FAKE_AGE_KEYGEN)):
            tool = tools / name
            tool.write_text(body)
            tool.chmod(0o755)

        manifest = self.home / ".config/restore-secrets/secrets.tsv"
        manifest.parent.mkdir(parents=True)
        manifest.write_text(
            "# comment\n"
            "~/keys/keys.txt\tnote one\tage1ONE\n"
            "~/keys/keys.txt\tnote two\tage1TWO\n"
            "~/keys/runner.txt\tnote three\tage1THREE\n"
        )
        (self.home / "bw-status").write_text('{"status":"locked"}\n')
        notes = self.home / "notes"
        notes.mkdir()
        (notes / "note one").write_text("AGE-SECRET-KEY-ONE")
        (notes / "note two").write_text("AGE-SECRET-KEY-TWO")
        (notes / "note three").write_text("AGE-SECRET-KEY-THREE\n")
        self.keys = self.home / "keys/keys.txt"
        self.runner = self.home / "keys/runner.txt"

        self.env = os.environ.copy()
        self.env["HOME"] = str(self.home)
        self.env["PATH"] = str(tools) + os.pathsep + self.env["PATH"]
        self.primary, self.terminal = os.openpty()
        self.addCleanup(os.close, self.primary)
        self.addCleanup(os.close, self.terminal)

    def run_script(self, *args, tty=True):
        return subprocess.run(
            [str(SCRIPT), *args],
            env=self.env,
            stdin=self.terminal if tty else subprocess.DEVNULL,
            capture_output=True,
            text=True,
            timeout=30,
        )

    def bw_calls(self):
        log = self.home / "bw.log"
        return log.read_text().splitlines() if log.exists() else []

    def write_all_keys(self):
        self.keys.parent.mkdir()
        self.keys.write_text("AGE-SECRET-KEY-ONE\nAGE-SECRET-KEY-TWO\n")
        self.runner.write_text("AGE-SECRET-KEY-THREE\n")

    def test_check_fails_without_touching_the_vault(self):
        result = self.run_script("--check")
        self.assertEqual(result.returncode, 1)
        self.assertIn("note one", result.stderr)
        self.assertEqual(self.bw_calls(), [])
        self.assertFalse(self.keys.exists())

    def test_check_passes_when_every_key_is_present(self):
        self.write_all_keys()
        self.assertEqual(self.run_script("--check").returncode, 0)
        self.assertEqual(self.bw_calls(), [])

    def test_restores_all_keys_privately_and_locks(self):
        result = self.run_script()
        self.assertEqual(result.returncode, 0, result.stderr)
        self.assertEqual(self.run_script("--check").returncode, 0)
        self.assertEqual(self.runner.stat().st_mode & 0o777, 0o600)
        self.assertEqual(self.keys.stat().st_mode & 0o777, 0o600)
        self.assertEqual(self.bw_calls()[-1], "lock")
        self.assertNotIn("AGE-SECRET-KEY", result.stdout + result.stderr)

    def test_keeps_existing_keys_and_adds_only_the_missing_one(self):
        self.keys.parent.mkdir()
        self.keys.write_text("# mine\nAGE-SECRET-KEY-OTHER\nAGE-SECRET-KEY-ONE")
        result = self.run_script()
        self.assertEqual(result.returncode, 0, result.stderr)
        content = self.keys.read_text()
        self.assertIn("AGE-SECRET-KEY-OTHER\nAGE-SECRET-KEY-ONE\n", content)
        self.assertEqual(content.count("AGE-SECRET-KEY-ONE"), 1)
        self.assertIn("AGE-SECRET-KEY-TWO\n", content)
        self.assertNotIn("get notes note one", self.bw_calls())

    def test_does_nothing_when_every_key_is_present(self):
        self.write_all_keys()
        self.assertEqual(self.run_script().returncode, 0)
        self.assertEqual(self.bw_calls(), [])

    def test_wrong_note_changes_nothing_and_still_locks(self):
        self.keys.parent.mkdir()
        original = "AGE-SECRET-KEY-ONE\n"
        self.keys.write_text(original)
        (self.home / "notes/note two").write_text("AGE-SECRET-KEY-WRONG\n")
        result = self.run_script()
        self.assertEqual(result.returncode, 1)
        self.assertIn("is unchanged", result.stderr)
        self.assertEqual(self.keys.read_text(), original)
        self.assertEqual(self.bw_calls()[-1], "lock")
        self.assertEqual(sorted(p.name for p in self.keys.parent.iterdir()), ["keys.txt"])

    def test_missing_note_changes_nothing(self):
        (self.home / "notes/note three").unlink()
        result = self.run_script()
        self.assertEqual(result.returncode, 1)
        self.assertIn('Couldn\'t read "note three"', result.stderr)
        self.assertFalse(self.runner.exists())
        self.assertEqual(sorted(p.name for p in self.runner.parent.iterdir()), ["keys.txt"])

    def test_corrupt_file_is_left_alone(self):
        self.keys.parent.mkdir()
        self.keys.write_text("BROKEN\n")
        result = self.run_script()
        self.assertEqual(result.returncode, 1)
        self.assertEqual(self.keys.read_text(), "BROKEN\n")

    def test_refuses_to_prompt_without_a_terminal(self):
        result = self.run_script(tty=False)
        self.assertEqual(result.returncode, 1)
        self.assertIn("in a terminal", result.stderr)
        self.assertEqual(self.bw_calls(), [])

    def test_logs_in_when_not_signed_in(self):
        (self.home / "bw-status").write_text('{"status":"unauthenticated"}\n')
        self.assertEqual(self.run_script().returncode, 0)
        self.assertIn("login --raw", self.bw_calls())
        self.assertNotIn("unlock --raw", self.bw_calls())

    def test_manifest_in_the_repo_is_well_formed(self):
        manifest = REPO / "dotfiles/secrets/bitwarden.tsv"
        rows = [
            line.split("\t")
            for line in manifest.read_text().splitlines()
            if line and not line.startswith("#")
        ]
        self.assertTrue(rows)
        for path, item, recipient in rows:
            self.assertTrue(path.startswith("~/"))
            self.assertTrue(item)
            self.assertTrue(recipient.startswith("age1"))


if __name__ == "__main__":
    unittest.main()
