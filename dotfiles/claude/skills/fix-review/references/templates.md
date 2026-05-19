# Templates

## GraphQL: Fetch Unresolved Review Threads

Use heredoc with stdin (`@-`) to avoid `$` encoding issues in GraphQL. Exclude `diffHunk` to avoid output truncation — read files directly when needed.

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

## Presentation Format

Group by category, prioritize by severity (Critical → Major → Minor):

```text
✓ [Critical] file.go:123 (@reviewer)
  Request: Add nil check before dereference
  Current: return result.Field
  Fix: if result == nil { return err }; return result.Field

? [Minor] handler.py:89 (@reviewer)
  Request: Use list comprehension
  Question: Codebase has both styles — apply?
  Options: a) Apply  b) Keep  c) Standardize all

✗ [N/A] utils.rs:45 (@reviewer)
  Request: Add error handling
  Invalid: Error handling already exists (lines 47-52)
  Action: Reply to reviewer with clarification
```
