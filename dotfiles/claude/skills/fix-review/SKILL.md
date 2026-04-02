---
name: fix-review
description: Analyze and apply unresolved PR review comment fixes. Use when user provides a GitHub PR URL with review comments to address, or asks to fix PR feedback.
argument-hint: [PR-URL] [--auto]
allowed-tools: Read,Grep,Bash(gh:*,git:*),Edit,AskUserQuestion
disable-model-invocation: true
user-invocable: true
---

# Fix PR Review Comments

Process unresolved review comments from `$1`, categorize validity, apply approved fixes.

## Arguments

- `PR-URL` (required) — GitHub PR URL, e.g. `https://github.com/owner/repo/pull/123`
- `--auto` — non-interactive: auto-apply valid fixes, skip questionable, auto-commit

## Phase 1: Fetch & Analyze

In interactive mode, start with EnterPlanMode and remain in plan mode through Phase 2.

### 1. Extract PR Info

```bash
OWNER=$(sed 's|.*github.com/\([^/]*\)/.*|\1|' <<< "$1")
REPO=$(sed 's|.*github.com/[^/]*/\([^/]*\)/.*|\1|' <<< "$1")
PR=$(sed 's|.*/pull/\([0-9]*\).*|\1|' <<< "$1")
```

### 2. Fetch Unresolved Review Threads

Read [references/templates.md](references/templates.md) for the GraphQL query and response filtering.

### 3. Categorize Each Comment

Read [references/examples.md](references/examples.md) in full before categorizing — contains criteria definitions, validation gates, and few-shot examples.

For each comment, reason explicitly:

```text
1. What does reviewer request?         → [extract specific request]
2. What code does it reference?        → [verify in diff via Read]
3. Is suggestion technically correct?  → [validate against language/framework]
4. Does similar code exist?            → [Grep codebase]
5. Would change break functionality?   → [trace execution path]
6. Is defensive code already present?  → [check guards/validation]

Conclusion: Valid/Invalid/Questionable because [specific reason]
```

**Language context:** Go → Effective Go, error handling, concurrency; Python → PEP 8, type hints; Rust → API Guidelines, ownership; TypeScript → type safety; Bash → Google Shell Style Guide.

## Phase 2: Present Findings

Read the presentation format in [references/templates.md](references/templates.md) and use it to group findings by category.

## Phase 3: Get User Decisions (interactive only)

Skip this phase when `--auto` is passed.

Call ExitPlanMode, then for each questionable comment use AskUserQuestion with these option types:

- **Style choices:** a) Apply reviewer's style b) Keep current style c) Standardize across file
- **Ambiguous requests:** a) Apply literal interpretation b) Apply inferred intent c) Skip and reply to reviewer
- **Trade-off decisions:** a) Apply change (accept trade-off) b) Skip (keep current behavior)

Add to valid or skip based on response.

## Phase 4: Apply Fixes

For each approved fix, verify against validation gates in [references/examples.md](references/examples.md), then:

1. Use Edit to apply the change
2. Continue to next fix — do NOT stage or commit yet

After all edits:

```bash
git status
git diff

# Run tests (non-blocking)
npm test 2>/dev/null || pytest 2>/dev/null || go test ./... 2>/dev/null || cargo test 2>/dev/null || true
```

**Auto mode only** — commit after valid fixes applied:

```bash
git add .
git commit -s -S -m "fix: address PR review comments"
```

## Phase 5: Summary

```text
Summary: Fixed N/M unresolved review comments

Applied:   ✓ N valid fixes across K files
Skipped:   ? X questionable  ✗ Y invalid

Changes: [staged and committed | NOT staged — manual review required]

Next steps (interactive):
1. git diff — review changes
2. git add -p — stage selectively
3. git commit -s -S — commit
4. Reply to invalid comments on GitHub
```

## Rules

- Search codebase for patterns before applying any fix (Grep)
- Follow linter rules from `~/.claude/CLAUDE.md` — fix root cause, never use skip/disable directives
- Never auto-commit in interactive mode
- Never mark review threads as resolved
- Never assume on ambiguous comments — ask (interactive) or skip (auto)

<example>
Input: /fix-review https://github.com/owner/repo/pull/42
Behavior: EnterPlanMode → fetch threads → categorize (read examples.md) →
          present findings → ExitPlanMode → ask about questionable → apply valid
          fixes with Edit → run tests → print summary (no auto-commit)
</example>

<example>
Input: /fix-review https://github.com/owner/repo/pull/42 --auto
Behavior: fetch threads → categorize → apply valid fixes immediately →
          skip questionable → git commit -s -S → print summary (no user interaction)
</example>

<example>
Input: PR comment "Add type hints" on def handle_request(req: Request) -> Response:
Phase 1 output: ✗ [N/A] api.py:89 — INVALID, type hints already present
Phase 4: no Edit applied; reply drafted for reviewer
</example>
