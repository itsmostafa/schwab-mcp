# CSET AI Chatbot Story

**Use for**: Round 3 AI Integration, Project Retro, and CEO conversation.

**Project**: `github.com/itsmostafa/cset-ai-chatbot`  
**One-liner**: Built a RAG chatbot for CISA's Cyber Security Evaluation Tool (CSET) knowledge base so users could ask natural-language questions over NIST-related cybersecurity guidance PDFs and get grounded, page-cited answers.

## Interview Version

> "A project I would use for this round is an AI knowledge assistant I built around CISA's Cyber Security Evaluation Tool, or CSET. CSET is used by asset owners and operators to evaluate IT and OT security posture, and the documents I loaded were NIST-related cybersecurity guidance PDFs meant to help public and private organizations improve their cybersecurity stance. A lot of that useful knowledge lives in dense PDFs, assessment guidance, tables, and standards-oriented documentation. The problem was not just 'put a chatbot in front of a PDF.' The real problem was how to make retrieval accurate enough that a cybersecurity user could trust the answer and inspect where it came from."
>
> "I built it as a FastAPI service with a RAG pipeline behind it. The ingestion path accepts PDFs, runs them through Docling with OCR and table-structure extraction, chunks them with a structure-aware hybrid chunker, embeds the chunks with `mxbai-embed-large-v1`, and stores them in Redis Stack as a vector index with metadata like filename, page, token count, and chunk id. On the query path, the pydantic-ai agent calls a `query_documents` tool. That tool embeds the query, retrieves the top 10 candidates from Redis using cosine similarity, then reranks them with a Hugging Face cross-encoder and passes the top 5 into the LLM context. The chat response streams back over SSE and the system tracks source filename and page citations."
>
> "The main design decision was to separate retrieval from generation. I did not want the LLM browsing the whole document mentally or relying on memory. Deterministic code handles PDF parsing, chunking, indexing, retrieval, reranking, response streaming, persistence, and error paths. The LLM is used for synthesis over already-selected evidence. That is the same line I would draw in a SOC product: let code enforce boundaries and let the model explain, summarize, and reason over bounded context."
>
> "A second tradeoff was retrieval quality versus latency. Basic vector search was fast, but for policy and security documentation, semantically similar chunks are not always the most answer-worthy chunks. I added a two-stage retrieval pipeline: bi-encoder retrieval for recall, cross-encoder reranking for precision. That adds cost and latency, but it improves the quality of the final context and makes the answer more defensible."
>
> "The third tradeoff was ingestion quality versus implementation complexity. CSET documents include scanned text and tables, so simple PDF text extraction would lose structure or produce bad reading order. I used Docling with OCR and table detection, then hybrid chunking that respects document structure before splitting by token budget. That made the retrieved chunks much closer to how a human would cite the material."
>
> "I also added production-minded pieces: MongoDB conversation history, Redis append semantics so uploading one PDF does not wipe the existing index, streaming responses so users are not blocked waiting for the full answer, Phoenix/OpenInference tracing for agent runs and tool calls, and a mocked unit test suite around ingestion, streaming behavior, PDF validation, source events, timeout/error propagation, and conversation logging."
>
> "What I would improve next is evaluation. The system has unit tests, but the next production step would be a retrieval and answer-quality eval set: known CSET questions, expected source pages, citation faithfulness, answer relevance, and refusal behavior when the docs do not contain the answer. I would also tighten security around prompt injection because in a security domain, document text and future user-provided artifacts should be treated as untrusted input."

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
