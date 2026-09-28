# environment-as-code

My macOS setup as code. `eac` (Go, in this repo) converges the machine to `eac.yaml`:

- **brew**: GUI apps and a few formulae from `Brewfile` (`brew bundle`, with zap cleanup)
- **mise**: every other CLI tool, pinned in `dotfiles/mise/config.toml` and bumped by Renovate
- **links / templates**: dotfiles symlinked or rendered into `$HOME`
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

Edit a dotfile under `dotfiles/` and it's live immediately (symlink). Add a tool to `dotfiles/mise/config.toml` or `Brewfile`, then `mise run apply`.

## Development

```bash
mise run test               # unit + e2e with fake brew/mise/defaults
mise run test:integration   # adds real `defaults` round-trip
mise run lint
```

See [CLAUDE.md](CLAUDE.md) for layout and conventions.
