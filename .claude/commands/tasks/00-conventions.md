# Ren — shared conventions and locked decisions

Every task file in this directory assumes these decisions — read this file
before starting any numbered task. Each numbered task file is self-contained:
it does not require reading anything outside this directory. If a task file
references something that turns out to be missing or ambiguous, that's a bug
in that task file: fix the task file (pulling in whatever detail is missing)
rather than inventing a second source of truth elsewhere.

Do not re-decide items on this list inside a task — if one turns out to be
wrong, update this file first so later tasks stay consistent, and say so in
your session summary.

These are working assumptions made to unblock incremental implementation, not
requirements handed down by the user unless marked otherwise. Push back and
change this file if a task reveals one is wrong.

## What Ren is

Ren (named after the ancient Egyptian concept of perception, knowledge, and
understanding) is a **private, local-first, agent-native personal knowledge
system**. It is not a Notion clone. Users dump raw thoughts, notes,
observations, project ideas, decisions, and Kaizen-style brain dumps into the
system with zero required organization; AI agents continuously organize,
connect, synthesize, and maintain the resulting knowledge base. The
application is open source; user data is private and never committed to the
application repository.

Ren has two interfaces onto the same underlying knowledge store:

1. **Astro website** — human-readable projection of the knowledge.
2. **MCP server** — machine-readable interface to the same knowledge.

Agents must never scrape or parse the rendered website to retrieve
knowledge — they go through MCP. The website is not a dependency of the MCP
server, and vice versa (see repository layout below).

## The core invariant

> **Files are truth. SQLite is acceleration.**

OKF/Markdown documents under `knowledge/` and `inbox/` are the single
canonical source of truth. SQLite (search index, vector store, knowledge
graph, provenance records, hot-context cache) is **entirely derived
acceleration infrastructure** that sits beside the pipeline below, never
inside it. Every piece of state that lives only in SQLite must be safely
deletable and fully rebuildable from the knowledge files, with one narrow,
explicitly-declared exception (the hot-context cache — see Task 13).

**Litmus test**, restated by name in Task 15 and used as an acceptance
criterion throughout: `rm -rf $REN_DATA_DIR/runtime && ren reindex` must lose
**nothing** — not search results, not the knowledge graph, not fact history,
not stated confidence, not unresolved conflicts, not consolidation decisions.
If a task's design can't survive that command, the design is wrong, not the
test.

The practical rule this implies: **if it's an LLM judgment, a human decision,
or unrecoverable without a network/model call, it lives in a file. If it's a
purely deterministic function of files (a hash, a chunk boundary, an
embedding, a BM25 score, a graph projection), it may live only in SQLite.**

## The pipeline

```
Capture → Bronze (inbox/, immutable) → Silver/OKF → consolidation → Gold/OKF
```

- **Bronze** = raw, immutable captures in `inbox/`. Append-only, never
  edited, never deleted, never labeled with a `tier`.
- **Silver** = structured AI-derived knowledge in `knowledge/`
  (`tier: silver`) — observations, friction, ideas, draft facts, still
  provisional.
- **Consolidation** = the pass that deduplicates, detects contradictions,
  reinforces matching knowledge, and promotes changed/confirmed knowledge
  forward (Task 11).
- **Gold** = curated, consolidated knowledge in `knowledge/` (`tier: gold`) —
  decisions, confirmed facts, summaries, active experiments, high-value
  knowledge.

SQLite, FTS5, embeddings, the knowledge graph, and the hot-context cache sit
**beside** this pipeline as derived acceleration infrastructure, not inside
the source-of-truth layer. If Ren disappears, the user should still have a
directory full of useful Markdown files that mean exactly what they meant
before.

## Design philosophy: simple yet powerful

The user has explicitly reinforced this. When a task's acceptance criteria
can be met with a smaller mechanism, prefer the smaller mechanism — a plain
SQL query over a query builder, a single file over a plugin system, one
provider adapter over an abstraction for six. Power should come from the
document model and the composability of a few well-chosen primitives (OKF
frontmatter, FTS5, the graph, MCP tools), not from configurability or layered
abstraction. If a task description below asks for more than it needs, trim
it and note the trim rather than building to the letter of the description.

Do not introduce external infrastructure unless genuinely necessary. Ren
targets a **single-user, always-on Mac mini**: local embeddings, low latency,
simple recovery, minimal operational complexity. No embedding compression in
v1 — the vector layer is abstracted (see below) so it can be added later
without a rewrite.

## Repository layout

Single Astro application, not a monorepo — nothing requires publishable
packages yet; revisit only in Task 16 if a real boundary pain shows up (e.g.
wanting to publish the MCP server standalone).

```
src/
  lib/          # core: config, document model, index, graph, retrieval, AI provider — no framework deps
  content/      # Astro content collections / page templates
  pages/        # Astro routes (website + capture/API endpoints)
  mcp/          # MCP server, imports src/lib directly
  components/   # Astro/UI islands
examples/
  knowledge/    # tracked, non-private demo OKF documents — used in docs, tests, screenshots
scripts/        # ren CLI entry points
```

`src/mcp` and `src/pages` both depend on `src/lib`; they must not depend on
each other. This is what keeps the MCP server from ever needing to scrape
rendered HTML — it reads the same document model and derived index the
website reads.

## Package manager & tooling

- **npm**, per the original quickstart (`npm install`, `npm run dev`). Don't introduce pnpm/yarn.
- **TypeScript**, strict mode.
- **Vitest** for unit tests (fits the Astro/TS stack with no extra config burden).
- Node.js current LTS.

## Data directory

`REN_DATA_DIR` env var, resolved in `src/lib/config`. Structure:

```
$REN_DATA_DIR/
  inbox/     # raw immutable captures (Bronze), one file per capture, append-only
  knowledge/ # OKF/Markdown structured knowledge (Silver + Gold), one tree
  runtime/   # SQLite db + indexes + logs + hot-context cache — fully derived, safe to delete and rebuild
  models/    # cached local embedding model weights — NOT derived, NOT private, NOT deleted by `ren repair`
```

`models/` is deliberately outside `runtime/`: it holds downloaded model
weights, not knowledge, and deleting `runtime/` must never create a network
dependency (re-downloading a model) as a side effect of an otherwise-offline
rebuild. It's safe to exclude from `ren export` by default and safe to
re-fetch if lost, but it is not part of the "derived, disposable" contract
the way `runtime/` is.

If `REN_DATA_DIR` is unset, fall back to `.ren-data/` at the repo root,
**gitignored**, for local dev only. On startup, if the fallback is in use, log
a visible warning that this is not where private production data should live
— never let `.gitignore` be the only protection.

`examples/knowledge/` (tracked, in-repo) is separate from all of the above —
it's demo content for docs/tests/screenshots, never written to at runtime,
never containing real user data.

## Document model / OKF

Ren's frontmatter targets **OKF v0.2** (github.com/GoogleCloudPlatform/open-knowledge-format)
as the wire format, extended with Ren-specific fields. Reuse OKF's own field
names for concepts OKF already defines, rather than inventing parallel ones.

| Field | Source | Notes |
|---|---|---|
| `type` | OKF (required) | note / project / concept / decision / observation / friction / experiment / question / person / resource / summary / capture / **fact** / **conflict** — see Data model below |
| `title`, `description`, `resource`, `tags` | OKF (recommended) | |
| `sources` | OKF | provenance: materials this document derives from (typically `inbox/` capture paths) |
| `generated` | OKF | who/what produced this — `{ by: "human" \| "agent", model?, prompt_version?, at }` |
| `verified` | OKF | human trust tier: `unverified` → `machine-confirmed` → `human-reviewed` |
| `status`, `stale_after` | OKF | lifecycle / staleness |
| `id` | Ren extension | stable ULID, independent of file path — see below |
| `created_at`, `updated_at` | Ren extension | OKF doesn't define timestamps beyond `stale_after`; keep these |
| `tier` | Ren extension | `silver` \| `gold`, `knowledge/` documents only — see Tiers below |
| `importance` | Ren extension | `0..1`, **stated** value — hand- or AI-set at write time. When absent, a *derived* default is computed from the by-type table below; the stored value, when present, is authoritative and need not match that formula — see Importance below |
| `confidence` | Ren extension | `0..1`, **stated** extractor/author confidence in this document as a whole — orthogonal to `verified` (see Two time axes below) |
| `pinned` | Ren extension | user pin; keeps a document in the hot-context working set regardless of recency/score (Task 13) |
| `valid_from`, `valid_to` | Ren extension | world-time validity window of this document's claims — see Two time axes below |
| `entities` | Ren extension | entity-linking decisions for the knowledge graph — array of `{ ref?: <entity ULID>, name: string, kind: string, role?: "subject" \| "object" }`; `ref` omitted means a mention-only entity — see Task 05 |
| `facts` | Ren extension | inline structured fact tuples — see Facts below |
| `supersedes` | Ren extension | **document-level**: array of document `id`s this document wholly replaces (used by Task 08's substantive updates) — see Forward-only consolidation edges below |
| `format` | Ren extension, `capture` only | `plain` \| `markdown` \| `html` — the raw format a Bronze capture was submitted in |

Document loading/validation (Task 02) must implement this table, the
type-specific fields below, and the default-by-type importance table as the
schema — not any illustrative example block that predates it.

### Facts

Each entry in a document's `facts:` array is a standalone structured claim:

```yaml
facts:
  - uid: 01K7Z3F8G2H9J...           # ULID, assigned once at extraction — the fact's identity
    subject: 01K...                  # entity ref (ULID) or a literal string
    predicate: uses
    object: "WhisperKit"              # or object_literal for a plain value; one of the two required
    confidence: 0.9                   # stated
    valid_from: 2026-08-30             # world time (see Two time axes)
    source_excerpt: "I need to benchmark WhisperKit"   # optional — a short verbatim quote from the
                                       # body supporting this fact; Task 04 hashes it (derived-only)
                                       # for drift detection, falling back to hashing the whole body
                                       # when omitted
    reviewed: true                    # optional — set by Task 11 once consolidation has considered
                                       # this fact, cosmetic bookkeeping on the fact's own (newly
                                       # written) document, not a mutation of an older document
    supersedes:   [f_01K5...]         # fact-level: arrays of *other facts'* uids, always written
    contradicts:  [f_01K4...]         # on the newer fact — see Forward-only consolidation edges
    duplicates:   [01K3...]
    reinforces:   [f_01K2...]
```

`uid` is the fact's own identity, assigned once and never re-validated
against anything else — Task 02 checks it's a syntactically valid ULID and
unique across the data dir, nothing more. The four edge fields *nested
inside a fact entry* reference other facts' `uid`s and are a different thing
from the document-level `supersedes` field above, which references a
document `id` — see Forward-only consolidation edges for which one to use
when.

The shared `status` field (generic, all types) is a free-text lifecycle tag
defaulting to `active`; `archived` is the one value every task must
recognize (a document a human or agent has retired without deleting it).
Types that need a richer status use a type-specific field instead of
overloading this one — see `experiment_status` and `conflict.status` below.

### Type-specific fields

Beyond the shared table above, these `type`s carry additional fields,
validated only when that `type` is set:

- **`conflict`**: `fact_uids` (array of ≥2 fact `uid`s in contention),
  `status: open | resolved | dismissed | accepted-both`, `detected_by`
  (`generated`-shaped), `resolution` (free text, once resolved),
  `resolved_at`.
- **`experiment`**: `experiment_status: proposed | active | concluded`,
  `friction` (a document `id` this experiment addresses — a plain reference
  field, scanned by Task 05's deterministic link extractor into the generic
  `links` table, not a new edge kind), `duration_days` or `end_date`,
  `outcome_result: positive | negative | neutral | inconclusive` (required
  once `experiment_status: concluded`).
- **`capture`**: `format` (see the shared table — listed there since every
  capture carries it, but it never appears on any other type).

### Importance: default by type

When `importance` is absent, the derived default (used for ranking, never
written back to the file) is:

| Type | Default |
|---|---|
| `decision`, `experiment` | 0.7 |
| `fact`, `summary` | 0.6 |
| `observation`, `friction`, `task` | 0.5 |
| `note`, `idea`, `question`, `insight` | 0.4 |
| everything else (including `capture`) | 0.5 |

Tier and `verified` still adjust the *effective* (derived) importance used
in ranking on top of this base — that layering is Task 12's job; this table
only fixes the base default so it's reproducible on rebuild.

## Stable IDs

ULID, stored as `id` in frontmatter, assigned at creation and never changed.
Filenames may be renamed freely by AI reorganization; links resolve through
`id`, not path, wherever both exist.

## Tiers

`tier: silver | gold` lives only on `knowledge/` documents. `inbox/` captures
are Bronze implicitly — they carry `type: capture` and **never** a `tier`
field; that keeps the immutability/byte-identical guarantee simple to state
and test through every later task, including consolidation.

Promotion Silver → Gold (Task 11) is never an in-place `tier` edit on an
existing document — that would mutate an older document, which the
forward-only rule below forbids. Promotion instead means writing a **new**
`tier: gold` document (a standalone `type: fact`, or a summary/decision) that
`reinforces`/`supersedes` the Silver material it promotes. The original
Silver document stays `silver`, untouched, and becomes an ancestor in the
provenance chain rather than being relabeled.

## Two time axes

Defined once here so no later task collapses them:

- **World time** — `valid_from` / `valid_to`: when a fact or document's
  claim is true *of the world*. Supersession sets the superseded fact's
  `valid_to` to the superseding fact's `valid_from` (or an explicit
  `effective_at` override).
- **System time** — `created_at` / `asserted_at`: when Ren *learned* the
  fact. Never conflate the two: a document can assert today that something
  was true last month.

Similarly, `verified` (human trust level: has a person looked at this?) and
`confidence` (extractor/author certainty at the moment of writing) are
independent axes. A document can be `verified: human-reviewed` with
`confidence: 0.6`, or `verified: unverified` with `confidence: 0.95`. Don't
let one stand in for the other anywhere in the schema or UI.

## Forward-only consolidation edges

Edges are always written on the **newer** record, at whichever of two
levels the change actually happened:

- **Fact-level** (Task 11's normal case): `supersedes`/`contradicts`/
  `duplicates`/`reinforces`, nested inside a `facts[]` entry (see Facts
  above), referencing other facts' `uid`s. Used when one claim replaces,
  conflicts with, duplicates, or reinforces another.
- **Document-level** (Task 08's substantive `update_document`): a top-level
  `supersedes` field on the document, referencing the document `id`(s) it
  wholly replaces. Used when an entire document is superseded, not just one
  of its facts.

Either way, an older record is never mutated to point forward. `valid_to`
and `superseded_by` are computed by reverse index at read time (in SQLite,
from file-resident forward edges — fact-level reverse-indexed in `fact_edges`,
document-level in `document_edges`, both derived-only). This keeps every
consolidation act a single-file write, keeps rebuild order-independent, and
means an old document is never touched by a later inference. Unresolved
contradictions become a standalone `type: conflict` document (see Type-
specific fields above) rather than being silently resolved in either
direction; a conflict's `status: dismissed` is itself a decision and
therefore file-resident — it must never be silently re-raised by a later
rebuild or reindex.

## Data model

Keep the ontology small. Core types:

```
note, project, concept, decision, observation, friction, experiment,
question, person, resource, summary, capture, fact, conflict,
idea, task, insight
```

`capture` is Bronze-only (`inbox/`). `fact` is for a promoted/merged Gold
fact with no natural host document — most extracted facts stay inline in
their source document's `facts:` array; only consolidated/promoted facts get
their own file, avoiding a file-per-triple explosion. `conflict` is one
document per unresolved contradiction detected during consolidation — see
Type-specific fields above for its schema. `idea`, `task`, and `insight` are
here specifically because Task 10's classifier produces them — see that
task's classification-label → `type` mapping table (a "possible action" is
stored as `task` with the shared `status: active` field marking it not yet
started; a "project update" is stored as `note` tagged `project-update`
rather than adding a fourteenth type for it).

Do not attempt a universal ontology upfront. Allow future extensibility.
Prefer conventions over a large schema system.

## Derived-state contract

**File-resident (must never live only in SQLite):** fact tuples and their
`uid`s, entity-linking decisions, stated `confidence`/`verified`,
supersede/contradict/duplicate/reinforce edges, conflict records including
dismissals, human verification, `pinned`, `importance`, `tier`.

**Derived-only (must never be written back into frontmatter):** content
hashes, chunk boundaries, chunk/document vectors, FTS rows, computed
`valid_to`/`superseded_by` (reverse-indexed from forward edges), effective
(reinforcement-adjusted) confidence, backlink/graph projections, hot-context
cache entries, index metadata/version numbers.

A content hash is computed *over* a file's bytes — it can never be stored
inside the frontmatter it would need to hash, which is why it's
derived-only.

## Version tracking

`runtime/` tracks `schema_version`, `parser_version`, `chunker_version`,
`embedding_model` + `embedding_version`, and `ranking_version` in an index
metadata table. Bumping any of these forces the incremental indexer to
reprocess exactly the rows it affects (see Task 04) — a version bump is
never a manual full-reindex requirement, though `ren reindex --full` must
always also work.

**Determinism is scoped, not absolute**: metadata, chunks, facts, graph
edges, and FTS content must rebuild bit-identically across repeated runs on
unchanged input, **except** a short, explicit list of operational
timestamps/counters that necessarily change on every run: `indexed_at`,
`last_full_reindex_at`, `index_revision`, `embedded_at`, and hot-context's
`built_at`/`expires_at`/`hit_count`/`last_accessed_at` (Task 13's cache is
exempt from determinism entirely — see that task). Every other column,
including file-projected
timestamps like `created_at`/`updated_at`/`asserted_at`, must match exactly
across rebuilds — don't widen the exclusion beyond this named list; a
silently-nondeterministic canonical projection is exactly the kind of bug
this contract exists to catch. Floating-point embedding vectors are the one
content column exempted from exact matching — compare those by matching
`embedding_version` plus a cosine-similarity tolerance, not byte equality.

## Hybrid retrieval

Ranking combines semantic similarity, FTS/BM25 relevance, importance,
recency, and relationship proximity into one score, with **configurable
weights** (config, not code) and **renormalization** when a signal is
unavailable (e.g. no embedding backend configured — semantic drops out and
the remaining weights renormalize over what's left; results are marked
`degraded` with a reason, never errored). Full ranking design lives in
Task 12; the config shape and default weights are specified there.

**Semantic search is on by default**, using a local embedding backend
(default: in-process ONNX, no external service — see Task 09), degrading
gracefully to FTS + importance + recency + proximity when unavailable. This
is a deliberate departure from treating embeddings as optional-and-off: Ren
targets an always-on Mac mini where local embeddings are cheap, and a
completely local non-vector mode still works as the fallback path, not the
default.

**`ai.provider` (the LLM used for classification/organization/consolidation
prose) and the embedding backend are two independent config axes.**
`ai.provider` defaults to `"none"` — the app still builds, runs, and passes
its full test suite with zero AI configured and zero network calls. The
embedding backend is separate and on by default. Don't let one default read
as contradicting the other; they answer different questions ("can an LLM
help me?" vs. "can I compute a vector locally?").

## MCP transport

stdio by default (the standard local MCP client pattern — Claude Code /
Claude Desktop config), documented as the only supported mode through
Task 08. An HTTP-over-Tailscale transport is optional, deferred to Task 16,
for agents that aren't local subprocesses.

MCP tools return raw markdown/frontmatter/JSON — never rendered HTML. This is
a hard requirement, not just a preference: agents must never be asked to
scrape the website.

## MCP permissions

Separate capability levels, checked before every write tool executes:

```
knowledge.read
knowledge.search
capture.create
knowledge.update
knowledge.admin
```

Read-only (`knowledge.read` + `knowledge.search`) is the out-of-the-box
default. Not every agent connection gets write/delete privileges.
Permanent deletion is privileged or avoided entirely — no delete tool ships
in the initial task set.

## AI provider

Default provider: `"none"`. The interface (Task 09) is shaped to cover
Anthropic, OpenAI, Ollama, LM Studio, llama.cpp, and MLX conceptually, but
only **Anthropic** and **Ollama** adapters get concrete implementations
initially — the rest stay interface-shaped stubs until something needs them.
This keeps Task 09 scoped to "prove the abstraction works end-to-end," not
"implement six providers."

## Embedding backend

Separate, pluggable interface (`EmbeddingBackend`: an `id` used as
`embedding_version`, a `dims` count, an `embed(texts) -> vectors` call) so
alternative representations (a different local model, a hosted API, future
compression) can be added later without touching the retrieval or indexing
code that consumes vectors. Default backend: **in-process ONNX**
(transformers.js, e.g. `Xenova/all-MiniLM-L6-v2`, 384 dimensions) — no
external service required, no compression in v1. Model weights cache under
`$REN_DATA_DIR/models/`, never under `runtime/`.

## Agent writes

AI agents modify the knowledge model, not application code. For structured
knowledge:

- preserve provenance
- never silently delete data
- prefer append/update/supersede/archive semantics — a changed fact
  supersedes the old one (see Forward-only consolidation edges), it doesn't
  overwrite it
- validate frontmatter/schema before writing
- use atomic writes (temp file + rename)
- avoid corrupting the knowledge directory
- maintain a change log where practical

A write lock (advisory, per-file or per-directory — single Mac mini, single
process is the target, not a distributed system) protects against
simultaneous writers.

**`runtime/changelog.jsonl` (Task 08) is a best-effort operational log, not
canonical audit history.** It lives under `runtime/` deliberately — it's
allowed to be lost on a `rm -rf runtime/`, because it isn't the thing that
makes history recoverable. What *does* make history recoverable is the
forward-only supersede model itself: an older document is never mutated, so
every past state is still sitting in `knowledge/` as an ancestor in a
supersede chain, reconstructable via Task 05's provenance/graph queries. The
changelog is a convenience for a human skimming recent activity, not the
source of truth for "what changed" — don't design anything to depend on it
surviving.

## Local-first & privacy

Prefer local files, local SQLite, local search, local embeddings, local
model support, Tailscale networking. Cloud services are optional
integrations, never foundational dependencies. Ren may contain extremely
private personal information, therefore: no telemetry by default, no
analytics by default, no silent cloud sync, no third-party tracking
scripts, no sending content to an AI provider without explicit
configuration, clearly identify when external inference is enabled, keep
secrets outside source files (env vars only), provide a strong
`.gitignore`, recommend storing user data outside the application
repository, never require public internet exposure, never open inbound
router ports.

## Design direction (UI)

Minimal, calm, modern, information-dense without feeling cluttered, fast,
keyboard-friendly, dark-mode friendly, suitable for long-form reading. Avoid
visual gimmicks. Not a generic SaaS dashboard, not a blog — a private
knowledge environment.

## What not to build

Not a collaborative Notion competitor, not a block editor, not a proprietary
rich-text format, not a cloud-first backend, no mandatory accounts, no
complicated permissions admin UI, no heavy distributed architecture, no
*required* vector database (the vector store lives inside the existing
SQLite runtime file), no microservices, no Kubernetes, no public SaaS
platform — unless a later requirement explicitly justifies one of these.
Task 16 checks the finished system against this list explicitly.

## Success criterion

A successful system makes this workflow effortless:

```
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

## Explicitly open / not decided here

- **License**: not chosen. Task 16 must surface this to the user rather than
  default silently — MIT is the common choice for a project like this, but
  it's a real decision (patent grant, copyleft, etc.), not a convention.
- **Deployment hostname / Tailscale Serve specifics**: left to Task 16,
  depends on the user's actual tailnet.
