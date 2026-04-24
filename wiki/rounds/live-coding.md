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
- Should the function return exactly `max_results` items, or up to `max_results`?
- What should happen if the query returns zero results — empty list, exception, or something else?
- Is there a maximum number of pages I should fetch, or should I exhaust `has_more` until I hit `max_results`?
- Should I filter to only answered questions, or return all results including unanswered?
- Are there any constraints on dependencies — stdlib only, or can I bring in third-party libraries?
- Should I handle the `backoff` field in the API response, or assume the caller manages rate limiting?
- Is caching a requirement here, or a nice-to-have I should mention in the debrief?
- Should the function be synchronous, or is there interest in an async version?

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

## Scalability Debrief — "How Would You Scale This?"

If asked about performance, reliability, or edge cases at scale, hit these in priority order. Each has a working implementation in `tech-interview/solution.py` you can reference.

### Performance

| Option | What to say | Tradeoff |
|--------|-------------|----------|
| **TTL caching** | "Identical queries hit the API at most once per 5 minutes. Cache key includes query, site, tags, and max_results as a sorted tuple so it's order-insensitive." | Stale data; fine for SE which changes slowly. TTL is tunable. |
| **Smarter page size** | "Cap pagesize at 25 instead of passing max_results directly. Avoids over-fetching quota on the first page when max_results is large." | More pages for large result sets, but smaller per-request cost. |
| **Async concurrent fetching** | "For multi-site search, fan out requests with `asyncio` + `aiohttp` or `ThreadPoolExecutor`. Useful when the same query needs results from stackoverflow, serverfault, etc." | SE rate limit (30 req/s) still applies — need a semaphore to bound concurrency. |

### Reliability

| Option | What to say | Tradeoff |
|--------|-------------|----------|
| **Exponential backoff + retry** | "Retry on 429 and 5xx with jitter: `wait = base * 2^attempt + random(0, 1)`. Respect the `backoff` field too — SE tells you exactly how long to wait." | Need a max-retries cap to avoid infinite loops. |
| **Thread-safe cache** | "Wrap cache reads and writes in a `threading.Lock`. The check-then-write is not atomic without it — two threads can race to fetch the same uncached key." | Already implemented. Lock is held only briefly; no performance issue at this scale. |
| **Bounded pagination** | "Add a max_pages guard (e.g. 20) so a runaway `has_more=True` loop can't exhaust quota. Return what you have and log a warning." | Truncates results for very broad queries, but prevents quota burn. |

### Edge Cases

| Case | What to say |
|------|-------------|
| **Empty query** | "Validate upfront — SE returns results for `intitle=` (blank), which is meaningless. Raise `ValueError` before making any request." |
| **max_results ≤ 0** | "Raise `ValueError` or return `[]` immediately — don't hit the API." |
| **All results unanswered** | "The `is_answered` filter can cause the loop to exhaust all pages and return fewer than max_results. That's correct behavior — document it." |
| **API quota exhausted** | "SE returns `quota_remaining: 0` in every response body. Check it and raise early rather than burning retries." |
| **`backoff` on final page** | "Still sleep — SE penalizes clients that ignore backoff even when they have their data. Skip it and you risk a temporary ban." |

### One-liner answer if time is short

> "I'd add TTL caching with a hashable key, cap the page size at 25, add bounded retries with exponential backoff on 5xx/429, and guard against empty query and infinite pagination. For multi-site use cases, `ThreadPoolExecutor` lets you fan out concurrently without adding a new dependency."

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
