---
type: plan
last-updated: 2026-04-26
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
  → analyst review / auto-close only high-confidence benign cases / escalation / ticket
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
- Embedding models: OpenAI `text-embedding-3-small` / `text-embedding-3-large`, BGE-style open models; domain-specific evaluation matters for security data
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
- Online (post-launch): analyst override rate, auto-close reversal rate, time-to-triage, latency against the stated SLA
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
| Vector DB vs pgvector | Dedicated (Qdrant, Pinecone) | pgvector | Need managed vector ops, strict latency/filtering, or independent scaling beyond Postgres |
| Streaming vs batch | Kafka + Flink | Spark batch | SLA is seconds/minutes and results must update continuously |
| Human-in-the-loop | Async review | Fully autonomous | High-stakes actions (blocking IPs, closing tickets) |

### Numbers to know
- Latency target: clarify first; use "seconds for first-pass enrichment, minutes for analyst-facing triage" as a defensible assumption if no SLA is given
- 1M events/day ≈ 11.5 events/second
- 1B events/day ≈ 11,500 events/second
- L1 cache: ~1ns, RAM: ~100ns, SSD read: ~100μs, network same-DC: ~0.5ms

---

## The Three Anti-Blank Scaffolds (memorize these cold)

When nerves hit, your prefrontal cortex temporarily checks out. These scaffolds are muscle memory that buys 30–60 seconds while it comes back online.

### Scaffold 1: The Universal Opener (use for ANY question, both rounds)

> "Great question. Before I dive in, let me make sure I'm solving the right problem — can I clarify a couple of things? *[ask 2–3 clarifying questions about scale, latency, and what we can't get wrong].* Okay, so I'm hearing X. Let me think about this in three layers: the architecture, the tradeoffs, and how this runs in production."

Buys 60 seconds, signals seniority, and locks you into the right frame every time.

### Scaffold 2: The Triage Agent Rosary (nine beads, say it until automatic)

> **Intake → Enrichment → Hybrid Retrieval → Rerank/Summarize → Prompt Assembly → LLM → Policy Gate → Analyst/Action → Feedback Loop**

For each bead: know what runs there and one tradeoff. That's all you need.

### Scaffold 3: The Systems-Design 6-Step (Round 4 only)

> **Clarify → Restate → High-level → Deep dive → Tradeoffs → Operational concerns**

If you blank in Round 4: "Let me work through this systematically — first I want to clarify…" and follow in order. A mediocre answer in this structure beats a brilliant one without it.

---

## 36-Hour Study Schedule

**Core principle: retrieval beats re-reading.** The wiki is ready. These sessions are almost entirely talking/writing, not reading. If your eyes are on the page > 30% of a session, you're doing it wrong.

**Memory-retention techniques embedded in the schedule:**

| Technique | Where used |
|---|---|
| Active retrieval | Every block — close the notes, then talk |
| Spaced repetition | Each anchor concept hit in ≥3 separate sessions with sleep between |
| Sleep consolidation | Both nights protected at 7+ hours — non-negotiable |
| Pre-sleep encoding | Last 10 min before sleep = rosary + 6-step only |
| Generation effect | Write the personal story by hand; write the cheat sheet from memory |
| Dual coding | Whiteboard the triage agent *while* narrating out loud |
| Interleaving | Round 3 and Round 4 topics mixed within sessions |
| Realistic test conditions | Timed mocks, no notes, out loud |

---

### Block A — Tonight, ~3 hours (Encoding & Generation)

Goal: generate the assets you'll rehearse from. Touch every topic once to refresh, not to learn.

| Time | Activity |
|---|---|
| 0:00–0:25 | Re-read this prep-plan and `rounds/ai-integration.md` — once, skimming. Don't try to memorize. |
| 0:25–0:50 | Whiteboard the alert triage agent from memory. Then check lines 38–49 above. Note gaps. Redraw with gaps filled. |
| 0:50–1:30 | **Write your personal system story by hand** (paper or doc). Use the 6-step structure. 400–600 words. |
| 1:30–1:50 | Read it once. Then say it out loud, timed. Target: 3:30–4:30. Trim hard if it runs long. |
| 1:50–2:20 | Skim the 3 micro-gaps: MCP (1 paragraph), confidence scoring (logprobs / self-consistency / verifier model), async inference (background enrichment, streaming partials). One sentence each only. |
| 2:20–2:50 | Rehearse the 3 scaffolds out loud, eyes closed — 5 reps each. Universal Opener, Rosary, 6-Step. |
| 2:50–3:00 | Pre-sleep review: read just the rosary and the 6-step one final time, in bed, then lights out. |

**Sleep: 7–8 hours. This night does more for retention than any extra study hour.**

---

### Block B — Tomorrow morning, ~2 hours (Active Retrieval)

Goal: pure recall under conditions that approximate the interview.

| Time | Activity |
|---|---|
| 0:00–0:10 | Say all 3 scaffolds out loud, no notes. Warm-up and overnight-consolidation check. |
| 0:10–0:40 | **Round 3 mock — solo, out loud, timer on.** Pick 3 of the 6 likely questions above. Answer each in 3–4 min. No notes. Then check the wiki and note what you missed. |
| 0:40–1:10 | **Round 4 mock — 25 min timed.** Prompt: *"Design a real-time alert triage system: 1B events/day, 100 customers, first-pass enrichment in seconds and analyst-ready triage in minutes."* Use the 6-step. Whiteboard or type. Out loud throughout. |
| 1:10–1:30 | Personal story rep: say it out loud once without notes. Then once more to a mirror or webcam. |
| 1:30–2:00 | **Tradeoff drill.** Cover the tradeoff table above. For each row, invent a hypothetical context and say out loud which you'd pick and why. Should feel automatic by the end. |

---

### Block C — Tomorrow afternoon/evening, ~1.5 hours (Polish + Cheat Sheet)

Goal: compress everything into one handwritten page and run one final mock.

| Time | Activity |
|---|---|
| 0:00–0:40 | **Build a 1-page cheat sheet by hand, from memory.** Must include: the rosary, the 6-step, the Universal Opener, the 4 metric axes (offline/online/safety/drift) with one example each, one prompt-injection sentence, the 5 numbers above, the personal-story 6-beat outline. Write it from memory, then check. The act of compressing is the retention. |
| 0:40–1:10 | **Second Round 4 mock** on a different prompt (text-to-search or detection-code-generation). Same 25-min timer. Out loud. |
| 1:10–1:30 | **Failure-mode drill.** For the design you just produced, name 5 things that could go wrong and how you'd detect/recover each. ("Failure Mode Awareness" is an explicit Round 4 rubric signal.) |

**No studying past 9pm. Re-read the cheat sheet for 5 min in bed, then lights out.**

**Sleep: 7–8 hours. This is the night that matters most.**

---

### Block D — Interview day, ~30 min (Activation only — no new material)

| Time | Activity |
|---|---|
| Wake → +20 min | Coffee/water/protein. Say the 3 scaffolds out loud. Then narrate the rosary while sketching the triage agent on paper — 4 min max. |
| +20 → +30 min | Re-read your cheat sheet once, slowly. Then put it away. |
| 15 min before Round 3 | 4-7-8 breathing × 4 cycles (inhale 4s, hold 7s, exhale 8s). Sip water. Small snack (nuts, banana). |

---

## Between-Round Reset

You'll be cognitively fatigued after Round 3. Round 4 needs your clearest thinking.

**5 minutes between rounds:**
1. Stand up and move — different room or step outside.
2. **Do not replay Round 3.** Whatever happened, happened. Rumination steals working memory.
3. Hydrate. Small snack.
4. 4-7-8 breathing × 2 cycles.
5. Re-read your cheat sheet for 60 seconds — only the 6-step and the rosary.
6. Say out loud: *"Clarify first. Then restate. Then high-level."* Walk back in.

---

## What We Are Deliberately NOT Doing

- **No deep dive on MCP, async inference, or capacity-math.** One-sentence inoculation only — new material under stress is a net negative.
- **No re-reading `agentic-ai.md` or `systems-design-patterns.md` end-to-end.** Use them as lookups only if a mock surfaces a gap.
- **No more than 2 timed systems-design mocks.** Sleep beats mock #3.
- **No new design scenarios beyond the 5 above.** The anchor is enough.
- **No studying past 9pm the night before.** Cortisol from late study impairs the sleep that matters most.
- **No wiki review the morning of.** New retrieval failures right before the interview shake confidence. Cheat sheet only.

---

## Verification Checklist

By end of Block C, you should be able to do all of these without notes:

- [ ] Recite the Triage Agent Rosary in < 20 seconds.
- [ ] Recite the Systems-Design 6-Step in < 10 seconds.
- [ ] Deliver the Universal Opener verbatim in one breath.
- [ ] Whiteboard the alert triage agent end-to-end in < 4 min, naming one tradeoff per stage.
- [ ] Tell your personal system story in 3:30–4:30, eye contact forward.
- [ ] Name the 4 metric axes and give one example metric per axis.
- [ ] Pick the right side of all 5 tradeoff-table rows given an arbitrary scenario.
- [ ] List 5 failure modes for the alert triage system and how you'd detect each.

If any item fails at the end of Block C: rehearse that item three more times out loud. Do **not** pull a new wiki page.

---

## Day-of Checklist

**AI Integration (Round 3)**
- [ ] Can I narrate the triage agent end-to-end without notes?
- [ ] Do I have a 3-5 min personal story ready?
- [ ] Do I know my production metrics (offline + online + safety + drift)?
- [ ] Can I describe the prompt architecture answer?
- [ ] Can I explain prompt injection risk and the defense?

**Systems Design (Round 4)**
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
