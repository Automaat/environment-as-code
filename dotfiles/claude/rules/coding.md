# Coding & Quality Gates

## Tooling

- Use `mise` for tool/dependency management

## Comments

- Only on public functions and the most problematic code
- Explain non-obvious "why", never "what"; terse phrase, not prose
- Default to none — add only when code can't be made self-explanatory

## Linter Errors

- Fix the root cause; research the proper fix if unclear
- Never use linter skip/disable directives (eslint-disable, ruff ignore, nolint, etc.)
- If still stuck after research, ask — don't skip/disable

## Hooks

- Always fix the root cause; never `--no-verify` or work around
- Nix: fix `.nix` files / `nix flake check` errors, don't skip `darwin-rebuild`

## Testing

- When debugging e2e failures, add extensive logging after the failure
- Verify against real/live clusters, not stubbed kubectl, before claiming a check works
- Restrict cluster testing to local unless explicitly told to use remote GKE/EKS

## Orchestrator

- Running the orchestrator monitor: do NOT register/schedule a cron — the Go backend owns the loop. Just scan the board and triage.

## Notes to Inbox

- Save to `/Users/marcin.skalski/Library/Mobile Documents/iCloud~md~obsidian/Documents/second-brain/0_Inbox`
- Markdown format; use a good amount of emojis
