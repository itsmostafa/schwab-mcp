---
type: audit
last-updated: 2026-04-26
tags: [audit, systems-design, validation]
---

# Systems Design Validation Audit

Scope: system-design concepts across `wiki/rounds/systems-design.md`, `wiki/prep-plan.md`, `wiki/study-prompt.md`, and linked concept/company pages.

Source standard: primary-first external sources. Private/raw interview material is allowed only for interview-process context, not as public technical validation.

## Summary

Most core concepts are valid: Panther's public job post confirms the role focus, four SOC agent capabilities, collective-intelligence language, 5-10x coverage goal, direct Founder/CTO collaboration, and production AI/security requirements. RAG, tool calling, prompt-injection risk, hybrid retrieval, HNSW/IVFFlat, and streaming/idempotency themes are broadly supported by primary sources.

Corrections made:
- Reframed auto-close as only for high-confidence benign/low-risk cases within policy.
- Replaced hard latency and vector-count thresholds with explicit assumptions.
- Clarified RAG vs fine-tuning tradeoffs: dynamic knowledge/evidence grounding vs stable behavior/style/task specialization.
- Clarified MCP as a protocol for tools/resources/prompts, not an agent framework.
- Clarified tool-calling provider mechanics and host-side execution/security boundaries.
- Replaced "force chain-of-thought" wording with production-safe structured reasoning, evidence, and tool traces.

## Claim Register

| Claim | Status | Evidence | Action |
|---|---|---|---|
| Panther role builds autonomous AI agents for alert triage, interactive chat, detection code generation, and text-to-search. | Confirmed | Panther job post says the role spans those four capabilities and expands the current suite of 4 SOC agents. | Kept; cited in validation notes. |
| Panther's goal is 5-10x more SOC data coverage without proportional headcount growth. | Confirmed | Panther job post states this vision directly. | Kept. |
| Panther has raised $140M from Coatue, Lightspeed, ICONIQ Growth, Snowflake Ventures, and others. | Confirmed | Panther job post "About Panther" section states this. | Kept. |
| Panther Detection-as-Code uses Python detections, version control, and testing workflows. | Confirmed | Panther docs describe detections-as-code, Python detections, and testing guidance. | Kept. |
| Panther AI Alert Triage performs autonomous investigations, uses MCP-connected tools, and shows evidence-backed classifications. | Confirmed | Panther Apr 1, 2026 blog describes autonomous investigation, MCP tool access, classification, evidence, and reasoning. | Kept; used as source for design framing. |
| Auto-close should happen for low-confidence false positives. | Incorrect | Panther describes auto-close when confidence on a benign classification meets a configured threshold. | Corrected to high-confidence benign/low-risk within policy. |
| SOC logs/event text should be treated as untrusted input for prompt injection. | Confirmed | OWASP LLM Top 10 lists prompt injection; OWASP cheat sheet recommends structured separation, least privilege, HITL, validation, and monitoring. | Kept; wording tightened. |
| Tool calling means the model executes external functions. | Needs nuance | OpenAI and Anthropic docs describe model-generated tool requests; Anthropic distinguishes client tools executed by the app from server tools. | Clarified client-side execution and host policy boundaries. |
| Anthropic tool-use loop uses `stop_reason: "tool_use"`, `tool_use`, and `tool_result`. | Confirmed | Anthropic tool-use docs describe this client-tool loop. | Kept. |
| OpenAI function tools are defined with JSON Schema-style parameters and strict mode can constrain outputs. | Confirmed | OpenAI function-calling docs describe function tools, JSON Schema parameters, and strict mode. | Reflected in tool-calling notes. |
| MCP is a protocol for integrating LLM apps with external context and tools. | Confirmed | MCP spec says servers provide Resources, Prompts, and Tools over JSON-RPC with security/user-consent principles. | Clarified wording. |
| RAG is appropriate for frequently changing knowledge and source-grounded answers. | Confirmed | Lewis et al. define retrieval-augmented generation; AWS guidance says RAG can incorporate latest documents quickly and provide source references. | Kept. |
| Fine-tuning is simply "high cost" and less interpretable than RAG. | Needs nuance | AWS guidance: fine-tuning helps style, compliance/domain behavior, and tasks; it does not inherently cite sources and is less suitable for frequently changing documents. | Updated comparison table. |
| Hybrid retrieval with lexical/BM25 plus vector search is valid for security data. | Confirmed with nuance | Weaviate/Qdrant docs describe hybrid search combining keyword/sparse and vector retrieval; this is especially plausible where exact IOCs matter. | Kept as a design pattern, not a guarantee. |
| RAGAS metrics include faithfulness, answer relevancy, and context precision. | Confirmed | RAGAS paper describes evaluation dimensions for retrieval and generation. | Kept. |
| HNSW and IVFFlat are common ANN vector index types. | Confirmed | pgvector docs support HNSW and IVFFlat; HNSW paper describes approximate k-NN with hierarchical navigable small-world graphs. | Kept. |
| pgvector should only be used below 10M vectors. | Unsupported | pgvector supports exact and approximate search; practical thresholds depend on workload, filtering, hardware, and operations. | Removed hard threshold; recommend benchmarking and managed-vector choice when ops/latency/filtering require it. |
| Real-time triage should target p99 < 5s. | Unsupported as a universal claim | Panther customer examples use minutes-to-triage outcomes; system-design latency should be clarified per prompt. | Replaced with explicit SLA assumptions. |
| Kafka/Flink "exactly once" solves duplicate triage. | Needs nuance | Kafka documents exactly-once processing semantics, but end-to-end side effects still require idempotency/deduplication. | Reworded to delivery semantics and idempotency at boundaries. |
| Chain-of-thought should be forced for production agents. | Needs nuance | CoT research supports reasoning benefits on some tasks, but production auditability should rely on evidence, tool traces, and structured rationales. | Reworded to structured reasoning. |
| Orchestrator + worker is Panther's internal architecture. | Unsupported | Public sources confirm multiple agent capabilities but not internal orchestration. | Reworded as plausible, not factual. |

## Primary Sources

- Panther job post: https://job-boards.greenhouse.io/pantherlabs/jobs/7622655003
- Panther AI Alert Triage: https://panther.com/blog/ai-powered-alert-triage
- Panther docs overview: https://docs.panther.com/
- Panther detections docs: https://docs.panther.com/detections
- Panther Python detections: https://docs.panther.com/detections/rules/python
- OpenAI function calling: https://developers.openai.com/api/docs/guides/function-calling
- Anthropic tool use: https://platform.claude.com/docs/en/agents-and-tools/tool-use/overview
- Anthropic models overview: https://platform.claude.com/docs/en/about-claude/models/overview
- MCP specification: https://modelcontextprotocol.io/specification/draft
- OWASP LLM Top 10: https://owasp.org/www-project-top-10-for-large-language-model-applications/
- OWASP Prompt Injection Prevention Cheat Sheet: https://cheatsheetseries.owasp.org/cheatsheets/LLM_Prompt_Injection_Prevention_Cheat_Sheet.html
- NCSC prompt injection blog: https://www.ncsc.gov.uk/blog-post/prompt-injection-is-not-sql-injection
- RAG paper: https://proceedings.neurips.cc/paper/2020/hash/6b493230205f780e1bc26945df7481e5-Abstract.html
- RAGAS paper: https://arxiv.org/abs/2309.15217
- AWS RAG vs fine-tuning guidance: https://docs.aws.amazon.com/prescriptive-guidance/latest/retrieval-augmented-generation-options/rag-vs-fine-tuning.html
- pgvector README: https://github.com/pgvector/pgvector
- HNSW paper: https://arxiv.org/abs/1603.09320
- Apache Kafka semantics: https://kafka.apache.org/documentation/#semantics
- ReAct paper: https://arxiv.org/abs/2210.03629
- Reflexion paper: https://arxiv.org/abs/2303.11366
- Chain-of-thought paper: https://arxiv.org/abs/2201.11903

## Follow-Up Watch Items

- Re-check provider model names before using code snippets in a real codebase; current model aliases change faster than interview concepts.
- If Panther publishes detailed architecture for its AI SOC platform, replace inferred architecture notes with product-specific details.
- Treat all numeric thresholds in interview prep as assumptions unless the interviewer provides an SLA or scale target.
