# Go Backend Interview Prep

Study notes and practice material for a Go backend interview (integrations team).

## Role Overview

The role is part of the **integrations team** — pulling and pushing data to many other services. Requirements:

- Experience with Go microservices
- Developing locally with containers
- Deploying and monitoring services on K8s, ArgoCD, Helm, Grafana
- HTTP integrations with third-party APIs
- Self-guided, able to read docs, figure out things on your own
- Good attitude and low ego
- Collaborative mindset

## Study Plan

Work through topics in order — each builds on the previous.

- [ ] [01 — Go Fundamentals](01-go-fundamentals/notes.md) — Goroutines, channels, `context.Context`, interfaces, error handling
- [ ] [02 — HTTP](02-http/notes.md) — `net/http` client & server, middleware, timeouts, retries
- [ ] [03 — Gin](03-gin/notes.md) — Router groups, binding, validation, middleware, error handling
- [ ] [04 — GORM](04-gorm/notes.md) — Models, migrations, associations, transactions, N+1
- [ ] [05 — Event Bus](05-event-bus/notes.md) — Bus pattern, pub/sub, in-process vs distributed
- [ ] [06 — OpenTelemetry](06-opentelemetry/notes.md) — Traces, spans, metrics, logs, SDK vs API, exporters
- [ ] [07 — Docker](07-docker/notes.md) — Multi-stage builds, compose for local dev
- [ ] [08 — Kubernetes](08-kubernetes/notes.md) — K8s objects, Helm, ArgoCD, Grafana

## Practice Project

[Integration Sync Service](project/README.md) — a Go microservice that exercises every topic above in one realistic project.
