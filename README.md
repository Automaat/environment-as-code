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
curl -fsSL https://raw.githubusercontent.com/Automaat/environment-as-code/main/bootstrap.sh | bash
```

Installs Xcode CLI tools, Homebrew and mise, clones the repo to `~/sideprojects/environment-as-code`, then runs `eac apply`.

## Day to day

```bash
mise run plan     # what would change
mise run apply    # show plan, confirm, apply
go run ./cmd/eac check   # exit 2 when the machine drifted
```

Installed dotfiles are read-only. Edit them under `dotfiles/` (or add a tool to `dotfiles/mise/config.toml` / `Brewfile`), then `mise run apply`. `check` reports any file edited in place; `apply` backs it up before restoring the repo version.

## Development

```bash
mise run test               # unit + e2e with fake brew/mise/defaults
mise run test:integration   # adds real `defaults` round-trip
mise run lint
```

See [CLAUDE.md](CLAUDE.md) for layout and conventions.
