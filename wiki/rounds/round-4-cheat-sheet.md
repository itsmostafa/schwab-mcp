---
type: cheat-sheet
last-updated: 2026-04-27
tags: [systems-design, round-4, interview-prep, panther]
---

# Round 4 Cheat Sheet - Systems Design

Round 4 is the **Systems Design** interview: 60 minutes, no coding, open-ended. Darwayne called this the **most revealing** round. The goal is not to guess the perfect architecture; the goal is to show senior judgment under ambiguity.

Use this page as the final pre-interview scan.

---

## What They Are Scoring

| Signal | Show It By Saying / Doing |
|---|---|
| Ambiguity handling | Clarify scale, SLA, product constraints, and success metric before drawing |
| Systems thinking | Draw data flow, components, storage, APIs, and operational boundaries |
| Tradeoff reasoning | For every major choice, state the alternative and why you chose this path |
| Failure awareness | Name what breaks, how you detect it, and how the system degrades |
| Production maturity | Cover monitoring, replay, audit logs, rollbacks, evals, and customer trust |
| Panther fit | Tie architecture back to SOC outcomes: faster triage, less analyst toil, safer automation |

---

## Universal Opener

Use this for any prompt:

> "Before I dive in, I want to make sure I'm solving the right problem. Can I clarify scale, latency, and what we absolutely cannot get wrong? Then I'll walk through the high-level architecture, pick one or two components to deep dive, and close with tradeoffs and failure modes."

Then ask 3-5 of these:

- "How many customers, alerts/day, and events/day should I assume?"
- "Is the output user-facing in seconds, minutes, or batch?"
- "Is this greenfield or does it need to integrate with existing Panther ingestion/query infrastructure?"
- "Are we optimizing first for accuracy, latency, cost, or analyst trust?"
- "What actions can the system take autonomously versus requiring human review?"
- "Do we need multi-tenant isolation and per-tenant policy from day one?"

Restate after answers:

> "So I am designing X for Y scale, with Z latency, prioritizing trust and analyst workflow over full autonomy."

---

## The 6-Step Flow

Memorize this:

```
Clarify -> Restate -> High-level design -> Deep dive -> Tradeoffs -> Operations
```

Suggested timing:

| Phase | Time | What To Do |
|---|---:|---|
| Clarify | 3 min | Scale, SLA, scope, autonomy boundary |
| Restate | 1 min | Repeat the problem and assumptions |
| High-level | 8-10 min | Boxes, arrows, storage, APIs |
| Deep dive | 15-20 min | One or two critical components |
| Tradeoffs | Throughout | Name options and decision criteria |
| Operations | Final 8-10 min | Failure modes, monitoring, replay, rollout |

If stuck, return to the flow out loud: "Let me step back and make the data flow explicit."

---

## Anchor Design: Real-Time Alert Triage

This design covers most Panther-flavored prompts.

```
Security events / alerts
  -> ingestion + tenant routing
  -> normalization + schema validation
  -> durable queue / stream
  -> enrichment fan-out
       - user history
       - asset context
       - threat intel
       - related alerts/logs
  -> retrieval
       - structured query / BM25 for exact IOCs, users, detection IDs
       - vector search for similar alerts and prior analyst decisions
  -> rerank + summarize + token budget
  -> agent orchestrator
       - prompt assembly
       - tool calls
       - structured JSON output
  -> policy gate outside the LLM
  -> analyst queue / escalation / limited auto-close
  -> feedback capture + audit log
  -> eval dataset / index updates / rollout loop
```

Key sentence:

> "The LLM synthesizes ambiguous evidence, but deterministic services own validation, authorization, policy enforcement, side effects, and auditability."

---

## Component Deep Dives

### Ingestion + Queue

Design choice:
- Use a durable stream/queue between ingestion and triage workers.
- Partition by `tenant_id` and possibly `alert_id` or `entity_id`.
- Make processing idempotent with event IDs and dedupe tables.

Tradeoff:
- Streaming adds operational complexity, but protects the system from bursty alert volume and lets workers scale independently.

Failure mode:
- Consumer lag grows during alert spikes.
- Detect with queue depth, age of oldest message, processing latency, and per-tenant lag.

### Enrichment Fan-Out

Design choice:
- Run independent enrichments in parallel: asset inventory, user history, threat intel, previous similar alerts.
- Put timeouts and circuit breakers around slow external dependencies.

Tradeoff:
- Waiting for all enrichment improves evidence quality but hurts latency. Prefer partial results with provenance when SLA matters.

Failure mode:
- Threat intel provider times out.
- Degrade gracefully: continue with missing evidence clearly marked.

### Retrieval Layer

Design choice:
- Use hybrid retrieval: structured filters / BM25 for exact matches, vector search for semantic similarity.
- Always filter by tenant and access policy before retrieval.

Tradeoff:
- Vector search improves recall for similar incidents; exact search is safer for IOCs and IDs.

Failure mode:
- Retrieval returns irrelevant examples and poisons the prompt.
- Detect with retrieval hit-rate, analyst override rate by retrieved-source cohort, and eval-set regression.

### Agent Orchestrator

Design choice:
- Assemble prompts from modules: base instructions, task, tenant policy, evidence, output schema.
- Require structured output with classification, confidence, evidence IDs, and next action.

Tradeoff:
- Modular prompts require more engineering discipline, but prevent drift across agents and make testing easier.

Failure mode:
- Model produces invalid JSON or unsupported action.
- Validate schema, retry once with repair, then route to human queue.

### Policy Gate

Design choice:
- Policy gate lives outside the LLM.
- It checks tenant policy, confidence, severity, action type, and required human approval.

Tradeoff:
- Reduces autonomy, but preserves customer trust and security boundaries.

Failure mode:
- Model recommends an action that policy disallows.
- Block it, log it, and expose the recommendation as analyst context only.

### Feedback Loop

Design choice:
- Capture analyst accept/reject/modify decisions as structured events.
- Use them for eval datasets, retrieval examples, prompt examples, and possible fine-tuning candidates.

Tradeoff:
- Fast learning from feedback is valuable, but blindly training on noisy labels can amplify mistakes.

Failure mode:
- New agent version regresses on one tenant or severity class.
- Use offline eval, shadow mode, canary rollout, A/B comparison, and rollback.

---

## Scenario Cards

### 1. Design A Real-Time Alert Triage System

Lead with:
- "I will optimize for analyst trust first, then latency/cost."

Architecture:
- Ingestion -> queue -> enrichment -> hybrid retrieval -> agent -> policy gate -> analyst workflow -> feedback.

Deep dive:
- Retrieval and policy gate.

Metrics:
- Time-to-triage, analyst override rate, auto-close reversal rate, precision/recall by severity, queue lag.

### 2. Design Text-To-Search For Security Logs

Lead with:
- "This should produce inspectable queries, not just natural-language answers."

Architecture:
- User question -> intent parser -> schema/entity resolver -> query generator -> query validator -> execute against log store -> rank/summarize results -> cite raw events.

Deep dive:
- Query validation and schema grounding.

Tradeoff:
- LLM-generated query gives flexibility; constrained query templates reduce dangerous or expensive queries.

Failure modes:
- Invalid query, expensive full scan, hallucinated field, unauthorized tenant access.

### 3. Design Detection Code Generation

Lead with:
- "Generated detections need tests and review before they affect production alerting."

Architecture:
- Natural-language spec -> RAG over detection examples/docs -> code generation -> static checks -> unit tests against sample logs -> sandbox run -> human review -> versioned deployment.

Deep dive:
- Test harness and human review workflow.

Tradeoff:
- Fast draft generation versus safety of deploying detection logic that can create false positives or miss real threats.

Failure modes:
- Overbroad rule floods analysts; underbroad rule misses attacks; generated code uses unsupported APIs.

### 4. Design Collective Intelligence For SOC Analysts

Lead with:
- "The core problem is turning analyst decisions into reliable product improvement without leaking tenant data."

Architecture:
- Analyst action stream -> normalize labels -> quality filters -> tenant/privacy boundary -> aggregate patterns -> update retrieval/evals/prompts -> canary rollout.

Deep dive:
- Signal quality and tenant isolation.

Tradeoff:
- Cross-customer learning improves coverage, but privacy and tenant boundaries must be explicit.

Failure modes:
- Noisy labels, feedback loops that reinforce bad decisions, tenant data leakage.

### 5. Design The Data Pipeline Feeding AI Agents

Lead with:
- "The AI system is only as good as the normalized, searchable, permissioned data layer beneath it."

Architecture:
- Log ingestion -> parsing/normalization -> schema evolution -> storage -> structured index + vector index -> retrieval APIs -> retention/deletion.

Deep dive:
- Schema evolution and dual indexing.

Tradeoff:
- Schema-on-write improves query quality; schema-on-read preserves flexibility for varied security logs.

Failure modes:
- Parser drift, index lag, duplicate events, retention mismatch, tenant isolation bug.

---

## Tradeoff Bank

| Decision | Default Answer | Alternative | Why |
|---|---|---|---|
| Streaming vs batch | Streaming for triage | Batch for reports/training | Triage needs continuous updates and bounded lag |
| RAG vs fine-tuning | RAG first | Fine-tune later | Security knowledge changes; evidence must be inspectable |
| Hybrid retrieval vs vector only | Hybrid | Vector only | IOCs and detection IDs need exact matching |
| Dedicated vector DB vs pgvector | Depends on scale | Either | pgvector for simplicity; dedicated service for independent scaling/filter latency |
| Smaller vs larger model | Tier by risk | Single big model | Fast model for first pass; stronger model/verifier for high-risk cases |
| Human review vs autonomy | Human gate for high-stakes actions | Full autonomy | Blocking, escalation, or closure affects trust and safety |
| Store summaries vs raw events | Store both | One only | Summaries help speed; raw events are needed for audit and reprocessing |
| Multi-tenant shared infra vs per-tenant isolation | Shared compute with hard isolation | Dedicated tenant stacks | Shared is cost-efficient; isolation must be enforced in auth, retrieval, and audit |

Pattern:

> "I would choose A under these assumptions, because B costs us X. If the interviewer changes the constraint to Y, I would switch."

---

## Failure Mode Bank

Use this whenever they ask "what could go wrong?"

| Failure | Detection | Mitigation |
|---|---|---|
| Queue backlog | Consumer lag, old message age | Autoscale workers, shed non-critical enrichment, prioritize severity |
| Duplicate processing | Duplicate event IDs/actions | Idempotency keys, dedupe table, exactly-once side-effect boundary |
| Enrichment dependency down | Timeout/error rate | Circuit breaker, cached/stale data, partial result |
| Retrieval pollution | Bad evidence in eval traces, overrides | Tenant filters, source scoring, reranking, retrieval evals |
| Prompt injection in logs | Suspicious instruction patterns, tool attempts | Treat logs as untrusted data, tool allowlists, policy gate |
| Invalid model output | Schema validation failures | JSON schema validation, retry/repair, human fallback |
| Bad autonomous action | Reversal rate, audit review | Limit autonomy, confidence/severity thresholds, approval workflow |
| Model drift | Distribution shift, precision drop | Eval suites, shadow mode, canary, rollback |
| Tenant data leak | Access audit anomalies | Authz at retrieval/tool layer, per-tenant filters, tests |
| Cost spike | Token spend, calls/alert | Caching, summarization, model routing, rate limits |

---

## Metrics To Name

### System Metrics

- Ingestion rate, queue lag, processing latency p50/p95/p99
- Enrichment timeout/error rate
- Retrieval latency and hit rate
- Model latency, token cost, schema-valid output rate
- Audit-log write success rate

### SOC Outcome Metrics

- Time-to-triage
- Analyst override rate
- Auto-close reversal rate
- False negative rate by severity
- Alerts handled per analyst
- Mean time to respond

### Safety Metrics

- Unauthorized tool attempt rate
- Policy-gate block rate
- Prompt-injection detection rate
- Human escalation rate for low-confidence cases

### Learning Metrics

- Eval-set precision/recall by tenant/severity
- Drift in classification distribution
- Retrieval quality by source type
- Canary version win/loss against baseline

---

## Numbers To Keep Handy

Use only after stating assumptions.

| Number | Use |
|---|---|
| 1M events/day ~= 11.5 events/sec | Quick event-volume conversion |
| 1B events/day ~= 11,500 events/sec | Panther-scale pipeline discussion |
| Same-DC network RTT ~= 0.5 ms | Service-call intuition |
| Cross-region RTT ~= 30-100 ms | Avoid cross-region synchronous calls |
| SSD read ~= 100 us | Storage intuition |
| Seconds | Reasonable first-pass enrichment target |
| Minutes | Reasonable analyst-ready triage target if evidence quality matters |

---

## Things To Say Explicitly

- "I would keep deterministic policy enforcement outside the model."
- "Every model decision should include evidence IDs so the analyst can inspect raw events."
- "I would design for graceful degradation: partial enrichment is better than blocking triage."
- "I would launch in shadow mode before allowing autonomous actions."
- "I would make tenant isolation a retrieval-layer and authorization-layer invariant, not a prompt instruction."
- "I would use analyst feedback, but I would not blindly train on it without quality filters."

---

## Red Flags To Avoid

- Jumping into boxes before clarifying requirements.
- Saying "just use an LLM" without data flow, tools, policy, and monitoring.
- Letting the LLM enforce permissions or decide side effects directly.
- Ignoring multi-tenancy and tenant data leakage.
- Discussing only AI metrics and not SOC/customer metrics.
- Claiming full autonomy for high-stakes security actions too early.
- Saying "it depends" without naming the dependency and decision rule.

---

## Strong Close

End with a concise summary:

> "The design is a streaming, multi-tenant triage pipeline with enrichment and hybrid retrieval feeding an agent, but the policy gate and side effects stay deterministic. The biggest risks are bad evidence, model drift, tenant isolation, and unsafe automation, so I would launch with shadow mode, strong audit trails, evals, canaries, and human review for high-stakes actions."

Then ask:

- "At Panther, where is the hardest bottleneck today: ingestion scale, retrieval quality, evals, or analyst workflow integration?"
- "How much autonomy do customers currently want from SOC agents versus analyst-assist workflows?"
- "How do you think about cross-customer learning while preserving customer trust and data boundaries?"

---

## See Also

- [[rounds/systems-design]] - full Round 4 guide
- [[concepts/systems-design-patterns]] - pattern reference
- [[concepts/soc-domain]] - SOC workflow context
- [[concepts/agentic-ai]] - RAG, feedback loops, and agent architecture
