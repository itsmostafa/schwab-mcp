# AI Integration Interview

**Format**: 45-60 min · Technical Design Discussion (no live coding)

**What this is**: A high-level design conversation to confirm you have the technical acumen to build production-grade AI agents. You'll design an AI-powered SOC alert triage agent, walk through architecture decisions, embedding and retrieval strategy, deterministic vs. LLM-driven logic, and how the system would operate in production. **Clarity of reasoning matters more than perfect syntax.**

## Signals Panther Is Evaluating

| Signal | What It Means |
|--------|--------------|
| **Agent Architecture Depth** | Layered agentic loops, tool design, composition, evaluation — not just LLM wrappers |
| **RAG & Retrieval Design** | Embeddings, vector DBs, multi-layered pipelines, token budget management |
| **Prompt Engineering at Scale** | Composable, maintainable prompt systems across multiple conversation types |
| **Security / SOC Fluency** | What's safe to automate vs. what requires human judgment in a SOC context |
| **Production Engineering Maturity** | Evaluation, monitoring, drift detection, failure modes — not prototype thinking |
| **AI Security Awareness** | Prompt injection, adversarial inputs, defense-in-depth in a security product |
| **Product & Customer Thinking** | Measurable outcomes, customer trust, phased autonomy over time |

## Official Prep Checklist

- **Know agent architecture fundamentals**: walk through an agentic loop end-to-end — context pre-loading, tool calling, deterministic vs. LLM steps, async inference, structured output, post-inference processing
- **Know your RAG stack**: discuss multi-layered retrieval, token budget constraints, handle "50K matching events in a 200K token context window"
- **Think prompt engineering at scale**: how do you avoid 12 monolithic prompt strings that drift independently? Composable modules, shared bases, testing
- **Know AI security blindspots**: prompt injection is real when log events are attacker-controlled; structural defenses + RBAC at the tool execution boundary
- **Think in production metrics**: offline eval sets, online metrics (analyst override rate, classification distribution), regression detection, feedback loops
- **Brush up on**: RAG, LLM tool calling, MCP (Model Context Protocol for exposing tools/resources/prompts), agentic loop design, prompt composition, confidence scoring, SOC alert triage workflows

## What They're Assessing (Deeper Detail)

From the job description and role context, the bar is not "can you define RAG?" It is whether you can reason through a production agent that a real SOC team would trust.

- Can you separate deterministic control flow from probabilistic model judgment?
- Can you design retrieval that handles noisy security data, high cardinality, and token limits?
- Can you define tool boundaries, especially read-only vs. side-effecting actions?
- Can you describe evaluation before launch, monitoring after launch, and rollback when quality drops?
- Can you speak fluently about analyst workflows, not just model APIs?
- Can you make safety, auditability, and customer trust first-class design constraints?

## Preparation Plan

### 1. Prepare one canonical design: alert triage agent

Be ready to whiteboard this end to end:

```
Alert fires
→ deterministic intake / schema validation
→ enrichment fan-out: alert details, user history, asset context, threat intel, related logs
→ retrieval: similar alerts, past analyst decisions, detection docs, threat intel notes
→ prompt assembly with ranked evidence and token budget
→ LLM reasoning with structured output
→ policy gate: confidence + severity + allowed action
→ analyst review / ticket / escalation / limited auto-close for high-confidence benign cases
→ feedback capture and evaluation dataset update
```

The key is to narrate where the LLM is useful and where it should not be trusted. Use code and policy for validation, permissions, routing, confidence thresholds, audit logs, and side effects. Use the LLM for synthesis, prioritization, natural-language explanation, and ambiguity handling.

### 2. Practice the "50K events / 200K context" answer

Do not say "put it all in context." A strong answer:

- Start with structured filtering: tenant, detection ID, time range, entities, severity, data source
- Use hybrid retrieval: SQL/BM25 for exact indicators and vector search for semantic similarity
- Retrieve in layers: top entities, representative event clusters, past similar alerts, related detection docs
- Rerank for precision and diversity
- Summarize or aggregate repetitive events before the final prompt
- Keep citations/evidence IDs so the analyst can inspect the raw events

### 3. Prepare a prompt architecture answer

For "12 monolithic prompt strings," answer with modular prompt composition:

- Shared base: role, safety constraints, output schema, "cite evidence or say unknown"
- Task module: triage vs. chat vs. detection generation vs. text-to-search
- Customer/tenant policy module: allowed actions, severity thresholds, integrations
- Retrieved evidence module: alert, enrichment, past decisions, threat intel
- Output contract: JSON schema with classification, confidence, rationale, evidence IDs, recommended next action
- Tests: golden prompts, regression cases, prompt diff review, offline eval before rollout

### 4. Prepare production metrics

Name both AI quality metrics and SOC outcome metrics:

- Offline: precision/recall by severity, false-negative rate, calibration, citation faithfulness, schema-valid output rate
- Online: analyst override rate, escalation acceptance rate, time-to-triage, auto-close reversal rate, tool error rate, latency p50/p95/p99, cost per alert
- Safety: unauthorized tool attempt rate, prompt-injection detections, policy-gate blocks, high-severity human-review coverage
- Drift: classification distribution changes, retrieval hit-rate changes, tenant-specific degradation, feedback disagreement trends

### 5. Prepare the security answer

SOC logs are attacker-controlled input. Treat raw event text as untrusted data:

- Keep instructions separate from retrieved/log content
- Use structured tool APIs, RBAC, and server-side authorization
- Sanitize or quote untrusted text in prompts
- Require human confirmation or policy gates for write tools
- Log every tool call and decision for auditability
- Use allowlists for tools and enforce tenant isolation at the data layer

### 6. Rehearse your own system story

Have a 3-5 minute story about an agentic, AI, security, or automation system you built. Structure it as:

1. Problem and user pain
2. Architecture and key data flow
3. Two tradeoffs you made
4. How you evaluated quality
5. What failed or surprised you
6. What you would change now

## Key Topics to Know Cold

### RAG (Retrieval-Augmented Generation)
- Why RAG over fine-tuning: better for frequently changing knowledge and inspectable evidence; fine-tuning is better for stable behavior, style, and task specialization
- Chunking strategies: fixed-size vs semantic, overlap, hierarchy
- Embedding models: OpenAI `text-embedding-3-small` / `text-embedding-3-large`, Sentence Transformers, BGE-style open models
- Vector stores: Pinecone, Weaviate, Qdrant, pgvector, Chroma
- Retrieval: cosine similarity, MMR (max marginal relevance), hybrid BM25+vector
- Reranking: cross-encoder rerankers for precision
- Evaluation: RAGAS, faithfulness, answer relevance, context precision

### Agentic Frameworks
- Core loop: perceive → reason → act → observe
- Tool use / function calling (Anthropic, OpenAI)
- Planning: ReAct, Reflexion, plan-then-execute, structured reasoning
- Multi-agent: orchestrator + worker pattern, message passing
- Agent memory: in-context, external (episodic, semantic, procedural)
- Temporal/workflow orchestration: Temporal, LangGraph, CrewAI

### Feedback Loops & Collective Intelligence
- RLHF vs RLAIF
- Implicit feedback signals (click-through, thumbs, escalation rates)
- How to use analyst feedback to improve agent accuracy over time
- Active learning: which examples to label next
- Evaluation pipelines for agents (LLM-as-judge, human eval)

### Vector Databases
- Indexing: HNSW, IVF-Flat, PQ
- Trade-offs: recall vs latency vs memory
- When to use pgvector vs dedicated vector DB

## Validation Notes

- Panther's public job post confirms the four agent capabilities: alert triage, interactive chat, detection code generation, and text-to-search.
- Panther's AI alert-triage blog describes autonomous investigation, MCP-connected tools, evidence-backed classifications, feedback into detection logic, and conservative auto-close thresholds for benign classifications.
- RAG and tool-calling guidance here is a practical interview synthesis; exact latency targets, retrieval thresholds, and autonomy levels should be stated as assumptions during the interview.

## Panther-Specific Angle

Connect everything back to SOC automation:
- Alert triage agent: RAG over known threat patterns + past analyst decisions
- Text-to-search: embedding-based semantic search over security logs
- Collective intelligence: analyst resolutions become training signal for future triage
- Detection code generation: LLM + examples (few-shot) from existing detection rules

## Likely Questions

- "Walk me through how you'd build a RAG pipeline for alert triage."
- "How would you design a feedback loop so the triage agent improves over time?"
- "What's the difference between RAG and fine-tuning? When would you use each?"
- "How do you evaluate an AI agent? What metrics matter for a SOC use case?"
- "Walk me through an agentic system you've built."
- "How do you handle hallucination in a security context where accuracy is critical?"

## Strong Answer Shape

For most questions, use this structure:

1. **Clarify**: scale, latency target, action scope, customer trust constraints
2. **State the architecture**: components and data flow first
3. **Call out boundaries**: deterministic code vs. retrieval vs. LLM judgment
4. **Discuss tradeoffs**: latency/cost/accuracy, autonomy/safety, recall/precision
5. **Make it production-grade**: evals, monitoring, rollback, audit, human review
6. **Tie back to SOC value**: less analyst toil, faster triage, better coverage, preserved trust

## See Also

- [[rounds/round-3-interview-questions]] — spoken-answer drill bank for likely Round 3 questions
- [[rounds/round-3-cheat-sheet]] — last-minute cheat sheet for this round
- [[projects/cset-ai-chatbot-story]] — personal AI/RAG system story for Round 3
- [[concepts/agentic-ai]] — deep reference on AI agents
- [[concepts/tool-calling]] — tool design, agent loop, read vs. write tools
- [[concepts/soc-domain]] — SOC workflow and Panther product context
- [[panther/role]] — what Panther specifically needs
