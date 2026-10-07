# Agent workflow failure corpus

This corpus turns recurring failures from local Claude Code, Codex, OpenCode, and multi-agent sessions into safe, repeatable workflow evaluations. Cases preserve the failure's control flow without copying prompts, source code, repository names, credentials, or tool output.

`cases.json` contains the corpus. `schema.json` defines its machine-readable contract. Each case records a synthetic task, covered harnesses, initial state, required behavior, forbidden behavior, grading observations, and privacy classification.

A case passes only when every required observation is present and every forbidden observation is absent. An unavailable harness is skipped, never passed.

## Privacy rules

- Use synthetic names, repositories, issue numbers, revisions, commands, and output.
- Preserve only the control-flow property needed to reproduce the failure.
- Never store prompt transcripts, source excerpts, filesystem paths, credentials, tokens, customer data, or private URLs.
- Mark transformed historical patterns as `synthetic-redaction` and public only after manual review.
- Reject fields outside the schema rather than retaining unclassified data.

## Adding a case

Add one independently gradable failure. Include first-run and error behavior when they affect the result. Keep identifiers stable because result histories refer to them.

Validate the corpus locally:

```sh
evals/agent-workflows/validate.sh
```

Each evaluation runner must validate the schema before execution so every harness receives the same input.
