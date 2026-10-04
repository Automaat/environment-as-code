# Global Instructions

Be extremely concise; sacrifice grammar for concision. Bullets over paragraphs. No hedging ("might", "could potentially") and no social validation ("great question!"). Explain only when asked.

## Detailed Rules

- [Repository map](#repository-map) — where pet (`~/sideprojects/`) and Kong (`~/kong/`) projects live; check before searching the filesystem
- [Git & GitHub](#git--github-conventions) — commits, PRs, `gh`, merging (smykla-skalski), subissues, umbrella issues
- [Coding & quality gates](#coding--quality-gates) — tooling, comments, linters, hooks, testing, notes inbox

## Workflow

- Explore → Plan → Code → Commit. Use the Explore agent to understand the codebase first; consider alternatives before implementing.
- Complex tasks: "ultrathink" for multi-dimensional analysis; ask clarifying questions upfront.
- `/clear` between unrelated tasks.
- Plans: list unresolved questions at the end (terse).
- Work in a repo other than the session's current repo → new git worktree the agent creates; never edit another repo's main checkout. Path: `~/sideprojects/worktrees/<repo>-<topic>` (pet) or `~/kong/worktrees/<repo>-<topic>` (Kong); branch from the freshly fetched default branch; remove the worktree after the work merges.

## Continuous Improvement

- Same instruction 3+ times, or a recurring error pattern → suggest a CLAUDE.md rule. Format: rule + one-line rationale + expected impact.

---

# Repository Map

Projects live in two roots. Don't search the filesystem — look here first. Owner = GitHub org/user (`—` = local only). Variants = extra worktrees/clones of the same repo; `*-worktrees`/`*.worktrees`/`worktrees` dirs are worktree holders.

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
| cache-buster | Automaat | Go | macOS dev cache manager with size limits |
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

## Work projects (Kong) — `~/kong/`

| Repo | Owner | Stack | What |
|------|-------|-------|------|
| afi | Kong | Rust | Small coding agent for self-hosted/remote models |
| agent-auth-cli | Kong | Go | Credential helper: Kong AI Gateway sign-in for coding agents |
| creative_tax | — | scripts | creative_tax (PR listing script; no git) |
| developer.konghq.com | Automaat | Ruby | Fork of Kong developer portal |
| draft-rfc-ai-cost | Kong | docs | Draft RFC: AI gateway/harness budget + cost exposure |
| engineering-ai-gw | Kong | config | GitOps (kongctl) for the engineering Konnect AI Gateway |
| envoy | envoyproxy | C++ | Envoy proxy |
| envoy-builds | kumahq | Shell | Envoy builds for Kuma |
| k8s-argo-apps | kong-konnect | Python/YAML | Konnect ArgoCD apps |
| k8s-cloud-aws | kong-konnect | YAML | Konnect k8s infra resources (AWS) |
| kong | Kong | Lua | Kong Gateway (API + AI gateway) |
| kong-ee | Kong | Lua | Kong Gateway Enterprise |
| kong-mesh | Kong | Go | Kong Mesh |
| kong-mesh-ecs | Kong | IaC | Kong Mesh on ECS + Fargate |
| kong-mesh-smoke | Kong | Go | Kong Mesh smoke tests |
| kong-mesh-v3-ai-driven-verification | Kong | Shell | Agent-driven Kong Mesh v3 upgrade verification |
| konnect-ui-apps | kong-konnect | TypeScript | Konnect UI micro frontends |
| konnect-velocity | — | Go | CLI: kong-konnect org PR efficiency (no per-person metrics) |
| ksai | Kong | Svelte | Kong Skills for Agentic Intelligence: AI review/testing + Claude plugin marketplace |
| ksai-bench | Kong | Python | Benchmark for AI code reviewer findings |
| ksai-cairn | Kong | TypeScript | Pull-through artifact cache + policy gate |
| ksai-cp | kong-konnect | Go | KSAI control plane: GitHub App deliveries → Actions runs |
| ksai-internal | Kong | JavaScript | KSAI internal |
| ksai-swe-bench | Kong | HTML | SWE-bench tasks/evals for the ksai agent stack |
| kubernetes-ingress-controller | Kong | Go | Kong Ingress Controller |
| kuma | kumahq | Go | Kuma service mesh |
| kuma-counter-demo | kumahq | Go | Kuma demo app |
| kuma-website | kumahq | Ruby | Kuma website |
| mink-charts | Kong | Go | MinK (Mesh in Konnect) Helm charts |
| mink-charts-v3 | kong-konnect | Go | MinK Helm charts for Kong Mesh v3 |
| mink-vcp-manager | Kong | Go | Mesh in Konnect VCP manager |
| opencode-plugin-kong-ai-gateway | Kong | TypeScript | opencode provider for Kong AI Gateway (OIDC) |
| oss-cc-scratch | — | scratch | Empty scratch dir |
| oss-cc-test | — | Python | OSS-model Claude Code experiments (glm, gpt-oss; no git) |
| perf-review | — | docs | Perf review notes (no git) |
| platform-api | Kong | OpenAPI | Konnect Platform OpenAPI specs |
| shared-go | Kong | Go | Internal Go modules monorepo (Konnect) |
| shared-ui-components | kong-konnect | TypeScript | Konnect shared UI components |
| sidecarless-gateway-poc | — | Go/K8s | POC: Kong Gateway in mesh without sidecar (SPIRE mTLS; no git) |
| team-mesh | Kong | Go | Mesh team repo: `team-mesh-tools` CLI, docs, automation |
| team-mesh-oncall | Kong | Shell | Agentic oncall runbooks for Kong Mesh / MinK / MoNK |
| tf-kmesh | kong-konnect | Terraform | Konnect mesh policies via TF/Atlantis |
| tf-kong-cloud-02-aws | kong-konnect | Terraform | kong-cloud-02 AWS infra |
| tf-kong-datadog | kong-konnect | Terraform | Datadog monitors/synthetics |
| tf-konnect-gh-runners | kong-konnect | Terraform | Konnect GitHub runners (Atlantis) |
| tf-resources-konnect | kong-konnect | Terraform | TF resources deployed to Konnect |
| v3-readiness | Kong | Go | Clone of `Kong/kong-mesh-v3-readiness`: Kuma 3.0 upgrade preflight CLI |

---

# Git & GitHub Conventions

## Git

- Sign commits with `-s -S`
- Forks: push to fork, open PR from fork
- Never use `-F` / file-based commit messages
- Comply with hooks — adjust the command to pass, never bypass

## Commit Messages

- Format: `type(scope): description` — scope required
- Types: feat, fix, docs, style, refactor, perf, test, build, ci, chore, revert
- Title ≤ 50 chars; body ≤ 72 chars/line
- Infra scopes use their own type: `ci(actions)`, `test(test)`, `docs(docs)`, `build(build)` — never `feat(ci)`
- No PR refs (`#123`, full URLs). No AI attribution (`Co-Authored-By: Claude`)
- Keep titles short; avoid bare `#` issue refs that hooks reject

Examples: ✅ `feat(api): add user endpoint` ✅ `ci(actions): update workflow` ❌ `feat(ci): update workflow` ❌ `feat: add endpoint`

## GitHub CLI

- Use `gh` as the primary GitHub interface
- New repos: enable squash + merge commit (PR_TITLE/PR_BODY), disable rebase:
  - `gh api repos/{owner}/{repo} --method PATCH -f allow_squash_merge=true -f allow_merge_commit=true -f allow_rebase_merge=false -f merge_commit_title=PR_TITLE -f merge_commit_message=PR_BODY`

## Merging — `smykla-skalski` org

- Agents may merge their own PRs once CI is green and every review comment (incl. Copilot) is fixed or answered and resolved
- Merge by posting a PR comment with the exact body `squash` (smyklot bot merges); never `gh pr merge` there
- Don't wait for Copilot to re-review fix commits

## Subissues

- Terminal: `gh-add-subissue <parent> <child> [owner/repo]` (script in `~/.local/bin/`)
- Manual GraphQL needs header `GraphQL-Features: sub_issues` on the `addSubIssue` mutation

## Umbrella Issues

- Prefix title with ☂️
- Add smaller issues as subissues; note which can run in parallel

## PR Description Format

```markdown
## Motivation

## Implementation information

## Supporting documentation
```

PR validation (currently disabled in klaudiush): title is conventional-commit ≤50 chars; body needs Motivation + Implementation sections; changelog line `> Changelog: type(scope): desc` or `> Changelog: skip` after Motivation; no `tmp/` or `/tmp` refs; avoid formal words (utilize, leverage, facilitate).

## Code Review

- Before posting a PR review, dismiss/handle any existing pending review — a leftover pending review blocks posting

---

# Coding & Quality Gates

## Tooling

- Use `mise` for tool/dependency management

## Comments

- Only on public functions and the most problematic code
- Explain non-obvious "why", never "what"; terse phrase, not prose
- Default to none — add only when code can't be made self-explanatory

## Linter Errors

- Fix the root cause; research the proper fix if unclear
- Never use linter skip/disable directives (eslint-disable, ruff ignore, nolint, etc.)
- If still stuck after research, ask — don't skip/disable

## Hooks

- Always fix the root cause; never `--no-verify` or work around
- zakwas: dotfiles are locked copies; change them in the repo and run `zakwas apply`, never unlock and edit the installed file

## Testing

- When debugging e2e failures, add extensive logging after the failure
- Verify against real/live clusters, not stubbed kubectl, before claiming a check works
- Restrict cluster testing to local unless explicitly told to use remote GKE/EKS

## Orchestrator

- Running the orchestrator monitor: do NOT register/schedule a cron — the Go backend owns the loop. Just scan the board and triage.

## Notes to Inbox

- Save to `/Users/marcin.skalski/Library/Mobile Documents/iCloud~md~obsidian/Documents/second-brain/0_Inbox`
- Markdown format; use a good amount of emojis
