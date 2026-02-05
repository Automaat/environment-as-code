# Triage Workflows Pattern

Quick filter (<3min) → Verification (5-15min) → Deep analysis (20-30min).

---

## When to Use

- High volume of items to evaluate
- Need quick reject before investing time
- Multi-stage evaluation process
- Progressive refinement of candidates
- Time-boxed analysis

## Pattern Structure

### Three Stages

1. **Triage** - Quick go/no-go (<3 min per item)
2. **Verification** - Validate key claims (5-15 min per item)
3. **Deep Analysis** - Comprehensive evaluation (20-30 min per item)

Each stage filters items for next stage.

## Template

```markdown
# Triage Stage (<3 minutes)

## Quick Checks

- [ ] [Check 1 - dealbreaker]
- [ ] [Check 2 - dealbreaker]
- [ ] [Check 3 - requirement]

**Decision:** [PASS to verification / REJECT]

**Reject reasons (if applicable):**
- [Reason 1]
- [Reason 2]

---

# Verification Stage (5-15 minutes)

**Only proceed if passed triage**

## Claims to Verify

1. [Claim 1] - [Verification method]
2. [Claim 2] - [Verification method]
3. [Claim 3] - [Verification method]

## Verification Results

| Claim | Status | Source | Notes |
|-------|--------|--------|-------|
| [Claim 1] | ✅/❌/⚠️ | [Source] | [Notes] |
| [Claim 2] | ✅/❌/⚠️ | [Source] | [Notes] |

**Decision:** [PASS to deep analysis / REJECT / FLAG for manual review]

**Flags (if applicable):**
- ⚠️ [Warning 1]
- ⚠️ [Warning 2]

---

# Deep Analysis Stage (20-30 minutes)

**Only proceed if passed verification**

## Comprehensive Evaluation

### [Analysis Area 1]

[Detailed analysis]

### [Analysis Area 2]

[Detailed analysis]

### [Analysis Area 3]

[Detailed analysis]

## Final Recommendation

**Status:** [ACCEPT / REJECT / CONDITIONAL]
**Score:** [X/Y]
**Confidence:** [High / Medium / Low]

**Next steps:**
1. [Step 1]
2. [Step 2]
```

## Example 1: Plot Triage (from plot-triage skill)

### Triage Stage (<3 min)

```markdown
## Quick Checks

- [x] Price within budget (max 300k PLN) - Listed: 250k PLN ✅
- [x] Location within 45min of Krakow - Podgórze, 20min ✅
- [x] Size adequate (min 800m²) - 1200m² ✅
- [ ] Zoning allows residential - Listing unclear ⚠️

**Decision:** PASS to verification (3 of 4 checks pass, unclear item needs verification)

**Time spent:** 2 minutes
```

### Verification Stage (5-15 min)

```markdown
## Claims to Verify

1. Zoning allows residential construction
2. Utilities available (power, water, sewage)
3. No ownership disputes

## Verification Results

| Claim | Status | Source | Notes |
|-------|--------|--------|-------|
| Residential zoning | ✅ | Local zoning map | Confirmed MN designation |
| Utilities available | ⚠️ | Utility company | Power yes, sewage requires septic |
| No disputes | ✅ | Public records | Clean title |

**Decision:** PASS to deep analysis

**Flags:**
- ⚠️ Septic tank required (no municipal sewage)
- ⚠️ Verify exact utility connection costs

**Time spent:** 12 minutes
```

### Deep Analysis Stage (20-30 min)

```markdown
## Comprehensive Evaluation

### Location Analysis
- 20min to center via Zakopiańska (good)
- Quiet residential area (excellent)
- Bus stop 800m (acceptable)
- Fiber internet available (excellent)

### Legal & Zoning
- Clean ownership history
- MN zoning confirmed
- Building permit feasible
- No environmental restrictions

### Land Quality
- Mostly flat (5% slope)
- South-facing (ideal)
- Some trees (pleasant)
- Soil test needed

### Cost Analysis
- Purchase: 250k PLN
- Utilities connection: ~40k PLN (includes septic)
- Total: 290k PLN (within budget)

## Final Recommendation

**Status:** ACCEPT for offer
**Score:** 14/20 points
**Confidence:** High

**Next steps:**
1. Order soil test (3000 PLN)
2. Legal verification (2000 PLN)
3. Get utility connection quotes
4. Make offer at 240k PLN

**Time spent:** 28 minutes
**Total time:** 42 minutes across all stages
```

## Example 2: Job Triage (from job-triage skill)

### Triage Stage (<3 min)

```markdown
## Quick Checks

- [x] Europe/Poland eligible - Remote Europe ✅
- [x] Tech match (Go, distributed systems) - Go, Kubernetes ✅
- [ ] Compensation stated - "Competitive salary" ❌
- [x] Not blockchain/crypto primary - Cloud infrastructure ✅

**Decision:** PASS to verification (3 of 4 pass, compensation needs clarification)

**Time spent:** 2 minutes
```

### Verification Stage (5-15 min)

```markdown
## Claims to Verify

1. Remote work policy (truly remote or hybrid)
2. Salary range (competitive = ?)
3. Company stage/funding
4. Team size and structure

## Verification Results

| Claim | Status | Source | Notes |
|-------|--------|--------|-------|
| Remote policy | ✅ | Company site | Fully remote, EU coverage |
| Salary range | ⚠️ | Glassdoor | 15-25k PLN/mo (mid-range) |
| Company stage | ✅ | Crunchbase | Series B, 80 employees |
| Team size | ✅ | LinkedIn | 8-person eng team |

**Decision:** PASS to deep analysis

**Flags:**
- ⚠️ Salary at market rate (not exceptional)
- ⚠️ Clarify equity offering

**Time spent:** 10 minutes
```

### Deep Analysis Stage (20-30 min)

```markdown
## Comprehensive Evaluation

### Role Analysis
- Senior Backend Engineer (Go)
- Building deployment pipeline
- Greenfield project (ideal)
- Technical ownership

### Culture Signals
- Async-first communication
- Written documentation culture
- Flexible hours
- No mandatory video

### Technology
- Go, Kubernetes, PostgreSQL (all good)
- AWS infrastructure
- Modern stack
- Some legacy Python services

### Compensation
- 18-22k PLN/mo (market rate)
- Equity: 0.1% (decent for Series B)
- Benefits: standard package
- Budget for conferences

## Final Recommendation

**Status:** ACCEPT - Apply
**Score:** 16/20 points
**Confidence:** Medium-High

**Strengths:**
- Greenfield project
- Good team size
- Modern tech stack
- Remote-first

**Concerns:**
- Salary not exceptional
- Some legacy maintenance likely

**Next steps:**
1. Apply via company site
2. Prepare questions about equity vesting
3. Ask about on-call expectations
4. Clarify greenfield vs maintenance split

**Time spent:** 25 minutes
**Total time:** 37 minutes across all stages
```

## Customization Points

**Adjust time limits:**
- Triage: 1-5 minutes (based on complexity)
- Verification: 5-20 minutes (based on claims)
- Deep: 15-60 minutes (based on depth needed)

**Adjust stage criteria:**
- Triage: Must-have requirements only
- Verification: Key claims to validate
- Deep: Comprehensive analysis areas

**Adjust filtering:**
- Strict: Reject on any triage failure
- Lenient: Pass with warnings to next stage
- Hybrid: Some failures are hard rejects, others are flags

**Adjust output:**
- Triage: One sentence (pass/reject)
- Verification: Table of results
- Deep: Full structured report

## Integration with Project Type

**CLI tools:** Quick test → Verify behavior → Deep code review
**Web apps:** Quick demo → Verify key features → Full audit
**API services:** Quick endpoint test → Verify responses → Load test
**Research:** Skim abstract → Verify key claims → Deep read
**Libraries:** Check docs → Verify API → Test integration

## Benefits

- Time-efficient (reject early)
- Progressive refinement
- Clear stage boundaries
- Predictable time investment
- Easy to parallelize (triage many, verify few)

## Anti-Patterns

**AVOID:**
- ❌ Skipping triage (wastes time on obvious rejects)
- ❌ Deep analysis without verification (risky)
- ❌ Vague triage criteria (everything passes)
- ❌ No time limits (defeats purpose)
- ❌ Same depth at all stages (inefficient)
