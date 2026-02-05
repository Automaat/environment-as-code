---
description: Find and merge Renovate PRs across sideprojects repos
argument-hint: [--dry-run]
allowed-tools: Read,Grep,Glob,Bash(gh:*,git:*,cd:*,ls:*,find:*,npm:*,pnpm:*,go:*,cargo:*,make:*)
---

# Renovate PR Merger

Merge Renovate dependency update PRs across all ~/sideprojects repositories.

## Arguments

- `--dry-run`: Discovery + analysis only, no merges/fixes

## Workflow

### Phase 1: Discovery 🔍

1. Find all git repos:
   ```bash
   find ~/sideprojects -maxdepth 2 -name ".git" -type d 2>/dev/null
   ```

2. For each repo with GitHub remote, list Renovate PRs:
   ```bash
   gh pr list --author "renovate[bot]" --state open --json number,title,headRefName,statusCheckRollup
   ```

3. Skip repos without GitHub remote or no open Renovate PRs

### Phase 2: Analysis 📊

For each discovered PR:

1. Check CI status: `gh pr checks <number>`
2. Categorize:
   - ✅ **Ready**: All checks passing, no conflicts
   - 🔧 **Need Fixes**: Failing checks that can be auto-fixed (lint, types)
   - 🔄 **Conflict**: Has merge conflicts → trigger `@renovate rebase`
   - ⛔ **Blocked**: Required reviews, unfixable failures

3. Use CoVe pattern: verify categorization by re-checking status before acting

### Phase 3: Fix Failing PRs 🔧

For PRs in "Need Fixes" category:

1. Checkout PR: `gh pr checkout <number>`
2. Identify failures from CI logs
3. Apply fixes:
   - Lint errors → fix properly (NEVER use skip/disable directives)
   - Type errors → fix types
   - Build errors → investigate and fix
4. Commit with signature: `git commit -s -S -m "fix: resolve CI failures"`
5. Push: `git push`
6. Wait for CI: `gh pr checks <number> --watch`
7. Re-categorize based on new status

### Phase 4: Merge 🚀

**Important**: Merging one PR can cause conflicts in others (lock files, etc.)

Merge loop (per repo):
1. Sort ready PRs: patch → minor → major
2. Merge first ready PR: `gh pr merge <number> --squash --delete-branch`
3. After merge, re-check remaining PRs for conflicts:
   ```bash
   gh pr view <number> --json mergeable
   ```
4. If conflicts detected:
   - Comment to trigger Renovate rebase: `gh pr comment <number> -b "@renovate rebase"`
   - Move PR to "Pending Rebase" category, continue with others
5. Repeat until no more ready PRs

This iterative approach ensures each merge is validated before proceeding.

### Phase 5: Report 📋

Generate summary table:

| Repo | PR | Status | Action |
|------|-----|--------|--------|
| repo-name | #123 update-dep | ✅ | Merged |
| repo-name | #124 update-dep | 🔧 | Fixed & Merged |
| repo-name | #125 update-dep | 🔄 | Rebase requested |
| repo-name | #126 update-dep | ❌ | Failed (unfixable) |

## Safety Gates

- ✅ Sign all commits: `-s -S`
- ✅ Never use linter skip/disable directives
- ✅ Squash merge (default)
- ✅ Delete branch after merge
- ❌ Never force push

## Error Handling

- **Rate limited** → wait and retry with backoff
- **No remote** → skip repo silently
- **Merge conflict** → trigger `@renovate rebase`, report as pending rebase
- **Auth failure** → prompt for `gh auth login`
- **CI timeout** → skip after 10 min, report as blocked

## Dry-Run Mode

When `--dry-run` is passed:
- Execute Phase 1 (Discovery) and Phase 2 (Analysis) only
- Report what WOULD be merged/fixed
- Make NO changes (no checkouts, commits, or merges)
