# Repository Map

Projects live in `~/sideprojects/`. Don't search the filesystem — look here first. Owner = GitHub org/user (`—` = local only). Variants = extra worktrees/clones of the same repo; `*-worktrees`/`*.worktrees`/`worktrees` dirs are worktree holders.

## Worktrees

- For a repository other than the session's current repository, create a new worktree from the freshly fetched default branch. Never edit another repository's main checkout.
- Pet project path: `~/sideprojects/worktrees/<repo>-<topic>`.
- Remove the worktree after its changes merge.

## Pet projects — `~/sideprojects/`

| Repo | Owner | Stack | What |
|------|-------|-------|------|
| 3d-projects | — | FreeCAD | 3D print models (`.FCStd`, `.stl`) |
| 3d-shop | — | Astro | 3D print shop site (no git) |
| ai-casino | Automaat | Python | Agentic stock trading system (variants: `ai-casino-remove-nest-asyncio`, `ai-casino-technical-debt` worktrees) |
| ai-training | — | MCP | AI endurance coaching → intervals.icu via Claude Code |
| akz-analysis | Automaat | Python/HTML | Sailing voyage history analysis + dashboard (AKŻ) |
| auto-claude | smykla-skalski | Go | PR automation daemon driving Claude Code CLI |
| baratie | Automaat | Go/SvelteKit | 🍳 Self-hosted kitchen companion: recipes, pantry, meal plans |
| benchee | smykla-skalski | Python | Reproducible benchmarks for AI code-review tools |
| bilgie | smykla-skalski | Go | Disk-pressure janitor for dev machines running many agents: auto-cleans caches and stale temp dirs (launchd/systemd/Task Scheduler) |
| creatorops | Automaat | Tauri/React | Photography workflow app for macOS |
| deepgram-stt | Automaat | Python | Deepgram speech-to-text for Home Assistant |
| doctorine | Automaat | Go/SvelteKit | Private health status + medical document manager |
| drafts-mcp | Automaat | TypeScript | Drafts MCP server |
| environment-as-code | Automaat | config | This machine's setup (zakwas config: dotfiles, Brewfile, mise) |
| falasurf | Automaat | Astro | Fala Surf & Wind School site (PL/EN) |
| finance-buddy | Automaat | Go/SvelteKit | Finansowa Forteca (personal finance, net worth) |
| flip | Automaat | Next.js | Spanish learning: FSRS + AI mnemonics |
| gh-runners | smykla-skalski | Ansible | Self-hosted GitHub Actions runners provisioning |
| golt | smykla-skalski | Go | Faster golangci-lint (patched x/tools, go-tools, revgrep) |
| gridfinity-generator | Automaat | Svelte | Gridfinity bin generator, 3D preview, STEP/STL export |
| harness | smykla-skalski | Rust | Local multi-agent swarm control plane (Claude, Codex, Gemini, Copilot, OpenCode) |
| home-lab-doc | Automaat | docs | Home lab knowledge base (variant: `home-lab-doc-gh-runners` worktree) |
| home-nas | Automaat | Ansible | Proxmox home NAS IaC (variant: `home-nas-gh-runners` worktree) |
| interview | — | Go/Python | Interview practice scratch (no git) |
| kaggle | Automaat | Python | Kaggle competitions: agents, experiment logs, tooling |
| kaggriculture-git-backup | — | backup | Git dir backup of a Kaggle agriculture project; don't edit |
| kengine | smykla-skalski | Go | klab lab engine daemon (Docker containers/networks) |
| klab | smykla-skalski | Rust | Hands-on troubleshooting/learning platform |
| klab-backend | smykla-skalski | Go | klab backend API |
| klaudiush | smykla-skalski | Go | Claude Code hook validation dispatcher (git/commit/quality rules) |
| klaudiusz-smart-home | Automaat | Nix | Smart home config (variants: `klaudiusz-gh-runners`, `klaudiusz-journald`, `klaudiusz-notifier` worktrees) |
| kuma | Automaat | Go | Kuma mesh fork (personal clone) |
| kuma-policy-validator | Automaat | JavaScript | Static Kuma policy validator app |
| lightroom-mcp | Automaat | Lua | Lightroom Classic MCP server |
| mskalski-blog | Automaat | Astro | Personal blog |
| mskalski-dev | Automaat | Astro | mskalski.dev personal site (variant: `mskalski-dev-zakwas` worktree) |
| opencode-plugin-comment-judge | smykla-skalski | TypeScript | opencode plugin: model judges/rewrites agent comments (`-worktrees` holder) |
| opencode-plugin-plan-review | smykla-skalski | TypeScript | opencode plugin: step-by-step plan review (`.worktrees` holder) |
| portfolio | Automaat | Astro | Portfolio site |
| project-fast-go-linter | — | Go | Workspace of golangci-lint/go-tools/x/tools forks behind golt (no git) |
| renovate-helper | smykla-skalski | Go | gh-renovate-tracker: TUI for Renovate PRs across repos/orgs |
| research-machine | Automaat | docs | Knowledge management (Notion, PARA) + AI research |
| sai | smykla-skalski | Python | Skills for Agentic Intelligence (Claude/Codex plugins) |
| sai-harness | smykla-skalski | Svelte | Clone of `smykla-skalski/sail`: Sail desktop workspace for OpenCode/Claude/Codex agents |
| sailing | — | docs | Cruise plans/menus (no git) |
| sailor-buddy | Automaat | Go/Svelte | Sailing cruise tracker + crew docs |
| security-research | — | docs | Security research/assessments (no git) |
| smycracker | smykla-skalski | Python | Learning/lab platform on Firecracker microVMs |
| smyklot | smykla-skalski | Go | GitHub App: PR approvals/merges via CODEOWNERS + `squash` comments |
| subgen | Automaat | Python | Fork: Whisper subtitle generation for Jellyfin/Plex |
| sui | smykla-skalski | Svelte | Shared Svelte UI components |
| sybra | Automaat | Go | Local agent orchestrator for Claude Code swarms (Wails + Svelte) |
| sybra-testbed | Automaat | JavaScript | Express app: manual-testing target for Sybra |
| sybra-website | Automaat | MDX | Sybra landing page + docs |
| synapse | Automaat | Go | Old-name clone of `Automaat/sybra`; prefer `sybra` |
| vault | — | docs | Second Brain (Obsidian, PARA); symlink to `~/Documents/Obsidian Vault` |
| whisper-hotkey | Automaat | Rust | macOS hotkey voice-to-text with local Whisper |
| zakwas | Automaat | Go | Declarative macOS setup CLI (OSS; config lives in environment-as-code) |
| zsh-clean-history | Automaat | Rust | zsh history cleaner |
