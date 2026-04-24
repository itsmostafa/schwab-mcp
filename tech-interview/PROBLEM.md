# Interview Problem — StackExchange Question Search

> **Start your timer before reading further.**

---

## Background

The team is building an internal knowledge tool that surfaces relevant StackOverflow Q&A for a given topic. Your task is to implement the backend search function that queries the StackExchange API.

---

## Task

Implement a Python function `search_questions` that queries the StackExchange API v2.3 `GET /search` endpoint and returns a structured list of matching questions.

### Function Signature

```python
def search_questions(
    query: str,
    site: str = "stackoverflow",
    tags: list[str] | None = None,
    max_results: int = 30,
) -> list[dict]:
    ...
```

### Requirements

1. **Search**: Use the `intitle` parameter to search by title text
2. **Tag filtering**: When `tags` is provided, pass them as the `tagged` parameter (semicolon-delimited)
3. **Pagination**: Collect results across pages using the `has_more` field; stop when you have `max_results` or there are no more pages
4. **Rate limiting**:
   - If the response includes a `backoff` field, wait that many seconds before the next request
   - Handle HTTP 429 responses gracefully (raise a clear exception — do not silently drop)
5. **Error handling**: Raise a descriptive exception for:
   - HTTP errors (4xx, 5xx)
   - API-level errors (response body contains `error_id`)
   - Network failures (timeout, connection error)
6. **Output shape**: Each returned dict must include at minimum:
   - `title` (str)
   - `link` (str)
   - `score` (int)
   - `answer_count` (int)
   - `is_answered` (bool)
   - `tags` (list[str])

### Constraints

- Base URL: `https://api.stackexchange.com/2.3`
- No API key — unauthenticated requests only (300 req/day quota)
- Python 3.11+
- Only `requests` and stdlib are available

---

## Clarifying Questions to Consider

Think through (or ask) these before starting:

- What should happen if `query` is an empty string?
- What if `max_results` is 0 or negative?
- Should the function return exactly `max_results` items, or up to `max_results`?
- Should pagination be bounded (e.g., max pages) to avoid runaway loops?

---

## After Implementation

Be ready to discuss:

1. How would you test this without hitting the real API?
2. What happens under high load if 100 callers hit this simultaneously?
3. How would you add caching? What's the cache key?
4. How would you make this production-ready (observability, retries, circuit breaker)?
