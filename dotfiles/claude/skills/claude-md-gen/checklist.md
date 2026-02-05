# CLAUDE.md Generation Verification Checklist

Use this checklist before finalizing generated CLAUDE.md.

---

## Technical Accuracy

- [ ] Language versions match package files exactly
- [ ] Framework versions correct and match dependencies
- [ ] Directory paths exist and are accurate
- [ ] Commands are executable (tested with --help at minimum)
- [ ] Package manager commands correct (npm/yarn/poetry/cargo/etc.)
- [ ] Environment variables match actual configuration
- [ ] File paths use correct separators for project

## Actionability

- [ ] Workflows have concrete, numbered steps
- [ ] Commands include actual syntax (not "run tests" but "pytest tests/")
- [ ] Quality gates list actual tool names (not "linter" but "ruff")
- [ ] Output templates show actual structure (not "create template")
- [ ] Examples use real project code/data
- [ ] Error messages show actual formats
- [ ] No generic phrases like "configure appropriately"

## Specificity

- [ ] No "TODO" or "Fill this in" placeholders
- [ ] No "[Project Name]" or generic placeholders remaining
- [ ] Anti-patterns specific to this project (not just generic)
- [ ] Examples use actual project context
- [ ] Directory structure matches actual project
- [ ] Tech stack includes versions (not just "Python" but "Python 3.11")
- [ ] Workflows reference actual files/directories

## Completeness

- [ ] All major project areas have sections
- [ ] Selected patterns (2-4) included and adapted
- [ ] Quality gates comprehensive for project type
- [ ] Common commands cover frequent operations
- [ ] Anti-patterns address actual project issues
- [ ] Workflow sections cover top 2-3 tasks
- [ ] Testing approach documented (if tests exist)
- [ ] Deployment/build process documented (if applicable)

## Pattern Integration

- [ ] 2-4 patterns selected (not 0, not 6+)
- [ ] Patterns chosen based on user needs (not arbitrary)
- [ ] Patterns adapted to project (not copy-pasted verbatim)
- [ ] Pattern examples use project context
- [ ] Pattern placeholders replaced with project specifics
- [ ] Patterns integrated into natural sections (not appendix dump)

## User Answers Integration

- [ ] User workflow answers reflected in workflow sections
- [ ] User tool answers reflected in tool integration sections
- [ ] User pitfall answers in anti-patterns section
- [ ] User output format answers in output templates
- [ ] User quality gate answers in quality gates section
- [ ] All 5-7 questions answered and incorporated

## Format & Style

- [ ] Consistent markdown formatting
- [ ] Code blocks have language tags
- [ ] Headings use proper hierarchy (# → ## → ###)
- [ ] Lists formatted consistently
- [ ] Commands in code blocks (not inline)
- [ ] File paths in code formatting
- [ ] No broken markdown syntax

## Extensibility

- [ ] Customization guide referenced
- [ ] Clear how to add new sections
- [ ] Pattern library location noted
- [ ] Mentions skill location for updates
- [ ] Indicates CLAUDE.md can evolve with project

## Self-Contained

- [ ] No references to global ~/.claude/CLAUDE.md
- [ ] All instructions complete within file
- [ ] No "see team documentation" without link
- [ ] No "ask your colleague" placeholders
- [ ] Patterns fully explained (not just referenced)

## Quality of Content

- [ ] Workflows are realistic (actually used, not idealized)
- [ ] Anti-patterns are real pitfalls (not made up)
- [ ] Commands tested or verified (not guessed)
- [ ] Examples compilable/runnable (if code)
- [ ] No contradictions between sections
- [ ] Consistent terminology throughout

---

## Pre-Generation Checks

Run before starting generation:

- [ ] At least 3 files read to understand structure
- [ ] Package file(s) parsed correctly
- [ ] Directory structure mapped
- [ ] Existing conventions identified (linters, tests, hooks)
- [ ] 5-7 questions asked (not more, not less)
- [ ] Questions relevant to detected type
- [ ] Answers captured for customization

---

## Post-Generation Actions

After generation, before presenting to user:

1. **Verify all placeholders replaced:**
   ```bash
   grep -E '\[(PROJECT|LANGUAGE|FRAMEWORK|COMMAND|DIR)' CLAUDE.md
   # Should return no matches
   ```

2. **Check for TODOs:**
   ```bash
   grep -i 'TODO\|FIXME\|XXX' CLAUDE.md
   # Should return no matches
   ```

3. **Verify commands (sample check):**
   ```bash
   # Extract first command from "Common Commands" section
   # Run with --help or --version (non-destructive)
   # Verify exit code 0 or expected error
   ```

4. **Review summary:**
   - Count sections generated
   - Verify patterns listed match included patterns
   - Check "Next Steps" are actionable

---

## Red Flags (Block Generation)

Do NOT generate if:

- ❌ Unable to detect project type and user didn't specify
- ❌ No package files found and unable to infer tech stack
- ❌ Asked <3 questions (insufficient context)
- ❌ User answers contradict analysis (resolve first)
- ❌ Cannot determine primary language
- ❌ Template file missing for detected type

---

## Severity Levels

**Critical (must fix before generation):**
- Missing tech stack info
- Wrong project type detected
- Unresolved contradictions

**High (fix during generation):**
- Placeholder not replaced
- Command syntax incorrect
- Wrong directory paths

**Medium (note in summary):**
- Command not testable (mark for verification)
- Optional section missing
- Pattern adaptation could be better

**Low (user can fix):**
- Minor formatting inconsistencies
- Could add more examples
- Could expand certain sections

---

## Success Criteria

CLAUDE.md is ready when:

✅ User can read and immediately understand project
✅ Commands are copy-paste ready
✅ Workflows are step-by-step actionable
✅ Anti-patterns address real issues
✅ Quality gates match actual project tooling
✅ No generic/placeholder content remains
✅ Patterns adapted to project context
✅ Extensibility path clear

---

## Checklist for Different Project Types

### CLI Tools - Additional Checks

- [ ] All commands shown with actual flags
- [ ] Exit codes documented
- [ ] Error message formats shown
- [ ] Output format examples included

### Web Apps - Additional Checks

- [ ] Responsive breakpoints actual values
- [ ] Component examples use real components
- [ ] Build commands tested
- [ ] Deployment process documented

### APIs - Additional Checks

- [ ] Endpoint examples show real routes
- [ ] Response formats match actual API
- [ ] Auth flow documented with project method
- [ ] Migration commands verified

### Research - Additional Checks

- [ ] Note templates match actual usage
- [ ] Folder structure matches vault
- [ ] Processing steps numbered and clear
- [ ] Citation format examples shown

### Libraries - Additional Checks

- [ ] Public API examples are runnable
- [ ] Versioning policy clear
- [ ] Publishing commands verified
- [ ] Breaking change policy documented
