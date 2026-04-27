---
type: question-bank
last-updated: 2026-04-27
tags: [round-4, systems-design, interview-questions, soc, architecture]
---

# Round 4 Interview Questions - Systems Design

Use this as a spoken-answer drill sheet for the systems design round. The best answers show senior judgment: clarify assumptions, draw the data flow, name boundaries, discuss tradeoffs, and close with operations.

## 1. What clarifying questions would you ask before designing the system?

Strong answer:

"Before I start, I want to clarify a few things:

- What scale should I assume: alerts per day, events per day, number of tenants, and number of analysts?
- What are the latency requirements: real-time seconds, near-real-time minutes, or batch?
- Is this greenfield, or does it need to integrate with existing infrastructure?
- What is the most important constraint: accuracy, latency, cost, analyst trust, or customer isolation?
- What enrichment sources and internal services are available?
- Which dependencies are reliable enough for the critical path?
- What actions can the system take autonomously versus requiring human review?

Unless you want me to optimize for something else, I will assume an event-driven system with real-time triage requirements and strong multi-tenant isolation."

Why this works:

- It frames ambiguity as part of the design problem.
- It gets scale, SLA, constraints, dependencies, and autonomy on the table before drawing boxes.

## 2. Design the high-level architecture for a real-time alert triage system.

Strong answer:

"At a high level, I would design it as an event-driven pipeline.

Alerts enter an ingestion pipeline, then go through deterministic intake: validation, schema normalization, tenant routing, and idempotency checks. From there, a durable queue decouples intake from triage workers.

Triage workers run enrichment in parallel against identity, asset inventory, threat intel, logs, and historical cases. Then the retrieval layer runs hybrid search: BM25 or structured search for exact IOCs and vector search for similar alerts or analyst decisions.

Retrieved evidence is reranked, deduped, and summarized. The agent orchestrator assembles a prompt from base instructions, task instructions, tenant policy, evidence, and output schema. The LLM produces a structured answer, and a policy gate outside the LLM decides whether to auto-close, escalate, or route to an analyst.

Every step writes audit events, and analyst feedback flows back into eval datasets, retrieval examples, and controlled rollout of improvements."

Architecture shorthand:

```text
ingestion
-> validation / normalization / tenant routing
-> queue
-> enrichment fan-out
-> hybrid retrieval
-> rerank / summarize
-> agent orchestrator
-> policy gate
-> analyst workflow
-> audit + feedback loop
```

## 3. What is enrichment?

Strong answer:

"Enrichment adds context to raw alerts so they become actionable.

For a SOC alert, enrichment could include user context, asset criticality and ownership, threat-intel reputation, historical alerts and decisions, related logs, recent authentication behavior, and customer runbooks.

I would implement enrichment as async fan-out to multiple services, with timeouts, circuit breakers, caching where appropriate, and partial-result handling. If threat intel is slow, the system should still produce a bounded answer with that missing evidence clearly marked."

Key tradeoff:

- Waiting for all enrichments improves evidence quality but hurts latency.
- Partial results preserve responsiveness but need clear provenance and confidence handling.

## 4. Where is the boundary between deterministic code and the LLM?

Strong answer:

"Deterministic code handles validation, routing, authorization, enrichment, retrieval, deduplication, prompt construction, schema validation, side effects, and policy enforcement.

The LLM is responsible for classification, summarization, prioritization, and reasoning over curated evidence.

The LLM must never control permissions, tenant isolation, policy enforcement, or direct write actions. Those are handled by services outside the model, especially the policy gate."

Good line:

> "The LLM can recommend; the platform enforces."

Examples:

- LLM can recommend `close_alert`; policy checks severity, confidence, tenant policy, and human-review requirements.
- LLM can request `query_logs`; tool runtime validates tenant, fields, time range, and cost limits.

## 5. Why use hybrid retrieval: BM25 plus vector search?

Strong answer:

"BM25 or structured search is high precision for exact matches like IOCs, hashes, usernames, hostnames, and detection IDs.

Vector search is higher recall for semantic similarity: behavioral patterns, similar investigations, runbook language, and analyst notes that do not share exact terms.

Using only exact search misses novel or paraphrased patterns. Using only vector search risks noise and can miss exact indicators.

So I would run both in parallel, merge results, apply tenant and access filters, optionally rerank, and pass a curated set to the LLM."

Design detail:

- Exact search is often the right first path for IOCs.
- Vector search helps with "this looks like prior suspicious behavior" or "similar investigation narrative."
- Reranking and diversity controls prevent 20 duplicate examples from filling the prompt.

## 6. How do you select what goes into the LLM prompt?

Strong answer:

"I optimize for information density per token.

First I pre-filter by tenant, access policy, time range, detection type, entity, severity, and data source. Then I deduplicate and cluster repetitive events, select representative samples, rank by relevance, optionally rerank with a cross-encoder, and summarize raw logs into structured evidence.

The prompt should include normalized alert fields, high-signal enrichment, retrieved evidence with IDs, tenant policy constraints, and a strict output schema.

The goal is high-signal, non-redundant context within token limits, with enough IDs for audit and analyst drill-down."

Do not say:

- "I would just put everything into the context window."

Say instead:

> "The model should reason over evidence, not become the storage layer."

## 7. What are key failure modes and mitigations?

Strong answer:

"I would group failure modes by layer.

At ingestion, duplicate or malformed alerts are handled with schema validation, idempotency keys, and dead-letter queues.

At enrichment, slow dependencies are handled with timeouts, circuit breakers, cached/stale context, and partial-result degradation.

At retrieval, bad or irrelevant evidence is handled with tenant filters, source scoring, reranking, retrieval evals, and monitoring hit rate.

At the model layer, hallucination and invalid output are handled with structured schemas, citation requirements, output validation, retry or repair, and human fallback.

At the action layer, unsafe actions are prevented by policy gates, RBAC, human approval, audit logging, and limited autonomy.

At the product layer, drift is handled with offline evals, shadow mode, canaries, tenant-specific monitoring, and rollback."

Key principle:

- Design failures to be observable, bounded, and safe.

## 8. How do you design the feedback loop?

Strong answer:

"I would capture structured analyst interactions: accept, reject, edit, escalation, final disposition, reopen events, severity changes, and time-to-triage.

Each feedback event should be stored with alert ID, evidence IDs, prompt version, model version, retrieval index version, confidence, and policy decision.

Then I would use that data to update the RAG index with confirmed examples, build eval datasets from overrides and reopens, tune prompts and thresholds through offline testing, and identify candidates for fine-tuning once the labels are high quality.

I would deploy improvements through shadow mode, canaries, or A/B tests, and monitor override rate, reversal rate, latency, cost, and tenant-specific regressions."

Important nuance:

- Feedback improves retrieval, evaluation, and decision quality; it is not just prompt editing.
- Do not blindly train on noisy clicks.
- Keep tenant privacy and cross-customer aggregation boundaries explicit.

## See Also

- [[rounds/systems-design]] - full Round 4 guide
- [[rounds/round-4-cheat-sheet]] - final-scan systems design cheat sheet
- [[concepts/systems-design-patterns]] - reusable architecture patterns
- [[concepts/soc-domain]] - SOC workflow context
