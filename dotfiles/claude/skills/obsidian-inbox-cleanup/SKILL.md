---
name: obsidian-inbox-cleanup
description: Interactive inbox cleanup with PARA categorization and AI-optimized formatting
argument-hint: "[--dry-run] [--file filename]"
allowed-tools: Read, Write, Bash, AskUserQuestion
user-invocable: true
---

# Obsidian Inbox Cleanup

Interactive review and categorization of inbox notes into PARA structure with AI-optimized frontmatter.

**Vault:** `/Users/marcin.skalski@konghq.com/Library/Mobile Documents/iCloud~md~obsidian/Documents/second-brain/`

---

## Arguments

Parse from `$ARGUMENTS`:

- **--dry-run:** Optional — Show suggestions without moving files
- **--file:** Optional — Process single file instead of full inbox

---

## Workflow

### Phase 1: Setup

1. Read vault CLAUDE.md for current PARA rules
2. List all files in `0_Inbox/` (exclude `.last-run`, `Tasks.md`, directories)
3. Report count: "Found X notes to process"

### Phase 2: Interactive Review

For each note in inbox:

#### 2a. Read & Analyze

- Read full content
- Identify content type and key topics
- Check for special patterns (see Routing Rules)

#### 2b. Present Suggestion

Show to user:

```markdown
## 📝 [filename]

**Content preview:** [first 200 chars or summary]
**Detected type:** [recipe/work/ai-research/travel/general]
**Suggested destination:** [full path]
**Suggested tags:** [#tag1, #tag2]
**Confidence:** [high/medium/low]

Move to suggested location?
```

Use AskUserQuestion with options:
- **Yes** — Move as suggested
- **Different folder** — Let me specify path
- **Skip** — Leave in inbox for now
- **Edit first** — Open for review before moving

#### 2c. Apply Decision

If approved:
1. Add/update YAML frontmatter (see Frontmatter Template)
2. Move file to destination
3. Log action

If "Different folder":
1. Ask for destination path
2. Suggest appropriate tags for that location
3. Move with updated frontmatter

### Phase 3: Summary

After all notes processed:

```markdown
## 📊 Inbox Cleanup Summary

**Processed:** X notes
**Moved:** Y notes
**Skipped:** Z notes

### Actions Taken
- [filename] → [destination]
- ...

### Remaining in Inbox
- [filename] (reason: skipped/unclear)
```

---

## Routing Rules

### Auto-Detection Patterns

| Pattern | Destination | Tags |
|---------|-------------|------|
| `przepis-*`, recipe keywords, ingredients list | `3_Resources/Cooking/` | `#resource/cooking` |
| `kuma-*`, Kong, mesh, dataplane | `1_Projects/0_Work/` | `#project/0_work` |
| LLM, AI, Claude, GPT, embeddings, RAG | `3_Resources/AI/` | `#resource/ai` |
| `plant-*`, gardening, watering | `3_Resources/Plants/` | `#resource/plants` |
| travel, country names, itinerary | `3_Resources/Travel/[Country]/` | `#resource/travel` |
| Spanish words, translations | Append to `2_Areas/Personal/Spanish_learning.md` | `#area/languages` |
| `.draftsExport` files | Process per CLAUDE.md rules | varies |
| `ai-digest-*` | `0_Inbox/ai-digest/` or `3_Resources/AI/Digests/` | `#resource/ai` |

### PARA Decision Framework

1. **Has deadline/goal?** → `1_Projects/`
2. **Ongoing responsibility?** → `2_Areas/`
3. **Reference/learning?** → `3_Resources/`
4. **Completed/inactive?** → `4_Archive/`

---

## Frontmatter Template

Add to all processed notes (AI-optimized):

```yaml
---
title: Descriptive Plain-Language Title
date: YYYY-MM-DD
tags: [para-tag, topic-tags]
type: note | moc | project | literature
summary: One-sentence TL;DR for AI context
source: URL or [[Internal Link]] if applicable
related: [[Note1]], [[Note2]]
---
```

**Rules:**
- `title`: Extract from content or filename (human readable)
- `date`: Use file modified date or today
- `tags`: PARA tag + topic tags (max 5)
- `type`: Usually `note` unless clearly MOC or project
- `summary`: CRITICAL — write concise 1-sentence summary
- `source`: Include if note references external content
- `related`: Add if obvious connections exist

---

## Special Processing

### .draftsExport Files

Per vault CLAUDE.md:
1. Extract only content field (ignore metadata)
2. Check for actionable tasks → add to `0_Inbox/Tasks.md`
3. Format URLs as markdown links
4. Process each extracted note individually

### Single-Link Notes

1. Follow the link
2. If article → create concise summary
3. If not article → create short description
4. Add source in References section

### Travel Notes

Route to `3_Resources/Travel/[Country]/[category].md`:
- Attractions → `attractions.md`
- Restaurants/food → `food.md`
- Hotels → `accommodation.md`
- Getting around → `transportation.md`
- Practical tips → `practical_info.md`

---

## Error Handling

- **File not found:** Skip, report in summary
- **Permission error:** Report, continue with next
- **Unclear categorization:** Ask user with AskUserQuestion
- **Duplicate exists:** Ask: overwrite, rename, or skip

---

## Quality Checklist

Before completing each note:

- [ ] Frontmatter has `summary` field (critical for AI)
- [ ] Tags follow `#para-type/topic` pattern
- [ ] Wikilinks added to related notes if obvious
- [ ] No placeholder text in content
- [ ] File moved to correct PARA location

---

## AI-Friendly Formatting

When processing notes, apply these optimizations:

1. **Section summaries** — Add brief summary after major headings
2. **Bullet points** — Convert tables to bullets where appropriate
3. **Liberal linking** — Add `[[wikilinks]]` to concepts that exist as notes
4. **Flat structure** — Avoid deep nesting (max 3 levels)
5. **Explicit connections** — Add "Related" section at bottom
