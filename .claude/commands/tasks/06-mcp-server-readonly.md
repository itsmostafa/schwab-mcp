# Task 06 — MCP server: read-only tools

**Read first:** `00-conventions.md` (MCP transport section — stdio, markdown/JSON only, never HTML).

**Depends on:** Task 02 (document model), Task 04 (search).

## Objective

Stand up the MCP server as the agent-facing counterpart to Task 05's
website, reading `src/lib` directly. This task covers read-only tools only —
the safest default per the spec ("expose read-only operation as the safest
default"). Write tools are Task 07.

## Scope

In scope:
- `src/mcp/server.ts`: MCP server over stdio, registered under a `knowledge.read` / `knowledge.search` capability tier (permission model itself formalized in Task 07, but this task's tools only ever read).
- Tools:
  - `search_knowledge(query, limit?, filters?)` — wraps Task 04's `search()`.
  - `get_document(id_or_path)` — wraps Task 02's loader; returns raw frontmatter + body (markdown), not HTML.
  - `get_recent_notes(limit?)` — most recently created/updated documents.
  - `get_project(name)`, `get_decisions(project?)` — filtered document queries by `type`/`tags`; thin wrappers, no new storage.
  - `get_related(id_or_topic)` — naive version: markdown-link scan + shared-tags overlap. Good enough for v1; richer traversal is Task 11's job.
  - `get_context(query, max_tokens)` — the spec's highest-priority tool. v1 implementation: run `search_knowledge`, pull top-N full documents (respecting `max_tokens` via a simple truncation/budget, not true token counting unless a tokenizer is already a dependency), concatenate into one compact bundle with source attribution per chunk. Rerank/semantic/relationship-following can be TODO-commented for Task 11 — this version must still work and be useful on its own.
- A short `docs/mcp.md` documenting how to point Claude Code / Claude Desktop at this server (stdio config example) and confirming every tool returns markdown/JSON, never rendered HTML.
- Tests: each tool against `examples/knowledge/`, confirming outputs are plain markdown/JSON (no HTML tags), and that `get_context` respects the `max_tokens` budget.

Out of scope: capture/write tools (Task 07), permission enforcement beyond "this server only exposes read tools right now" (Task 07 adds the actual capability-level gating), embeddings/reranking (Task 11).

## Deliverables

- `src/mcp/server.ts` and one file per tool (or a single `tools.ts` if that stays simpler — prefer fewer files while it's still small).
- `docs/mcp.md`.
- Tests.

## Acceptance criteria

- [ ] An MCP-compatible client (Claude Code, Claude Desktop) configured per `docs/mcp.md` can call every tool listed above successfully against the example knowledge dir.
- [ ] No tool output contains HTML markup — plain markdown/JSON only.
- [ ] `get_context` never returns more than its `max_tokens` budget and cites which source documents contributed.
- [ ] This server has no route through which it could write to `knowledge/` or `inbox/` — verified by inspection, not just by tool list.

## Notes / assumptions

- Resist building a plugin/tool-registry abstraction for six-or-so tools — a flat list registered directly on the server is simple yet powerful enough here.
