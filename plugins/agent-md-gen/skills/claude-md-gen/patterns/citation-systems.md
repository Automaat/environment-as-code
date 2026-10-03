# Citation Systems Pattern

Source attribution formats for traceability and verification.

---

## When to Use

- Research and knowledge work
- Technical documentation
- Code analysis with references
- Fact-based reports
- Claims requiring evidence

## Pattern Structure

### Citation Types

1. **Internal Knowledge Base (KB)** - Files in project
2. **External Sources** - Web, documentation, papers
3. **Listing/Structured Data** - Real estate, jobs, products
4. **Code References** - File:line for code locations
5. **Tools/Commands** - CLI output, API responses

## Citation Formats

### Format 1: Knowledge Base (KB)

For internal project files:

```
[KB: file.ext:line]
[KB: path/to/file.ext:line-range]
[KB: file.ext] (if whole file)
```

**Examples:**
- `[KB: handlers/auth.go:45]`
- `[KB: docs/api.md:100-120]`
- `[KB: README.md]`

### Format 2: External Sources

For web content, docs, papers:

```
[Source: Title](URL)
[Doc: Title](URL)
[Paper: Authors, Year](URL)
```

**Examples:**
- `[Source: React Documentation](https://react.dev/reference)`
- `[Doc: FastAPI Tutorial](https://fastapi.tiangolo.com/tutorial/)`
- `[Paper: Raft Consensus, Ongaro & Ousterhout 2014](https://raft.github.io/raft.pdf)`

### Format 3: Listing Data

For structured sources (real estate, jobs, etc.):

```
[Listing: Item Title](URL)
[Job: Company - Role](URL)
[Plot: Location - ID](URL)
```

**Examples:**
- `[Listing: Działka Podgórze 1200m²](https://otodom.pl/...)`
- `[Job: Acme Corp - Senior Go Engineer](https://jobs.acme.com/...)`

### Format 4: Code References (File:Line)

For code locations in reports:

```
file.ext:line
path/to/file.ext:line
```

**Examples:**
- `handlers/auth.go:45`
- `src/components/Button.tsx:120`

### Format 5: Tool/Command Output

For CLI or API results:

```
[Tool: command](output snippet)
[API: endpoint](response excerpt)
```

**Examples:**
- `[Tool: go version](go version go1.21.0 darwin/arm64)`
- `[API: GET /api/status]({ "status": "healthy" })`

## Template

```markdown
## [Section Title]

[Content with inline citations]

According to [KB: file.go:45], the authentication middleware checks...

The official documentation states [Doc: FastAPI Auth](url) that...

Research by [Paper: Authors, Year](url) demonstrates...

### Sources

- [KB: file1.ext:line] - [Brief description]
- [KB: file2.ext:line-range] - [Brief description]
- [Source: Title](URL) - [Brief description]
- [Listing: Title](URL) - [Brief description]
```

## Example 1: Go Code Review (from go-review skill)

```markdown
## Race Condition Analysis

### Findings

Shared map `activeConnections` is accessed from multiple goroutines without
synchronization. Three concurrent access points identified:

1. Write operation in handlers/websocket.go:78
2. Read operation in handlers/metrics.go:145
3. Delete operation in handlers/cleanup.go:89

According to [KB: handlers/websocket.go:78], the connection is added:
```go
activeConnections[conn.ID] = conn  // No mutex
```

The metrics handler reads this map [KB: handlers/metrics.go:145]:
```go
count := len(activeConnections)  // Concurrent read
```

### Recommendation

Add sync.RWMutex as shown in [Doc: Go Memory Model](https://go.dev/ref/mem).

### References

- handlers/websocket.go:78 - Write operation
- handlers/metrics.go:145 - Read operation
- handlers/cleanup.go:89 - Delete operation
- [Doc: Go Memory Model](https://go.dev/ref/mem) - Synchronization guidance
```

## Example 2: Plot Analysis (from plot-analysis skill)

```markdown
## Zoning Verification

### Land Use Designation

According to [Listing: Działka Podgórze 1200m²](https://otodom.pl/...),
the plot is designated for residential use.

**Verification:** Checked against [Source: Krakow Zoning Map](url) which
confirms MN (residential low-density) designation for parcel 123/45.

### Building Permit Requirements

Local regulations [Doc: Podgórze Building Code](url) specify:
- Max building height: 12m (3 floors)
- Min setbacks: 6m from property line
- Max coverage: 40% of plot area

**Calculation:**
- Plot size: 1200m² [Listing]
- Max building footprint: 480m² (40% × 1200m²)
- Required setbacks: 6m × 4 sides = feasible on 1200m² plot

### Sources

- [Listing: Działka Podgórze 1200m²](https://otodom.pl/...) - Original listing
- [Source: Krakow Zoning Map](url) - MN designation verification
- [Doc: Podgórze Building Code](url) - Building regulations
```

## Example 3: API Documentation

```markdown
## Authentication Endpoints

### POST /api/auth/login

Authenticates user and returns JWT token.

**Implementation:** handlers/auth.go:34

**Request:**
```json
{
  "email": "user@example.com",
  "password": "secret"
}
```

**Response:** [KB: handlers/auth.go:67-72]
```json
{
  "token": "eyJhbG...",
  "expires_at": "2024-01-27T12:00:00Z"
}
```

**Error handling:** [KB: middleware/errors.go:23] wraps all auth errors
as 401 responses.

**Token format:** JWT with HS256, as specified in [KB: config/auth.go:12].
Follows [Doc: JWT Best Practices](https://datatracker.ietf.org/doc/html/rfc8725).

### References

- handlers/auth.go:34 - Login endpoint
- handlers/auth.go:67-72 - Success response
- middleware/errors.go:23 - Error handling
- config/auth.go:12 - Token configuration
- [Doc: JWT Best Practices](https://datatracker.ietf.org/doc/html/rfc8725)
```

## Example 4: Research Note

```markdown
# Raft Consensus Algorithm

## Core Concept

Raft provides distributed consensus through leader election and log replication
[Paper: Ongaro & Ousterhout, 2014](https://raft.github.io/raft.pdf).

## Leader Election

Election timeout randomized between 150-300ms [Paper: Section 5.2] to minimize
split votes. If follower receives no heartbeat within timeout, it transitions
to candidate state.

## Log Replication

Leader sends AppendEntries RPC to all followers [Paper: Figure 2]. Entry
committed once replicated to majority of servers.

**Implementation example:** etcd uses Raft for distributed key-value storage
[Source: etcd Architecture](https://etcd.io/docs/latest/learning/design/).

## Comparison with Paxos

Raft designed for understandability compared to Paxos [Paper: Section 1].
Key difference: Raft restricts possible states to simplify reasoning.

### Sources

- [Paper: In Search of an Understandable Consensus Algorithm, Ongaro & Ousterhout, 2014](https://raft.github.io/raft.pdf) - Original Raft paper
- [Source: etcd Architecture](https://etcd.io/docs/latest/learning/design/) - Production implementation
- [Doc: Raft Visualization](https://raft.github.io/) - Interactive explanation
```

## Customization Points

**Adjust citation density:**
- High density: Every claim cited (academic style)
- Medium density: Key claims cited (technical docs)
- Low density: Just controversial/non-obvious claims

**Adjust format:**
- Inline: `[KB: file:line]` within text
- Footnote-style: Numbered references at end
- Section-based: "Sources" section per major section

**Adjust granularity:**
- Line-level: `file.go:45` for precise location
- Function-level: `file.go:functionName` for broader context
- File-level: `file.go` for whole-file reference

**Adjust required fields:**
- Minimal: Just URL or file:line
- Standard: Format + brief description
- Detailed: Format + description + relevance note

## Integration with Project Type

**CLI tools:** Code references for implementation details
**Web apps:** Component references, design system citations
**API services:** Endpoint implementations, middleware references
**Research:** Papers, docs, web sources for facts
**Libraries:** API references, example code locations

## Benefits

- Traceability (can verify claims)
- Reproducibility (can check sources)
- Credibility (shows evidence)
- Navigability (can jump to references)
- Accountability (clear attribution)

## Anti-Patterns

**AVOID:**
- ❌ Vague references ("according to docs")
- ❌ Missing line numbers for code (can't locate)
- ❌ Dead links (always check URLs)
- ❌ Circular citations (citing own output)
- ❌ Inconsistent formats (mixing styles)
- ❌ Over-citing obvious facts (noise)
