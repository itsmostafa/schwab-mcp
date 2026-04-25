---
type: plan
last-updated: 2026-04-25
tags: [prep, ai-integration, systems-design]
---

# Interview Prep Plan — Rounds 3 & 4

Covers the **AI Integration** (Round 3, next up) and **Systems Design** (Round 4, Darwayne's "most revealing") interviews.

---

## What Each Round Actually Tests

### AI Integration (Round 3) — 45-60 min, no coding
You're being evaluated as a senior engineer who can reason about production AI systems — not just someone who knows the vocabulary. The bar is: could you own this system at Panther on day one?

Key signals they're watching for:
- Layered agent architecture (not just "call the LLM")
- RAG design that handles security-specific constraints (noisy data, token limits, adversarial inputs)
- Production-grade thinking: evals, monitoring, rollback, feedback loops
- Safety-first instincts: prompt injection, RBAC, human-in-the-loop for high-stakes actions
- SOC domain fluency: you speak analyst, not just engineer

### Systems Design (Round 4) — 60 min, no coding
Per Darwayne, this one reveals how you think. They care about:
- How you handle ambiguity — do you clarify before diving in?
- Can you articulate the tradeoff behind every decision?
- Do you proactively surface failure modes?
- Is your process systematic, not just pattern-matching to a memorized answer?

---

## The Anchor: Alert Triage Agent

Both rounds can be anchored to a single canonical design. Learn this cold — it covers ~70% of the material for both interviews.

```
Alert fires
  → deterministic intake (schema validation, tenant routing)
  → enrichment fan-out (user history, asset context, threat intel, related logs)
  → hybrid retrieval (SQL/BM25 for IOCs, vector search for similar alerts + past decisions)
  → rerank + summarize + token budget
  → prompt assembly (shared base + task module + tenant policy + evidence)
  → LLM → structured JSON output (classification, confidence, rationale, evidence IDs)
  → policy gate (confidence threshold, severity, allowed action scope)
  → analyst review / auto-close / escalation / ticket
  → feedback capture → eval dataset → model/index update
```

Key narration beats to hit:
- Where deterministic code runs vs where the LLM runs (be explicit)
- Why hybrid retrieval (BM25 catches exact IOC matches; vector catches semantic similarity)
- Why the policy gate lives outside the LLM (LLMs cannot enforce permissions)
- What metrics prove this system is working (override rate, auto-close reversal rate)
- How it improves over time (analyst decisions → RAG index or fine-tune candidates)

---

## Round 3 Prep: AI Integration

### High-value topics (know these well)

**1. RAG pipeline design**
- Chunking strategies: fixed-size+overlap (simple) vs semantic/hierarchical (better recall for docs)
- Embedding models: OpenAI text-embedding-3, BGE; domain-specific matters for security data
- Hybrid retrieval: BM25 for exact keywords (IOCs, detection IDs) + vector for semantic similarity
- Reranking: cross-encoder after initial retrieval improves precision, costs latency
- Filtering: always pre-filter by tenant, time range, severity before vector search

**2. The "50K events / 200K context" answer**
Don't say "put it all in context." Walk through:
1. Structured filtering first (tenant, detection ID, time range, entity)
2. Cluster and summarize repetitive events (dedup by signature)
3. Retrieve representative examples, not exhaustive hits
4. Keep citation IDs — analyst must be able to inspect raw events
5. Budget tokens: base prompt + task + retrieved evidence + output schema

**3. Prompt architecture at scale**
For "12 monolithic prompt strings that drift independently":
- Shared base: role, safety constraints, output schema
- Task module: triage vs chat vs detection generation vs text-to-search
- Tenant policy module: per-customer allowed actions, severity thresholds
- Evidence module: alert + enrichment + past decisions + threat intel
- Tests: golden prompts, regression cases, offline eval before rollout

**4. Production metrics — name both AI and SOC metrics**
- Offline (pre-launch): precision/recall by severity, false-negative rate, schema-valid output rate
- Online (post-launch): analyst override rate, auto-close reversal rate, time-to-triage, p99 latency
- Safety: unauthorized tool attempt rate, prompt-injection detection rate, policy-gate blocks
- Drift signals: classification distribution shift, retrieval hit-rate drop, tenant-specific degradation

**5. AI security answer (prompt injection)**
SOC logs are attacker-controlled. Treat raw event text as untrusted:
- Keep system instructions separate from retrieved/log content (structural separation)
- Structured tool APIs with RBAC + server-side authorization
- Quote or sanitize untrusted text in prompts
- Allowlists for tool calls; require policy gate for any write action
- Log every tool call and model decision for audit

**6. Your personal system story (3-5 min, practice out loud)**
Structure:
1. Problem and user pain
2. Architecture and key data flow
3. Two tradeoffs you made
4. How you evaluated quality
5. What failed or surprised you
6. What you'd change now

### Likely questions to practice
- "Walk me through how you'd build a RAG pipeline for alert triage."
- "How would you design a feedback loop so the triage agent improves over time?"
- "RAG vs fine-tuning — when would you use each?"
- "How do you evaluate an AI agent in a SOC use case?"
- "Walk me through an agentic system you've built."
- "How do you handle hallucination where accuracy is critical?"

### Answer shape (use this template every time)
1. Clarify scope (scale, latency, action boundaries, customer trust constraints)
2. State the architecture — components + data flow
3. Call out deterministic vs LLM boundaries
4. Discuss tradeoffs (latency/cost/accuracy, autonomy/safety)
5. Make it production-grade (evals, monitoring, rollback, audit)
6. Tie to SOC value (analyst toil, triage speed, trust)

---

## Round 4 Prep: Systems Design

### Process is the product — internalize this sequence

**Step 1 — Clarify (2-3 min, always)**
Ask before drawing anything:
- "How many alerts/day? How many analysts?"
- "What's the latency SLA — real-time or batch?"
- "Greenfield or integrating with existing Panther infra?"
- "What's the one thing we can't get wrong — accuracy, latency, or cost?"

**Step 2 — Restate the problem in your own words (1-2 min)**
Get buy-in. Shows you understood before you started.

**Step 3 — High-level design (5-10 min)**
Major components only. Data flows. Where data lives. Integration points. Don't go deep yet.

**Step 4 — Deep dive on 1-2 key components (10-15 min)**
Pick the most interesting. Schema, API contract, algorithms, failure modes.

**Step 5 — Tradeoffs throughout**
Every decision: state what you chose and what you gave up. Never just "it depends."

**Step 6 — Operational concerns (5 min)**
Monitoring, failure detection, recovery, scalability bottlenecks.

### Design scenarios to prepare (have answers for all 5)

**1. Real-time alert triage system** — most likely
Components: ingestion pipeline → enrichment fan-out → embedding + retrieval → agent orchestrator → feedback loop → audit log

**2. Text-to-search over security logs**
Components: LLM query understanding → embedding search over log index → structured query rewrite → result ranking + explanation

**3. Detection code generation**
Components: natural language spec intake → RAG over existing detection rules → LLM code generation → automated testing + validation → human review queue

**4. Collective intelligence for SOC analysts**
Components: signal capture (analyst accept/reject/modify) → feedback aggregation → RAG index update → A/B test new agent version → rollback mechanism

**5. Data pipeline feeding AI agents**
Components: log ingestion → normalization/parsing → schema-on-read vs schema-on-write decision → dual indexing (SQL + vector) → retention policies

### Key tradeoffs to have ready

| Decision | Choice A | Choice B | When to pick A |
|----------|----------|----------|----------------|
| Latency vs accuracy | Smaller/faster model | Larger/slower model | High alert volume, first-pass triage |
| RAG vs fine-tuning | RAG | Fine-tune | Data changes frequently; interpretability matters |
| Vector DB vs pgvector | Dedicated (Qdrant, Pinecone) | pgvector | > 10M vectors, need ANN performance |
| Streaming vs batch | Kafka + Flink | Spark batch | SLA < 30s |
| Human-in-the-loop | Async review | Fully autonomous | High-stakes actions (blocking IPs, closing tickets) |

### Numbers to know
- p99 triage latency target: < 5s
- 1M events/day ≈ 11.5 events/second
- 1B events/day ≈ 11,500 events/second
- L1 cache: ~1ns, RAM: ~100ns, SSD read: ~100μs, network same-DC: ~0.5ms

---

## Study & Retention Strategy

### Core principle: retrieval beats re-reading

Reading notes again doesn't build fluency. What builds fluency is retrieving information under pressure. Every study session should have more talking/writing than reading.

### Technique stack

**1. Whiteboard the anchor design from memory (daily)**
Close all notes. Draw the alert triage agent end-to-end. Narrate it out loud. Then check what you missed.
This single exercise covers ~70% of both rounds. 15 min/day.

**2. Question-first review**
Don't re-read notes top-to-bottom. Open the "Likely Questions" section, pick one, close everything, and answer it out loud for 2-3 minutes before looking. This mirrors what you'll be doing in the actual interview.

**3. Personal story polish (for Round 3)**
Write the story down once in your own words. Then practice saying it out loud 3+ times without reading. Time it — it should land in under 5 minutes.

**4. Systems design mock run**
Pick one of the 5 design scenarios. Set a 25-min timer. Use the 6-step process (clarify → restate → high-level → deep dive → tradeoffs → ops). Talk out loud or type your answer. Grade yourself: did you clarify first? Did you name tradeoffs? Did you surface failure modes?

**5. Teach-back test**
If you can't explain RAG, the agent loop, or hybrid retrieval to a non-technical PM in 60 seconds, you don't know it cold enough. Practice this. It also prepares you for the communication signal Panther evaluates.

**6. Tradeoff flashcards**
For every key decision in the tradeoffs table, practice: given a context, which do you choose and why? The answer should come automatically, not feel calculated.

### Session plan (days until Round 3)

**Day 1 (today): Foundation**
- Read `rounds/ai-integration.md` and `concepts/agentic-ai.md` fully — once
- Whiteboard the triage agent from memory (first pass will be rough — that's fine)
- Write your personal system story draft

**Day 2: AI Integration depth**
- Answer all 6 likely questions out loud without notes
- Practice the "50K events" answer, prompt architecture answer, and metrics answer
- Refine your personal system story — say it out loud 3x

**Day 3: Systems Design depth**
- Run the 6-step mock on "real-time alert triage system" (25 min)
- Run it again on "text-to-search" (25 min)
- Review tradeoffs table until each answer is automatic

**Day 4: Integration**
- Full mock: AI Integration likely questions, no notes (30 min)
- Full mock: Systems design with a new scenario (25 min)
- Fill any gaps you discovered

**Day before each interview: Light polish**
- Re-read `rounds/ai-integration.md` or `rounds/systems-design.md` (not the deep dives — just the signals and checklist)
- Whiteboard the triage agent one more time
- Say your personal story out loud twice
- Sleep — cramming the night before degrades performance

---

## Day-of Checklist

**AI Integration**
- [ ] Can I narrate the triage agent end-to-end without notes?
- [ ] Do I have a 3-5 min personal story ready?
- [ ] Do I know my production metrics (offline + online + safety + drift)?
- [ ] Can I describe the prompt architecture answer?
- [ ] Can I explain prompt injection risk and the defense?

**Systems Design**
- [ ] Will I ask clarifying questions before drawing anything?
- [ ] Do I have a tradeoff ready for every major design decision?
- [ ] Can I name failure modes and monitoring for the system I design?
- [ ] Am I ready to say "I'd approach it differently now" for at least one decision?

---

## See Also

- [[rounds/ai-integration]] — detailed prep for Round 3
- [[rounds/systems-design]] — detailed prep for Round 4
- [[concepts/agentic-ai]] — agent and RAG reference
- [[concepts/systems-design-patterns]] — patterns library
- [[concepts/soc-domain]] — SOC domain context
