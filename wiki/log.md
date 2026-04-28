# Activity Log

Append-only. Format: `## [YYYY-MM-DD] type | description`

---

## [2026-04-28] update | Round 4 cheat sheet table of contents added

Added a clickable table of contents near the top of `wiki/rounds/round-4-cheat-sheet.md` so the main Systems Design cheat-sheet sections are easier to jump to.

---

## [2026-04-28] update | Failure-mode brush-up added

Added a drill-oriented failure-mode brush-up section to `wiki/rounds/round-4-cheat-sheet.md`, including the answer shape, fast scenario prompts, a one-minute spoken drill, and a memory hook. Updated the index hook for the cheat sheet.

---

## [2026-04-28] update | Round 4 cheat sheet made glanceable

Reworked `wiki/rounds/round-4-cheat-sheet.md` into a faster quick-reference format that leans on `raw/round-4-cheat-sheet.md`: 5-second map, compact component cards, scenario cards, failure counters, tradeoffs, metrics, scaling, safe autonomy, and interview-ready one-liners.

---

## [2026-04-28] cleanup | Completed-round wiki material removed

Removed completed technical-design-round prep material from the maintained wiki. Deleted the dedicated completed-round guide, question bank, cheat sheet, and obsolete prep-method audit. Reworked the active prep plan and guided study prompt around Systems Design. Updated the index, overview, cross-links, project story metadata, and concept pages so the wiki points to the next active round.

---

## [2026-04-28] cleanup | Project retrospective prep removed

Marked Project Retrospective complete and removed the dedicated round prep page from the maintained wiki. Updated the index, overview, repo instructions, and CSET story metadata so the remaining active prep is Systems Design, Culture, and CEO-facing material.

---

## [2026-04-27] ingest | Systems design interview questions

Ingested `raw/round-4-interview-questions.md` into the maintained wiki question bank. Created `wiki/rounds/round-4-interview-questions.md` with spoken-answer drills for Systems Design. Updated `wiki/index.md` and `wiki/rounds/systems-design.md`.

---

## [2026-04-28] update | Circuit breaker section added

Updated `wiki/rounds/round-4-cheat-sheet.md` with a Circuit Breakers component card covering dependency protection, breaker states, trip signals, fallbacks, recovery, and a spoken threat-intel timeout example. Updated `wiki/index.md`.

---

## [2026-04-28] update | Round 4 cheat sheet TOC expanded

Updated `wiki/rounds/round-4-cheat-sheet.md` table of contents to include missing component, scenario, and failure drill subsections, including Backpressure and Multi-Tenancy + Isolation.

---

## [2026-04-28] update | Backpressure and multi-tenancy sections added

Updated `wiki/rounds/round-4-cheat-sheet.md` with component cards for handling backpressure and multi-tenancy/isolation, including triggers, mitigations, tenant-scoped retrieval/cache guidance, and spoken interview examples. Updated `wiki/index.md`.

---

## [2026-04-28] update | Idempotency key section added to Round 4 cheat sheet

Updated `wiki/rounds/round-4-cheat-sheet.md` with an Idempotency + Idempotency Keys component card covering key shape, storage, write flow, side effects, TTL, failure modes, and a spoken Panther alert duplicate example. Updated `wiki/index.md`.

---

## [2026-04-27] resource | CSET AI chatbot interview story created

Researched private GitHub repo `itsmostafa/cset-ai-chatbot` and CISA CSET public context. Created `wiki/projects/cset-ai-chatbot-story.md` with an interview talk track, 60-second version, architecture anchor, tradeoffs, likely follow-up answers, Panther bridge, and overclaiming guardrails. Updated `wiki/index.md`.

---

## [2026-04-27] update | CSET story NIST document context added

Updated `wiki/projects/cset-ai-chatbot-story.md` to clarify that the uploaded knowledge-base PDFs were NIST-related cybersecurity guidance intended to help public and private organizations improve cybersecurity posture.

---

## [2026-04-27] resource | Tool calling cheat sheet created

Created `wiki/concepts/tool-calling-cheat-sheet.md` as a concise recall sheet for tool-call mechanics, tool schema anatomy, read/write tool boundaries, SOC tool inventory, handler rules, parallel calls, failure modes, security guardrails, testing strategy, and interview answer templates. Updated `wiki/index.md` and linked it from `wiki/concepts/tool-calling.md`.

---

## [2026-04-27] resource | Systems design cheat sheet created

Created `wiki/rounds/round-4-cheat-sheet.md` as a final-scan guide for the Systems Design interview. Covers the universal opener, 6-step interview flow, Panther-style alert triage anchor architecture, scenario cards, component deep dives, tradeoff bank, failure modes, metrics, useful scale numbers, red flags, and strong closing questions. Updated `wiki/index.md` and `wiki/rounds/systems-design.md`.

---

## [2026-04-27] update | CSET chatbot interview version converted to quick-reference outline

Reworked `wiki/projects/cset-ai-chatbot-story.md` so the Interview Version is no longer a long spoken script. It now highlights the opening hook, architecture, ingestion/query paths, design principle, tradeoffs, production-minded pieces, next improvements, and the strongest closing line for quick interview recall.

## [2026-04-26] audit | System design concepts externally validated

Created `wiki/audits/ai-system-design-validation.md` with a primary-source claim register for role/product claims, retrieval, tool calling, MCP, prompt-injection risk, vector-index concepts, and systems-design heuristics. Corrected overstatements around auto-close confidence, fixed latency/vector-count thresholds to be explicit assumptions, clarified RAG vs fine-tuning, MCP, tool-calling execution boundaries, and structured reasoning vs chain-of-thought.

---

## [2026-04-25] plan | Compressed prep plan created for technical design interviews

Designed a memory-retention-optimized study plan for a compressed interview window. Key insight: the wiki content was strong; the gap was rehearsed retrieval under pressure. Plan built around active recall, out-loud mocks, a canonical alert triage design, tradeoffs, failure modes, and sleep.

---

## [2026-04-25] resource | Guided study prompt created

Created `wiki/study-prompt.md` as a self-contained paste-ready prompt for any AI chat session. Sets up Socratic quiz mode, mock interview mode, whiteboard narration, and weak-spots drill.

---

## [2026-04-24] ingest | candidate-prep-guide-staff-ai-engineer-soc-agent-platform.pdf

Live coding round completed. Ingested official Panther prep guide for the remaining main loop interviews.
Updated: `overview.md` with current round status and priorities.
Updated: `rounds/systems-design.md` and `rounds/culture.md` with official signals and prep guidance.
Updated: `index.md`.

---

## [2026-04-24] update | Scalability debrief section added to live-coding.md

Implemented TDD-driven improvements in `tech-interview/solution.py`:
- TTL caching with stdlib dict and `threading.Lock`, 5-minute TTL, order-insensitive tuple key
- Smarter pagination with `pagesize` capped at 25

Added `wiki/rounds/live-coding.md` section: "Scalability Debrief — How Would You Scale This?" Covers performance, reliability, and edge cases. All 18 tests passing.

---

## [2026-04-23] setup | Mock interview environment created in tech-interview/

Created: `tech-interview/README.md`, `PROBLEM.md`, `solution.py`, `test_solution.py`. Covers a realistic problem prompt, 60-minute session flow, self-eval rubric, and full pytest scaffold.

---

## [2026-04-23] query | Tool calling for AI agents

Created `wiki/concepts/tool-calling.md`. Covers protocol mechanics, Anthropic API implementation, agent loop pattern, parallel tool calls, SOC tool inventory, design principles, failure modes, and interview talking points.

---

## [2026-04-22] init | Wiki initialized from raw sources

Ingested `raw/job-description.md` and `raw/additional-interview-info.md`. Created the initial wiki structure, schema, and pages.
