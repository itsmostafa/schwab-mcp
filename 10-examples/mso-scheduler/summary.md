# MSO Patient Onboarding Scheduler — Project Summary

## What It Is

This was a patient onboarding orchestration service for an MSO (Medical Service Organization). When a new patient enrolls, there's a multi-step process — validate their intake form, create a patient record, verify insurance eligibility, schedule their first appointment, route prescriptions to a pharmacy, collect required documents (photo ID, insurance card, consent forms), and send a welcome notification. The scheduler doesn't own patient data — it orchestrates these steps across six downstream services using Temporal for durable workflow execution.

## The Problem It Solved

Patient onboarding touches multiple independent systems that each have their own failure modes, latency profiles, and consistency requirements. Insurance verification alone can take 48 hours if the insurer responds asynchronously. Without orchestration, the ops team was manually tracking which patients were stuck at which step, re-triggering failed steps via admin panels, and hoping nothing got dropped. The scheduler automated the entire flow as a single workflow that survives crashes, handles retries, waits for async events (document uploads, insurance responses), and rolls back completed steps if a later step fails.

## Architecture

The system has three main components:

**API Server (Gin)** — Four HTTP endpoints. `POST /onboard` starts a workflow with a deterministic WorkflowID (hash of patient email, preventing duplicate onboardings). `GET /status` queries the workflow's current state via Temporal's query mechanism. Two signal endpoints let the patient portal and insurance webhook push events into a running workflow — document uploads and insurance verification responses.

**Temporal Server** — Manages workflow state, task queues, timers, and event history. The workflow's progress is durably persisted — if a worker crashes mid-step, the workflow resumes on another worker exactly where it left off. Temporal also handles the 48-hour insurance verification wait and 7-day document collection timeout without consuming worker resources.

**Worker Process** — Executes the workflow logic and activity code. The main `PatientOnboardingWorkflow` runs seven sequential steps, each calling an activity that makes HTTP requests to downstream services. The `InsuranceVerificationChildWorkflow` handles the async insurance flow separately — it calls the insurance API, then waits for a response signal with a 48-hour timeout.

## Data Layer

The scheduler is stateless — it doesn't have its own database. Temporal's event history is the source of truth for onboarding state. Each workflow execution stores its inputs, intermediate results (patient ID, appointment ID, prescription IDs), and current status as part of the workflow's queryable state. This means the API server can answer "what step is this patient on?" by querying Temporal, not a database. The downstream services (Patient DB, Scheduling, Pharmacy, Document Management) own their respective data.

## Key Design Decisions I'd Highlight

**Temporal over message queues.** A queue-based approach (RabbitMQ, SQS) would require building our own state machine, compensation logic, timer system, and progress tracking — essentially reimplementing what Temporal provides. The multi-step flow with human-timescale waits (48-hour insurance, 7-day document collection) made Temporal's durable execution model a natural fit. The tradeoff is operational complexity: running a Temporal cluster and teaching the team workflow determinism constraints.

**Saga pattern for distributed rollback.** Each step that creates a resource in a downstream service pushes a compensation function onto a slice. If a later step fails, compensations run in reverse order — cancel prescriptions, cancel appointment, delete patient record. Compensations are best-effort with limited retries. This provides eventual consistency without distributed transactions, which aren't feasible across six independent services.

**Idempotency keys on all external calls.** Temporal provides at-least-once activity execution — if a worker crashes after completing an HTTP call but before reporting success, the activity retries. Every activity computes a deterministic `Idempotency-Key` header from stable inputs (email + DOB for patient creation, patientID + medication name for prescriptions). The downstream service uses this key to return the existing resource instead of creating a duplicate.

**Child workflow for insurance verification.** Insurance can take 48 hours. Holding an activity slot for that long wastes worker resources. A child workflow parks the wait in Temporal's timer infrastructure (zero resource consumption while waiting) and gives the insurance verification its own independent lifecycle — queryable and cancellable from the Temporal UI.

**Task queue strategy.** Single task queue (`patient-onboarding`) for simplicity. In production with more volume, I'd split into separate queues per downstream service type (e.g., `insurance-verification` with more workers for the slow insurance API, `notifications` with fewer workers for rate-limited email/SMS).

**HIPAA concerns addressed.** Patient data flows through Temporal's event history. For production, a custom `DataConverter` encrypts all workflow payloads (inputs, outputs, signals) at rest using AES-256. Temporal's namespace-level access controls restrict who can view workflow history. TLS on all gRPC connections. The append-only event history provides a built-in audit trail.

## What I Learned

This project taught me the value and cost of workflow orchestration frameworks. The value is massive — what would have been thousands of lines of state machine, retry, and compensation code collapsed into a workflow function that reads like a Go function with sequential steps. The 48-hour insurance wait and 7-day document collection timeout, which would have been painful with queues and cron jobs, became simple `workflow.NewTimer` calls. But the cost is real: workflow determinism is a hard constraint that's invisible to the compiler and only surfaces at runtime during replay. I had to internalize that workflow code is re-executed and must be treated as a pure function of its event history. The Temporal linter catches common violations, but the mental model shift from "code runs once" to "code replays from history" takes deliberate practice.
