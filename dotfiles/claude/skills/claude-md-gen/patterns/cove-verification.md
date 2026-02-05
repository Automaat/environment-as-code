# Chain of Verification (CoVe) Pattern

Two-phase analysis: internal reasoning → verify with external search/questions.

---

## When to Use

- High-stakes decisions requiring accuracy
- Claims that need factual verification
- Technical assessments with objective criteria
- Research requiring source validation
- Complex analysis prone to errors

## Pattern Structure

### Phase 1: Internal Analysis

Generate initial answer/conclusion using existing knowledge:

```
1. Analyze problem with available information
2. Form hypothesis or draft answer
3. Identify claims that need verification
4. Note assumptions made
```

### Phase 2: Verification

Verify internal analysis with external sources:

```
1. List specific claims to verify
2. Search/question for each claim
3. Compare findings with internal analysis
4. Adjust conclusions based on verification
5. Flag unverifiable claims
```

## Template

```markdown
## Internal Analysis

### Initial Assessment

[Your analysis based on available information]

### Key Claims

1. [Claim 1 that needs verification]
2. [Claim 2 that needs verification]
3. [Claim 3 that needs verification]

### Assumptions

- [Assumption 1]
- [Assumption 2]

## Verification Phase

### Claim 1: [Claim]

**Verification method:** [Search, documentation, test, etc.]
**Result:** [Verified / Contradicted / Inconclusive]
**Source:** [Where verified]
**Adjustment:** [Any changes to initial assessment]

### Claim 2: [Claim]

**Verification method:** [Method]
**Result:** [Status]
**Source:** [Source]
**Adjustment:** [Changes]

## Final Conclusion

[Adjusted conclusion after verification]

### Confidence Level

[High / Medium / Low] - [Reasoning]

### Unverified Claims

[Any claims that couldn't be verified]
```

## Example 1: Go Code Review (from go-review skill)

### Internal Analysis

```
Initial assessment: Function appears to have race condition
Key claim: Multiple goroutines access shared map without mutex
Assumption: This code runs concurrently
```

### Verification Phase

```
Claim: Multiple goroutines access shared map
Verification: Grep for goroutine launches calling this function
Result: Verified - found 3 call sites in goroutines
Source: KB: handlers/process.go:45, KB: workers/sync.go:89
Adjustment: Confirmed race condition exists
```

### Final Conclusion

```
Race condition verified. Shared map accessed from 3 goroutines
without synchronization.

Confidence: High - verified via code search
```

## Example 2: Plot Analysis Verification (from plot-analysis skill)

### Internal Analysis

```
Initial: Listing claims "fiber internet available"
Key claim: Fiber ISP operates in this location
Assumption: Listing information is accurate
```

### Verification Phase

```
Claim: Fiber internet available at Krakow, Podgórze
Verification: Search "[ISP] coverage Podgórze Krakow"
Result: Contradicted - ISP only covers Stare Miasto
Source: [ISP coverage map](url)
Adjustment: Listing claim is incorrect, only DSL available

Recommended action: Contact ISP directly for confirmation
```

### Final Conclusion

```
Fiber claim unverified and likely incorrect based on ISP coverage map.
Flag as potential misrepresentation.

Confidence: Medium - based on public coverage map, needs direct confirmation
```

## Customization Points

**Adjust verification methods:**
- Web search for factual claims
- Code search for technical claims
- API calls for data verification
- Documentation for spec compliance
- Tests for behavior verification

**Adjust confidence levels:**
- High: Multiple independent sources confirm
- Medium: Single authoritative source
- Low: No verification possible, relying on reasoning

**Adjust output format:**
- Inline verification (short checks)
- Separate section (complex analysis)
- Checklist format (multiple claims)
- Table format (many items)

## Integration with Project Type

**CLI tools:** Verify flag behavior, error codes, output formats
**Web apps:** Verify API contracts, browser compatibility, accessibility
**API services:** Verify endpoint behavior, status codes, auth requirements
**Research:** Verify facts, statistics, quotes, claims
**Libraries:** Verify API compatibility, dependency versions, behavior

## Benefits

- Reduces hallucination/incorrect claims
- Provides traceable reasoning
- Identifies confidence gaps
- Enables fact-checking
- Improves accuracy

## Anti-Patterns

**AVOID:**
- ❌ Skipping verification step (defeats purpose)
- ❌ Weak verification (searching own output, circular reasoning)
- ❌ Ignoring contradictions (must adjust conclusion)
- ❌ No confidence assessment (always state certainty level)
- ❌ Verifying everything (focus on key claims)
