# environment-as-code

This Mac's config for [zakwas](https://github.com/Automaat/zakwas) (Go CLI, lives in `~/sideprojects/zakwas`), which converges the machine to `zakwas.yaml`. No Go code here: CLI changes go to the zakwas repo.

## Layout

| Path | What |
|---|---|
| `zakwas.yaml` | what zakwas manages: files, links, templates, brew, mise, agents, defaults, system, commands |
| `Brewfile` | GUI apps (casks) + formulae mise can't install |
| `dotfiles/mise/config.toml` | pinned CLI tools incl. zakwas (global mise config; Renovate bumps it) |
| `dotfiles/` | sources installed into `$HOME` as protected copies |
| `bootstrap.sh` | fresh Mac: zakwas `install.sh`, SSH key, apply, `link.sh` |
| `mise.toml` | repo toolchain (zakwas, shellcheck, actionlint) + tasks |
| `plugins/<plugin>/` | portable agent plugins, one skill each: root `plugin.json` (Agent Plugins, Codex) + `.claude-plugin/plugin.json` + `skills/<skill>/SKILL.md` |
| `.claude-plugin/marketplace.json`, `.agents/plugins/marketplace.json` | Claude Code and Codex marketplaces listing `plugins/` |

## Commands

```bash
zakwas plan [--diff]          # ZAKWAS_CONFIG (zshenv) points at this repo
zakwas apply                  # plan, confirm, apply
zakwas upgrade                # brew update, then apply
zakwas check                  # exit 2 on drift
zakwas plan --only brew,defaults
mise run lint                 # shellcheck + actionlint
```

State: `~/.local/state/zakwas/` (`files.json` hashes, `history.jsonl` applies). Backups: `<file>.zakwas-bak`.

## Where things go

| Change | Edit |
|---|---|
| CLI tool | `dotfiles/mise/config.toml` (pin exact version; prefer aqua/github backends) |
| CLI tool not in mise registry / needs system libs | `Brewfile` `brew` |
| GUI app, font | `Brewfile` `cask` |
| Dotfile | put under `dotfiles/`, add to `files` in `zakwas.yaml` (a dir src copies every file) |
| Config an app must write itself | `links` in `zakwas.yaml` (plain symlink, unprotected) |
| Dotfile needing `$HOME`/vars, or must be a real file | `templates` in `zakwas.yaml` (Go `text/template`: `.Home`, `.Vars.x`) |
| macOS setting | `defaults` in `zakwas.yaml` (YAML type picks `-bool/-int/-float/-string`; `currentHost: true` for ByHost prefs) |
| Agent plugin / marketplace | `agents` in `zakwas.yaml` (`name@marketplace`; `providers` limits it to some of claude, codex, opencode) |
| One-off setup step | `commands` in `zakwas.yaml` (`run` executes only while `check` fails; both run in `$HOME`, `check` times out after 30s) |
| zakwas behavior / new module | `~/sideprojects/zakwas`, release a tag, bump the pin here |

Find a macOS preference key: `defaults read > a`, toggle in System Settings, `defaults read > b`, `diff a b`.

## Protected dotfiles

Installed dotfiles are read-only and `uchg` (`protect.immutable`). To change one: edit it under `dotfiles/`, then `zakwas apply`. For a quick experiment: `chflags nouchg F && chmod u+w F`; `zakwas check` flags it until the change is ported to the repo.

## Rules

- `brew.cleanup: zap` removes anything not in the Brewfile: read the `-` lines of `plan` before `apply`.
- Every third-party `tap` in the Brewfile needs `trusted: true`: `brew bundle cleanup --force` resets Homebrew's trust store to the Brewfile (plan fails otherwise).
- `mise.prune: true` removes installed tool versions no mise config on the machine references.
- zakwas is pinned in both `mise.toml` and `dotfiles/mise/config.toml`; keep them equal (Renovate groups them).
- Renovate bumps `dotfiles/mise/config.toml`, `mise.toml`, GitHub Actions and auto-merges non-major bumps; CI job `e2e on a fresh Mac` applies the config to a clean macOS runner to gate them.
- Claude/Codex/Copilot/opencode instructions are linked by `dotfiles/claude/link.sh`, not zakwas; skills come only from plugins in `agents` (link.sh removes its old skill and command links). CI runs it against a temp `$HOME` and fails if the committed `dotfiles/claude/AGENTS.md` is stale.
- `agents.prune: true` removes undeclared user-scope plugins and marketplaces of Claude Code, Codex and opencode: declare a plugin before installing it by hand. Skills that run only on request (`disable-model-invocation`, or Codex `allow_implicit_invocation: false`) get `providers: [claude, codex]`, since opencode can't stop implicit invocation.
- Shell scripts pass `shellcheck` and workflows pass `actionlint` (both in CI).
- `commands` stops at the first failing entry, so the `git signing key imported` check (fails until the GPG key is imported) stays last.
- Skills follow the Agent Skills spec so opencode loads them: `name` equals the dir name, frontmatter only `name`, `description` (≤1024 chars), `allowed-tools` (space-separated), plus `disable-model-invocation` where a skill must run only on request (Codex: `agents/openai.yaml` `policy.allow_implicit_invocation: false`). Read arguments from the user's request, not `$ARGUMENTS`/`$1`. Plugin names can't start with `claude-`.
- Validate with `claude plugin validate .` and `claude plugin validate plugins/<plugin>`; bump `version` in both `plugin.json` files on functional skill changes.
