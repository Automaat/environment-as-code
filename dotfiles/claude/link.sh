#!/usr/bin/env bash
# Install direct symlinks for the agent-instruction configs (Claude, Codex,
# Copilot, opencode). eac doesn't manage these: agents write next to them, so
# they stay plain symlinks. Edit the sources in this repo and re-run.
# Idempotent; backs up any real (non-symlink) file it replaces.
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

# Agents still try to load a dangling link left by a renamed source.
prune_dangling() { # prune_dangling <dir>
  local dir="$1" existing
  mkdir -p "$dir"
  for existing in "$dir"/*; do
    [ -L "$existing" ] || continue
    [ -e "$existing" ] && continue
    rm -f "$existing"
    echo "pruned dangling $existing"
  done
}

# Claude (global) — modular; Claude inlines the rules/ links itself.
link "$src_dir/CLAUDE.md" "$HOME/.claude/CLAUDE.md"
link "$src_dir/rules" "$HOME/.claude/rules"
link "$src_dir/statusline-command.sh" "$HOME/.claude/statusline-command.sh"

# Claude skills and commands — linked one by one; both directories also hold
# entries installed by plugins and other repos.
prune_dangling "$HOME/.claude/skills"
for skill in "$src_dir"/skills/*/; do
  [ -d "$skill" ] || continue
  skill="${skill%/}"
  link "$skill" "$HOME/.claude/skills/$(basename "$skill")"
done
prune_dangling "$HOME/.claude/commands"
for cmd in "$src_dir"/commands/*.md; do
  [ -e "$cmd" ] || continue
  link "$cmd" "$HOME/.claude/commands/$(basename "$cmd")"
done

# Codex (global) — flat AGENTS.md; Codex does not follow markdown links.
link "$src_dir/AGENTS.md" "$HOME/.codex/AGENTS.md"

# Copilot CLI (global) — documented global instruction directory.
link "$src_dir/AGENTS.md" "$HOME/.copilot/instructions/global.instructions.md"

# Copilot CLI compatibility — older/global parent walk-up setups.
link "$src_dir/AGENTS.md" "$HOME/AGENTS.md"

# opencode (global) — flat AGENTS.md. Linked explicitly rather than relying on
# opencode's ~/.claude/CLAUDE.md fallback, which loses the rules/ links.
link "$src_dir/AGENTS.md" "$HOME/.config/opencode/AGENTS.md"

# opencode commands — opencode has no .claude/commands fallback.
oc_commands="$HOME/.config/opencode/commands"
prune_dangling "$oc_commands"
for cmd in "$src_dir"/commands/*.md; do
  [ -e "$cmd" ] || continue
  link "$cmd" "$oc_commands/$(basename "$cmd")"
done

echo "done."
