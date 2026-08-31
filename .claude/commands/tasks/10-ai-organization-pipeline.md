# Task 10 — AI organization: raw capture → structured Silver knowledge

**Read first:** `00-conventions.md` (the pipeline, tiers, two time axes).

**Depends on:** Task 03 (captures exist), Task 09 (AI provider interface), Task 02/08 (document write path), Task 05 (knowledge graph — this task's fact extraction feeds it).

## Objective

Turn a Bronze inbox capture into one or more structured Silver `knowledge/`
documents, classified per the product's concept list, with the raw capture
preserved unchanged and provenance linking every derived document back to
it. This task also does the first pass of structured fact extraction that
Task 11's consolidation will act on.

## Scope

In scope:
- `src/lib/organize/classify.ts`: given a capture's text, use the AI provider (Task 09) to split it into distinct thoughts (a brain dump is often several unrelated sentences) and classify each using this mapping as few-shot guidance in the classification prompt — it's the product's own definition of the categories, not just an example:

  ```text
  "I noticed..."          → observation
  "I should..."           → possible action
  "This keeps happening..." → recurring friction
  "What if I tried..."    → experiment
  "I decided..."          → decision
  ```

  Do not turn every observation into a task — that distinction (observation vs. possible action vs. task) is a real quality bar for this classifier, not a nicety. The classifier's labels are guided by the prose above, but the **stored document `type`** must be one of `00-conventions.md`'s ontology — map explicitly:

  | Classifier label | Stored `type` | Notes |
  |---|---|---|
  | observation | `observation` | |
  | recurring friction / friction | `friction` | |
  | idea | `idea` | |
  | possible action | `task` | `status: active` (not yet started) distinguishes it from a firmer plan |
  | task | `task` | |
  | decision | `decision` | |
  | experiment | `experiment` | `experiment_status: proposed` |
  | question | `question` | |
  | insight | `insight` | |
  | project update | `note` | tagged `tags: [project-update]` — not a distinct type, per `00-conventions.md`'s Data model |

- For each classified thought, write a `knowledge/` document (via Task 08's write path) with `type` per the mapping above, `tier: silver`, `sources: [inbox/<capture file>]` for provenance, `generated: {by: "agent", model, prompt_version, at}`, `verified: "unverified"`, `confidence` set from the classifier's own certainty (this is the *stated* confidence field defined in `00-conventions.md` — distinct from `verified`).
- `src/lib/organize/extract-facts.ts`: alongside classification, extract structured fact tuples (subject, predicate, object-or-literal, confidence) from the capture where the text supports it, and write them into the new document's `facts:` frontmatter array per `00-conventions.md`'s Facts schema — `uid` freshly assigned per fact (the `uid` is the persistence mechanism), `valid_from` defaulted to the capture's `created_at` date unless the text states an explicit different date (a simple, deterministic default — not an AI judgment call), and `source_excerpt` set to the specific sentence(s) supporting the fact when identifiable (Task 05 hashes this for drift detection; when extraction can't cleanly isolate an excerpt, omit it and let Task 05 fall back to hashing the whole body). Include `entities:` linking decisions for any named subject/object. This is deliberately a lighter pass than Task 05's schema demands in full — not every classified thought yields a fact; skip extraction rather than force one. These are the inputs Task 11's consolidation reads.
- Explicitly do not rewrite or summarize away the raw capture — it stays in `inbox/` untouched; the new document links back to it, it doesn't replace it.
- Reprocessing guard: instead of a separate derived marker file, use Task 05's `provenance` table (indexed on `capture_path`) as the reverse lookup — a capture is "already processed" if any document's provenance chain names it as a source. This keeps the guard itself rebuildable from files rather than being an undeletable derived artifact that would violate the core invariant if lost or, worse, if kept and `runtime/` were deleted (a stale marker with no matching source doc would incorrectly report "processed"). Edge case: a capture that legitimately produces **zero** documents (e.g. pure noise) must still leave a trace or it will be reprocessed every run — write a minimal `type: note` document with `sources: [<capture>]` and `status: no-op` in that case, rather than leaving it silent.
- A manual trigger (`ren organize` CLI command, or an MCP tool if that's simpler — pick one, don't build both) rather than fully automatic background processing for v1 — keeps this task testable and avoids surprising the user with silent AI writes.
- Tests split into two kinds, deliberately not conflated: **plumbing tests** (using a mocked provider that returns canned, fixture-matched classifications) prove the pipeline correctly turns AI output into the right number of documents, correct `type` mapping, correct provenance, and correct idempotence — these run in CI and don't require a live model. **Prompt-quality** — whether the classifier actually produces the "don't turn every observation into a task" distinction against real, unscripted input — is not something a mocked-provider test can prove; call this out explicitly as a manual/live-model spot-check to run before shipping a prompt change, not a CI-asserted acceptance criterion. A capture with several distinct thoughts produces the expected number of documents with correct `type` classification for the fixture sentences below (plumbing test, mocked), provenance links back to the source capture, extracted facts get a sensible `valid_from`, a second run doesn't duplicate documents for an already-processed capture (verified via the provenance reverse-lookup, not a marker file), and a fixture producing zero real documents still leaves a `status: no-op` trace that prevents reprocessing.

Out of scope: recurring-pattern detection across many captures (Task 14 — that's a separate, higher-level pass), consolidation/deduplication/contradiction handling of the extracted facts (Task 11 — this task only extracts and writes them, it doesn't compare them against existing knowledge), a fully automatic/always-on pipeline (defer to Task 16 if wanted; explicit trigger is safer for v1 given "do not turn every observation into a task" — a human should be able to review before it runs unattended).

## Deliverables

- `src/lib/organize/classify.ts`, `src/lib/organize/extract-facts.ts`, `src/lib/organize/run.ts`.
- CLI or MCP trigger (pick one).
- Tests.

## Acceptance criteria

- [ ] Given this fixture brain dump —

  ```text
  I keep checking Slack too much.

  Lumi transcription still feels slow.

  Maybe I should batch meetings in the afternoon.

  I need to benchmark WhisperKit.

  The kitchen drawer keeps becoming a mess.
  ```

  — classification produces 5 distinct documents: two recurring frictions ("checking Slack too much", "kitchen drawer") — or friction/observation depending on wording, an observation (transcription feels slow), a possible action (batch meetings), and a task (benchmark WhisperKit). Not one blob and not five tasks.
- [ ] Every derived document's `sources` field points back to the originating capture file, and that capture file is byte-identical to before processing. Every derived document has `tier: silver`.
- [ ] At least one fixture sentence with a clear subject/predicate/object (e.g. "I need to benchmark WhisperKit") produces a `facts:` entry with a `uid`, `subject`/`predicate`/`object`, `confidence`, and a `valid_from` (defaulted from the capture's `created_at` since the fixture states no explicit date) — a document satisfying this criterion also passes Task 02's schema validation unmodified.
- [ ] Re-running organization on an already-processed capture does not create duplicate documents, including for a capture whose only prior result was a `status: no-op` document.
- [ ] With `ai.provider: "none"`, running the pipeline is a safe no-op (queues/skips, doesn't error) — confirms Task 09's default is honored here too.

## Notes / assumptions

- "Do not turn every observation into a task" is a real quality bar, not just a guideline. The mocked-provider test only proves the pipeline *can* produce the right shape when the model does its job — it doesn't prove the model does its job. Run a live-model spot-check against the fixture brain dump (and a few unscripted ones) before considering the prompt done, and don't weaken the prompt just to make the mocked test pass more easily.
- Fact extraction here is intentionally shallow. Task 11 is where duplicate/contradiction detection and confidence reinforcement happen — this task's job is only to get candidate facts into files with correct provenance.
