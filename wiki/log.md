# Activity Log

Append-only. Format: `## [YYYY-MM-DD] type | description`

---

## [2026-04-27] resource | CSET AI chatbot interview story created

Researched private GitHub repo `itsmostafa/cset-ai-chatbot` and CISA CSET public context. Created wiki/projects/cset-ai-chatbot-story.md with a Round 3 AI Integration talk track, 60-second version, architecture anchor, tradeoffs, likely follow-up answers, Panther bridge, and overclaiming guardrails. Updated wiki/index.md, wiki/rounds/ai-integration.md, and wiki/rounds/project-retro.md.

---

## [2026-04-27] update | CSET story NIST document context added

Updated wiki/projects/cset-ai-chatbot-story.md to clarify that the uploaded knowledge-base PDFs were NIST-related cybersecurity guidance intended to help public and private organizations improve cybersecurity posture.

---

## [2026-04-27] resource | Tool calling cheat sheet created

Created wiki/concepts/tool-calling-cheat-sheet.md as a concise recall sheet for tool-call mechanics, tool schema anatomy, read/write tool boundaries, SOC tool inventory, handler rules, parallel calls, failure modes, security guardrails, testing strategy, and interview answer templates. Updated wiki/index.md and linked it from wiki/concepts/tool-calling.md.

---

## [2026-04-27] resource | Round 4 systems design cheat sheet created

Created wiki/rounds/round-4-cheat-sheet.md as a final-scan guide for the Round 4 Systems Design interview. Covers the universal opener, 6-step interview flow, Panther-style alert triage anchor architecture, scenario cards, component deep dives, tradeoff bank, failure modes, metrics, useful scale numbers, red flags, and strong closing questions. Updated wiki/index.md and wiki/rounds/systems-design.md to link the new page.

---

## [2026-04-27] resource | Round 3 AI Integration cheat sheet generated

Created wiki/rounds/round-3-cheat-sheet.md as a last-minute interview reference for the AI Integration round. Condenses the full Round 3 guide into answer frames, canonical alert-triage architecture, RAG/token-budget handling, prompt architecture, tool-calling guardrails, AI security, evaluation metrics, feedback loops, Panther-specific hooks, likely questions, and red flags. Updated wiki/index.md and linked it from wiki/rounds/ai-integration.md.

---

## [2026-04-27] update | CSET chatbot interview version converted to quick-reference outline

Reworked wiki/projects/cset-ai-chatbot-story.md so the Interview Version is no longer a long spoken script. It now highlights the opening hook, architecture, ingestion/query paths, design principle, tradeoffs, production-minded pieces, next improvements, and the strongest closing line for quick interview recall.

---

## [2026-04-26] query | Prep plan method audit

Audited wiki/prep-plan.md against the official candidate prep guide, job description, referral notes, and round-specific wiki pages. Verdict: close to the best practical 36-hour plan for AI Integration and Systems Design because it prioritizes active retrieval, out-loud mocks, a canonical alert triage design, tradeoffs, failure modes, and sleep. Added wiki/audits/prep-plan-method-audit.md with remaining gaps: Panther questions drill, one interviewer-assisted mock, stronger personal story emphasis, no live AI-generated answers, and a non-triage systems-design mock.

---

## [2026-04-26] audit | AI/system design concepts externally validated

Created wiki/audits/ai-system-design-validation.md with a primary-source claim register for AI Integration and Systems Design material. Validated Panther role/product claims, RAG/tool-calling/MCP/prompt-injection/vector-index concepts, and systems-design heuristics. Corrected overstatements around auto-close confidence, fixed latency/vector-count thresholds to be explicit assumptions, clarified RAG vs fine-tuning, MCP, tool-calling execution boundaries, and structured reasoning vs chain-of-thought.

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
