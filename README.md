# environment-as-code

My macOS setup as code: the config for [zakwas](https://github.com/Automaat/zakwas), which converges the machine to `zakwas.yaml`:

- **brew**: GUI apps and a few formulae from `Brewfile` (`brew bundle`, with zap cleanup)
- **mise**: every other CLI tool, pinned in `dotfiles/mise/config.toml` and bumped by Renovate
- **files / templates**: dotfiles installed as read-only, immutable copies; **links** for the few configs apps must write
- **defaults**: macOS preferences
- **system**: Touch ID for sudo, directories, SSH key
- **commands**: guarded one-off steps

## Fresh Mac

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Automaat/environment-as-code/main/bootstrap.sh)"
```

Runs the zakwas installer (Xcode CLI tools, Homebrew, mise, zakwas) with the pinned zakwas release, clones the repo to `~/sideprojects/environment-as-code`, creates an SSH key and waits until it's added to GitHub, then runs `zakwas apply` and links the agent configs. It asks for an optional GitHub token so tool downloads don't hit the API rate limit.

Git signs every commit with the GPG key in `dotfiles/git/config`, which bootstrap can't create. Until it's imported, `zakwas apply` fails on the `git signing key imported` command with the steps:

```bash
gpg --export-secret-keys --armor C25DFDF396ADF455 > key.asc   # on the old Mac
gpg --import key.asc && rm key.asc                             # on the new one
zakwas apply
```

## Day to day

```bash
zakwas plan            # what would change (--diff shows file contents)
zakwas apply           # show plan, confirm, apply
zakwas upgrade         # refresh Homebrew's package list, then apply: upgrades casks and formulae
zakwas check           # exit 2 when the machine drifted
```

zakwas is pinned in `dotfiles/mise/config.toml` and `mise.toml` (Renovate bumps both in one PR); `ZAKWAS_CONFIG` in `zshenv` points it at this repo from any directory.

Installed dotfiles are read-only. Edit them under `dotfiles/` (or add a tool to `dotfiles/mise/config.toml` / `Brewfile`), then `zakwas apply`. Removing a file from `zakwas.yaml` removes it from `$HOME` on the next apply.

A launchd agent runs `zakwas check` every day at 10:00 and shows a notification when the Mac drifted.

## Development

```bash
mise run plan
mise run lint    # shellcheck + actionlint
```

See [CLAUDE.md](CLAUDE.md) for layout and conventions; zakwas itself is documented in [its repo](https://github.com/Automaat/zakwas).
