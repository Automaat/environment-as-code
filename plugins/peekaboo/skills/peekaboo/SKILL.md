---
name: peekaboo
description: Inspect and operate visible native macOS apps through Peekaboo when a task requires GUI interaction or manual app testing. Use for screenshots, Accessibility inspection, clicks, typing, menus, and dialogs; prefer programmatic interfaces when they cover the task.
---

# Peekaboo

The `peekaboo` CLI and MCP server are pinned by mise in environment-as-code. Check `peekaboo --version` and `peekaboo permissions status` before live work. Screen capture needs Screen Recording permission; UI control needs Accessibility permission on the process hosting Peekaboo.

Observe the exact app or window with `peekaboo window list --app APP --json` and `peekaboo see --app APP --json`. Prefer Accessibility element IDs and actions over coordinates. Reobserve after each UI change before reusing an element or snapshot ID. Verify the resulting app state; successful input dispatch alone does not prove the expected behavior.

Use the `peekaboo` MCP tools when available. For current CLI syntax, use `peekaboo help COMMAND`; for tool schemas, use `peekaboo tools --json`. Follow the upstream [Peekaboo documentation](https://github.com/openclaw/Peekaboo/tree/main/docs) for detailed commands and permission troubleshooting.
