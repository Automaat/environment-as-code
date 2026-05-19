# Fix Review — Categorization Reference

## Categorization Criteria

### Valid (ALL must be true)

- References specific code in PR diff
- Technically correct suggestion
- Doesn't break existing functionality
- Aligns with codebase patterns (verified via search)
- Sound reasoning from reviewer
- No defensive code already prevents the issue

### Invalid (ANY true)

- References code not in this PR
- Already implemented
- Conflicts with existing codebase patterns
- Technically incorrect
- Misunderstands code purpose
- Opinion without technical basis

### Questionable (ANY doubt)

- Ambiguous request
- Multiple valid interpretations
- Trade-offs not clear
- Needs user decision on approach

## Validation Gates

Report a fix ONLY if ALL true:

1. Comment references exact code in diff
2. Traced execution path showing issue/improvement
3. Change doesn't break functionality (verified)
4. Searched codebase for patterns (consistent)
5. Aligns with language/framework best practices
6. No existing defensive code prevents issue

If ANY doubt → questionable → ask user (interactive) or skip (auto).

## Few-Shot Examples

<example>
Comment: "@reviewer: Extract validation to helper function"
Code: Lines 45-78 with inline validation
Reasoning:
1. Request: extract validation logic
2. References: lines 45-78 in diff
3. Technically correct: yes, reduces duplication
4. Similar patterns: found 3 other validation helpers
5. Breaks functionality: no, pure refactor
6. Defensive code: none
Conclusion: VALID — create validateRequest() helper
</example>

<example>
Comment: "@reviewer: Add type hints"
Code: def handle_request(req: Request) -> Response:
Reasoning:
1. Request: add type hints
2. References: lines 89-92
3. Already has type hints: YES
4. Similar code: all handlers typed
5. Breaks functionality: n/a
6. Reviewer misread code
Conclusion: INVALID — type hints already present
</example>

<example>
Comment: "@reviewer: Use async/await instead of promises"
Code: Mixed async/await and .then() chains
Reasoning:
1. Request: convert to async/await
2. References: lines 120-145
3. Technically correct: both work
4. Codebase has both styles
5. Breaks functionality: no
6. Style preference, inconsistent patterns exist
Conclusion: QUESTIONABLE — ask user about style preference
</example>
