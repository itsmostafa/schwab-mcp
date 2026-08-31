# Task 05 — Website core: navigation, knowledge pages, search UI, dashboard

**Read first:** `00-conventions.md`.

**Depends on:** Task 02 (document model), Task 04 (search) — this task is where earlier plumbing becomes visible/usable.

## Objective

Make the knowledge store browsable and searchable as calm, mostly-static
HTML. This completes Phase 1: after this task, the end-to-end loop —
capture → see it saved → search it — should work through the website.

Design direction: minimal, calm, information-dense without feeling
cluttered, fast, keyboard-friendly, dark-mode friendly, suitable for
long-form reading. Not a generic SaaS dashboard, not a blog — closer to a
private knowledge environment. Avoid visual gimmicks.

## Scope

In scope:
- Navigation matching (or a trimmed version of) the spec's suggested nav: Home, Inbox, Knowledge, Search at minimum. Projects/Kaizen/Decisions/Experiments can be thin placeholders or omitted until their data exists (Tasks 09-10) — don't build empty chrome for features with no data yet; a flat "Knowledge" listing filterable by type covers Projects/Decisions for now.
- Statically/server-rendered knowledge document pages: render frontmatter (type, tags, verification state, timestamps) + Markdown body for any document under `knowledge/`.
- An Inbox view listing raw captures (read-only, reverse-chronological).
- A search island (client-side input, calls a small search API route backed by Task 04's `search()`), server-rendered results list is also fine if simpler.
- A home dashboard with a small subset of the spec's suggested widgets that actually have data at this point: Recently updated knowledge, Unprocessed captures. Leave the rest (Active projects, Open questions, Recurring friction, Active experiments) as later additions once Tasks 09-10 produce that data — do not fake empty sections.
- Baseline styling: minimal, dark-mode friendly (respect `prefers-color-scheme`), readable typography for long-form Markdown. No design system dependency — plain CSS is enough.

Out of scope: Kaizen view (Task 10), graph visualization, command palette, AI query interface (later, optional per spec's own "islands" list — not required for v1).

## Deliverables

- `src/layouts/`, `src/pages/index.astro`, `src/pages/inbox.astro`, `src/pages/knowledge/[...slug].astro`, `src/pages/search.astro` (or equivalent routing).
- Search UI island.
- Base stylesheet.

## Acceptance criteria

- [ ] Every document in `examples/knowledge/` is reachable by navigation and renders correctly, including its frontmatter (type, tags, verification, timestamps).
- [ ] The Inbox page shows captures made via Task 03, newest first.
- [ ] Typing a query in Search returns results sourced from Task 04's index, not a page scrape.
- [ ] Site is legible and usable in both light and dark OS theme settings.
- [ ] No client-side JS ships on pages that don't need it (knowledge pages, inbox should work with JS disabled; only search needs an island, and even that can degrade to a server round-trip).

## Notes / assumptions

- This is where "the website is a human-readable projection of the knowledge, not a place where knowledge is scraped for agents" gets tested — nothing built here should become a dependency of the MCP server (Task 06 reads `src/lib` directly, never this page layer).
