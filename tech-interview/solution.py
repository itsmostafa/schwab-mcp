import json
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
    response = []
    page = 1
    while len(response) < max_results:
        resp = requests.get(
            f"{BASE_URL}/search",
            params={
                "intitle": query,
                "site": site,
                "tagged": ",".join(tagged) if tagged else None,
                "filter": "withbody",
                "sort": "relevance",
                "page": page,
                "pagesize": max_results,
                "nottagged": ",".join(not_tagged) if not_tagged else None,
            },
        )
        resp.raise_for_status()
        body = resp.json()

        if "error_id" in body:
            raise Exception(f"{body['error_id']}: {body.get('error_message', '')}")

        if "backoff" in body:
            time.sleep(body["backoff"])

        response = [
            *response,
            *[res for res in body["items"] if res["is_answered"]],
        ]

        if not body["has_more"]:
            break

        page += 1

    return response[:max_results]


if __name__ == "__main__":
    results = search_questions(query="Golang interface", tagged=["go"])
    with open("results.json", "w") as f:
        json.dump(results, f, indent=2)
