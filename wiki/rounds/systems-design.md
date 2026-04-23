# Systems Design Round

Per Darwayne: **"The most revealing one IMO."** They care about how you approach the problem, recognize it's open-ended, and speak to your decisions.

## What They Evaluate

- Do you recognize the problem is open-ended and clarify scope?
- How do you break down the problem?
- Can you speak to trade-offs in your choices?
- Do you think about scale, failure modes, and operational concerns?
- Do you consider the security domain specifically?

## Approach Framework

### 1. Clarify First (2-3 min)
Ask before drawing anything:
- "What scale are we designing for? How many alerts/day? How many analysts?"
- "What's the latency requirement? Real-time vs batch?"
- "Is this a greenfield design or integrating with existing Panther infrastructure?"
- "What's most important to get right — accuracy, latency, or cost?"

### 2. Define the Problem (2 min)
Restate what you're designing in your own words. Get agreement.

### 3. High-Level Design (5-10 min)
Draw the major components first. Don't go deep yet.
- Data flows
- Major services/components
- Where data is stored
- Integration points

### 4. Deep Dive on Key Components (10-15 min)
Pick the 1-2 most interesting/important components and go deep.
- Schema design
- API contracts
- Algorithms / data structures
- Failure modes and mitigations

### 5. Address Trade-offs (throughout)
For every decision: say why you chose it, and what you gave up.

### 6. Operational Concerns (5 min)
- Monitoring: what metrics, what alerts
- Failure modes: what breaks, how to detect, how to recover
- Scalability: what's the bottleneck, how to scale it

---

## Likely Panther System Design Prompts

These are common for security + AI platforms. Prepare answers for at least 2-3:

### "Design a real-time alert triage system"
Key components: ingestion pipeline, enrichment service, embedding + retrieval, agent orchestrator, analyst feedback loop, audit log.

### "Design a text-to-search system for security logs"
Key components: query understanding (LLM), embedding search over log index, structured query rewrite, result ranking.

### "Design a detection code generation system"
Key components: spec intake (natural language), RAG over existing detection rules, code generation (LLM), testing/validation, human review workflow.

### "Design a collective intelligence system for SOC analysts"
Key components: signal capture (analyst decisions), feedback aggregation, embedding updates, A/B testing of agent improvements, rollback mechanism.

### "Design the data pipeline that feeds your AI agents"
Key components: log ingestion, normalization/parsing, schema-on-read vs schema-on-write, indexing for both structured queries and semantic search, retention policies.

---

## Key Trade-offs to Know

| Decision | Option A | Option B | When to use A |
|----------|----------|----------|---------------|
| Latency vs accuracy | Smaller, faster model | Larger, slower model | Alert volume is high, triage just needs first-pass |
| RAG vs fine-tuning | RAG | Fine-tune | Data changes frequently; interpretability matters |
| Vector DB vs pgvector | Dedicated (Pinecone, Qdrant) | pgvector | Scale > 10M vectors or need ANN performance |
| Real-time vs batch | Streaming (Kafka) | Batch (Spark) | SLA < 30s for alert triage |
| Human-in-the-loop | Async review | Fully autonomous | High-stakes actions (blocking IPs, closing tickets) |

---

## See Also

- [[concepts/systems-design-patterns]] — patterns library
- [[concepts/soc-domain]] — SOC domain context for designing around the right problems
