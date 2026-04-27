# Round 3: AI Integrations Questions

1. Walk me through a RAG pipeline for alert triage

Answer:

“When an alert fires, I’d treat the system as event-driven. The alert gets published to a queue, then a deterministic pipeline handles intake: schema validation, tenant routing, and enrichment with asset, user, and threat intel context.

For retrieval, I’d use a hybrid approach:

* BM25 or structured search for exact IOCs like hashes, IPs, detection IDs
* Vector search for semantic similarity across past investigations, runbooks, and threat intel

Before retrieval, I’d pre-filter by tenant, time range, detection ID, severity, and access scope to ensure relevance and RBAC compliance.

After retrieval, I’d optionally rerank results, dedupe or cluster repetitive events, and pass a compact, high-signal evidence set into the LLM.

The LLM generates a structured triage output, and a deterministic policy gate decides whether to auto-close, escalate, or send to an analyst.”

⸻

2. What goes into the LLM, and what comes out? Where is the policy gate?

Answer:

“I pass a structured prompt containing:

* normalized alert fields
* enriched context (asset/user/threat intel)
* top retrieved evidence with citation IDs
* similar past investigations
* tenant policy constraints
* strict output schema

The LLM returns structured JSON:

* classification
* confidence
* rationale
* evidence IDs
* recommended action

I use two gates:

* pre-LLM: enforce RBAC, sanitize input, filter data
* post-LLM: validate schema, enforce permissions, apply confidence thresholds

The key principle is: the LLM can recommend, but deterministic systems enforce decisions.”

⸻

3. How do you evaluate this system in production?

Answer:

“I evaluate at three levels:

Offline:

* precision/recall by severity
* false-negative rate
* schema-valid output rate
* agreement with analyst decisions

Online:

* analyst override rate
* auto-close reversal rate
* time-to-triage
* latency and cost

Safety:

* prompt-injection detection rate
* unauthorized tool attempts
* policy-gate blocks

I log every prompt, retrieval, response, and decision so the full chain is auditable. Improvements come from analyzing failure modes, not just aggregate metrics.”

⸻

4. RAG vs fine-tuning — when and why?

Answer:

“I’d start with RAG because:

* security data changes constantly
* we need citations and auditability
* we must enforce tenant/RBAC filtering before retrieval

Fine-tuning is slower to update and less interpretable.

That said, I’d use fine-tuning for:

* consistent output formatting
* classification behavior
* repeated decision patterns once we have labeled data

So: RAG for dynamic knowledge, fine-tuning for stable behavior.”

⸻

5. How do you defend against prompt injection in SOC logs?

Answer:

“I treat all logs as attacker-controlled.

Key defenses:

* structural separation between system instructions and retrieved data
* quote or sanitize untrusted log content
* pre-filter data with RBAC and tenant boundaries
* allowlist tool access
* enforce all permissions server-side

I don’t rely on detecting injection alone. Even if the model is tricked, it cannot take unauthorized actions because enforcement is outside the LLM.”

⸻

6. Why not just put 50K events into a 200K-token context?

Answer:

“Because more context ≠ better performance.

Large context increases:

* latency and cost
* noise and irrelevant data
* prompt-injection surface area
* risk of missing key signals (‘context rot’)

Instead:

* filter deterministically (tenant, time, entity, severity)
* dedupe and cluster repetitive events
* summarize clusters
* select representative examples with citations

The LLM sees a high-signal summary, not raw logs.”

⸻

7. We have 12 monolithic prompts drifting — how do you fix it?

Answer:

“I’d move to a modular prompt architecture:

* Base prompt: role, safety rules, output schema
* Task modules: triage, chat, detection generation, etc.
* Tenant policy module: allowed actions, thresholds
* Evidence module: retrieved context

Then I’d:

* version prompts like code
* add golden test cases and evals
* log prompt versions in production
* use staged rollout / A/B testing

The goal is to prevent shared logic from drifting while keeping task-specific flexibility.”

⸻

8. How does the system improve over time (feedback loop)?

Answer:

“The feedback loop is driven by analyst decisions:

Capture:

* accept / reject / edit
* final disposition
* severity changes
* reopen rate

Log everything:

* prompt version
* retrieved evidence
* model output
* confidence
* policy decisions

Use feedback to:

* update retrieval index (better similar cases)
* improve prompts and thresholds
* create eval datasets
* identify failure modes

I wouldn’t blindly train on all data — I’d curate high-quality signals and deploy changes through evals and staged rollout.”
