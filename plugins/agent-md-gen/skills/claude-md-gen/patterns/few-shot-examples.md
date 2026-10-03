# Few-Shot Examples Pattern

Examples as executable specifications showing expected behavior.

---

## When to Use

- Teaching complex concepts
- Defining API contracts
- Showing CLI usage patterns
- Test-driven development
- Documentation with examples

## Pattern Structure

### Components

1. **Context** - When/why to use this
2. **Input** - Example input data/command
3. **Process** - What happens (optional)
4. **Output** - Expected result
5. **Variations** - Edge cases, alternatives

### Example Format

```
## Example [N]: [Descriptive Title]

**Context:** [When this example applies]

**Input:**
```
[Input code/command/data]
```

**Expected Output:**
```
[Output code/response/result]
```

**Explanation:** [Why this works / what to note]
```

## Meta-Example: How to Write Examples

### Example 1: Basic Usage

**Context:** User wants to check tool version

**Input:**
```bash
mytool --version
```

**Expected Output:**
```
mytool version 1.2.3 (build abc123)
```

**Explanation:** Always show semantic version plus build identifier

### Example 2: Error Handling

**Context:** User provides invalid flag

**Input:**
```bash
mytool --invalid-flag
```

**Expected Output:**
```
Error: Unknown flag '--invalid-flag'

Did you mean: --input-file?

Run 'mytool --help' for usage.
```

**Explanation:** Suggest similar flags, point to help

### Example 3: Edge Case

**Context:** User provides empty input

**Input:**
```bash
mytool process --input ""
```

**Expected Output:**
```
Error: Input cannot be empty

Usage: mytool process --input FILE
```

**Explanation:** Validate inputs, show usage hint

## Example Set 1: CLI Tool Usage (from CLI skills)

### Example 1: Basic Command

**Context:** Run basic analysis

**Input:**
```bash
plot-triage https://otodom.pl/listing/123
```

**Expected Output:**
```
🏡 Plot Triage Analysis

Location: Podgórze, Kraków
Price: 250,000 PLN
Size: 1200 m²

⚡ Quick Assessment:
✅ Within budget
✅ Good location
⚠️  Zoning unclear - needs verification

Recommendation: PASS to verification stage

Next: Run 'plot-verify 123' for deep analysis
```

**Explanation:** Quick triage output with emoji, clear next step

### Example 2: With Options

**Context:** Output as JSON for parsing

**Input:**
```bash
plot-triage https://otodom.pl/listing/123 --output json
```

**Expected Output:**
```json
{
  "status": "pass",
  "location": "Podgórze, Kraków",
  "price": 250000,
  "size": 1200,
  "checks": {
    "budget": true,
    "location": true,
    "zoning": null
  },
  "recommendation": "verify",
  "next_steps": ["plot-verify 123"]
}
```

**Explanation:** Machine-readable format for automation

### Example 3: Error Case

**Context:** Invalid URL provided

**Input:**
```bash
plot-triage invalid-url
```

**Expected Output:**
```
Error: Invalid listing URL

Expected format: https://otodom.pl/listing/[ID]

Examples:
  plot-triage https://otodom.pl/listing/12345
  plot-triage https://otodom.pl/oferta/some-listing-ID123
```

**Explanation:** Clear error, show expected format with examples

## Example Set 2: API Requests (API service docs)

### Example 1: Success Response

**Context:** Authenticate user with valid credentials

**Request:**
```http
POST /api/auth/login HTTP/1.1
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "correct_password"
}
```

**Response:**
```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "expires_at": "2024-01-27T12:00:00Z",
  "user": {
    "id": "usr_123",
    "email": "user@example.com",
    "role": "user"
  }
}
```

**Explanation:** Returns JWT token valid for 24 hours

### Example 2: Validation Error

**Context:** Missing required field

**Request:**
```http
POST /api/auth/login HTTP/1.1
Content-Type: application/json

{
  "email": "user@example.com"
}
```

**Response:**
```http
HTTP/1.1 400 Bad Request
Content-Type: application/json

{
  "error": {
    "code": "validation_error",
    "message": "Validation failed",
    "details": [
      {
        "field": "password",
        "error": "required",
        "message": "Password is required"
      }
    ]
  }
}
```

**Explanation:** Returns 400 with specific field errors

### Example 3: Authentication Error

**Context:** Invalid credentials

**Request:**
```http
POST /api/auth/login HTTP/1.1
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "wrong_password"
}
```

**Response:**
```http
HTTP/1.1 401 Unauthorized
Content-Type: application/json

{
  "error": {
    "code": "invalid_credentials",
    "message": "Email or password is incorrect"
  }
}
```

**Explanation:** Vague message for security (don't reveal which field is wrong)

## Example Set 3: Test Fixtures (Python TDD)

### Example 1: Happy Path Test

**Context:** Test successful user creation

**Test Code:**
```python
def test_create_user_success():
    # Given
    user_data = {
        "email": "new@example.com",
        "password": "SecurePass123!"
    }

    # When
    response = client.post("/api/users", json=user_data)

    # Then
    assert response.status_code == 201
    assert response.json()["email"] == user_data["email"]
    assert "password" not in response.json()  # Never return password
    assert "id" in response.json()
```

**Explanation:** Verify status, returned data, and security (no password leak)

### Example 2: Duplicate Email Test

**Context:** Test duplicate email rejection

**Test Code:**
```python
def test_create_user_duplicate_email(existing_user):
    # Given - existing_user fixture creates user with test@example.com
    user_data = {
        "email": "test@example.com",  # Same as fixture
        "password": "DifferentPass456!"
    }

    # When
    response = client.post("/api/users", json=user_data)

    # Then
    assert response.status_code == 409  # Conflict
    assert response.json()["error"]["code"] == "duplicate_email"
```

**Explanation:** Use fixture for setup, verify specific error code

### Example 3: Input Validation Test

**Context:** Test password requirements

**Test Code:**
```python
@pytest.mark.parametrize("password,expected_error", [
    ("short", "min_length"),
    ("nouppercaseornumbers", "complexity"),
    ("", "required"),
])
def test_create_user_invalid_password(password, expected_error):
    # Given
    user_data = {
        "email": "test@example.com",
        "password": password
    }

    # When
    response = client.post("/api/users", json=user_data)

    # Then
    assert response.status_code == 400
    assert expected_error in response.json()["error"]["details"][0]["error"]
```

**Explanation:** Parametrize to test multiple invalid inputs efficiently

## Example Set 4: Library API Usage (library docs)

### Example 1: Basic Usage

**Context:** Parse and validate JSON schema

**Code:**
```python
from mylib import Schema, ValidationError

# Define schema
user_schema = Schema({
    "name": str,
    "age": int,
    "email": str
})

# Valid data
data = {"name": "Alice", "age": 30, "email": "alice@example.com"}
validated = user_schema.validate(data)  # ✓ Returns validated data

# Invalid data
invalid = {"name": "Bob", "age": "thirty"}  # age should be int
try:
    user_schema.validate(invalid)
except ValidationError as e:
    print(e.errors)  # {"age": "expected int, got str"}
```

**Explanation:** Shows both success and error paths

### Example 2: Advanced Usage

**Context:** Custom validators and optional fields

**Code:**
```python
from mylib import Schema, Field, validators

# Schema with custom validation
user_schema = Schema({
    "email": Field(str, validators=[validators.email()]),
    "age": Field(int, validators=[validators.range(18, 120)]),
    "nickname": Field(str, required=False),  # Optional
})

# Valid with optional field omitted
data = {"email": "alice@example.com", "age": 25}
validated = user_schema.validate(data)  # ✓ nickname defaults to None

# Invalid email format
invalid = {"email": "not-an-email", "age": 25}
try:
    user_schema.validate(invalid)
except ValidationError as e:
    print(e.errors)  # {"email": "invalid email format"}
```

**Explanation:** Shows optional fields and custom validators

### Example 3: Nested Schemas

**Context:** Validate nested objects

**Code:**
```python
from mylib import Schema

address_schema = Schema({
    "street": str,
    "city": str,
    "zip": str
})

user_schema = Schema({
    "name": str,
    "address": address_schema  # Nested schema
})

# Valid nested data
data = {
    "name": "Alice",
    "address": {
        "street": "123 Main St",
        "city": "Krakow",
        "zip": "30-001"
    }
}
validated = user_schema.validate(data)  # ✓ Works

# Invalid nested data
invalid = {
    "name": "Bob",
    "address": {
        "street": "456 Oak Ave"
        # Missing city and zip
    }
}
try:
    user_schema.validate(invalid)
except ValidationError as e:
    print(e.errors)  # {"address.city": "required", "address.zip": "required"}
```

**Explanation:** Shows schema composition and nested error paths

## Customization Points

**Adjust example count:**
- Quick reference: 1-2 examples
- Tutorial: 3-5 examples
- Comprehensive docs: 5-10 examples with variations

**Adjust detail level:**
- Minimal: Just input/output
- Standard: Add explanation
- Detailed: Include process/reasoning

**Adjust format:**
- Inline: Small snippets in prose
- Dedicated section: "Examples" section
- Separate file: `examples/` directory

**Adjust variations:**
- Happy path only: Simplest case
- Happy + error: Common patterns
- Comprehensive: All edge cases

## Integration with Project Type

**CLI tools:** Command examples, flag combinations, error cases
**Web apps:** Component usage, prop variations, styling examples
**API services:** Request/response pairs, auth flows, error responses
**Research:** Note-taking examples, citation formats, templates
**Libraries:** API usage, advanced patterns, integration examples

## Benefits

- Concrete understanding (not abstract)
- Executable specs (can test examples)
- Quick reference (copy-paste ready)
- Coverage of edge cases
- Learning by example

## Anti-Patterns

**AVOID:**
- ❌ Too few examples (just one happy path)
- ❌ Unrealistic examples (toy data only)
- ❌ No error examples (hides failure modes)
- ❌ Inconsistent format across examples
- ❌ Outdated examples (don't match current API)
- ❌ No explanation (just code dumps)
