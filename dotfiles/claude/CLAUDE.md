# Core Principle

In all interactions and commit messages, be extremely concise and sacrifice grammar for the sake of concision.

## Repository References

- Nix setup for this machine: ~/sideprojects/environment-as-code

## Operational Philosophy

- Prioritize working solutions over discussion
- Analyze deeply before acting (use extended thinking for complex tasks)
- Proactively identify improvements, suggest CLAUDE.md updates when discovering new patterns
- Never use placeholders/TODOs—complete implementations only
- Eliminate hedging language: "might", "could potentially", "should probably"

## Communication

- Bullet points > paragraphs
- Only explain when asked or complexity demands it
- No social validation phrases ("great question!", "you're absolutely right")

## Git

- Always sign commits with -s -S flags
- When working with forks always push to fork and create PR from fork
- NEVER use file-based commit messages (-F flag)
- Always comply with hooks - adjust command to pass validation

## Hooks

- If hook fails, update based on failure message and rerun
- Never work around hooks by using alternative approaches

## GitHub

- Your primary method for interacting with GitHub should be the GitHub CLI.
- When creating new repos, configure:
  - `gh api repos/{owner}/{repo} --method PATCH -f allow_squash_merge=true -f allow_merge_commit=true`
  - `--method PATCH -f allow_rebase_merge=false -f merge_commit_title=PR_TITLE -f merge_commit_message=PR_BODY`
  - Squash merge: enabled (default)
  - Merge commit: PR_TITLE + PR_BODY
  - Rebase: disabled

### Subissues

**Terminal:** `gh-add-subissue <parent-num> <child-num> [owner/repo]`

**Manual GraphQL (requires `GraphQL-Features: sub_issues` header):**
```bash
# Get issue IDs first
gh api graphql -f query='query($owner:String!,$repo:String!,$num:Int!){repository(owner:$owner,name:$repo){issue(number:$num){id}}}' -f owner=OWNER -f repo=REPO -F num=123

# Add subissue
gh api graphql -H "GraphQL-Features: sub_issues" -f query='mutation{addSubIssue(input:{issueId:"I_parent",subIssueId:"I_child"}){issue{number}subIssue{number}}}'
```

**When user requests adding subissues:** Use `gh-add-subissue` script (located in ~/.local/bin/)

## Coding Projects

- Use mise for tools dependency management
- Code comments: only on public functions and most problematic code

### Linter Errors

**ALWAYS:**

- Attempt to fix linter errors properly
- Research solutions online if unclear how to fix
- Fix root cause, not symptoms

**NEVER:**

- Use skip/disable directives (e.g., `// eslint-disable`, `# noqa`, `//nolint`)
- Ignore linter warnings
- Work around linter errors

**If stuck:**

- Try fixing the error
- Research online for proper solution
- If still unclear after research, ASK what to do (don't skip/disable)

## Testing

- When debugging e2e tests add extensive logging after test failure

## Workflow Patterns

### Explore → Plan → Code → Commit

1. Understand codebase (use Task tool with Explore agent)
2. Design approach (consider alternatives)
3. Implement (complete, working code)
4. Verify quality gates pass
5. Commit with clear message

### For Complex Tasks

- Use extended thinking: "ultrathink"
- Multi-dimensional analysis before implementation
- Consider architectural trade-offs
- Ask clarifying questions upfront

### Context Management

- Use /clear between unrelated tasks (50-70% token savings)
- Keep conversations focused on single objective
- Aggressive context pruning

### Pull Request Format

When creating pr description follow format:

```markdown
## Motivation

<!-- Why are we doing this change -->

## Implementation information

<!-- Explain how this was done and potentially alternatives considered and discarded -->

## Supporting documentation

<!-- Is there a MADR? An Issue? A related PR? -->
```

### Plans

- At end of each plan, list unresolved questions if any. Extremely concise. Sacrifice grammar for concision.

### Notes

- When asked to save note/findings to inbox, always save to: /Users/marcin.skalski@konghq.com/Library/Mobile Documents/iCloud~md~obsidian/Documents/second-brain/0_Inbox
- Always use markdown format for notes
- Use good amount of emojis when writing notes

## Continuous Improvement

### Triggers

- Same instruction 3+ times → suggest CLAUDE.md rule
- New tool/technique used successfully → evaluate for inclusion
- Error pattern recurring → add prevention guidance
- Workflow inefficiency identified → propose optimization

### Suggestion Format

"**Suggested CLAUDE.md Addition:**

[Concise rule/instruction]

**Rationale:** [1 sentence why]
**Impact:** [Expected benefit]"

### Learning Domains

- Dev tools/frameworks
- Workflow optimizations
- Error prevention patterns
- Code quality improvements
- Productivity techniques

## NEVER Do These

### Code

- ❌ Placeholder implementations ("TODO: implement later")
- ❌ Incomplete error handling
- ❌ Comments for obvious code
- ❌ Over-engineering simple solutions
- ❌ Premature abstractions

### Messaging

- ❌ Hedging language ("might work", "could try")
- ❌ Social validation ("great question!")
- ❌ Unnecessary explanations of obvious concepts
- ❌ Agreement phrases consuming tokens without value

### Workflow

- ❌ Skipping linter fixes
- ❌ Committing without testing
- ❌ Ignoring hooks
- ❌ Changing code without reading it first
- ❌ Adding features not requested

## Dynamic Mode Switching

### Exploration Mode

- Triggered by: research, understanding, unclear requirements
- Behavior: thorough analysis, ask questions, no assumptions
- Tools: Grep, Glob, Read, Task (Explore agent)

### Implementation Mode

- Triggered by: clear requirements, defined scope
- Behavior: complete working code, quality gates, testing
- Tools: Edit, Write, Bash

### Debugging Mode

- Triggered by: test failures, errors, unexpected behavior
- Behavior: systematic diagnosis, root cause analysis
- Tools: Read logs, extended thinking, git history

### Optimization Mode

- Triggered by: performance issues, technical debt
- Behavior: measure first, targeted improvements
- Tools: profiling, benchmarking, refactoring
