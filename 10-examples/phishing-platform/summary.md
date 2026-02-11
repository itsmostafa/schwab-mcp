# Simulated Phishing Platform — Project Summary

## What It Is

This was a security awareness training platform built in Go. Organizations would sign up, import their employee roster, and run simulated phishing campaigns — basically sending realistic but harmless phishing emails to their own employees to measure susceptibility and train them to recognize threats. It's the same model as products like KnowBe4 or Gophish.

## The Problem It Solved

Companies needed a way to proactively test whether employees would fall for phishing attacks without putting them at actual risk. The platform let security teams create campaigns with customizable email templates, send them across a randomized time window so they'd feel organic, and then track who opened the email and who clicked the link. Employees who clicked were redirected to an educational landing page — not a malicious site — explaining what to look out for.

## Architecture

The system had four main components:

**API Server (Gin)** — RESTful API for managing organizations, employees, and campaigns. Cursor-based pagination on list endpoints to handle large employee rosters efficiently.

**Campaign Scheduler** — When a campaign launches, the scheduler generates a random send time for each recipient within the configured window using `crypto/rand` (not `math/rand`, since we needed the timing to be unpredictable). Each email gets enqueued as an individual asynq task with `ProcessAt` for delayed execution. Asynq stores these in Redis sorted sets and delivers them when the time arrives.

**Email Worker Pool** — A pool of ~20 concurrent asynq workers that dequeue tasks, render the email template with per-recipient personalization, embed a tracking pixel and rewritten links, then send via SMTP. The SMTP call is wrapped in a circuit breaker (sony/gobreaker) — after 5 consecutive failures, the breaker opens and workers fail fast instead of blocking on TCP timeouts against a dead server.

**Tracking Service** — Two Gin endpoints. The tracking pixel endpoint (`/t/open/:token`) serves a 43-byte transparent GIF and records an "opened" event. The click endpoint (`/t/click/:token`) records a "clicked" event and 302-redirects to the educational landing page. Both update Redis HyperLogLog counters for unique open/click counts in real time.

## Data Layer

**MongoDB** for the primary store — four collections: organizations, employees, campaigns, and email events. I chose MongoDB over Postgres here because the email events are high-volume append-only writes with variable metadata per event type (an "opened" event carries IP and user-agent, a "clicked" event carries the link URL). The schema flexibility made documents a natural fit. TTL indexes on `occurred_at` handle automatic 90-day data expiration for GDPR without needing batch cleanup jobs.

**Redis** served seven roles: job queue (asynq), sliding-window rate limiting per org to avoid SMTP throttling, HyperLogLog for memory-efficient unique counting (~12 KB per key regardless of cardinality), template caching, distributed locks to prevent duplicate campaign scheduling, real-time dashboard counters via hashes, and campaign progress tracking.

## Key Design Decisions I'd Highlight

**Idempotent email sends.** Each recipient gets a unique cryptographic token. Before sending, the worker checks MongoDB for an existing "sent" event with that token. The collection has a unique index on `token`, so even if two workers race past the check, only one insert succeeds — the other gets a duplicate key error and acknowledges the task without re-sending. This handles asynq's at-least-once delivery guarantee safely.

**Campaign state machine.** Campaigns follow a strict `draft → scheduled → sending → completed` lifecycle with cancellation allowed from any pre-completed state. Transitions are validated in application code and enforced at the database level with conditional updates that filter on current status, preventing race conditions.

**Circuit breaker on SMTP.** Without it, an SMTP outage would cause every worker goroutine to block on TCP timeouts, stalling the entire pipeline. The breaker opens after 5 consecutive failures, fails fast for 30 seconds, then enters half-open state and lets 3 test requests through. If they succeed, normal operation resumes automatically.

**HyperLogLog for analytics.** Tracking unique opens across a campaign with potentially thousands of recipients would be expensive with a Redis set (O(n) memory). HyperLogLog gives us unique counts with fixed 12 KB memory and 0.81% standard error — more than accurate enough for an analytics dashboard.

## What I Learned

This project pushed me to think carefully about failure modes in distributed systems. The email pipeline has at least seven distinct failure points — worker crashes, SMTP outages, Redis downtime, MongoDB failures, duplicate deliveries, network partitions, scheduler crashes — and each one needs a specific mitigation strategy. The idempotency pattern with unique tokens became the backbone of reliability: no matter what fails or retries, the system converges to the correct state.

I also gained a much deeper understanding of Redis as more than just a cache. Using it simultaneously for job queues, rate limiting, probabilistic counting, distributed locking, and real-time counters showed me how versatile the data structures are when you match the right one to the access pattern.
