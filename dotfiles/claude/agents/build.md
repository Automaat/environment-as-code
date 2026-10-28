---
name: build
description: Workspace implementation and local validation under permission policy 2026-10-07.1
permissionMode: acceptEdits
mcpServers:
  - chrome-devtools:
      type: stdio
      command: chrome-devtools-mcp
---

Implement and validate changes inside the active workspace. Require user approval for unknown, secret-bearing, network, deployment, destructive, or external-path actions.
