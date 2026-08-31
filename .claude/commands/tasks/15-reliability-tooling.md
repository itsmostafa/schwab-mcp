# Task 15 — Reliability tooling: doctor, status, verify, reindex, repair

**Read first:** `00-conventions.md` (the core invariant and its litmus
test — this task is where that test becomes an actual command).

**Depends on:** Task 02 (document schema — `verify`'s malformed-OKF check
runs it directly), Task 04 (derived index — the minimal `ren reindex` this
task generalizes), Task 05 (graph/provenance — what `verify` checks),
Task 09 (embedding backend — reachability check), Task 11 (consolidation —
open conflicts this task reports), Task 13 (hot context — the one store
this task must *not* flag as damaged when cold).

## Objective

Give the system a set of commands that make the core invariant — files are
truth, SQLite is acceleration, deletable and rebuildable — something a user
can actually verify and act on, not just trust. Detect malformed OKF, broken
links, stale indexes, missing embeddings, unresolved conflicts, and database
inconsistencies; repair what's safely repairable without ever touching
`inbox/` or `knowledge/` files.

## Scope

In scope:
- `ren doctor`: environment/config sanity check — config validity, data-dir
  structure and permissions (`inbox/`, `knowledge/`, `runtime/`, `models/`
  all present and writable), `PRAGMA integrity_check` on `runtime/ren.db`,
  WAL mode confirmed active, model-cache presence (Task 09's `models/`
  contents match the configured embedding backend), and — only if an
  external `ai.provider` is configured — a reachability check. Exits
  non-zero on any failure, with each finding actionable (what's wrong, what
  to run next), not just logged.
- `ren status`: a fast, read-only summary — document/chunk/fact/vector
  counts, index freshness (documents whose on-disk `file_hash` differs from
  the indexed one), open `type: conflict` count, embedding coverage
  percentage (chunks with vs. without a `vectors` row), current
  `degraded`/fallback state, and `index_meta` version numbers. This data is
  suitable to power a future index-health indicator on Task 06's dashboard —
  that's not part of Task 06's own required scope, so don't treat its
  absence there as a Task 15 blocker; this task's job is to make the data
  available, not to build the widget.
- `ren verify`: the deeper, still read-only consistency check —
  - malformed OKF (frontmatter that fails Task 02's schema) across the
    whole data dir, reported per file, not just the first failure —
    including a `facts[].uid` that collides with another fact's `uid`
    elsewhere (a uniqueness violation, not a broken reference — `uid`s are
    identities, not references, per Task 02);
  - broken links (Task 05's `links` table entries with `to_doc IS NULL`, or
    a `sources`/document-level `supersedes`/fact-level `supersedes`-etc./
    `conflict.fact_uids`/`experiment.friction` reference that doesn't
    resolve to anything);
  - stale index rows (a `documents` row whose `file_hash` doesn't match a
    fresh hash of the file on disk — the file changed since last index);
  - missing embeddings (chunks with no matching `vectors` row while
    `search.semantic: true`, distinguished from the expected state when
    semantic search is off);
  - drift (a `facts.source_chunk_hash` that no longer matches its source
    chunk — the file was hand-edited after fact extraction);
  - unresolved conflicts (open `type: conflict` documents, surfaced as a
    work-queue list, not an error to auto-fix);
  - orphaned derived rows (a `vectors` row whose `chunk_hash` matches no
    live `chunks` row, `documents_fts`/`chunks_fts` row-count mismatch
    against their source tables).

  `ren verify` never modifies anything — it's the detect-and-report half of
  the doctor/repair pair.
- `ren reindex [--full]`: generalizes Task 04's minimal version — targeted
  reprocessing driven by `index_meta` version bumps (the common case), or a
  full wipe-and-rebuild from `knowledge/`/`inbox/` with `--full`. Must
  complete with **zero** AI/network calls when the embedding backend and AI
  provider are both unavailable (degrading per `00-conventions.md`, not
  failing) — this is the command the litmus test names directly.
- `ren repair`: fixes what `ren verify` found, but strictly scoped to
  **derived state only** — rebuilding damaged/orphaned index rows, graph
  projections, or vector rows from source files. `ren repair` must never
  write to a file under `inbox/` or `knowledge/`; a malformed OKF file or a
  broken link is a file-level problem and stays in `ren verify`'s report for
  a human (or an explicit write-path call — Task 08) to fix, not something
  `repair` silently patches. Test this as a hard boundary, not a convention.
- Tests: `ren verify` against a deliberately corrupted fixture set (a
  truncated `ren.db`, an orphaned vector row, a document with a dangling
  `supersedes` reference, a hand-edited document whose `source_chunk_hash`
  no longer matches) reports every seeded problem and nothing else; `ren
  repair` after such corruption restores derived-state correctness while
  leaving every file under `inbox/`/`knowledge/` byte-identical to before,
  verified by hash comparison, not just spot-checking; `ren reindex --full`
  from an empty database completes with zero AI/network calls (mock and
  assert zero invocations) and produces content-identical output to a second
  independent full rebuild, excluding **only** the named operational
  keys in `00-conventions.md`'s Version tracking section (`indexed_at`,
  `last_full_reindex_at`, `index_revision`, `embedded_at`) — every other
  column, including file-projected `created_at`/`updated_at`/`asserted_at`,
  must match exactly; vectors compared by `embedding_version` + cosine
  tolerance per that same section; `ren doctor` and `ren status` do not flag
  a cold/empty
  `hot_context` table (Task 13) as a problem; `ren status` counts match a
  hand-verified fixture set exactly.

Out of scope: fixing file-level problems automatically (explicitly excluded
from `ren repair`'s scope above), a delete tool (still excluded per
`00-conventions.md`), backups/export (Task 16's `ren export`), license/deployment docs (Task 16).

## Deliverables

- `src/lib/reliability/doctor.ts`, `status.ts`, `verify.ts`, `repair.ts`.
- `ren doctor` / `ren status` / `ren verify` / `ren reindex` / `ren repair`
  CLI commands (extending whatever CLI entry point Task 01/04 established).
- Tests.

## Acceptance criteria

- [ ] `ren verify` against a fixture set with seeded problems (malformed
      OKF, broken link, orphaned vector row, dangling `supersedes`
      reference, drifted `source_chunk_hash`) reports every one of them,
      identifiable by file/row, and nothing else on a clean fixture set.
- [ ] `ren repair` fixes derived-state damage from that same fixture set
      and leaves every source file under `inbox/`/`knowledge/`
      byte-identical to before repair, verified by hash.
- [ ] `ren reindex --full` from an empty `runtime/` makes zero AI/network
      calls and reconstructs a system functionally equivalent to a second
      independent full rebuild.
- [ ] `ren doctor` exits non-zero with an actionable message on each of:
      bad config, missing/unwritable data-dir subfolder, failed
      `integrity_check`, WAL not active, missing model cache when semantic
      search is enabled, unreachable configured `ai.provider`.
- [ ] `ren status` counts and freshness/coverage percentages match a
      hand-verified fixture set exactly, and neither `doctor` nor `status`
      treats an empty hot-context cache as a fault.
- [ ] `ren repair` never writes to any file under `inbox/` or `knowledge/`
      — verified by a test that corrupts derived state, runs repair, and
      hashes every source file before and after.

## Notes / assumptions

- This task is what turns "files are truth, SQLite is acceleration" from a
  design principle into something a user can run and trust — treat the
  litmus-test acceptance criteria above as non-negotiable, not aspirational.
- Resist folding file-level auto-fixing into `ren repair` even when it would
  be easy (e.g. "just regenerate the missing `id`") — that blurs the
  derived/canonical boundary this whole task exists to protect. Report it
  and let a human or the normal write path (Task 08) fix it instead.
