# Mock Interview — StackExchange API

**Interview date**: Friday, April 25, 2026  
**Format**: CoderPad IDE, 60 minutes, Zoom  
**Language**: Python

This directory is your practice environment. Run the mock as close to real conditions as possible.

---

## How to Run a Mock Session

1. Close all tabs except this directory and the [SE API docs](https://api.stackexchange.com/docs/search)
2. Set a 60-minute timer before opening `PROBLEM.md`
3. Read `PROBLEM.md` cold — pretend you've never seen it
4. Work in `solution.py` and `test_solution.py` only
5. Talk out loud throughout (or write reasoning as comments you'll delete)
6. Stop at 60 minutes regardless of where you are
7. Self-score using the rubric at the bottom of this file

---

## Rules (Mirrors Real Interview)

| Allowed | Not Allowed |
|---------|-------------|
| SE API docs (live endpoint) | Claude Code / Cursor / ChatGPT |
| Python stdlib + `requests` | Any other external resources |
| CoderPad AI Assist equivalent: ask yourself what you'd prompt | Looking up code solutions |

---

## Setup

```bash
cd tech-interview
python -m venv .venv
source .venv/bin/activate
pip install requests pytest pytest-mock
```

Run tests:

```bash
pytest test_solution.py -v
```

---

## Session Flow (follow this timing)

| Timestamp | Phase | What to do |
|-----------|-------|------------|
| 0:00–0:05 | Read + Clarify | Read `PROBLEM.md`. Write down your clarifying questions. |
| 0:05–0:15 | Design | Talk through (out loud): function signature, data structures, pagination strategy, error handling plan. Don't write code yet. |
| 0:15 | Align | If real interview: confirm with interviewer. Here: write a 3-line summary of your plan in `solution.py` as comments. |
| 0:15–0:45 | Execute | Implement. Narrate as you go. |
| 0:45–0:55 | Test | Fill in `test_solution.py`. Run pytest. |
| 0:55–1:00 | Debrief | Walk through what you'd improve, production considerations, tradeoffs. |

---

## Self-Evaluation Rubric

After each mock, score yourself 1–3 on each dimension:

| Dimension | 1 — Needs work | 2 — Solid | 3 — Strong |
|-----------|----------------|-----------|------------|
| **Communication** | Long silences, didn't explain approach | Explained approach, some narration | Narrated throughout, no silent gaps |
| **Collaboration** | Didn't ask clarifying questions | Asked 1–2 questions | Surfaced all ambiguity upfront |
| **Coding Fluency** | Messy, hard to follow | Clean structure, readable | Clean + used helpers well |
| **Data Structures** | Wrong or naive choice | Reasonable choice | Optimal + justified out loud |
| **Scalability** | Didn't mention | Mentioned pagination/rate limits | Discussed at volume, retries, backoff |
| **Testing** | No tests | Basic happy path | Happy path + edge cases + errors |
| **Feature Completeness** | Missing requirements | Met core requirements | All requirements + bonus edge cases |

Target: all 2s before Friday, at least two 3s.

---

## Key Things to Nail

- **Say what you're doing before you do it** — this is the #1 differentiator
- **Respect `backoff`** — interviewers specifically look for this
- **Handle API-level errors** — `error_id` can appear even in a 200 response
- **Pagination via `has_more`** — don't assume single page
- **Test the sad paths** — 429, 400, network timeout, empty results

See `wiki/rounds/live-coding.md` and `wiki/concepts/stackexchange-api.md` for full reference.
