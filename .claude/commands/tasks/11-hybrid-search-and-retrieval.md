# Task 11 — Embeddings, hybrid search, richer retrieval

**Read first:** `00-conventions.md`.

**Depends on:** Task 04 (BM25 search, which must keep working standalone), Task 06 (`get_context`, to be upgraded), Task 08 (AI/embedding provider), Task 10 (relationship links, friction clustering — both get better inputs here).

## Objective

Add optional semantic retrieval and combine it with keyword search,
metadata, recency, and relationships into better-ranked results — without
making embeddings a requirement. "A completely local, non-vector mode should
work" is a hard constraint carried over from Phase 1, not relaxed here.

## Scope

In scope:
- `search.semantic: true|false` in config (default `false`, per `00-conventions.md`/the spec).
- An embeddings step in `src/lib/search`: when enabled, embed document chunks on index (Task 04's incremental indexer calls this too) and store vectors — SQLite is fine for a personal-scale corpus (e.g. a vectors table + brute-force cosine similarity, or `sqlite-vec` if it doesn't add meaningful operational weight; avoid standing up a separate vector database — explicitly excluded in the spec's "What not to build").
- Hybrid combination: run BM25 + vector search (when enabled) + metadata filters + recency boost + relationship expansion (from Task 10's link data), then a simple reranking step (weighted score combination is enough; don't reach for a learned reranker) before returning results.
- Upgrade Task 06's `get_context` to use this pipeline and do real token-budgeted assembly (now worth using a real tokenizer if one isn't already a dependency, or a well-calibrated character-based approximation — document which).
- Tests: with `search.semantic: false`, everything from Task 04 still passes unchanged; with it `true` (using a mocked/local embedding provider in tests — no live API calls in CI), a semantically related but keyword-dissimilar query surfaces the right document; `get_context` output stays within its token budget on a larger fixture set than Task 06 used.

Out of scope: anything not already covered by Tasks 04/06/10's data — this task is a retrieval-quality upgrade, not a new data model.

## Deliverables

- `src/lib/search/embeddings.ts`.
- Updated `src/lib/search/index.ts` (hybrid `search()`).
- Updated `src/mcp` `get_context`.
- Tests.

## Acceptance criteria

- [ ] `search.semantic: false` reproduces Task 04's exact behavior and test results — this task adds a mode, it doesn't change the default path.
- [ ] With semantic search enabled, a query with no keyword overlap but clear topical overlap with an example document ranks that document in the top results.
- [ ] `get_context` respects its token budget and its ranking visibly improves (documented with a before/after example in the task's PR/summary) over Task 06's naive top-N.
- [ ] No new required external service — the vector store lives in the existing SQLite runtime file.

## Notes / assumptions

- If `sqlite-vec` (or an equivalent lightweight extension) turns out to be awkward to install/ship on the target Mac mini, brute-force cosine similarity over a personal-scale document count (thousands, not millions) is genuinely fine — don't over-engineer the vector index for a scale this system will never reach.
