# Task 16 — Hardening, deployment, open-source release polish

**Read first:** `00-conventions.md`.

**Depends on:** everything (01-15) — this is the closing task, not a fixed
scope like the others. Treat the checklist below as a menu to work through
and report on, not all of which may be needed.

## Objective

Take the working system from Tasks 01-15 and make it something a stranger
can clone, run, and trust with private data.

The end-to-end flow this whole project exists to make effortless:

```text
User: "Brain dump: I keep losing context when I switch between projects.
Maybe Ren should save a quick checkpoint whenever I switch."
    → Ren preserves the raw capture (Bronze, immutable)
    → AI derives an observation + a potential experiment (Silver), linked to
      related existing notes on attention switching via the knowledge graph
    → Consolidation reinforces the pattern across multiple mentions and
      promotes it toward Gold with full provenance intact
    → User later asks an agent: "What problems have I repeatedly mentioned
      about switching projects?"
    → MCP returns the relevant synthesized knowledge, its confidence/
      verification state, and its raw sources — via hybrid retrieval, not a
      page scrape.
```

Confirm this works end-to-end using only the README and example data before
considering this task done.

## Scope

Work through, dropping anything that genuinely isn't needed rather than
padding for completeness (simple yet powerful applies here too). Four items
are **mandatory release gates**, not menu items — License, no-public-bind
confirmation, the "what not to build" audit, and the core-invariant audit —
everything else below is judged against whether the project actually needs
it:

- **License**: surface this to the user explicitly (don't silently pick one) — MIT is the likely default for a project like this, but confirm. Add `LICENSE`.
- **Tailscale deployment docs**: `docs/deployment.md` covering running on an always-on Mac mini behind Tailscale, and Tailscale Serve for exposing the website within the tailnet. No inbound router ports, no public exposure by default — confirm the app doesn't bind `0.0.0.0` unexpectedly.
- **Concurrency safety review**: revisit Task 08's write lock under realistic conditions — website, MCP server, file-watcher, and Task 11's consolidation pass all live at once now. This is the point where a bug there actually bites; Task 11's frontmatter-editing writes (supersede edges) are the newest and least-exercised source of concurrent writes and deserve specific attention here.
- **Backups/export**: `ren export` — a straightforward archive to a single portable bundle. Default contents: `knowledge/` only. `inbox/` (raw Bronze captures — the most sensitive, highest-volume material) is **excluded by default** and included only with an explicit `--include-inbox` flag; document the tradeoff plainly — without it, an exported bundle's documents still carry `sources` paths pointing at captures the bundle doesn't contain, so provenance chains reference material the recipient won't have. `runtime/` (derived, rebuildable) and `models/` (re-fetchable, not knowledge) are excluded unconditionally, no flag to include them.
- **`ren doctor`/`ren status`/`ren verify`/`ren reindex`/`ren repair`**: these are Task 15's job, not this task's — confirm they're actually present, documented, and exercised end-to-end here rather than reimplementing any of them. If Task 15 was skipped or partially done, this is the task that must notice and either complete it or flag the gap explicitly.
- **Audit/change history**: decide whether Task 08's `changelog.jsonl` is enough or needs a small viewer/CLI (`ren log`) — don't build more than "maintain a change log where practical" asks for.
- **Test coverage pass**: revisit the full testing surface end-to-end — parsing, schema validation, content hashing, incremental indexing, reindex determinism (scoped per `00-conventions.md`), stable IDs, graph/provenance correctness, temporal queries, consolidation outcomes (reinforce/supersede/flag), hot-context cache invalidation, capture append/immutability, MCP read and write operations, permission enforcement, safe concurrent writes, malformed input handling — and fill any real gaps, using sample/demo data only, never private fixtures.
- **README / repo polish**: architecture overview, dev setup, example knowledge directory pointer, example OKF documents (including ones showing `tier`, `facts`, `supersedes`), security/privacy notes, MCP configuration example, contribution guide, screenshots of the website (Task 06) and Kaizen view (Task 14).
- **Sanity check against "what not to build"**: confirm nothing built across Tasks 01-15 drifted toward any of: a collaborative Notion competitor, a block editor, a proprietary rich-text format, a cloud-first backend, mandatory accounts, a complicated permissions admin UI, a heavy distributed architecture, a required (external) vector database, microservices, Kubernetes, or a public SaaS platform. If something did, this task should flag it rather than quietly keep it.
- **Core-invariant audit**: grep the finished codebase and docs for any language that treats SQLite/the derived index as authoritative, and for any code path that writes consolidation/temporal/provenance state *only* to SQLite without a corresponding file write. This is the project-wide version of Task 15's litmus test — confirm it holds everywhere, not just in the reliability commands built to test it.

## Deliverables

- `LICENSE`, `docs/deployment.md`, `ren export` command.
- Confirmation (or completion) that Task 15's reliability commands are in place and documented.
- Updated `README.md` meeting the open-source checklist above.
- A closing summary of what was deferred beyond v1 and why, for whoever picks up Ren next.

## Acceptance criteria

- [ ] A person who has never seen the project can `git clone`, `npm install`, `npm run dev`, and reach the flow above (capture → structured knowledge → consolidation → search/graph query → MCP query) using only the README and example data.
- [ ] License is present and was confirmed with the user, not assumed.
- [ ] Nothing in the app binds to a public interface or requires cloud auth by default.
- [ ] The "What not to build" list has been explicitly checked against, not just implicitly avoided.
- [ ] The core-invariant audit above has been run and its findings (or a clean result) are recorded in the closing summary.

## Notes / assumptions

- This task can reasonably span more than one session — split it further at that point rather than trying to land it all at once.
