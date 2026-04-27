# CSET AI Chatbot Story

**Use for**: Round 3 AI Integration, Project Retro, and CEO conversation.

**Project**: `github.com/itsmostafa/cset-ai-chatbot`  
**One-liner**: Built a RAG chatbot for CISA's Cyber Security Evaluation Tool (CSET) knowledge base so users could ask natural-language questions over NIST-related cybersecurity guidance PDFs and get grounded, page-cited answers.

## Interview Version

Use this as a quick-reference outline, not a script.

**Opening hook**
- Built an AI knowledge assistant around CISA's Cyber Security Evaluation Tool (CSET).
- CSET helps asset owners/operators evaluate IT and OT cybersecurity posture.
- Documents were NIST-related cybersecurity guidance PDFs for public/private org security improvement.
- Core problem: not "chat with a PDF"; make retrieval trustworthy, inspectable, and citation-backed for cybersecurity users.

**Architecture**
- Backend: FastAPI service with a RAG pipeline.
- Agent: pydantic-ai agent using a `query_documents` tool.
- Storage: Redis Stack for vector index, MongoDB for conversation history.
- Streaming: SSE response deltas plus source filename/page events.
- Observability: Phoenix/OpenInference tracing for agent runs and tool calls.

**Ingestion path**
- Accept PDF upload through `/api/load-data/`.
- Validate and stage temp file.
- Parse with Docling using OCR + table-structure extraction.
- Chunk with structure-aware hybrid chunker and `mxbai` tokenizer.
- Embed with `mxbai-embed-large-v1`.
- Store in Redis with metadata: filename, page, token count, chunk id.
- Append semantics: adding one PDF does not wipe the existing index.

**Query path**
- User asks question through `/api/chat/`.
- Save user message to MongoDB.
- Agent calls `query_documents`.
- Embed query.
- Redis cosine search retrieves top 10 candidates.
- Hugging Face cross-encoder reranks for precision.
- Top 5 chunks go into LLM context.
- LLM synthesizes over bounded evidence and returns cited answer.
- Save assistant message to MongoDB.

**Main design principle**
- Separate retrieval from generation.
- Deterministic code owns parsing, chunking, indexing, retrieval, reranking, citations, streaming, persistence, and errors.
- LLM only synthesizes over selected evidence.
- Panther bridge: same boundary applies to SOC agents; code enforces access/tool boundaries, model explains and reasons over bounded context.

**Tradeoff 1: retrieval quality vs latency**
- Basic vector search is faster.
- Security/policy docs need precision because semantically similar chunks can be misleading.
- Two-stage retrieval: bi-encoder for recall, cross-encoder for precision.
- Accepted extra latency to make answers more defensible.

**Tradeoff 2: ingestion quality vs complexity**
- CSET docs can include scanned pages, tables, and standards-style layouts.
- Simple PDF extraction risks bad reading order and lost structure.
- Docling OCR/table extraction plus hybrid chunking makes chunks more human-citable.

**Production-minded pieces**
- MongoDB conversation history.
- Redis append mode for multi-document growth.
- SSE streaming for perceived latency and source events.
- Phoenix/OpenInference traces.
- Unit tests around ingestion, streaming, PDF validation, source events, timeout/error propagation, and conversation logging.
- CPU-only Torch pinning / Docker Compose setup.

**What to improve next**
- Add retrieval and answer-quality eval set.
- Start with 50-100 representative CSET questions.
- Track expected source pages, retrieval hit rate at top-k, citation faithfulness, answer relevance, latency, schema validity, and refusal behavior.
- Add stronger prompt-injection boundaries because document text and future user-provided artifacts are untrusted input.

**Best one-line close**
- "The key lesson was that in security workflows, AI value depends less on having a chatbot and more on controlling the evidence path: what was retrieved, why it was trusted, and how the user can audit the answer."

## 60-Second Version

> "I built a RAG chatbot for CISA's CSET knowledge base, using NIST-related cybersecurity guidance PDFs intended to help public and private organizations improve their security posture. The goal was to make dense cybersecurity assessment documentation usable through natural-language questions without losing trust or traceability. The backend is FastAPI with pydantic-ai, Redis Stack for vector search, MongoDB for conversation history, and Docling for PDF ingestion. PDFs are OCRed, table-aware parsed, chunked with a structure-aware hybrid chunker, embedded with `mxbai-embed-large-v1`, stored with page metadata, retrieved with Redis cosine search, and reranked with `mxbai-rerank-large-v1` before the LLM answers. The important design choice was keeping the LLM bounded: deterministic code owns ingestion, retrieval, reranking, source tracking, and persistence; the model synthesizes over retrieved evidence. If I were hardening it for production, I would add a golden eval set for retrieval/citation quality and stronger prompt-injection boundaries around retrieved content."

## Architecture Anchor

```text
PDF upload
  -> FastAPI /api/load-data/
  -> temp file validation
  -> Docling conversion with OCR + table structure
  -> HybridChunker with mxbai tokenizer
  -> contextualized chunks with filename/page/tokens
  -> mxbai embeddings
  -> Redis Stack vector index, flat cosine, append mode

User question
  -> FastAPI /api/chat/
  -> save user message to MongoDB
  -> pydantic-ai agent
  -> query_documents tool
  -> embed query
  -> Redis top-10 vector retrieval
  -> cross-encoder rerank to top 5
  -> LLM synthesizes grounded answer
  -> stream deltas over SSE
  -> emit source filename/page event
  -> save assistant message to MongoDB
```

## Strong Technical Points To Emphasize

- **Not a wrapper**: The LLM is behind a tool-using agent and only receives bounded, retrieved context.
- **Two-stage retrieval**: Bi-encoder for recall, cross-encoder for precision.
- **Document-aware ingestion**: OCR and table detection matter because cybersecurity assessment docs are often scanned, tabular, or standards-heavy.
- **NIST/CSET context**: The uploaded PDFs were NIST-related cybersecurity guidance for helping public and private organizations assess and improve cybersecurity posture.
- **Citation path**: Chunks retain filename and page, and streaming chat emits source events.
- **Operational thinking**: Docker Compose, Redis Stack, MongoDB, SSE streaming, Phoenix traces, unit tests, and CPU-only Torch pinning.
- **Security relevance**: CSET is tied to cyber posture assessment for IT/OT and critical infrastructure users, so trust, traceability, and failure behavior matter.

## Tradeoffs

**Redis flat vector index vs approximate ANN**  
Chose flat cosine search because the target knowledge base is small enough to fit in memory and exact search avoids recall loss. If the corpus grew substantially, the next step would be HNSW or a dedicated vector DB with recall/latency benchmarking.

**Cross-encoder reranking vs pure vector search**  
Accepted extra latency to improve relevance. This is worth it for security guidance because a superficially similar chunk can lead to a confident but wrong answer.

**Docling plus hybrid chunking vs simpler PDF extraction**  
Accepted more dependencies and slower ingestion for better OCR, tables, layout, and semantic chunk quality.

**Streaming vs simple JSON response**  
Streaming adds async complexity, but it improves perceived latency and creates room for source events and future tool-progress events.

## Questions They May Ask

**How did you evaluate it?**  
"The current repo has unit tests around ingestion and chat streaming, but for production I would add an eval harness. I would create 50-100 representative CSET questions with expected source pages, test retrieval hit rate at top-k, citation faithfulness, answer relevance, schema validity, latency, and refusal behavior when the answer is not in the corpus."

**How would you prevent hallucination?**  
"Constrain the system prompt to documentation-backed answers, pass only retrieved context, preserve source metadata, require citations, and evaluate citation faithfulness. For high-stakes answers, I would expose the source chunks and page links directly rather than asking users to trust prose alone."

**What did you learn?**  
"Chunk quality matters as much as model choice. With technical PDFs, bad parsing or arbitrary character chunks can make a strong LLM look unreliable. The most important work was upstream: OCR, table structure, contextualized chunks, metadata, and reranking."

**How does this relate to Panther?**  
"The same architecture maps to SOC alert triage. Replace CSET PDFs with detection docs, alert history, analyst notes, asset context, and threat intelligence. Keep deterministic services responsible for retrieval, RBAC, tenant isolation, and tool execution. Let the LLM synthesize evidence and produce an explanation with citations."

## Panther Bridge

Use this line to connect the story to the role:

> "This project is a smaller version of the pattern Panther cares about: take messy security knowledge, retrieve the right evidence, let an agent reason over it, and keep the output auditable enough that a security user can trust it."

## Do Not Overclaim

- Do not claim it was deployed across CISA unless that is independently true.
- Do not claim measured user impact unless you have numbers.
- Say "built for CISA/CSET use case" or "built around CISA's CSET knowledge base" rather than "production CISA system" unless confirmed.
- Say "unit-tested and production-minded" rather than "fully production hardened."

## Sources

- Repo inspected locally from private GitHub repository `itsmostafa/cset-ai-chatbot` on 2026-04-27.
- CISA describes CSET as a no-cost tool that guides asset owners and operators through evaluating IT and OT cybersecurity posture: https://www.cisa.gov/ics/Downloading-and-Installing-CSET
- CISA CSET public safety fact sheet notes CSET is offered by DHS/CISA and supports infrastructure security assessments, baseline reports, system security plans, diagrams, JSON import/export, custom assessments, and Cyber Performance Goal assessment support: https://www.cisa.gov/resources-tools/resources/cyber-security-evaluation-tool-cset-fact-sheet-public-safety
