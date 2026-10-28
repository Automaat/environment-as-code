# Agent context budget

Measured as UTF-8 bytes of instruction Markdown on 2026-10-07. The previous global `AGENTS.md` was 13,909 bytes.

| Representative session | Loaded guidance | Bytes | Reduction |
|---|---|---:|---:|
| General conversation | global core | 1,702 | 87.8% |
| Code change in current repository | core + skill router + quality | 4,201 | 69.8% |
| GitHub work in current repository | core + skill router + Git/GitHub | 5,170 | 62.8% |
| Cross-repository delivery | core + skill router + all references | 15,700 | -12.9% |

The last case intentionally pays for the complete policy because it needs repository discovery, worktree rules, implementation quality gates, and GitHub delivery. Other sessions avoid the 9,167-byte repository inventory and unrelated workflow sections.

Reproduce from the repository root:

```bash
wc -c dotfiles/agents/AGENTS.md
wc -c plugins/agent-workflow/skills/agent-workflow/SKILL.md
wc -c plugins/agent-workflow/skills/agent-workflow/references/*.md
```
