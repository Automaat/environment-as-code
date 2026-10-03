# Mixed Project Template

Base structure for hybrid projects combining multiple component types.

---

## Project Overview

[PROJECT_DESCRIPTION]

**Components:**
- [COMPONENT_1]: [TYPE] - [PURPOSE]
- [COMPONENT_2]: [TYPE] - [PURPOSE]
- [COMPONENT_3]: [TYPE] - [PURPOSE]

## Project Structure

```
[DIRECTORY_STRUCTURE]
```

## Tech Stack

### [Component 1]

**Language:** [LANGUAGE_VERSION]
**Framework:** [FRAMEWORK_VERSION]
**Purpose:** [PURPOSE]

### [Component 2]

**Language:** [LANGUAGE_VERSION]
**Framework:** [FRAMEWORK_VERSION]
**Purpose:** [PURPOSE]

### [Component 3]

[IF_COMPONENT_3_EXISTS]
**Language:** [LANGUAGE_VERSION]
**Framework:** [FRAMEWORK_VERSION]
**Purpose:** [PURPOSE]
[END_IF_COMPONENT_3]

### Shared

**Testing:** [TEST_FRAMEWORKS]
**Linting:** [LINTERS]
**Build Tool:** [BUILD_TOOLS]
**Monorepo Tool:** [MONOREPO_TOOL] (if applicable)

## Development Workflow

### Working Across Components

[IF_MONOREPO]

**Monorepo structure:** [MONOREPO_STRUCTURE]

**Commands:**
```bash
# Run command in specific workspace
[WORKSPACE_COMMAND]

# Run command in all workspaces
[ALL_WORKSPACES_COMMAND]
```

[END_IF_MONOREPO]

[IF_SEPARATE_REPOS]

**Repository structure:** [REPO_STRUCTURE]

**Local development:**
```bash
[LOCAL_DEV_SETUP]
```

[END_IF_SEPARATE_REPOS]

### Integration Points

**[Component 1] ↔ [Component 2]:**
- [INTEGRATION_METHOD_1]
- [INTEGRATION_METHOD_2]

**[Component 2] ↔ [Component 3]:**
[IF_COMPONENT_3_EXISTS]
- [INTEGRATION_METHOD_3]
- [INTEGRATION_METHOD_4]
[END_IF_COMPONENT_3]

### Shared Code

[IF_SHARED_CODE]

**Location:** `[SHARED_CODE_DIR]/`

**Contents:**
- [SHARED_MODULE_1] - [PURPOSE]
- [SHARED_MODULE_2] - [PURPOSE]
- [SHARED_MODULE_3] - [PURPOSE]

**Import convention:**
```[LANGUAGE]
[SHARED_IMPORT_EXAMPLE]
```

[END_IF_SHARED_CODE]

## Component-Specific Workflows

### [Component 1]: [Component Name]

#### Adding Feature

1. [STEP_1]
2. [STEP_2]
3. [STEP_3]
4. [STEP_4]

#### Testing

```bash
[COMPONENT_1_TEST_COMMAND]
```

#### Common Commands

```bash
# [Command purpose]
[COMPONENT_1_COMMAND_1]

# [Command purpose]
[COMPONENT_1_COMMAND_2]
```

### [Component 2]: [Component Name]

#### Adding Feature

1. [STEP_1]
2. [STEP_2]
3. [STEP_3]
4. [STEP_4]

#### Testing

```bash
[COMPONENT_2_TEST_COMMAND]
```

#### Common Commands

```bash
# [Command purpose]
[COMPONENT_2_COMMAND_1]

# [Command purpose]
[COMPONENT_2_COMMAND_2]
```

### [Component 3]: [Component Name]

[IF_COMPONENT_3_EXISTS]

#### Adding Feature

1. [STEP_1]
2. [STEP_2]
3. [STEP_3]

#### Testing

```bash
[COMPONENT_3_TEST_COMMAND]
```

#### Common Commands

```bash
# [Command purpose]
[COMPONENT_3_COMMAND_1]

# [Command purpose]
[COMPONENT_3_COMMAND_2]
```

[END_IF_COMPONENT_3]

## Integration Testing

**Approach:** [INTEGRATION_TEST_APPROACH]

**Test scenarios:**
1. [SCENARIO_1]
2. [SCENARIO_2]
3. [SCENARIO_3]

**Run integration tests:**
```bash
[INTEGRATION_TEST_COMMAND]
```

## Configuration

### Environment Variables

**Shared:**
```
[SHARED_ENV_VAR_1]=[DESCRIPTION]
[SHARED_ENV_VAR_2]=[DESCRIPTION]
```

**[Component 1]:**
```
[COMPONENT_1_ENV_VAR_1]=[DESCRIPTION]
[COMPONENT_1_ENV_VAR_2]=[DESCRIPTION]
```

**[Component 2]:**
```
[COMPONENT_2_ENV_VAR_1]=[DESCRIPTION]
[COMPONENT_2_ENV_VAR_2]=[DESCRIPTION]
```

### Config Files

**Shared:** `[SHARED_CONFIG_FILE]`
**[Component 1]:** `[COMPONENT_1_CONFIG_FILE]`
**[Component 2]:** `[COMPONENT_2_CONFIG_FILE]`

## Build & Deployment

### Build Order

[IF_BUILD_DEPENDENCIES]

**Dependencies:**
1. [BUILD_STEP_1]
2. [BUILD_STEP_2]
3. [BUILD_STEP_3]

**Build all:**
```bash
[BUILD_ALL_COMMAND]
```

[END_IF_BUILD_DEPENDENCIES]

### Deployment Strategy

**[Component 1]:** [DEPLOYMENT_1]
**[Component 2]:** [DEPLOYMENT_2]
**[Component 3]:** [DEPLOYMENT_3] (if applicable)

**Deploy commands:**
```bash
# Deploy [Component 1]
[DEPLOY_1_COMMAND]

# Deploy [Component 2]
[DEPLOY_2_COMMAND]
```

## Quality Gates

Before committing:

- [ ] All component linters pass
- [ ] All component tests pass
- [ ] Integration tests pass
- [ ] Shared code tests pass
- [ ] No breaking changes in component APIs
- [ ] [HOOK_NAME] hooks pass (if configured)

Commands:
```bash
# Lint all
[LINT_ALL_COMMAND]

# Test all
[TEST_ALL_COMMAND]

# Integration tests
[INTEGRATION_TEST_COMMAND]
```

## Common Commands

### Project-Wide

```bash
# Install all dependencies
[INSTALL_ALL_COMMAND]

# Run all tests
[TEST_ALL_COMMAND]

# Lint all
[LINT_ALL_COMMAND]

# Build all
[BUILD_ALL_COMMAND]

# Clean all
[CLEAN_ALL_COMMAND]
```

### Component-Specific

**[Component 1]:**
```bash
[COMPONENT_1_COMMANDS]
```

**[Component 2]:**
```bash
[COMPONENT_2_COMMANDS]
```

## Documentation

**Project docs:** `[PROJECT_DOCS_DIR]/`
**[Component 1] docs:** `[COMPONENT_1_DOCS_DIR]/`
**[Component 2] docs:** `[COMPONENT_2_DOCS_DIR]/`

**Generate all docs:**
```bash
[DOCS_GENERATE_COMMAND]
```

## Anti-Patterns

**AVOID:**

- ❌ Tight coupling between components
- ❌ Shared state without clear ownership
- ❌ Circular dependencies between components
- ❌ Inconsistent coding standards across components
- ❌ Duplicating code instead of using shared modules
- ❌ Testing components in isolation only (no integration tests)
- ❌ Different error handling patterns per component
- ❌ Inconsistent logging formats
- ❌ No clear API contracts between components
- ❌ Building components in wrong order

**REASON:** Mixed projects require clear boundaries and interfaces. Tight coupling makes components hard to develop and test independently.

## Component-Specific Sections

[INCLUDE_RELEVANT_SECTIONS_FROM_COMPONENT_TEMPLATES]

### [Section from Component 1 Template]

[COMPONENT_1_SPECIFIC_CONTENT]

### [Section from Component 2 Template]

[COMPONENT_2_SPECIFIC_CONTENT]

## Extensibility

Add sections as needs evolve:
- New components
- New integration points
- Shared utilities
- Cross-component workflows
- Performance optimization

Follow section structure above. Keep concrete and actionable.
