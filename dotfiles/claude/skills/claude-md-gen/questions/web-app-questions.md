# Web App Questions

Questions for web application projects to customize CLAUDE.md.

---

## Domain Workflows (2-3 questions)

### Q1: What are the 2-3 most common development tasks?

**Purpose:** Define workflow sections

**Options:**
- Adding new components
- Creating new pages/routes
- Implementing API integrations
- Building forms with validation
- Styling and theming
- State management patterns
- Optimizing performance

**Follow-up:** For each task, how many steps typically involved?

### Q2: Is there a design system or component library?

**Purpose:** Determine if design system section needed

**Options:**
- Yes, custom design system
- Yes, using [Library] (MUI, Chakra, etc.)
- No, but consistent patterns exist
- No design system

**Follow-up if yes:** Key components? Breakpoints? Color system?

### Q3: What responsive breakpoints are used?

**Purpose:** Document responsive approach

**Options:**
- Mobile-first with specific breakpoints
- Desktop-first with specific breakpoints
- Fluid/flexible (no fixed breakpoints)
- Not responsive (fixed layout)

**Follow-up:** Specific pixel values? Approach to testing?

## Output Formats (1-2 questions)

### Q4: Are there component templates or scaffolding tools?

**Purpose:** Include templates in CLAUDE.md

**Options:**
- Yes, use generator/scaffolding tool
- Yes, copy existing component patterns
- No, write from scratch each time

**Follow-up if yes:** What templates? Where are they?

### Q5: How are error states and loading states handled?

**Purpose:** Document patterns

**Options:**
- Consistent pattern across app
- Per-component basis
- Using library (React Query, SWR, etc.)
- Not standardized

**Follow-up:** Example of the pattern?

## Tool Integrations (1-2 questions)

### Q6: What build/deploy workflow is used?

**Purpose:** Document commands

**Options:**
- CI/CD automated
- Manual deploy with scripts
- Platform-specific (Vercel, Netlify, etc.)
- Local only (no deployment)

**Follow-up:** Key commands? Environment variables?

### Q7: Are there specific linters, formatters, or type checkers?

**Purpose:** Add to quality gates

**Options:**
- ESLint + Prettier + TypeScript
- Just ESLint
- Other tools: [specify]
- No linting

**Follow-up:** Any custom rules? Pre-commit hooks?

## Quality Gates (1 question)

### Q8: What must pass before committing code?

**Purpose:** Define quality gates section

**Options (select multiple):**
- Linter passes
- Tests pass
- Type checking passes
- Build succeeds
- Accessibility tests pass
- Visual regression tests pass
- Manual review only

**Follow-up:** Commands to run these checks?

## Common Pitfalls (1 question)

### Q9: What mistakes do you frequently encounter?

**Purpose:** Populate anti-patterns section

**Options (select multiple):**
- Prop drilling (passing props too deep)
- Missing alt text on images
- Hardcoded breakpoints
- Client-side only routing
- Unoptimized images
- Non-semantic HTML
- Missing loading/error states
- Other: [specify]

**Follow-up:** Any project-specific pitfalls?

---

## Question Selection Strategy

**Always ask:** Q1, Q8 (workflows and quality gates are essential)

**Ask if design system detected:** Q2, Q3, Q4

**Ask if TypeScript detected:** Q7 (type checking)

**Ask if testing setup detected:** Q8 (include test requirements)

**Total questions:** 5-7 based on project characteristics
