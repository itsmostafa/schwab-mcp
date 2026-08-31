# Task 08 — AI provider abstraction

**Read first:** `00-conventions.md`.

## Objective

A minimal, swappable interface for "ask a model to do X" so nothing later
(Task 09's organization pipeline, Task 11's reranking) is hardwired to one
vendor, and the app keeps working with zero AI configured.

## Scope

In scope:
- `src/lib/ai/provider.ts`: a small interface — something like `complete(prompt, opts) -> text` and/or `classify(text, labels) -> label` — sized to what Task 09 actually needs, not a general chat-completion SDK wrapper. Keep the surface area small.
- Two concrete adapters: **Anthropic** and **Ollama** (per `00-conventions.md` — the rest of the spec's provider list stays as documented-but-unimplemented interface targets, not stub files, unless a stub genuinely clarifies the interface).
- A `"none"` provider (default) that makes AI-dependent code paths no-op cleanly — e.g. Task 09's classification step should skip/queue rather than error when `ai.provider: "none"`.
- Config wiring in `sia.config.ts`: `ai.provider`, plus whatever each adapter needs (API key env var name for Anthropic, base URL for Ollama) — secrets via env vars only, never in config files, per the spec's privacy section.
- A visible indicator (log line at minimum; UI badge is a nice-to-have, not required) whenever external inference is actually enabled — "clearly identify when external inference is enabled" is an explicit privacy requirement, not just a nicety.
- Tests: provider selection from config, `"none"` provider never throws and never makes a network call, Anthropic/Ollama adapters mockable (don't require live credentials in CI).

Out of scope: the organization pipeline itself (Task 09), embeddings provider (Task 11 — may reuse this interface or need a small addition, decide there).

## Deliverables

- `src/lib/ai/provider.ts` (interface + `"none"`).
- `src/lib/ai/anthropic.ts`, `src/lib/ai/ollama.ts`.
- Config additions.
- Tests.

## Acceptance criteria

- [ ] With `ai.provider: "none"` (the default), the app builds, runs, and passes its full test suite with zero network calls to any AI provider.
- [ ] Switching to `anthropic` or `ollama` in config routes calls through the matching adapter with no other code changes.
- [ ] No API key or secret appears in any committed file; both adapters read credentials from env vars.
- [ ] Enabling either external provider produces a visible log statement confirming external inference is active.

## Notes / assumptions

- Do not build a provider registry/plugin loader for six providers per the spec's list — a small switch/factory over the two implemented adapters plus `"none"` is enough until a third is actually needed. Simple yet powerful.
