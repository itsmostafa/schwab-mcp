# Task 09 — AI provider abstraction and local embedding backend

**Read first:** `00-conventions.md` (AI provider vs. embedding backend — two
independent config axes; semantic search on by default).

**Depends on:** Task 01 (config — `ai.*`/`search.embeddings.*` blocks and
`$REN_DATA_DIR/models/` resolution this task extends), Task 04 (derived
index — the `vectors` table and no-op `EmbeddingBackend` interface this task
fills in).

## Objective

Two separate, minimal, swappable interfaces:

1. "ask a model to do X" (`ai.provider`) — so nothing later (Task 10's
   organization pipeline, Task 11's consolidation prose) is hardwired to one
   vendor, and the app keeps working with zero AI configured.
2. "turn text into a vector" (the embedding backend) — so semantic search
   (Task 12) works locally by default without depending on the LLM
   interface at all.

Keep these genuinely independent: `ai.provider: "none"` must not disable
embeddings, and the embedding backend being unavailable must not be treated
as "AI is off" anywhere in later tasks.

## Scope

In scope:
- `src/lib/ai/provider.ts`: a small interface — something like `complete(prompt, opts) -> text` and/or `classify(text, labels) -> label` — sized to what Task 10 actually needs, not a general chat-completion SDK wrapper. Keep the surface area small.
- Two concrete adapters: **Anthropic** and **Ollama** (per `00-conventions.md` — the rest of the documented provider list stays as documented-but-unimplemented interface targets, not stub files, unless a stub genuinely clarifies the interface).
- A `"none"` provider (default) that makes AI-dependent code paths no-op cleanly — e.g. Task 10's classification step should skip/queue rather than error when `ai.provider: "none"`.
- `src/lib/search/embeddings-onnx.ts`: an implementation of Task 04's `EmbeddingBackend` interface (`id`, `dims`, `embed(texts) -> vectors`) backed by an in-process ONNX model (transformers.js, default `Xenova/all-MiniLM-L6-v2`, 384 dimensions). Model weights download once and cache under `$REN_DATA_DIR/models/` — **never** under `runtime/` (per `00-conventions.md`: deleting `runtime/` must never trigger a re-download as a side effect of an otherwise-offline rebuild). `id` (e.g. `"onnx:Xenova/all-MiniLM-L6-v2"`) is what Task 04 stores as `embedding_version` in `index_meta`; changing the model bumps this and Task 04's version-tracking machinery handles reprocessing.
- Wire this backend into Task 04's incremental indexer in place of the no-op provider — chunks needing an embedding (per the `chunk_hash` reuse check Task 04 already does) get one computed and written to the `vectors` table.
- Offline/first-run behavior: the first embed call downloads the model; document this explicitly, and make the "model not yet cached, network unavailable" case degrade (skip embedding, log, leave `vectors` empty for that chunk) rather than crash the indexer — this is what lets Task 12's hybrid search fall back to FTS-only per `00-conventions.md`'s degrade contract.
- Config wiring in `ren.config.ts`: `ai.provider` (+ per-adapter secrets via env vars only, never in config files) and `search.embeddings.{backend,model,dim,cacheDir}` as two separate blocks — see the config shape in Task 01.
- A visible indicator (log line at minimum; UI badge is a nice-to-have, not required) whenever external inference is actually enabled — "clearly identify when external inference is enabled" is an explicit privacy requirement, not just a nicety. The embedding backend running locally does not count as "external inference" and does not need this indicator; only `ai.provider` set to a network-calling adapter does.
- Tests: provider selection from config, `"none"` AI provider never throws and never makes a network call, Anthropic/Ollama adapters mockable (don't require live credentials in CI), the ONNX embedding backend is mockable/swappable for a deterministic fake in tests (no live model download in CI), `embed()` output dimension matches `dims`, an indexing run with a mocked "model unavailable" embedding backend completes successfully with empty `vectors` rows rather than failing.

Out of scope: the organization pipeline itself (Task 10), hybrid ranking that consumes the vectors (Task 12), embedding compression (explicitly deferred — the vector layer is abstracted precisely so this can be added later without a rewrite).

## Deliverables

- `src/lib/ai/provider.ts` (interface + `"none"`).
- `src/lib/ai/anthropic.ts`, `src/lib/ai/ollama.ts`.
- `src/lib/search/embeddings-onnx.ts`.
- Config additions (both `ai.*` and `search.embeddings.*`).
- Tests.

## Acceptance criteria

- [ ] With `ai.provider: "none"` (the default), the app builds, runs, and passes its full test suite with zero network calls to any AI provider — and this holds independently of whether the embedding backend is enabled.
- [ ] Switching to `anthropic` or `ollama` in config routes calls through the matching adapter with no other code changes.
- [ ] No API key or secret appears in any committed file; both adapters read credentials from env vars.
- [ ] Enabling either external `ai.provider` produces a visible log statement confirming external inference is active; running the local embedding backend does not.
- [ ] Reindexing a fixture set with the ONNX backend enabled populates `vectors` rows with the correct `dim` and `model`/`embedding_version`; with the backend mocked as unavailable, indexing still completes and leaves those rows empty rather than erroring.
- [ ] Model weights are cached under `$REN_DATA_DIR/models/`, confirmed by inspection not to be written under `runtime/`.

## Notes / assumptions

- Do not build a provider registry/plugin loader for six LLM providers — a small switch/factory over the two implemented adapters plus `"none"` is enough until a third is actually needed. Simple yet powerful.
- Same principle applies to the embedding backend: one concrete implementation (ONNX) behind the interface Task 04 already defined is enough for v1 — the interface exists so a second backend is a new file, not a rewrite.
