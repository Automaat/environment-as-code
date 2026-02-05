# Output Templates Pattern

Structured outputs with consistent sections and metadata.

---

## When to Use

- Standardized reports (analysis, reviews, research)
- CLI tool outputs (JSON, structured text)
- Documentation generation
- Note-taking systems
- Repeatable workflows

## Pattern Structure

### Components

1. **Header** - Metadata (date, type, version)
2. **Summary** - Quick overview (TL;DR)
3. **Body Sections** - Main content (structured)
4. **Metadata Footer** - Sources, tags, next steps
5. **Format Options** - Text, JSON, markdown, etc.

## Template

```markdown
# [Template Name]

**Purpose:** [What this template is for]
**When to use:** [Conditions for using this template]

## Structure

```
[FORMAT_NAME] Format:

[HEADER_SECTION]
---
[SUMMARY_SECTION]
---
[BODY_SECTION_1]
[BODY_SECTION_2]
[BODY_SECTION_3]
---
[METADATA_FOOTER]
```

## Placeholder Syntax

- `[PLACEHOLDER]` - Required field
- `[OPTIONAL_FIELD]` - Optional field
- `[LIST_ITEM]` - Repeating element
- `{dynamic|value}` - Alternative values
```

## Example 1: Plot Analysis Report (from plot-analysis skill)

```markdown
# 🏡 Plot Analysis Report

**Plot:** [Address or ID]
**Date:** [YYYY-MM-DD]
**Analyst:** [Name or AI]
**Status:** {✅ Recommended | ⚠️ Conditional | ❌ Not Recommended}

---

## 🎯 Executive Summary

**Recommendation:** [One sentence: accept/reject/investigate]

**Key Strengths:**
- [Strength 1]
- [Strength 2]
- [Strength 3]

**Key Concerns:**
- [Concern 1]
- [Concern 2]

**Score:** [X/20] points

---

## 📍 Location Analysis

**Address:** [Full address]
**Coordinates:** [Lat, Lon]
**Distance to Krakow center:** [X] minutes

### Accessibility
- Public transport: [Details]
- Roads: [Quality/access]
- Highways: [Proximity]

### Neighborhood
- Character: [Description]
- Noise level: [Assessment]
- Amenities: [Nearby facilities]

---

## 🏛️ Legal & Zoning

**Zoning designation:** [MN, RM, etc.]
**Building permit:** {Required | Not required}
**Restrictions:** [Any limitations]

### Ownership
- Title: {Clean | Disputed | Unknown}
- Encumbrances: [Any liens, easements]

---

## 🌍 Land Quality

**Size:** [X m²]
**Terrain:** {Flat | Sloped X%}
**Orientation:** [North/South/East/West facing]
**Soil:** [Type if known, or "needs testing"]
**Vegetation:** [Description]

---

## 🔌 Utilities

| Utility | Available | Connection Cost | Notes |
|---------|-----------|----------------|-------|
| Electric | {✅|❌} | [Est. cost] | [Details] |
| Water | {✅|❌} | [Est. cost] | [Details] |
| Sewage | {✅|❌} | [Est. cost] | [Details] |
| Gas | {✅|❌} | [Est. cost] | [Details] |
| Internet | {✅|❌} | [Est. cost] | [Details] |

---

## 💰 Cost Analysis

**Purchase price:** [X PLN]
**Utility connections:** [X PLN]
**Est. permit/fees:** [X PLN]
**Total initial cost:** [X PLN]

**Compared to market:** {Below | At | Above} market by [X%]

---

## ⚠️ Risk Factors

- [Risk 1 with severity: Low/Medium/High]
- [Risk 2 with severity]
- [Risk 3 with severity]

---

## 📋 Next Steps

1. [Action item 1]
2. [Action item 2]
3. [Action item 3]

**Recommended timeline:** [Time estimate]
**Budget for next steps:** [Cost estimate]

---

## 📚 Sources

- [KB: file.ext:line] - [Description]
- [Listing: Title](URL) - [Description]
- [Source: Title](URL) - [Description]

**Tags:** #plot-analysis #location #status
**Related:** [[Other Analysis]], [[Comparison]]
```

## Example 2: CLI Tool JSON Output

```json
{
  "version": "1.0",
  "timestamp": "2024-01-27T12:00:00Z",
  "command": "[tool] [subcommand] [args]",
  "status": "success|error|warning",
  "data": {
    "[primary_field]": "[value]",
    "[secondary_field]": "[value]",
    "[list_field]": [
      {
        "id": "[id]",
        "name": "[name]",
        "value": "[value]"
      }
    ]
  },
  "metadata": {
    "duration_ms": 1234,
    "items_processed": 42,
    "warnings": [
      "[warning message 1]",
      "[warning message 2]"
    ]
  },
  "errors": [
    {
      "code": "[error_code]",
      "message": "[user-facing message]",
      "details": "[technical details]"
    }
  ]
}
```

## Example 3: Research Note (Obsidian-style)

```markdown
# [Topic Name]

**Source:** [Source Title](URL) or [[Internal Note]]
**Date:** YYYY-MM-DD
**Tags:** #research #domain #subdomain
**Type:** {Research | Technical | Meeting | Literature}

---

## 📝 Summary

[2-3 sentence overview of the topic/source]

---

## 🔑 Key Points

- **[Point 1]:** [Explanation with details]
- **[Point 2]:** [Explanation with details]
- **[Point 3]:** [Explanation with details]

---

## 💡 Personal Insights

[Your analysis, connections to other knowledge, implications]

**Connections:**
- Links to [[Related Concept 1]]
- Contrasts with [[Related Concept 2]]
- Extends [[Related Concept 3]]

---

## ❓ Questions

- [Question 1 raised by this material]
- [Question 2 to explore further]

---

## 📖 References

- [Source 1](URL)
- [[Internal Note]]
- [Source 2](URL)

---

**Related MOCs:** [[Topic MOC]], [[Domain MOC]]
**Next Steps:** [[Follow-up Note]], [[Exploration Area]]
```

## Example 4: Go Code Review Output

```markdown
# Go Code Review: [Package/File]

**Date:** YYYY-MM-DD
**Files Reviewed:** [Count] files, [LOC] lines
**Severity:** {🔴 Critical | 🟡 Medium | 🟢 Low}

---

## ⚠️ Critical Issues (Must Fix)

### [Issue 1 Title]

**Location:** file.go:line
**Category:** {Race Condition | Memory Leak | Security | ...}
**Severity:** Critical

**Problem:**
[Description of the issue]

**Code:**
```go
[Problematic code snippet]
```

**Recommendation:**
[How to fix it]

**Fixed Code:**
```go
[Corrected code snippet]
```

**Reference:** [Doc: Go Best Practice](URL)

---

## 🔍 Medium Issues (Should Fix)

### [Issue 2 Title]

**Location:** file.go:line
**Category:** {Performance | Error Handling | API Design}
**Severity:** Medium

[Same structure as critical issues]

---

## 💡 Suggestions (Consider)

- [Suggestion 1] - [file.go:line]
- [Suggestion 2] - [file.go:line]

---

## ✅ Positive Patterns

- [Good pattern 1] in file.go:line
- [Good pattern 2] in file.go:line

---

## 📊 Summary

**Total Issues:** [X critical, Y medium, Z low]
**Recommendation:** {Block merge | Fix critical first | Approve with suggestions}

**Estimated Fix Time:** [Time estimate]

---

## 📚 Resources

- handlers/auth.go:45 - [Issue location]
- [Doc: Go Code Review Guide](URL)
- [KB: CONTRIBUTING.md] - Project standards
```

## Example 5: API Test Results

```json
{
  "test_suite": "API Integration Tests",
  "timestamp": "2024-01-27T12:00:00Z",
  "environment": "staging",
  "summary": {
    "total": 45,
    "passed": 42,
    "failed": 2,
    "skipped": 1,
    "duration_ms": 3456
  },
  "failures": [
    {
      "test": "POST /api/users - should reject invalid email",
      "endpoint": "POST /api/users",
      "expected": 400,
      "actual": 200,
      "error": "Validation not triggered",
      "location": "tests/users.test.ts:78"
    }
  ],
  "performance": {
    "slowest_tests": [
      {
        "test": "GET /api/reports - generate large report",
        "duration_ms": 2345,
        "threshold_ms": 1000,
        "status": "warning"
      }
    ]
  },
  "coverage": {
    "endpoints": 23,
    "tested": 21,
    "percentage": 91.3
  }
}
```

## Customization Points

**Adjust sections:**
- Add domain-specific sections
- Remove irrelevant sections
- Reorder by priority
- Nest subsections for depth

**Adjust metadata:**
- More/fewer fields in header
- Different footer elements
- Versioning approach
- Timestamps and tracking

**Adjust format:**
- Markdown for human reading
- JSON for machine processing
- YAML for configuration
- Plain text for simplicity

**Adjust emoji usage:**
- Heavy emoji (research notes, visual scanning)
- Light emoji (section markers only)
- No emoji (formal/technical docs)

**Adjust granularity:**
- Executive: summary only
- Standard: summary + key sections
- Detailed: all sections with examples

## Integration with Project Type

**CLI tools:** JSON outputs, error reports, status displays
**Web apps:** Component documentation, design specs
**API services:** API docs, test reports, monitoring alerts
**Research:** Note templates, literature reviews, analysis reports
**Libraries:** API documentation, example templates, release notes

## Benefits

- Consistency across outputs
- Easy to parse (humans and machines)
- Completeness (no missed sections)
- Navigability (clear structure)
- Professionalism (polished appearance)

## Anti-Patterns

**AVOID:**
- ❌ Skipping sections (breaks consistency)
- ❌ Inconsistent placeholder syntax
- ❌ Too many sections (cognitive overload)
- ❌ No examples (hard to understand)
- ❌ Rigid templates (allow customization)
- ❌ Missing metadata (can't track or verify)
