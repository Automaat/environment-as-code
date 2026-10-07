---
name: agent-workflow
description: Load Marcin's binding repository map and workflow rules on demand. Use before locating a repository, working outside the current repository, changing code or configuration, running tests or linters, using Git or GitHub, creating issues or PRs, merging, or writing notes.
allowed-tools: Read
---

# Agent Workflow

Load only the references required for the current task. Read each selected reference completely before acting. The selected rules are binding.

## Routing

| Task | Read |
|---|---|
| Find or identify a repository, owner, stack, or worktree root | `references/repositories.md` |
| Work in a repository other than the session repository | `references/repositories.md`, `references/git-github.md` |
| Commit, push, create or edit issues/PRs, review, or merge | `references/git-github.md` |
| Change code/config, run hooks, linters, tests, clusters, or the orchestrator | `references/quality.md` |
| Save a note | `references/quality.md` |

For tasks spanning several rows, read the union once. Do not load the repository map merely because the current repository name is already known.
