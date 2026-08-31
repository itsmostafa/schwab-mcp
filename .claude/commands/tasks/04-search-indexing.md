# Task 04 — SQLite search index, file watching, reindex

**Read first:** `00-conventions.md`.

**Depends on:** Task 01 (config), Task 02 (document loader/schema), Task 03 (inbox exists, though indexing targets `knowledge/` primarily — see scope).

## Objective

Derived, rebuildable SQLite state that makes keyword search fast, without the
Markdown/OKF files ever stopping being canonical. If the SQLite file is
deleted, `sia reindex` must reconstruct it fully from source.

## Scope

In scope:
- `runtime/sia.db` schema: a documents table (id, path, type, title, tags, timestamps, verified/status) and an FTS5 virtual table for full-text/BM25 search over title + body.
- `src/lib/search/index.ts`: `indexDocument(doc)`, `removeDocument(id)`, `search(query, opts)` returning ranked results with snippets.
- `sia reindex` CLI command (or `npm run reindex` if the CLI itself is deferred — see Task 01/12 on CLI scope): wipes and rebuilds the whole index from `knowledge/` (and optionally `inbox/`) deterministically — same input files must produce the same index content on every run.
- A file watcher (dev-mode, e.g. via `chokidar` or Node's built-in `fs.watch` if sufficient — prefer the smaller dependency) over `knowledge/`: on add/change, re-validate + re-index that one document; on delete, remove it from the index. No full rebuild on every change.
- Tests: reindex determinism (two runs on the same fixture set produce identical index contents), incremental update on file change, incremental removal on file delete, search ranks an exact-title match above an incidental body match.

Out of scope: semantic/vector search (Task 11), reranking/hybrid combination (Task 11), relationship/backlink traversal (folded into Task 06's `get_related` at the reading level — a naive markdown-link scan is enough here if useful, but don't build a graph engine in this task).

## Deliverables

- `src/lib/search/schema.sql` (or inline schema).
- `src/lib/search/index.ts`.
- File watcher wiring (dev server integration).
- `sia reindex` command.
- Tests.

## Acceptance criteria

- [ ] Deleting `runtime/sia.db` and running `sia reindex` fully restores search functionality with no data loss (source files are canonical).
- [ ] Editing a file in `knowledge/` while the dev server runs updates search results without a full reindex.
- [ ] `search("some phrase")` returns ranked, relevant results with snippets for the example knowledge docs.
- [ ] Reindex is deterministic across repeated runs on unchanged input.

## Notes / assumptions

- SQLite FTS5 is sufficient for BM25-style ranking out of the box (`bm25()` ranking function) — no need for a separate search library. Keep this to one dependency (the sqlite driver) if at all possible.
