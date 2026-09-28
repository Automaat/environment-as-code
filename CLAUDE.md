# environment-as-code

Declarative macOS setup driven by `eac`, a small Go CLI that converges this Mac to `eac.yaml`.

## Layout

| Path | What |
|---|---|
| `eac.yaml` | what eac manages: files, links, templates, brew, mise, defaults, system, commands |
| `Brewfile` | GUI apps (casks) + formulae mise can't install |
| `dotfiles/mise/config.toml` | pinned CLI tools (global mise config; Renovate bumps it) |
| `dotfiles/` | sources installed into `$HOME` as protected copies |
| `cmd/eac/` | entrypoint + testscript e2e tests (`testdata/script/*.txtar`) |
| `internal/cli/` | flag parsing, module wiring, exit codes |
| `internal/engine/` | `Module` interface, plan/apply |
| `internal/modules/<name>/` | one package per module |
| `internal/install/` | protected-copy writer + hash state (`~/.local/state/eac/files.json`) |
| `internal/runner/` | exec abstraction; `runnertest.Fake` for tests |
| `mise.toml` | repo toolchain (go, golangci-lint) + tasks |

## Commands

```bash
eac plan [--diff]             # ~/.local/bin/eac wraps `go run ./cmd/eac` in the repo
eac apply                     # plan, confirm, apply
eac upgrade                   # brew update, then apply (casks and formulae pick up new versions)
eac check                     # exit 2 on drift
eac plan --only brew,defaults
mise run test                 # unit + e2e (fake binaries, no system changes)
mise run test:integration     # + real `defaults` round-trip (macOS only)
mise run lint
```

Exit codes: 0 ok, 1 error, 2 drift (`check`), 64 usage. A module that fails to plan is reported and skipped; the others still apply, and the run exits 1.

Each apply is appended to `~/.local/state/eac/history.jsonl` (time, commit, changes, error). Applying from a dirty checkout, a branch other than `main`, or behind upstream prints a warning.

## Where things go

| Change | Edit |
|---|---|
| CLI tool | `dotfiles/mise/config.toml` (pin exact version; prefer aqua/github backends) |
| CLI tool not in mise registry / needs system libs | `Brewfile` `brew` |
| GUI app, font | `Brewfile` `cask` |
| Dotfile | put under `dotfiles/`, add to `files` in `eac.yaml` (a dir src copies every file) |
| Config an app must write itself | `links` in `eac.yaml` (plain symlink, unprotected) |
| Dotfile needing `$HOME`/vars, or must be a real file | `templates` in `eac.yaml` (Go `text/template`: `.Home`, `.Vars.x`) |
| macOS setting | `defaults` in `eac.yaml` (YAML type picks `-bool/-int/-float/-string`; `currentHost: true` for ByHost prefs) |
| One-off setup step | `commands` in `eac.yaml` (`run` executes only while `check` fails; both run in `$HOME`, `check` times out after 30s) |

Find a macOS preference key: `defaults read > a`, toggle in System Settings, `defaults read > b`, `diff a b`.

## Protected dotfiles

Like the Nix store, installed dotfiles can't be edited in place: `files` and `templates` write copies with write bits stripped, and `protect.immutable` adds the macOS `uchg` flag (blocks `:w!`, `chmod`, `rm`). To change one: edit it under `dotfiles/`, then `eac apply`.

eac records the hash of each file it writes. On `plan`/`check`:
- repo changed → `~ … (content)`, overwritten
- installed copy edited anyway → `~ … (edited in place, back up to X.eac-bak)`
- protection removed → `~ … (mode 644 → 444)` / `(set immutable)`
- file dropped from `files`/`templates` (or from a dir source) → `- … (no longer managed)`, deleted; backed up first if it was edited

To hand-edit an installed file for a quick experiment: `chflags nouchg F && chmod u+w F`; `eac check` flags it until you port the change to the repo.

## Module order

system → files → links → templates → brew → mise → commands → defaults. Files put the mise config in place; brew installs mise; commands may need tools from either.

## Adding a module

1. `internal/modules/<name>/` implementing `engine.Module` (`Name`, `Plan`); each `Change.Apply` does one idempotent step.
2. Run external commands only through `runner.Runner`; test with `runnertest.Fake`.
3. Config in `internal/config` (validate in `Validate`), wire in `cli.Modules`.
4. Tests: table-driven unit tests + converge twice and assert the second plan is empty.

## Rules

- `plan`/`check` must never change the system; brew calls set `HOMEBREW_NO_AUTO_UPDATE=1`.
- Replaced files that eac didn't write (or that were edited) are backed up to `<file>.eac-bak`, never deleted.
- `brew.cleanup: zap` removes anything not in the Brewfile: read the `-` lines of `plan` before `apply`.
- `mise.prune: true` removes installed tool versions no mise config on the machine references (other projects' configs count, so their tools stay).
- The mise module runs mise from `/`: from `$HOME`, mise treats `~/.config/mise/config.toml` as a project config that outranks the repo file, hiding bumped pins until after `apply`.
- Renovate bumps `dotfiles/mise/config.toml`, `go.mod`, `mise.toml`, GitHub Actions; CI job `e2e on a fresh Mac` applies the config to a clean macOS runner to gate them.
- Claude/Codex/Copilot config is linked by `dotfiles/claude/link.sh`, not eac.
