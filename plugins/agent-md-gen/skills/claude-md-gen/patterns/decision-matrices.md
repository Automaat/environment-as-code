# Decision Matrices Pattern

Hard/soft criteria with auto-reject conditions for systematic evaluation.

---

## When to Use

- Evaluating options against requirements
- Go/no-go decisions with clear criteria
- Scoring candidates (jobs, plots, libraries, etc.)
- Requirements validation
- Multi-criteria decision making

## Pattern Structure

### Components

1. **Hard Requirements** - Must-haves (auto-reject if missing)
2. **Soft Criteria** - Nice-to-haves (scored)
3. **Scoring System** - Point values for soft criteria
4. **Decision Threshold** - Minimum score to pass
5. **Auto-Reject Conditions** - Immediate disqualification

## Template

```markdown
## Hard Requirements (Must-Have)

- [ ] [Requirement 1]
- [ ] [Requirement 2]
- [ ] [Requirement 3]

**If any hard requirement fails: REJECT immediately**

## Soft Criteria (Scored)

### [Category 1]

- [+X pts] [Positive criterion]
- [+Y pts] [Positive criterion]
- [-Z pts] [Negative criterion]

### [Category 2]

- [+X pts] [Positive criterion]
- [+Y pts] [Positive criterion]
- [-Z pts] [Negative criterion]

## Auto-Reject Conditions

Immediate rejection if:
- [Condition 1]
- [Condition 2]
- [Condition 3]

## Scoring

**Maximum possible:** [MAX] points
**Minimum to pass:** [THRESHOLD] points
**Actual score:** [SCORE] points

**Result:** [PASS / REJECT]

## Decision

[Final recommendation based on score and criteria]
```

## Example 1: Job Listing Evaluation (from job-triage skill)

```markdown
## Hard Requirements

- [ ] Europe-based or remote-OK for Europe
- [ ] Poland work authorization OR sponsorship offered
- [ ] Technology match (Go, Python, distributed systems)

**Status:** ✅ All hard requirements met

## Soft Criteria

### Positive Signals (+points)

- [+3 pts] Greenfield project
- [+2 pts] Small team (< 20 people)
- [+2 pts] Modern tech stack
- [+1 pt] Remote-first culture
- [+1 pt] Flexible hours

### Negative Signals (-points)

- [-2 pts] Equity-only compensation
- [-2 pts] Enterprise/corporate environment
- [-1 pt] On-call rotation

## Auto-Reject Conditions

- ❌ Unpaid or "token" compensation
- ❌ Blockchain/crypto primary focus
- ❌ Required relocation outside Europe

## Scoring

**Maximum:** 15 points
**Threshold:** 5 points
**Actual:** 8 points (+3+2+2+1, -0)

**Result:** PASS ✅

## Decision

**Recommendation:** Worth applying
**Reason:** Meets all requirements, strong positive signals
**Next step:** Review company culture and team
```

## Example 2: Plot Evaluation (from plot-analysis skill)

```markdown
## Hard Requirements

- [ ] Within budget (max 300k PLN)
- [ ] Buildable (zoning allows residential)
- [ ] Utilities accessible (power, water, sewage)
- [ ] Legal (no ownership disputes)

**Status:** ✅ 4/4 met

## Soft Criteria

### Location (+/-points)

- [+5 pts] < 30min to Krakow center
- [+3 pts] Quiet area (no highways nearby)
- [-3 pts] Flood zone
- [-2 pts] High-traffic road

### Infrastructure (+/-points)

- [+3 pts] Fiber internet available
- [+2 pts] Paved road access
- [+1 pt] Public transport < 1km
- [-2 pts] Septic tank required (no sewage)

### Land Quality (+/-points)

- [+3 pts] Flat terrain (< 10% slope)
- [+2 pts] South-facing
- [+2 pts] Trees/mature vegetation
- [-3 pts] Contaminated soil

## Auto-Reject Conditions

- ❌ Disputed ownership
- ❌ No building permit possible
- ❌ Severe environmental hazards
- ❌ Price > 20% above market

## Scoring

**Maximum:** 21 points
**Threshold:** 10 points
**Actual:** 14 points (+5+3+3+2+1+2, -2)

**Result:** PASS ✅

## Decision

**Recommendation:** Strong candidate for deep due diligence
**Reason:** Meets requirements, good location, minor infrastructure gap
**Flag:** Septic tank requirement (cost +30k PLN, maintenance)
**Next step:** Legal verification, soil test, neighbor interviews
```

## Example 3: Library Selection

```markdown
## Hard Requirements

- [ ] Actively maintained (commit in last 6 months)
- [ ] Compatible license (MIT, Apache, BSD)
- [ ] Supports target platform (Linux, macOS, Windows)
- [ ] TypeScript/type definitions available

**Status:** ✅ 4/4 met

## Soft Criteria

### Quality (+points)

- [+3 pts] > 1k GitHub stars
- [+2 pts] Comprehensive documentation
- [+2 pts] Test coverage > 80%
- [+1 pt] CI/CD configured

### Community (+points)

- [+2 pts] Active maintainers (> 1)
- [+1 pt] Responsive to issues
- [+1 pt] Regular releases

### Concerns (-points)

- [-3 pts] Large bundle size (> 100KB)
- [-2 pts] Complex API
- [-1 pt] Brings many dependencies

## Auto-Reject Conditions

- ❌ Abandoned (no activity in 1+ year)
- ❌ Known security vulnerabilities
- ❌ Incompatible license
- ❌ No TypeScript support

## Scoring

**Maximum:** 12 points
**Threshold:** 6 points
**Actual:** 10 points (+3+2+2+1+2+1+1, -2)

**Result:** PASS ✅

## Decision

**Recommendation:** Use this library
**Trade-off:** Bundle size is moderate but acceptable
**Alternative considered:** [Other library] (scored 7)
```

## Customization Points

**Adjust categories:**
- Technical criteria
- Business criteria
- Risk factors
- Cost factors
- Timeline factors

**Adjust scoring:**
- Weight critical factors higher (+5 instead of +1)
- Add severity to negatives (-5 for dealbreakers)
- Use ranges for uncertainty (+1 to +3)

**Adjust thresholds:**
- Simple pass/fail: threshold = 0 (just count positives/negatives)
- Tiered: < 5 reject, 5-10 maybe, > 10 accept
- Percentage: score must be > 60% of maximum

**Adjust auto-reject:**
- Project-specific dealbreakers
- Regulatory requirements
- Technical constraints
- Budget limits

## Integration with Project Type

**CLI tools:** Evaluate flag designs, error handling approaches
**Web apps:** Choose libraries, frameworks, design patterns
**API services:** Select databases, auth methods, hosting
**Research:** Evaluate sources, plots, jobs, opportunities
**Libraries:** Select dependencies, design patterns

## Benefits

- Systematic evaluation
- Traceable decisions
- Objective criteria
- Easy to compare options
- Clear pass/fail

## Anti-Patterns

**AVOID:**
- ❌ Too many criteria (analysis paralysis)
- ❌ All criteria weighted equally (no priorities)
- ❌ Ignoring auto-reject conditions
- ❌ Changing criteria mid-evaluation
- ❌ No threshold (unclear when to accept)
