1. What clarifying questions would you ask before designing the system?
    “Before I start, I want to clarify a few things:
    * What is the scale (alerts per day, number of tenants/analysts)?
    * What are the latency requirements—real-time seconds vs batch minutes?
    * Is this greenfield or integrating with existing infrastructure?
    * What’s the most critical constraint—accuracy, latency, or cost?
    * What enrichment sources and internal services are available, and which are reliable enough for the critical path?
        I’ll assume an event-driven system with real-time triage requirements unless told otherwise.”
2. High-level architecture for a real-time alert triage system
    “At a high level:
    * Ingestion pipeline receives alerts (Kafka or similar queue)
    * Deterministic intake: validation, schema normalization, tenant routing
    * Enrichment fan-out: parallel calls to identity, asset inventory, threat intel, logs, and history
    * Retrieval layer: hybrid BM25/SQL + vector search
    * Rerank + summarization to compress evidence
    * Prompt assembly: base + task + tenant policy + evidence
    * LLM produces structured output (classification, confidence, rationale)
    * Policy gate enforces thresholds and permissions
    * Action layer: auto-close, escalate, or analyst review
    * Feedback loop + audit logging
        Data flows sequentially but enrichment/retrieval are parallelized for latency.”
3. What is enrichment?
    “Enrichment adds context to raw alerts so they’re actionable. It includes:
    * User context (role, behavior history)
    * Asset context (criticality, ownership)
    * Threat intel (IP/domain reputation)
    * Historical alerts and decisions
    * Related logs/events
        It’s implemented as an async fan-out to multiple services, with timeouts and partial results handling.”
4. Where is the boundary between deterministic code and the LLM?
    “Deterministic code handles validation, routing, enrichment, retrieval, deduplication, and prompt construction.
    The LLM is responsible for classification, summarization, and reasoning over curated evidence.
    The LLM must never control permissions, policy enforcement, or direct write actions—those are handled by a policy gate outside the model.”
5. Why use hybrid retrieval (BM25 + vector)?
    “BM25/SQL is high-precision for exact matches like IOCs, hashes, and IDs.
    Vector search is high-recall for semantic similarity like behavioral patterns.
    Using only BM25 misses novel patterns; only vector risks noise and missing exact signals.
    So I run both in parallel, merge results, optionally rerank, and pass a curated set to the LLM.”
6. How do you select what goes into the LLM prompt?
    “I optimize for information density per token:
    * Pre-filter by tenant, time range, detection type
    * Deduplicate and cluster repetitive events
    * Select representative samples
    * Rank by relevance/confidence
    * Optionally rerank with a cross-encoder
    * Summarize raw logs into structured evidence
    * Include IDs for traceability
        The goal is high-signal, non-redundant context within token limits.”
7. What are key failure modes and mitigations?
    “Examples:
    * Retrieval failure → monitor hit-rate, use hybrid fallback
    * Enrichment failure → timeouts, partial results, degrade gracefully
    * Hallucination → structured outputs, evidence grounding, monitor override rate
    * Unsafe actions → policy gate, human-in-the-loop, RBAC
    * Data drift → monitor distribution shifts, update RAG/index, continuous evals
        I design failures to be observable, bounded, and safe.”
8. How do you design the feedback loop?
    “I capture structured analyst interactions: accept/reject, edits, escalation, reopen events, time-to-triage.
    Store alongside alert, evidence IDs, prompt/model version.
    Use it to:
    * Update RAG index with confirmed examples
    * Build eval datasets from overrides/reopens
    * Tune prompts/models via offline testing
    * Adjust policy thresholds (e.g., auto-close)
    * Monitor metrics like override rate and reversal rate
        The loop improves retrieval, evaluation, and decision quality—not just prompts.”
