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

## Orchestrator

- When running the orchestrator monitor, do not register or schedule a cron. The Go backend owns the loop. Scan the board and triage.

## Notes

- Save notes to `/Users/marcin.skalski/Library/Mobile Documents/iCloud~md~obsidian/Documents/second-brain/0_Inbox`.
- Use Markdown and a good amount of emojis.
