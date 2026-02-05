# Mixed Project Questions

Questions for hybrid projects with multiple component types to customize CLAUDE.md.

---

## Domain Workflows (2-3 questions)

### Q1: What are the main components of this project?

**Purpose:** Identify component types and structure

**Options (select multiple):**
- Frontend (web app)
- Backend (API service)
- CLI tool
- Library/package
- Database/data layer
- Background workers
- Other: [specify]

**Follow-up:** Primary purpose of each? How do they interact?

### Q2: Is this a monorepo or separate repositories?

**Purpose:** Document project structure

**Options:**
- Monorepo (all components in one repo)
- Separate repos with integration
- Hybrid (some together, some separate)

**Follow-up if monorepo:** Tool used (Nx, Turborepo, Lerna, workspaces)? Shared code location?

### Q3: What are the most common cross-component workflows?

**Purpose:** Define integration workflows

**Options (select multiple):**
- API client updates (backend → frontend)
- Shared type definitions
- End-to-end testing
- Coordinated deployments
- Shared utilities/libraries
- Other: [specify]

**Follow-up:** Steps for each workflow? Integration points?

## Output Formats (1-2 questions)

### Q4: How do components communicate?

**Purpose:** Document integration patterns

**Options (select multiple):**
- REST API (JSON)
- GraphQL
- gRPC
- Message queue
- Shared database
- File system
- Other: [specify]

**Follow-up:** Data formats? Versioning strategy?

### Q5: Are there component-specific output requirements?

**Purpose:** Include templates per component

**Options:**
- Yes, each component has specific formats
- Shared format across components
- Not applicable

**Follow-up if yes:** What formats per component?

## Tool Integrations (1-2 questions)

### Q6: How is shared code organized?

**Purpose:** Document code sharing strategy

**Options:**
- Shared directory/package in monorepo
- Published internal library
- Code duplication (each component independent)
- No shared code yet

**Follow-up:** What's shared? Import conventions? Versioning?

### Q7: How are integration tests organized?

**Purpose:** Document testing strategy

**Options:**
- Separate integration test suite
- Within each component's tests
- End-to-end test suite
- Manual integration testing
- No integration tests yet

**Follow-up:** How are components connected for tests? Test data management?

## Quality Gates (1 question)

### Q8: What must pass before committing?

**Purpose:** Define quality gates section

**Options (select multiple):**
- All component linters pass
- All component tests pass
- Integration tests pass
- Shared code tests pass
- No breaking changes in component APIs
- Build succeeds for all components
- Type checking passes (if shared types)
- Other: [specify]

**Follow-up:** Commands to verify all components? Order of checks?

## Common Pitfalls (1 question)

### Q9: What integration challenges occur frequently?

**Purpose:** Populate anti-patterns section

**Options (select multiple):**
- Tight coupling between components
- Shared state without clear ownership
- Circular dependencies
- Inconsistent coding standards across components
- Duplicating code instead of sharing
- Testing in isolation only (no integration tests)
- Different error handling per component
- Inconsistent logging formats
- No clear API contracts
- Building components in wrong order
- Version mismatches between components
- Other: [specify]

**Follow-up:** Project-specific integration issues?

---

## Question Selection Strategy

**Always ask:** Q1, Q2, Q8 (components, structure, quality gates are essential)

**Ask if monorepo:** Q6 (shared code)

**Ask if multiple languages:** Q3 (cross-component workflows)

**Ask if APIs detected:** Q4 (communication patterns)

**Ask if tests detected:** Q7 (integration testing)

**Total questions:** 5-7 based on project characteristics

---

## Component-Specific Follow-ups

After initial questions, load relevant questions from component-specific files:

- If frontend component: Ask 1-2 key questions from `web-app-questions.md`
- If backend API: Ask 1-2 key questions from `api-service-questions.md`
- If CLI tool: Ask 1-2 key questions from `cli-tool-questions.md`
- If library: Ask 1-2 key questions from `library-questions.md`

**Total across all questions:** Max 7 (mixed questions + component questions)

**Example question flow:**
1. Q1: Components? (Answer: API + Frontend)
2. Q2: Monorepo? (Answer: Yes, Turborepo)
3. Q8: Quality gates? (Answer: Linters + tests + integration)
4. From api-service-questions: Auth approach? (Answer: JWT)
5. From web-app-questions: Design system? (Answer: Custom)
6. Q6: Shared code? (Answer: Shared types package)
7. Q9: Common pitfalls? (Answer: Type mismatches)
