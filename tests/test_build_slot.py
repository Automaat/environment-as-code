import os
from pathlib import Path
import shutil
import signal
import subprocess
import tempfile
import time
import unittest


SCRIPT = Path(__file__).resolve().parents[1] / "dotfiles/bin/build-slot"


class BuildSlotTest(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory(prefix="build-slot-test-")
        self.root = Path(self.temp.name)
        self.processes = []
        self.default_root = Path("/tmp") / ("build-slots-test-" + self.root.name)
        tools = self.root / "bin"
        tools.mkdir()
        identity = tools / "id"
        identity.write_text("#!/bin/sh\nprintf '%s\\n' test-" + self.root.name + "\n")
        identity.chmod(0o755)
        self.env = os.environ.copy()
        self.env.pop("BUILD_SLOT_DIR", None)
        self.env.pop("BUILD_SLOTS", None)
        self.env["PATH"] = str(tools) + os.pathsep + self.env["PATH"]

    def tearDown(self):
        for process in self.processes:
            if process.poll() is None:
                os.killpg(process.pid, signal.SIGTERM)
        for process in self.processes:
            try:
                process.wait(timeout=5)
            except subprocess.TimeoutExpired:
                os.killpg(process.pid, signal.SIGKILL)
                process.wait()
        shutil.rmtree(self.default_root, ignore_errors=True)
        self.temp.cleanup()

    def start(self, name, directory=None, command=None):
        temporary = self.root / name
        temporary.mkdir()
        environment = dict(self.env, TMPDIR=str(temporary))
        if directory is not None:
            environment["BUILD_SLOT_DIR"] = str(directory)
        marker = temporary / "started"
        if command is None:
            command = ["sh", "-c", 'touch "$1"; while [ ! -e "$2" ]; do sleep 0.05; done',
                       "worker", str(marker), str(temporary / "release")]
        process = subprocess.Popen(["bash", str(SCRIPT), *command], env=environment,
                                   start_new_session=True)
        self.processes.append(process)
        return process, marker

    def wait_for(self, predicate):
        deadline = time.monotonic() + 8
        while not predicate():
            self.assertLess(time.monotonic(), deadline, "timed out waiting for worker")
            time.sleep(0.03)

    def test_private_temporary_directories_share_default_three_slots(self):
        holders = [self.start(str(index)) for index in range(3)]
        self.wait_for(lambda: all(marker.exists() for _, marker in holders))
        waiter, marker = self.start("waiter")
        time.sleep(2.3)
        self.assertFalse(marker.exists())
        self.assertIsNone(waiter.poll())
        holders[0][1].with_name("release").touch()
        self.assertEqual(holders[0][0].wait(timeout=5), 0)
        self.wait_for(marker.exists)
        self.assertTrue(all(process.poll() is None for process, _ in holders[1:]))

    def test_explicit_directory_overrides_default(self):
        holders = [self.start(str(index)) for index in range(3)]
        self.wait_for(lambda: all(marker.exists() for _, marker in holders))
        process, marker = self.start("override", self.root / "override-slots")
        self.wait_for(marker.exists)
        self.assertIsNone(process.poll())

    def test_failure_releases_slot(self):
        process, _ = self.start("failure", command=["sh", "-c", "exit 7"])
        self.assertEqual(process.wait(timeout=5), 7)
        self.assertEqual(list(self.default_root.glob("slot-*")), [])

    def test_cancel_waiter_preserves_owned_slots(self):
        holders = [self.start(str(index)) for index in range(3)]
        self.wait_for(lambda: all(marker.exists() for _, marker in holders))
        waiter, marker = self.start("waiter")
        time.sleep(0.2)
        os.killpg(waiter.pid, signal.SIGTERM)
        self.assertEqual(waiter.wait(timeout=5), 130)
        self.assertFalse(marker.exists())
        self.assertEqual(len(list(self.default_root.glob("slot-*"))), 3)
        os.killpg(holders[0][0].pid, signal.SIGTERM)
        self.assertEqual(holders[0][0].wait(timeout=5), 130)
        self.assertEqual(len(list(self.default_root.glob("slot-*"))), 2)
        self.assertTrue(all(process.poll() is None for process, _ in holders[1:]))


if __name__ == "__main__":
    unittest.main()
