# Build Ren — an agent-native personal knowledge system

Build an open-source project called **Ren**, named after the ancient Egyptian concept/personification of perception, knowledge, and understanding.

Ren is a **private, local-first, agent-native personal knowledge system**. It is not meant to be a Notion clone. The core idea is that users should be able to dump raw thoughts, notes, observations, project ideas, decisions, and Kaizen-style brain dumps into the system, while AI agents continuously organize, connect, synthesize, and maintain the resulting knowledge base.

The application itself should be open source and live in GitHub. User data must remain private and must never be committed to the application repository.

## Core product philosophy

Ren should treat the user's knowledge as portable data, not as application-owned state.

The system should have two primary interfaces:

1. **Astro-based website** for humans.
2. **MCP server** for AI agents.

Both should operate on the same underlying knowledge store.

The website is a human-readable projection of the knowledge.

The MCP server is the machine-readable interface to that same knowledge.

Do not make AI agents scrape or parse the rendered website to retrieve knowledge.

## Technology choices

Use:

- Astro
- TypeScript
- Markdown
- YAML frontmatter
- OKF, Google's Open Knowledge Format, as the preferred structured knowledge representation
- SQLite for derived metadata/search state where appropriate
- BM25/full-text search
- Optional vector embeddings for semantic retrieval
- MCP for agent access
- Git for application source control
- Tailscale for private deployment/access

Prefer simple, local-first technologies over cloud infrastructure.

Avoid unnecessary dependencies, databases, services, and infrastructure.

Ren should be able to run comfortably on an always-on Mac mini.

## Repository and private-data separation

The GitHub repository should contain only application code, schemas, documentation, templates, tests, and example/demo content.

Private user data should live outside the repository whenever possible.

Support something like:

```text
~/code/ren/
    src/
    public/
    packages/
    package.json
    astro.config.*
    ren.config.*

~/data/ren/
    inbox/
    knowledge/
    runtime/
```

Use a configurable environment variable such as:

```text
REN_DATA_DIR=~/data/ren
```

The application may optionally default to a local ignored directory for development, but production usage should encourage external data storage.

Never assume `.gitignore` is sufficient protection for private knowledge.

## OKF

Use OKF as the primary format for structured and curated knowledge.

Knowledge should remain:

- human-readable
- Markdown-based
- Git-friendly if users choose to version their private store separately
- portable
- agent-readable
- minimally proprietary

Do not introduce a Ren-specific proprietary document format unless absolutely necessary.

Ren may extend OKF metadata with application-specific fields where useful.

Example:

```md
---
type: decision
title: Use WhisperKit for Lumi
tags:
  - lumi
  - transcription
created_at: 2026-08-29
updated_at: 2026-08-29
created_by: ren-agent
verification: unverified
sources:
  - inbox/2026-08-29-001.md
---

# Use WhisperKit for Lumi

WhisperKit is currently preferred because...
```

Support provenance wherever practical.

The user should be able to tell:

- what was directly written by them
- what was inferred by AI
- which raw input produced a derived conclusion
- whether something has been human verified
- when information was last updated
- whether a piece of knowledge may be stale

## Capture should be frictionless

The most important interaction is capturing thoughts.

Users should be able to submit:

- plain text
- Markdown
- HTML
- brain dumps
- observations
- project notes
- ideas
- decisions
- questions
- potentially transcripts later

A capture should require essentially no organization from the user.

Do not force the user to choose:

- folders
- tags
- page names
- taxonomies
- categories

before saving a thought.

The AI layer should handle organization after capture.

## Kaizen brain dumps

Kaizen-style brain dumps are a first-class feature.

A user may submit something messy such as:

```text
I keep checking Slack too much.

Lumi transcription still feels slow.

Maybe I should batch meetings in the afternoon.

I need to benchmark WhisperKit.

The kitchen drawer keeps becoming a mess.
```

Ren should preserve the original input and then derive useful structured knowledge.

The AI may classify content into concepts such as:

- observation
- friction
- recurring friction
- idea
- potential action
- task
- decision
- experiment
- question
- insight
- project update

Do not turn every observation into a task.

Distinguish between:

```text
"I noticed..."
→ observation

"I should..."
→ possible action

"This keeps happening..."
→ recurring friction

"What if I tried..."
→ experiment

"I decided..."
→ decision
```

A key goal of Ren is to identify repeated patterns across time.

For example, multiple independent notes about context switching should eventually be recognized as one recurring friction pattern.

Ren should be able to generate something like:

```text
Recurring friction:
Context restoration after project switching

Mentions:
4 times in 25 days

Possible intervention:
Create automatic project checkpoints before switching.

Suggested experiment:
Try for 14 days.

Possible success metric:
Time required to resume work.
```

## Preserve raw input

Raw user input should be immutable or treated as append-only.

Never rewrite the user's original thought as part of normal AI organization.

Derived knowledge can change over time.

Think in terms of:

```text
RAW THOUGHT
    ↓
INTERPRETATION
    ↓
STRUCTURED KNOWLEDGE
    ↓
SYNTHESIS
    ↓
ACTION / DECISION / EXPERIMENT
```

Provenance between these stages should be retained where practical.

## Medallion architecture

Medallion architecture is a useful mental model but is **not a primary product requirement**.

Do not overengineer Ren around Bronze/Silver/Gold terminology.

If it naturally improves the internal design, the approximate mapping is:

```text
Bronze
Raw immutable captures

Silver
Structured AI-derived knowledge

Gold
Curated conclusions, decisions, summaries, active experiments, and high-value knowledge
```

However:

- OKF matters more than medallion terminology.
- Simple file organization matters more than forcing data tiers.
- Do not expose Bronze/Silver/Gold concepts prominently in the UI unless there is a clear UX reason.
- It is perfectly acceptable to implement the same conceptual separation without using those names.

## Knowledge relationships

Ren should automatically support relationships between knowledge.

Examples:

- project ↔ decision
- project ↔ person
- project ↔ technology
- observation ↔ recurring friction
- friction ↔ experiment
- experiment ↔ outcome
- concept ↔ related concept
- note ↔ source
- derived knowledge ↔ original capture

Backlinks and related-content discovery should be automatic where practical.

Avoid requiring users to manually maintain a graph.

## Astro website

Build the human interface with Astro.

Use Astro because most of the application is content-oriented and should be rendered as lightweight HTML.

Avoid shipping unnecessary client-side JavaScript.

Use interactive islands only where useful, such as:

- global search
- command palette
- AI query interface
- graph visualization
- capture box
- filtering
- dynamic dashboards

Most knowledge pages should work as simple server-rendered or statically rendered HTML.

The UI should feel more like a calm personal operating system than a blog.

Possible main navigation:

```text
Home
Inbox
Knowledge
Projects
Kaizen
Decisions
Experiments
Recent
Search
```

## Suggested home/dashboard

Create a useful dashboard that may include:

```text
Recently updated knowledge

Active projects

Open questions

Recent decisions

Recurring friction

Active experiments

Unprocessed captures

Recently resolved issues
```

Do not overload the interface.

## Kaizen view

Create a dedicated Kaizen view.

Possible layout:

```text
KAIZEN

Capture
[ Dump anything on your mind... ]

Recurring Friction
- Context switching
- Meeting fragmentation
- Repeated administrative work

Active Experiments
- Project checkpoint before switching
- Check Slack at fixed intervals

Recent Improvements
- Morning planning workflow
- Grocery planning automation
```

Support frequency or recurrence indicators when useful.

## Search architecture

Knowledge retrieval must be fast.

Build a local search layer.

Preferred approach:

```text
OKF / Markdown source files
        ↓
indexing
        ↓
SQLite
├── metadata
├── full-text/BM25
└── optional embeddings
```

Search indexes are derived state.

The source Markdown/OKF files are canonical.

It should always be possible to delete the SQLite/search state and rebuild it from source files.

## Hybrid search

Provide hybrid retrieval where practical:

- keyword/BM25
- semantic/vector retrieval
- metadata filters
- recency
- relationships/backlinks

Then rerank or combine results.

Do not require embeddings for the system to function.

A completely local, non-vector mode should work.

## MCP server

Build an MCP server as a first-class part of Ren.

The MCP server should expose structured access to knowledge rather than rendered HTML.

Initial tools should include concepts similar to:

```text
search_knowledge(query, limit, filters)

get_document(id_or_path)

get_context(query, max_tokens)

get_recent_notes()

get_project(name)

get_decisions(project?)

get_related(id_or_topic)

capture(text, format?)

kaizen_dump(text)

create_note(...)

record_decision(...)

create_experiment(...)

update_document(...)

review_kaizen()
```

Names may differ if better MCP conventions exist.

The most important tool is something conceptually similar to:

```text
get_context(query, max_tokens)
```

It should gather the most relevant information for an AI agent without requiring the agent to manually inspect many files.

Internally it may:

1. run keyword search
2. run semantic search if enabled
3. consider metadata
4. follow relevant relationships
5. consider recent captures
6. rerank
7. return a compact context bundle

The result should be optimized for agent context windows.

## MCP permissions

Design access control with separate capability levels.

Conceptually:

```text
knowledge.read
knowledge.search
capture.create
knowledge.update
knowledge.admin
```

Not every agent should receive write/delete privileges.

Where practical, expose read-only operation as the safest default.

Permanent deletion should be privileged or avoided entirely.

## Authentication and networking

The initial deployment target is:

```text
Always-on Mac mini
        ↓
Tailscale
        ↓
user devices + trusted agents
```

Ren should not require public internet exposure.

Do not require Auth0, Clerk, Firebase Auth, Supabase Auth, or another cloud identity system for the default local deployment.

For the initial version, trust the Tailscale network boundary.

Design the architecture so stronger authentication can be added later if Ren is exposed beyond the tailnet.

Do not open inbound router ports.

Document recommended deployment using Tailscale Serve where appropriate.

## Agent writes

AI agents should preferably modify the knowledge model rather than application code.

For structured knowledge:

- preserve provenance
- avoid silently deleting data
- prefer append/update/archive semantics
- validate frontmatter/schema
- use atomic writes
- avoid corrupting the knowledge directory
- maintain a change log where practical

Consider filesystem locking or equivalent protection against simultaneous writers.

## AI provider independence

Ren should not depend on one AI provider.

Design clear interfaces for AI capabilities.

Potential providers may include:

- OpenAI
- Anthropic
- local Ollama models
- LM Studio
- llama.cpp
- MLX

The core application should still function as a knowledge browser/search system without an AI provider configured.

AI enrichment should be modular.

## Local-first

Local-first is a core principle.

Prefer:

```text
local files
local SQLite
local search
local embeddings when desired
local model support
Tailscale networking
```

Cloud services should be optional integrations rather than foundational dependencies.

## Privacy

Assume Ren may contain extremely private personal information.

Therefore:

- no telemetry by default
- no analytics by default
- no silent cloud sync
- no third-party tracking scripts
- no sending content to an AI provider without explicit configuration
- clearly identify when external inference is enabled
- keep secrets outside source files
- provide a strong `.gitignore`
- recommend storing user data outside the application repository

## Configuration

Create a straightforward configuration system.

Possible example:

```ts
export default {
  dataDir: process.env.REN_DATA_DIR,
  search: {
    semantic: false
  },
  ai: {
    provider: "none"
  },
  network: {
    trustProxy: true
  }
}
```

Prefer environment variables for secrets.

## CLI

Add a minimal CLI if it helps the architecture.

Useful commands could include:

```bash
ren init

ren dev

ren capture "..."

ren reindex

ren doctor

ren serve

ren export

ren validate
```

Do not let CLI work delay the core website/MCP implementation.

## File watching and indexing

Watch the data directory for changes.

When knowledge files change:

- validate them
- update metadata
- update the search index
- update relationships where applicable

Avoid rebuilding the entire system unnecessarily.

Also support a full deterministic rebuild:

```bash
ren reindex
```

## Data model

Keep the initial ontology small.

Potential core types:

```text
note
project
concept
decision
observation
friction
experiment
question
person
resource
summary
```

Do not attempt to design a universal ontology upfront.

Allow future extensibility.

Prefer conventions over a large schema system.

## IDs

Documents should have stable IDs independent of their filename where practical.

This helps preserve links if the AI later reorganizes or renames files.

A document may contain:

```yaml
id: 01K...
```

Use a suitable sortable identifier such as UUIDv7 or ULID if useful.

## Human verification

Support AI-derived knowledge being marked with trust/provenance state.

Possible states:

```text
human-authored
ai-derived
human-verified
stale
```

Do not overcomplicate this in v1, but design metadata so it can evolve.

## Editing model

The user should be able to manually edit Markdown files without breaking Ren.

Ren must not require all mutations to occur through the UI.

External editors such as:

- VS Code
- Obsidian
- Vim
- AI coding agents

should be able to safely modify the knowledge store.

The indexing layer should notice and incorporate changes.

## Open-source experience

The public repository should be polished enough that another person can clone and run Ren.

Provide:

- README
- architecture overview
- development setup
- example knowledge directory
- example OKF documents
- security/privacy notes
- Tailscale deployment instructions
- MCP configuration example
- contribution guide
- license
- screenshots when UI exists

The project should have an obvious initial workflow:

```bash
git clone ...
cd ren
npm install
npm run setup
npm run dev
```

Then:

```text
Open Ren
→ paste a brain dump
→ see it saved
→ see structured knowledge appear
→ search it
→ query it through MCP
```

## Design direction

The UI should be:

- minimal
- calm
- modern
- information-dense without feeling cluttered
- fast
- keyboard-friendly
- dark-mode friendly
- suitable for long-form reading

Avoid visual gimmicks.

Do not design it as a generic SaaS dashboard.

Ren should feel like a private knowledge environment.

## Important architectural principle

Maintain a strict distinction between:

```text
Canonical knowledge
= Markdown / OKF

Derived runtime state
= SQLite indexes, embeddings, caches, relationships that can be rebuilt

Application
= Astro + MCP + indexing + AI orchestration
```

The application must never become the only place where the user's knowledge can be interpreted.

If Ren disappears, the user should still have a directory full of useful Markdown files.

## Implementation priorities

Build incrementally.

### Phase 1 — foundation

Implement:

- Astro project
- external configurable data directory
- Markdown/OKF document loading
- schema/frontmatter validation
- basic navigation
- knowledge page rendering
- local full-text search
- capture endpoint/UI
- immutable raw captures
- indexing/reindexing

### Phase 2 — MCP

Implement:

- MCP server
- search
- document retrieval
- context retrieval
- capture
- read-only access mode

### Phase 3 — AI organization

Implement an abstraction for AI processing.

Allow raw captures to be transformed into:

- structured notes
- observations
- project updates
- decisions
- questions
- friction
- experiments

Preserve source provenance.

### Phase 4 — Kaizen

Implement:

- recurring-friction detection
- experiment tracking
- Kaizen dashboard
- weekly/review synthesis
- links between observations → friction → experiments → outcomes

### Phase 5 — enhanced retrieval

Add where useful:

- embeddings
- hybrid search
- reranking
- richer relationship traversal
- token-budgeted context assembly

### Phase 6 — hardening

Add:

- concurrency safety
- backups/export
- validation tooling
- audit/change history
- enhanced permissions
- tests
- deployment documentation

Do not start with every Phase 6 capability.

## Testing

Include tests for the highest-risk parts:

- knowledge file parsing
- schema validation
- search indexing
- reindex determinism
- stable IDs
- source/provenance preservation
- capture append behavior
- MCP read operations
- safe file writes
- malformed Markdown/frontmatter handling

Use sample/demo data rather than private fixtures.

## What not to build

Do not build:

- a collaborative Notion competitor
- a block editor
- a proprietary rich-text format
- a cloud-first backend
- mandatory accounts
- a complicated permissions admin UI
- a heavy distributed architecture
- a required vector database
- microservices
- Kubernetes
- a public SaaS platform

unless later requirements explicitly justify them.

## Main success criterion

A successful initial version should make this workflow feel effortless:

```text
User:
"Brain dump: I keep losing context when I switch between projects.
Maybe Ren should save a quick checkpoint whenever I switch."

        ↓

Ren preserves raw capture.

        ↓

AI derives:
Observation:
Frequent project switching causes context-restoration overhead.

Potential experiment:
Create a project checkpoint before switching.

Related:
Existing notes about attention switching and project workflow.

        ↓

User can later ask an AI agent:

"What problems have I repeatedly mentioned about switching projects?"

        ↓

MCP returns the relevant synthesized knowledge and supporting raw sources.
```

The defining characteristic of Ren is:

**Users capture thoughts. Ren turns those thoughts into evolving understanding without requiring users to manually organize their lives into folders, databases, and tags.**

Start by inspecting the repository if one already exists, then produce a concise implementation plan and begin building the smallest coherent end-to-end version. Prefer working software over elaborate abstractions. Make sensible decisions without repeatedly asking for clarification, and document any assumptions you make.
