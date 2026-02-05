# Python CLI Questions

Questions for Python CLI projects with TDD focus to customize CLAUDE.md.

---

## Domain Workflows (2-3 questions)

### Q1: Where are shared utilities organized?

**Purpose:** Document project organization

**Options:**
- `utils/` or `lib/` directory
- Within each command module
- Separate package
- No shared utilities yet

**Follow-up:** What are the main utility modules? Import conventions?

### Q2: What's the TDD workflow?

**Purpose:** Define development workflow section

**Options:**
- Strict TDD (always write test first)
- Test after implementation
- Tests for critical paths only
- No consistent approach yet

**Follow-up:** Coverage requirements? Test structure preferences?

### Q3: Which CLI framework is used?

**Purpose:** Include framework-specific patterns

**Options:**
- Click
- Typer
- argparse
- Other: [specify]

**Follow-up:** Command groups? Decorators? Common patterns?

## Output Formats (1-2 questions)

### Q4: What output formats does the CLI support?

**Purpose:** Create output templates

**Options (select multiple):**
- Plain text
- JSON
- Table format (rich, tabulate, etc.)
- Colored terminal output
- Quiet mode (minimal output)
- Verbose mode (detailed output)

**Follow-up:** Default format? How to control output level?

### Q5: How are API responses or external data handled?

**Purpose:** Document data handling patterns

**Options:**
- Parse and display immediately
- Cache locally
- Transform to internal models
- Pass through raw
- Not applicable (no external data)

**Follow-up:** Where are models defined? Serialization approach?

## Tool Integrations (1-2 questions)

### Q6: How are tests fixtures and mocking organized?

**Purpose:** Document test patterns

**Options:**
- Fixtures in `conftest.py`
- Fixtures in test files
- Mock external APIs (vcr.py, responses, etc.)
- Use test database
- Mix of approaches

**Follow-up:** Key fixtures? Mocking strategy for APIs?

### Q7: What's the dependency management approach?

**Purpose:** Document setup/install commands

**Options:**
- Poetry
- PDM
- pip-tools (requirements.txt)
- Just pip (no lock file)
- mise for task running

**Follow-up:** Key commands? Development setup steps?

## Quality Gates (1 question)

### Q8: What must pass before committing?

**Purpose:** Define quality gates section

**Options (select multiple):**
- Ruff (or other linter) passes
- mypy (or other type checker) passes
- pytest tests pass
- Coverage threshold met (X%)
- Black/isort formatting
- Pre-commit hooks
- Other: [specify]

**Follow-up:** Specific coverage percentage? Command to run all checks?

## Common Pitfalls (1 question)

### Q9: What mistakes are commonly made in this project?

**Purpose:** Populate anti-patterns section

**Options (select multiple):**
- Skipping tests (violates TDD)
- No type annotations
- Mutable default arguments
- Bare except clauses
- Print instead of click.echo/logging
- Hardcoded paths (not using pathlib)
- Tests depend on execution order
- Missing fixtures for repeated setup
- Other: [specify]

**Follow-up:** Project-specific patterns to avoid?

---

## Question Selection Strategy

**Always ask:** Q1, Q2, Q8 (organization, TDD, quality gates are essential)

**Ask if Click/Typer detected:** Q3 (framework patterns)

**Ask if external APIs used:** Q5, Q6 (data handling and mocking)

**Ask if pytest detected:** Q6 (fixtures and test patterns)

**Total questions:** 5-7 based on project characteristics
