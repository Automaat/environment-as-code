# CLI Tool Template

Base structure for command-line tools and utilities.

---

## Project Structure

```
[DIRECTORY_STRUCTURE]
```

## Tech Stack

**Language:** [LANGUAGE_VERSION]
**CLI Framework:** [CLI_FRAMEWORK] (cobra, clap, commander, etc.)
**Testing:** [TEST_FRAMEWORK]
**Linting:** [LINTERS]
**Build Tool:** [BUILD_TOOL]
**Dependency Mgmt:** [DEP_MGMT]

## Development Workflow

### Adding New Command

1. Create command file in `[COMMANDS_DIR]/`
2. Define flags/arguments: [FLAG_DEFINITION_APPROACH]
3. Implement command logic in `[LOGIC_DIR]/`
4. Add tests in `[TEST_DIR]/`
5. Register command in root: [REGISTRATION_APPROACH]
6. Update help text
7. Test manually: `[RUN_COMMAND] [new-command] --help`

### Adding New Flag

1. Define flag in command struct/config
2. Add validation logic
3. Update help documentation
4. Add test cases for flag combinations
5. Test: `[RUN_COMMAND] [command] --[flag] value`

### Error Handling

1. Use error types: [ERROR_TYPE_APPROACH]
2. Provide actionable messages
3. Include exit codes: [EXIT_CODE_CONVENTIONS]
4. Log to stderr for errors
5. Test error paths

## Architecture

### Command Structure

```
[COMMAND_HIERARCHY]
```

**Command pattern:**
```[LANGUAGE]
[COMMAND_DEFINITION_EXAMPLE]
```

### Input Validation

**Approach:** [VALIDATION_APPROACH]

**Rules:**
- Required vs optional arguments
- Type validation (string, int, bool, enum)
- Format validation (URL, path, email, etc.)
- Range validation (min/max, allowed values)

**Validation example:**
```[LANGUAGE]
[VALIDATION_EXAMPLE]
```

### Output Formatting

**Supported formats:** [OUTPUT_FORMATS] (text, JSON, YAML, table, etc.)

**Flag:** `--output=[format]` or `-o [format]`

**Implementation:**
```[LANGUAGE]
[OUTPUT_FORMAT_EXAMPLE]
```

### Configuration

**Config locations:**
1. `[GLOBAL_CONFIG_PATH]` - Global defaults
2. `[PROJECT_CONFIG_PATH]` - Project-specific
3. Environment variables: `[ENV_VAR_PREFIX]_*`
4. CLI flags (highest priority)

**Config format:** [CONFIG_FORMAT] (YAML, TOML, JSON)

**Loading order:** CLI flags > env vars > project config > global config > defaults

## Error Handling

### Exit Codes

```
0   - Success
1   - General error
2   - Invalid input
3   - File not found
4   - Permission denied
[CUSTOM_EXIT_CODES]
```

### Error Messages

**Format:** `Error: [clear description]. [actionable suggestion]`

**Examples:**
```
Error: Config file not found at ~/.config/[tool]/config.yml. Run '[tool] init' to create it.
Error: Invalid format 'xml'. Supported formats: json, yaml, text.
Error: Permission denied writing to /etc/[tool]/. Try running with sudo.
```

## Testing

### Test Categories

1. **Unit tests:** Individual functions
2. **Command tests:** Full command execution
3. **Integration tests:** External dependencies
4. **E2E tests:** Real-world scenarios

### Test Structure

```[LANGUAGE]
[TEST_EXAMPLE]
```

### Test Data

**Location:** `[TEST_DATA_DIR]/`

**Fixtures:** [FIXTURE_APPROACH]

## Build & Distribution

### Building

```bash
# Development build
[DEV_BUILD_COMMAND]

# Production build
[PROD_BUILD_COMMAND]

# Cross-compile (if applicable)
[CROSS_COMPILE_COMMAND]
```

**Output:** `[BUILD_OUTPUT_DIR]/[BINARY_NAME]`

### Installation

[INSTALLATION_APPROACH]

```bash
# Install locally
[LOCAL_INSTALL_COMMAND]

# Install globally
[GLOBAL_INSTALL_COMMAND]
```

### Versioning

**Format:** [VERSION_FORMAT] (semver, calver, etc.)

**Show version:**
```bash
[TOOL_NAME] --version
```

## Quality Gates

Before committing:

- [ ] [LINTER_NAME] passes
- [ ] [TEST_FRAMEWORK] tests pass
- [ ] All commands run with `--help`
- [ ] Exit codes correct
- [ ] Error messages actionable
- [ ] [HOOK_NAME] hooks pass (if configured)

Commands:
```bash
[LINT_COMMAND]
[TEST_COMMAND]
[BUILD_COMMAND]
```

## Common Commands

```bash
# Run in development
[DEV_RUN_COMMAND]

# Run tests
[TEST_COMMAND]

# Run specific test
[TEST_SPECIFIC_COMMAND]

# Run linter
[LINT_COMMAND]

# Fix linting issues
[LINT_FIX_COMMAND]

# Build binary
[BUILD_COMMAND]

# Install locally
[INSTALL_COMMAND]

# Show help
[HELP_COMMAND]
```

## Output Templates

[IF_STRUCTURED_OUTPUT]

### JSON Output

```json
[JSON_OUTPUT_TEMPLATE]
```

### Text Output

```
[TEXT_OUTPUT_TEMPLATE]
```

### Table Output

```
[TABLE_OUTPUT_TEMPLATE]
```

[END_IF_STRUCTURED_OUTPUT]

## Help Documentation

### Command Help Format

```
[TOOL_NAME] [COMMAND] - [Short description]

Usage:
  [TOOL_NAME] [COMMAND] [FLAGS] [ARGS]

Arguments:
  [ARG1]  [Description]
  [ARG2]  [Description]

Flags:
  -f, --flag1  [Description]
  -o, --output [Description] (default: "[default_value]")

Examples:
  [TOOL_NAME] [COMMAND] [example1]
  [TOOL_NAME] [COMMAND] --flag1 [example2]
```

## Anti-Patterns

**AVOID:**

- ❌ Silent failures (always output error messages)
- ❌ Unclear error messages ("something went wrong")
- ❌ Using stdout for errors (use stderr)
- ❌ Inconsistent flag names across commands
- ❌ Missing `--help` for any command
- ❌ Hardcoded paths (use config or flags)
- ❌ Assuming file permissions (check and error gracefully)
- ❌ No progress indicators for long operations
- ❌ Breaking changes without version bump
- ❌ Undefined exit codes

**REASON:** CLI tools must be predictable, scriptable, and provide clear feedback. Poor error handling breaks automation.

## Extensibility

Add sections as needs evolve:
- Shell completion scripts
- Plugin system
- Interactive mode
- Progress indicators
- Logging levels

Follow section structure above. Keep concrete and actionable.
