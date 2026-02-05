---
description: Apply unresolved PR review comment fixes
argument-hint: [PR-URL]
allowed-tools: Read,Grep,Glob,Bash(gh:*,git:*)
---

# Fix PR Review Comments

**Role**: You are a senior software engineer analyzing PR review feedback to determine validity and implement fixes.

**Task**: Process unresolved review comments from $1, research validity, propose fixes.

**IMPORTANT**: Start by using EnterPlanMode tool for thorough research before making any changes.

## Phase 1: Fetch & Analyze (Plan Mode)

### 1. Extract PR Info

Parse PR URL to get owner/repo/number:

```bash
# From URL like: https://github.com/owner/repo/pull/123
OWNER=$(echo "$1" | sed 's|.*github.com/\([^/]*\)/.*|\1|')
REPO=$(echo "$1" | sed 's|.*github.com/[^/]*/\([^/]*\)/.*|\1|')
PR=$(echo "$1" | sed 's|.*/pull/\([0-9]*\).*|\1|')
```

### 2. Fetch Unresolved Review Threads

**IMPORTANT**: Use heredoc with stdin (`@-`) to avoid character encoding issues with `$` in GraphQL.

**NOTE**: Query excludes `diffHunk` to avoid output truncation (~80KB+). Read files directly when analyzing.

```bash
gh api graphql -f owner="$OWNER" -f repo="$REPO" -F pr="$PR" -F query=@- << 'EOF'
query($owner: String!, $repo: String!, $pr: Int!) {
  repository(owner: $owner, name: $repo) {
    pullRequest(number: $pr) {
      reviewThreads(first: 100) {
        nodes {
          isResolved
          isOutdated
          path
          line
          comments(first: 100) {
            nodes {
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

### 3. Research Each Comment (Chain-of-Thought)

For each unresolved comment, use explicit reasoning:

#### Step 1: Context Gathering

- Read affected file at specified `path:line` (use Read tool)
- Search codebase for similar patterns (Grep)
- Check language/framework best practices

#### Step 2: Validity Analysis

Use CoT reasoning before categorizing:

```text
Reasoning checklist:
1. What does reviewer request? → [extract specific request]
2. What code does it reference? → [verify in diff]
3. Is suggestion technically correct? → [validate against language/framework]
4. Does similar code exist? → [search codebase]
5. Would change break functionality? → [trace execution path]
6. Is there defensive code preventing issue? → [check guards/validation]

Conclusion: Valid/Invalid/Questionable because [specific reason]
```

#### Step 3: Categorize

**Valid** (ALL must be true):

- References specific code in PR diff
- Technically correct suggestion
- Doesn't break existing functionality
- Aligns with codebase patterns (verified via search)
- Sound reasoning from reviewer
- No defensive code prevents issue

**Invalid** (ANY true):

- References code not in this PR
- Already implemented
- Conflicts with existing patterns
- Technically incorrect
- Misunderstands code purpose
- Opinion without technical basis

**Questionable** (if ANY doubt):

- Ambiguous request
- Multiple valid interpretations
- Trade-offs not clear
- Need user decision on approach

### 4. Few-Shot Examples

**Example 1 - Valid:**

```text
Comment: "@reviewer: Extract validation to helper function"
Code: Lines 45-78 with inline validation
Reasoning:
1. Request: extract validation logic
2. References: lines 45-78 in diff
3. Technically correct: yes, reduces duplication
4. Similar patterns: found 3 other validation helpers
5. Breaks functionality: no, pure refactor
6. Defensive code: none
Conclusion: VALID - create validateRequest() helper
```

**Example 2 - Invalid:**

```text
Comment: "@reviewer: Add type hints"
Code: def handle_request(req: Request) -> Response:
Reasoning:
1. Request: add type hints
2. References: lines 89-92
3. Already has type hints: YES
4. Similar code: all handlers typed
5. Breaks functionality: n/a
6. Reviewer misread code
Conclusion: INVALID - type hints already present
```

**Example 3 - Questionable:**

```text
Comment: "@reviewer: Use async/await instead of promises"
Code: Mixed async/await and .then() chains
Reasoning:
1. Request: convert to async/await
2. References: lines 120-145
3. Technically correct: both work
4. Codebase has both styles
5. Breaks functionality: no
6. Style preference, inconsistent patterns exist
Conclusion: QUESTIONABLE - ask user about style preference
```

## Phase 2: Present Findings (Still in Plan Mode)

Group by category, prioritize by severity:

### Valid Fixes

```text
✓ [Critical] file.go:123 (@reviewer)
  Request: Add nil check before dereference
  Current: result.Field without check
  Fix: Add `if result == nil { return err }`
  Why: Prevents panic, aligns with error handling pattern

✓ [Major] api.ts:56 (@reviewer)
  Request: Extract duplicate validation
  Current: Same validation in 3 places
  Fix: Create validateInput() function
  Why: DRY principle, found 3 duplicates
```

### Questionable Comments

```text
? [Minor] handler.py:89 (@reviewer)
  Request: Use list comprehension
  Current: for loop building list
  Question: Codebase has both styles. Prefer comprehensions for consistency?
  Options:
    a) Apply (use comprehensions)
    b) Keep (both styles ok)
    c) Standardize all (larger refactor)
```

### Invalid Comments

```text
✗ [N/A] utils.rs:45 (@reviewer)
  Request: Add error handling
  Invalid: Error handling already exists (lines 47-52)
  Evidence: `match result { Ok(v) => ..., Err(e) => ... }`
  Action: Reply to reviewer with clarification
```

## Phase 3: Get User Decisions

For questionable comments:

- Present question clearly
- Provide context + options
- Wait for user decision
- Add to valid or skip based on response

## Phase 4: Apply Fixes (After ExitPlanMode)

### Process Order

1. Critical (bugs, security, correctness)
2. Major (refactoring, performance)
3. Minor (style, naming, comments)

Within each: batch by file

### Application

```bash
# For each approved fix:
# 1. Use Edit tool to apply change
# 2. Continue to next fix
# 3. DO NOT stage/commit

# After all fixes applied:
git status
git diff

# Run tests (non-blocking)
npm test 2>/dev/null || pytest 2>/dev/null || go test ./... 2>/dev/null || cargo test 2>/dev/null || true
```

### Self-Review Before Applying

For each fix, verify:

1. Does this solve reviewer's concern? [yes/no + why]
2. Could it break anything? [check imports, callers, tests]
3. Aligns with language best practices? [verify]
4. Follows linter rules from CLAUDE.md? [check]

#### If ANY concern → ask user first

## Phase 5: Final Report

```text
Summary: Fixed 5/8 unresolved review comments

Applied:
✓ 5 valid fixes across 3 files
  - 2 critical (nil checks, error handling)
  - 2 major (refactoring, performance)
  - 1 minor (naming)

Skipped:
? 2 questionable (user declined)
✗ 1 invalid (already implemented)

Changes: NOT staged (manual review required)

Next steps:
1. Review changes: git diff
2. Run full test suite
3. Stage selectively: git add -p
4. Commit with -s -S flags
5. Resolve review threads on GitHub
6. Reply to invalid comments with clarification
```

## Validation Gates

**Report fix ONLY if ALL true:**

1. Comment references exact code in diff
2. Traced execution path showing issue/improvement
3. Change doesn't break functionality (verified)
4. Searched codebase for patterns (consistent)
5. Aligns with language/framework best practices
6. No existing defensive code prevents issue

### If ANY doubt → questionable category → ask user

## Language/Framework Context

Include in analysis:

- **Go**: Effective Go, error handling, concurrency patterns
- **Python**: PEP 8, Pythonic idioms, type hints
- **Rust**: API Guidelines, ownership, lifetimes
- **TypeScript**: Best practices, type safety
- **Terraform**: HashiCorp conventions
- **Bash**: Google Shell Style Guide

## Key Rules

**DO:**

- Use EnterPlanMode at start
- Show CoT reasoning for each comment
- Search codebase for patterns (Grep)
- Ask about questionable comments
- Run tests (non-blocking)
- Batch changes by file
- Follow linter rules from `~/.claude/CLAUDE.md`
- Provide precise file:line references
- Include reviewer rationale in fixes

**DON'T:**

- Auto-commit or stage changes
- Mark review threads as resolved
- Skip questionable without asking
- Use linter skip/disable directives
- Auto-push to remote
- Apply fixes without research
- Ignore codebase patterns
- Make assumptions on ambiguous comments

**Reference**: Linter error handling from `~/.claude/CLAUDE.md`:

- Fix properly, research if unclear
- Never use skip/disable directives
- Fix root cause, not symptoms

## Output Format

**Planning phase:**

- Concise per-comment summary
- Clear categorization with reasoning
- Specific questions for questionable
- Evidence for invalid

**Execution phase:**

- file:line references (clickable)
- Before/after code snippets
- Reviewer rationale quoted
- Verification results
- Test status

Extremely concise. Sacrifice grammar for concision.
