# Task 08 — MCP write tools and permission levels

**Read first:** `00-conventions.md` (forward-only consolidation edges — this task is where "supersede, don't overwrite" first gets implemented).

**Depends on:** Task 07 (read-only MCP server), Task 03 (capture lib), Task 02 (document model), Task 04 (incremental indexer — write tools trigger it directly), Task 05 (knowledge graph, for provenance stamping and `document_edges`).

## Objective

Add write capability to the MCP server behind an explicit permission model,
so not every agent connection gets write/delete access by default.

## Scope

In scope:
- Capability levels: `knowledge.read`, `knowledge.search`, `capture.create`, `knowledge.update`, `knowledge.admin`. Implementation: a config-driven allowlist (e.g. in `ren.config.ts` or a separate `mcp.permissions.json`) naming which levels this server instance grants; tools check their required level before executing and return a clear permission-denied error otherwise. Read-only is the out-of-the-box default (matches Task 07).
- Write tools:
  - `capture(text, format?)` / `kaizen_dump(text)` — both wrap Task 03's capture lib (kaizen_dump is just capture with a hint, not a separate storage path).
  - `create_note(...)`, `record_decision(...)`, `create_experiment(...)` — construct valid `knowledge/` documents via Task 02's schema (assign `id`, `created_at`, `tier: silver` unless explicitly promoted, `generated: {by: "agent", model, at}` provenance), write atomically (temp file + rename), then trigger Task 04's incremental indexer (which in turn feeds Task 05's graph projection).
  - `update_document(...)` — two distinct semantics, not one blurred operation: a **cosmetic** update (typo fix, formatting, tag addition — same `id`, bumped `updated_at`, no meaning change) edits in place; a **substantive** update that changes what a document *claims* must instead write a **new** document whose top-level `supersedes` field names the old document's `id` (per `00-conventions.md`'s Forward-only consolidation edges — this is the **document-level** edge, projected by Task 05 into `document_edges`; it is a different thing from the fact-level `facts[].supersedes` edges Task 11 writes, which reference fact `uid`s, not document `id`s), never mutate the old document's claims directly. The tool's contract: preserve `id`/`created_at` on cosmetic edits, refuse (or require an explicit `supersede: true` flag that creates a new document) on substantive ones. Every update — cosmetic or superseding — appends a `runtime/changelog.jsonl` entry (id, timestamp, diff summary — not a full VCS; this changelog is a best-effort operational log, not the canonical audit trail, per `00-conventions.md` — losing it loses convenience, not history, since the supersede chain itself is the recoverable record).
- Concurrency safety: a simple write lock (e.g. per-file or per-`knowledge/`-directory advisory lock) so two simultaneous writers can't corrupt a file. Doesn't need to be distributed — single Mac mini, single process is the target. Task 16 revisits this under more realistic concurrent load (website + MCP + watcher + Task 11's consolidation all live); this task only needs it to hold for two racing MCP writes.
- Tests: permission denial when a level isn't granted, `create_note` produces a schema-valid document findable by search immediately after, a cosmetic `update_document` never loses the prior `created_at`/`id`, a substantive update produces a new document with a correct `supersedes` edge and leaves the old document's file untouched, concurrent writes to the same document don't corrupt it (simulate two near-simultaneous calls).

Out of scope: a permissions admin UI (explicitly excluded — see "what not to build" in `00-conventions.md`), full audit history / diffing beyond the changelog (Task 16 can deepen this), deletion (permanent deletion is privileged or avoided entirely — do not implement a delete tool in this task; if one is added later it belongs at `knowledge.admin` behind extra confirmation, deferred), consolidation logic itself — deciding *when* something is a duplicate/contradiction (Task 11); this task only implements the primitive (supersede) that consolidation will call.

## Deliverables

- `src/mcp/permissions.ts` — capability check.
- Write tools added to `src/mcp/server.ts` (or `tools.ts`).
- `src/lib/document/write.ts` — shared atomic write + lock helper, including the supersede primitive, used by both write tools and Task 03/10/11.
- `runtime/changelog.jsonl` writer.
- Tests.

## Acceptance criteria

- [ ] With only `knowledge.read`/`knowledge.search` granted, every write tool call is rejected with a clear error, not a silent no-op or a crash.
- [ ] With `capture.create` and `knowledge.update` granted, all write tools function and their output is immediately visible to search (Task 04), the graph (Task 05), and the website (Task 06).
- [ ] No write tool ever deletes a file; a cosmetic `update_document` preserves prior `id`/`created_at`; a substantive update creates a new superseding document rather than mutating the old one's claims. Every write appends a changelog entry.
- [ ] A concurrency test (two writes racing) leaves the target file valid and parseable, not interleaved/corrupted.

## Notes / assumptions

- Keep the permission check itself trivial (a level → boolean lookup). Don't build role/user management — there's one operator (the Mac mini owner) and possibly multiple agent connections, not multiple human users.
- The cosmetic/substantive distinction in `update_document` is a judgment call the caller makes explicitly (a flag), not something this task infers from a diff — inferring "did the meaning change" is Task 11's job when it runs consolidation over incoming content; this task just makes sure the *correct primitive exists* for both cases.
