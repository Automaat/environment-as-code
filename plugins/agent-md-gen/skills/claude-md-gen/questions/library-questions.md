# Library Questions

Questions for reusable library/package projects to customize CLAUDE.md.

---

## Domain Workflows (2-3 questions)

### Q1: What's the process for adding new public API functions?

**Purpose:** Define API development workflow

**Options:**
- Design API first, then implement
- TDD (tests first)
- Implement, then add tests
- Ad-hoc (no consistent process)

**Follow-up:** How do you decide what's public vs internal? Review process?

### Q2: What's the backwards compatibility policy?

**Purpose:** Document versioning and breaking changes

**Options:**
- Strict semver (no breaking changes in minor)
- Deprecation period before removal
- Breaking changes allowed with major bump
- No formal policy yet

**Follow-up:** How long is deprecation period? How are breaking changes communicated?

### Q3: Where are usage examples maintained?

**Purpose:** Document examples organization

**Options:**
- `examples/` directory with runnable code
- In documentation only
- In docstrings/doc comments
- Tests serve as examples
- No formal examples yet

**Follow-up:** How are examples kept up to date? Testing examples?

## Output Formats (1-2 questions)

### Q4: What documentation format is used?

**Purpose:** Document docs approach

**Options:**
- Generated from code (JSDoc, Sphinx, rustdoc, etc.)
- Handwritten markdown
- Both (code docs + guides)
- Minimal (README only)

**Follow-up:** Generation commands? Hosting (ReadTheDocs, GitHub Pages)?

### Q5: Are there specific output templates or formats users expect?

**Purpose:** Include in library patterns

**Options:**
- Structured data (JSON, objects)
- Text reports
- Visualization (charts, graphs)
- Not applicable (library doesn't generate output)

**Follow-up if applicable:** Output format conventions?

## Tool Integrations (1-2 questions)

### Q6: What's the testing strategy?

**Purpose:** Document test approach

**Options (select multiple):**
- Unit tests (all public APIs)
- Integration tests (component interactions)
- Example tests (examples work)
- Property-based tests
- Benchmarks/performance tests
- Manual testing only

**Follow-up:** Coverage requirements? Test organization?

### Q7: How is the library published/distributed?

**Purpose:** Document release process

**Options:**
- Package registry (npm, PyPI, crates.io, etc.)
- Git tags only
- Manual distribution
- Not published yet

**Follow-up:** Publishing commands? Versioning process? Release checklist?

## Quality Gates (1 question)

### Q8: What must pass before releasing a new version?

**Purpose:** Define quality gates section

**Options (select multiple):**
- Linter passes
- All tests pass
- Coverage threshold met (X%)
- All examples run successfully
- Documentation builds without errors
- CHANGELOG updated
- Version bumped correctly
- No breaking changes in minor/patch
- Manual review/approval
- Other: [specify]

**Follow-up:** Specific coverage percentage? Commands to verify?

## Common Pitfalls (1 question)

### Q9: What mistakes are commonly made?

**Purpose:** Populate anti-patterns section

**Options (select multiple):**
- Breaking API in minor version
- Undocumented public functions
- No usage examples
- Cryptic error messages
- Exporting internal utilities
- Tight coupling to dependency versions
- Missing type definitions (for JS libraries)
- Inconsistent naming across API
- Missing CHANGELOG entries
- Publishing without testing examples
- Other: [specify]

**Follow-up:** Project-specific patterns to avoid?

---

## Question Selection Strategy

**Always ask:** Q1, Q8 (API workflow and quality gates are essential)

**Ask if examples/ detected:** Q3 (examples maintenance)

**Ask if docs tool detected:** Q4 (documentation approach)

**Ask if published:** Q7 (publishing process)

**Ask about versioning:** Q2 (compatibility policy)

**Total questions:** 5-7 based on project characteristics
