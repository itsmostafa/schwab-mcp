# Task 14 — Kaizen: recurring friction, experiments, dedicated view

**Read first:** `00-conventions.md`.

**Depends on:** Task 10 (classified friction/observation/experiment
documents exist to detect patterns over), Task 05 (knowledge graph and
consolidation edges — recurrence is read from the graph, not
reclustered from scratch), Task 11 (consolidated/reinforced facts —
better inputs for detecting a real recurring pattern vs. noise), Task 06
(website shell to add the view to), Task 03 (capture — the Kaizen page
reuses its component), Task 08 (write path, for creating experiment
documents from the UI).

## Objective

Surface recurring patterns across independently-captured friction/observation
documents, and give Kaizen its own view. Target output shape, for a friction
detected 4 times in 25 days:

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

## Scope

In scope:
- `src/lib/kaizen/detect.ts`: find recurring friction by reading Task 05's
  graph and Task 11's consolidation edges — a friction that has accumulated
  `reinforces` edges across independently-captured documents *is* a
  recurring pattern; this task queries that structure (`getRelated`,
  `getFactHistory`, reinforcement counts) within a rolling window (e.g. last
  90 days) and surfaces clusters that recur ≥ N times (N configurable,
  default e.g. 3), rather than reimplementing similarity clustering. Where a
  friction hasn't gone through consolidation yet (e.g. Task 11 hasn't run),
  fall back to a shared-tags + fuzzy keyword grouping as a lighter-weight
  local heuristic — keep this fallback small and inspectable, it's a
  bridge, not the primary mechanism.
- For each detected recurring pattern, generate a summary (recurring
  friction description, mention count + timespan, possible intervention,
  suggested experiment, suggested success metric) — via the AI provider if
  configured, or a simpler templated summary (mention count + linked
  sources, no generated prose) when `ai.provider: "none"`, so this still
  works with zero AI configured.
- `knowledge/` document type `experiment` gets first-class support using the
  exact fields `00-conventions.md`'s Type-specific fields section already
  defines — no new schema invented here: `experiment_status: proposed |
  active | concluded`, `friction` (the friction document's `id`),
  `duration_days` or `end_date`, and `outcome_result: positive | negative |
  neutral | inconclusive` (required once `experiment_status: concluded`).
  Friction → experiment relationships are **not** a new graph edge kind —
  the `friction` field is a plain reference, scanned by Task 05's
  deterministic extractor into the generic `links` table
  (`kind: "frontmatter"`) exactly like any other frontmatter reference, and
  queried through `getRelated` like any other link. "Positive outcome" (see
  Recent Improvements below) means, precisely, `outcome_result: "positive"`.
- Kaizen page (`src/pages/kaizen.astro`): Capture box (reuses Task 03's
  component), Recurring Friction list, Active Experiments list, Recent
  Improvements (experiments with `experiment_status: concluded` and
  `outcome_result: positive`).
- Tests: the worked example above (repeated context-switching mentions),
  reproduced with fixture documents carrying real `reinforces` edges,
  produces a detected recurring-friction cluster with correct mention count;
  an experiment marked concluded shows up under Recent Improvements; the
  pre-consolidation fallback heuristic is exercised separately and produces
  a reasonable (if less precise) grouping on fixtures that haven't been
  through Task 11 yet.

Out of scope: weekly/automated synthesis scheduling (Task 16 can wire a
cron-style trigger if wanted), the consolidation/reinforcement logic itself
(Task 11 — this task only reads its output), ranking of Kaizen results
(Task 12's hybrid retrieval, if this page ever needs search over its own
content).

## Deliverables

- `src/lib/kaizen/detect.ts`.
- Experiment document type support in the schema (extends Task 02's schema — coordinate rather than duplicate).
- `src/pages/kaizen.astro` + supporting components.
- Tests.

## Acceptance criteria

- [ ] Fixture data reproducing the "4 mentions in 25 days" example above,
      with real graph reinforcement edges, is correctly detected as one
      recurring-friction cluster, not 4 separate items.
- [ ] The Kaizen page renders Recurring Friction, Active Experiments, and
      Recent Improvements sections sourced from real `knowledge/` data, with
      graceful empty states (not fake placeholder content) when a section
      has none yet.
- [ ] Works with `ai.provider: "none"` — recurring friction is still
      detected and shown, just without generated prose summaries.
- [ ] An experiment created via the UI or MCP, then marked
      `experiment_status: concluded` with `outcome_result: positive`,
      appears under Recent Improvements; one concluded with a non-positive
      `outcome_result` does not.
- [ ] The pre-consolidation fallback path (shared tags + keyword overlap)
      is covered by its own test and clearly distinguished, in code and in
      output, from graph-backed detection.

## Notes / assumptions

- This task previously specified avoiding a typed graph for friction ↔
  experiment ↔ outcome relationships ("don't build a typed graph schema for
  this"). That's now reversed: Task 05 builds exactly that graph, and this
  task's job is to consume it well, not to reimplement clustering on top of
  it. Keep whatever remains of the local fallback heuristic small and
  inspectable (a human should be able to see *why* two notes were grouped)
  — that's still good practice even once the graph exists.
