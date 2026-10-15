#!/usr/bin/env -S uv run --script
# /// script
# requires-python = ">=3.11"
# dependencies = ["tomlkit==0.13.3"]
# ///
"""Converge Codex agent capacity without replacing mutable user configuration."""

from __future__ import annotations

import argparse
import os
import tempfile
from pathlib import Path

import tomlkit


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--config", type=Path, default=Path.home() / ".codex/config.toml")
    parser.add_argument("--check", action="store_true")
    args = parser.parse_args()
    path = args.config
    original = path.read_text() if path.exists() else ""
    document = tomlkit.parse(original)
    if "agents" not in document:
        document["agents"] = tomlkit.table()
    agents = document["agents"]
    agents["enabled"] = True
    agents["max_concurrent_threads_per_session"] = 8
    if "max_threads" in agents:
        del agents["max_threads"]
    updated = tomlkit.dumps(document)
    if args.check:
        return int(updated != original)
    if updated == original:
        return 0
    path.parent.mkdir(parents=True, exist_ok=True)
    descriptor, temporary = tempfile.mkstemp(prefix=".codex-agents-", dir=path.parent)
    try:
        with os.fdopen(descriptor, "w") as output:
            output.write(updated)
        os.chmod(temporary, path.stat().st_mode & 0o777 if path.exists() else 0o600)
        current = path.read_text() if path.exists() else ""
        if current != original:
            raise RuntimeError("Codex configuration changed during update; rerun zakwas apply")
        os.replace(temporary, path)
    finally:
        if os.path.exists(temporary):
            os.unlink(temporary)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
