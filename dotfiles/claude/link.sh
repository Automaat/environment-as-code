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

link() { # link <target> <linkpath> [backup]
  local target="$1" path="$2"
  local backup="$path.pre-zakwas.bak"
  if [ "$#" -ge 3 ]; then backup="$3"; fi
  mkdir -p "$(dirname "$path")"
  if [ -e "$path" ] && [ ! -L "$path" ]; then
    mkdir -p "$(dirname "$backup")"
    mv "$path" "$backup"
    echo "backed up $path -> $backup"
  fi
  ln -sfn "$target" "$path"
  echo "linked $path -> $target"
}

# Agents still try to load a dangling link left by a renamed source. Only
# links into this repo go: plugins and other repos link into the same dirs,
# and their targets may just be missing for now (e.g. repo not cloned yet).
prune_dangling() { # prune_dangling <dir>
  local dir="$1" existing
  mkdir -p "$dir"
  for existing in "$dir"/*; do
    [ -L "$existing" ] || continue
    [ -e "$existing" ] && continue
    case "$(readlink "$existing")" in
      "$repo_dir"/*) ;;
      *) continue ;;
    esac
    rm -f "$existing"
    echo "pruned dangling $existing"
  done
}

# Claude (global) — modular; Claude inlines the rules/ links itself.
link "$src_dir/CLAUDE.md" "$HOME/.claude/CLAUDE.md"
link "$src_dir/rules" "$HOME/.claude/rules"
link "$src_dir/statusline-command.sh" "$HOME/.claude/statusline-command.sh"

# Claude skills (opencode reads them too) — linked one by one from the
# marketplace plugins; the directory also holds entries installed by plugins and
# other repos. Backups go outside it: agents would load a backed-up skill dir.
prune_dangling "$HOME/.claude/skills"
for skill in "$repo_dir"/plugins/*/skills/*/; do
  [ -d "$skill" ] || continue
  skill="${skill%/}"
  name="$(basename "$skill")"
  link "$skill" "$HOME/.claude/skills/$name" "$HOME/.claude/skills.pre-zakwas.bak/$name"
done
# No commands ship any more; this only prunes links to removed ones.
prune_dangling "$HOME/.claude/commands"

# Codex (global) — flat AGENTS.md; Codex does not follow markdown links.
link "$src_dir/AGENTS.md" "$HOME/.codex/AGENTS.md"

# Copilot CLI (global) — documented global instruction directory.
link "$src_dir/AGENTS.md" "$HOME/.copilot/instructions/global.instructions.md"

# Copilot CLI compatibility — older/global parent walk-up setups.
link "$src_dir/AGENTS.md" "$HOME/AGENTS.md"

# opencode (global) — flat AGENTS.md. Linked explicitly rather than relying on
# opencode's ~/.claude/CLAUDE.md fallback, which loses the rules/ links.
link "$src_dir/AGENTS.md" "$HOME/.config/opencode/AGENTS.md"

prune_dangling "$HOME/.config/opencode/commands"

echo "done."
