# Systems Design Patterns

Reference patterns for the systems design round.

## Patterns Relevant to Panther

### Data Pipeline / Event-Driven Architecture
```
Source → Ingestion (Kafka/Kinesis) → Processing (Flink/Spark/Lambda) → Storage → Query layer
```
- Backpressure handling: Kafka consumer lag monitoring
- Exactly-once vs at-least-once: important for security (don't lose alerts, don't double-triage)
- Schema evolution: how to handle log format changes without breaking pipeline

### Real-Time AI Inference
```
Event → Feature extraction → Embedding → Vector search → LLM prompt → Output
```
- Latency budget: p99 < 5s for real-time triage
- Caching: embed common queries, cache frequent retrievals
- Fallback: if LLM is slow/down, fall back to rule-based triage

### CQRS (Command Query Responsibility Segregation)
Separate write path (ingestion, detection, triage actions) from read path (search, dashboards, reporting).
Useful for Panther: high-throughput ingestion doesn't block analyst query performance.

### Fan-out / Scatter-Gather
One event triggers multiple parallel enrichment lookups (IP reputation, user history, asset context), results aggregated before LLM call. Reduces latency vs sequential calls.

### Circuit Breaker
If an external enrichment service (VirusTotal, Shodan) is slow, break the circuit and proceed without that enrichment rather than blocking triage.

### Feedback Loop Architecture
```
Action → Outcome observed → Signal captured → Model/index updated → New action improved
```
Key: make signals explicit, structured, attributable to specific decisions.

---

## General Systems Design Checklist

When approaching any design question:

**Scope (first)**
- [ ] Users / scale (req/s, data volume, SLAs)
- [ ] What's in scope for this design session?
- [ ] Any non-functional requirements (latency, availability, consistency)?

**High-level design**
- [ ] Major components sketched
- [ ] Data flow between components
- [ ] Where data lives (databases, caches, queues)

**Data model**
- [ ] What entities? What relationships?
- [ ] SQL vs NoSQL trade-offs for this use case
- [ ] Indexing strategy

**APIs / interfaces**
- [ ] What are the key API contracts between components?
- [ ] Sync vs async communication?

**Scaling**
- [ ] What's the bottleneck?
- [ ] Horizontal vs vertical scaling?
- [ ] Sharding / partitioning strategy?

**Failure modes**
- [ ] What happens if component X fails?
- [ ] How do we detect failures?
- [ ] Recovery strategy?

**Monitoring**
- [ ] Key metrics for each component
- [ ] Alerting thresholds
- [ ] Logging strategy

---

## Numbers to Know

| Thing | Approximate value |
|-------|-------------------|
| L1 cache access | ~1 ns |
| L2 cache access | ~10 ns |
| RAM access | ~100 ns |
| SSD read | ~100 μs |
| HDD seek | ~10 ms |
| Network round trip (same DC) | ~0.5 ms |
| Network round trip (cross-region) | ~30-100 ms |
| 1 Gbps network throughput | ~125 MB/s |
| 1M events/day | ~11.5 events/second |
| 1B events/day | ~11,500 events/second |

## See Also

- [[rounds/systems-design]] — how to approach the systems design round
- [[concepts/soc-domain]] — security domain context for design prompts
