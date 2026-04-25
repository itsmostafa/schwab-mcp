# CLAUDE.md — Interview Prep Wiki Schema

## Purpose

This repo is a persistent, compounding knowledge base for interview prep at Panther Labs. The wiki is the compiled output — you (the LLM) write and maintain it; the human reads it and directs what to work on next. Raw source documents live in `raw/` and are never modified. The wiki in `wiki/` is the living artifact.

## Directory Layout

```
raw/                        # Source documents — read-only, never modify
  job-description.md        # Panther job posting
  additional-interview-info.md  # Text messages with Darwayne (internal tips)
  candidate-prep-guide-staff-ai-engineer-soc-agent-platform.pdf  # Official prep guide for the 4 main loop interviews

wiki/                       # LLM-maintained wiki — you own this
  index.md                  # Master catalog of all wiki pages (update on every change)
  log.md                    # Append-only activity log
  overview.md               # Big picture: company, role, full interview structure
  rounds/                   # One page per interview round
    live-coding.md          # Friday's 60-min CoderPad round (most urgent)
    ai-integration.md       # AI Integration interview
    systems-design.md       # Systems design round
    project-retro.md        # Project retrospective round
    culture.md              # Culture fit round
  concepts/                 # Deep-dive reference pages
    stackexchange-api.md    # StackExchange API v2.3 — GET /search reference + practice
    agentic-ai.md           # AI agents, RAG, embeddings, vector DBs, feedback loops
    soc-domain.md           # SOC workflows, alert triage, XDR, threat analysis
    systems-design-patterns.md  # Patterns for the systems design round
  panther/
    company.md              # Panther context, values, product, culture
    role.md                 # Role expectations, key differentiators, what they care about
```

## Key Context

- **Candidate**: Mostafa — strong Go + Python + security background, AI experience
- **Company**: Panther Labs — AI SOC platform, raised $140M, remote-first
- **Role**: AI engineer for SOC automation (alert triage, agentic AI, detection-as-code)
- **Referral**: Darwayne referred Mostafa; know each other through Jamie + AI work
- **Interview status**: Live coding round (Round 2) — COMPLETE
- **Next interview**: AI Integration (Round 3) — upcoming
- **Final gate**: CEO (founder, former CTO, very AI-focused)

## Interview Rounds (in order)

1. ~~Chat with EM/VP~~ — DONE (Justin, went well)
2. ~~Live coding~~ — DONE (CoderPad, StackExchange API)
3. **AI Integration interview** — NEXT
4. Systems design — "most revealing" per Darwayne
5. Project retrospective
6. Culture
7. CEO final gate

## Operations

### Ingest a new source
1. Read the source document in `raw/`
2. Discuss key takeaways with Mostafa
3. Write or update relevant wiki pages
4. Update `wiki/index.md` to reflect any new pages or major updates
5. Append an entry to `wiki/log.md` in format: `## [YYYY-MM-DD] ingest | Source Title`

### Answer a query
1. Read `wiki/index.md` to find relevant pages
2. Read those pages and synthesize an answer
3. If the answer is substantive enough to be useful in future sessions, file it back as a new wiki page or addition to an existing one
4. Append to `wiki/log.md`: `## [YYYY-MM-DD] query | Question summary`

### Lint the wiki
Check for: contradictions between pages, stale info, orphan pages, concepts mentioned but not having their own page, missing cross-references, data gaps worth filling.

## Conventions

- All wiki pages use `[[wikilink]]` style cross-references where relevant (Obsidian-compatible)
- Frontmatter is optional but encouraged for entity pages: `tags`, `type`, `last-updated`
- The index uses one line per page: `- [Title](path) — one-line hook`
- Log entries start with `## [YYYY-MM-DD]` so they're greppable
- Raw sources are immutable — never edit files in `raw/`
- Prioritize the AI Integration round as the next immediate focus
