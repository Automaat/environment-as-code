"""Exercise Codex config convergence against mutable user settings."""

import importlib.util
import sys
import tempfile
import unittest
from pathlib import Path
from unittest.mock import patch

import tomlkit

SCRIPT = Path(__file__).resolve().parents[2] / "dotfiles/codex/configure-agents.py"
SPEC = importlib.util.spec_from_file_location("configure_agents", SCRIPT)
MODULE = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(MODULE)


class ConfigureAgentsTest(unittest.TestCase):
    def setUp(self):
        self.directory = tempfile.TemporaryDirectory()
        self.addCleanup(self.directory.cleanup)
        self.path = Path(self.directory.name) / "config.toml"

    def invoke(self, *arguments):
        with patch.object(sys, "argv", [str(SCRIPT), "--config", str(self.path), *arguments]):
            return MODULE.main()

    def test_missing_config_check_is_read_only_then_apply_converges(self):
        self.assertEqual(self.invoke("--check"), 1)
        self.assertFalse(self.path.exists())
        self.assertEqual(self.invoke(), 0)
        self.assertEqual(self.invoke("--check"), 0)
        agents = tomlkit.parse(self.path.read_text())["agents"]
        self.assertTrue(agents["enabled"])
        self.assertEqual(agents["max_concurrent_threads_per_session"], 8)
        self.assertEqual(self.path.stat().st_mode & 0o777, 0o600)

    def test_preserves_comments_other_settings_and_permissions(self):
        self.path.write_text('# keep this comment\nmodel = "example"\n[mcp_servers.custom]\ncommand = "custom-mcp"\n[agents]\nmax_threads = 2\ndefault_subagent_model = "example-worker"\n')
        self.path.chmod(0o640)
        self.invoke()
        first = self.path.read_text()
        self.assertIn("# keep this comment", first)
        config = tomlkit.parse(first)
        self.assertEqual(config["model"], "example")
        self.assertEqual(config["mcp_servers"]["custom"]["command"], "custom-mcp")
        self.assertEqual(config["agents"]["default_subagent_model"], "example-worker")
        self.assertNotIn("max_threads", config["agents"])
        self.assertEqual(self.path.stat().st_mode & 0o777, 0o640)
        self.invoke()
        self.assertEqual(self.path.read_text(), first)

    def test_inline_agents_table_is_supported(self):
        self.path.write_text("agents = {enabled = false, max_threads = 1}\n")
        self.invoke()
        self.assertEqual(self.invoke("--check"), 0)

    def test_concurrent_mcp_registration_is_never_overwritten(self):
        original = 'model = "example"\n'
        concurrent = '[mcp_servers.new]\ncommand = "new-mcp"\n'
        self.path.write_text(original)
        chmod = MODULE.os.chmod

        def register_during_staging(path, mode):
            chmod(path, mode)
            self.path.write_text(original + concurrent)

        with patch.object(MODULE.os, "chmod", side_effect=register_during_staging):
            with self.assertRaisesRegex(RuntimeError, "changed during update"):
                self.invoke()
        self.assertEqual(self.path.read_text(), original + concurrent)
        self.assertEqual(list(self.path.parent.glob(".codex-agents-*")), [])

    def test_invalid_config_is_never_overwritten(self):
        original = "[agents\n"
        self.path.write_text(original)
        with self.assertRaises(tomlkit.exceptions.ParseError):
            self.invoke()
        self.assertEqual(self.path.read_text(), original)


if __name__ == "__main__":
    unittest.main()
