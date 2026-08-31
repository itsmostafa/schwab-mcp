# Sia — shared conventions and locked decisions

Every task file in this directory assumes these decisions — read this file
before starting any numbered task. Each numbered task file is self-contained
and does not require reading `.claude/commands/project-implementation.md`
(the original, much longer product spec that this whole task breakdown was
derived from) — it's kept in the repo as background/provenance, not as
required reading per task. If a task file references something that turns
out to be missing or ambiguous, it's a bug in that task file: fix the task
file (pulling in whatever detail is missing, from the original spec if
needed) rather than falling back to reading the full spec every time.

Do not re-decide items on this list inside a task — if one turns out to be
wrong, update this file first so later tasks stay consistent, and say so in
your session summary.

These are working assumptions made to unblock incremental implementation, not
requirements handed down by the user unless marked otherwise. Push back and
change this file if a task reveals one is wrong.

## Design philosophy: simple yet powerful

The user has explicitly reinforced this. When a task's acceptance criteria
can be met with a smaller mechanism, prefer the smaller mechanism — a plain
SQL query over a query builder, a single file over a plugin system, one
provider adapter over an abstraction for six. Power should come from the
document model and the composability of a few well-chosen primitives (OKF
frontmatter, FTS5, MCP tools), not from configurability or layered
abstraction. If a task description below asks for more than it needs, trim
it and note the trim rather than building to the letter of the description.

## Repository layout

Single Astro application, not a monorepo. The spec's directory sketch shows a
top-level `packages/`, but nothing requires it, and splitting into publishable
packages now would be premature — revisit only in Task 12 if a real boundary
pain shows up (e.g. wanting to publish the MCP server standalone).

```
src/
  lib/          # core: config, document model, search, AI provider interface — no framework deps
  content/      # Astro content collections / page templates
  pages/        # Astro routes (website + capture/API endpoints)
  mcp/          # MCP server, imports src/lib directly
  components/   # Astro/UI islands
examples/
  knowledge/    # tracked, non-private demo OKF documents — used in docs, tests, screenshots
scripts/        # sia CLI entry points
```

`src/mcp` and `src/pages` both depend on `src/lib`; they must not depend on
each other. This is what keeps the MCP server from ever needing to scrape
rendered HTML — it reads the same document model the website reads.

## Package manager & tooling

- **npm**, per the spec's own quickstart (`npm install`, `npm run dev`). Don't introduce pnpm/yarn.
- **TypeScript**, strict mode.
- **Vitest** for unit tests (assumption — spec doesn't name a framework; Vitest fits the Astro/TS stack with no extra config burden).
- Node.js current LTS.

## Data directory

`SIA_DATA_DIR` env var, resolved in `src/lib/config`. Structure:

```
$SIA_DATA_DIR/
  inbox/     # raw immutable captures, one file per capture, append-only
  knowledge/ # OKF/Markdown structured knowledge (notes, projects, decisions, ...)
  runtime/   # SQLite db + indexes + logs — fully derived, safe to delete and rebuild
```

If `SIA_DATA_DIR` is unset, fall back to `.sia-data/` at the repo root,
**gitignored**, for local dev only. On startup, if the fallback is in use, log
a visible warning that this is not where private production data should live
(per the spec's privacy section — never let `.gitignore` be the only
protection).

`examples/knowledge/` (tracked, in-repo) is separate from both — it's demo
content for docs/tests/screenshots, never written to at runtime, never
containing real user data.

## Document model / OKF

Sia's frontmatter targets **OKF v0.2** (github.com/GoogleCloudPlatform/open-knowledge-format)
as the wire format, extended with Sia-specific fields. Reuse OKF's own field
names for concepts OKF already defines, rather than inventing parallel ones —
this is a correction to the illustrative frontmatter example in the spec,
which uses `created_by`/`verification` where OKF already has `generated`/`verified`:

| Field | Source | Notes |
|---|---|---|
| `type` | OKF (required) | note / project / concept / decision / observation / friction / experiment / question / person / resource / summary — see spec's Data Model section |
| `title`, `description`, `resource`, `tags` | OKF (recommended) | |
| `sources` | OKF | provenance: materials this concept derives from |
| `generated` | OKF | who/what produced this (human vs. agent, model id) |
| `verified` | OKF | trust tier: unverified → machine-confirmed → human-reviewed |
| `status`, `stale_after` | OKF | lifecycle / staleness |
| `id` | Sia extension | stable ULID, independent of file path — see below |
| `created_at`, `updated_at` | Sia extension | OKF doesn't define timestamps beyond `stale_after`; keep these |

Document loading/validation (Task 02) must implement this table as the schema,
not the spec's literal example block.

## Stable IDs

ULID, stored as `id` in frontmatter, assigned at creation and never changed.
Filenames may be renamed freely by AI reorganization; links resolve through
`id`, not path, wherever both exist.

## MCP transport

stdio by default (the standard local MCP client pattern — Claude Code /
Claude Desktop config), documented as the only supported mode through Task 7.
An HTTP-over-Tailscale transport is optional, deferred to Task 12, for agents
that aren't local subprocesses.

MCP tools return raw markdown/frontmatter/JSON — never rendered HTML. This is
a hard requirement (explicitly confirmed by the user), not just a spec
preference: agents must never be asked to scrape the website.

## AI provider

Default provider: `"none"`. Interface defined in Task 08 covers Anthropic,
OpenAI, Ollama, LM Studio, llama.cpp, MLX per the spec, but only **Anthropic**
and **Ollama** adapters get concrete implementations initially — the rest stay
interface-shaped stubs until something needs them. This keeps Task 08 scoped
to "prove the abstraction works end-to-end," not "implement six providers."

## Explicitly open / not decided here

- **License**: not chosen. Task 12 must surface this to the user rather than default silently — MIT is the common choice for a project like this, but it's a real decision (patent grant, copyleft, etc.), not a convention.
- **Deployment hostname / Tailscale Serve specifics**: left to Task 12, depends on the user's actual tailnet.
