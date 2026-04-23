# AI Integration Interview

## What They're Assessing

This round evaluates depth in the AI/ML stack that powers Panther's SOC agents. From the job description, they care about:
- Embeddings, vector databases, RAG
- Agentic engineering frameworks
- ML/AI agent architectures and feedback loops
- Collective intelligence / agents that learn over time

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
