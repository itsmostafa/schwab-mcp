# Activity Log

Append-only. Format: `## [YYYY-MM-DD] type | description`

---

## [2026-04-24] update | Scalability debrief section added to live-coding.md

Implemented TDD-driven improvements in tech-interview/solution.py:
- TTL caching (stdlib dict + threading.Lock, 5-min TTL, order-insensitive tuple key)
- Smarter pagination (pagesize capped at 25)
Added wiki/rounds/live-coding.md section: "Scalability Debrief — How Would You Scale This?"
Covers performance (caching, page size, async), reliability (retries, thread safety, bounded pagination), edge cases (empty query, quota exhaustion, all-unanswered results)
All 18 tests passing (13 original + 5 new).

---

## [2026-04-23] setup | Mock interview environment created in tech-interview/

Created: tech-interview/README.md, PROBLEM.md, solution.py, test_solution.py
Covers: realistic problem prompt, 60-min session flow, self-eval rubric, full pytest scaffold (happy path, pagination, rate limiting, API errors, network failures)

---

## [2026-04-23] query | Tool calling for AI agents — Darwayne tip for separate interview

Created: wiki/concepts/tool-calling.md
Covers: protocol mechanics, Anthropic API implementation, agent loop pattern, parallel tool calls, SOC tool inventory, design principles, failure modes, interview talking points

---

## [2026-04-22] init | Wiki initialized from raw sources

Ingested: job-description.md, additional-interview-info.md, backend-paired-coding-interview.pdf
Created: full wiki structure, CLAUDE.md schema, all initial pages
Priority set: live coding round (Friday April 25) is most urgent
