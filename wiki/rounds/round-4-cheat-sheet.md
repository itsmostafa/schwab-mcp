---
type: cheat-sheet
last-updated: 2026-04-28
tags: [systems-design, round-4, interview-prep, panther]
---

# Round 4 Cheat Sheet - Systems Design

**Goal:** be clear under ambiguity. Round 4 is 60 minutes, no coding, open-ended. Darwayne called it the **most revealing** round.

**Default posture:** SOC trust first, then latency/cost/autonomy.

**One-line anchor:** "Event-driven, decoupled, observable; the LLM recommends, the platform enforces."

---

## Table of Contents

- [0. 5-Second Map](#0-5-second-map)
- [1. Universal Opener](#1-universal-opener)
- [2. Interview Flow](#2-interview-flow)
- [3. Anchor Architecture: AI Alert Triage](#3-anchor-architecture-ai-alert-triage)
- [4. Component Cards](#4-component-cards)
- [5. Scenario Cards](#5-scenario-cards)
- [6. Failure Mode Bank](#6-failure-mode-bank)
- [7. Failure Mode Brush-Up](#7-failure-mode-brush-up)
- [8. Tradeoff Bank](#8-tradeoff-bank)
- [9. Metrics to Name](#9-metrics-to-name)
- [10. Scaling at 10x](#10-scaling-at-10x)
- [11. Safe Autonomy](#11-safe-autonomy)
- [12. Things to Say](#12-things-to-say)
- [13. Red Flags to Avoid](#13-red-flags-to-avoid)
- [14. Strong Close](#14-strong-close)
- [See Also](#see-also)

---

## 0. 5-Second Map

| Topic | Say | Remember |
|---|---|---|
| Start | "Let me clarify assumptions before designing." | Frame ambiguity before boxes |
| Architecture | "I would structure this as an event-driven pipeline." | Queue decouples bursty work |
| Enrichment | "Better partial answers than blocked systems." | Timeouts, retries, circuit breakers |
| Retrieval | "BM25 for exact, vector for similar." | Precision + recall = signal |
| Prompting | "Optimize for signal per token." | Evidence IDs + strict schema |
| LLM boundary | "LLM recommends; system enforces." | Auth/policy/actions outside model |
| Failure | "Observable, bounded, safe." | Detect + degrade + recover |
| Feedback | "Feedback improves retrieval, evals, and decisions." | Do not blindly train on noisy labels |
| Autonomy | "Controlled autonomy, not blind automation." | Low risk auto; high risk HITL |
| Scaling | "Design for graceful degradation." | Backpressure, cache, shard, route models |

---

## 1. Universal Opener

Use this verbatim:

> "Before I draw the system, I want to clarify scale, latency, success metrics, and what the system must not get wrong. Then I will propose the high-level architecture, deep dive on the riskiest components, and close with tradeoffs and failure modes."

Ask 3-5:

| Clarify | Why |
|---|---|
| Alerts/day, events/day, tenants, analysts? | Scale + partitioning |
| Real-time, minutes, or batch? | Latency + architecture |
| Optimize for accuracy, cost, latency, or trust? | Tradeoff order |
| What can AI do autonomously? | Policy gate + HITL |
| Existing Panther infra or greenfield? | Integration constraints |
| Multi-tenant isolation from day one? | Auth/retrieval/storage design |
| Dependencies: SIEM, logs, threat intel, identity? | Enrichment + failure modes |

Restate:

> "So I am designing X at Y scale with Z latency, prioritizing analyst trust over full autonomy."

---

## 2. Interview Flow

```text
Clarify -> Restate -> Architecture -> Deep dive -> Tradeoffs -> Operations
```

| Phase | Time | Output |
|---|---:|---|
| Clarify | 3 min | Scale, SLA, scope, autonomy |
| Restate | 1 min | Assumptions + success metric |
| Architecture | 8-10 min | Boxes, arrows, stores, APIs |
| Deep dive | 15-20 min | 1-2 risky components |
| Tradeoffs | Throughout | A vs B + decision rule |
| Operations | 8-10 min | Failures, metrics, rollout |

If stuck:

> "Let me step back and make the data flow explicit."

---

## 3. Anchor Architecture: AI Alert Triage

```text
Alerts/events
-> validation + normalization + tenant routing
-> durable queue
-> enrichment fan-out
-> hybrid retrieval
-> rerank + summarize
-> LLM / agent reasoning
-> policy gate
-> analyst action / limited auto action
-> audit log + feedback loop
```

| Stage | Job | Watch For |
|---|---|---|
| Ingestion | Accept events, validate schema, route tenant | duplicates, parser drift |
| Queue | Decouple spikes from workers | lag, backpressure |
| Enrichment | Add user/asset/log/threat intel/history | slow dependencies |
| Retrieval | Find exact IOCs + similar incidents | tenant filters, bad context |
| Rerank/summarize | Fit best evidence into prompt | token waste, lost provenance |
| LLM/agent | Classify, summarize, recommend | hallucination, invalid output |
| Policy gate | Enforce rules outside model | unsafe action, RBAC |
| Analyst workflow | Human review, escalation, closure | trust, override rate |
| Feedback | Capture decisions for eval/retrieval/prompting | noisy labels, drift |

Key sentence:

> "The LLM synthesizes ambiguous evidence, but deterministic services own auth, validation, policy enforcement, side effects, and auditability."

---

## 4. Component Cards

### Ingestion + Queue

| Field | Quick Reference |
|---|---|
| Say | "Use a durable stream partitioned by tenant, with idempotent processing." |
| Why | Decouples alert spikes from triage workers. |
| Design | `tenant_id` partitioning, event IDs, idempotency keys, dedupe table, DLQ. |
| Failure | Queue lag, duplicate processing, poison messages. |
| Detect | Oldest message age, consumer lag, retry/DLQ rate. |
| Mitigate | Autoscale workers, backpressure, severity priority, idempotency keys. |

### Idempotency + Idempotency Keys

| Field | Quick Reference |
|---|---|
| Say | "Retries are inevitable, so writes and side effects need idempotency keys." |
| Meaning | Same request/event can be processed more than once but produces one durable outcome. |
| Key shape | `tenant_id + source + external_event_id` for incoming events; client-provided `Idempotency-Key` for API writes/actions. |
| Store | Idempotency table/cache with key, request hash, status, result pointer, created/expires timestamps. |
| Flow | On write: check key -> if completed, return stored result -> if in progress, reject or wait -> else reserve key and process. |
| Side effects | Apply the same key to alert creation, ticket creation, notification sends, and auto-actions. |
| TTL | Keep keys long enough to cover retry windows and replay risk; use longer retention for high-impact actions. |
| Failure | Duplicate alerts, duplicate tickets, repeated analyst notifications, repeated containment actions. |
| Mitigate | Atomic key reservation, unique DB constraint, request-hash validation, transaction/outbox for external calls. |

Spoken example:

> "If Panther receives the same alert twice because a producer retries after a timeout, I would compute an idempotency key from tenant, source, and source alert ID. The triage service first reserves that key in a dedupe/idempotency store. If the key already completed, it returns the existing alert/triage result instead of creating another alert or rerunning side effects. If the same key arrives with a different request hash, that is a client or producer bug and should be rejected or sent to review."

### Enrichment Fan-Out

| Field | Quick Reference |
|---|---|
| Say | "Better partial answers than blocked systems." |
| Adds | User, asset, identity, threat intel, prior alerts, related logs. |
| Design | Parallel async calls with timeouts, retries, circuit breakers, cache. |
| Tradeoff | Wait for all evidence = better quality; partial evidence = lower latency. |
| Failure | Threat intel or identity provider times out. |
| Mitigate | Return partial results with missing evidence clearly marked. |

### Hybrid Retrieval

| Field | Quick Reference |
|---|---|
| Say | "BM25 gives precision; vectors give recall." |
| BM25 | Exact IPs, hashes, usernames, detection IDs, raw log terms. |
| Vector | Similar cases, analyst decisions, semantic patterns. |
| Design | Run both -> tenant filter -> merge -> rerank -> dedupe -> cite evidence IDs. |
| Failure | Irrelevant examples poison prompt. |
| Mitigate | Source scoring, reranking, retrieval evals, analyst override tracking. |

### Prompt Construction

| Field | Quick Reference |
|---|---|
| Say | "Optimize for signal per token." |
| Inputs | Tenant policy, task, evidence, output schema, allowed tools. |
| Prep | Filter, dedupe, cluster, rerank, summarize. |
| Require | Classification, confidence, evidence IDs, rationale, next action. |
| Failure | Hallucinated fields, missing evidence, invalid JSON. |
| Mitigate | Strict schema, validation, one repair retry, human fallback. |

### Deterministic vs LLM Boundary

| Deterministic | LLM |
|---|---|
| Auth/RBAC | Summarization |
| Tenant isolation | Classification |
| Validation | Reasoning over evidence |
| Routing | Recommendation |
| Tool execution | Explanation |
| Retries/timeouts | Draft query/action |
| Policy gates | Confidence estimate |
| Audit logs | Analyst-facing narrative |

> "The model reasons over data; it does not store authority or execute side effects."

### Policy Gate

| Field | Quick Reference |
|---|---|
| Say | "Policy gate lives outside the LLM." |
| Checks | Tenant rules, severity, confidence, action type, RBAC, approval requirement. |
| Risk model | Low risk -> auto; high risk -> human in the loop. |
| Failure | Model recommends an unsupported or unsafe action. |
| Mitigate | Block action, log it, show recommendation as analyst context only. |

### Feedback Loop

| Field | Quick Reference |
|---|---|
| Say | "Feedback improves the system, not just prompts." |
| Log | Alert ID, evidence IDs, prompt/model version, tool calls, confidence, analyst action. |
| Use | Eval datasets, retrieval examples, thresholds, prompts, fine-tuning candidates. |
| Warning | Noisy labels can reinforce bad decisions. |
| Rollout | Offline eval -> shadow mode -> canary -> A/B -> rollback. |

---

## 5. Scenario Cards

### Real-Time Alert Triage

| Field | Quick Reference |
|---|---|
| Lead | "I will optimize for analyst trust first, then latency and cost." |
| Flow | Ingest -> queue -> enrich -> hybrid retrieval -> agent -> policy -> analyst -> feedback. |
| Deep dive | Retrieval + policy gate. |
| Metrics | Time-to-triage, override rate, auto-close reversal, queue lag, precision/recall by severity. |
| Risks | Bad evidence, tenant leaks, unsafe action, model drift. |

### Text-to-Search for Security Logs

| Field | Quick Reference |
|---|---|
| Lead | "This should produce inspectable queries, not opaque answers." |
| Flow | Question -> intent -> schema/entity resolver -> query generator -> validator -> execute -> cite raw events. |
| Deep dive | Query validation + schema grounding. |
| Tradeoff | Flexible LLM query vs safer constrained templates. |
| Risks | Full scan, hallucinated field, invalid query, unauthorized tenant access. |

### Detection Code Generation

| Field | Quick Reference |
|---|---|
| Lead | "Generated detections need tests and review before production alerting." |
| Flow | Spec -> RAG examples/docs -> code -> static checks -> unit tests -> sandbox -> human review -> deploy. |
| Deep dive | Test harness + review workflow. |
| Tradeoff | Fast draft generation vs false-positive/false-negative risk. |
| Risks | Overbroad rule floods analysts; underbroad rule misses attacks; unsupported APIs. |

### Collective Intelligence for SOC Analysts

| Field | Quick Reference |
|---|---|
| Lead | "The hard part is learning across decisions without leaking tenant data." |
| Flow | Analyst actions -> normalize labels -> quality filters -> privacy boundary -> aggregate patterns -> eval/retrieval/prompt updates. |
| Deep dive | Signal quality + tenant isolation. |
| Tradeoff | Cross-customer learning improves coverage; privacy boundaries preserve trust. |
| Risks | Noisy labels, feedback loops, tenant data leakage. |

### Data Pipeline Feeding AI Agents

| Field | Quick Reference |
|---|---|
| Lead | "AI is only as good as the normalized, searchable, permissioned data beneath it." |
| Flow | Logs -> parse/normalize -> schema evolution -> storage -> structured index + vector index -> retrieval APIs -> retention/deletion. |
| Deep dive | Schema evolution + dual indexing. |
| Tradeoff | Schema-on-write improves query quality; schema-on-read preserves flexibility. |
| Risks | Parser drift, index lag, duplicate events, retention mismatch, tenant bug. |

---

## 6. Failure Mode Bank

| Layer | Failure | Counter |
|---|---|---|
| Ingestion | duplicates, poison messages | idempotency keys, dedupe table, DLQ |
| Queue | backlog, old messages | autoscale, backpressure, severity priority |
| Enrichment | slow/down dependencies | timeout, retry, circuit breaker, cache, partial result |
| Retrieval | bad evidence | tenant filters, rerank, retrieval evals |
| Prompt | token bloat, lost evidence | cluster, summarize, evidence IDs |
| LLM | hallucination, invalid JSON | strict schema, validator, repair retry |
| Tools | unsafe call | allowlist, RBAC, policy gate |
| Actions | bad auto-close/block | confidence + severity thresholds, human approval |
| Model | drift/regression | offline evals, shadow mode, canary, rollback |
| Tenant | data leak | authz in retrieval/tool layer, isolation tests |
| Cost | token/API spike | caching, summarization, model routing, rate limits |

Interview line:

> "Failures must be observable, bounded, and safe."

---

## 7. Failure Mode Brush-Up

Use this when you have 5-10 minutes and want the failure-mode muscle memory.

### Default Answer Shape

```text
Failure -> Detection -> Containment -> Recovery -> Prevention
```

| Step | Say |
|---|---|
| Failure | "The likely failure is X at layer Y." |
| Detection | "I would detect it with metric/log/trace Z." |
| Containment | "The system should degrade safely by doing A." |
| Recovery | "Then retry/replay/rebuild through B." |
| Prevention | "Longer term, add test/eval/guardrail C." |

### Fast Drills

| Prompt | Strong Answer |
|---|---|
| Threat intel is down | Timeout, circuit breaker, cached/stale intel, partial triage, mark missing source. |
| Vector retrieval returns bad context | Tenant filter, source scores, rerank, cap low-quality examples, track analyst rejects. |
| Model emits invalid JSON | Schema validation, one repair retry, deterministic fallback, human review if still invalid. |
| Agent recommends unsafe action | Policy gate blocks, log recommendation, show as context only, require analyst approval. |
| Queue is falling behind | Consumer lag alert, autoscale, severity priority, backpressure, shed non-critical enrichment. |
| Duplicate alerts flood triage | Idempotency keys, dedupe table, grouping by detection/entity/window, analyst-visible cluster. |
| Parser/schema changes break normalization | Schema versioning, parser tests, DLQ, replay after fix, compatibility checks. |
| Cross-tenant evidence appears | Treat as severity-one incident: block result, audit retrieval filters, isolation tests, rotate affected cache/index if needed. |
| Cost spikes | Token budgets, summarization, cache, smaller-model routing, rate limits per tenant/workflow. |
| Model quality regresses | Offline eval catches it, shadow/canary rollout, rollback, compare by severity and tenant segment. |

### One-Minute Spoken Drill

> "For each component I would ask: how does it fail, how do we know, what is the safe degraded behavior, and how do we recover? For example, if enrichment is slow, I do not block triage indefinitely. I use timeouts, retries, circuit breakers, and cached data, then return a partial result with missing evidence explicitly marked. The same pattern applies across the system: queues degrade with backpressure and priority, retrieval degrades with stricter filters and reranking, LLM output degrades through schema validation and human fallback, and actions are always bounded by deterministic policy gates."

Memory hook:

> "Detect fast, degrade safely, recover by replay, prevent with evals and tests."

---

## 8. Tradeoff Bank

| Decision | Default | Switch If |
|---|---|---|
| Streaming vs batch | Streaming for triage | Reports/training can be batch |
| RAG vs fine-tuning | RAG first | Stable task + enough high-quality labels |
| Hybrid vs vector only | Hybrid | Vector-only is acceptable only for non-exact semantic search |
| pgvector vs dedicated vector DB | pgvector early | Dedicated service when scale/filter latency demands it |
| Big vs small model | Route by risk | Single model if simplicity matters more |
| Human review vs autonomy | HITL for high stakes | Low-risk, high-confidence, reversible action |
| Store summaries vs raw | Store both | Retention/cost forces tiering |
| Shared vs per-tenant infra | Shared with hard isolation | Regulated/large tenant needs dedicated stack |
| Schema-on-write vs schema-on-read | Schema-on-write for quality | Unknown log shapes need flexibility |

Pattern:

> "I would choose A under these assumptions because B costs us X. If the constraint changes to Y, I would switch."

---

## 9. Metrics to Name

| Category | Metrics |
|---|---|
| System | ingestion rate, queue lag, p95/p99 latency, enrichment timeout rate |
| Retrieval | hit rate, rerank quality, evidence usefulness, source quality |
| Model | latency, token cost, schema-valid rate, repair rate |
| SOC outcome | time-to-triage, analyst override rate, auto-close reversal, MTTR |
| Safety | policy blocks, unauthorized tool attempts, prompt-injection detections |
| Learning | eval precision/recall by severity/tenant, drift, canary win/loss |
| Trust | audit completeness, evidence click-through, analyst accept/edit/reject |

---

## 10. Scaling at 10x

| Bottleneck | Move |
|---|---|
| Ingestion throughput | partitioning, autoscaling, backpressure |
| Queue depth | more consumers, severity priority, shed non-critical work |
| Enrichment APIs | cache, batch, timeout, stale reads |
| Retrieval latency | shard indexes, precompute features, source filters |
| LLM cost/rate limit | summarize, cache, route smaller models, batch where possible |
| Storage volume | hot/cold tiering, retention policy, compression |
| Tenant isolation | authz tests, per-tenant filters, audit trails |

Scale numbers:

| Number | Recall |
|---|---|
| 1M events/day | about 11.5 events/sec |
| 1B events/day | about 11,500 events/sec |
| Same-DC RTT | about 0.5 ms |
| Cross-region RTT | about 30-100 ms |
| SSD read | about 100 us |
| Seconds | first-pass enrichment target |
| Minutes | analyst-ready triage target when evidence quality matters |

---

## 11. Safe Autonomy

| Risk | Action |
|---|---|
| Low risk + high confidence + reversible | Auto-tag, auto-summarize, maybe auto-close low-severity duplicate |
| Medium risk | Recommend + require analyst approval |
| High severity or irreversible | Human review always |
| Policy conflict | Block and log |
| Unknown confidence | Escalate |

Hard rule:

> "The LLM never directly executes side effects; tools and policy enforce action boundaries."

---

## 12. Things to Say

- "Let me clarify assumptions before designing."
- "I would structure this as an event-driven pipeline."
- "The LLM recommends; the platform enforces."
- "I optimize for signal per token."
- "Every decision should cite evidence IDs."
- "Tenant isolation is an auth/retrieval invariant, not a prompt instruction."
- "Partial enrichment is better than blocking triage."
- "I would launch in shadow mode before autonomous actions."
- "Feedback improves retrieval, evals, thresholds, and prompts."
- "We design for graceful degradation at scale."

---

## 13. Red Flags to Avoid

| Avoid | Replace With |
|---|---|
| Drawing before clarifying | Ask scale, latency, trust, autonomy |
| "Just use an LLM" | Show data flow, tools, policy, monitoring |
| LLM enforces permissions | Authz outside model |
| Vector-only retrieval | Hybrid retrieval for exact + semantic |
| Full autonomy too early | Risk-based autonomy + HITL |
| Only AI metrics | SOC/customer metrics too |
| "It depends" | Name the dependency and decision rule |
| Ignoring multi-tenancy | Tenant filters, authz, audit, tests |

---

## 14. Strong Close

Use this:

> "The design is a streaming, multi-tenant triage pipeline with enrichment and hybrid retrieval feeding an agent, but policy and side effects stay deterministic. The biggest risks are bad evidence, model drift, tenant isolation, and unsafe automation, so I would launch with shadow mode, audit trails, evals, canaries, and human review for high-stakes actions."

Then ask one:

- "At Panther, is the hardest bottleneck ingestion scale, retrieval quality, evals, or analyst workflow integration?"
- "How much autonomy do customers want today versus analyst-assist workflows?"
- "How do you approach cross-customer learning while preserving data boundaries?"

---

## See Also

- [[rounds/systems-design]] - full Round 4 guide
- [[rounds/round-4-interview-questions]] - spoken-answer drill bank
- [[concepts/systems-design-patterns]] - pattern reference
- [[concepts/soc-domain]] - SOC workflow context
- [[concepts/agentic-ai]] - RAG, feedback loops, and agent architecture
