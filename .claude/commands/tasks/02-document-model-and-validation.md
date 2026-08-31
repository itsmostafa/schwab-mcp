# Task 02 — Document model, schema validation, stable IDs

**Read first:** `.claude/commands/tasks/00-conventions.md` (frontmatter field table, ID scheme).

**Depends on:** Task 01 (`src/lib/config`, `examples/knowledge/`).

## Objective

A `src/lib/document` module that can load a Markdown+YAML-frontmatter file
from disk, validate it against Ren's OKF-based schema, and return a typed
document object — or a structured error for malformed input. This is the
foundation every later task (search, MCP, capture, AI organization) reads
through; nothing later should parse frontmatter itself.

## Scope

In scope:
- Frontmatter schema (e.g. via `zod` or a hand-rolled validator — pick the smaller option) implementing the field table in `00-conventions.md`: `type` required; `title`, `description`, `resource`, `tags`, `sources`, `generated`, `verified`, `status`, `stale_after`, `id`, `created_at`, `updated_at` all optional/typed per that table.
- ULID generation for new documents; a stable `id` field, independent of filename.
- A loader: given a path (or the whole `knowledge/` tree), parse frontmatter + body, validate, return `{ id, type, frontmatter, body, path, raw }` or a typed validation error — never throw on malformed input in a way that crashes a caller; malformed files should be reportable, not fatal.
- Handling for documents missing optional fields (most fields are optional — only `type` is required per OKF).
- Unit tests (Vitest) covering: valid document, missing required `type`, malformed YAML, missing frontmatter entirely, duplicate `id` across two files, a document using the spec's old example field names (`created_by`/`verification`) to confirm they're rejected or clearly flagged as non-canonical.

Out of scope: writing documents (Task 03/07), search indexing (Task 04), any UI.

## Deliverables

- `src/lib/document/schema.ts` — the validator.
- `src/lib/document/load.ts` (or similar) — file → typed document.
- `src/lib/document/id.ts` — ULID helper.
- Tests in `src/lib/document/*.test.ts`.

## Acceptance criteria

- [ ] Loading every file in `examples/knowledge/` succeeds and produces a stable `id` for each.
- [ ] A file with no frontmatter, or frontmatter missing `type`, produces a clear validation error identifying the file and the problem — not an unhandled exception.
- [ ] Two documents that happen to share an `id` are detected (at minimum: loader surfaces both paths; full collision handling can defer to Task 04's indexing).
- [ ] `npm test` (or equivalent) passes and covers the cases listed above.

## Notes / assumptions

- Do not build a general-purpose schema DSL. A single schema for the fields in `00-conventions.md`, plus room to add a `type`-specific extra-fields bag later, is enough — simple yet powerful.
