# Integration Sync Service

A practice Go microservice that exercises every interview topic in one realistic project.

## What It Is

A service that syncs data from a third-party API into a local Postgres database, publishing events on an internal bus, with full OpenTelemetry instrumentation.

## Why

Covers all interview topics in a single codebase:

- **Go fundamentals** — context propagation, interfaces, error handling
- **HTTP** — outbound client calls to third-party API with retries/timeouts
- **Gin** — inbound REST API with middleware
- **GORM** — Postgres persistence with migrations and associations
- **Event bus** — publish domain events (e.g., `sync.completed`, `record.created`)
- **OpenTelemetry** — traces across HTTP calls, metrics on sync operations
- **Docker** — multi-stage build, docker-compose for local stack
- **Kubernetes** — Helm chart for deployment (stretch goal)

## Architecture

```
┌─────────────────────────────────────────────┐
│                 Gin Handlers                │
│            (REST API endpoints)             │
└──────────────────┬──────────────────────────┘
                   │
┌──────────────────▼──────────────────────────┐
│              Service Layer                  │
│     (business logic, orchestration)         │
└───┬──────────────┬─────────────────┬────────┘
    │              │                 │
┌───▼────┐   ┌────▼─────┐   ┌──────▼───────┐
│  GORM  │   │  HTTP    │   │  Event Bus   │
│  Repo  │   │  Client  │   │  (pub/sub)   │
└───┬────┘   └────┬─────┘   └──────────────┘
    │              │
┌───▼────┐   ┌────▼─────┐
│Postgres│   │Third-party│
│        │   │   API     │
└────────┘   └──────────┘

OTel instrumentation throughout all layers
```

## How to Run

TODO: Set up docker-compose with:
- App (Go service)
- Postgres
- Jaeger (trace viewer)

```bash
# TODO
docker-compose up
```

## Implementation Plan

- [ ] Project skeleton — `cmd/`, `internal/`, `go.mod`
- [ ] GORM models and migrations
- [ ] HTTP client for third-party API
- [ ] Event bus implementation
- [ ] Gin handlers and routes
- [ ] Service layer wiring
- [ ] OTel instrumentation
- [ ] Dockerfile (multi-stage)
- [ ] docker-compose.yml
- [ ] Helm chart (stretch)
