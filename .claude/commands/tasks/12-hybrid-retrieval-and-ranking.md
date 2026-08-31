# Task 12 — Hybrid retrieval and configurable ranking

**Read first:** `00-conventions.md` (hybrid retrieval section — semantic on
by default, degrade-gracefully contract, two config axes).

**Depends on:** Task 04 (BM25/FTS + vectors table, which must keep working
standalone when semantic is unavailable), Task 05 (knowledge graph — the
proximity signal and `getRelated`), Task 07 (`get_context`, `get_related`,
to be upgraded), Task 09 (embedding backend), Task 11 (consolidated
facts/importance signals get better inputs here).

## Objective

Combine semantic similarity, FTS/BM25 relevance, importance, recency, and
relationship proximity into one configurably-weighted ranking, exposed
through Ren's MCP interface — without ever making embeddings a hard
requirement. "A completely local, non-vector mode should work" is a hard
constraint carried over from Task 04, not relaxed here: it's the fallback
path, not a legacy mode being phased out.

## Scope

In scope:
- Ranking config (already sketched in Task 01; this task implements it):

  ```ts
  search: {
    semantic: true,   // default on — see 00-conventions.md
    ranking: {
      weights: { lexical: 0.40, semantic: 0.35, importance: 0.10, recency: 0.10, proximity: 0.05 },
      recencyHalfLifeDays: 90,
      proximityMaxHops: 2,
      candidateLimit: 200,
      includeSuperseded: false,
    },
  }
  ```

- `src/lib/search/rank.ts`: candidate generation is the union of BM25 top-K
  (Task 04's `documents_fts`/`chunks_fts`), vector top-K (Task 04's
  `vectors` table via Task 09's backend, brute-force cosine similarity is
  fine at this scale — don't stand up a separate vector database, that's
  explicitly excluded per `00-conventions.md`'s "what not to build"), and
  graph expansion via Task 05's `getRelated` up to `proximityMaxHops`. Score
  each candidate:

  ```
  score(d) = Σᵢ wᵢ' · nᵢ(d)     where wᵢ' = wᵢ / Σⱼ∈available wⱼ

  n_lexical    = 1 - minmax(bm25_scores)   -- SQLite's bm25() is LOWER-is-better;
                                            -- inverting after min-max normalization
                                            -- makes n_lexical HIGHER-is-better like
                                            -- every other signal. With exactly one
                                            -- candidate (no spread to normalize),
                                            -- n_lexical = 1.0 by convention.
  n_semantic   = (cosine_similarity + 1) / 2
  n_importance = frontmatter importance, or the by-type default from
                 00-conventions.md's Importance table when unset
  n_recency    = 0.5 ^ (age_days / recencyHalfLifeDays)
  n_proximity  = 0.5 ^ hops_from_seed   (0 if unreachable within proximityMaxHops)
  ```

  Every result carries a `signals: { lexical, semantic, importance, recency,
  proximity }` breakdown in addition to the combined score — this is cheap
  to compute here and is what makes weight tuning and Task 15's `ren
  verify`/explain story possible later. Don't ship ranking as a black box.
- **Graceful degradation**: when the embedding backend is unavailable (not
  configured, model not cached and offline, or a dimension mismatch against
  `index_meta`), `semantic` drops out of the weight set entirely and the
  remaining weights renormalize (per the formula above — this is exactly
  what "renormalize over available signals" in `00-conventions.md` means
  concretely). Results carry `degraded: true` and a `degradedReason`; this
  is never an error path.
- Temporal filtering: `includeSuperseded: false` (the default) excludes
  results whose derived `valid_to` is in the past relative to `asOf`
  (defaults to now) — for a **fact**-shaped result this reads `facts.valid_to`
  (Task 05's fact-level reverse index); for a **document**-shaped result
  (e.g. a substantively-updated document with a document-level `supersedes`
  edge pointing at it) this reads the equivalent reverse index over Task 05's
  `document_edges` table, which the same task defines and this task only
  consumes. Set `includeSuperseded: true` or pass an explicit `asOf` to see
  historical state either way, using Task 05's temporal query functions
  rather than reimplementing either filter here.
- Upgrade Task 07's `get_context(query, max_tokens)` to use this pipeline:
  real token-budgeted assembly (worth using a real tokenizer now if one
  isn't already a dependency, or a well-calibrated character-based
  approximation — document which), ranked by the hybrid score above,
  including graph-adjacent documents where they fit the budget, with source
  attribution per contributing document. Upgrade `get_related` to rank by
  the `proximity`/`semantic` signals rather than returning an unordered set.
- Tests: with the embedding backend mocked as unavailable, `semantic` drops
  out and the remaining weights (lexical/importance/recency/proximity)
  renormalize per the formula above — the *relative order* this produces is
  not required to match a pure-BM25 ranking (importance/recency/proximity
  are still contributing), but the result set is marked `degraded: true`
  with a reason, and a fixture query with a clear single best match still
  ranks it first under the renormalized weights; this task adds a mode, it
  doesn't remove the fallback path; with it available (mocked/local
  provider in tests — no live API calls in CI), a semantically related but
  keyword-dissimilar query surfaces the right document, and a document one
  hop away in the graph from a strong keyword match outranks an unrelated
  document with a weaker keyword match; `get_context` output stays within
  its token budget on a larger fixture set than Task 07 used, and its
  ranking visibly improves over Task 07's naive top-N (document a
  before/after example); `includeSuperseded: false` excludes a fixture's
  superseded fact from results by default, `true` includes it.

Out of scope: anything not already covered by Tasks 04/05/09/11's data —
this task is a retrieval-quality upgrade over existing derived state, not a
new data model. The hot-context cache (Task 13) sits in front of this
pipeline's output, it doesn't change the ranking formula itself.

## Deliverables

- `src/lib/search/rank.ts` (hybrid scoring + degrade logic).
- Updated `src/lib/search/index.ts` (`search()` now returns ranked, signal-annotated results).
- Updated `src/mcp` `get_context`, `get_related`.
- Config additions (`search.ranking.*`).
- Tests.

## Acceptance criteria

- [ ] With the embedding backend unavailable, `semantic` drops out and the
      remaining weights renormalize (not a reproduction of Task 04's
      pure-BM25 ordering, since importance/recency/proximity still
      contribute) — results are marked `degraded: true` with a reason, never
      an error.
- [ ] With semantic search available, a query with no keyword overlap but
      clear topical overlap with an example document ranks that document in
      the top results.
- [ ] A document one graph-hop from a strong keyword match outranks an
      unrelated document with a weaker keyword match, demonstrating the
      proximity signal contributes to ranking.
- [ ] `get_context` respects its token budget and its ranking visibly
      improves (documented with a before/after example) over Task 07's naive
      top-N.
- [ ] `includeSuperseded` correctly includes/excludes historical facts, using
      Task 05's temporal queries.
- [ ] No new required external service — the vector store lives in the
      existing SQLite runtime file; ranking weights are read from config, not
      hardcoded.

## Notes / assumptions

- If `sqlite-vec` (or an equivalent lightweight extension) turns out to be
  awkward to install/ship on the target Mac mini, brute-force cosine
  similarity over a personal-scale document count (thousands, not millions)
  is genuinely fine — don't over-engineer the vector index for a scale this
  system will never reach.
- Resist adding a learned reranker or a query-planner abstraction — weighted
  score combination over five signals is enough. Simple yet powerful.
