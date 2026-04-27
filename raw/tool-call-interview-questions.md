1. Explain tool calling end-to-end.
    Answer:
    Tool calling is when an LLM emits a structured request (tool name + arguments) instead of text when it needs external data or actions. The runtime/orchestrator validates the request (schema + semantics), enforces authorization and policy (RBAC, tenant scope, safety), executes the tool, and returns the result to the model. The model then continues reasoning or produces a final answer. The model proposes; the application owns execution, permissions, and safety.

⸻

2. What specifically are you validating before executing a tool call?
    Answer:
    I validate at three levels:

* Structure: required fields, types, enums, JSON shape
* Semantics: valid IDs, realistic values, correct formats (e.g., valid IP), bounds (time ranges, limits)
* Policy: RBAC, tenant isolation, allowed tool/action, rate limits, read vs. write permissions
    Schema-valid is not enough; the handler must validate meaning and authorization.

⸻

3. How do you decide which tools are safe vs. require human approval?
    Answer:
    I evaluate tools along read vs. write and risk level. Read tools are generally safe but still require RBAC and tenant enforcement. Write tools are gated based on severity, confidence, and blast radius. Low-risk/high-confidence actions may be automated, but high-impact actions (e.g., closing alerts, blocking IPs) require policy gates or human approval. The model proposes actions; the runtime decides if they are allowed.

⸻

4. How do you defend against prompt injection from attacker-controlled logs?
    Answer:
    I treat all retrieved data as untrusted and never rely on prompts for enforcement. I enforce safety at the system level: treat tool output as data (not instructions), enforce RBAC and tenant isolation server-side, restrict tools via allowlists, gate write actions with policy/human approval, and log all actions. Even if the model is injected, it cannot execute unsafe actions because the runtime enforces boundaries.

⸻

5. How do you prevent infinite tool-calling loops?
    Answer:
    I enforce a max-iteration limit and track progress. If the agent repeats failed or irrelevant calls, the loop stops and falls back to a final response or human escalation. Tool handlers return structured errors so the model can recover, but retries are bounded. The agent can recover, but not loop indefinitely.

⸻

6. What tool calls happen before classification in alert triage?
    Answer:
    I fan out read-only enrichment calls such as: get_alert_details, query_logs, get_user_history, get_asset_context, search_threat_intel, and find_similar_alerts. These results are compiled into a compact evidence bundle with IDs/citations for the LLM to reason over.

⸻

7. How do you handle 50K logs with limited context?
    Answer:
    I first apply deterministic filters (tenant, detection ID, time range, entities). Then I use hybrid retrieval (keyword + vector search), followed by reranking for relevance and diversity. I cluster or aggregate repetitive events and summarize representative patterns while preserving evidence IDs for auditability. Summarize patterns, not evidence away.

⸻

8. When do you use parallel vs. sequential tool calls?
    Answer:
    Parallel calls are used for independent enrichment tasks (e.g., logs, user history, threat intel) to reduce latency. Sequential calls are used when one result informs the next (e.g., extract IP → then query threat intel). I also enforce concurrency limits and timeouts.

⸻

9. What should tools return to the model?
    Answer:
    Tools should return structured, compact, evidence-backed data. This includes stable fields, summaries, and evidence IDs. Avoid dumping large raw data; instead, preserve references so analysts can inspect underlying data when needed.

⸻

10. Design a schema for search_threat_intel.
    Answer:
    Include constrained, unambiguous inputs:

* indicator (string): the exact IP/domain/URL/hash
* indicator_type (enum): [“ip”, “domain”, “url”, “hash”]
* optional constraints like confidence thresholds

Use enums and validation to reduce ambiguity and hallucination. The description should clearly state when to use the tool.

⸻

11. What happens if the model hallucinates an invalid indicator?
    Answer:
    The handler rejects it during semantic validation (e.g., invalid IP format). It returns a structured error like invalid_indicator_format, allowing the agent to retry with corrected data or escalate. Schema validation checks structure; semantic validation catches impossible values.

⸻

12. How do you make write tools idempotent?
    Answer:
    I use idempotency keys or deterministic identifiers so repeated calls produce the same result instead of duplicates. For example, create_ticket can include a unique alert ID or hash so retries don’t create multiple tickets. I also check existing state before execution (e.g., if alert is already closed). This ensures safe retries and prevents duplicate side effects.
