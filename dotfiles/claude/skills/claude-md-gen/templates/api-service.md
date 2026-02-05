# API Service Template

Base structure for backend APIs and services.

---

## Project Structure

```
[DIRECTORY_STRUCTURE]
```

## Tech Stack

**Language:** [LANGUAGE_VERSION]
**Framework:** [FRAMEWORK_VERSION] (Express, FastAPI, Gin, etc.)
**Database:** [DATABASE] + [ORM/DRIVER]
**Testing:** [TEST_FRAMEWORK]
**Linting:** [LINTERS]
**Dependency Mgmt:** [DEP_MGMT]

## Development Workflow

### Adding New Endpoint

1. Define route in `[ROUTES_DIR]/[resource].py|.go|.ts`
2. Implement handler in `[HANDLERS_DIR]/[resource].py|.go|.ts`
3. Add validation schema: [VALIDATION_APPROACH]
4. Write unit tests for handler logic
5. Write integration test for full request/response
6. Update API documentation
7. Test manually: `curl [METHOD] [ENDPOINT]`

### Adding Middleware

1. Create middleware file in `[MIDDLEWARE_DIR]/`
2. Implement middleware function: [MIDDLEWARE_PATTERN]
3. Add tests for middleware behavior
4. Register in application setup
5. Test on affected routes

### Database Changes

1. Create migration: `[MIGRATION_COMMAND]`
2. Update models in `[MODELS_DIR]/`
3. Update repository/query functions
4. Add tests with test database
5. Apply migration: `[MIGRATION_APPLY_COMMAND]`
6. Verify schema: `[MIGRATION_STATUS_COMMAND]`

## API Design

### Endpoint Structure

**Pattern:** `[API_PREFIX]/[VERSION]/[RESOURCE]/[ID]`

**Example routes:**
```
[ENDPOINT_EXAMPLES]
```

### Request/Response Format

**Content-Type:** `application/json`

**Request example:**
```json
[REQUEST_EXAMPLE]
```

**Response example:**
```json
[RESPONSE_EXAMPLE]
```

**Error response:**
```json
[ERROR_RESPONSE_EXAMPLE]
```

### Versioning

**Strategy:** [VERSIONING_STRATEGY] (URL path, header, etc.)

**Current version:** [CURRENT_VERSION]

**Deprecation policy:** [DEPRECATION_POLICY]

## Authentication & Authorization

**Method:** [AUTH_METHOD] (JWT, OAuth, API keys, etc.)

### Authentication

**Implementation:**
```[LANGUAGE]
[AUTH_IMPLEMENTATION_EXAMPLE]
```

**Token format:** [TOKEN_FORMAT]
**Token expiry:** [TOKEN_EXPIRY]
**Refresh strategy:** [REFRESH_STRATEGY]

### Authorization

**Approach:** [AUTHZ_APPROACH] (RBAC, ABAC, etc.)

**Roles:** [ROLE_LIST]

**Permission check:**
```[LANGUAGE]
[PERMISSION_CHECK_EXAMPLE]
```

## Database

### Models/Entities

**Location:** `[MODELS_DIR]/`

**Example model:**
```[LANGUAGE]
[MODEL_DEFINITION_EXAMPLE]
```

### Migrations

**Tool:** [MIGRATION_TOOL]

**Location:** `[MIGRATIONS_DIR]/`

**Commands:**
```bash
# Create migration
[MIGRATION_CREATE_COMMAND]

# Apply migrations
[MIGRATION_UP_COMMAND]

# Rollback migration
[MIGRATION_DOWN_COMMAND]

# Check status
[MIGRATION_STATUS_COMMAND]
```

### Query Patterns

[IF_ORM]
**ORM:** [ORM_NAME]

**Query example:**
```[LANGUAGE]
[ORM_QUERY_EXAMPLE]
```
[END_IF_ORM]

[IF_RAW_SQL]
**Approach:** Raw SQL with [DRIVER_NAME]

**Query example:**
```[LANGUAGE]
[RAW_SQL_EXAMPLE]
```
[END_IF_RAW_SQL]

## Validation

**Library:** [VALIDATION_LIBRARY]

**Schema definition:**
```[LANGUAGE]
[VALIDATION_SCHEMA_EXAMPLE]
```

**Validation in handler:**
```[LANGUAGE]
[VALIDATION_USAGE_EXAMPLE]
```

## Error Handling

### Error Types

```[LANGUAGE]
[ERROR_TYPES_DEFINITION]
```

### HTTP Status Codes

```
200 - OK
201 - Created
400 - Bad Request (validation errors)
401 - Unauthorized (missing/invalid auth)
403 - Forbidden (insufficient permissions)
404 - Not Found
409 - Conflict (duplicate resource)
422 - Unprocessable Entity
429 - Too Many Requests
500 - Internal Server Error
[CUSTOM_STATUS_CODES]
```

### Error Response Format

```json
{
  "error": {
    "code": "[ERROR_CODE]",
    "message": "[USER_FACING_MESSAGE]",
    "details": [OPTIONAL_DETAILS],
    "request_id": "[REQUEST_ID]"
  }
}
```

## Middleware Stack

**Order matters:**

1. [MIDDLEWARE_1] - [PURPOSE]
2. [MIDDLEWARE_2] - [PURPOSE]
3. [MIDDLEWARE_3] - [PURPOSE]
4. [MIDDLEWARE_4] - [PURPOSE]
5. Route handlers
6. Error handler (last)

## Rate Limiting

**Strategy:** [RATE_LIMIT_STRATEGY]

**Limits:**
- Authenticated: [AUTH_RATE_LIMIT]
- Anonymous: [ANON_RATE_LIMIT]
- By endpoint: [ENDPOINT_SPECIFIC_LIMITS]

**Implementation:**
```[LANGUAGE]
[RATE_LIMIT_IMPLEMENTATION]
```

## Logging

**Library:** [LOGGING_LIBRARY]

**Log levels:** DEBUG, INFO, WARN, ERROR

**Format:** [LOG_FORMAT] (JSON, text, etc.)

**Example:**
```[LANGUAGE]
[LOGGING_EXAMPLE]
```

**What to log:**
- All requests (method, path, status, duration)
- Authentication attempts
- Authorization failures
- Database errors
- External API calls
- Rate limit hits

**What NOT to log:**
- Passwords or tokens
- Full request/response bodies (unless debug)
- PII (unless required and secured)

## Testing

### Test Categories

1. **Unit tests:** Handler logic, utilities
2. **Integration tests:** Full request/response with test DB
3. **API tests:** Contract testing against specs

### Test Database

**Approach:** [TEST_DB_APPROACH]

**Setup:**
```bash
[TEST_DB_SETUP_COMMAND]
```

**Teardown:** [TEST_DB_TEARDOWN_APPROACH]

### Test Example

```[LANGUAGE]
[API_TEST_EXAMPLE]
```

## Configuration

**File:** `[CONFIG_FILE_LOCATION]`

**Environment-specific:** [ENV_CONFIG_APPROACH]

**Required variables:**
```
[ENV_VAR_1]=[DESCRIPTION]
[ENV_VAR_2]=[DESCRIPTION]
[ENV_VAR_3]=[DESCRIPTION]
```

**Loading:**
```[LANGUAGE]
[CONFIG_LOADING_EXAMPLE]
```

## Deployment

**Platform:** [DEPLOYMENT_PLATFORM]

**Build:**
```bash
[BUILD_COMMAND]
```

**Run:**
```bash
[RUN_COMMAND]
```

**Health check endpoint:** `[HEALTH_ENDPOINT]`

**Readiness check endpoint:** `[READINESS_ENDPOINT]`

## Quality Gates

Before committing:

- [ ] [LINTER_NAME] passes
- [ ] [TEST_FRAMEWORK] tests pass
- [ ] API tests pass
- [ ] Database migrations apply cleanly
- [ ] No sensitive data in logs
- [ ] Auth middleware on protected routes
- [ ] [HOOK_NAME] hooks pass (if configured)

Commands:
```bash
[LINT_COMMAND]
[TEST_COMMAND]
[MIGRATION_CHECK_COMMAND]
```

## Common Commands

```bash
# Start development server
[DEV_COMMAND]

# Run tests
[TEST_COMMAND]

# Run integration tests
[INTEGRATION_TEST_COMMAND]

# Run linter
[LINT_COMMAND]

# Create migration
[MIGRATION_CREATE_COMMAND]

# Apply migrations
[MIGRATION_UP_COMMAND]

# Seed database
[SEED_COMMAND]

# Build for production
[BUILD_COMMAND]

# Generate API docs
[DOCS_COMMAND]
```

## API Documentation

**Tool:** [DOCS_TOOL] (OpenAPI, Swagger, etc.)

**Location:** `[DOCS_LOCATION]`

**Generate docs:**
```bash
[DOCS_GENERATE_COMMAND]
```

**View docs:** `[DOCS_URL]`

## Anti-Patterns

**AVOID:**

- ❌ No authentication on endpoints (default to authenticated)
- ❌ Exposing internal error details to clients
- ❌ No rate limiting (enables abuse)
- ❌ Inconsistent error response formats
- ❌ Logging sensitive data (passwords, tokens, PII)
- ❌ No input validation (trust no input)
- ❌ Blocking operations in request handlers
- ❌ No database connection pooling
- ❌ Missing request IDs for tracing
- ❌ Returning 200 for errors

**REASON:** API security and reliability require defensive coding. These patterns lead to vulnerabilities and operational issues.

## Extensibility

Add sections as needs evolve:
- Caching strategy
- Background job processing
- WebSocket support
- File upload handling
- Pagination standards
- Filtering/sorting conventions

Follow section structure above. Keep concrete and actionable.
