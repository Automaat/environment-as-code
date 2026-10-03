---
name: kong-ai-intel
description: Search Kong Slack for new messages about agentic AI, Claude Code, AI agents, and coding-agent tooling since the last run, and produce a short digest with permalinks, filtered for kreview/ksai relevance, internal AI adoption, and building agentic tools. Use when asked to "check Slack for AI intel", "run the daily AI digest", "what's new in agentic AI at Kong", or invoked as /kong-ai-intel.
allowed-tools: Bash Read Write ToolSearch Agent mcp__plugin_slack_slack__slack_search_public_and_private mcp__plugin_slack_slack__slack_search_channels mcp__plugin_slack_slack__slack_read_thread
---

# Kong AI/Agentic Intel Digest

Goal: find new Kong Slack activity about agentic AI, Claude Code, and AI agents since the last run. Report it short. Filter for what matters to Marcin: kreview/ksai work, driving AI adoption inside Kong, and building agentic tools (ours or the industry's). Drop routine AI Gateway product/support traffic even when it mentions "LLM" or an agent keyword — that's a different domain and floods the digest (tested and confirmed noisy: 2026-08-07).

## Arguments

Read them from the user's request: `--since YYYY-MM-DD` and `--dry-run`, both optional.

## State file

Path: `${XDG_STATE_HOME:-$HOME/.local/state}/kong-ai-intel/last_run.txt`. Resolve it with Bash. It holds one line: the last run date, `YYYY-MM-DD`, UTC. It lives outside the skill directory, which may be a read-only plugin cache or a symlink into a git repo.

- No file yet → use the date from the old unmanaged copy's state, the first that exists of `~/.claude/skills/kong-ai-intel/last_run.txt` and `~/.claude/skills.pre-zakwas.bak/kong-ai-intel/last_run.txt`; otherwise use 2 days back as the start date.
- `--since <date>` → use that date as the start, for this run only.
- On a normal (non-`--dry-run`) run that completes, overwrite the file with today's date.

## Step 1 — Load tools and dates

1. The Slack tools are `slack_search_public_and_private`, `slack_search_channels` and `slack_read_thread` from the Slack MCP server. In Claude Code they are `mcp__plugin_slack_slack__<tool>`; other agents expose them under their own prefix. If they are deferred, load them in one call:
   `ToolSearch(query: "select:mcp__plugin_slack_slack__slack_search_public_and_private,mcp__plugin_slack_slack__slack_search_channels,mcp__plugin_slack_slack__slack_read_thread", max_results: 5)`
2. Get today's date with Bash — do not guess it: `date -u +%Y-%m-%d`.
3. Read the state file for the start date (or apply the defaults above).

## Step 2 — Search

Run these queries in parallel against `slack_search_public_and_private`. Each query: append `after:<start> before:<tomorrow> -in:ask-agentic-engineering -in:ask-ai-working-group` to the query string, `sort: timestamp`, `sort_dir: desc`, `include_bots: false`, `response_format: detailed`, `include_context: false`, `limit: 20`.

Use `response_format: detailed`, not `concise` — only `detailed` returns the `Permalink` field, which Step 4 needs.

A page holds at most 20 results. While a query's page is full and returns a next cursor, call it again with that `cursor` until a page comes back short or without a cursor, so the whole interval is read. Track whether every query finished this way; Step 5 depends on it.

- `"Claude Code"`
- `agentic`
- `"AI agent"`
- `"agentic AI"`
- `kreview`
- `ksai`

Skip `"coding agent"` — tested empty historically, drop it unless the user asks to widen scope.

Channels `ask-agentic-engineering` and `ask-ai-working-group` are excluded from every query (`-in:` modifier above) — Marcin already sees everything there directly, so it's never new signal for this digest.

If any single call errors with "exceeds maximum allowed tokens", it saved its output to a file — spawn a background subagent (`Agent`, general-purpose, in Claude Code; without subagents, do the same reading yourself) with the file path and this exact instruction: slice-read the file in ~80,000-char spans via Python until it's fully read, then return only the items not already covered by the other queries' results (listed inline in the prompt), each with channel, author, date, one-line summary, and permalink. Keep the agent's report under 300 words.

## Step 3 — Merge, dedupe, classify

1. Dedupe by `message_ts` — the same message often matches more than one query.
2. Drop noise:
   - Routine bot digests (security/CVE feeds, dependency-advisory bots, PR-merge notices).
   - Keyword collisions unrelated to agentic AI (e.g. "Datadog agent", generic "AI" mentions with no agent/Claude Code/MCP tie-in).
   - Routine AI Gateway product/support traffic: model-routing configs, plugin bugs, provider API quirks, customer LLM-proxy setup questions. This dominates channels like `#ask-ai-gateway`, `#team-ai-gateway-manager`, `#ai-gateway-2-testing-feedback`, and customer `#internal-*` channels, and it will match "agentic"/"AI agent" purely on keyword collision. Keep a message from these channels only if it's specifically about agent identity/auth, agentic traffic governance, or an agentic-tooling capability — not routine model-proxying work.
3. Bucket what's left. Test each surviving item against: *would this help drive AI adoption inside Kong, or does it touch building agentic tools (ours or the industry's)?* If neither, drop it even if it matched a query term.
   - **kreview/ksai-relevant** — mentions ksai, kreview, AI-review agents/reviewers, or a direct ask related to that work.
   - **Internal AI adoption** — Kong people/teams adopting AI coding agents or agentic workflows internally: usage stories, hackathons/pilots, internal tooling built to support adoption (e.g. SE Harness, KODE, muthur-style plugins), blockers to adoption (auth, rate/session limits, cost).
   - **Building agentic tools** — work on agent-facing capabilities, ours or the industry's: MCP servers, Catalog Agents, agent identity/auth, agent plugin formats, agent security/governance controls, competitor agent-gateway or AMP moves that inform what to build next.
   - Anything left over that's genuinely new and on-topic but doesn't fit any bucket — drop it rather than force one.

## Step 4 — Report

One line per kept item: `permalink — channel — one-line summary`. Lead with the permalink so the report is a scannable, copyable list of thread links. Group by bucket; skip empty buckets. Always use plain permalinks (not markdown links) — never suppress this even if the user asked for markdown links elsewhere in the conversation, since the point of this report is one-click copying.

Write the report in ASD-STE100 Simplified Technical English: short sentences, one idea per sentence, plain approved words, active voice.

If nothing new turned up, say that in one line. Do not pad the report with old or low-value items to look productive.

## Step 5 — Update state

Unless `--dry-run` was passed, write today's date (from Step 1) to the state file, creating its directory first: `mkdir -p "${XDG_STATE_HOME:-$HOME/.local/state}/kong-ai-intel"`.

Write it only when every query in Step 2 succeeded and was read to its last page. If any query failed, or its results could not all be read, keep the old state file and say in the report that the run was incomplete and which queries were affected, so the next run covers the same interval again.

## Notes

- Slack access here comes through the Slack plugin's MCP server (OAuth). That auth may not carry over to a scheduled/headless run (cron, `/schedule`). If a search call returns an auth error instead of results, say so plainly — do not report "nothing new" when the real cause was no Slack access.
- There is no "list all channels" API here — only keyword search across channels the user already belongs to. A topic with no matching keyword will not surface.
- If Marcin adds recurring topics to track (a competitor name, a specific project), add them as extra query terms in Step 2 rather than starting a new skill.
