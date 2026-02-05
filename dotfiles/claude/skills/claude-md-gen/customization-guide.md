# CLAUDE.md Customization Guide

How to extend and adapt the generated CLAUDE.md as your project evolves.

---

## Philosophy

Generated CLAUDE.md is a **starting point**, not a final artifact. As your project evolves:

- Add sections for new patterns
- Refine workflows based on actual usage
- Update anti-patterns as you discover them
- Expand templates for new output types
- Document new tools and integrations

**Goal:** Keep CLAUDE.md synchronized with how you actually work.

---

## Adding New Sections

### When to Add a Section

Add when you:
- Introduce new workflow (e.g., deployment, releases)
- Adopt new tool (e.g., monitoring, analytics)
- Establish new pattern (e.g., caching strategy)
- Document new constraint (e.g., performance budgets)

### Section Structure Template

```markdown
## [Section Name]

**Purpose:** [Why this section exists]

[Main content - concrete and actionable]

### [Subsection if needed]

[Details]

**Example:**
```
[Code or command example]
```
```

### Where to Place New Sections

**Guideline:** Place near related content

- **Workflow sections:** After "Development Workflow"
- **Tool integrations:** After "Quality Gates" or "Common Commands"
- **Patterns:** Before "Anti-Patterns"
- **Templates:** In "Output Templates" or after relevant workflow
- **Quality/standards:** Near "Quality Gates"

---

## Updating Existing Sections

### Workflows

**When workflows change:**

1. Update step-by-step instructions
2. Add/remove steps as needed
3. Update related commands
4. Adjust time estimates if included
5. Add examples if helpful

**Example addition:**
```markdown
### Adding Feature with Feature Flag

1. Create flag in `config/flags.go`
2. Implement feature behind flag check
3. Add tests with flag enabled/disabled
4. Deploy to staging with flag off
5. Enable flag for internal testing
6. Monitor metrics before production rollout
```

### Quality Gates

**When gates change:**

1. Add new checks to checklist
2. Update commands
3. Adjust pass/fail criteria
4. Document new tools/versions

**Example:**
```markdown
- [ ] golangci-lint passes (added 2024-01)
- [ ] gosec security scan passes (added 2024-02)
```

### Anti-Patterns

**As you discover pitfalls:**

1. Add to "AVOID" list
2. Explain why problematic
3. Link to incident/issue if applicable
4. Show correct approach

**Example:**
```markdown
- ❌ Storing tokens in localStorage (XSS risk)
  - **Correct:** httpOnly cookies or memory-only storage
  - **Incident:** Issue #342
```

---

## Adding Custom Patterns

### Pattern Template

Create new section using this structure:

```markdown
## [Pattern Name]

**When to use:** [Conditions/context]

**Structure:**

1. [Step 1]
2. [Step 2]
3. [Step 3]

**Example:**
```
[Concrete example from your project]
```

**Benefits:**
- [Benefit 1]
- [Benefit 2]

**Trade-offs:**
- [Trade-off 1]
- [Trade-off 2]
```

### Pattern Categories

**Workflow patterns:**
- Development workflows
- Testing workflows
- Deployment workflows
- Review workflows

**Code patterns:**
- Error handling
- Logging
- Validation
- Performance optimization

**Organizational patterns:**
- File organization
- Code sharing
- Versioning
- Documentation

### Example: Adding Release Workflow

```markdown
## Release Workflow

**When to use:** When cutting a new version

**Steps:**

1. Create release branch: `git checkout -b release/v1.2.0`
2. Update version in `package.json`, `Cargo.toml`, etc.
3. Update CHANGELOG.md with version and date
4. Run full test suite: `npm test`
5. Build: `npm run build`
6. Tag: `git tag -s v1.2.0 -m "Release v1.2.0"`
7. Push: `git push --tags origin main`
8. Publish: `npm publish`
9. Create GitHub release with notes

**Checklist:**
- [ ] CHANGELOG updated
- [ ] Version bumped
- [ ] Tests pass
- [ ] Build succeeds
- [ ] Tag created
- [ ] Published to registry
- [ ] GitHub release created

**Common issues:**
- Forgot to update version → Publishing fails
- Didn't run tests → Publish broken version
- Wrong version bump (breaking change as minor)
```

---

## Adding New Project Types

If the generated template doesn't fit, create custom template.

### Template Location

Create in skill directory:
```
.claude/skills/claude-md-gen/templates/my-custom-type.md
```

### Template Structure

Use existing templates as reference. Include:

1. **Project Structure** (with placeholders)
2. **Tech Stack** (with placeholders)
3. **Development Workflow** (type-specific)
4. **Domain-Specific Sections** (2-3 unique to this type)
5. **Quality Gates** (type-specific)
6. **Common Commands** (with placeholders)
7. **Anti-Patterns** (type-defaults)

**Placeholders to use:**
- `[PROJECT_NAME]`
- `[LANGUAGE_VERSION]`
- `[FRAMEWORK_VERSION]`
- `[DIRECTORY_STRUCTURE]`
- `[COMMANDS]`
- `[LINTERS]`

### Question Framework

Create matching questions file:
```
.claude/skills/claude-md-gen/questions/my-custom-type-questions.md
```

Follow structure from existing question files (5-7 questions across categories).

### Update Auto-Detection

Edit `.claude/skills/claude-md-gen/SKILL.md`:

```markdown
my-custom-type:
  IF [detection indicators]
  → my-custom-type
```

---

## Evolving with Project Stages

### Early Stage (0-6 months)

**Focus:**
- Core workflows
- Essential quality gates
- Basic anti-patterns

**Lightweight:** Don't over-document yet

### Growth Stage (6-18 months)

**Add:**
- Team conventions
- Deployment workflows
- Performance standards
- Security guidelines

**Document patterns that emerged**

### Mature Stage (18+ months)

**Expand:**
- Complex workflows
- Integration patterns
- Troubleshooting guides
- Architecture decision records

**Maintain as knowledge base**

---

## Integrating with .claude/ Structure

If you create `.claude/` directory structure:

### Option 1: Reference Files

Keep CLAUDE.md at root, reference detailed files:

```markdown
## Workflow Details

See `.claude/workflows/` for detailed workflows:
- `.claude/workflows/feature-development.md`
- `.claude/workflows/bug-fixes.md`
- `.claude/workflows/releases.md`
```

### Option 2: Move to .claude/

Move main file:
```bash
mv CLAUDE.md .claude/CLAUDE.md
```

Tell AI where to find it in conversations.

### Option 3: Split by Concern

```
.claude/
├── CLAUDE.md (main entry point)
├── workflows/
│   ├── development.md
│   ├── testing.md
│   └── deployment.md
├── patterns/
│   ├── error-handling.md
│   └── logging.md
└── templates/
    ├── component.tsx
    └── api-endpoint.go
```

Reference from main CLAUDE.md:
```markdown
## Workflows

See `.claude/workflows/` for detailed workflows.
```

---

## Maintenance Schedule

### Weekly (Active Development)

- Add new anti-patterns discovered
- Update commands if changed
- Note new tools adopted

### Monthly

- Review workflows for accuracy
- Update examples to reflect current code
- Adjust quality gates if tooling changed

### Quarterly

- Major section additions (new patterns, workflows)
- Refactor for clarity
- Remove obsolete sections
- Align with project current state

### Annually

- Comprehensive review
- Consider regeneration if project drastically changed
- Update project type if evolved (e.g., CLI → API service)

---

## Regeneration

### When to Regenerate

Consider regenerating when:
- Project type fundamentally changed
- Tech stack completely replaced
- CLAUDE.md no longer matches reality
- Starting fresh is easier than updating

### How to Regenerate

1. **Backup current CLAUDE.md:**
   ```bash
   cp CLAUDE.md CLAUDE.md.backup
   ```

2. **Run skill again:**
   ```bash
   /claude-md-gen
   ```

3. **Merge custom sections:**
   - Extract custom sections from backup
   - Integrate into new generated file
   - Preserve custom patterns/workflows

4. **Review and adjust:**
   - Verify accuracy
   - Update outdated parts
   - Remove obsolete sections

---

## Contributing Improvements

### To the Skill Itself

If you create valuable patterns or templates:

1. Consider adding to skill library
2. Update pattern files in `.claude/skills/claude-md-gen/patterns/`
3. Update templates in `.claude/skills/claude-md-gen/templates/`
4. Share with team or community

### Pattern Library Additions

Create new pattern file:
```
.claude/skills/claude-md-gen/patterns/my-pattern.md
```

Follow structure from existing patterns:
- When to use
- Structure/template
- Examples (2-3)
- Customization points

Update SKILL.md to reference new pattern.

---

## Tips for Effective CLAUDE.md

**DO:**
- ✅ Keep concrete and specific
- ✅ Use actual commands and examples
- ✅ Update as project changes
- ✅ Focus on what's actually used
- ✅ Document the "why" not just "what"
- ✅ Include common pitfalls
- ✅ Show real examples from codebase

**DON'T:**
- ❌ Document ideal processes that aren't followed
- ❌ Leave outdated information
- ❌ Over-document early (start small, grow)
- ❌ Copy-paste from other projects without adapting
- ❌ Use vague language ("properly configure")
- ❌ Skip examples (show, don't just tell)

---

## Getting Help

**Skill location:** `.claude/skills/claude-md-gen/`

**Files:**
- `SKILL.md` - Main skill logic
- `templates/` - Project type templates
- `patterns/` - Reusable patterns
- `questions/` - Question frameworks
- `checklist.md` - Quality verification
- `customization-guide.md` - This file

**Regenerate sections:**
Ask AI to regenerate specific sections using patterns from the skill library.

**Example:**
"Using the decision-matrices pattern from claude-md-gen, create a section for evaluating database choices in my API project."
