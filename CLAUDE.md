# environment-as-code

Declarative macOS setup driven by `eac`, a small Go CLI that converges this Mac to `eac.yaml`.

## Layout

| Path | What |
|---|---|
| `eac.yaml` | what eac manages: links, templates, brew, mise, defaults, system, commands |
| `Brewfile` | GUI apps (casks) + formulae mise can't install |
| `dotfiles/mise/config.toml` | pinned CLI tools (global mise config; Renovate bumps it) |
| `dotfiles/` | source files symlinked or rendered into `$HOME` |
| `cmd/eac/` | entrypoint + testscript e2e tests (`testdata/script/*.txtar`) |
| `internal/cli/` | flag parsing, module wiring, exit codes |
| `internal/engine/` | `Module` interface, plan/apply |
| `internal/modules/<name>/` | one package per module |
| `internal/runner/` | exec abstraction; `runnertest.Fake` for tests |
| `mise.toml` | repo toolchain (go, golangci-lint) + tasks |

## Commands

```bash
mise run plan                 # go run ./cmd/eac plan
mise run apply                # plan, confirm, apply
go run ./cmd/eac check        # exit 2 on drift
go run ./cmd/eac plan --only brew,defaults
mise run test                 # unit + e2e (fake binaries, no system changes)
mise run test:integration     # + real `defaults` round-trip (macOS only)
mise run lint
```

Exit codes: 0 ok, 1 error, 2 drift (`check`), 64 usage.

## Where things go

| Change | Edit |
|---|---|
| CLI tool | `dotfiles/mise/config.toml` (pin exact version; prefer aqua/github backends) |
| CLI tool not in mise registry / needs system libs | `Brewfile` `brew` |
| GUI app, font | `Brewfile` `cask` |
| Dotfile | put under `dotfiles/`, add to `links` in `eac.yaml` |
| Dotfile needing `$HOME`/vars, or must be a real file | `templates` in `eac.yaml` (Go `text/template`: `.Home`, `.Vars.x`) |
| macOS setting | `defaults` in `eac.yaml` (YAML type picks `-bool/-int/-float/-string`) |
| One-off setup step | `commands` in `eac.yaml` (`run` executes only while `check` fails) |

Find a macOS preference key: `defaults read > a`, toggle in System Settings, `defaults read > b`, `diff a b`.

## Module order

system → links → templates → brew → mise → commands → defaults. Links put the mise config in place; brew installs mise; commands may need tools from either.

## Adding a module

1. `internal/modules/<name>/` implementing `engine.Module` (`Name`, `Plan`); each `Change.Apply` does one idempotent step.
2. Run external commands only through `runner.Runner`; test with `runnertest.Fake`.
3. Config in `internal/config` (validate in `Validate`), wire in `cli.Modules`.
4. Tests: table-driven unit tests + converge twice and assert the second plan is empty.

## Rules

- `plan`/`check` must never change the system; brew calls set `HOMEBREW_NO_AUTO_UPDATE=1`.
- Existing regular files are backed up to `<file>.eac-bak` before linking, never deleted.
- `brew.cleanup: zap` removes anything not in the Brewfile: read the `-` lines of `plan` before `apply`.
- Renovate bumps `dotfiles/mise/config.toml`, `go.mod`, `mise.toml`, GitHub Actions; CI job `install pinned tools` installs every pinned tool on macOS to gate them.
- Claude/Codex/Copilot config is linked by `dotfiles/claude/link.sh`, not eac.
