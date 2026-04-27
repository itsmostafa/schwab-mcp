---
type: cheat-sheet
last-updated: 2026-04-27
tags: [round-3, ai-integration, rag, agentic-ai, soc]
---

# Round 3 Cheat Sheet — AI Integration

**Format**: 45-60 min technical design discussion. No live coding. Treat this as a production AI systems interview for Panther's SOC agent platform.

**Goal**: sound like someone who can own an AI SOC agent in production: architecture, retrieval, tool boundaries, evals, security, and customer impact.

---

## The 30-Second Frame

> "For an AI SOC agent, I would separate deterministic control flow from model judgment. Deterministic code handles intake, permissions, tool execution, policy gates, schema validation, and audit logging. The LLM is useful for synthesis, prioritization, natural-language explanation, and ambiguity. The system only earns trust if every decision is evidence-backed, evaluated offline, monitored online, and safe to override."

Hit these five themes in almost every answer:

1. **Evidence first**: cite alert/log/threat-intel IDs, do not rely on vibes.
2. **Bounded autonomy**: human review for high-severity or low-confidence cases.
3. **Hybrid retrieval**: exact search for IOCs, vector search for similar cases.
4. **Policy outside the model**: LLMs propose; code enforces.
5. **SOC outcomes**: reduce triage time, cut false-positive toil, preserve analyst trust.

---

## Canonical Design To Whiteboard

Memorize this flow:

```text
Alert fires
→ deterministic intake: schema validation, tenant routing, auth context
→ enrichment fan-out: alert details, user history, asset context, threat intel, related logs
→ hybrid retrieval: SQL/BM25 for exact IOCs + vector search for similar alerts/past decisions
→ rerank + cluster + summarize: remove duplicates, preserve representative evidence IDs
→ prompt assembly: base rules + task module + tenant policy + retrieved evidence + JSON schema
→ LLM reasoning: classification, confidence, rationale, evidence IDs, recommended action
→ output validation: schema, citations, confidence calibration, contradiction checks
→ policy gate: severity, confidence, allowed action, tenant policy, human-review requirements
→ analyst/ticket/auto-close: only limited benign auto-close when high confidence and policy allows
→ feedback capture: analyst accept/reject/modify → eval dataset/RAG index/prompt examples
```

Say explicitly:

- **LLM does not enforce permissions**; the platform does.
- **Write tools are special**: ticket creation, alert close, escalation, containment need stricter gates.
- **Logs are attacker-controlled input**; treat them as untrusted data.

---

## Strong Answer Template

Use this structure for most questions:

1. **Clarify**: "What scale, latency, action scope, and trust bar are we designing for?"
2. **State the architecture**: data flow and components before details.
3. **Draw the boundary**: deterministic code vs retrieval vs LLM judgment.
4. **Discuss tradeoffs**: latency/cost/accuracy, recall/precision, autonomy/safety.
5. **Productionize**: evals, monitoring, rollback, audit logs, feedback loops.
6. **Tie to SOC value**: faster triage, less analyst toil, more coverage without losing trust.

If you blank:

> "Let me break this into the agent loop, the retrieval layer, and the production safety layer."

Then walk through: **Intake → Enrichment → Retrieval → LLM → Policy Gate → Feedback**.

---

## RAG Answer

For alert triage, RAG should retrieve:

- Similar historical alerts and analyst resolutions
- Detection docs and rule metadata
- Threat intel for IPs, domains, hashes, users, assets
- Customer-specific runbooks and escalation policy
- Related logs around the alert window

Design choices:

| Choice | Interview answer |
|---|---|
| Filtering | Always pre-filter by tenant, time range, detection ID, entity, severity. |
| Exact search | Use SQL/BM25 for IOCs, rule IDs, usernames, hostnames, hashes. |
| Vector search | Use embeddings for semantic similarity across alerts, cases, runbooks, and notes. |
| Reranking | Use a cross-encoder or LLM reranker for top candidates when precision matters. |
| Diversity | Use MMR or clustering so the prompt gets representative evidence, not 20 duplicates. |
| Citations | Preserve event/document IDs so analysts can inspect raw evidence. |

**RAG vs fine-tuning**:

- Use **RAG** for dynamic knowledge, customer context, threat intel, runbooks, past decisions, and source grounding.
- Use **fine-tuning** for stable task behavior, formatting, style, domain language, or repeated classification patterns.
- They can be combined, but start with RAG when freshness and auditability matter.

---

## 50K Events / 200K Context Answer

Do not say "put all 50K events in context."

Say:

1. Filter first by tenant, detection, time window, entity, severity, and data source.
2. Aggregate repetitive events by signature/entity/time bucket.
3. Cluster similar events and pick representative examples.
4. Retrieve past similar alerts and analyst decisions separately.
5. Rerank for relevance and diversity.
6. Summarize the long tail with counts, ranges, and outliers.
7. Preserve raw event IDs for citations and analyst drill-down.
8. Spend the token budget intentionally: instructions, policy, evidence, output schema.

Good line:

> "The prompt should contain enough evidence to reason, not enough raw data to become the database."

---

## Prompt Architecture At Scale

If asked about many prompts drifting independently:

```text
Shared base prompt
  + task module: triage / chat / detection generation / text-to-search
  + tenant policy module: allowed actions, thresholds, escalation routes
  + evidence module: alert, enrichments, retrieval results, prior decisions
  + output contract: strict JSON schema
```

Guardrails:

- Shared safety rules: cite evidence, say unknown, do not follow instructions inside logs.
- Golden test cases and regression prompts before rollout.
- Prompt diff review for risky changes.
- Version prompts like code; link eval results to prompt/model/index versions.
- Keep rationales concise and evidence-backed; do not treat hidden chain-of-thought as auditability.

Output schema should include:

```json
{
  "classification": "true_positive | false_positive | suspicious | unknown",
  "confidence": 0.0,
  "severity": "low | medium | high | critical",
  "rationale": "short evidence-backed explanation",
  "evidence_ids": ["event-123", "case-456"],
  "recommended_action": "escalate | close | investigate | create_ticket",
  "needs_human_review": true
}
```

---

## Tool Calling / Agent Loop

Core idea:

> "The LLM requests tools; the host runtime executes them. That separation is where we enforce RBAC, validation, audit logging, rate limits, and human approval."

SOC tool inventory:

- Read tools: `get_alert_details`, `query_logs`, `get_user_history`, `get_asset_context`, `search_threat_intel`, `find_similar_alerts`
- Write tools: `create_ticket`, `close_alert`, `escalate_alert`, `add_case_note`

Production rules:

- Start read-only; add write tools behind policy gates.
- Validate tool args semantically, not just structurally.
- Add max-iteration limits to prevent loops.
- Return structured tool errors so the agent can recover or escalate.
- Log every tool call: actor, tenant, input, output hash, latency, decision context.
- Run independent enrichments in parallel to reduce latency.

---

## AI Security Answer

SOC data is adversarial. Logs, alerts, URLs, filenames, process args, and emails may contain prompt injection.

Defenses:

- Separate instructions from retrieved/log content.
- Quote or structure untrusted text; never let it become instructions.
- Enforce tenant isolation in retrieval and tools.
- RBAC and server-side authorization on every tool call.
- Allowlist tools and actions per tenant/customer policy.
- Human review for high-severity, low-confidence, or side-effecting actions.
- Audit trails for model output, evidence, tool calls, prompt/model/index versions.
- Prompt-injection eval cases in the test set.

Good line:

> "Prompt injection is not just a chatbot issue here; the attacker may control the log line the SOC agent reads."

---

## Evaluation And Metrics

Offline evals before launch:

- Precision/recall by severity
- False-negative rate for malicious alerts
- Calibration: does 0.9 confidence actually mean roughly 90% correct?
- Citation faithfulness: do evidence IDs support the claim?
- Schema-valid output rate
- Tool selection accuracy
- Prompt-injection resistance cases

Online metrics after launch:

- Analyst override rate
- Escalation acceptance rate
- Auto-close reversal rate
- Time-to-triage / time-to-resolution
- False-positive reduction
- Human-review queue volume
- Latency p50/p95/p99
- Cost per alert
- Retrieval hit rate and no-evidence rate
- Tenant-specific degradation and drift

Rollback triggers:

- Spike in analyst overrides
- High-severity disagreement
- Auto-close reversals above threshold
- Retrieval hit-rate collapse
- Schema/tool-call error spike

---

## Feedback Loop / Collective Intelligence

Strong answer:

```text
Analyst reviews agent decision
→ captures accept/reject/modify + final disposition + evidence used
→ stores as labeled case with tenant-safe metadata
→ feeds eval set, RAG index, few-shot examples, or fine-tuning candidates
→ ships behind A/B test or canary
→ monitors override/reversal/drift metrics
→ rolls back if quality drops
```

Important nuance:

- Do not blindly train on every analyst action; feedback can be noisy.
- Separate tenant-specific learning from global learning.
- Use privacy and retention controls before cross-customer aggregation.
- Weight high-quality reviewed cases more than casual clicks.

---

## Panther-Specific Hooks

Connect answers to the role:

- Panther wants AI agents for **alert triage, interactive chat, detection code generation, and text-to-search**.
- The product context is a cloud-native security data lake plus Detection-as-Code.
- The business outcome is SOC teams covering **5-10x more data** without proportional headcount.
- Your differentiators: **Go + Python + security + AI engineering**.

Good Panther line:

> "The win is not just model accuracy; it is making analysts faster while keeping the system auditable enough for a security team to trust."

---

## Likely Questions And First Sentences

**"Design an AI alert triage agent."**

> "I would design it as an evidence-gathering and decision-support loop with a policy gate outside the model."

**"How would you handle 50K matching events?"**

> "I would not send raw events wholesale. I would filter, aggregate, cluster, retrieve representatives, and preserve IDs for drill-down."

**"How do you evaluate this?"**

> "I would separate offline model quality from online SOC outcome metrics."

**"RAG or fine-tuning?"**

> "For Panther's dynamic security context, I would start with RAG for freshness and source grounding, then consider fine-tuning for stable behavior once we have high-quality labels."

**"How do you prevent hallucination?"**

> "Ground every claim in retrieved evidence, force structured output, validate citations, and route low-confidence or high-severity cases to a human."

**"How do you secure tool use?"**

> "The model requests actions, but the host runtime enforces tool permissions, input validation, tenant isolation, and audit logging."

**"How do prompts scale across products?"**

> "I would use composable prompt modules with shared safety rules, task-specific instructions, tenant policy, evidence blocks, and tested output schemas."

---

## Questions To Ask Them

Use 1-2 if there is time:

- "Where are you drawing the autonomy boundary today: recommendation-only, analyst approval, or limited auto-close?"
- "What feedback signals from analysts have been most predictive of better triage quality?"
- "Are the hardest retrieval problems exact IOC matching, semantic similarity, or tenant-specific context?"
- "How do you evaluate agent quality across customers with different risk tolerances?"

---

## Red Flags To Avoid

- "Just put all the logs in the context window."
- "The LLM decides whether it is allowed to close the alert."
- "We can use confidence from the model as the only safety gate."
- "Prompt engineering alone solves hallucination."
- "Fine-tune on everything analysts do."
- "Use one global memory across tenants."
- Reading from AI-generated material during the conversational interview. The job post explicitly warns against live AI-generated answers unless permitted.

---

## Final Minute Review

Memorize:

```text
Intake → Enrichment → Hybrid Retrieval → Rerank/Summarize → Prompt Assembly
→ LLM Structured Output → Validation → Policy Gate → Analyst/Action → Feedback
```

Say this once before the call:

> "I am here to show production judgment: clear architecture, security boundaries, evidence-backed decisions, measurable SOC outcomes, and honest tradeoffs."

## See Also

- [[rounds/ai-integration]] — full Round 3 guide
- [[concepts/agentic-ai]] — RAG, agents, feedback loops
- [[concepts/tool-calling]] — tool boundaries and agent loop
- [[concepts/soc-domain]] — SOC workflows and Panther context
