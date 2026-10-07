# Git and GitHub

## Git

- Sign commits with `-s -S`.
- Forks: push to the fork, then open the PR from the fork.
- Never use `-F` or file-based commit messages.
- Comply with hooks. Adjust commands to pass; never bypass them.

## Commit messages

- Format: `type(scope): description`; scope required.
- Types: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `build`, `ci`, `chore`, `revert`.
- Title ≤50 characters; body lines ≤72 characters.
- Infra scopes use their own type: `ci(actions)`, `test(test)`, `docs(docs)`, `build(build)`; never `feat(ci)`.
- No PR references or full URLs. No AI attribution or `Co-Authored-By` lines.
- Keep titles short; avoid bare `#` issue references that hooks reject.

Examples: `feat(api): add user endpoint`; `ci(actions): update workflow`.

## GitHub CLI

- Use `gh` as the primary GitHub interface.
- Configure new repositories with squash and merge commits enabled, rebase disabled, and PR title/body used for merge commits:

```bash
gh api repos/{owner}/{repo} --method PATCH -f allow_squash_merge=true -f allow_merge_commit=true -f allow_rebase_merge=false -f merge_commit_title=PR_TITLE -f merge_commit_message=PR_BODY
```

## Merging in `smykla-skalski`

- Agents may merge their own PRs after CI is green and every review comment, including Copilot, is fixed or answered and resolved.
- Merge by posting a PR comment with exact body `squash`; the smyklot bot merges. Never use `gh pr merge` there.
- Do not wait for Copilot to re-review fix commits.

## Subissues and umbrellas

- Add a subissue with `gh-add-subissue <parent> <child> [owner/repo]`.
- Manual GraphQL requires header `GraphQL-Features: sub_issues` on the `addSubIssue` mutation.
- Prefix umbrella titles with ☂️.
- Add smaller issues as subissues and note which can run in parallel.

## PR descriptions

```markdown
## Motivation

> Changelog: type(scope): description

## Implementation information

## Supporting documentation
```

- The title is a conventional commit ≤50 characters.
- Use `> Changelog: skip` when no changelog entry applies.
- Do not mention temporary-directory paths.
- Avoid formal words such as `utilize`, `leverage`, and `facilitate`.

## Code review

- Before posting a PR review, dismiss or handle any existing pending review. A leftover pending review blocks posting.
