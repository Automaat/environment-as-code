# CLI Tool Questions

Questions for command-line tool projects to customize CLAUDE.md.

---

## Domain Workflows (2-3 questions)

### Q1: What are the primary commands/subcommands?

**Purpose:** Define command structure section

**Options:**
- Single command with flags
- Multiple subcommands (like git)
- Hierarchical commands (nested subcommands)
- Interactive mode

**Follow-up:** List 2-3 most common commands? Typical usage?

### Q2: How should input validation work?

**Purpose:** Define validation approach

**Options:**
- Strict (reject immediately on invalid input)
- Lenient (warn but continue when possible)
- Interactive (prompt for corrections)
- Context-dependent

**Follow-up:** Required vs optional flags? Type validation?

### Q3: What's the testing approach?

**Purpose:** Document test strategy

**Options:**
- Unit tests for functions
- Command-level integration tests
- End-to-end tests with real inputs
- Manual testing only

**Follow-up:** How are external dependencies mocked?

## Output Formats (1-2 questions)

### Q4: What output formats are needed?

**Purpose:** Create output templates

**Options (select multiple):**
- Plain text (human-readable)
- JSON (machine-readable)
- YAML
- Table format
- Colored/formatted terminal output
- Other: [specify]

**Follow-up:** Default format? How to switch between formats?

### Q5: How should errors be displayed?

**Purpose:** Define error handling pattern

**Options:**
- Detailed error messages with suggestions
- Concise error messages
- Structured error output (JSON)
- Exit codes only

**Follow-up:** Exit code conventions? Error message format?

## Tool Integrations (1-2 questions)

### Q6: Are there configuration files?

**Purpose:** Document config approach

**Options:**
- Yes, [format: YAML/TOML/JSON]
- Environment variables only
- CLI flags only (no config file)
- Mix of config file + env vars + flags

**Follow-up:** Config file location? Precedence order?

### Q7: Does the tool integrate with external services/APIs?

**Purpose:** Document API interactions

**Options:**
- Yes, requires API keys/auth
- Yes, but optional functionality
- No external dependencies
- Local filesystem only

**Follow-up:** How are credentials managed? Testing strategy?

## Quality Gates (1 question)

### Q8: What must pass before committing?

**Purpose:** Define quality gates section

**Options (select multiple):**
- Linter passes
- Tests pass
- All commands run with `--help`
- Manual smoke test
- Build succeeds
- Binary size check
- Other: [specify]

**Follow-up:** Specific commands to run?

## Common Pitfalls (1 question)

### Q9: What mistakes are frequently made?

**Purpose:** Populate anti-patterns section

**Options (select multiple):**
- Silent failures (no error output)
- Unclear error messages
- Using stdout for errors (should use stderr)
- Inconsistent flag names
- Missing --help text
- Hardcoded paths
- No progress indicators for long operations
- Breaking changes without version bump
- Other: [specify]

**Follow-up:** Any project-specific anti-patterns?

---

## Question Selection Strategy

**Always ask:** Q1, Q4, Q8 (commands, output, and quality gates are essential)

**Ask if config detected:** Q6 (configuration)

**Ask if external deps detected:** Q7 (APIs/services)

**Ask if tests exist:** Q3 (testing approach)

**Total questions:** 5-7 based on project characteristics
