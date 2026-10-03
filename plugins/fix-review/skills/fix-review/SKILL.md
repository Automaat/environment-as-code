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

Otherwise stage only the files you edited in Phase 2 (never `git add .`), then write the message.

### Commit message

Format: `<type>(<scope>): <description>`, title ≤50 characters, no PR refs, no AI attribution.

- **type**: `fix` by default; `docs`, `test`, `refactor`, `style`, `perf`, `ci`, or `build` when every fix is of that kind.
- **scope**: reuse a scope from recent subjects that matches the touched paths:

  ```bash
  git log -n 50 --format=%s
  git diff --cached --name-only
  ```

  Prefer the scope of recent commits that touched the same files (`git log -n 20 --format=%s -- <path>`). If no recent scope fits, use the narrowest directory name that contains every staged file (e.g. `auth` for `internal/auth/*`). If that directory is the repo root, use the top-level directory of the most significant fix, or the repo name for root-only files. Never leave the scope empty. Lowercase, no spaces.
- **description**: imperative, lowercase, no trailing period; name what the fixes changed (e.g. `fix(auth): handle nil token in refresh`), not that review comments were addressed. For several unrelated fixes, name the most significant one or the common theme.
- Count the full title; if it exceeds 50 characters, shorten the description, not the scope.

Commit signed, with the title as the message: `git commit -s -S -m '<title>'` (single quotes, so backticks and `$` stay literal).

### Hook rejection

If a hook rejects the commit, read its error, adjust the message (or fix the flagged content and re-stage those files with `git add`) to satisfy it, and retry. Never use `--no-verify`, never disable hooks, never drop `-s -S`. After three failed attempts, stop committing: leave the fixes uncommitted, still post the Phase 4 replies for questionable and invalid threads, post no **Applied** replies (there is no SHA), and report the uncommitted fixes plus the hook output in the Phase 5 summary.

### Push

```bash
git push
```

Never use `git push --no-verify`. If a pre-push hook rejects the push, fix what it flags, commit with the rules above, and retry, at most three times. If the push still fails, handle the replies the same way as a failed commit and report the unpushed commit.

Push is required before replying so reviewers see the new SHA alongside the replies.

## Phase 4: Reply to Every Unresolved Thread

For **every** thread processed in Phase 1 — applied, questionable, or invalid — post one reply. Silent skips are the failure mode this skill exists to fix; reviewers must learn the outcome without re-pinging.

### Reply template

```
**Applied** — <one-line description of the change> (<short-sha>).
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
