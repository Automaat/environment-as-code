#!/usr/bin/env bash
# Install direct symlinks for the agent-instruction configs (Claude, Codex,
# Copilot, opencode). zakwas doesn't manage these: agents write next to them, so
# they stay plain symlinks. Edit the sources in this repo and re-run.
# Idempotent; backs up any real (non-symlink) file it replaces.
set -euo pipefail

src_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repo_dir="$(cd "$src_dir/../.." && pwd)"

# Regenerate the flat, self-contained AGENTS.md from CLAUDE.md + rules/.
"$src_dir/build-agents.sh"

link() { # link <target> <linkpath>
  local target="$1" path="$2"
  local backup="$path.pre-zakwas.bak"
  mkdir -p "$(dirname "$path")"
  if [ -e "$path" ] && [ ! -L "$path" ]; then
    mv "$path" "$backup"
    echo "backed up $path -> $backup"
  fi
  ln -sfn "$target" "$path"
  echo "linked $path -> $target"
}

# Skills and commands now come from plugins (zakwas `agents`), so this script
# links nothing into those dirs any more. It removes its old links there, live
# or dangling, since agents would load them next to the plugin copies. Only
# links into this repo go: plugins and other repos link into the same dirs.
prune_own_links() { # prune_own_links <dir>
  local dir="$1" existing
  [ -d "$dir" ] || return 0
  for existing in "$dir"/*; do
    [ -L "$existing" ] || continue
    case "$(readlink "$existing")" in
      "$repo_dir"/*) ;;
      *) continue ;;
    esac
    rm -f "$existing"
    echo "pruned $existing"
  done
}

# Claude (global) — modular; Claude inlines the rules/ links itself.
link "$src_dir/CLAUDE.md" "$HOME/.claude/CLAUDE.md"
link "$src_dir/rules" "$HOME/.claude/rules"
link "$src_dir/statusline-command.sh" "$HOME/.claude/statusline-command.sh"

prune_own_links "$HOME/.claude/skills"
prune_own_links "$HOME/.claude/commands"

# Codex (global) — flat AGENTS.md; Codex does not follow markdown links.
link "$src_dir/AGENTS.md" "$HOME/.codex/AGENTS.md"

# Copilot CLI (global) — documented global instruction directory.
link "$src_dir/AGENTS.md" "$HOME/.copilot/instructions/global.instructions.md"

# Copilot CLI compatibility — older/global parent walk-up setups.
link "$src_dir/AGENTS.md" "$HOME/AGENTS.md"

# opencode (global) — flat AGENTS.md. Linked explicitly rather than relying on
# opencode's ~/.claude/CLAUDE.md fallback, which loses the rules/ links.
link "$src_dir/AGENTS.md" "$HOME/.config/opencode/AGENTS.md"

prune_own_links "$HOME/.config/opencode/commands"

echo "done."
