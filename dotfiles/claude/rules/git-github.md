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
