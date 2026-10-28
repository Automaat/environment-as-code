# Python CLI Template

Base structure for Python CLI tools with TDD focus.

---

## Project Structure

```
[DIRECTORY_STRUCTURE]
```

## Tech Stack

**Language:** Python [PYTHON_VERSION]
**CLI Framework:** [CLI_FRAMEWORK] (click, typer, argparse)
**Testing:** [TEST_FRAMEWORK] (pytest, unittest)
**Linting:** [LINTERS] (ruff, pylint, mypy)
**Dependency Mgmt:** [DEP_MGMT] (poetry, pdm, pip-tools)
**Task Runner:** [TASK_RUNNER] (mise, make, invoke)

## Development Workflow

### Adding New Command (TDD)

1. Write test first in `[TEST_DIR]/test_[command].py`
2. Run tests (should fail): `[TEST_COMMAND]`
3. Implement command in `[COMMANDS_DIR]/[command].py`
4. Run tests (should pass)
5. Refactor if needed
6. Add integration test
7. Update help documentation

### Adding New Utility Function

1. Write test in `[TEST_DIR]/[module]/test_[function].py`
2. Implement in `[UTILS_DIR]/[module].py`
3. Run tests: `[TEST_SPECIFIC_COMMAND]`
4. Add docstring with examples
5. Type annotate fully

### Refactoring Existing Code

1. Ensure tests exist and pass
2. Make changes
3. Run full test suite: `[TEST_COMMAND]`
4. Check coverage: `[COVERAGE_COMMAND]`
5. Update docstrings if needed

## TDD Practices

### Test Structure

**Location:** `[TEST_DIR]/` mirrors `[SOURCE_DIR]/` structure

**Naming:**
- Test files: `test_[module].py`
- Test functions: `test_[function]_[scenario]`
- Test classes: `Test[ClassName]`

**Example:**
```python
[TEST_EXAMPLE]
```

### Fixtures and Mocking

**Shared fixtures:** `[TEST_DIR]/conftest.py`

**Common fixtures:**
- `tmp_path` - Temporary directories
- `monkeypatch` - Environment variables, attributes
- `[CUSTOM_FIXTURE_1]` - [Purpose]
- `[CUSTOM_FIXTURE_2]` - [Purpose]

**Mocking approach:** [MOCKING_APPROACH]

```python
[MOCKING_EXAMPLE]
```

### Coverage Requirements

**Minimum coverage:** [COVERAGE_PERCENTAGE]%

**Check coverage:**
```bash
[COVERAGE_COMMAND]
```

**Exclude from coverage:**
- `[EXCLUDE_PATTERN_1]`
- `[EXCLUDE_PATTERN_2]`

## Project Organization

### Shared Utilities

**Location:** `[UTILS_DIR]/`

**Modules:**
- `[UTILS_DIR]/[module1].py` - [Purpose]
- `[UTILS_DIR]/[module2].py` - [Purpose]
- `[UTILS_DIR]/[module3].py` - [Purpose]

**Import convention:**
```python
[IMPORT_EXAMPLE]
```

### CLI Commands

**Location:** `[COMMANDS_DIR]/`

**Structure:**
```python
[COMMAND_STRUCTURE_EXAMPLE]
```

### Entry Point

**File:** `[ENTRY_POINT_FILE]`

**Registration:**
```python
[ENTRY_POINT_EXAMPLE]
```

## Click/Typer Patterns

[IF_CLICK_OR_TYPER]

### Command Groups

```python
[COMMAND_GROUP_EXAMPLE]
```

### Options and Arguments

```python
[OPTIONS_ARGUMENTS_EXAMPLE]
```

### Output Formatting

```python
[OUTPUT_FORMATTING_EXAMPLE]
```

[END_IF_CLICK_OR_TYPER]

## Type Annotations

**Approach:** [TYPE_CHECKING_APPROACH]

**Run type checker:**
```bash
[TYPE_CHECK_COMMAND]
```

**Example:**
```python
[TYPE_ANNOTATION_EXAMPLE]
```

## Error Handling

### Custom Exceptions

**Location:** `[EXCEPTIONS_FILE]`

```python
[EXCEPTION_DEFINITION_EXAMPLE]
```

### Error Display

```python
[ERROR_DISPLAY_EXAMPLE]
```

### Exit Codes

```python
[EXIT_CODE_EXAMPLE]
```

## Configuration

**Config file:** `[CONFIG_FILE_LOCATION]`
**Format:** [CONFIG_FORMAT]

**Loading:**
```python
[CONFIG_LOADING_EXAMPLE]
```

**Precedence:** CLI args > env vars > config file > defaults

## Testing External APIs

[IF_EXTERNAL_APIS]

**Approach:** [API_TESTING_APPROACH] (vcr.py, responses, mock)

**Example:**
```python
[API_MOCKING_EXAMPLE]
```

**Fixture location:** `[API_FIXTURE_DIR]/`

[END_IF_EXTERNAL_APIS]

## Quality Gates

Before committing:

- [ ] [LINTER_NAME] passes (formatting, imports, style)
- [ ] [TYPE_CHECKER] passes (type annotations)
- [ ] [TEST_FRAMEWORK] tests pass
- [ ] Coverage >= [COVERAGE_PERCENTAGE]%
- [ ] All commands run with `--help`
- [ ] [HOOK_NAME] hooks pass (if configured)

Commands:
```bash
[LINT_COMMAND]
[TYPE_CHECK_COMMAND]
[TEST_COMMAND]
[COVERAGE_COMMAND]
```

## Common Commands

```bash
# Run tests
[TEST_COMMAND]

# Run specific test
[TEST_SPECIFIC_COMMAND]

# Run with coverage
[COVERAGE_COMMAND]

# Run linter
[LINT_COMMAND]

# Fix linting issues
[LINT_FIX_COMMAND]

# Run type checker
[TYPE_CHECK_COMMAND]

# Install dependencies
[INSTALL_COMMAND]

# Run CLI in development
[DEV_RUN_COMMAND]

# Build distribution
[BUILD_COMMAND]
```

## Output Templates

[IF_STRUCTURED_OUTPUT]

### JSON Output

```python
[JSON_OUTPUT_CODE]
```

Output:
```json
[JSON_OUTPUT_EXAMPLE]
```

### Table Output

```python
[TABLE_OUTPUT_CODE]
```

Output:
```
[TABLE_OUTPUT_EXAMPLE]
```

[END_IF_STRUCTURED_OUTPUT]

## Anti-Patterns

**AVOID:**

- ❌ Skipping tests (TDD is core to workflow)
- ❌ Testing implementation details (test behavior, not internals)
- ❌ No type annotations on public functions
- ❌ Mutable default arguments (`def func(arg=[])`)
- ❌ Bare `except:` clauses (catch specific exceptions)
- ❌ Print statements for output (use click.echo or logging)
- ❌ Hardcoded paths (use pathlib, config, or arguments)
- ❌ Mixing test and production code in same file
- ❌ Not using fixtures for repeated setup
- ❌ Tests that depend on execution order

**REASON:** Python best practices and TDD require discipline. These patterns lead to brittle tests and unmaintainable code.

## Extensibility

Add sections as needs evolve:
- Async command support
- Plugin architecture
- Configuration validation schemas
- Shell completion generation
- Documentation generation

Follow section structure above. Keep concrete and actionable.
