# Task 13 — Hot context cache

**Read first:** `00-conventions.md` (derived-state contract — this task is
the one narrow, explicitly-declared exception to "everything derived
survives `rm -rf runtime/`").

**Depends on:** Task 07 (MCP server — this is exposed as a fast path in
front of its retrieval tools), Task 12 (hybrid retrieval — the cache stores
that pipeline's output, it doesn't reimplement ranking).

## Objective

Maintain a small, bounded, TTL-scoped cache of recently active projects,
topics, decisions, notes, and related concepts, so an agent asking "what's
going on right now" gets an answer in constant time instead of re-running
full hybrid retrieval on every call. stdio MCP server processes are
typically short-lived (one per client session), so this must be a
**persisted** cache (in `runtime/ren.db`, alongside Task 04's tables — not
an in-memory-only structure that resets every process start).

This is the one derived store explicitly **exempt** from the deterministic-
rebuild litmus test in `00-conventions.md`: deleting it changes latency, not
correctness. Say this explicitly wherever the cache is documented, so
Task 15's `ren verify` knows not to flag a cold cache as damage.

## Scope

In scope:
- `hot_context` table: `key` (PK — see key composition below), `scope`,
  `input_hash` (see invalidation below), `payload_json` (the assembled
  bundle: recently active projects, topics, decisions, notes, related
  concepts), `score`, `built_at`, `expires_at`, `hit_count`,
  `last_accessed_at`.
- **Key composition**: a request's cache key is `"<scope>:<params_hash>"`,
  where `scope` is a coarse identifier (`"recent"`, `"project:<id>"`,
  `"topic:<slug>"`) and `params_hash` is a hash of the request's *actual*
  parameters — normalized query text, `limit`/`max_tokens`, any filters,
  `asOf`, and the active `ranking_version`. Two `get_context` calls with the
  same scope but a different query or token budget are different cache
  entries, never sharers of one payload.
- `src/lib/hotcontext/build.ts`: assembles a bundle for a scope by calling
  Task 12's hybrid retrieval (recency-weighted, pulling in `pinned: true`
  documents unconditionally regardless of their score — pinning is a
  file-resident decision per `00-conventions.md`, so it must survive a cold
  cache exactly as it was before), and writes the result with a computed
  `expires_at`.
- **Invalidation**: `input_hash` covers the `file_hash` set of every
  document that *actually contributed* to the bundle, plus the active
  ranking config and `embedding_version` — but a contributor-only hash can't
  detect a *new or newly-relevant* document that should now be in the bundle
  but wasn't a contributor before, so it isn't part of the hash yet either.
  Close that gap with a coarse, cheap **global index-revision counter**
  (bumped in `index_meta` on every successful index write, by Task 04) folded
  into every entry's `input_hash`: broad scopes (`"recent"`, `"topic:*"`)
  invalidate on *any* index change — correct and cheap to rebuild, since
  their bundles are small — while narrow, expensive-to-rebuild scopes may
  additionally track a tighter file-set hash if a real performance need
  shows up later (not required for v1). On a miss (absent, expired, or
  hash-mismatched), rebuild and write through; on a hit, serve directly and
  bump `hit_count`/`last_accessed_at`.
- Bounded size + TTL, both config-driven (`hotContext.maxEntries`,
  `hotContext.ttlSeconds` from Task 01's config block): eviction is real
  **LRU by `last_accessed_at`** (not `hit_count`/`built_at`, neither of
  which records last access) once `maxEntries` is exceeded; an entry past
  `expires_at` is never served stale, only treated as a miss.
- MCP integration: `get_context` and `get_recent_notes` (Task 07) check this
  cache before falling through to Task 12's full pipeline — this task wires
  the fast path in, it doesn't add new MCP tools of its own.
- Tests: deleting the entire `hot_context` table (or `runtime/` wholesale)
  changes only latency — a fixture query returns identical ranked results
  and identical `signals` values cold vs. warm; editing any document bumps
  the global index revision and invalidates broad-scope entries (verified
  against `"recent"`); a newly created, highly-relevant fixture document
  becomes visible in a broad-scope bundle on the next request rather than
  being masked by a stale contributor-only hash; a `pinned: true` fixture
  document is present in the working-set bundle both cold and warm; an entry
  past `expires_at` is never served (rebuilt instead); cache size never
  exceeds `maxEntries` under a fixture load that would otherwise overflow
  it, with eviction removing the least-recently-*accessed* entry, not
  simply the oldest-built one; two requests differing only in `max_tokens`
  produce two distinct cache entries, never a shared/wrong-budget hit.

Out of scope: any new ranking logic (this task is a cache in front of
Task 12's pipeline, not a second ranking implementation), cross-session
sharing beyond what SQLite already provides (no distributed cache — single
Mac mini, single database file).

## Deliverables

- `hot_context` schema addition to `runtime/ren.db`.
- `src/lib/hotcontext/build.ts`, `src/lib/hotcontext/invalidate.ts`.
- MCP wiring in `get_context`/`get_recent_notes`.
- Tests.

## Acceptance criteria

- [ ] Cold vs. warm cache produce identical ranked results and identical
      per-signal breakdowns for the same fixture query — only latency
      differs.
- [ ] `ren verify` (Task 15) does not flag an empty/cold `hot_context` table
      as an inconsistency.
- [ ] Editing any fixture document invalidates broad-scope cache entries
      (via the global index-revision counter) so a newly relevant document
      is never masked by a stale contributor-only hash.
- [ ] `pinned: true` documents appear in the relevant bundle regardless of
      recency/score, both cold and warm.
- [ ] Cache respects configured `maxEntries` (evicting by least-recently-
      accessed) and `ttlSeconds`; no stale entry is ever served past
      `expires_at`; requests differing only in query/limit/`max_tokens`
      never collide on one cache entry.

## Notes / assumptions

- Keep the cache dumb on purpose: it stores what Task 12 computed, it
  doesn't second-guess it. If ranking needs to change, change Task 12 — this
  task should never grow its own scoring logic.
