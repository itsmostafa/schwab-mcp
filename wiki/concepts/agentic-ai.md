# Agentic AI

Reference for the AI Integration round and general interview prep.

## Core Agent Loop

```
Perceive (input + context) → Reason (LLM) → Act (tool call / output) → Observe (result) → loop
```

A "tool" can be: a function call, a search, a code executor, a database write, another agent.

## Agent Memory Types

| Type | Description | Example |
|------|-------------|---------|
| In-context | The prompt/conversation window | Recent messages, system prompt |
| Episodic | Past interactions, stored externally | Vector DB of past analyst decisions |
| Semantic | Structured facts/entities | Knowledge graph of threat actors |
| Procedural | Learned behaviors | Fine-tuned weights, few-shot examples |

## Planning Strategies

**ReAct** (Reason + Act): LLM interleaves reasoning steps and tool calls. Simple, interpretable.

**Reflexion**: Agent reflects on past failures and adjusts. Useful when first-pass accuracy is insufficient.

**Chain-of-Thought**: Force step-by-step reasoning before action. Reduces errors on complex tasks.

**Plan-then-Execute**: Generate full plan first, then execute each step. Better for long-horizon tasks.

## RAG (Retrieval-Augmented Generation)

### Pipeline
```
Query → Embed query → Search vector index → Retrieve chunks → Augment prompt → LLM → Answer
```

### Key Design Decisions
- **Chunking**: fixed-size with overlap (simple) vs semantic/hierarchical (better recall)
- **Embedding model**: general (OpenAI, BGE) vs domain-specific (fine-tuned on security data)
- **Retrieval**: pure vector (cosine similarity) vs hybrid (BM25 + vector, MMR for diversity)
- **Reranking**: cross-encoder after retrieval for precision (slower but better)
- **Filtering**: metadata filters before vector search to reduce noise

### RAG vs Fine-tuning
| | RAG | Fine-tuning |
|--|-----|-------------|
| Cost | Low (indexing only) | High (training) |
| Update frequency | Real-time | Periodic retraining |
| Interpretability | High (can inspect retrieved docs) | Low |
| Best for | Dynamic knowledge, factual recall | Style, format, behavior changes |

## Multi-Agent Patterns

**Orchestrator + Workers**: Central orchestrator assigns tasks to specialized workers. Panther's 4 agents likely follow this pattern.

**Supervisor**: One agent oversees others, can re-try or re-route.

**Peer-to-peer**: Agents communicate directly. More complex state management.

**Pipeline**: Linear chain of agents, each processes output of previous.

## Feedback Loops for Agent Improvement

How Panther's "collective intelligence framework" likely works:
1. Analyst reviews agent's triage decision (accept/reject/modify)
2. Decision + context captured as a training signal
3. Signals aggregated → either:
   - Added to RAG index (retrieval-based improvement)
   - Used for fine-tuning (model-based improvement)
   - Used to update prompt examples (few-shot improvement)
4. A/B test new agent version vs baseline before rollout
5. Rollback mechanism if accuracy drops

## Agent Evaluation

- **Task success rate**: did the agent accomplish the goal?
- **Accuracy**: for triage, is the classification correct?
- **Latency**: acceptable for the use case?
- **Tool call efficiency**: is it over-calling tools? Under-calling?
- **LLM-as-judge**: use a stronger LLM to evaluate weaker agent's outputs
- **Human eval**: analyst ratings on triage quality (ground truth)
- **RAGAS metrics**: faithfulness, answer relevance, context precision (for RAG specifically)

## Hallucination Mitigation (Critical in Security Context)

- Ground responses in retrieved evidence; cite sources
- Use smaller, scoped prompts over large open-ended ones
- Add explicit "I don't know" behavior for low-confidence cases
- Structured output (JSON schema) reduces drift from format
- Constitutional AI / guardrails layer on top of raw LLM output
- Human-in-the-loop for high-stakes actions (blocking, closing tickets)

## See Also

- [[rounds/ai-integration]] — what this round evaluates
- [[panther/role]] — Panther's specific agent architecture
