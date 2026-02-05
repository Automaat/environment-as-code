# Library Template

Base structure for reusable libraries and packages.

---

## Project Structure

```
[DIRECTORY_STRUCTURE]
```

## Tech Stack

**Language:** [LANGUAGE_VERSION]
**Testing:** [TEST_FRAMEWORK]
**Linting:** [LINTERS]
**Build Tool:** [BUILD_TOOL]
**Dependency Mgmt:** [DEP_MGMT]
**Documentation:** [DOCS_TOOL]

## Development Workflow

### Adding New Public API

1. Design API surface (function signature, types)
2. Write tests first (TDD approach)
3. Implement in `[SOURCE_DIR]/`
4. Add docstring/doc comments
5. Add usage example in `[EXAMPLES_DIR]/`
6. Update API documentation
7. Consider backwards compatibility

### Adding Internal Utility

1. Implement in `[INTERNAL_DIR]/` or similar
2. Keep private (not exported)
3. Add tests
4. No documentation needed (internal only)

### Making Breaking Change

1. Check [BREAKING_CHANGE_POLICY]
2. Document in CHANGELOG under "Breaking Changes"
3. Bump version: [VERSION_BUMP_RULE]
4. Add migration guide
5. Consider deprecation period first

## Public API Design

### Exported Functions/Classes

**Location:** `[PUBLIC_API_FILES]`

**Naming convention:** [NAMING_CONVENTION]

**Example:**
```[LANGUAGE]
[PUBLIC_API_EXAMPLE]
```

### API Surface Principles

- **Minimal:** Expose only what's necessary
- **Consistent:** Similar operations use similar patterns
- **Typed:** Full type annotations/declarations
- **Documented:** Every public symbol has docs
- **Stable:** Avoid breaking changes

### Backwards Compatibility

**Policy:** [COMPATIBILITY_POLICY]

**Versioning:** [VERSIONING_SCHEME] (Semantic Versioning, etc.)

**Breaking changes:**
[BREAKING_CHANGE_RULES]

**Deprecation process:**
1. Mark as deprecated in version X
2. Add warning in documentation
3. Remove in version X+2 (minimum)

## Documentation

### API Documentation

**Tool:** [DOCS_TOOL] (JSDoc, Sphinx, godoc, rustdoc, etc.)

**Location:** `[DOCS_DIR]/`

**Generate docs:**
```bash
[DOCS_GENERATE_COMMAND]
```

**Docstring format:**
```[LANGUAGE]
[DOCSTRING_EXAMPLE]
```

### Usage Examples

**Location:** `[EXAMPLES_DIR]/`

**Coverage:**
- Basic usage (getting started)
- Common use cases
- Advanced patterns
- Edge cases

**Example file structure:**
```
[EXAMPLES_STRUCTURE]
```

**Running examples:**
```bash
[RUN_EXAMPLE_COMMAND]
```

### README

**Sections:**
1. Installation
2. Quick start (1-2 minute example)
3. Features
4. Documentation link
5. Examples link
6. Contributing
7. License

## Testing

### Test Coverage

**Minimum:** [COVERAGE_PERCENTAGE]%

**Focus areas:**
- All public APIs (100% coverage)
- Edge cases
- Error conditions
- Type checking (if applicable)

### Test Structure

```[LANGUAGE]
[TEST_EXAMPLE]
```

### Test Categories

1. **Unit tests:** Individual functions
2. **Integration tests:** Component interactions
3. **API tests:** Public interface contracts
4. **Example tests:** Examples in docs work

## Versioning

**Scheme:** [VERSION_SCHEME]

**Format:** [VERSION_FORMAT]

**Bump rules:**
- MAJOR: Breaking changes
- MINOR: New features (backwards compatible)
- PATCH: Bug fixes (backwards compatible)

**Version file:** `[VERSION_FILE]`

**Tagging:**
```bash
git tag -a v[VERSION] -m "Release v[VERSION]"
git push origin v[VERSION]
```

## Publishing

### Build Artifacts

**Build command:**
```bash
[BUILD_COMMAND]
```

**Output:** `[BUILD_OUTPUT_DIR]/`

**Artifacts:**
- [ARTIFACT_1]
- [ARTIFACT_2]
- [ARTIFACT_3]

### Release Process

1. Update CHANGELOG with version and date
2. Bump version in [VERSION_FILES]
3. Run full test suite: `[TEST_COMMAND]`
4. Build: `[BUILD_COMMAND]`
5. Tag release: `git tag v[VERSION]`
6. Publish: `[PUBLISH_COMMAND]`
7. Push tags: `git push --tags`
8. Create GitHub release with notes

### Package Registry

**Registry:** [PACKAGE_REGISTRY] (npm, PyPI, crates.io, etc.)

**Package name:** `[PACKAGE_NAME]`

**Publish command:**
```bash
[PUBLISH_COMMAND]
```

## Dependencies

### Dependency Policy

**Approach:** [DEPENDENCY_APPROACH] (minimal, batteries-included, etc.)

**Allowed dependencies:**
- [DEPENDENCY_CATEGORY_1]: [EXAMPLES]
- [DEPENDENCY_CATEGORY_2]: [EXAMPLES]

**Avoid:**
- Large dependencies for small features
- Dependencies with poor maintenance
- Dependencies with different license requirements

### Peer Dependencies

[IF_PEER_DEPENDENCIES]

**Peer dependencies:**
- [PEER_DEP_1] - [REASON]
- [PEER_DEP_2] - [REASON]

**Version ranges:** [VERSION_RANGE_STRATEGY]

[END_IF_PEER_DEPENDENCIES]

## Error Handling

### Error Types

```[LANGUAGE]
[ERROR_TYPES_DEFINITION]
```

### Error Messages

**Guidelines:**
- Clear and actionable
- Include context (what was attempted)
- Suggest fix when possible
- Consistent format across library

**Example:**
```[LANGUAGE]
[ERROR_EXAMPLE]
```

## Type Definitions

[IF_TYPED_LANGUAGE]

**Type export:**
```[LANGUAGE]
[TYPE_EXPORT_EXAMPLE]
```

**Generic types:**
```[LANGUAGE]
[GENERIC_TYPE_EXAMPLE]
```

[END_IF_TYPED_LANGUAGE]

[IF_TYPE_DEFINITIONS]

**TypeScript definitions:** `[TYPE_DEFS_FILE]`

**Generation:** [TYPE_GEN_APPROACH]

[END_IF_TYPE_DEFINITIONS]

## Quality Gates

Before releasing:

- [ ] [LINTER_NAME] passes
- [ ] [TEST_FRAMEWORK] tests pass
- [ ] Coverage >= [COVERAGE_PERCENTAGE]%
- [ ] All examples run successfully
- [ ] Documentation generated without errors
- [ ] CHANGELOG updated
- [ ] Version bumped according to changes
- [ ] No breaking changes in minor/patch releases
- [ ] [HOOK_NAME] hooks pass (if configured)

Commands:
```bash
[LINT_COMMAND]
[TEST_COMMAND]
[COVERAGE_COMMAND]
[DOCS_COMMAND]
[EXAMPLES_TEST_COMMAND]
```

## Common Commands

```bash
# Run tests
[TEST_COMMAND]

# Run tests with coverage
[COVERAGE_COMMAND]

# Run linter
[LINT_COMMAND]

# Fix linting issues
[LINT_FIX_COMMAND]

# Build library
[BUILD_COMMAND]

# Generate documentation
[DOCS_COMMAND]

# Run examples
[RUN_EXAMPLES_COMMAND]

# Publish package
[PUBLISH_COMMAND]
```

## Examples

[IF_EXAMPLES_EXIST]

### Example 1: [EXAMPLE_1_TITLE]

```[LANGUAGE]
[EXAMPLE_1_CODE]
```

### Example 2: [EXAMPLE_2_TITLE]

```[LANGUAGE]
[EXAMPLE_2_CODE]
```

[END_IF_EXAMPLES]

## Anti-Patterns

**AVOID:**

- ❌ Breaking API changes in minor versions
- ❌ Undocumented public functions
- ❌ No usage examples
- ❌ Cryptic error messages
- ❌ Exporting internal utilities
- ❌ Tight coupling to specific versions of dependencies
- ❌ No type definitions (for JS libraries)
- ❌ Inconsistent naming across API
- ❌ Missing CHANGELOG entries
- ❌ Publishing without testing examples

**REASON:** Libraries require stability and clarity. Users depend on predictable behavior and clear documentation.

## Extensibility

Add sections as needs evolve:
- Plugin system design
- Custom configuration options
- Hooks/callbacks for extensions
- Performance benchmarks
- Migration guides between versions

Follow section structure above. Keep concrete and actionable.
