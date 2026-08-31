# Task 07 — MCP write tools and permission levels

**Read first:** `00-conventions.md`.

**Depends on:** Task 06 (read-only MCP server), Task 03 (capture lib), Task 02 (document model).

## Objective

Add write capability to the MCP server behind an explicit permission model,
so not every agent connection gets write/delete access by default.

## Scope

In scope:
- Capability levels per the spec: `knowledge.read`, `knowledge.search`, `capture.create`, `knowledge.update`, `knowledge.admin`. Implementation: a config-driven allowlist (e.g. in `ren.config.ts` or a separate `mcp.permissions.json`) naming which levels this server instance grants; tools check their required level before executing and return a clear permission-denied error otherwise. Read-only is the out-of-the-box default (matches Task 06).
- Write tools:
  - `capture(text, format?)` / `kaizen_dump(text)` — both wrap Task 03's capture lib (kaizen_dump is just capture with a hint, not a separate storage path).
  - `create_note(...)`, `record_decision(...)`, `create_experiment(...)` — construct valid `knowledge/` documents via Task 02's schema (assign `id`, `created_at`, `generated: {by: "agent", ...}` provenance), write atomically (temp file + rename), then trigger Task 04's incremental indexer.
  - `update_document(...)` — update semantics only: preserve `id`, bump `updated_at`, never silently delete prior content; if the update is substantive, keep the previous version retrievable (e.g. a lightweight `runtime/changelog.jsonl` entry: id, timestamp, diff summary — not a full VCS).
- Concurrency safety: a simple write lock (e.g. per-file or per-`knowledge/`-directory advisory lock) so two simultaneous writers can't corrupt a file. Doesn't need to be distributed — single Mac mini, single process is the target.
- Tests: permission denial when a level isn't granted, `create_note` produces a schema-valid document findable by search immediately after, `update_document` never loses the prior `created_at`/`id`, concurrent writes to the same document don't corrupt it (simulate two near-simultaneous calls).

Out of scope: a permissions admin UI (explicitly excluded in the spec's "What not to build"), full audit history / diffing (Task 12 can deepen this), deletion (spec: "permanent deletion should be privileged or avoided entirely" — do not implement a delete tool in this task; if one is added later it belongs at `knowledge.admin` behind extra confirmation, deferred).

## Deliverables

- `src/mcp/permissions.ts` — capability check.
- Write tools added to `src/mcp/server.ts` (or `tools.ts`).
- `src/lib/document/write.ts` — shared atomic write + lock helper used by both write tools and Task 03/09.
- `runtime/changelog.jsonl` writer.
- Tests.

## Acceptance criteria

- [ ] With only `knowledge.read`/`knowledge.search` granted, every write tool call is rejected with a clear error, not a silent no-op or a crash.
- [ ] With `capture.create` and `knowledge.update` granted, all write tools function and their output is immediately visible to search (Task 04) and the website (Task 05).
- [ ] No write tool ever deletes a file; `update_document` preserves prior `id`/`created_at` and appends a changelog entry.
- [ ] A concurrency test (two writes racing) leaves the target file valid and parseable, not interleaved/corrupted.

## Notes / assumptions

- Keep the permission check itself trivial (a level → boolean lookup). Don't build role/user management — there's one operator (the Mac mini owner) and possibly multiple agent connections, not multiple human users.
