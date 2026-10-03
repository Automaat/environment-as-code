# Web App Template

Base structure for frontend/fullstack web applications.

---

## Project Structure

```
[DIRECTORY_STRUCTURE]
```

## Tech Stack

**Language:** [LANGUAGE_VERSION]
**Framework:** [FRAMEWORK_VERSION]
**Testing:** [TEST_FRAMEWORK]
**Linting:** [LINTERS]
**Build Tool:** [BUILD_TOOL]
**Dependency Mgmt:** [DEP_MGMT]

## Development Workflow

### Adding New Component

1. Create component file in `[COMPONENTS_DIR]/`
2. Follow naming convention: [NAMING_CONVENTION]
3. Add tests in `[TEST_DIR]/`
4. Import in parent component/page
5. Run `[DEV_COMMAND]` to verify

### Making Responsive Changes

1. Check breakpoints: [BREAKPOINTS]
2. Test on mobile viewport first
3. Use responsive utilities: [RESPONSIVE_UTILS]
4. Verify on all breakpoint sizes
5. Run visual regression tests: `[VISUAL_TEST_COMMAND]`

### Optimizing Assets

1. Images: [IMAGE_OPTIMIZATION_APPROACH]
2. Fonts: [FONT_LOADING_STRATEGY]
3. Code splitting: [CODE_SPLITTING_APPROACH]
4. Check bundle size: `[BUNDLE_ANALYZE_COMMAND]`

## Design System

[IF_DESIGN_SYSTEM_EXISTS]

### Components

Core components: [COMPONENT_LIST]

**Usage:**
```[LANGUAGE]
[COMPONENT_EXAMPLE]
```

### Breakpoints

```
Mobile:  [MOBILE_BREAKPOINT]
Tablet:  [TABLET_BREAKPOINT]
Desktop: [DESKTOP_BREAKPOINT]
Wide:    [WIDE_BREAKPOINT]
```

### Color System

[COLOR_PALETTE_APPROACH]

### Typography

[TYPOGRAPHY_SCALE]

[END_IF_DESIGN_SYSTEM]

## Accessibility

**Standards:** [ACCESSIBILITY_LEVEL] (WCAG 2.1 AA, etc.)

**Requirements:**
- Semantic HTML elements
- ARIA labels for interactive elements
- Keyboard navigation support
- Focus indicators visible
- Color contrast ratios: [CONTRAST_REQUIREMENTS]

**Testing:**
```bash
[ACCESSIBILITY_TEST_COMMAND]
```

## State Management

**Approach:** [STATE_MANAGEMENT_APPROACH]

**State structure:**
```[LANGUAGE]
[STATE_STRUCTURE_EXAMPLE]
```

**Best practices:**
- [STATE_BEST_PRACTICE_1]
- [STATE_BEST_PRACTICE_2]
- [STATE_BEST_PRACTICE_3]

## Routing

**Router:** [ROUTER_NAME]

**Route structure:**
```
[ROUTE_STRUCTURE]
```

**Dynamic routes:** [DYNAMIC_ROUTE_PATTERN]
**Protected routes:** [AUTH_ROUTE_PATTERN]

## API Integration

**Client:** [API_CLIENT]
**Base URL:** [API_BASE_URL_CONFIG]

**Request pattern:**
```[LANGUAGE]
[API_REQUEST_EXAMPLE]
```

**Error handling:**
```[LANGUAGE]
[ERROR_HANDLING_EXAMPLE]
```

## Build & Deploy

### Development

```bash
[DEV_SERVER_COMMAND]
```

### Production Build

```bash
[BUILD_COMMAND]
```

**Output:** `[BUILD_OUTPUT_DIR]/`

### Deployment

[DEPLOYMENT_APPROACH]

**Environment variables:**
- `[ENV_VAR_1]` - [PURPOSE]
- `[ENV_VAR_2]` - [PURPOSE]

## Quality Gates

Before committing:

- [ ] [LINTER_NAME] passes
- [ ] [TEST_FRAMEWORK] tests pass
- [ ] TypeScript/type checking passes (if applicable)
- [ ] Accessibility tests pass
- [ ] Build succeeds
- [ ] [HOOK_NAME] hooks pass (if configured)

Commands:
```bash
[LINT_COMMAND]
[TEST_COMMAND]
[TYPE_CHECK_COMMAND]
[BUILD_COMMAND]
```

## Common Commands

```bash
# Start development server
[DEV_COMMAND]

# Run tests
[TEST_COMMAND]

# Run linter
[LINT_COMMAND]

# Fix linting issues
[LINT_FIX_COMMAND]

# Type check
[TYPE_CHECK_COMMAND]

# Build for production
[BUILD_COMMAND]

# Analyze bundle size
[BUNDLE_ANALYZE_COMMAND]

# Run accessibility checks
[A11Y_COMMAND]
```

## Component Templates

[IF_COMPONENT_TEMPLATES_NEEDED]

### Functional Component

```[LANGUAGE]
[FUNCTIONAL_COMPONENT_TEMPLATE]
```

### Page Component

```[LANGUAGE]
[PAGE_COMPONENT_TEMPLATE]
```

### Layout Component

```[LANGUAGE]
[LAYOUT_COMPONENT_TEMPLATE]
```

[END_IF_COMPONENT_TEMPLATES]

## Anti-Patterns

**AVOID:**

- ❌ Inline styles in components (use design system/CSS modules)
- ❌ Fetching data in render methods (use effects/loaders)
- ❌ Prop drilling beyond 2 levels (use context/state management)
- ❌ Hardcoded breakpoints (use design system tokens)
- ❌ Missing alt text on images
- ❌ Non-semantic HTML (divs for buttons, etc.)
- ❌ Client-side only routing without fallbacks
- ❌ Unoptimized images in production
- ❌ Missing loading/error states for async operations
- ❌ Skipping accessibility attributes

**REASON:** These lead to maintenance issues, poor UX, accessibility violations, and performance problems.

## Extensibility

Add sections as needs evolve:
- Animation/transition guidelines
- Internationalization approach
- Analytics integration
- Error monitoring setup
- Performance budgets

Follow section structure above. Keep concrete and actionable.
