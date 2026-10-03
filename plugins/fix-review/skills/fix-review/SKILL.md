---
name: fix-review
description: Non-interactive PR review fixer. Fetches unresolved review threads of a GitHub pull request, applies the valid fixes, commits, pushes, and replies to every thread. Run only when the user explicitly invokes fix-review by name; never start it on your own.
allowed-tools: Read Grep Glob Edit Bash(gh:*) Bash(git:*)
disable-model-invocation: true
---

# Fix PR Review Comments (Non-Interactive)

**Role**: Senior software engineer autonomously fixing PR review feedback.

**Task**: Process unresolved review comments on the target PR, research validity, auto-apply valid fixes, skip questionable/invalid.

**Target PR**: the PR URL the user passed when invoking this skill. If none was given, use the PR of the current branch: `gh pr view --json url -q .url`. Stop if neither resolves.

**IMPORTANT**: Work directly — no plan mode (EnterPlanMode or equivalent). Apply fixes immediately after research.

## Phase 1: Fetch & Analyze

### 1. Extract PR Info

Parse PR URL to get owner/repo/number:

```bash
PR_URL="<target PR URL>"
OWNER=$(echo "$PR_URL" | sed 's|.*github.com/\([^/]*\)/.*|\1|')
REPO=$(echo "$PR_URL" | sed 's|.*github.com/[^/]*/\([^/]*\)/.*|\1|')
PR=$(echo "$PR_URL" | sed 's|.*/pull/\([0-9]*\).*|\1|')
```

### 2. Fetch Unresolved Review Threads

```bash
gh api graphql -f owner="$OWNER" -f repo="$REPO" -F pr="$PR" -F query=@- << 'EOF'
query($owner: String!, $repo: String!, $pr: Int!) {
  repository(owner: $owner, name: $repo) {
    pullRequest(number: $pr) {
      reviewThreads(first: 100) {
        nodes {
          id
          isResolved
          isOutdated
          path
          line
          comments(first: 100) {
            nodes {
              databaseId
              author { login }
              body
              createdAt
            }
          }
        }
      }
    }
  }
}
EOF
```

Filter: `isResolved: false` AND `isOutdated: false`

For each surviving thread, capture:
- `threadId` — GraphQL node id (used for `addPullRequestReviewThreadReply`)
- first comment's `databaseId` — REST fallback id for `POST /repos/{O}/{R}/pulls/{N}/comments/{cid}/replies`

### 3. Research Each Comment

For each unresolved comment:

#### Context Gathering
- Read affected file at specified `path:line`
- Search codebase for similar patterns (Grep)
- Check language/framework best practices

#### Categorize

**Valid** (ALL must be true):
- References specific code in PR diff
- Technically correct suggestion
- Doesn't break existing functionality
- Aligns with codebase patterns
- Sound reasoning from reviewer

**Invalid** (ANY true):
- References code not in this PR
- Already implemented
- Conflicts with existing patterns
- Technically incorrect
- Misunderstands code purpose

**Questionable** (ANY doubt):
- Ambiguous request
- Multiple valid interpretations
- Trade-offs not clear

## Phase 2: Apply Fixes

Before the first edit, record the files that already have local changes, staged or not: `git status --porcelain`. Do not edit those files; skip such a fix as questionable ("file has local changes not from this run") so the user's work never lands in the fix commit.

### Process Order
1. Critical (bugs, security, correctness)
2. Major (refactoring, performance)
3. Minor (style, naming)

### For Valid Fixes
- Edit the file to apply the change
- Record `threadId` + one-line note on the fix (for the post-commit reply)

### For Questionable Comments
- **SKIP the fix** — do NOT attempt in non-interactive mode
- Record `threadId` + reasoning (the tradeoff, the open question)

### For Invalid Comments
- **SKIP the fix**
- Record `threadId` + evidence (e.g. "already implemented at file:line", "not in PR diff", "conflicts with pattern X")

## Phase 3: Commit

If no valid fix was applied, make no commit and no push; go to Phase 4.

Otherwise list the files you edited in Phase 2 as `<paths>`. Never `git add .`, and leave anything already staged alone: only `<paths>` go into this commit.

### Commit message

Format: `<type>(<scope>): <description>`, title ≤50 characters, no PR refs, no AI attribution.

- **type**: `fix` by default; `docs`, `test`, `refactor`, `style`, `perf`, `ci`, or `build` when every fix is of that kind. If the scope is `ci`, `test`, `docs`, or `build`, use it as the type too (`test(test): ...`, never `fix(test): ...`).
- **scope**: reuse a scope from recent subjects that matches the touched paths:

  ```bash
  git log -n 50 --format=%s
  git diff --name-only -- <paths>
  ```

  Prefer the scope of recent commits that touched the same files (`git log -n 20 --format=%s -- <path>`). If no recent scope fits, use the narrowest directory name that contains every file in `<paths>` (e.g. `auth` for `internal/auth/*`). If that directory is the repo root, use the top-level directory of the most significant fix, or the repo name for root-only files. Never leave the scope empty. Use only lowercase letters, digits, `-`, `_` and `/`: drop leading dots and turn other dots into `-` (`.github` becomes `github`, `api.v2` becomes `api-v2`).
- **description**: imperative, lowercase, no trailing period; name what the fixes changed (e.g. `fix(auth): handle nil token in refresh`), not that review comments were addressed. For several unrelated fixes, name the most significant one or the common theme.
- Count the full title; if it exceeds 50 characters, shorten the description, not the scope.

Keep apostrophes, backticks, `$` and `!` out of the title (write "do not", not "don't"), then stage and commit `<paths>` in one command, signed, with the title in single quotes. The trailing `-- <paths>` keeps anything else in the index out of the commit:

`git add -- <paths> && git commit -s -S -m '<title>' -- <paths>`

Files you removed with `git rm` are already staged: keep them in the trailing `-- <paths>` of `git commit` but leave them out of `git add`.

### Commit failure

Check the exit code. If the commit fails for any reason (hook rejection, signing error, shell error), read the error and fix the cause: adjust the message, or fix the flagged content, then retry. Change only lines you edited in Phase 2; if a hook flags other code, do not touch it and count the attempt as failed. Never use `--no-verify`, never disable hooks, never drop `-s -S`.

After three failed attempts, stop committing and do not push. Unstage with `git reset -q -- <paths>` and leave the fixes in the working tree for the user to commit, still reply to every thread in Phase 4 (applied fixes use the **Fixed, not pushed** template), and report the uncommitted fixes plus the last error in the Phase 5 summary.

### Push

```bash
git push
```

Never use `git push --no-verify`. If a pre-push hook rejects the push, fix what it flags, commit with the rules above, and retry, at most three times. As for commits, change only lines you edited in Phase 2; if the hook flags other code, do not touch it and count the attempt as failed. If the push still fails, reply as for a failed commit (applied fixes use **Fixed, not pushed**, naming the local SHA) and report the unpushed commit.

Push is required before replying so reviewers see the new SHA alongside the replies.

## Phase 4: Reply to Every Unresolved Thread

For **every** thread processed in Phase 1 — applied, questionable, or invalid — post one reply. Silent skips are the failure mode this skill exists to fix; reviewers must learn the outcome without re-pinging.

### Reply template

```
**Applied** — <one-line description of the change> (<short-sha>).
```

Only when Phase 3 could not commit or push:

```
**Fixed, not pushed** — <one-line description of the change>. <Why it is not on the PR yet, e.g. the hook error>. The change is ready locally and needs a manual commit or push.
```

```
**Skipped (questionable)** — <one-line reasoning>. <Concrete question for reviewer, or tradeoff statement>. Leaving the thread open for your call.
```

```
**Skipped (invalid)** — <one-line evidence>. Happy to revisit if I'm reading this wrong.
```

### Posting

GraphQL (preferred):

```bash
gh api graphql -f threadId="$THREAD_ID" -f body="$REPLY_BODY" -F query=@- << 'EOF'
mutation($threadId: ID!, $body: String!) {
  addPullRequestReviewThreadReply(input: { pullRequestReviewThreadId: $threadId, body: $body }) {
    comment { id url }
  }
}
EOF
```

REST fallback (if mutation unavailable for the repo):

```bash
gh api -X POST "/repos/$OWNER/$REPO/pulls/$PR/comments/$FIRST_COMMENT_DBID/replies" \
  -f body="$REPLY_BODY"
```

### Reply rules

- One reply per thread. Never spam.
- Match the reviewer's terseness — no apologies, no filler, no AI attribution.
- Reference the fix commit's short SHA on applied replies so the link to the diff is obvious.
- Never mark threads as resolved — the reviewer decides.

## Phase 5: Summary (to user)

```text
Summary: Fixed N/M unresolved review comments

Applied (replied + pushed in <sha>):
✓ <thread1>: <one-liner>
✓ ...

Fixed, not pushed (only when Phase 3 failed; <last commit or push error>):
! <thread4>: <one-liner> — in working tree / local commit <sha>

Skipped (replied with reasoning):
? <thread2>: <one-liner> — questionable, awaiting reviewer
✗ <thread3>: <one-liner> — invalid, <evidence>

Threads replied: N+X+Y / total processed
```

## Key Rules

**DO:**
- Research each comment before acting
- Search codebase for patterns
- Apply only clearly valid fixes
- Commit with -s -S flags and a scoped `type(scope): description` title of at most 50 chars
- **Reply to every processed thread** — applied, questionable, invalid
- Reference fix SHA in applied replies
- Log all decisions

**DON'T:**
- Use EnterPlanMode
- Ask for user input
- Apply questionable fixes
- Use linter skip/disable directives
- Bypass git hooks (commit or push), or commit when no fix was applied
- Mark review threads as resolved
- Silently drop a comment — if you read it, you reply to it
