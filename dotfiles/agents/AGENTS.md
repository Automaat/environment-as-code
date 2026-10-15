# Global Instructions

Use standard technical English. Bullets over paragraphs. No hedging ("might", "could potentially") or social validation ("great question!"). Explain only when asked.

## Workflow

- Complex tasks: use deep, multi-dimensional analysis; ask material clarifying questions upfront.
- Clear context between unrelated tasks. End plans with terse unresolved questions.
- Use `$ship-it` only when asked or for complex issues.
- For Kuma changes, open pull requests from a fork.
- A request to implement, fix, or ship work authorizes creating a pull request and writing or updating its title and description for that work. Do not ask for a separate draft approval or PR-creation confirmation, including when using `issue-authoring`. Honor explicit draft-only or no-PR instructions; follow the repository's checks, review, and merge rules.
- Use the `agent-workflow` skill before locating a repository, working outside the current repository, changing code or configuration, running quality gates, using Git/GitHub, or writing notes. Read only the references it routes for the task; their rules are binding. If the provider lacks skills, read its `SKILL.md` from `~/sideprojects/environment-as-code/plugins/agent-workflow/skills/agent-workflow/`.
- For tasks that require controlling or manually testing a visible macOS app, use the `peekaboo` skill. Prefer a programmatic interface when it covers the task.
- Keep shell commands plain so permission allowlists match them: no `VAR=` prefixes, `env`/`timeout` wrappers, or redirects when a flag, config file or tool option does the same. For another directory, use the tool's working-directory parameter, else `cd <dir> && git …`, not `git -C <dir>`.
- In a sandbox (Codex), run commands sandboxed first. Request escalation only after a sandbox denial or for network access, and escalate one plain command at a time.

## Non-negotiable safety

- Work outside the session repository only in a fresh worktree created from the fetched default branch. Never edit another repository's main checkout.
- Never bypass hooks, signatures, reviews, permissions, linters, or failing checks. Fix the root cause.
- If a klaudiush hook fails and the failure looks like a hook bug (wrong rule, false positive, crash), stop and report it to me with the command, hook output, and why it looks wrong. Do not work around it.
- When running inside the Sail harness, report every potential bug you hit in it (UI, session, tool, or permission misbehavior) to me with repro steps and observed vs. expected behavior, so I can act on it. Continue the task unless blocked.
- Never add linter-disable directives to hide an error.
- Never use destructive Git commands unless explicitly requested.
- Restrict cluster testing to local clusters unless explicitly authorized to use remote infrastructure.
- Treat existing processes, worktrees, and test environments as owned by others unless this task created them.

## Continuous improvement

- Same instruction 3+ times, or a recurring error pattern → suggest a project instruction rule: rule + one-line rationale + expected impact.
