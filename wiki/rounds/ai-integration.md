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
- **Brush up on**: RAG, LLM tool calling, MCP (Model Context Protocol), agentic loop design, prompt composition, confidence scoring, SOC alert triage workflows

## What They're Assessing (Deeper Detail)

From the job description and role context:

## Key Topics to Know Cold

### RAG (Retrieval-Augmented Generation)
- Why RAG over fine-tuning: cheaper, more up-to-date, more controllable
- Chunking strategies: fixed-size vs semantic, overlap, hierarchy
- Embedding models: OpenAI text-embedding-3, sentence-transformers, BGE
- Vector stores: Pinecone, Weaviate, Qdrant, pgvector, Chroma
- Retrieval: cosine similarity, MMR (max marginal relevance), hybrid BM25+vector
- Reranking: cross-encoder rerankers for precision
- Evaluation: RAGAS, faithfulness, answer relevance, context precision

### Agentic Frameworks
- Core loop: perceive → reason → act → observe
- Tool use / function calling (Anthropic, OpenAI)
- Planning: ReAct, Reflexion, Chain-of-Thought
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

## See Also

- [[concepts/agentic-ai]] — deep reference on AI agents
- [[panther/role]] — what Panther specifically needs
