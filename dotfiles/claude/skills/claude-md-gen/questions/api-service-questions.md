# API Service Questions

Questions for backend API/service projects to customize CLAUDE.md.

---

## Domain Workflows (2-3 questions)

### Q1: What are the most common development tasks?

**Purpose:** Define workflow sections

**Options (select multiple):**
- Adding new endpoints
- Adding middleware
- Database schema changes/migrations
- Writing integration tests
- Adding authentication/authorization
- Other: [specify]

**Follow-up:** Typical steps for each? Where are files organized?

### Q2: What's the authentication/authorization approach?

**Purpose:** Document auth patterns

**Options:**
- JWT tokens
- OAuth 2.0
- API keys
- Session-based
- No auth (internal service)
- Multiple methods

**Follow-up:** Token format? Permission model (RBAC, ABAC)? Middleware setup?

### Q3: How are database changes managed?

**Purpose:** Document migration workflow

**Options:**
- Migration tool (Alembic, Flyway, etc.)
- ORM auto-migrations
- Manual SQL scripts
- No migrations (schemaless)

**Follow-up:** Migration commands? Rollback strategy?

## Output Formats (1-2 questions)

### Q4: What response format(s) does the API use?

**Purpose:** Document response structure

**Options:**
- JSON (standard REST)
- JSON:API spec
- GraphQL
- Protocol Buffers
- XML
- Other: [specify]

**Follow-up:** Error response format? Pagination format?

### Q5: How are errors structured?

**Purpose:** Create error handling section

**Options:**
- Custom error format
- RFC 7807 (Problem Details)
- Simple message + status code
- Framework default
- Inconsistent (needs standardization)

**Follow-up:** Error codes? Status code conventions?

## Tool Integrations (1-2 questions)

### Q6: What's the testing approach?

**Purpose:** Document test strategy

**Options (select multiple):**
- Unit tests (handler logic)
- Integration tests (with test DB)
- API contract tests
- E2E tests
- Manual testing only

**Follow-up:** Test database setup? How are external services mocked?

### Q7: Are there rate limiting or middleware requirements?

**Purpose:** Document middleware stack

**Options (select multiple):**
- Rate limiting
- CORS configuration
- Request logging
- Request ID tracing
- Compression
- Custom middleware
- None yet

**Follow-up:** Middleware order? Configuration approach?

## Quality Gates (1 question)

### Q8: What must pass before committing?

**Purpose:** Define quality gates section

**Options (select multiple):**
- Linter passes
- Tests pass (unit + integration)
- Database migrations apply cleanly
- API tests pass
- No sensitive data in logs
- Auth middleware on protected routes
- Security checks (SQL injection, etc.)
- Other: [specify]

**Follow-up:** Specific commands? Pre-commit hooks?

## Common Pitfalls (1 question)

### Q9: What mistakes are frequently encountered?

**Purpose:** Populate anti-patterns section

**Options (select multiple):**
- Missing authentication on endpoints
- Exposing internal errors to clients
- No rate limiting
- Inconsistent error formats
- Logging sensitive data
- No input validation
- Blocking operations in handlers
- Missing database connection pooling
- No request tracing
- Returning 200 for errors
- Other: [specify]

**Follow-up:** Project-specific anti-patterns?

---

## Question Selection Strategy

**Always ask:** Q1, Q4, Q8 (workflows, response format, quality gates are essential)

**Ask if auth detected:** Q2 (authentication details)

**Ask if database detected:** Q3 (migrations)

**Ask if middleware detected:** Q7 (middleware stack)

**Ask if tests exist:** Q6 (testing approach)

**Total questions:** 5-7 based on project characteristics
