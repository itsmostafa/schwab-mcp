---
tags: [agentic-ai, rag, embeddings, evaluation, interview-prep]
type: concept
last-updated: 2026-04-26
---

# Agentic AI

Reference for agentic AI concepts that may come up across Panther interviews.

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

**ReAct** (Reason + Act): LLM interleaves reasoning traces and task-specific actions/tool calls. Useful as a conceptual pattern for agents that need to inspect external state before deciding.

**Reflexion**: Agent incorporates feedback from previous attempts into future attempts. Useful as a research pattern for tasks where retry and feedback loops are available.

**Chain-of-Thought / structured reasoning**: Asking for intermediate reasoning can improve some multi-step tasks, but do not rely on hidden reasoning as an audit trail. In production, prefer explicit evidence, concise rationales, tool traces, and structured outputs.

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
| Cost | Usually lower to update knowledge; still pays retrieval/storage/inference cost | Training/eval cost; can reduce prompt/retrieval overhead for stable behavior |
| Update frequency | Can incorporate new documents quickly once indexed | Requires retraining/redeployment to absorb new knowledge |
| Source grounding | Can cite retrieved evidence if the pipeline preserves source IDs | Does not inherently cite sources |
| Best for | Dynamic knowledge, factual lookup, inspectable evidence | Style, format, domain behavior, task specialization |

## Multi-Agent Patterns

**Orchestrator + Workers**: Central orchestrator assigns tasks to specialized workers. This is a plausible pattern for SOC agents, but the exact Panther internal architecture is not public.

**Supervisor**: One agent oversees others, can re-try or re-route.

**Peer-to-peer**: Agents communicate directly. More complex state management.

**Pipeline**: Linear chain of agents, each processes output of previous.

## Feedback Loops for Agent Improvement

How Panther's "collective intelligence framework" could work, inferred from the job post and Panther's public AI alert-triage material:
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
- Structured output (JSON schema) reduces format drift but does not replace semantic validation
- Policy checks, output validators, and tool-permission gates outside the model
- Human-in-the-loop for high-stakes actions (blocking, closing tickets)

## Validation Sources

- RAG: Lewis et al., "Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks" (NeurIPS 2020).
- Agent patterns: ReAct (Yao et al., 2022), Reflexion (Shinn et al., 2023), and chain-of-thought prompting (Wei et al., 2022).
- RAG evaluation: RAGAS describes faithfulness, answer relevancy, and context precision as evaluation dimensions.
- RAG vs fine-tuning: AWS Prescriptive Guidance recommends starting with RAG for custom-document QA and using fine-tuning for additional behavior/task needs; the two can also be combined.

## See Also

- [[panther/role]] — Panther's specific agent architecture
