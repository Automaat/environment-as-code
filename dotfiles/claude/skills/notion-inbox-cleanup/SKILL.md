---
name: notion-inbox-cleanup
description: Interactive inbox cleanup with PARA categorization and AI-optimized formatting
argument-hint: "[--dry-run] [--page page-id]"
allowed-tools: Read, Write, Bash, AskUserQuestion, ToolSearch, mcp__notion__notion-search, mcp__notion__notion-fetch, mcp__notion__notion-update-page, mcp__notion__notion-move-pages
user-invocable: true
---

# Notion Inbox Cleanup

Interactive review and categorization of inbox pages into PARA structure with AI-optimized formatting.

**Workspace:** Connected via Notion MCP

---

## Arguments

Parse from `$ARGUMENTS`:

- **--dry-run:** Optional — Show suggestions without moving pages
- **--page:** Optional — Process single page ID instead of full inbox

---

## Workflow

### Phase 1: Setup

1. Load Notion MCP tools using ToolSearch
2. Search workspace for pages in inbox (search for pages tagged with inbox or in inbox parent)
3. Report count: "Found X pages to process"

### Phase 2: Interactive Review

For each page in inbox:

#### 2a. Read & Analyze

- Fetch page content using `notion-fetch`
- Identify content type and key topics
- Check for special patterns (see Routing Rules)

#### 2b. Present Suggestion

Show to user:

```markdown
## 📝 [page-title]

**Content preview:** [first 200 chars or summary]
**Detected type:** [recipe/work/ai-research/travel/general]
**Suggested destination:** [PARA category]
**Suggested tags:** [#tag1, #tag2]
**Confidence:** [high/medium/low]

Move to suggested location?
```

Use AskUserQuestion with options:
- **Yes** — Move as suggested
- **Different location** — Let me specify destination
- **Skip** — Leave in inbox for now
- **Edit first** — Review before moving

#### 2c. Apply Decision

If approved:
1. Update page properties (see Properties Template)
2. Move page to destination using `notion-move-pages`
3. Log action

If "Different location":
1. Ask for destination
2. Suggest appropriate tags for that location
3. Move with updated properties

### Phase 3: Summary

After all pages processed:

```markdown
## 📊 Inbox Cleanup Summary

**Processed:** X pages
**Moved:** Y pages
**Skipped:** Z pages

### Actions Taken
- [page-title] → [destination]
- ...

### Remaining in Inbox
- [page-title] (reason: skipped/unclear)
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
| Spanish words, translations | Link to `2_Areas/Personal/Spanish_learning` page | `#area/languages` |
| `ai-digest-*` | `3_Resources/AI/Digests/` | `#resource/ai` |

### PARA Decision Framework

1. **Has deadline/goal?** → `1_Projects/`
2. **Ongoing responsibility?** → `2_Areas/`
3. **Reference/learning?** → `3_Resources/`
4. **Completed/inactive?** → `4_Archive/`

---

## Properties Template

Update all processed pages (AI-optimized):

Properties to add/update:
- **Tags:** PARA tag + topic tags (max 5)
- **Type:** note | moc | project | literature
- **Summary:** One-sentence TL;DR for AI context (rich text property)
- **Date:** Current date
- **Source:** URL if applicable
- **Related:** Links to related pages

**Rules:**
- `Tags`: Use multi-select property with PARA tag + topic tags
- `Type`: Use select property (usually `note` unless clearly MOC or project)
- `Summary`: CRITICAL — write concise 1-sentence summary
- `Source`: URL property if note references external content
- `Related`: Relation property linking to related pages

---

## Special Processing

### Recipe Pages

1. Ensure proper formatting with ingredients and instructions
2. Add cooking tags (#vegetarian, #quick, etc.)
3. Move to `3_Resources/Cooking/`

### Single-Link Pages

1. Follow the link using WebFetch if needed
2. If article → create concise summary
3. If not article → create short description
4. Add source in properties

### Travel Pages

Route to `3_Resources/Travel/[Country]/`:
- Create or link to country-specific pages
- Categorize by type (attractions, food, accommodation, etc.)

---

## Error Handling

- **Page not found:** Skip, report in summary
- **Permission error:** Report, continue with next
- **Unclear categorization:** Ask user with AskUserQuestion
- **Duplicate exists:** Ask: merge, keep both, or skip

---

## Quality Checklist

Before completing each page:

- [ ] Properties include `Summary` field (critical for AI)
- [ ] Tags follow `#para-type/topic` pattern
- [ ] Wikilinks added to related pages if obvious
- [ ] No placeholder text in content
- [ ] Page moved to correct PARA location

---

## AI-Friendly Formatting

When processing pages, apply these optimizations:

1. **Section summaries** — Add brief summary after major headings
2. **Bullet points** — Use bullets for better readability
3. **Liberal linking** — Add page links to concepts that exist as pages
4. **Flat structure** — Avoid deep nesting (max 3 levels)
5. **Explicit connections** — Link to related pages
