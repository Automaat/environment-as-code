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

Runs the zakwas installer (Xcode CLI tools, pinned zakwas release to `~/.local/bin`), clones the repo to `~/sideprojects/environment-as-code`, creates an SSH key and waits until it's added to GitHub, then runs `zakwas apply` (which installs Homebrew and mise itself) and runs it again for the agent plugins and instructions, whose plan needs the CLIs the first one installed. It asks for an optional GitHub token so tool downloads don't hit the API rate limit.

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

Another launchd agent runs `zsh-clean-history` daily at 03:15 instead of starting a cleaner after every interactive shell exits.

Reef is installed from `smykla-skalski/tap`. A launchd agent runs `reef observe run` at login and restarts it after a crash. It records a baseline without starting the scheduler, sampling every 60 seconds with a 512 MiB storage cap. The effective history length depends on sample size. Inspect collection with `launchctl print "gui/$(id -u)/dev.reef.observe"`, `reef report`, and `~/Library/Logs/reef-observe.log`. Stop it with `reef observe stop`; restart it with `launchctl kickstart "gui/$(id -u)/dev.reef.observe"`. A stopped recorder resumes at the next login or `zakwas apply`. Reef keeps observations in its private, bounded state directory.

`zakwas apply` runs `reef agents setup` after Reef is installed from Homebrew. Reef detects Codex, Claude Code, and OpenCode and owns their observation integrations. Review and trust the Codex hook with `/hooks` after setup. EAC no longer ships a copy of the OpenCode plugin or a separate hook installer.

A separate `dev.reef.serve` launchd agent keeps Reef's scheduler available at login, with one 1-CPU/1024-MiB admission slot and live pressure checks. It does not opt any agent into canary treatment; only a Reef-wrapped agent session uses that scheduler. Inspect it with `launchctl print "gui/$(id -u)/dev.reef.serve"` and `reef queue`. Stop new treatment assignments with `reef agents canary off`; restart the scheduler after a Homebrew upgrade with `launchctl kickstart -k "gui/$(id -u)/dev.reef.serve"`.

## Agent plugins

The `agents` section of `zakwas.yaml` declares the plugins of Claude Code, Codex and opencode and the marketplaces they come from. `zakwas apply` installs what's missing, updates plugins to their marketplace versions and, with `prune: true`, removes user-scope plugins and marketplaces that aren't declared. Plugins installed for a single project are left alone. opencode has no plugin system, so zakwas links each declared plugin's skills into `~/.config/opencode/skills`.

Klaudiush hook registrations for enabled providers are managed by `agents.hooks.klaudiush`. The provider settings live in `dotfiles/klaudiush/config.toml.tmpl`; `zakwas apply` refreshes registrations after Klaudiush or its config changes.

To add a plugin, list it under `agents.plugins` as `name@marketplace` (the marketplace must be declared under `agents.marketplaces`), then `zakwas apply`:

```yaml
agents:
  providers: [claude, codex, opencode]   # the default for every entry
  marketplaces:
    sai: smykla-skalski/sai                                       # all providers
    environment-as-code: ./                                       # this repo
    claude-plugins-official: {source: anthropics/claude-plugins-official, providers: [claude]}
  plugins:
    - humanize@sai                                                # Claude Code, Codex and opencode
    - {id: fix-review@environment-as-code, providers: [claude]}   # narrow one plugin
    - gopls-lsp@claude-plugins-official                           # follows its marketplace: Claude Code only
```

`plan` and `apply` read the local copy of each marketplace and never fetch it, so a plugin added to its marketplace after that copy fails the plan with "not in marketplace". `zakwas upgrade` refreshes the marketplaces first, then applies.

A plugin targets its marketplace's providers unless it lists its own, which must be a subset. Every skill here installs for all three agents. opencode can't keep an on-request skill (`disable-model-invocation`, or Codex `allow_implicit_invocation: false`) from running on its own, so those skills rely on their own wording there. A marketplace source is a GitHub `owner/repo`, a git URL or a local path; its name must match the `name` in its manifest.

`plugins/` holds my personal skills as portable [Agent Plugins](https://agent-plugins.org), one skill each: `agent-workflow`, `claude-md-gen`, `go-code-review`, `fix-review` (runs only when invoked) and `kong-ai-intel`. The repo itself is the `environment-as-code` marketplace (`.claude-plugin/marketplace.json` for Claude Code, `.agents/plugins/marketplace.json` for Codex), declared in `zakwas.yaml` as a local path.

The small, provider-neutral instruction core lives in `dotfiles/agents/AGENTS.md`. `agents.instructions` links it as `~/.claude/CLAUDE.md`, `~/.codex/AGENTS.md` and `~/.config/opencode/AGENTS.md`; plain `links` entries give it to Copilot CLI (`~/.copilot/instructions/global.instructions.md`) and `~/AGENTS.md`. Detailed repository and workflow rules load on demand from the shared `agent-workflow` skill.

See [the agent context budget](docs/agent-context-budget.md) for representative payload measurements.

OpenCode's declarative configuration lives in `dotfiles/opencode/`. `zakwas apply` protects `opencode.json`, `cli.json` and the Kong themes; the config enables only Kong AI Gateway models. It deliberately excludes generated plugins and skills, local package state, auth caches and `service.json` because they are machine state or secrets.

## Agent permission profiles

`dotfiles/agents/permission-profiles.json` is the provider-neutral policy. Every rule carries the policy revision and a reason so adapters and task telemetry can explain automated decisions. Explore and review are read-only; build can edit the workspace and run recognized local checks; release reviews remote delivery at the provider permission boundary. Unknown, destructive, secret-bearing, network, deployment and out-of-workspace actions require provider review (OpenCode is the exception, see below).

Claude Code loads the matching protected user agent with `claude --agent explore`, `review`, `build` or `release`. Read-only profiles omit mutation tools, build accepts workspace edits, and release keeps approvals manual. These agents expose no MCP servers.

Codex loads the matching protected profile with `codex --profile explore`, `review`, `build` or `release`. Explore uses automatic review to reject requests outside its read-only boundary. Build and release route eligible approval requests to automatic review; review keeps user approvals. Ordinary workspace reads, edits and local checks stay inside the sandbox. Automatic review evaluates boundary-crossing requests and does not override sandbox decisions or managed requirements. Review exposes Datadog and build exposes Chrome DevTools, with every MCP tool call requiring approval; the other servers stay disabled.

OpenCode exposes the same names as agents but deliberately runs in full auto mode, unlike the Claude Code and Codex profiles above: no tool prompts, including `bash`, pushes, merges, deletes and network access. It still denies direct `.env` reads and the edit tool for `explore` and `review`, but `bash` is unrestricted for every agent, so those two are not a read-only boundary. Datadog tools stay off outside review and Chrome DevTools outside build. Select a profile with the agent picker or `opencode --agent <name>`.

## Development

```bash
mise run plan
mise run lint    # shellcheck + actionlint
```

See [CLAUDE.md](CLAUDE.md) for layout and conventions; zakwas itself is documented in [its repo](https://github.com/Automaat/zakwas).

Codex agent capacity is converged by `dotfiles/codex/configure-agents.py`: eight concurrent subagent threads, excluding the primary agent. Zakwas installs the helper and runs it through the commands module, preserving comments, MCP registrations and other mutable `~/.codex/config.toml` settings. Restart Codex after applying. Hosted sessions retain their host-supplied limit.
