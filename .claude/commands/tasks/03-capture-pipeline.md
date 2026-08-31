# Task 03 — Capture pipeline (frictionless input, immutable Bronze inbox)

**Read first:** `00-conventions.md` (the pipeline: Capture → Bronze → Silver/OKF → consolidation → Gold/OKF).

**Depends on:** Task 01 (config), Task 02 (document model, for the `id`/timestamp helpers — captures are the source; they still get an `id`. Content hashing is computed later, when Task 04 indexes the capture — this task's write path does not compute or store one).

## Objective

Let a user (or, later, an MCP client) submit raw text with zero required
structure, and have it land immutably in `$REN_DATA_DIR/inbox/` as a Bronze
record. This is the single most important interaction in the product — get
it simple and reliable before anything fancier, and get its immutability
guarantee airtight, because every later task (organization, consolidation,
provenance) depends on Bronze never changing under them.

## Scope

In scope:
- `src/lib/capture`: given raw text (+ optional format hint: plain/markdown/html), write one new immutable file under `inbox/`, named/timestamped so ordering is obvious (e.g. `inbox/2026-08-30-<ulid>.md`), with minimal frontmatter (`id`, `created_at`, `type: capture`, raw `format`). **Never** a `tier` field — `tier` is a `knowledge/`-only concept per `00-conventions.md`; Bronze has no tier. Body is the raw text, untouched.
- Atomic write (write to temp file, rename) so a crash never leaves a partial capture.
- Append-only guarantee: nothing in this task ever edits or deletes an existing inbox file. This guarantee must hold not just for this task but through every later task that reads captures (organization in Task 10, consolidation in Task 11) — say so as a note for implementers of those tasks.
- One Astro API route (e.g. `src/pages/api/capture.ts`) accepting POST text and calling the lib function.
- One minimal capture UI island (a textarea + submit) on the homepage — no formatting toolbar, no required fields, matches "Do not force the user to choose folders/tags/taxonomies before saving."
- Unit tests: capture produces exactly one new file, original text is byte-for-byte preserved, concurrent captures don't collide (two captures in the same tick get distinct files), a captured file never validates with a `tier` field present.

Out of scope: classification into observation/friction/decision/etc. (Task 10 — AI organization), the Kaizen view UI (Task 14), MCP capture tools (Task 07/08 expose this same lib function, don't reimplement it).

## Deliverables

- `src/lib/capture/index.ts`.
- `src/pages/api/capture.ts`.
- A capture UI island component.
- Tests.

## Acceptance criteria

- [ ] Submitting text via the UI or a direct POST to the API route produces exactly one new immutable file in `inbox/` with the raw text preserved verbatim.
- [ ] No existing file is ever modified by capture.
- [ ] The user is never asked to pick a folder, tag, title, or category before the capture succeeds.
- [ ] Killing the process mid-write (simulate: write-then-rename pattern reviewed, or a test that interrupts the temp-file step) never leaves a corrupt or half-written file visible in `inbox/`.

## Notes / assumptions

- This task deliberately does not touch `knowledge/` at all — captures only ever land in `inbox/`. Turning captures into structured `knowledge/` documents is Task 10's job, later, and explicitly preserves the source capture rather than replacing it.
- Every later stage of the pipeline (Silver, consolidation, Gold) traces back to a Bronze capture via `sources`. If this task's immutability guarantee ever breaks, provenance breaks everywhere downstream — treat it as load-bearing, not incidental.
- Content hashing, indexing, and provenance lookup by capture path are all Task 04/05's job, not this task's — Task 04 indexes `inbox/` unconditionally (not optionally) precisely so those later tasks have something to read.
