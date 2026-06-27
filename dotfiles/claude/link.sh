#!/usr/bin/env bash
# Install direct symlinks for the agent-instruction configs (Claude, Codex,
# Copilot). These are intentionally NOT nix-managed — edit the sources in this
# repo and re-run. Idempotent; backs up any real (non-symlink) file it replaces.
set -euo pipefail

src_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Regenerate the flat, self-contained AGENTS.md from CLAUDE.md + rules/.
"$src_dir/build-agents.sh"

link() { # link <target> <linkpath>
  local target="$1" path="$2"
  mkdir -p "$(dirname "$path")"
  if [ -e "$path" ] && [ ! -L "$path" ]; then
    mv "$path" "$path.pre-eac.bak"
    echo "backed up $path -> $path.pre-eac.bak"
  fi
  ln -sfn "$target" "$path"
  echo "linked $path -> $target"
}

# Claude (global) — modular; Claude inlines the rules/ links itself.
link "$src_dir/CLAUDE.md" "$HOME/.claude/CLAUDE.md"
link "$src_dir/rules" "$HOME/.claude/rules"
link "$src_dir/statusline-command.sh" "$HOME/.claude/statusline-command.sh"

# Codex (global) — flat AGENTS.md; Codex does not follow markdown links.
link "$src_dir/AGENTS.md" "$HOME/.codex/AGENTS.md"

# Copilot CLI (global) — documented global instruction directory.
link "$src_dir/AGENTS.md" "$HOME/.copilot/instructions/global.instructions.md"

# Copilot CLI compatibility — older/global parent walk-up setups.
link "$src_dir/AGENTS.md" "$HOME/AGENTS.md"

echo "done."
