---
description: Non-interactive PR review fix (auto-apply valid fixes)
argument-hint: [PR-URL]
allowed-tools: Read,Grep,Glob,Bash(gh:*,git:*)
---

# Fix PR Review Comments (Non-Interactive)

**Role**: Senior software engineer autonomously fixing PR review feedback.

**Task**: Process unresolved review comments from $1, research validity, auto-apply valid fixes, skip questionable/invalid.

**IMPORTANT**: Work directly — no EnterPlanMode. Apply fixes immediately after research.

## Phase 1: Fetch & Analyze

### 1. Extract PR Info

Parse PR URL to get owner/repo/number:

```bash
OWNER=$(echo "$1" | sed 's|.*github.com/\([^/]*\)/.*|\1|')
REPO=$(echo "$1" | sed 's|.*github.com/[^/]*/\([^/]*\)/.*|\1|')
PR=$(echo "$1" | sed 's|.*/pull/\([0-9]*\).*|\1|')
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
- Use Edit tool to apply change
- Record `threadId` + one-line note on the fix (for the post-commit reply)

### For Questionable Comments
- **SKIP the fix** — do NOT attempt in non-interactive mode
- Record `threadId` + reasoning (the tradeoff, the open question)

### For Invalid Comments
- **SKIP the fix**
- Record `threadId` + evidence (e.g. "already implemented at file:line", "not in PR diff", "conflicts with pattern X")

## Phase 3: Commit

After all valid fixes applied:

```bash
git add .
git commit -s -S -m "fix: address PR review comments"
git push
```

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
- Commit with -s -S flags
- **Reply to every processed thread** — applied, questionable, invalid
- Reference fix SHA in applied replies
- Log all decisions

**DON'T:**
- Use EnterPlanMode
- Ask for user input
- Apply questionable fixes
- Use linter skip/disable directives
- Mark review threads as resolved
- Silently drop a comment — if you read it, you reply to it
