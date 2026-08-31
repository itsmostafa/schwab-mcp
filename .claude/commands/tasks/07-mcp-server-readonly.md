# Task 07 — MCP server: read-only tools

**Read first:** `00-conventions.md` (MCP transport section — stdio, markdown/JSON only, never HTML; MCP permissions).

**Depends on:** Task 02 (document model), Task 04 (derived index), Task 05
(knowledge graph + provenance).

## Objective

Stand up the MCP server as the agent-facing counterpart to Task 06's
website, reading `src/lib` directly. This task covers read-only tools
only — the safest default ("expose read-only operation as the safest
default"). Write tools are Task 08.

## Scope

In scope:
- `src/mcp/server.ts`: MCP server over stdio, registered under a `knowledge.read` / `knowledge.search` capability tier (permission model itself formalized in Task 08, but this task's tools only ever read).
- Tools:
  - `search_knowledge(query, limit?, filters?)` — wraps Task 04's `search()`.
  - `get_document(id_or_path)` — wraps Task 02's loader; returns raw frontmatter + body (markdown), not HTML.
  - `get_recent_notes(limit?)` — most recently created/updated documents.
  - `get_project(name)`, `get_decisions(project?)` — filtered document queries by `type`/`tags`; thin wrappers, no new storage.
  - `get_related(id_or_topic)` — backed by Task 05's `getRelated` graph traversal (backlinks + shared entities), not a reimplementation. Ranking/proximity scoring on top of this is Task 12's job — this tool just returns the related set.
  - `get_provenance(id)` — wraps Task 05's `getProvenance`: the chain from a document or fact back to its Bronze/Silver sources, with author/agent, timestamp, confidence, and verification status at each step.
  - `get_timeline(entity_or_topic, as_of?)` — wraps Task 05's `getFacts`/`getFactHistory`: current state by default, or state as of a given date; the full supersession chain when asked.
  - `get_context(query, max_tokens)` — the highest-priority tool. v1 implementation: run `search_knowledge` (BM25 ranking, ties broken by `updated_at` desc — no weighted recency scoring until Task 12 defines one), pull top-N full documents, concatenate into one compact bundle with source attribution per chunk. Budget enforcement: use a conservative, documented characters-per-token ratio (deliberately over-estimating tokens per character, e.g. 3 chars/token for English prose) rather than a real tokenizer, so the `max_tokens` guarantee is provably honest even without true token counting — a real tokenizer is a Task 12 upgrade, not a v1 requirement. Hybrid ranking, semantic search, and relationship-following are TODO-commented for Task 12 — this version must still work and be useful on its own.
- A short `docs/mcp.md` documenting how to point Claude Code / Claude Desktop at this server (stdio config example) and confirming every tool returns markdown/JSON, never rendered HTML.
- Tests: each tool against `examples/knowledge/`, confirming outputs are plain markdown/JSON (no HTML tags), that `get_context` respects the `max_tokens` budget, and that `get_provenance`/`get_timeline` correctly surface a hand-authored fixture supersession chain (reuse Task 05's fixtures).

Out of scope: capture/write tools (Task 08), permission enforcement beyond "this server only exposes read tools right now" (Task 08 adds the actual capability-level gating), hybrid/semantic search ranking (Task 12).

## Deliverables

- `src/mcp/server.ts` and one file per tool (or a single `tools.ts` if that stays simpler — prefer fewer files while it's still small).
- `docs/mcp.md`.
- Tests.

## Acceptance criteria

- [ ] An MCP-compatible client (Claude Code, Claude Desktop) configured per `docs/mcp.md` can call every tool listed above successfully against the example knowledge dir.
- [ ] No tool output contains HTML markup — plain markdown/JSON only.
- [ ] `get_context` never returns more than its `max_tokens` budget and cites which source documents contributed.
- [ ] `get_provenance` on an AI-derived fixture document returns its full chain back to a Bronze capture; on a hand-authored fixture document it correctly terminates at that document, per Task 05. `get_timeline` on a fixture entity returns the correct current fact and, with `as_of` set to a past date, the correct historical one.
- [ ] This server has no route through which it could write to `knowledge/` or `inbox/` — verified by inspection, not just by tool list.

## Notes / assumptions

- Resist building a plugin/tool-registry abstraction for eight-or-so tools — a flat list registered directly on the server is simple yet powerful enough here.
