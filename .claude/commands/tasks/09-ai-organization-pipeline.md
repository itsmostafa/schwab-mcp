# Task 09 — AI organization: raw capture → structured knowledge

**Read first:** `00-conventions.md`.

**Depends on:** Task 03 (captures exist), Task 08 (AI provider interface), Task 02/07 (document write path).

## Objective

Turn an inbox capture into one or more structured `knowledge/` documents,
classified per the spec's concept list, with the raw capture preserved
unchanged and provenance linking derived documents back to it.

## Scope

In scope:
- `src/lib/organize/classify.ts`: given a capture's text, use the AI provider (Task 08) to split it into distinct thoughts (a brain dump is often several unrelated sentences) and classify each into one of: observation, friction, idea, possible action, task, decision, experiment, question, insight, project update. Use this mapping as few-shot guidance in the classification prompt — it's the product's own definition of the categories, not just an example:

  ```text
  "I noticed..."          → observation
  "I should..."           → possible action
  "This keeps happening..." → recurring friction
  "What if I tried..."    → experiment
  "I decided..."          → decision
  ```

  Do not turn every observation into a task — that distinction (observation vs. possible action vs. task) is a real quality bar for this classifier, not a nicety.
- For each classified thought, write a `knowledge/` document (via Task 07's write path) with `type` matching the classification, `sources: [inbox/<capture file>]` for provenance, `generated: {by: "ai", ...}`, `verified: "unverified"`.
- Explicitly do not rewrite or summarize away the raw capture — it stays in `inbox/` untouched; the new document links back to it, it doesn't replace it.
- A processing marker so a capture isn't reprocessed twice (e.g. a small `runtime/processed.json` set of already-organized capture ids, or a frontmatter-free sentinel file — pick whichever is simpler; avoid mutating the immutable capture file itself to mark it processed).
- A manual trigger (`sia organize` CLI command, or an MCP tool if that's simpler — pick one, don't build both) rather than fully automatic background processing for v1 — keeps this task testable and avoids surprising the user with silent AI writes.
- Tests (using the `"none"`/mocked provider): a capture with several distinct thoughts produces the expected number of documents with correct `type` classification for the spec's five example sentences, provenance links back to the source capture, and a second run doesn't duplicate documents for an already-processed capture.

Out of scope: recurring-pattern detection across many captures (Task 10 — that's a separate, higher-level pass), a fully automatic/always-on pipeline (defer to Task 12 if wanted; explicit trigger is safer for v1 given "do not turn every observation into a task" — a human should be able to review before it runs unattended).

## Deliverables

- `src/lib/organize/classify.ts`, `src/lib/organize/run.ts`.
- Processing-marker storage.
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
- [ ] Every derived document's `sources` field points back to the originating capture file, and that capture file is byte-identical to before processing.
- [ ] Re-running organization on an already-processed capture does not create duplicate documents.
- [ ] With `ai.provider: "none"`, running the pipeline is a safe no-op (queues/skips, doesn't error) — confirms Task 08's default is honored here too.

## Notes / assumptions

- "Do not turn every observation into a task" is a real quality bar, not just a guideline — the test above using the spec's own example is the acceptance check for it, so don't weaken the classifier to pass a looser test.
