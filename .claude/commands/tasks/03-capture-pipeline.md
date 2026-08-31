# Task 03 — Capture pipeline (frictionless input, immutable inbox)

**Read first:** `00-conventions.md`.

**Depends on:** Task 01 (config), Task 02 (document model, for the `id`/timestamp helpers — captures are the source; they still get an `id`).

## Objective

Let a user (or, later, an MCP client) submit raw text with zero required
structure, and have it land immutably in `$SIA_DATA_DIR/inbox/`. This is the
single most important interaction in the product per the spec — get it
simple and reliable before anything fancier.

## Scope

In scope:
- `src/lib/capture`: given raw text (+ optional format hint: plain/markdown/html), write one new immutable file under `inbox/`, named/timestamped so ordering is obvious (e.g. `inbox/2026-08-30-<ulid>.md`), with minimal frontmatter (`id`, `created_at`, `type: capture`, raw `format`). Body is the raw text, untouched.
- Atomic write (write to temp file, rename) so a crash never leaves a partial capture.
- Append-only guarantee: nothing in this task ever edits or deletes an existing inbox file.
- One Astro API route (e.g. `src/pages/api/capture.ts`) accepting POST text and calling the lib function.
- One minimal capture UI island (a textarea + submit) on the homepage — no formatting toolbar, no required fields, matches "Do not force the user to choose folders/tags/taxonomies before saving."
- Unit tests: capture produces exactly one new file, original text is byte-for-byte preserved, concurrent captures don't collide (two captures in the same tick get distinct files).

Out of scope: classification into observation/friction/decision/etc. (Task 09 — AI organization), the Kaizen view UI (Task 10), MCP capture tools (Task 06/07 expose this same lib function, don't reimplement it).

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

- This task deliberately does not touch `knowledge/` at all — captures only ever land in `inbox/`. Turning captures into structured `knowledge/` documents is Task 09's job, later, and explicitly preserves the source capture rather than replacing it.
