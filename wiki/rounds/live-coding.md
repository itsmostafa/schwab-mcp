# Live Coding Round

**Date**: Friday, April 25, 2026
**Duration**: 60 minutes via Zoom
**Environment**: CoderPad IDE + CoderPad AI Assist ONLY
**Problem**: StackExchange API v2.3 — GET /search endpoint

> This is the most time-sensitive round. Full problem statement shared at interview start.

---

## What They're Evaluating (7 Dimensions)

| Dimension | What It Means |
|-----------|---------------|
| 🗣 Communication | Explain approach before starting. Narrate decisions. No long silences. |
| 🤝 Collaboration | Ask clarifying questions upfront. Treat the interviewer as a partner. |
| ⚡ Coding Fluency | Readable, well-structured code. Use AI Assist effectively — prompt well, review critically. |
| 🗂 Data Structure | Choose appropriate data structures for the problem. |
| 📈 Scalability | Think about scale: performance, reliability, edge cases at volume. |
| 🧪 Debugging/Testing | Structure for testability. Write unit tests or document test scenarios. |
| ✅ Feature Completeness | Meet all requirements, return required fields, finish on time. |

---

## Session Flow (60 Minutes)

| Time | Phase | What to Do |
|------|-------|------------|
| 0–5 min | **Intro/Read** | Read problem carefully. Confirm understanding of input, output format, objectives. Ask questions. |
| 5–15 min | **Discuss** | Talk through approach BEFORE writing a line. Walk through data structures, strategy, test plan. |
| 15 min | **Align** | Confirm you and interviewer are aligned before starting to code. |
| 15–45 min | **Execute** | Implement. Narrate thought process. Use AI Assist, but review every line before accepting. |
| 45–55 min | **Test** | Write unit tests. If short on time, write detailed test scenario notes instead. |
| 55–60 min | **Debrief** | Walk through solution. Discuss improvements, production considerations, tradeoffs. |

---

## Rules

- ✅ CoderPad IDE
- ✅ CoderPad AI Assist (can change model in AI agent tab)
- ✅ StackExchange API docs (the live endpoint)
- ❌ External IDEs (Claude Code, Cursor, Copilot, ChatGPT, etc.)
- ❌ Any online resources outside CoderPad + SE API docs

**Any programming language is allowed** — switch in the left panel. Go or Python recommended given Mostafa's strength.

---

## The Problem: StackExchange API v2.3 GET /search

Full instructions shared at interview start. Based on the prep guide:
- Build a backend function using `GET /search`
- Show it works, discuss edge cases, discuss scalability considerations
- Think about: error responses, empty result sets, pagination, rate limiting, compression

See [[concepts/stackexchange-api]] for full API reference and practice scenarios.

---

## Strategy for the Session

### Before starting to code

Ask these questions:
- What should the function signature look like? What are the inputs?
- What's the expected output format/fields?
- What's the error handling expectation?
- Should I handle pagination?
- Any rate limiting considerations?
- What language do you prefer / is the team comfortable reviewing?

### During execution

- State what you're about to do before doing it
- When using AI Assist: say "I'm going to prompt AI Assist for the HTTP client setup, then review it"
- When accepting AI output: explicitly call out what you checked ("looks good, it handled the error status codes correctly")
- When rejecting/modifying AI output: say why ("AI Assist didn't handle the 400 bad request case, I'm adding that")

### Clarify AI usage upfront

Ask the interviewer at the start: "What's your preference on AI Assist usage — do you want to see me use it heavily, or demonstrate more without it first?" This signals self-awareness.

---

## Common Mistakes to Avoid

- 🤐 **Coding in silence** — most common way to lose points
- 🧪 **No testing at all** — even notes on scenarios counts; silence does not
- 🤷 **Assuming instead of asking** — use 0-5 min to surface all ambiguity
- 🤖 **Rubber-stamping AI output** — own every line; "it looked right" is a red flag to interviewers
- ⏱ **Running out of time** — time-box execution to 30 min; leave buffer for testing

---

## Practice Checklist

- [ ] Read the StackExchange API v2.3 docs, especially GET /search
- [ ] Write a working GET /search function in Python or Go
- [ ] Handle: 200 OK, 400, 404, 429 (rate limit), 500, network errors
- [ ] Handle: empty results, pagination (page + pagesize params)
- [ ] Handle: gzip compression (SE API returns compressed responses by default)
- [ ] Write unit tests for the above scenarios
- [ ] Rehearse talking through approach out loud before coding
- [ ] Practice narrating while coding (say what you're doing as you do it)
