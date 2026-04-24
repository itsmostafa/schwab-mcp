import time
from typing import Optional

import requests

BASE_URL = "https://api.stackexchange.com/2.3"


def search_questions(
    query: str,
    site: str = "stackoverflow",
    tagged: Optional[list[str]] = None,
    not_tagged: Optional[list[str]] = None,
    max_results: int = 30,
) -> list[dict]:
    """
    Search StackExchange for questions matching `query`.

    Returns up to `max_results` questions, each with:
    title, link, score, answer_count, is_answered, tags
    """
    resp = requests.get(
        f"{BASE_URL}/search",
        params={
            "intitle": query,
            "site": site,
            "tagged": ",".join(tagged) if tagged else None,
            "filter": "withbody",
            "sort": "relevance",
            "page": 1,
            "pagesize": max_results,
            "nottagged": ",".join(not_tagged) if not_tagged else None,
        },
    )
    return [res for res in resp.json()["items"] if res["is_answered"]]


if __name__ == "__main__":
    print(search_questions(query="Golang interface", tagged=["go"]))
