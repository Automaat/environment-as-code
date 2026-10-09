# Coding and quality gates

## Tooling

- Use `mise` for tool and dependency management.

## Comments

- Comment only public functions and the most problematic code.
- Explain non-obvious why, never what. Use a terse phrase, not prose.
- Default to no comment when clearer code suffices.

## Linters and hooks

- Fix the root cause of every error. Research the proper fix when unclear.
- Never add linter skip or disable directives such as `eslint-disable`, `ruff ignore`, or `nolint`.
- If proper research does not resolve the problem, ask instead of suppressing it.
- Never use `--no-verify` or work around a hook.
- Zakwas-managed dotfiles are locked copies. Change their source in `environment-as-code`, then run `zakwas apply`; never unlock and edit the installed file.

## Testing

- When debugging end-to-end failures, add extensive logging after the failure.
- Verify cluster checks against real local clusters, not stubbed `kubectl`, before claiming they work.
- Do not use remote GKE or EKS clusters unless explicitly authorized.

## Scratch and build output

- Create one scratch root per task with `mktemp -d "${TMPDIR:-/tmp}/<repo>-<task>.XXXXXX"` and record the path. Point every cache and temp location at it (`CARGO_TARGET_DIR`, `GOCACHE`, `GOTMPDIR`, `npm_config_cache`, `UV_CACHE_DIR`) so one `rm -rf` of that exact path removes everything.
- Remove the root when the task ends, whether it merged, failed or was cancelled. Do it in the final workflow step and in every stop or blocked path. Only the agent that created the root removes it; subagents never do.
- Remove only directories you created in this task. Never delete other sessions' or agents' directories (`sail*`, other worktrees' targets) or anything outside your scratch root.
- Reuse one target directory per worktree (`<root>/target-<worktree>`) for repeated gates of the same revision. Pass the scratch root to review and test subagents so they reuse it instead of creating their own.
- Before a build-heavy task, run `df -h "${TMPDIR:-/tmp}"`. Under 20 GB available: say so before continuing and clean your own leftovers first. If still low, stop and ask.
- Run heavy builds and e2e/browser suites as `build-slot <command>` (`~/.local/bin/build-slot`): it queues until one of 3 machine-wide slots is free, so parallel agents do not push the Mac into swap. Applies to full `cargo build/test`, `go test ./...`, `wdio`/Playwright runs and Docker builds; skip it for quick single-file checks. If `build-slot` is missing, run the command directly.
- This is not a cleanup tool. Sweeping of stale directories is handled by `smykla-skalski/sail#339`.

## Orchestrator

- When running the orchestrator monitor, do not register or schedule a cron. The Go backend owns the loop. Scan the board and triage.

## Notes

- Save notes to `/Users/marcin.skalski/Library/Mobile Documents/iCloud~md~obsidian/Documents/second-brain/0_Inbox`.
- Use Markdown and a good amount of emojis.
