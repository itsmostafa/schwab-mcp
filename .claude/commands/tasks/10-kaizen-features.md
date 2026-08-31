# Task 10 — Kaizen: recurring friction, experiments, dedicated view

**Read first:** `00-conventions.md`.

**Depends on:** Task 09 (classified friction/observation/experiment documents exist to detect patterns over), Task 05 (website shell to add the view to), Task 07 (write path, for creating experiment documents from the UI).

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
- `src/lib/kaizen/detect.ts`: group `friction`/`observation` documents by similarity (start simple — shared tags + fuzzy title/body keyword overlap is enough for v1; do not require embeddings here, Task 11 can strengthen this later) within a rolling window (e.g. last 90 days), surfacing clusters that recur ≥ N times (N configurable, default e.g. 3).
- For each detected recurring pattern, generate the spec's summary shape (recurring friction description, mention count + timespan, possible intervention, suggested experiment, suggested success metric) — via the AI provider if configured, or a simpler templated summary (mention count + linked sources, no generated prose) when `ai.provider: "none"`, so this still works with zero AI configured.
- `knowledge/` document type `experiment` gets first-class support: status (proposed/active/concluded), linked friction, a duration/end date, and an outcome field once concluded.
- Kaizen page (`src/pages/kaizen.astro`): Capture box (reuses Task 03's component), Recurring Friction list, Active Experiments list, Recent Improvements (concluded experiments with a positive outcome) — matching the spec's suggested layout.
- Relationship links: observation ↔ recurring friction ↔ experiment ↔ outcome, stored via `sources`/body markdown links per the OKF convention (links form untyped relationships; semantics come from surrounding prose per OKF's own spec) — don't build a typed graph schema for this.
- Tests: the spec's worked example (repeated context-switching mentions) reproduced with fixture documents produces a detected recurring-friction cluster with correct mention count; an experiment marked concluded shows up under Recent Improvements.

Out of scope: weekly/automated synthesis scheduling (Task 12 can wire a cron-style trigger if wanted), semantic clustering via embeddings (Task 11).

## Deliverables

- `src/lib/kaizen/detect.ts`.
- Experiment document type support in the schema (extends Task 02's schema — coordinate rather than duplicate).
- `src/pages/kaizen.astro` + supporting components.
- Tests.

## Acceptance criteria

- [ ] Fixture data reproducing the "4 mentions in 25 days" example above is correctly detected as one recurring-friction cluster, not 4 separate items.
- [ ] The Kaizen page renders Recurring Friction, Active Experiments, and Recent Improvements sections sourced from real `knowledge/` data, with graceful empty states (not fake placeholder content) when a section has none yet.
- [ ] Works with `ai.provider: "none"` — recurring friction is still detected and shown, just without generated prose summaries.
- [ ] An experiment created via the UI or MCP, then marked concluded, appears under Recent Improvements.

## Notes / assumptions

- Keep the similarity/clustering heuristic small and inspectable (a human should be able to see *why* two notes were grouped) rather than a black-box scoring function — fits "simple yet powerful" and makes the eventual Task 11 embeddings upgrade a clear improvement, not a rewrite.
