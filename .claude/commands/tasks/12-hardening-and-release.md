# Task 12 — Hardening, deployment, open-source release polish

**Read first:** `00-conventions.md`.

**Depends on:** everything (01-11) — this is the closing task, not a fixed scope like the others. Treat the checklist below as a menu to work through and report on, not all of which may be needed.

## Objective

Take the working system from Tasks 01-11 and make it something a stranger
can clone, run, and trust with private data.

The end-to-end flow this whole project exists to make effortless:

```text
User: "Brain dump: I keep losing context when I switch between projects.
Maybe Ren should save a quick checkpoint whenever I switch."
    → Ren preserves raw capture
    → AI derives an observation + a potential experiment, linked to related
      existing notes on attention switching
    → User later asks an agent: "What problems have I repeatedly mentioned
      about switching projects?"
    → MCP returns the relevant synthesized knowledge and its raw sources.
```

Confirm this works end-to-end using only the README and example data before
considering this task done.

## Scope

Work through, dropping anything that genuinely isn't needed rather than
padding for completeness (simple yet powerful applies here too):

- **License**: surface this to the user explicitly (don't silently pick one) — MIT is the likely default for a project like this, but confirm. Add `LICENSE`.
- **Tailscale deployment docs**: `docs/deployment.md` covering running on an always-on Mac mini behind Tailscale, and Tailscale Serve for exposing the website within the tailnet. No inbound router ports, no public exposure by default — confirm the app doesn't bind `0.0.0.0` unexpectedly.
- **Concurrency safety review**: revisit Task 07's write lock under more realistic conditions (website + MCP server + file-watcher all live at once) — this is the point where a bug there actually bites.
- **Backups/export**: `ren export` — a straightforward archive of `knowledge/` (+ optionally `inbox/`) to a single portable bundle. Not a sync service, just an export command.
- **`ren validate`**: run schema validation (Task 02) across the whole data dir and report problems, without needing a full reindex.
- **`ren doctor`**: sanity-check config, data dir permissions/structure, DB integrity, and (if configured) AI provider reachability.
- **Audit/change history**: decide whether Task 07's `changelog.jsonl` is enough or needs a small viewer/CLI (`ren log`) — don't build more than the spec's "maintain a change log where practical" asks for.
- **Test coverage pass**: revisit the "Testing" section's list end-to-end (parsing, schema validation, indexing, reindex determinism, stable IDs, provenance preservation, capture append behavior, MCP read operations, safe writes, malformed input) and fill any real gaps — using sample/demo data only, never private fixtures, per the spec.
- **README / repo polish**: architecture overview, dev setup, example knowledge directory pointer, example OKF documents, security/privacy notes, MCP configuration example, contribution guide, screenshots of the website (Task 05) and Kaizen view (Task 10).
- **Sanity check against "what not to build"**: confirm nothing built across Tasks 01-11 drifted toward any of: a collaborative Notion competitor, a block editor, a proprietary rich-text format, a cloud-first backend, mandatory accounts, a complicated permissions admin UI, a heavy distributed architecture, a required vector database, microservices, Kubernetes, or a public SaaS platform. If something did, this task should flag it rather than quietly keep it.

## Deliverables

- `LICENSE`, `docs/deployment.md`, `ren export`/`validate`/`doctor` commands (as much as is genuinely useful — not all are required if the checklist above says otherwise).
- Updated `README.md` meeting the spec's open-source checklist.
- A closing summary of what was deferred beyond v1 and why, for whoever picks up Ren next.

## Acceptance criteria

- [ ] A person who has never seen the project can `git clone`, `npm install`, `npm run dev`, and reach the flow above (capture → structured knowledge → search → MCP query) using only the README and example data.
- [ ] License is present and was confirmed with the user, not assumed.
- [ ] Nothing in the app binds to a public interface or requires cloud auth by default.
- [ ] The "What not to build" list has been explicitly checked against, not just implicitly avoided.

## Notes / assumptions

- This task can reasonably span more than one session — split it further at that point rather than trying to land it all at once.
