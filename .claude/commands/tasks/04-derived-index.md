# Task 04 — Derived SQLite index: FTS5, WAL, incremental indexing

**Read first:** `00-conventions.md` (the core invariant, derived-state
contract, version tracking, tiers, two time axes).

**Depends on:** Task 01 (config), Task 02 (document loader/schema/hash),
Task 03 (inbox exists — this task indexes it unconditionally, see scope).

## Objective

Derived, rebuildable SQLite state that makes keyword and (later) semantic
search fast, without the Markdown/OKF files ever stopping being canonical.
If `runtime/ren.db` is deleted, `ren reindex` must reconstruct it fully from
source — this is the first task where the litmus test in `00-conventions.md`
(`rm -rf $REN_DATA_DIR/runtime && ren reindex` loses nothing) becomes
concrete and testable.

This task owns the **shape** of the derived index, including the vectors
table and version-tracking columns that later tasks fill in. It does not
implement embedding generation (Task 09), hybrid ranking (Task 12), the
knowledge graph (Task 05), or consolidation (Task 11) — but its schema must
not need to change shape when those land, only get populated.

## Scope

In scope:
- `runtime/ren.db`, opened with `journal_mode=WAL`, `synchronous=NORMAL`,
  `busy_timeout=5000`, `foreign_keys=ON` on every connection. WAL matters
  here specifically because the website (Astro SSR), the MCP server, and the
  file watcher are three separate processes that will read this database
  concurrently once later tasks exist — get the pragma right now even though
  this task itself is still single-process.
- Schema (`src/lib/search/schema.sql` or inline):
  - `documents` — one row per document in **either** `knowledge/` or
    `inbox/` — indexing both is **mandatory**, not optional: Task 05's
    provenance reverse-lookup by `capture_path`, Task 10's already-processed
    check, and Task 15's full-reindex all assume `inbox/` has rows here.
    Columns: `id`, `path`, `zone` (`"inbox"` | `"knowledge"`), `tier`,
    `type`, `title`, `description`, `tags_json`, `sources_json`,
    `generated_json`, `verified`, `status`, `importance`, `confidence`,
    `valid_from`, `valid_to`, `created_at`, `updated_at`, `stale_after` —
    all projections of frontmatter (rebuildable by reparsing) — plus
    derived-only bookkeeping columns: `file_hash`, `frontmatter_hash`,
    `body_hash` (from Task 02's hash helper), `mtime_ns`, `size`,
    `parser_version`, `indexed_at`. Index on `(type, updated_at DESC)`,
    `(file_hash)`, `(tier, importance DESC)`, `(zone)`. **`search()`
    defaults to `zone = "knowledge"` unless a caller passes a filter
    explicitly including `"inbox"`** — indexing both zones is about making
    provenance/reprocessing lookups possible, not about surfacing raw Bronze
    captures in ordinary keyword search results.
  - `chunks` — one row per chunked span of a document body: `id`, `doc_id`,
    `ordinal`, `heading_path`, `char_start`, `char_end`, `text`,
    `token_est`, `chunk_hash`, `chunker_version`. Entirely derived. Unique on
    `(doc_id, ordinal)`; index on `chunk_hash` — this is the reuse key that
    lets an edit to one paragraph re-embed one chunk, not the whole
    document, once Task 09 exists.
  - `vectors` — **primary key is `(chunk_hash, model)`, not `chunk_id`**.
    Columns: `chunk_hash`, `model`, `dim`, `embedding` (BLOB, float32),
    `norm`, `embedded_at`. Deliberately keyed by content, not by the chunk
    row's identity: when a document is rechunked, `chunks` rows are
    recreated (new `id` per ordinal even if the text is unchanged), but a
    `vectors` row keyed by `(chunk_hash, model)` is trivially reusable by any
    chunk — in this document or another — whose `chunk_hash` matches, with a
    plain lookup join (`chunks.chunk_hash = vectors.chunk_hash`), no
    copy/rekey step needed. Entirely derived, entirely empty until Task 09
    populates it — this task defines the table and an `EmbeddingBackend`
    interface (`id`/`dims`/`embed(texts) -> vectors[]` — name matches
    `00-conventions.md` exactly, do not call it `EmbeddingProvider`) with a
    **no-op implementation** that the indexer calls unconditionally (so
    wiring an embedder later is a one-file change, not a schema change). A
    brute-force cosine-similarity query helper belongs here too, even though
    nothing calls it until Task 12.
  - `documents_fts` / `chunks_fts` — two separate FTS5 virtual tables
    (title/description/tags vs. body chunk text) so `bm25()` weights differ
    between the two — a title match should outrank an incidental body match.
    Keep sync between the source tables and their FTS shadow explicit in the
    indexer code (on write), not implicit via SQLite triggers, so a partial
    write can't leave FTS out of sync silently.
  - `index_meta` — key/value: `schema_version`, `parser_version`,
    `chunker_version`, `embedding_model`, `embedding_version`,
    `ranking_version`, `last_full_reindex_at`, `index_revision` (an integer,
    bumped by 1 on every successful write to `documents`/`chunks`/`vectors`
    — a cheap global "something changed" signal Task 13's cache invalidation
    reads; not itself meaningful content, just a counter). A version bump
    here is what later tasks use to force targeted reprocessing (see below).
    `last_full_reindex_at` and `index_revision` (and any other operational
    counter/timestamp added here later) are explicitly **excluded** from the
    determinism check below — see `00-conventions.md`'s Version tracking
    section for the exact excluded-key list; every other key must be
    reproducible.
  - `index_errors` — `path`, `phase`, `message`, `seen_at`: a quarantine
    feed for documents that fail to parse/index, read by Task 15's
    `ren doctor`/`ren verify`. A malformed file must land here, not crash
    the indexer.
- `src/lib/search/index.ts`: `indexDocument(doc)`, `removeDocument(id)`,
  `search(query, opts)` returning ranked results with snippets (BM25-only
  ranking for now, scoped to `zone: "knowledge"` by default per above —
  Task 12 replaces this with hybrid scoring without changing the function's
  shape).
- Incremental indexing keyed on content hash: `mtime_ns`+`size` is a cheap
  pre-check; if unchanged, skip entirely without reading the file. If
  changed, recompute `file_hash`; if `frontmatter_hash` alone changed,
  reproject metadata + FTS row without rechunking; if `body_hash` changed,
  rechunk, and for each new chunk look up `vectors` by `chunk_hash` before
  treating it as needing a fresh embedding (the actual embedding call is
  Task 09's job — this task's indexer just needs to leave the right hook and
  not blow away reusable vector rows on an unrelated edit).
- `ren reindex` CLI command (or `npm run reindex` if the CLI itself is
  deferred — see Task 01/16 on CLI scope): wipes and rebuilds the whole index
  from **both** `knowledge/` and `inbox/` deterministically. A version bump
  in `index_meta` (e.g. `chunkerVersion` changes) must also be able to
  trigger a **targeted** reprocess of only the affected rows, not force a
  full `ren reindex` — this is the seed of Task 15's fuller `ren reindex`.
- A file watcher (dev-mode, e.g. via `chokidar` or Node's built-in
  `fs.watch` if sufficient — prefer the smaller dependency): over
  `knowledge/`, on add/change, re-validate + re-index that one document; on
  delete, remove it from the index. Over `inbox/`, **add events only** — a
  Bronze capture is never modified or deleted once written (per Task 03), so
  the watcher only ever needs to index a new file once; wire it either as a
  restricted watch (ignore change/unlink under `inbox/`) or, more simply,
  have Task 03's write path call `indexDocument` directly right after its
  atomic write — pick whichever is simpler to implement correctly, but don't
  leave `inbox/` unindexed until the next full reindex. No full rebuild on
  every change, either way.
- Tests: reindex determinism, incremental update on file change, incremental
  removal on file delete, search ranks an exact-title match above an
  incidental body match, a document whose only change is a typo in one
  paragraph re-chunks that paragraph without invalidating unrelated chunks'
  `vectors` rows (verify via `chunk_hash` reuse, even with the no-op
  embedder — the row simply won't exist yet, but the reuse *lookup* path
  must be exercised), a malformed document lands in `index_errors` instead
  of crashing the indexer or the file watcher, a new `inbox/` capture gets a
  `documents` row (zone `"inbox"`) without being returned by a default
  (zone-unfiltered) `search()` call.

Out of scope: semantic/vector search generation (Task 09 fills `vectors`),
hybrid ranking/reranking (Task 12), the knowledge graph — entities,
relationships, provenance (Task 05), consolidation state (Task 11).

**Determinism is scoped, not absolute** (per `00-conventions.md`): two
reindex runs over the same fixture set must produce identical `documents`,
`chunks`, FTS, and `index_meta` content, **excluding** the named operational
timestamps (`indexed_at`, `last_full_reindex_at`) — everything else,
including file-projected `created_at`/`updated_at`, must match exactly.
Since `vectors` is empty in this task (no-op embedder), that table is
trivially deterministic here too — the float-tolerance carve-out in
`00-conventions.md` becomes relevant only once Task 09 supplies real
embeddings.

## Deliverables

- `src/lib/search/schema.sql` (or inline schema) — documents, chunks,
  vectors, `documents_fts`, `chunks_fts`, `index_meta`, `index_errors`.
- `src/lib/search/embeddings.ts` — `EmbeddingBackend` interface + no-op
  implementation.
- `src/lib/search/index.ts`.
- File watcher wiring (dev server integration) covering both `knowledge/`
  and `inbox/` per above.
- `ren reindex` command.
- Tests.

## Acceptance criteria

- [ ] Deleting `runtime/ren.db` and running `ren reindex` fully restores
      search functionality with no data loss (source files are canonical),
      covering both `knowledge/` and `inbox/`.
- [ ] Editing a file in `knowledge/` while the dev server runs updates search
      results without a full reindex; a new file appearing in `inbox/` is
      indexed (zone `"inbox"`) without a full reindex either.
- [ ] `search("some phrase")` with no zone filter returns ranked, relevant
      results with snippets from `knowledge/` only, with title matches
      ranked above incidental body matches; passing an explicit filter that
      includes `"inbox"` also surfaces matching captures.
- [ ] Reindex is deterministic across repeated runs on unchanged input,
      excluding only the named operational timestamps (`indexed_at`,
      `last_full_reindex_at`) — every other column, including file-projected
      timestamps, is content-equal; `vectors` is trivially deterministic
      while empty.
- [ ] Editing one paragraph of a multi-chunk document only rechunks that
      document's affected chunks — unaffected chunks' rows (`id`,
      `chunk_hash`) are untouched, verified by row identity across the
      before/after index state.
- [ ] A document with malformed frontmatter is skipped with an entry in
      `index_errors`, not a crash of the indexer or the dev server.

## Notes / assumptions

- SQLite FTS5 is sufficient for BM25-style ranking out of the box (`bm25()`
  ranking function) — no need for a separate search library. Keep this to
  one dependency (the sqlite driver) if at all possible.
- The `EmbeddingBackend` interface and `vectors` table exist now
  specifically so Task 09 doesn't have to touch this task's schema — decide
  the shape carefully here, since it's the seam three later tasks (09, 11,
  12) build on. Use this exact interface name everywhere it's referenced —
  `00-conventions.md` is the canonical spelling.
