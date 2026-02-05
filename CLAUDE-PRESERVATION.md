# Claude Code Configuration Preservation

## What We Preserve (Already Tracked in Nix)

### Core Config - `modules/dotfiles/claude.nix`

#### 1. Settings (`~/.claude/settings.json`)
Generated via nix, includes:
- Permissions (allow list for Bash, WebSearch, WebFetch, etc.)
- Hooks (klaudiush PreToolUse + Notification)
- Status line command
- Environment variables
- Model selection (opus)
- alwaysThinkingEnabled

#### 2. Status Line Script (`~/.claude/statusline-command.sh`)
Custom bash script showing:
- Project name (github.com/org/repo format)
- Git branch (colored output)
- Current file context

#### 3. Global Instructions (`~/.claude/CLAUDE.md`)
Source: `dotfiles/claude/CLAUDE.md`
- Core principles (concise, no placeholders)
- Git workflow (sign commits, hooks compliance)
- Coding standards
- Communication style
- Project-specific patterns

#### 4. Commands (3/6 tracked)
**Tracked:**
- `cleanup-code.md` - Pre-commit cleanup
- `docs-review.md` - Documentation PR review
- `pr-review.md` - Automated PR review

**NOT Tracked:**
- `fix-review.md` - Apply unresolved PR review comments
- `go-review.md` - Manual Go code review vs 100go.co
- `merge-renovate.md` - Merge Renovate PRs across repos

#### 5. Skills (1/3 tracked)
**Tracked:**
- `go-code-review/` - Auto-review Go code for 100+ mistakes
  - SKILL.md
  - knowledge-base.md
  - real-world-patterns.md

**NOT Tracked:**
- `claude-md-gen/` - Generate project-specific CLAUDE.md
  - SKILL.md (18KB)
  - checklist.md, customization-guide.md
  - patterns/ directory (8 files)
  - questions/ directory (9 files)
  - templates/ directory (9 files)
- `obsidian-inbox-cleanup/` - Interactive inbox cleanup
  - SKILL.md (5.5KB)

## What We DON'T Preserve (Ephemeral/Generated)

### Runtime State - Don't Track
1. **history.jsonl** (3MB) - Conversation history
2. **session-env/** (908 dirs) - Session environments
3. **file-history/** (835 dirs) - File version history
4. **shell-snapshots/** (356 dirs) - Shell state snapshots
5. **todos/** (5319 dirs) - Task tracking data
6. **debug/** (2580 dirs) - Debug logs
7. **plans/** (347 files) - Generated plans
8. **paste-cache/** (37 files) - Clipboard cache
9. **tasks/** (46 files) - Task outputs
10. **telemetry/** - Usage telemetry
11. **usage-data/** - Usage statistics
12. **stats-cache.json** (12KB) - Statistics cache
13. **statsig/** - Feature flags cache
14. **plugins/** (7 files) - Plugin state
15. **ide/** (6 files) - IDE integration state
16. **hooks/dispatcher.log** (7MB) - Hook execution logs
17. **cache/changelog.md** - Cached changelog

### Project-Specific Memory - Conditional
**~/.claude/projects/** (36 project directories)
- Each project may have MEMORY.md with learnings
- These are **machine-independent** but **project-specific**
- Consider backing up separately if valuable
- Will regenerate on new machine as you work

**Decision**: Don't track in nix-darwin (machine config)
**Alternative**: Backup manually or via git if projects tracked

### Local Settings Override
**settings.local.json** (387 bytes)
- Machine-specific permission overrides
- Example: additional Bash permissions (cp, rm, git add/commit/push, etc.)
- Should be **machine-specific**, don't track

## What to Add (Optional)

### Missing Commands (3)
If actively used, add to `dotfiles/claude/commands/`:

1. **fix-review.md** - Apply PR review fixes
   - Uses EnterPlanMode for research
   - Fetches unresolved threads via gh CLI
   - Proposes + implements fixes

2. **go-review.md** - Manual Go review
   - 100go.co best practices
   - Works with files, patterns, or PR URLs
   - Focused on changed lines only

3. **merge-renovate.md** - Renovate PR merger
   - Discovers all ~/sideprojects repos
   - Lists open Renovate PRs
   - Checks CI, merges if passing
   - Supports --dry-run

### Missing Skills (2)
If actively used, add to `dotfiles/claude/skills/`:

1. **claude-md-gen/** - CLAUDE.md generator
   - Full directory structure (patterns, questions, templates)
   - ~25 files total
   - Generates project-specific instructions

2. **obsidian-inbox-cleanup/** - Inbox cleanup
   - Single SKILL.md file
   - Interactive PARA categorization
   - AI-optimized formatting

## Implementation Plan

### To Add Missing Commands:
```bash
# Add to dotfiles/claude/commands/
cp ~/.claude/commands/fix-review.md dotfiles/claude/commands/
cp ~/.claude/commands/go-review.md dotfiles/claude/commands/
cp ~/.claude/commands/merge-renovate.md dotfiles/claude/commands/

# Update modules/dotfiles/claude.nix:
home.file.".claude/commands/fix-review.md".source = ../../dotfiles/claude/commands/fix-review.md;
home.file.".claude/commands/go-review.md".source = ../../dotfiles/claude/commands/go-review.md;
home.file.".claude/commands/merge-renovate.md".source = ../../dotfiles/claude/commands/merge-renovate.md;
```

### To Add Missing Skills:
```bash
# Copy full skill directories
cp -r ~/.claude/skills/claude-md-gen dotfiles/claude/skills/
cp -r ~/.claude/skills/obsidian-inbox-cleanup dotfiles/claude/skills/

# Update modules/dotfiles/claude.nix with all files
# Use wildcards or explicit paths for subdirectories
```

## Migration Strategy

### New Machine Setup:
1. **Nix applies base config** → settings.json, CLAUDE.md, tracked commands/skills
2. **Create settings.local.json** → machine-specific permissions
3. **Project memory regenerates** → as you work on projects
4. **Runtime state rebuilds** → history, sessions, tasks naturally accumulate

### What to Backup Separately:
1. **Project MEMORY.md files** (if valuable learnings)
   - Find: `find ~/.claude/projects -name "MEMORY.md"`
   - Backup to Obsidian or separate repo
2. **settings.local.json** (optional, for reference)
3. **Custom commands/skills not tracked** (if exist)

## Current Coverage

**Tracked (via nix):**
- ✅ Core settings.json (permissions, hooks, statusline)
- ✅ CLAUDE.md (global instructions)
- ✅ 3/6 commands (50%)
- ✅ 1/3 skills (33%)
- ✅ Statusline script

**Not Tracked:**
- ❌ 3 commands (fix-review, go-review, merge-renovate)
- ❌ 2 skills (claude-md-gen, obsidian-inbox-cleanup)
- ❌ settings.local.json (machine-specific, intentional)
- ❌ Runtime state (ephemeral, intentional)
- ❌ Project memory (project-specific, intentional)

## Recommendation

**High Priority - Add to Nix:**
- `fix-review.md`, `go-review.md`, `merge-renovate.md` commands (if actively used)
- Skills only if frequently used (can reinstall via skill marketplace)

**Low Priority:**
- Project MEMORY.md files → backup separately if valuable
- settings.local.json → document pattern, regenerate per machine

**Don't Track:**
- Runtime state (history, sessions, etc.) → regenerates naturally
- Logs → ephemeral debugging data
