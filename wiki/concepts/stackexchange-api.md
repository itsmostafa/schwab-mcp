# StackExchange API v2.3

Reference for the live coding interview on April 25. The problem uses `GET /search`.

## Base URL

```
https://api.stackexchange.com/2.3
```

## GET /search

Search for questions matching given criteria.

### Key Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `intitle` | string | Text that must appear in the title |
| `tagged` | string | Semicolon-delimited tags to filter by |
| `nottagged` | string | Semicolon-delimited tags to exclude |
| `site` | string | **Required.** The SE site (e.g. `stackoverflow`) |
| `q` | string | Free-form search query |
| `page` | int | Page number (default 1) |
| `pagesize` | int | Results per page (default 30, max 100) |
| `order` | string | `desc` or `asc` |
| `sort` | string | `activity`, `votes`, `creation`, `relevance` |
| `fromdate` | unix timestamp | Min creation date |
| `todate` | unix timestamp | Max creation date |
| `accepted` | boolean | Filter to accepted-answer-only questions |
| `min` / `max` | int | Min/max value for the sort field |
| `filter` | string | Controls which fields are returned |

### Example Request

```
GET https://api.stackexchange.com/2.3/search?intitle=python&site=stackoverflow&pagesize=10
```

### Response Shape

```json
{
  "items": [
    {
      "question_id": 123,
      "title": "...",
      "link": "https://...",
      "score": 42,
      "answer_count": 5,
      "is_answered": true,
      "tags": ["python", "list"],
      "creation_date": 1234567890,
      "last_activity_date": 1234567890,
      "view_count": 1000,
      "owner": { "user_id": 456, "display_name": "...", "reputation": 9999 }
    }
  ],
  "has_more": true,
  "quota_max": 300,
  "quota_remaining": 297,
  "page": 1,
  "pagesize": 10,
  "total": 150
}
```

### Important Response Fields

- `has_more` — whether more pages exist; use for pagination
- `quota_max` / `quota_remaining` — rate limit info (300/day unauthenticated, 10,000 with key)
- `backoff` — if present in response, **must wait this many seconds** before next request

---

## Critical Gotchas

### Gzip Compression
The API returns gzip-compressed responses by default. In Python, `requests` handles this automatically. In Go, you need to handle it or set the `Accept-Encoding` header:
```go
// Go: decompress manually if needed
resp, _ := http.Get(url)
reader, _ := gzip.NewReader(resp.Body)
```

### Rate Limiting
- Without API key: 300 requests/day
- With API key: 10,000 requests/day
- `backoff` field in response = mandatory wait (violating this gets you throttled harder)
- HTTP 429 = explicitly throttled — back off with exponential retry

### Unauthenticated vs Authenticated
- No auth needed for read-only endpoints
- API key (not OAuth) is sufficient for quota increase
- OAuth only needed for write operations

### Error Responses
```json
{
  "error_id": 502,
  "error_message": "...",
  "error_name": "throttle_violation"
}
```
Common error codes: 400 (bad request), 401 (unauthorized), 403 (forbidden), 404 (not found), 429 (throttled)

---

## Edge Cases to Address in the Interview

1. **Empty results** — `items` is `[]`, `has_more` is false, `total` is 0
2. **Pagination** — loop using `has_more` flag; don't assume single page
3. **Rate limiting** — check `backoff` field; handle HTTP 429
4. **Network errors** — timeouts, connection refused, DNS failure
5. **API errors** — `error_id` present in body even with HTTP 200 sometimes
6. **Large result sets** — `pagesize` max is 100; paginate if `total` >> 100
7. **Bad parameters** — `site` is required; missing it returns 400
8. **Compression** — ensure gzip decompression works
9. **Quota exhaustion** — `quota_remaining` = 0; handle gracefully

---

## Practice Function Skeleton (Python)

```python
import requests
from typing import Optional

def search_stack_overflow(
    query: str,
    tags: Optional[list[str]] = None,
    site: str = "stackoverflow",
    max_results: int = 30,
) -> list[dict]:
    """
    Search StackExchange for questions matching the given query.
    Returns a list of question objects.
    """
    base_url = "https://api.stackexchange.com/2.3/search"
    params = {
        "intitle": query,
        "site": site,
        "pagesize": min(max_results, 100),
        "order": "desc",
        "sort": "relevance",
    }
    if tags:
        params["tagged"] = ";".join(tags)

    results = []
    page = 1

    while len(results) < max_results:
        params["page"] = page
        resp = requests.get(base_url, params=params, timeout=10)
        resp.raise_for_status()  # raises HTTPError for 4xx/5xx

        data = resp.json()

        # Check for API-level errors
        if "error_id" in data:
            raise ValueError(f"API error {data['error_id']}: {data.get('error_message')}")

        # Respect backoff
        if "backoff" in data:
            import time
            time.sleep(data["backoff"])

        results.extend(data.get("items", []))

        if not data.get("has_more"):
            break
        page += 1

    return results[:max_results]
```

---

## Test Scenarios to Cover

| Scenario | What to test |
|----------|-------------|
| Happy path | Valid query returns list of questions |
| Empty results | `items: []` — function returns empty list, no error |
| Pagination | `has_more: true` — function fetches next page |
| HTTP 429 | Rate limit hit — function raises or retries with backoff |
| HTTP 400 | Bad request — function raises with clear message |
| `backoff` in response | Function sleeps for the required duration |
| Network timeout | `requests.Timeout` — function propagates or retries |
| Missing `site` param | Should never happen if function signature enforces it |
| Large query results | Stops at `max_results`, doesn't over-fetch |
