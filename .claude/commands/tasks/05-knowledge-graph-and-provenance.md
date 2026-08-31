# Task 05 — Knowledge graph and provenance store (temporal)

**Read first:** `00-conventions.md` (the core invariant, forward-only
consolidation edges, two time axes, derived-state contract).

**Depends on:** Task 02 (document model — `entities`/`facts`/edge fields),
Task 04 (derived index — `documents` table, `index_meta` version tracking).

## Objective

Represent entity relationships and fact-level provenance as a **temporal**
graph, derived from file-resident data, so agents can ask both "what's true
now" and "what was true as of a given date." This task deliberately runs
**before** the AI provider (Task 09) and the organization pipeline
(Task 10) exist: it is built and tested entirely against hand-authored
`entities:`/`facts:` frontmatter fixtures and a deterministic (non-LLM)
extractor. That ordering is intentional — it's what lets the graph's rebuild
guarantee be stated as "zero LLM or network calls," and it means later tasks
(10, 11) have a working graph to write into rather than having to build one
themselves.

## Scope

In scope:
- Schema additions to `runtime/ren.db` (extends Task 04's database, same
  file, same WAL settings — do not open a second database):
  - `entities` — `id` (ULID), `kind`, `canonical_name`, `slug`, `doc_id`
    (nullable — a mention-only entity with no dedicated document has
    `doc_id = NULL` until one exists). Projection of frontmatter
    `entities:` blocks across all documents.
  - `entity_aliases` — `entity_id`, `alias`, `source_doc_id`. Lets the same
    entity be referred to differently across documents while resolving to
    one row.
  - `facts` — `uid` (ULID, **assigned in the source document's frontmatter**
    — this is the persistence mechanism, not a derived key; validated by
    Task 02 for syntactic validity and uniqueness only), `doc_id`,
    `subject_entity_id`, `predicate`, `object_entity_id` /
    `object_literal`, `confidence` (stated), `valid_from` (world time),
    `valid_to` (**derived** — computed by reverse index from another fact's
    `supersedes`, never stored directly from frontmatter), `asserted_at`
    (system time), `source_chunk_hash` (**derived**, for drift detection —
    see Task 15: computed by hashing `facts[].source_excerpt` when the
    frontmatter provides one, otherwise hashing the document's whole body as
    a coarser fallback — the extractor never requires an author to locate an
    exact chunk, it just gets a less precise drift signal without one),
    `superseded_by` (derived, reverse index), `retracted` (derived),
    `effective_confidence` (derived — stated confidence adjusted by
    reinforcement, computed in Task 11 but the column belongs here).
    Partial index on `(subject_entity_id, predicate)` `WHERE valid_to IS
    NULL AND superseded_by IS NULL` — this is the "current facts" hot path
    every read goes through.
  - `fact_edges` — `from_uid`, `kind` (`supersedes` | `contradicts` |
    `duplicates` | `reinforces`), `to_uid`, both `uid`s referencing `facts`.
    Direct projection of the **fact-level** edge arrays nested inside a
    `facts[]` entry (per `00-conventions.md`'s Facts section) — **forward-
    only**, always populated from the newer fact's frontmatter entry, never
    inferred or written back onto the older one.
  - `document_edges` — `from_doc`, `kind` (currently just `supersedes`),
    `to_doc`, both referencing `documents.id`. Direct projection of a
    document's **top-level** `supersedes` field (whole-document replacement,
    used by Task 08's substantive `update_document`) — a different thing
    from `fact_edges` above, which operates one level down inside `facts[]`.
    Don't conflate the two: a document can supersede another document as a
    whole (`document_edges`) independently of, or alongside, individual
    facts superseding other facts (`fact_edges`).
  - `links` — `from_doc`, `to_doc` (nullable — unresolved target),
    `target_raw`, `kind` (`wikilink` | `md` | `sources` | `frontmatter`).
    Entirely derived from a deterministic markdown/frontmatter scan — no AI
    involved. Indexed on `to_doc` for backlinks. This is also where
    type-specific reference fields land — e.g. an `experiment` document's
    `friction` field (Task 14) is scanned into a `links` row with
    `kind: "frontmatter"`, rather than inventing a new edge table per
    relationship type.
  - `provenance` — `subject_kind` (`doc` | `fact`), `subject_id`,
    `source_doc_id`, `source_chunk_hash`, `capture_path`, `extractor`,
    `model`, `prompt_version`, `extracted_at`. Every derived record —
    document or fact — gets a row here linking it back to the Bronze/Silver
    material it came from. Indexed on `source_doc_id` and on `capture_path`
    (the latter is the reverse lookup Task 10 uses to tell whether a capture
    has already been processed — see that task's notes).
- `src/lib/graph/extract.ts`: a **deterministic** extractor that reads a
  document's `entities:`/`facts:` frontmatter plus a markdown-link/`sources`
  scan of the body, and produces graph rows. No AI call anywhere in this
  path. This is what Task 10/11 replace/extend with LLM-derived extraction
  later — the deterministic path must keep working standalone even after
  they exist (e.g. for a user who hand-writes `facts:` in a note).
- `src/lib/graph/query.ts`: `getFacts(subject, predicate, { asOf? })` —
  returns the current fact by default (`valid_to IS NULL`), or the fact
  valid as of a given date; `getFactHistory(subject, predicate)` — the full
  supersession chain, oldest to newest, following `fact_edges`;
  `getRelated(idOrEntity, { maxHops? })` — graph traversal over `links` and
  `fact_edges`/`document_edges` for backlinks/relationship proximity (feeds
  Task 07's `get_related` and Task 12's proximity signal); `getProvenance(id)`
  — full chain from a document or fact back to its source(s), terminating
  either at a Bronze capture (for AI-derived material) or at a
  human-authored document with no further source (see the provenance
  acceptance criterion below for which is required when).
- Wire this extractor into Task 04's incremental indexer: when a document is
  (re)indexed, its graph rows are (re)projected in the same pass, keyed off
  the same content-hash change detection — no separate "graph reindex" step
  in normal operation, though `ren reindex --full` (Task 15) rebuilds both
  together.
- Tests: rebuilding the graph from a fixture tree makes zero calls to any AI
  provider or embedding backend (assert this, don't just assume it — mock
  the provider with a call-counting stub that fails the test on any
  invocation); a fixture where document B has a `facts[]` entry whose
  `supersedes: [<uid of a fact in doc A>]` yields, after indexing, A's fact
  with `valid_to == B's fact's valid_from` and `superseded_by == B`'s fact
  uid, with document A itself byte-identical to before; a second fixture
  where document B's top-level `supersedes: [<A's document id>]` produces a
  `document_edges` row (not a `fact_edges` row) — the two edge kinds tested
  separately so they can't be silently conflated; `getFacts` with no `asOf`
  returns exactly the current fact, `asOf` in the past returns the
  historical one, and `getFactHistory` returns the full chain; an entity
  referenced by `entities[].ref` in two documents under different alias text
  resolves to one row, and a mention with no `ref` becomes a `doc_id = NULL`
  row that later links up when a document using that same name is created;
  reindex determinism — two runs over an unchanged fixture tree produce
  content-identical `entities`/`facts`/`fact_edges`/`document_edges`
  (excluding `indexed_at`).

Out of scope: AI-driven entity/fact extraction (Task 10), consolidation
logic — dedup/contradiction detection/promotion (Task 11), ranking/ proximity
scoring itself (Task 12 consumes `getRelated`, doesn't reimplement graph
traversal), the hot-context cache (Task 13).

## Deliverables

- Schema additions in `src/lib/search/schema.sql` (same database as
  Task 04) or a clearly separated `src/lib/graph/schema.sql` applied to the
  same file.
- `src/lib/graph/extract.ts`.
- `src/lib/graph/query.ts`.
- Indexer wiring.
- Tests.

## Acceptance criteria

- [ ] Deleting `runtime/` and running `ren reindex` reconstructs `entities`,
      `facts`, `fact_edges`, `document_edges`, `links`, and `provenance` with
      **zero** calls to any AI provider or embedding backend — a test fails
      the run if either is invoked.
- [ ] A fixture where a newer fact supersedes an older fact produces the
      derived `valid_to`/`superseded_by` correctly in `fact_edges`, with the
      older document's file untouched; a separate fixture where a newer
      *document* supersedes an older document as a whole produces a
      `document_edges` row instead — the two are not conflated.
- [ ] `getFacts(subject, predicate, { asOf })` returns the correct current or
      historical value; `getFactHistory` returns the full supersession
      chain in order.
- [ ] Entity alias resolution: two documents referencing the same
      `entities[].ref` with different alias text resolve to one `entities`
      row; a `ref`-less mention becomes a linkable placeholder row.
- [ ] Every fact whose host document's `generated.by` is `"agent"` resolves,
      via `provenance`, back to its source document and (transitively, via
      that document's `sources`) to its originating `inbox/` capture; a
      fixture fact on a hand-authored document (`generated.by: "human"`,
      no `sources`) is expected to terminate its provenance chain at that
      document itself, not to fail the check for lacking a Bronze ancestor.
- [ ] Reindex is deterministic: two runs over the same fixture tree produce
      identical `entities`/`facts`/`fact_edges`/`document_edges` content.

## Notes / assumptions

- This task's extractor is intentionally dumb. Resist the urge to make it
  smarter than "read the frontmatter, scan for markdown links" — the value
  of running it before the AI provider exists is exactly that it has no
  judgment calls to get wrong. Task 10/11 own the judgment calls.
- Keep this in the same SQLite file as Task 04's tables, not a second
  database — one WAL-mode connection, one set of pragmas, one thing to back
  up or delete.
