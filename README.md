# environment-as-code

My macOS setup as code: the config for [zakwas](https://github.com/Automaat/zakwas), which converges the machine to `zakwas.yaml`. It's also a real-world example of a zakwas config (~65 pinned CLI tools, casks, dotfiles, macOS defaults, a daily drift check); for a fresh one, `zakwas init` builds a starter repo from your own Mac.

What it manages:

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

Runs the zakwas installer (Xcode CLI tools, pinned zakwas release to `~/.local/bin`), clones the repo to `~/sideprojects/environment-as-code`, creates an SSH key and waits until it's added to GitHub, then runs `zakwas apply` (which installs Homebrew and mise itself) and links the agent configs. It asks for an optional GitHub token so tool downloads don't hit the API rate limit.

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

## Agent plugins

The `agents` section of `zakwas.yaml` declares the plugins of Claude Code, Codex and opencode and the marketplaces they come from. `zakwas apply` installs what's missing, updates plugins to their marketplace versions and, with `prune: true`, removes user-scope plugins and marketplaces that aren't declared. Plugins installed for a single project are left alone. opencode has no plugin system, so zakwas links each declared plugin's skills into `~/.config/opencode/skills`.

To add a plugin, list it under `agents.plugins` as `name@marketplace` (the marketplace must be declared under `agents.marketplaces`), then `zakwas apply`:

```yaml
agents:
  providers: [claude, codex, opencode]   # the default for every entry
  marketplaces:
    sai: smykla-skalski/sai                                       # all providers
    ksai: {source: Kong/ksai, providers: [claude]}                # Claude Code only
  plugins:
    - humanize@sai                                                # Claude Code, Codex and opencode
    - {id: fix-review@environment-as-code, providers: [claude, codex]}
    - issue-authoring@ksai                                        # follows its marketplace: Claude Code only
```

`plan` and `apply` read the local copy of each marketplace and never fetch it, so a plugin added to its marketplace after that copy fails the plan with "not in marketplace". `zakwas upgrade` refreshes the marketplaces first, then applies.

A plugin targets its marketplace's providers unless it lists its own, which must be a subset. Skills that must run only on request (`disable-model-invocation`, usually for side effects like pushing commits or writing notes) skip opencode, which can't keep a skill from running on its own. A marketplace source is a GitHub `owner/repo`, a git URL or a local path; its name must match the `name` in its manifest.

`plugins/` holds my personal skills as portable [Agent Plugins](https://agent-plugins.org), one skill each: `claude-md-gen`, `go-code-review`, `fix-review` (runs only when invoked) and `kong-ai-intel`. The repo itself is the `environment-as-code` marketplace (`.claude-plugin/marketplace.json` for Claude Code, `.agents/plugins/marketplace.json` for Codex), declared in `zakwas.yaml` as a local path.

`dotfiles/claude/link.sh` links only the instruction files (`CLAUDE.md`, the rules, the status line, `AGENTS.md`); it removes links it used to make for skills and commands.

## Development

```bash
mise run plan
mise run lint    # shellcheck + actionlint
```

See [CLAUDE.md](CLAUDE.md) for layout and conventions; zakwas itself is documented in [its repo](https://github.com/Automaat/zakwas).
