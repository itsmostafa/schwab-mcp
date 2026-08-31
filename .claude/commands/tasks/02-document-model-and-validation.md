# Task 02 — Document model, schema validation, stable IDs

**Read first:** `.claude/commands/tasks/00-conventions.md` (frontmatter field
table, Facts, Type-specific fields, ID scheme).

**Depends on:** Task 01 (`src/lib/config`, `examples/knowledge/`).

## Objective

A `src/lib/document` module that can load a Markdown+YAML-frontmatter file
from disk, validate it against Ren's OKF-based schema, and return a typed
document object — or a structured error for malformed input. This is the
foundation every later task (search, MCP, capture, AI organization) reads
through; nothing later should parse frontmatter itself.

## Scope

In scope:
- Frontmatter schema (e.g. via `zod` or a hand-rolled validator — pick the smaller option) implementing the full field table in `00-conventions.md`: `type` required; `title`, `description`, `resource`, `tags`, `sources`, `generated`, `verified`, `status`, `stale_after`, `id`, `created_at`, `updated_at`, `tier`, `importance`, `confidence`, `pinned`, `valid_from`, `valid_to`, `entities`, `facts`, `supersedes`, `format` — all optional/typed per that table. `type` must accept the full ontology including `capture`, `fact`, `conflict`, `idea`, `task`, `insight`.
- **Nested shapes**, per `00-conventions.md`'s Facts and Type-specific fields sections — these are part of the schema, not left to the implementer to invent:
  - `entities[]`: `{ ref?: ULID, name: string, kind: string, role?: "subject" | "object" }`.
  - `facts[]`: `{ uid: ULID, subject: ULID | string, predicate: string, object?: ULID | string, object_literal?: string, confidence: number, valid_from: string, source_excerpt?: string, reviewed?: boolean, supersedes?: ULID[], contradicts?: ULID[], duplicates?: ULID[], reinforces?: ULID[] }` — exactly one of `object`/`object_literal` required.
  - `conflict`-type extra fields: `fact_uids: ULID[]` (≥2 entries), `status: "open" | "resolved" | "dismissed" | "accepted-both"`, `detected_by`, `resolution?`, `resolved_at?`.
  - `experiment`-type extra fields: `experiment_status: "proposed" | "active" | "concluded"`, `friction?: ULID`, `duration_days?: number`, `end_date?: string`, `outcome_result?: "positive" | "negative" | "neutral" | "inconclusive"` (required when `experiment_status: "concluded"`).
- ULID generation for new documents; a stable `id` field, independent of filename.
- Temporal/consolidation field validation: `valid_to` (if present) must not precede `valid_from`; `tier`, when present, only valid on `knowledge/` documents (never `inbox/`); `format`, when present, only valid on `capture`-type documents. `facts[].uid` is validated for **syntactic validity always, and uniqueness at whatever scope the caller is loading** — within one document by the single-file loader (no two `facts[]` entries in the same file share a `uid`), and across the whole data dir when loading the `knowledge/` tree (mirroring the existing `id`-collision check above: the whole-tree loader surfaces both offending paths; deferring full cross-run collision handling to Task 04's indexing is acceptable, same as for `id`). It is the fact's own identity, not a reference, so there is nothing for it to "resolve" to. The four edge arrays nested inside a `facts[]` entry (`supersedes`/`contradicts`/`duplicates`/`reinforces`), the document-level `supersedes` field, and `conflict.fact_uids`/`experiment.friction` **are** references (to other facts' `uid`s or other documents' `id`s) — this task validates that each entry is a syntactically valid ULID; whether it resolves to something that actually exists is a cross-document check, deferred to Task 04/05's indexing and reported by Task 15's `ren verify`.
- A loader: given a path (or the whole `knowledge/` tree), parse frontmatter + body, validate, return either a fully-valid typed document, or a typed error for a top-level problem (missing/malformed frontmatter, missing `type`). For a document that is otherwise valid but contains one or more malformed `facts[]` entries, return `{ document, warnings, invalidFacts }`: `document.facts` contains only the entries that validated, `invalidFacts` carries the rejected entries with their specific errors, and `warnings` is a flat list for logging/`ren verify` — the document as a whole is not failed by a bad fact entry. Never throw in a way that crashes a caller.
- Handling for documents missing optional fields (most fields are optional — only `type` is required per OKF).
- Unit tests (Vitest) covering: valid document, missing required `type`, malformed YAML, missing frontmatter entirely, duplicate `id` across two files, a `knowledge/` document with `tier` set correctly and one with `tier` on an `inbox/`-style capture rejected, a `valid_to` before `valid_from` rejected, a `facts[]` entry missing `uid`/`confidence` landing in `invalidFacts` with the rest of the document (and its other, valid facts) still loading successfully, a `facts[].uid` that happens to equal another fact's `uid` elsewhere in the data dir flagged as a uniqueness violation (not silently accepted as if it were a valid self-reference).

Out of scope: writing documents (Task 03/08), search indexing (Task 04), the knowledge graph itself (Task 05 — this task validates the *shape* of `entities`/`facts`/edges, it doesn't resolve or index them), any UI.

## Deliverables

- `src/lib/document/schema.ts` — the validator.
- `src/lib/document/load.ts` (or similar) — file → typed document.
- `src/lib/document/id.ts` — ULID helper.
- `src/lib/document/hash.ts` — content-hash helper (sha256 over raw file bytes; the hash itself is never written into the frontmatter it hashes — it's derived-only per `00-conventions.md` and only ever stored in SQLite by Task 04).
- Tests in `src/lib/document/*.test.ts`.

## Acceptance criteria

- [ ] Loading every file in `examples/knowledge/` succeeds and produces a stable `id` for each.
- [ ] A file with no frontmatter, or frontmatter missing `type`, produces a clear validation error identifying the file and the problem — not an unhandled exception.
- [ ] Two documents that happen to share an `id` are detected (at minimum: loader surfaces both paths; full collision handling can defer to Task 04's indexing).
- [ ] A document with `valid_to` before `valid_from`, or with `tier` set on a document loaded from `inbox/`, is rejected with a clear per-field error.
- [ ] A document with one malformed `facts[]` entry among otherwise-valid ones loads successfully with that entry excluded from `document.facts` and reported in `invalidFacts` — the document is not failed outright.
- [ ] `npm test` (or equivalent) passes and covers the cases listed above.

## Notes / assumptions

- Do not build a general-purpose schema DSL. One schema implementing the shared table plus the small set of type-specific extra-fields blocks in `00-conventions.md` is enough — simple yet powerful.
- This task validates shape only. Whether a `supersedes`/`contradicts`/`fact_uids`/`friction` reference actually resolves to something that exists is a cross-document concern, checked by Task 04/05's indexer and reported by Task 15's `ren verify` — don't build a resolver here. `facts[].uid` is the one exception worth stating twice: it's an identity, checked for uniqueness, never a reference to resolve.
