from unittest.mock import MagicMock, patch

import pytest
import requests
from solution import clear_cache, search_questions

# ---------------------------------------------------------------------------
# Helpers
# ---------------------------------------------------------------------------


def make_response(items=None, has_more=False, backoff=None, error_id=None, status=200):
    """Build a mock requests.Response for a StackExchange API call."""
    mock = MagicMock()
    mock.status_code = status
    mock.raise_for_status = MagicMock()
    if status >= 400:
        mock.raise_for_status.side_effect = Exception(f"HTTP {status}")

    body = {
        "items": items or [],
        "has_more": has_more,
        "quota_max": 300,
        "quota_remaining": 299,
    }
    if backoff is not None:
        body["backoff"] = backoff
    if error_id is not None:
        body["error_id"] = error_id
        body["error_message"] = "test error"
        body["error_name"] = "test_error"

    mock.json.return_value = body
    return mock


SAMPLE_ITEM = {
    "question_id": 1,
    "title": "How do I foo?",
    "link": "https://stackoverflow.com/q/1",
    "score": 10,
    "answer_count": 3,
    "is_answered": True,
    "tags": ["python", "foo"],
}


# ---------------------------------------------------------------------------
# Happy path
# ---------------------------------------------------------------------------


@patch("solution.requests.get")
def test_returns_questions_on_valid_query(mock_get):
    mock_get.return_value = make_response(items=[SAMPLE_ITEM], has_more=False)
    results = search_questions("foo")
    assert len(results) == 1
    assert results[0]["title"] == "How do I foo?"


@patch("solution.requests.get")
def test_output_fields_present(mock_get):
    mock_get.return_value = make_response(items=[SAMPLE_ITEM], has_more=False)
    results = search_questions("foo")
    required = {"title", "link", "score", "answer_count", "is_answered", "tags"}
    assert required.issubset(results[0].keys())


@patch("solution.requests.get")
def test_tag_filtering_passed_correctly(mock_get):
    mock_get.return_value = make_response(items=[SAMPLE_ITEM], has_more=False)
    search_questions("foo", tagged=["python", "django"])
    call_kwargs = mock_get.call_args
    params = call_kwargs[1].get("params") or call_kwargs[0][1]
    assert params["tagged"] == "python,django"


# ---------------------------------------------------------------------------
# Empty results
# ---------------------------------------------------------------------------


@patch("solution.requests.get")
def test_empty_results_returns_empty_list(mock_get):
    mock_get.return_value = make_response(items=[], has_more=False)
    results = search_questions("xyzzy-no-results")
    assert results == []


# ---------------------------------------------------------------------------
# Pagination
# ---------------------------------------------------------------------------


@patch("solution.requests.get")
def test_paginates_when_has_more_is_true(mock_get):
    page1 = make_response(items=[SAMPLE_ITEM] * 10, has_more=True)
    page2 = make_response(items=[SAMPLE_ITEM] * 5, has_more=False)
    mock_get.side_effect = [page1, page2]
    results = search_questions("foo", max_results=15)
    assert len(results) == 15
    assert mock_get.call_count == 2


@patch("solution.requests.get")
def test_stops_at_max_results(mock_get):
    page1 = make_response(items=[SAMPLE_ITEM] * 100, has_more=True)
    mock_get.return_value = page1
    results = search_questions("foo", max_results=5)
    assert len(results) == 5


# ---------------------------------------------------------------------------
# Rate limiting
# ---------------------------------------------------------------------------


@patch("solution.time.sleep")
@patch("solution.requests.get")
def test_respects_backoff_field(mock_get, mock_sleep):
    page1 = make_response(items=[SAMPLE_ITEM], has_more=True, backoff=3)
    page2 = make_response(items=[SAMPLE_ITEM], has_more=False)
    mock_get.side_effect = [page1, page2]
    search_questions("foo", max_results=2)
    mock_sleep.assert_called_once_with(3)


@patch("solution.requests.get")
def test_raises_on_http_429(mock_get):
    mock_get.return_value = make_response(status=429)
    with pytest.raises(Exception):
        search_questions("foo")


# ---------------------------------------------------------------------------
# HTTP errors
# ---------------------------------------------------------------------------


@patch("solution.requests.get")
def test_raises_on_http_400(mock_get):
    mock_get.return_value = make_response(status=400)
    with pytest.raises(Exception):
        search_questions("foo")


@patch("solution.requests.get")
def test_raises_on_http_500(mock_get):
    mock_get.return_value = make_response(status=500)
    with pytest.raises(Exception):
        search_questions("foo")


# ---------------------------------------------------------------------------
# API-level errors (error_id in 200 body)
# ---------------------------------------------------------------------------


@patch("solution.requests.get")
def test_raises_on_api_error_in_body(mock_get):
    mock_get.return_value = make_response(items=[], error_id=502)
    with pytest.raises(Exception, match="502"):
        search_questions("foo")


# ---------------------------------------------------------------------------
# Network failures
# ---------------------------------------------------------------------------


@patch("solution.requests.get", side_effect=requests.exceptions.Timeout)
def test_raises_on_timeout(mock_get):
    with pytest.raises(Exception):
        search_questions("foo")


@patch("solution.requests.get", side_effect=requests.exceptions.ConnectionError)
def test_raises_on_connection_error(mock_get):
    with pytest.raises(Exception):
        search_questions("foo")


# ---------------------------------------------------------------------------
# Cache isolation
# ---------------------------------------------------------------------------


@pytest.fixture(autouse=True)
def reset_cache():
    clear_cache()
    yield
    clear_cache()


# ---------------------------------------------------------------------------
# TTL caching
# ---------------------------------------------------------------------------


@patch("solution.requests.get")
def test_same_query_uses_cache(mock_get):
    mock_get.return_value = make_response(items=[SAMPLE_ITEM], has_more=False)
    result1 = search_questions("foo")
    result2 = search_questions("foo")
    assert result1 == result2
    assert mock_get.call_count == 1


@patch("solution.requests.get")
def test_different_queries_bypass_cache(mock_get):
    mock_get.return_value = make_response(items=[SAMPLE_ITEM], has_more=False)
    search_questions("foo")
    search_questions("bar")
    assert mock_get.call_count == 2


@patch("solution.requests.get")
def test_different_tags_bypass_cache(mock_get):
    mock_get.return_value = make_response(items=[SAMPLE_ITEM], has_more=False)
    search_questions("foo", tagged=["python"])
    search_questions("foo", tagged=["go"])
    assert mock_get.call_count == 2


@patch("solution.requests.get")
def test_clear_cache_forces_refetch(mock_get):
    mock_get.return_value = make_response(items=[SAMPLE_ITEM], has_more=False)
    search_questions("foo")
    clear_cache()
    search_questions("foo")
    assert mock_get.call_count == 2


# ---------------------------------------------------------------------------
# Smarter pagination
# ---------------------------------------------------------------------------


@patch("solution.requests.get")
def test_pagesize_capped_at_25(mock_get):
    mock_get.return_value = make_response(items=[], has_more=False)
    search_questions("foo", max_results=100)
    params = mock_get.call_args[1]["params"]
    assert params["pagesize"] == 25


# ---------------------------------------------------------------------------
# Debrief prompts (not tests — read these after each session)
# ---------------------------------------------------------------------------
#
# After your mock, think through:
#
# 1. How would you cache responses? What's the cache key? TTL?
# 2. Under high load (100 concurrent callers), what breaks first?
# 3. How would you add retry logic with exponential backoff?
# 4. What observability would you add (metrics, logs, traces)?
# 5. If `max_results` is 10,000, what's the minimum number of API calls?
