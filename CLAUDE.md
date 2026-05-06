# CLAUDE.md — CISSP Study-Notes Wiki Schema

This repo is an LLM-maintained personal wiki for CISSP exam prep. The agent (Claude Code) reads
the raw sources, builds and maintains the wiki/, and serves queries against it.

## Roles

- **User**: curates raw sources, asks questions, directs analysis, reviews wiki edits.
- **Agent (you)**: extracts knowledge from `raw/`, writes/updates `wiki/`, maintains
  cross-references and the index/log, answers queries with citations.

The user does not edit `wiki/` by hand. You are responsible for every page in there.
You never modify `raw/`. (docling-extracted .md files in raw/ are part of the source layer.)

## Layout

- `raw/` — immutable sources. PDFs + their docling-extracted .md siblings.
- `wiki/` — LLM-owned markdown wiki (see structure below).
- `CLAUDE.md` — this file. The schema. Co-evolves with the wiki.

```
wiki/
  index.md             # catalog of all pages, organized by category
  log.md               # append-only timeline of ingests/queries/lint passes
  overview.md          # CISSP exam map (8 domains, weights, top-level structure)
  cheatsheet.md        # living weak-areas doc, written by you, revised by you
  acronyms.md          # alphabetized acronym glossary
  mnemonics.md         # mnemonics for the topics that benefit from them
  domains/             # one page per CISSP domain (8 total)
  concepts/            # cross-domain concepts (AAA, asymmetric crypto, BCP/DRP, etc.)
  standards/           # NIST CSF, ISO 27001, PCI-DSS, COBIT, etc.
  sources/             # one summary page per ingested raw source
  practice-questions/  # one page per domain, holding Q&A items for that domain
```

## Page conventions

Every wiki page begins with YAML frontmatter:

```yaml
---
title: <human-readable title>
type: domain | concept | standard | source | practice | overview | cheatsheet | glossary
domain: <1-8 or "cross">    # primary CISSP domain this page belongs to
tags: [tag1, tag2, ...]     # free-form
sources: [cissp-ultimate-guide-rb, cissp-exam-outline, ...]   # source slugs that informed this page
updated: YYYY-MM-DD
---
```

After the frontmatter, the page body uses these conventions:

- **File names**: kebab-case, ASCII only. Domain pages prefix with their number (`01-...`).
- **Cross-references**: relative markdown links, e.g. `[AAA](../concepts/aaa.md)`. Do not use
  Obsidian `[[wikilinks]]` — keep portability for plain markdown viewers.
- **Sources section**: each page ends with a `## Sources` section listing raw sources with
  section/page pointers when known (e.g. `cissp-ultimate-guide-rb §5.3`).
- **Domain pages**: include the official subtopic outline (from the Exam Outline) as the spine,
  and link out to concept/standard pages for depth. The domain page is a hub, not a dump.
- **Concept/standard pages**: definition → key facts → exam-relevant nuance → cross-links → sources.
- **Practice-question pages**: each Q is `### Q<n>: <stem>` with `**Answer:**` and `**Why:**` blocks
  underneath. Tag with the relevant subtopic and link to the concept page that explains it.

## index.md

Single source of truth for "what's in the wiki." Organized by category. Every page in `wiki/`
must appear in `index.md` with a one-line summary. Update `index.md` on every ingest and
whenever you create a new page. Format:

```markdown
## Domains
- [01 — Security and Risk Management](domains/01-security-and-risk-management.md) — governance, compliance, risk frameworks (16%)
- ...

## Concepts
- [AAA](concepts/aaa.md) — authentication, authorization, accounting
- ...
```

## log.md

Append-only. Every entry starts with `## [YYYY-MM-DD] <op> | <subject>` so it's grep-able:

```
## [2026-05-05] init | wiki bootstrapped, both source PDFs extracted
## [2026-05-06] ingest | cissp-ultimate-guide-rb ch.1 (security and risk management)
## [2026-05-06] query | "difference between qualitative and quantitative risk analysis"
## [2026-05-10] lint | found 3 orphan pages, 1 contradiction, fixed
```

Use `grep "^## \[" wiki/log.md | tail -10` to recover recent activity.

## Operations

### Ingest

When the user says "ingest `<source>`" or "process raw/X":

1. Read the relevant raw markdown (docling-extracted, so just `Read` it; very large files need
   to be read in sections with offset/limit).
2. Discuss key takeaways with the user before writing — confirm scope and what to emphasize.
3. Create or update `wiki/sources/<slug>.md` with: bibliographic info, scope, quality
   assessment, a structured summary, and a list of every wiki page this source touches.
4. For each domain/concept/standard the source covers:
   - If a page exists, update it. Mark new claims, flag contradictions with existing claims
     (do NOT silently overwrite — surface them to the user).
   - If no page exists and the topic warrants one, create it.
5. Update `wiki/index.md`: add new pages, refresh one-liners that changed.
6. Append a `## [date] ingest | <slug>` entry to `wiki/log.md`.
7. Bump `updated:` frontmatter on every page touched.

A single ingest typically touches 5–15 pages.

### Query

When the user asks a question:

1. Read `wiki/index.md` first to find candidate pages.
2. Read those pages. If the answer requires source-level detail, fall through to `raw/*.md`
   (cite section/page number).
3. Answer with inline citations (link to wiki pages and source slugs).
4. If the answer surfaces a connection or analysis worth keeping, ask whether to file it back
   into the wiki. If yes, write it and update index/log.

### Lint

When the user says "lint" or "health-check":

- Find orphan pages (no inbound links from index or other pages).
- Find concepts mentioned ≥3 times across the wiki but lacking their own page.
- Find contradictions: identical claims with different values across pages.
- Find stale `updated:` dates on pages whose source has been re-ingested.
- Suggest pages where a quick web search would close a data gap.

Report findings in a markdown table. Apply fixes only with the user's go-ahead.

## Active-study artifacts

### cheatsheet.md

Living weak-areas doc. Each entry is a `### <topic>` section with the minimum facts the user
needs to remember, plus a link to the full concept page. Reorder by frequency-of-confusion
when asked.

### acronyms.md

Alphabetized. Format: `**ABC** — Full Form. (Short gloss. [Concept page](concepts/...).)`.
Add new acronyms on every ingest when you encounter a term in a source.

### mnemonics.md

Curated. Each mnemonic: the trick + what it stands for + which domain/concept it applies to.
Only keep ones that actually help; don't invent low-quality ones.

### practice-questions/\<domain\>.md

One file per domain. Q items as `### Q<n>: <stem>` with `**Answer:**` and `**Why:**` blocks.
When a query produces a question-style insight worth quizzing on later, ask the user whether
to file it here.

## House rules

- **CISSP scope only.** Do not file material outside the 8 CISSP domains. If a source covers
  adjacent material, summarize it in the source page but do not create concept/standard pages.
- **Cite, don't paraphrase silently.** Every non-trivial claim links to at least one source.
- **Surface contradictions, never resolve them silently.** If two sources disagree, both
  positions appear on the page with attribution; the user decides.
- **Currency.** Note publication years on standard pages (e.g. NIST SP 800-53 Rev. 5).
- **Today's date** for `updated:` and `log.md` entries: use the date from the user's session
  or `CLAUDE.md` if provided.

## Raw source notes

Both PDFs have been extracted to markdown by docling (granite VLM pipeline):

- `raw/CISSP-Exam-Outline.md` — structural source-of-truth for domain/subtopic naming and
  exam weights. Contains inline base64 images; text content is at specific line ranges.
  Domain 1 starts at line 70, Domain 2 at 149, Domain 3 at 178, Domain 4 at 260,
  Domain 5 at 301, Domain 6 at 349, Domain 7 at 399, Domain 8 at 493.
- `raw/CISSP-Ultimate-Guide-RB.md` — comprehensive study guide. Very large; read in chunks
  using `offset` and `limit` parameters. Use `grep` on the raw file to locate section headers
  before reading full sections.

## Quick commands

```sh
grep "^## \[" wiki/log.md | tail -10        # recent activity
find wiki -name "*.md" | wc -l              # page count
rg -l "topic" wiki/                         # pages mentioning a topic
grep -n "^## " raw/CISSP-Ultimate-Guide-RB.md | head -50  # chapter list
```
