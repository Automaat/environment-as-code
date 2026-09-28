# environment-as-code

My macOS setup as code. `eac` (Go, in this repo) converges the machine to `eac.yaml`:

- **brew**: GUI apps and a few formulae from `Brewfile` (`brew bundle`, with zap cleanup)
- **mise**: every other CLI tool, pinned in `dotfiles/mise/config.toml` and bumped by Renovate
- **files / templates**: dotfiles installed as read-only, immutable copies (like the Nix store); **links** for the few configs apps must write
- **defaults**: macOS preferences
- **system**: Touch ID for sudo, directories, SSH key
- **commands**: guarded one-off steps

## Fresh Mac

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Automaat/environment-as-code/main/bootstrap.sh)"
```

Installs Xcode CLI tools, Homebrew and mise, clones the repo to `~/sideprojects/environment-as-code`, creates an SSH key and waits until it's added to GitHub, then runs `eac apply`. It asks for an optional GitHub token so tool downloads don't hit the API rate limit.

## Day to day

```bash
eac plan            # what would change (--diff shows file contents)
eac apply           # show plan, confirm, apply
eac upgrade         # refresh Homebrew's package list, then apply: upgrades casks and formulae
eac check           # exit 2 when the machine drifted
```

`eac` is installed to `~/.local/bin/eac` and always runs the repo's current code.

Installed dotfiles are read-only. Edit them under `dotfiles/` (or add a tool to `dotfiles/mise/config.toml` / `Brewfile`), then `eac apply`. `check` reports any file edited in place; `apply` backs it up before restoring the repo version. Removing a file from `eac.yaml` removes it from `$HOME` on the next apply.

A launchd agent runs `eac check` every day at 10:00 and shows a notification when the Mac drifted. Every apply is logged to `~/.local/state/eac/history.jsonl` with the commit it ran from.

## Development

```bash
mise run test               # unit + e2e with fake brew/mise/defaults
mise run test:integration   # adds real `defaults` round-trip
mise run lint
```

See [CLAUDE.md](CLAUDE.md) for layout and conventions.
