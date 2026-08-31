# Task 11 — Knowledge consolidation: dedupe, contradiction, supersede, reinforce

**Read first:** `00-conventions.md` (Facts, forward-only consolidation
edges, tiers, the derived-state contract — dismissals must be file-resident).

**Depends on:** Task 10 (Silver documents with extracted `facts:` exist),
Task 05 (knowledge graph — candidate lookup and edge storage), Task 04
(FTS for candidate generation), Task 09 (embedding backend and `ai.provider`
— vector-similarity candidates and semantic contradiction judgment both use
this), Task 08 (write path — this task's supersede/new-document writes
reuse it).

## Objective

For each incoming fact, decide what it means for existing knowledge, and
write that decision into files — never only into SQLite. This is the pass
that promotes Silver toward Gold: extracting structured facts (already done
by Task 10), detecting duplicates and contradictions, reinforcing matching
knowledge, adding genuinely new knowledge, superseding changed facts, and
flagging ambiguous conflicts rather than silently resolving them. Bronze is
never touched by this task, directly or indirectly.

## Scope

In scope:
- `src/lib/consolidate/candidates.ts`: for an incoming fact (subject,
  predicate, object), find existing candidate facts to compare against —
  FTS match on subject/predicate text, vector similarity on the surrounding
  chunk (Task 04/09's index; degrade to FTS-only if the embedding backend is
  unavailable per `00-conventions.md`'s degrade contract), and exact entity
  match via Task 05's graph. Keep this a bounded, inspectable candidate set
  (e.g. top-K), not an all-pairs comparison.
- `src/lib/consolidate/decide.ts`: for each incoming fact against its
  candidates, decide one of — all four write into the **nested** edge arrays
  inside the *new* fact's own `facts[]` entry (per `00-conventions.md`'s
  Facts schema: `facts[].supersedes`/`contradicts`/`duplicates`/`reinforces`,
  each an array of *other facts'* `uid`s) — never the document-level
  `supersedes` field, which is Task 08's whole-document primitive, not this
  task's:
  - **Reinforce** — same claim, same subject/predicate/object (or
    equivalent phrasing): add the candidate's `uid` to the new fact's
    `reinforces` array; do not create a redundant fact document. The
    *derived* `effective_confidence` (computed in SQLite, not stored in
    frontmatter) rises with reinforcement count — the file-resident record
    is the edge, not the number.
  - **Add** — genuinely new, no real candidate: leave the fact where
    Task 10 wrote it, and mark it `reviewed: true` on the same (newly
    written) document — this is cosmetic bookkeeping on a document this
    task's own run is creating/touching for the first time, not a mutation
    of an older document, so it doesn't conflict with the forward-only rule.
    Without this marker the fact would look "unconsolidated" forever and be
    re-considered on every run — see the trigger bullet and the idempotence
    acceptance criterion below. Promotion to a standalone Gold document (see
    below) can also apply to an "Add" outcome once its criteria are met.
  - **Supersede** — same subject/predicate, changed object/value: add the
    candidate's `uid` to the new fact's `supersedes` array (forward-only,
    per `00-conventions.md` — the old fact's document is never edited).
    Task 05's derived `valid_to`/`superseded_by` pick this up automatically
    on reindex.
  - **Flag (ambiguous conflict)** — plausible contradiction that isn't a
    clear supersession (e.g. two claims that could both be true depending
    on context, or contradictory claims with comparable confidence and no
    clear temporal ordering): write a standalone `type: conflict` document
    per `00-conventions.md`'s Type-specific fields (`fact_uids`, `status:
    open`, `detected_by`). **Never silently resolve this case** by picking a
    side — an open conflict is a human work queue item, not a bug to be
    quietly fixed. A human (or an explicit later action) can mark a
    conflict `resolved`/`dismissed`/`accepted-both`; a `dismissed` status is
    itself file-resident and must not be re-raised by a later consolidation
    pass or a full reindex. Mark the newer fact `reviewed: true` regardless
    of outcome (including Flag) so it isn't re-considered on the next run —
    the conflict document, not the reviewed marker, is what tracks whether
    the underlying disagreement is resolved.
- Promotion (Silver → Gold): **never** an in-place `tier` edit on an
  existing document (that would mutate an older document — forbidden, see
  `00-conventions.md`'s Tiers section). When a fact's reinforcement count
  passes a configurable threshold, or a human/agent explicitly requests
  promotion, this task writes a **new** standalone `tier: gold`,
  `type: fact` document that `reinforces` the promoted fact's `uid` — the
  Silver material that earned the promotion stays exactly as it was, now an
  ancestor in the provenance chain rather than a relabeled file.
- A manual trigger (`ren consolidate`, or an MCP tool — match Task 10's
  choice) run over unconsolidated facts — the eligibility query is "facts
  with no `reviewed` marker set" (not "no outgoing edge", since an Add
  outcome legitimately has no edge but must still not be re-processed) —
  rather than a fully automatic pipeline, for the same reasons Task 10 stays
  manually triggered.
- Tests: a fixture where the same claim appears in three independently
  captured documents produces two `reinforces` edges (not three duplicate
  fact documents) and a measurably higher derived `effective_confidence`
  than a single-mention fact; a fixture where a later document changes a
  fact's value produces a `supersedes` edge (nested inside that fact's
  `facts[]` entry, distinct from Task 08's document-level `supersedes`) and
  leaves the earlier document's file byte-identical; a fixture with two
  genuinely conflicting, comparably-confident claims produces a
  `type: conflict` document with `status: open`, not a silently-picked
  winner; a `conflict` document marked `dismissed` in a fixture is not
  re-flagged by a second consolidation run or by `ren reindex`; a fixture
  fact with no real candidate ("Add") is marked `reviewed: true` and is
  correctly excluded from the eligibility query on a second run; running
  consolidation twice over unchanged input produces zero new documents and
  zero new edges (idempotent, including for the Add case); with
  `ai.provider: "none"`, exact-duplicate detection (normalized-text/hash
  match) and clear temporal-overlap contradiction detection still run and
  still produce file-resident results — only prose-level semantic
  contradiction judgment is skipped, and that's reported, not silently
  dropped; a promoted fact produces a *new* Gold document while its Silver
  source document's `tier` remains `silver` and its file is untouched.

Out of scope: fact extraction itself (Task 10 — this task consumes what's
already in `facts:`), ranking/ retrieval (Task 12 — consolidation improves
what retrieval finds, it doesn't rank results), the hot-context cache
(Task 13), recurring-friction detection across many Kaizen entries
(Task 14 — a related but higher-level pattern-detection pass; it should
read consolidated facts and the graph rather than duplicate this task's
comparison logic).

## Deliverables

- `src/lib/consolidate/candidates.ts`.
- `src/lib/consolidate/decide.ts`.
- `src/lib/consolidate/run.ts` (trigger).
- Tests.

## Acceptance criteria

- [ ] Three independently captured documents making the same claim produce
      exactly two `reinforces` edges and a derived `effective_confidence`
      above any single-mention baseline — with zero duplicate fact
      documents created.
- [ ] A later document that changes a fact's value produces a fact-level
      `supersedes` edge on the new document's `facts[]` entry; the earlier
      document's file is byte-identical before and after.
- [ ] A genuinely ambiguous contradiction produces an open `type: conflict`
      document naming both facts via `fact_uids` — never a silently-resolved
      outcome.
- [ ] A `dismissed` conflict is never re-raised by a later consolidation run
      or by `rm -rf runtime/ && ren reindex`.
- [ ] A fact with no real candidate is marked `reviewed: true` and excluded
      from re-consideration; consolidation is idempotent overall: running it
      twice over unchanged input creates no new documents or edges.
- [ ] With `ai.provider: "none"`, deterministic dedupe/contradiction paths
      (exact-match, temporal-overlap) still function; the pipeline does not
      error, and the skipped semantic-judgment cases are reported, not
      silently dropped.
- [ ] A promoted fact produces a new Gold `type: fact` document; the
      originating Silver document's `tier` field and file content are
      unchanged.

## Notes / assumptions

- Every consolidation outcome must be a file write, never only a SQLite row
  or in-memory decision — this is the task where the "safely deletable and
  rebuildable" invariant is most at risk of being quietly violated, because
  it's tempting to store dedupe/conflict state only in the derived layer.
  Don't.
- Reinforcement never mutates the reinforced (older) document — only the
  newer document gets a new edge pointing at it. This keeps consolidation
  order-independent and every write a single-file operation, consistent
  with Task 05/08's forward-only model. The same rule is why promotion
  writes a new document instead of flipping an existing `tier` field.
