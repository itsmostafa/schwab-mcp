# Activity Log

Append-only. Format: `## [YYYY-MM-DD] type | description`

---

## [2026-04-25] plan | 36-hour compressed prep plan created for back-to-back Rounds 3 & 4

Designed a memory-retention-optimized study plan for ~36-hour window with both rounds same day. Key insight from audit: wiki content is strong; the gap is rehearsed retrieval under pressure. Plan built around 3 anti-blank scaffolds (Universal Opener, Triage Rosary, Systems Design 6-Step), 4 study blocks (~7 hrs total), two protected sleep nights, and a deliberate between-round reset protocol. All study is retrieval-based (no re-reading). Personal system story draft is the single highest-priority Block A action.

---

## [2026-04-25] resource | Guided study prompt created for Rounds 3 & 4

Created wiki/study-prompt.md — a self-contained paste-ready prompt for any AI chat session. Sets up Socratic quiz mode, mock interview mode, whiteboard narration, and weak-spots drill. Embeds the triage agent anchor design, all 10 AI Integration topics, all 5 systems design scenarios, and explicit rubrics for strong vs weak answers for both rounds.

---

## [2026-04-25] plan | Prep plan created for Round 3 (AI Integration) and Round 4 (Systems Design)

Created wiki/prep-plan.md with: per-round breakdown of what's being evaluated, anchor alert triage agent design, topic-by-topic prep for both rounds, 5-technique retention strategy, 4-day session plan, and day-of checklists for both interviews.
Updated wiki/index.md to reference the new page.

---

## [2026-04-25] query | AI Integration interview preparation

Reviewed wiki/ for the upcoming AI Integration interview. Expanded rounds/ai-integration.md with a concrete preparation plan, canonical alert triage agent design, token-budget/RAG answer, prompt architecture answer, production metrics, security guardrails, personal system story structure, and strong answer framework.

## [2026-04-24] ingest | candidate-prep-guide-staff-ai-engineer-soc-agent-platform.pdf

Live coding round completed. Ingested official Panther prep guide for the 4 remaining main loop interviews.
Updated: overview.md (live coding marked DONE, AI Integration now NEXT, priorities reordered, Key Themes Across All Interviews table added)
Updated: rounds/ai-integration.md (format, 7 official signals, official prep checklist)
Updated: rounds/systems-design.md (format, 6 official signals, official prep guidance)
Updated: rounds/project-retro.md (format, 6 official signals, official project selection criteria, come prepared list)
Updated: rounds/culture.md (format, 6 official signals incl. Grit & Ambiguity / Startup Mindset / Self-Awareness, story prep format)
Updated: index.md (live coding marked DONE)

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
