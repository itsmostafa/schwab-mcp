---
type: question-bank
last-updated: 2026-04-27
tags: [round-3, ai-integration, interview-questions, rag, soc]
---

# Round 3 Interview Questions - AI Integration

Use this as a spoken-answer drill sheet. The goal is not to memorize every word; it is to keep the same architecture, boundaries, metrics, and safety posture under pressure.

## 1. Walk me through a RAG pipeline for alert triage.

Strong answer:

"When an alert fires, I would treat the system as event-driven. The alert is published to a queue, then a deterministic pipeline handles intake: schema validation, tenant routing, and enrichment with asset, user, and threat-intel context.

For retrieval, I would use a hybrid approach:

- BM25 or structured search for exact IOCs like hashes, IPs, detection IDs, users, and hostnames
- Vector search for semantic similarity across past investigations, runbooks, threat intel, and analyst notes

Before retrieval, I pre-filter by tenant, time range, detection ID, severity, and access scope so relevance and RBAC are enforced before the model sees anything.

After retrieval, I optionally rerank results, dedupe or cluster repetitive events, and pass a compact, high-signal evidence set into the LLM.

The LLM generates a structured triage output, and a deterministic policy gate decides whether to auto-close, escalate, or send to an analyst."

Key points to hit:

- Event-driven intake
- Hybrid retrieval
- Tenant and RBAC filtering before retrieval
- Rerank, dedupe, summarize
- LLM recommends; policy gate enforces

## 2. What goes into the LLM, what comes out, and where is the policy gate?

Strong answer:

"I pass a structured prompt containing normalized alert fields, enriched context, top retrieved evidence with citation IDs, similar past investigations, tenant policy constraints, and a strict output schema.

The LLM returns structured JSON: classification, confidence, rationale, evidence IDs, and recommended action.

I use two gates. The pre-LLM gate enforces RBAC, sanitizes input, and filters data. The post-LLM gate validates schema, enforces permissions, checks confidence thresholds, and applies tenant policy.

The key principle is that the LLM can recommend, but deterministic systems enforce decisions."

Possible output schema:

```json
{
  "classification": "true_positive | false_positive | suspicious | unknown",
  "confidence": 0.0,
  "rationale": "short evidence-backed explanation",
  "evidence_ids": ["event-123", "case-456"],
  "recommended_action": "escalate | close | investigate | create_ticket",
  "needs_human_review": true
}
```

## 3. How do you evaluate this system in production?

Strong answer:

"I evaluate it at three levels: offline quality, online product impact, and safety.

Offline, I track precision and recall by severity, false-negative rate, schema-valid output rate, citation faithfulness, and agreement with analyst decisions.

Online, I track analyst override rate, escalation acceptance rate, auto-close reversal rate, time-to-triage, latency, and cost per alert.

For safety, I track prompt-injection detection, unauthorized tool attempts, policy-gate blocks, and high-severity human-review coverage.

I log every prompt, retrieval result, model response, policy decision, and tool call so the full chain is auditable. Improvements should come from failure-mode analysis, not just aggregate metrics."

Key production detail:

- Tie each metric to prompt version, model version, retrieval index version, and tenant-safe evaluation slices.

## 4. RAG vs. fine-tuning: when and why?

Strong answer:

"I would start with RAG because security data changes constantly, we need citations and auditability, and tenant/RBAC filtering must happen before retrieval and prompt assembly.

Fine-tuning is slower to update and less inspectable as a source of factual knowledge.

That said, I would use fine-tuning for stable behavior: consistent output formatting, classification behavior, domain phrasing, and repeated decision patterns once we have high-quality labeled data.

So my shorthand is: RAG for dynamic knowledge and evidence grounding; fine-tuning for stable behavior."

Watch out:

- Do not imply fine-tuning replaces retrieval for fresh threat intel, runbooks, or tenant-specific history.

## 5. How do you defend against prompt injection in SOC logs?

Strong answer:

"I treat all logs as attacker-controlled input.

The core defenses are structural separation between system instructions and retrieved data, quoting or sanitizing untrusted log content, pre-filtering data with RBAC and tenant boundaries, allowlisting tool access, and enforcing all permissions server-side.

I do not rely on detecting injection alone. Even if the model is tricked, it cannot take unauthorized actions because enforcement is outside the LLM."

Good line:

> "Prompt injection is not just a chatbot issue here; the attacker may control the log line the SOC agent reads."

## 6. Why not just put 50K events into a 200K-token context?

Strong answer:

"Because more context does not automatically mean better reasoning.

Large context increases latency, cost, noise, prompt-injection surface area, and the risk that key signals get lost in irrelevant data.

Instead, I would filter deterministically by tenant, time window, entity, severity, detection ID, and data source. Then I would dedupe and cluster repetitive events, summarize clusters, select representative examples, rerank for relevance and diversity, and preserve raw event IDs for citation and drill-down.

The LLM should see a high-signal evidence package, not act as the database."

Key phrase:

> "The prompt should contain enough evidence to reason, not enough raw data to become the database."

## 7. We have 12 monolithic prompts drifting independently. How do you fix it?

Strong answer:

"I would move to a modular prompt architecture.

There would be a shared base prompt for role, safety rules, citation expectations, and output schema. Then task modules for triage, chat, detection generation, and text-to-search. Then a tenant policy module for allowed actions and thresholds. Finally, an evidence module with alert context and retrieved sources.

I would version prompts like code, add golden test cases and evals, log prompt versions in production, and use staged rollout or A/B testing.

The goal is to prevent shared logic from drifting while keeping task-specific flexibility."

Practical details:

- Require schema validation for model outputs.
- Diff-review prompt changes that affect safety or action policy.
- Attach eval results to prompt/model/index versions.

## 8. How does the system improve over time?

Strong answer:

"The feedback loop is driven by analyst decisions.

I would capture accept, reject, edit, final disposition, severity changes, reopen rate, and time-to-triage. I would store that alongside prompt version, retrieved evidence IDs, model output, confidence, and policy decisions.

Then I would use feedback to update the retrieval index with confirmed cases, improve prompts and thresholds, create eval datasets, and identify recurring failure modes.

I would not blindly train on all data. Analyst feedback can be noisy, so I would curate high-quality signals and deploy changes through evals, shadow mode, canaries, and rollback."

Key nuance:

- Separate tenant-specific learning from any cross-customer aggregation.
- Keep privacy, retention, and access boundaries explicit.

## See Also

- [[rounds/ai-integration]] - full Round 3 guide
- [[rounds/round-3-cheat-sheet]] - final-scan AI Integration cheat sheet
- [[concepts/tool-calling]] - tool execution boundaries and guardrails
- [[concepts/agentic-ai]] - deeper RAG and agent architecture notes
