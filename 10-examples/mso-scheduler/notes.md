# MSO Patient Onboarding Scheduler (Temporal Workflow Orchestration)

## Key Concepts

### System Architecture Overview

An MSO (Medical Service Organization) scheduler orchestrates patient onboarding across multiple downstream services. It doesn't own patient data — it coordinates the sequence: validate intake, create a patient record, verify insurance, schedule appointments, route prescriptions, collect documents, and send notifications. Temporal provides durable workflow execution — if the process crashes mid-step, it resumes exactly where it left off.

```
┌─────────────────────────────────────────────────────────┐
│                   API Server (Gin)                       │
│  Start Workflow │ Query Status │ Signal Events           │
└────────┬────────────────┬──────────────┬────────────────┘
         │                │              │
         ▼                ▼              ▼
┌─────────────────────────────────────────────────────────┐
│                  Temporal Server                         │
│  Workflow State │ Task Queues │ Timers │ Event History   │
└────────┬────────────────────────────────────────────────┘
         │ dispatches tasks
         ▼
┌─────────────────────────────────────────────────────────┐
│                   Worker Process                         │
│  PatientOnboardingWorkflow                               │
│  InsuranceVerificationChildWorkflow                      │
│                                                          │
│  Activities:                                             │
│    ValidatePatientData     → internal validation         │
│    CreatePatientRecord     → Patient DB Service          │
│    VerifyInsurance         → Insurance Eligibility API   │
│    ScheduleAppointment     → Scheduling Service          │
│    RoutePrescription       → Pharmacy Service            │
│    CollectDocuments        → Document Management Service │
│    SendNotification        → Notification Service        │
└──────────────────────────────────────────────────────────┘
```

| Component | Responsibility | Key Libraries |
|-----------|---------------|---------------|
| API Server | HTTP endpoints to start/query/signal workflows | `gin-gonic/gin`, `go.temporal.io/sdk/client` |
| Temporal Server | Durable workflow state, task dispatch, timers, event history | Temporal (self-hosted or Temporal Cloud) |
| Worker | Executes workflow logic and activity side effects | `go.temporal.io/sdk/worker`, `go.temporal.io/sdk/workflow` |
| PatientOnboardingWorkflow | Main orchestration — 7 sequential steps with saga compensation | `go.temporal.io/sdk/workflow` |
| InsuranceVerificationChildWorkflow | Sub-process for async insurance verification with 48h signal wait | `go.temporal.io/sdk/workflow` |
| Activities | HTTP calls to downstream services with idempotency keys | `go.temporal.io/sdk/activity`, `net/http` |

---

### Data Models

The scheduler is stateless — it doesn't persist patient data in its own database. Temporal's workflow history is the source of truth for onboarding state. The data models represent transient workflow inputs, signal payloads, and queryable state.

**`OnboardingRequest`** — Workflow input containing everything needed to onboard a patient:

```go
type OnboardingRequest struct {
    Patient       Patient
    Insurance     InsuranceInfo
    Intake        IntakeQuestionnaire
    Prescriptions []Prescription
    Appointment   Appointment
    Documents     []DocumentRequirement
}
```

**`OnboardingState`** — Queryable workflow state, updated after each step:

```go
type OnboardingState struct {
    Status              OnboardingStatus
    PatientID           string
    InsuranceResult     *InsuranceVerificationResult
    AppointmentID       string
    PrescriptionIDs     []string
    DocumentsCollected  int
    DocumentsRequired   int
    CurrentStep         string
    FailureReason       string
    StartedAt           time.Time
    CompletedAt         *time.Time
}
```

**State machine transitions:**

```
pending → validating → creating_record → verifying_insurance →
  scheduling_appointment → routing_prescription → collecting_documents →
  sending_notification → completed

Any state → failed
```

The state machine is enforced in application code. Invalid transitions return an error. Any state can transition to `failed` — this represents unrecoverable errors where the saga compensation must run.

**Signal payloads:**

| Signal | Channel Name | Payload | Sender |
|--------|-------------|---------|--------|
| Document uploaded | `document_uploaded` | `DocumentUploadedSignal{DocumentType, StorageURL}` | Patient portal |
| Insurance response | `insurance_response` | `InsuranceResponseSignal{Approved, Copay, ...}` | Insurance company webhook |

---

### Temporal Workflow Orchestration

**Workflow determinism** is the most important constraint in Temporal. Workflows are replayed from event history on every task execution — so the code must produce the same result each time. Violations cause nondeterminism errors at runtime.

Banned in workflows:
- `time.Now()` — use `workflow.Now(ctx)`
- `time.Sleep()` — use `workflow.Sleep(ctx, d)` or `workflow.NewTimer(ctx, d)`
- `rand.Intn()` — use `workflow.SideEffect(ctx, func() interface{} { ... })`
- `goroutines` — use `workflow.Go(ctx, func(ctx workflow.Context) { ... })`
- Any I/O (HTTP calls, DB queries, file reads) — must happen in activities

**Activity options** vary per activity type because different downstream services have different reliability characteristics:

| Activity | StartToCloseTimeout | HeartbeatTimeout | MaxAttempts | Backoff |
|----------|-------------------|-----------------|-------------|---------|
| ValidatePatientData | 10s | — | 1 | — |
| CreatePatientRecord | 30s | — | 5 | exponential (1s→30s) |
| VerifyInsurance | 60s | 15s | 5 | exponential (2s→60s) |
| ScheduleAppointment | 30s | — | 5 | exponential (1s→30s) |
| RoutePrescription | 30s | — | 5 | exponential (1s→30s) |
| CollectDocuments | 30s | — | 3 | exponential (1s) |
| SendNotification | 30s | — | 3 | exponential (1s) |

`StartToCloseTimeout` is the maximum time for a single activity attempt. `HeartbeatTimeout` applies to long-running activities that call `activity.RecordHeartbeat` — if the server doesn't receive a heartbeat within this window, it considers the activity stuck and retries it.

---

### Signal and Query Patterns

**Signals** push data into a running workflow from the outside. They're asynchronous — the sender doesn't wait for the workflow to process the signal. Two signal patterns in this project:

1. **Document collection** — The workflow enters a loop waiting for `document_uploaded` signals. Each signal increments `DocumentsCollected`. When `DocumentsCollected >= DocumentsRequired`, the loop exits. A 7-day timer runs concurrently — if it fires first, the workflow proceeds with whatever documents have been collected.

2. **Insurance verification** — The child workflow calls the insurance API, then waits for an `insurance_response` signal if the API doesn't return an immediate result. A 48-hour timer races against the signal via `workflow.Selector`.

```go
// workflow.Selector multiplexes channels and futures (like Go's select).
selector := workflow.NewSelector(ctx)

selector.AddReceive(signalCh, func(c workflow.ReceiveChannel, more bool) {
    var signal DocumentUploadedSignal
    c.Receive(ctx, &signal)
    state.DocumentsCollected++
})

selector.AddFuture(timer, func(f workflow.Future) {
    // timeout — proceed without all documents
})

selector.Select(ctx) // blocks until one branch fires
```

**Queries** provide synchronous read access to workflow state. The API server calls `client.QueryWorkflow` to invoke the `onboarding_status` query handler, which returns the current `OnboardingState`. Queries don't block the workflow and don't appear in the workflow history.

```go
// In the workflow — register a query handler.
workflow.SetQueryHandler(ctx, "onboarding_status", func() (*OnboardingState, error) {
    return state, nil
})

// In the API — query the workflow.
resp, _ := client.QueryWorkflow(ctx, workflowID, "", "onboarding_status")
var state OnboardingState
resp.Get(&state)
```

---

### Saga Pattern

The patient onboarding workflow uses the saga pattern for distributed transactions. Unlike a database transaction, there's no global rollback — each step's side effects (HTTP calls to downstream services) must be explicitly compensated.

Implementation:

```go
var compensations []func(workflow.Context) error

// After creating a patient record:
compensations = append(compensations, func(ctx workflow.Context) error {
    return workflow.ExecuteActivity(ctx, activities.DeletePatientRecord, patientID).Get(ctx, nil)
})

// On failure at any later step:
for i := len(compensations) - 1; i >= 0; i-- {
    compensations[i](compensationCtx) // reverse order
}
```

Compensation order matters. If Step 3 fails:
1. Cancel prescriptions (Step 5 compensation — if reached)
2. Cancel appointment (Step 4 compensation — if reached)
3. Delete patient record (Step 2 compensation)

Compensations use short timeouts (30s) and limited retries (3 attempts). They're best-effort — if a compensation fails, it's logged but doesn't block other compensations. An operator can manually reconcile.

**What's not compensable:** The notification in Step 7 has no compensation — you can't unsend an email. Notification failure is also non-critical: the workflow still completes even if the notification fails.

---

### Child Workflows

The insurance verification is a child workflow rather than a plain activity because:

1. **Long duration** — Insurance verification can take 48 hours. An activity holding a worker slot for that long wastes resources. A child workflow parks the wait in Temporal's timer infrastructure, consuming no worker resources.

2. **Independent lifecycle** — The child workflow has its own WorkflowID (`insurance-verify-{patientID}`), making it independently queryable and cancellable from the Temporal UI or CLI.

3. **Own signal handling** — The child workflow registers its own signal channel (`insurance_response`), keeping the parent workflow's signal space clean.

```go
childCtx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
    WorkflowID:         fmt.Sprintf("insurance-verify-%s", patientID),
    TaskQueue:          TaskQueuePatientOnboarding,
    WorkflowRunTimeout: 72 * time.Hour,
    ParentClosePolicy:  1, // TERMINATE child if parent completes
})

var result InsuranceVerificationResult
err := workflow.ExecuteChildWorkflow(childCtx, InsuranceVerificationChildWorkflow, insurance).Get(ctx, &result)
```

`ParentClosePolicy: TERMINATE` ensures the child workflow is cleaned up if the parent workflow completes or is cancelled. Without this, orphaned child workflows would accumulate.

---

### Activity Patterns

**Heartbeats** — `activity.RecordHeartbeat(ctx, details)` sends a progress update to the Temporal server during long-running activities. If the server doesn't receive a heartbeat within `HeartbeatTimeout`, it considers the activity stuck and schedules a retry on a different worker. The insurance verification activity heartbeats before and after the HTTP call because the insurance API can be slow:

```go
activity.RecordHeartbeat(ctx, "sending verification request")
resp, err := a.HTTPClient.Do(req)
activity.RecordHeartbeat(ctx, "processing verification response")
```

**Idempotency** — Every activity that creates external resources uses an `Idempotency-Key` HTTP header derived from a SHA-256 hash of stable inputs. This ensures Temporal retries don't create duplicate records:

| Activity | Idempotency Key Inputs |
|----------|----------------------|
| CreatePatientRecord | email + DOB |
| ScheduleAppointment | patientID + appointment type |
| RoutePrescription | patientID + medication name |
| SendNotification | notification type + patientID + time bucket (hourly) |

Why SHA-256 instead of UUID? UUIDs are random — a new one on each retry defeats idempotency. A hash of stable inputs produces the same key on every retry.

**Struct-based registration** — Activities are methods on `PatientActivities`, which holds downstream service URLs and an HTTP client. Registering the struct with `w.RegisterActivity(&PatientActivities{...})` makes all exported methods available as activities, with dependencies injected via the struct fields:

```go
type PatientActivities struct {
    PatientServiceURL string
    HTTPClient        *http.Client
}

func (a *PatientActivities) CreatePatientRecord(ctx context.Context, patient Patient) (string, error) {
    // a.PatientServiceURL and a.HTTPClient are available here
}
```

**Error types** — Temporal distinguishes retryable and non-retryable errors:
- Default errors are retryable (network timeouts, 5xx responses).
- `temporal.NewNonRetryableApplicationError` stops retries (invalid input, denied insurance). Use this for permanent failures where retrying with the same input will always fail.

---

### Worker Configuration

The worker polls a single task queue (`patient-onboarding`) for both workflow tasks and activity tasks. Key configuration:

```go
w := worker.New(c, "patient-onboarding", worker.Options{
    MaxConcurrentActivityExecutionSize:     10,
    MaxConcurrentWorkflowTaskExecutionSize: 5,
})
```

- **`MaxConcurrentActivityExecutionSize: 10`** — At most 10 activities run in parallel on this worker. This prevents overwhelming downstream services (patient DB, insurance API). Tune based on downstream rate limits and worker resources.
- **`MaxConcurrentWorkflowTaskExecutionSize: 5`** — At most 5 workflow tasks replay simultaneously. Workflow tasks are CPU-bound (replaying event history), so this should roughly match available cores.

**Graceful shutdown** — `worker.InterruptCh()` returns a channel that receives SIGINT/SIGTERM. On signal:
1. Worker stops polling for new tasks.
2. In-progress activities continue to completion.
3. Once all activities finish, the worker exits.

This prevents partially executed HTTP calls and unnecessary retries during clean deployments.

---

### Production Concerns

**Workflow versioning** — When deploying changes to workflow logic, use `workflow.GetVersion` to branch between old and new code paths. Running workflows continue using the old path; new workflows use the new one:

```go
v := workflow.GetVersion(ctx, "add-lab-work-step", workflow.DefaultVersion, 1)
if v == 1 {
    // new code: schedule lab work after appointment
    workflow.ExecuteActivity(ctx, activities.ScheduleLabWork, ...).Get(ctx, nil)
}
// old code path: no lab work step
```

**Payload encryption** — Patient data is HIPAA-sensitive. In production, configure a `DataConverter` that encrypts workflow inputs, outputs, and signal payloads at rest. Temporal stores these in its event history database, so encryption ensures PHI is not stored in plaintext.

**Monitoring** — Temporal exposes Prometheus metrics out of the box:
- `temporal_workflow_completed` / `temporal_workflow_failed` — onboarding success rate
- `temporal_activity_execution_latency` — downstream service response times
- `temporal_workflow_task_schedule_to_start_latency` — worker capacity (high latency = workers overloaded)
- `temporal_activity_task_schedule_to_start_latency` — activity task queue backlog

**Task queue strategy** — This project uses a single task queue. For production with multiple downstream services at different scales, consider separate task queues (e.g., `insurance-verification` with more workers for slow insurance APIs, `notifications` with fewer workers for rate-limited email/SMS).

---

## Interview Questions

1. **Why Temporal over a message queue (RabbitMQ, SQS) for multi-step orchestration?**

   Message queues give you at-most-once or at-least-once delivery for individual messages, but orchestrating a multi-step process across them requires you to build your own state machine, compensation logic, retry policies, timeout handling, and progress tracking — essentially reimplementing what Temporal provides out of the box. With a queue-based approach, you'd need a database to track which step each patient is on, custom code to handle step failures and compensations, polling or pub/sub to wait for async events (like insurance responses), and a separate timer system for timeouts. Temporal encapsulates all of this in a workflow function that reads like sequential code. The tradeoff is operational complexity — you need to run and maintain the Temporal server (or pay for Temporal Cloud), and your team needs to understand workflow determinism constraints. For simple single-step jobs, a queue is simpler. For multi-step orchestration with human-timescale waits (hours/days), Temporal's value proposition is clear.

2. **What is workflow determinism, and what happens if you violate it?**

   Temporal replays workflow code from event history to reconstruct state after a worker restart or task reassignment. This means the workflow function must produce the same sequence of commands (schedule activity, start timer, etc.) every time it's executed with the same history. If you use `time.Now()`, `rand.Intn()`, or make a direct HTTP call in workflow code, the replay will produce different results than the original execution, causing a "non-determinism detected" error that halts the workflow. You use `workflow.Now(ctx)` instead of `time.Now()`, `workflow.SideEffect()` for random values, and activities for all I/O. The compiler won't catch these violations — they surface at runtime during replay, which makes them especially tricky. The Temporal linter (`go.temporal.io/sdk/contrib/tools/workflowcheck`) can catch common violations statically.

3. **How does the saga pattern handle compensation in this workflow?**

   Each step that creates an external side effect (patient record, appointment, prescription) pushes a compensation function onto a slice. If any later step fails, the compensations execute in reverse order (LIFO) to undo completed steps. For example, if prescription routing fails after the appointment was scheduled and the patient record was created, the saga runs `CancelAppointment` then `DeletePatientRecord`. Compensations use short timeouts and limited retries (3 attempts) — they're best-effort. If a compensation itself fails, it's logged but doesn't block other compensations. An operator would need to manually reconcile. This is the fundamental tradeoff of sagas vs. distributed transactions: eventual consistency with manual intervention for edge cases, rather than guaranteed atomicity.

4. **Why is insurance verification a child workflow instead of just an activity?**

   The insurance company may not respond immediately — the verification can take up to 48 hours. If this were an activity, a worker goroutine would be blocked for the entire duration, wasting resources. A child workflow parks the wait in Temporal's timer infrastructure (backed by the Temporal server's timer queue), consuming zero worker resources while waiting. The child workflow also gets its own WorkflowID, making it independently queryable from the Temporal UI — operators can check insurance verification status without querying the parent workflow. Additionally, the child workflow has its own signal channel and retry policy, keeping the parent workflow's concerns separate.

5. **How do you ensure idempotency in activities that call external services?**

   Every activity that creates a resource computes a deterministic `Idempotency-Key` header from stable inputs — for `CreatePatientRecord`, it's `SHA256(email + DOB)`. The downstream service stores this key and returns the existing resource if the same key is seen again. This is critical because Temporal provides at-least-once activity execution: if a worker crashes after completing the HTTP call but before reporting success to the Temporal server, the activity will be retried on a different worker. Without idempotency, the retry would create a duplicate patient record. SHA-256 is used instead of UUID because UUIDs are random — a new one on each retry would defeat the purpose.

6. **How do signals work in Temporal, and how are they used here?**

   Signals are fire-and-forget messages sent to a running workflow from external systems. The sender calls `client.SignalWorkflow(workflowID, signalName, payload)` — this is asynchronous and returns immediately. Inside the workflow, `workflow.GetSignalChannel(ctx, signalName)` returns a channel that receives the signal payload. The workflow blocks on this channel using `workflow.Selector` (like Go's `select`). In this project, signals handle two async events: document uploads from the patient portal and insurance verification responses from the insurance company's webhook. Signals are durable — if the workflow is not currently running (e.g., between tasks), the signal is buffered and delivered when the workflow resumes.

7. **How do you handle worker deployments without breaking running workflows?**

   Use `workflow.GetVersion(ctx, changeID, minSupported, maxSupported)` to branch between old and new code paths. Running workflows that started before the deployment continue using the old path; new workflows use the new one. For example, if you add a new "lab work scheduling" step between appointment scheduling and prescription routing, you'd wrap the new step in a version check. Workers running the new code can handle both old and new workflows because the version check reads from the workflow's event history to determine which branch to take. Eventually, once all old workflows complete, you can remove the version check and the old code path.

8. **What is `workflow.Selector` and when do you use it?**

   `workflow.Selector` is Temporal's equivalent of Go's `select` statement, but for workflow-safe futures and channels. You add branches for signal channels (`AddReceive`), activity/child workflow futures (`AddFuture`), and timers (`AddFuture` with `workflow.NewTimer`). `selector.Select(ctx)` blocks until one branch fires. It's used whenever you need to race multiple async sources — for example, racing a 48-hour timer against an insurance response signal, or racing a document upload signal against a 7-day collection timeout. Unlike Go's `select`, `Selector` is replay-safe: the same branch fires on replay because Temporal records which branch was taken in the event history.

9. **How do activity heartbeats work, and when are they necessary?**

   `activity.RecordHeartbeat(ctx, progressDetails)` sends a heartbeat to the Temporal server from a running activity. If the server doesn't receive a heartbeat within `HeartbeatTimeout` (configured in `ActivityOptions`), it considers the activity stuck and schedules a retry on a different worker. Heartbeats are necessary for long-running activities where the `StartToCloseTimeout` is significantly longer than the expected failure detection time. In this project, the insurance verification activity heartbeats before and after the HTTP call because the insurance API can be slow (up to 30 seconds). Without heartbeats, a worker crash during the API call wouldn't be detected until the `StartToCloseTimeout` expires (60 seconds). With a 15-second `HeartbeatTimeout`, the failure is detected much faster.

10. **How would you handle HIPAA compliance for patient data in Temporal workflows?**

    Temporal stores workflow inputs, outputs, and signal payloads in its event history database. For HIPAA compliance, configure a custom `DataConverter` that encrypts all payloads with AES-256 before they're stored and decrypts them on read. Temporal's `PayloadCodec` interface supports this — the Temporal team provides a reference implementation for encryption. Additionally, enable TLS for all gRPC connections between workers and the Temporal server, use Temporal's namespace-level access controls to restrict who can view workflow history, and ensure the Temporal server's backing store (Cassandra, MySQL, or PostgreSQL) is encrypted at rest. For audit trails, Temporal's event history is append-only and timestamped, providing a natural audit log of every step, signal, and state change.

11. **What's the difference between retryable and non-retryable errors in Temporal activities?**

    By default, all errors returned from activities are retryable — Temporal will retry the activity according to the `RetryPolicy` in `ActivityOptions`. `temporal.NewNonRetryableApplicationError` creates an error that Temporal will not retry, regardless of the retry policy. Use non-retryable errors for permanent failures where retrying with the same input will always produce the same error — for example, invalid patient data (validation failures), denied insurance (the member ID is invalid), or a 422 response from a downstream service. Use retryable errors (plain `fmt.Errorf` or wrapping with `%w`) for transient failures — network timeouts, 5xx responses, temporary unavailability. Misclassifying errors wastes resources: marking transient errors as non-retryable causes premature workflow failure; marking permanent errors as retryable wastes retry attempts.

12. **How would you design the task queue strategy for production?**

    This project uses a single task queue (`patient-onboarding`) for both workflows and activities. In production, I'd consider splitting into multiple task queues based on downstream service characteristics. For example, `insurance-verification` could have its own queue with more workers (insurance API is slow, so more concurrency helps throughput), while `notifications` could have fewer workers (email/SMS APIs often have rate limits). Workflow tasks should generally stay on one queue since they're CPU-bound replays. The tradeoff with multiple queues is operational complexity — more queues means more worker deployments to manage. Start with one queue and split only when monitoring shows bottlenecks (high `schedule_to_start_latency` on one activity type holding up others).

13. **How do you test Temporal workflows?**

    Temporal provides `testsuite.WorkflowTestSuite` with a test environment that replays workflows in-memory without a Temporal server. You mock activities to return predefined results and verify the workflow's behavior:

    ```go
    func TestPatientOnboardingWorkflow(t *testing.T) {
        suite := &testsuite.WorkflowTestSuite{}
        env := suite.NewTestWorkflowEnvironment()

        // Mock activities.
        env.OnActivity(activities.ValidatePatientData, mock.Anything, mock.Anything).Return(nil)
        env.OnActivity(activities.CreatePatientRecord, mock.Anything, mock.Anything).Return("patient-123", nil)
        // ... mock remaining activities

        // Execute workflow.
        env.ExecuteWorkflow(PatientOnboardingWorkflow, testRequest)
        require.True(t, env.IsWorkflowCompleted())
        require.NoError(t, env.GetWorkflowError())

        // Verify result.
        var state OnboardingState
        env.GetWorkflowResult(&state)
        assert.Equal(t, StatusCompleted, state.Status)
    }
    ```

    For signal testing, use `env.SignalWorkflow(signalName, payload)` and `env.RegisterDelayedCallback(duration, func())` to simulate time passing and signal arrival. Integration tests use `TestServer` from `go.temporal.io/sdk/testsuite` to run a real Temporal server in-process.

14. **What happens if the Temporal server goes down while workflows are running?**

    Running workflows pause — they're not lost. Temporal's workflow state is durably persisted in its backing store (Cassandra, MySQL, or PostgreSQL). When the server comes back up, it resumes all workflows from their last recorded state. Workers reconnect automatically via gRPC retries. Activities that were in-progress during the outage may time out and be retried. The workflow itself doesn't lose progress because its state is reconstructed from the event history. The main impact is latency — any pending timers, signals, or activity dispatches are delayed until the server recovers. This is one of Temporal's key advantages over in-process state machines: the workflow state survives not just worker crashes but infrastructure failures. For high availability, run Temporal in a multi-node cluster with replicated storage.

15. **How would you safely add a new step to the onboarding workflow (e.g., credit check)?**

    Use `workflow.GetVersion` to introduce the new step without breaking running workflows:

    ```go
    v := workflow.GetVersion(ctx, "add-credit-check", workflow.DefaultVersion, 1)
    if v == 1 {
        err = workflow.ExecuteActivity(ctx, activities.RunCreditCheck, patientID).Get(ctx, nil)
        if err != nil { ... }
        compensations = append(compensations, func(ctx workflow.Context) error {
            return workflow.ExecuteActivity(ctx, activities.VoidCreditCheck, patientID).Get(ctx, nil)
        })
    }
    ```

    Running workflows (version `DefaultVersion`) skip the credit check. New workflows (version `1`) execute it. Both versions run on the same workers. Once all old workflows complete, remove the version check and the `DefaultVersion` path. Also add the credit check compensation to the saga slice so it's rolled back if a later step fails. Register the new activities (`RunCreditCheck`, `VoidCreditCheck`) on the worker before deploying — unregistered activities cause task failures. Deploy the worker update before any new workflows can use the new version.

---

## Resources

- Temporal Go SDK: https://github.com/temporalio/sdk-go
- Temporal documentation: https://docs.temporal.io
- Temporal Go SDK samples: https://github.com/temporalio/samples-go
- Temporal workflow determinism: https://docs.temporal.io/workflows#deterministic-constraints
- Temporal data encryption (codec server): https://docs.temporal.io/production-deployment/data-encryption
- Saga pattern: https://microservices.io/patterns/data/saga.html
- gin-gonic/gin: https://github.com/gin-gonic/gin
