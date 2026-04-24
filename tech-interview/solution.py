import json
import threading
import time
from typing import Optional

import requests

BASE_URL = "https://api.stackexchange.com/2.3"
_CACHE_TTL = 300  # seconds — StackExchange throttles aggressively; 5-min TTL avoids repeat hits

# Module-level in-memory cache keyed by normalized request params.
# Stores (result_list, expiry_monotonic_timestamp) so we can check staleness without wall-clock drift.
_cache: dict[tuple, tuple[list, float]] = {}
# Lock guards both reads and writes to _cache — dict ops are not atomic across threads.
_cache_lock = threading.Lock()


def _make_key(query, site, tagged, not_tagged, max_results) -> tuple:
    # Sort tag lists before hashing so ["go", "python"] and ["python", "go"] map to the same key.
    return (
        query,
        site,
        tuple(sorted(tagged)) if tagged else (),
        tuple(sorted(not_tagged)) if not_tagged else (),
        max_results,
    )


def clear_cache() -> None:
    with _cache_lock:
        _cache.clear()


def search_questions(
    query: str,
    site: str = "stackoverflow",
    tagged: Optional[list[str]] = None,
    not_tagged: Optional[list[str]] = None,
    max_results: int = 30,
) -> list[dict]:
    """Return up to max_results answered questions matching query, with caching."""
    key = _make_key(query, site, tagged, not_tagged, max_results)

    # Check cache before hitting the network.
    # We hold the lock only long enough to read + optionally evict the expired entry —
    # the actual fetch happens outside the lock so we don't block other threads for the full RTT.
    with _cache_lock:
        if key in _cache:
            result, expires_at = _cache[key]
            if time.monotonic() < expires_at:
                return result
            # Entry exists but is stale — remove it so we don't serve it to a concurrent caller
            # who might sneak in between this eviction and the store below.
            del _cache[key]

    result = _fetch_questions(query, site, tagged, not_tagged, max_results)

    with _cache_lock:
        _cache[key] = (result, time.monotonic() + _CACHE_TTL)

    return result


def _fetch_questions(
    query: str,
    site: str,
    tagged: Optional[list[str]],
    not_tagged: Optional[list[str]],
    max_results: int,
) -> list[dict]:
    """Paginate through /search until we have max_results answered questions or exhaust pages."""
    collected: list[dict] = []
    page = 1
    # API cap per page is 100, but 25 is the sweet spot: fewer wasted items when we hit max_results
    # mid-page while still being large enough to minimize round-trips.
    page_size = min(25, max_results)

    while len(collected) < max_results:
        resp = requests.get(
            f"{BASE_URL}/search",
            params={
                "intitle": query,       # matches questions whose title contains the query
                "site": site,
                "tagged": ",".join(tagged) if tagged else None,
                "filter": "withbody",   # include full question body in response
                "sort": "relevance",
                "page": page,
                "pagesize": page_size,
                "nottagged": ",".join(not_tagged) if not_tagged else None,
            },
        )
        resp.raise_for_status()  # surface 4xx/5xx as exceptions immediately
        body = resp.json()

        # The API returns HTTP 200 even for application-level errors (e.g. bad site param);
        # check the error_id field explicitly.
        if "error_id" in body:
            raise Exception(f"{body['error_id']}: {body.get('error_message', '')}")

        # The API may include a `backoff` field instructing clients to pause before the next request.
        # Ignoring it risks a temporary ban — sleep before continuing pagination.
        if "backoff" in body:
            time.sleep(body["backoff"])

        # Only keep answered questions — unanswered ones aren't useful for study material.
        collected.extend(item for item in body["items"] if item["is_answered"])

        # has_more=False means we've exhausted all matching questions; no point requesting page N+1.
        if not body["has_more"]:
            break

        page += 1

    # Slice to exact max_results in case the last page pushed us over.
    return collected[:max_results]


if __name__ == "__main__":
    results = search_questions(query="Golang interface", tagged=["go"])
    with open("results.json", "w") as f:
        json.dump(results, f, indent=2)
