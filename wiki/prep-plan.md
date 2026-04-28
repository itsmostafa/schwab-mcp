---
type: plan
last-updated: 2026-04-28
tags: [prep, systems-design]
---

# Interview Prep Plan — Systems Design

Covers the **Systems Design** round, Darwayne's "most revealing" interview.

---

## What This Round Tests

Per Darwayne, this one reveals how you think. Panther cares about:

- How you handle ambiguity: clarify before diving in.
- Whether every technical decision has an explicit tradeoff.
- Whether you proactively surface failure modes.
- Whether your process is systematic, not just pattern-matching to a memorized answer.
- Whether you connect architecture choices to SOC customer outcomes.

---

## The Anchor: Alert Triage System

Use this as the default mental model for Panther-style design prompts:

```text
Alert fires
  -> deterministic intake: schema validation, tenant routing, auth context
  -> enrichment fan-out: alert details, user history, asset context, threat intel, related logs
  -> retrieval/index lookups: exact IOC/rule/entity search plus similar historical cases
  -> ranking and summarization: dedupe, cluster, preserve evidence IDs
  -> decision service: classification, confidence, rationale, recommended next action
  -> policy gate: severity, confidence, allowed action scope, tenant policy
  -> analyst review / ticket / escalation / limited auto-close for high-confidence benign cases
  -> feedback capture: analyst accept/reject/modify, final disposition, metrics
```

Key narration beats:

- Where deterministic services run vs. where probabilistic model judgment is acceptable.
- Why exact search and semantic retrieval complement each other.
- Why the policy gate lives outside the model.
- What metrics prove the system is working.
- How analyst feedback improves future system behavior.

---

## Process To Use In The Interview

### 1. Clarify First

Ask before drawing:

- "What scale are we designing for: alerts/day, customers, analysts?"
- "What is the latency SLA: seconds, minutes, or batch?"
- "Is this greenfield or integrating with Panther's existing ingestion and detection stack?"
- "What is the one thing we cannot get wrong: accuracy, latency, cost, or auditability?"
- "What actions are in scope: recommendation only, ticket creation, alert closure, containment?"

### 2. Restate The Problem

Summarize what you heard and lock assumptions:

> "So I am designing a multi-tenant alert triage system that takes normalized alerts, enriches them, ranks supporting evidence, recommends a disposition, and routes uncertain or risky cases to analysts. I will assume analyst-facing triage can complete in minutes unless you want a stricter SLA."

### 3. Draw The High-Level Design

Start with components and data flow:

```text
Ingestion / alert stream
  -> Alert intake service
  -> Enrichment orchestrator
  -> Log/threat-intel/user/asset stores
  -> Retrieval/index layer
  -> Triage decision service
  -> Policy gate
  -> Case/ticket workflow
  -> Analyst feedback store
  -> Monitoring + audit log
```

### 4. Deep Dive On One Or Two Components

Good components to choose:

- Enrichment fan-out and concurrency controls.
- Retrieval/index layer and tenant-safe filtering.
- Policy gate and human-review workflow.
- Feedback loop and evaluation pipeline.
- Audit trail and observability.

### 5. Name Tradeoffs Throughout

Useful tradeoff pairs:

| Decision | Option A | Option B | When to pick A |
|---|---|---|---|
| Latency vs accuracy | Smaller/faster model | Larger/slower model | High alert volume, first-pass triage |
| Retrieval vs fine-tuning | Retrieval | Fine-tune | Data changes frequently; evidence must be inspectable |
| Vector DB vs pgvector | Dedicated vector DB | pgvector | Need independent scaling, managed ops, or strict latency/filtering |
| Streaming vs batch | Kafka/Flink-style streaming | Batch jobs | SLA is seconds/minutes and results update continuously |
| Human review vs autonomy | Analyst approval | Autonomous action | High-severity, low-confidence, or side-effecting actions |

### 6. Finish With Operations

Always cover:

- Metrics: time-to-triage, analyst override rate, escalation acceptance, false-negative rate, latency p95/p99, cost per alert.
- Safety: prompt-injection detections, policy-gate blocks, unauthorized tool attempts, high-severity human-review coverage.
- Reliability: retries, idempotency, circuit breakers, degraded mode when enrichment sources fail.
- Auditability: prompt/model/index versions, retrieved evidence IDs, tool calls, policy decisions, analyst actions.
- Rollback: staged rollout, canary by tenant, shadow mode, automatic rollback triggers.

---

## Scenarios To Practice

### 1. Real-Time Alert Triage System

Most likely prompt. Anchor on ingestion, enrichment, retrieval, decisioning, policy, workflow, feedback, audit.

### 2. Text-To-Search Over Security Logs

Key components:

- Natural-language query intake.
- Intent/entity extraction.
- Structured query generation with validation.
- Permission and tenant filtering.
- Query execution against the log store.
- Result ranking and explanation.
- Analyst correction feedback.

### 3. Detection Code Generation

Key components:

- Natural-language rule spec.
- Retrieval over existing detections and docs.
- Code generation.
- Static checks and unit tests.
- Sandbox execution against sample logs.
- Human review and version-controlled promotion.

### 4. Collective Intelligence For SOC Analysts

Key components:

- Capture analyst decisions and edits.
- Normalize feedback signals.
- Separate tenant-specific learning from shared product learning.
- Update eval sets, retrieval examples, and thresholds.
- A/B test or canary improvements.
- Roll back on quality regression.

### 5. Data Pipeline Feeding AI Agents

Key components:

- Log ingestion.
- Parsing and normalization.
- Schema-on-read vs schema-on-write.
- Storage for raw and normalized data.
- Exact and semantic indexing.
- Retention, privacy, and tenant isolation.

---

## Numbers To Keep Handy

- 1M events/day is about 11.5 events/second.
- 1B events/day is about 11,500 events/second.
- Memory is much faster than SSD; same-DC network calls are usually sub-millisecond to low-millisecond, but service fan-out dominates tail latency.
- Use "seconds for first-pass enrichment, minutes for analyst-ready triage" as a defensible assumption if no SLA is given.

---

## Practice Schedule

### Block A — Process Rehearsal

- Recite the six-step flow: clarify, restate, high-level, deep dive, tradeoffs, operations.
- Whiteboard the alert triage system from memory in under 5 minutes.
- For each component, name one failure mode and one metric.

### Block B — Timed Mock

Prompt:

> "Design a real-time alert triage system for 1B events/day, 100 customers, first-pass enrichment in seconds, and analyst-ready triage in minutes."

Run it for 30 minutes:

- 3 minutes clarify.
- 2 minutes restate.
- 8 minutes high-level design.
- 10 minutes deep dive.
- 5 minutes failure modes and metrics.
- 2 minutes close with tradeoffs and next steps.

### Block C — Non-Triage Mock

Pick one:

- Text-to-search over security logs.
- Detection code generation.
- Collective intelligence for SOC analysts.

Force yourself to use the same process without relying entirely on the triage anchor.

### Block D — Final Review

- Memorize 4 Panther architecture questions.
- Review the tradeoff table.
- Practice one concise close:

> "The design I chose optimizes for trusted analyst assistance over blind automation. I would start with recommendation and evidence packaging, prove quality with analyst feedback and safety metrics, then expand autonomy only where policy and data show it is safe."

---

## Questions To Ask Panther

- "Where do you draw the line today between autonomous action and analyst review?"
- "What signals have mattered most in evaluating SOC agent quality with real customers?"
- "How do you think about tenant-specific behavior versus shared product behavior?"
- "What constraints from Panther's existing ingestion pipeline most shape the agent architecture?"

---

## Day-Of Checklist

- [ ] Ask clarifying questions before drawing.
- [ ] State scale and SLA assumptions explicitly.
- [ ] Keep the design layered and easy to follow.
- [ ] Name the tradeoff behind each major decision.
- [ ] Surface at least three failure modes.
- [ ] Cover metrics, auditability, rollback, and human review.
- [ ] Connect decisions back to analyst trust and customer outcomes.

---

## See Also

- [[rounds/systems-design]] — detailed round guide
- [[rounds/round-4-interview-questions]] — spoken-answer drill bank
- [[rounds/round-4-cheat-sheet]] — final-scan cheat sheet
- [[concepts/systems-design-patterns]] — patterns library
- [[concepts/soc-domain]] — SOC domain context
