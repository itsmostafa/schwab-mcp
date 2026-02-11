package main

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// Temporal workflow definitions for patient onboarding orchestration.
//
// Key interview concepts:
//   - Workflow determinism: workflows must produce the same result when replayed.
//     No time.Now(), rand, or goroutines — use workflow.Now(), workflow.SideEffect(),
//     and workflow.Go() instead.
//   - Query handlers: synchronous read-only access to workflow state. The caller
//     gets the current OnboardingState without waiting for the workflow to finish.
//   - Signal channels: async events pushed into a running workflow from outside
//     (e.g., "patient uploaded a document"). The workflow blocks on the channel
//     until the signal arrives or a timer fires.
//   - Saga pattern: track compensating actions in a slice. On failure, execute
//     them in reverse order to undo completed steps.
//   - Child workflows: encapsulate a sub-process with its own retry policy,
//     timeout, and signal handling (insurance verification here).
//   - workflow.Selector: multiplexes multiple futures and channels (like select
//     for Temporal). Used to race a timer against a signal channel.
//   - ActivityOptions: configure per-activity timeouts and retry policies.
//     Validation needs short timeouts; insurance verification needs long ones.

// Task queue name. Workers listen on this queue for workflow and activity tasks.
const TaskQueuePatientOnboarding = "patient-onboarding"

// Signal and query channel names (constants prevent typos).
const (
	SignalDocumentUploaded  = "document_uploaded"
	SignalInsuranceResponse = "insurance_response"
	QueryOnboardingStatus   = "onboarding_status"
)

// PatientOnboardingWorkflow orchestrates the full patient onboarding process.
//
// Steps (sequential):
//  1. Validate patient data
//  2. Create patient record in downstream service
//  3. Verify insurance (child workflow — may wait for async response)
//  4. Schedule appointment
//  5. Route prescriptions (one per rx)
//  6. Collect documents (wait for signals from patient portal)
//  7. Send welcome notification
//
// On failure at any step, the saga compensations run in reverse order to
// undo all previously completed steps.
func PatientOnboardingWorkflow(ctx workflow.Context, req OnboardingRequest) (*OnboardingState, error) {
	logger := workflow.GetLogger(ctx)

	// Initialize workflow state.
	state := &OnboardingState{
		Status:            StatusPending,
		DocumentsRequired: countRequiredDocs(req.Documents),
		StartedAt:         workflow.Now(ctx), // deterministic — not time.Now()
	}

	// --- Register query handler ---
	// Callers can query "onboarding_status" at any time to get the current
	// state without blocking the workflow. This powers the GET /status endpoint.
	err := workflow.SetQueryHandler(ctx, QueryOnboardingStatus, func() (*OnboardingState, error) {
		return state, nil
	})
	if err != nil {
		return nil, fmt.Errorf("registering query handler: %w", err)
	}

	// --- Saga compensation tracking ---
	// As each step succeeds, we push a compensation function onto this slice.
	// On failure, we pop and execute them in reverse order (LIFO).
	var compensations []func(workflow.Context) error

	// runCompensations executes all registered compensations in reverse.
	// Uses short timeouts — compensations should be fast cleanup operations.
	runCompensations := func(ctx workflow.Context) {
		compensationCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			StartToCloseTimeout: 30 * time.Second,
			RetryPolicy: &temporal.RetryPolicy{
				InitialInterval: time.Second,
				MaximumAttempts: 3,
			},
		})
		for i := len(compensations) - 1; i >= 0; i-- {
			if err := compensations[i](compensationCtx); err != nil {
				// Log but don't fail — best-effort cleanup.
				logger.Error("compensation failed", "index", i, "error", err)
			}
		}
	}

	// Helper to transition state and set the current step description.
	transition := func(next OnboardingStatus, step string) error {
		if err := state.TransitionTo(next); err != nil {
			return err
		}
		state.CurrentStep = step
		return nil
	}

	// --- Step 1: Validate patient data ---
	// Short timeout, no retries — validation is local and fast.
	transition(StatusValidating, "Validating patient data")
	validateCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1, // no retries for validation
		},
	})

	var activities *PatientActivities
	err = workflow.ExecuteActivity(validateCtx, activities.ValidatePatientData, req).Get(ctx, nil)
	if err != nil {
		state.TransitionTo(StatusFailed)
		state.FailureReason = fmt.Sprintf("validation failed: %v", err)
		return state, err
	}

	// --- Step 2: Create patient record ---
	transition(StatusCreatingRecord, "Creating patient record")
	createCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    30 * time.Second,
			MaximumAttempts:    5,
		},
	})

	err = workflow.ExecuteActivity(createCtx, activities.CreatePatientRecord, req.Patient).Get(ctx, &state.PatientID)
	if err != nil {
		state.TransitionTo(StatusFailed)
		state.FailureReason = fmt.Sprintf("failed to create patient record: %v", err)
		return state, err
	}

	// Register compensation: delete the patient record if a later step fails.
	compensations = append(compensations, func(ctx workflow.Context) error {
		return workflow.ExecuteActivity(ctx, activities.DeletePatientRecord, state.PatientID).Get(ctx, nil)
	})

	// --- Step 3: Verify insurance (child workflow) ---
	// Child workflow has its own lifecycle — it can wait for an async signal
	// from the insurance company (up to 48 hours) independently.
	transition(StatusVerifyingInsurance, "Verifying insurance eligibility")
	childCtx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
		WorkflowID:          fmt.Sprintf("insurance-verify-%s", state.PatientID),
		TaskQueue:           TaskQueuePatientOnboarding,
		WorkflowRunTimeout:  72 * time.Hour, // insurance can take days
		ParentClosePolicy:   1,              // TERMINATE child if parent completes
	})

	var insuranceResult InsuranceVerificationResult
	err = workflow.ExecuteChildWorkflow(childCtx, InsuranceVerificationChildWorkflow, req.Insurance).Get(ctx, &insuranceResult)
	if err != nil {
		state.TransitionTo(StatusFailed)
		state.FailureReason = fmt.Sprintf("insurance verification failed: %v", err)
		runCompensations(ctx)
		return state, err
	}
	state.InsuranceResult = &insuranceResult

	if !insuranceResult.Verified {
		state.TransitionTo(StatusFailed)
		state.FailureReason = "insurance verification denied"
		runCompensations(ctx)
		return state, fmt.Errorf("insurance not verified")
	}

	// --- Step 4: Schedule appointment ---
	transition(StatusSchedulingAppointment, "Scheduling appointment")
	scheduleCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    30 * time.Second,
			MaximumAttempts:    5,
		},
	})

	err = workflow.ExecuteActivity(scheduleCtx, activities.ScheduleAppointment, state.PatientID, req.Appointment).Get(ctx, &state.AppointmentID)
	if err != nil {
		state.TransitionTo(StatusFailed)
		state.FailureReason = fmt.Sprintf("failed to schedule appointment: %v", err)
		runCompensations(ctx)
		return state, err
	}

	// Compensation: cancel the appointment.
	compensations = append(compensations, func(ctx workflow.Context) error {
		return workflow.ExecuteActivity(ctx, activities.CancelAppointment, state.AppointmentID).Get(ctx, nil)
	})

	// --- Step 5: Route prescriptions ---
	transition(StatusRoutingPrescription, "Routing prescriptions")
	rxCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    30 * time.Second,
			MaximumAttempts:    5,
		},
	})

	for _, rx := range req.Prescriptions {
		var prescriptionID string
		err = workflow.ExecuteActivity(rxCtx, activities.RoutePrescription, state.PatientID, rx).Get(ctx, &prescriptionID)
		if err != nil {
			state.TransitionTo(StatusFailed)
			state.FailureReason = fmt.Sprintf("failed to route prescription %s: %v", rx.MedicationName, err)
			runCompensations(ctx)
			return state, err
		}
		state.PrescriptionIDs = append(state.PrescriptionIDs, prescriptionID)

		// Compensation: cancel this prescription.
		rxID := prescriptionID // capture loop variable
		compensations = append(compensations, func(ctx workflow.Context) error {
			return workflow.ExecuteActivity(ctx, activities.CancelPrescription, rxID).Get(ctx, nil)
		})
	}

	// --- Step 6: Collect documents (signal-driven) ---
	transition(StatusCollectingDocuments, "Waiting for document uploads")
	docCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval: time.Second,
			MaximumAttempts: 3,
		},
	})

	// Initialize document collection in the document service.
	err = workflow.ExecuteActivity(docCtx, activities.CollectDocuments, state.PatientID, req.Documents).Get(ctx, nil)
	if err != nil {
		state.TransitionTo(StatusFailed)
		state.FailureReason = fmt.Sprintf("failed to initialize document collection: %v", err)
		runCompensations(ctx)
		return state, err
	}

	// Wait for document upload signals. The patient portal sends a signal
	// each time a document is uploaded. We wait until all required documents
	// are collected or a 7-day timeout expires.
	docChannel := workflow.GetSignalChannel(ctx, SignalDocumentUploaded)
	docTimeout := workflow.NewTimer(ctx, 7*24*time.Hour)

	for state.DocumentsCollected < state.DocumentsRequired {
		selector := workflow.NewSelector(ctx)

		// Option 1: document uploaded signal arrives.
		selector.AddReceive(docChannel, func(c workflow.ReceiveChannel, more bool) {
			var signal DocumentUploadedSignal
			c.Receive(ctx, &signal)
			state.DocumentsCollected++
			logger.Info("document uploaded",
				"type", signal.DocumentType,
				"collected", state.DocumentsCollected,
				"required", state.DocumentsRequired,
			)
		})

		// Option 2: 7-day timeout fires — proceed without all documents.
		selector.AddFuture(docTimeout, func(f workflow.Future) {
			logger.Warn("document collection timed out, proceeding with available documents")
			// Set collected = required to break the loop.
			state.DocumentsCollected = state.DocumentsRequired
		})

		selector.Select(ctx)
	}

	// --- Step 7: Send welcome notification ---
	transition(StatusSendingNotification, "Sending welcome notification")
	notifyCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval: time.Second,
			MaximumAttempts: 3,
		},
	})

	message := fmt.Sprintf("Welcome %s! Your onboarding is complete. "+
		"Appointment scheduled for %s.",
		req.Patient.FirstName,
		req.Appointment.ScheduledAt.Format("Jan 2, 2006 at 3:04 PM"))

	err = workflow.ExecuteActivity(notifyCtx, activities.SendNotification,
		state.PatientID, "onboarding_complete", message).Get(ctx, nil)
	if err != nil {
		// Notification failure is non-critical — don't roll back the entire onboarding.
		logger.Error("failed to send notification", "error", err)
	}

	// --- Complete ---
	state.TransitionTo(StatusCompleted)
	state.CurrentStep = "Onboarding complete"
	now := workflow.Now(ctx)
	state.CompletedAt = &now

	logger.Info("patient onboarding completed",
		"patient_id", state.PatientID,
		"duration", now.Sub(state.StartedAt),
	)

	return state, nil
}

// InsuranceVerificationChildWorkflow handles the insurance verification
// sub-process as a child workflow.
//
// Why a child workflow instead of just an activity?
//   - The insurance company may respond asynchronously (hours or days later).
//     An activity would need to poll or block for that entire duration, tying
//     up a worker slot.
//   - A child workflow can wait on a signal channel with a 48-hour timer
//     without consuming worker resources. Temporal manages the wait.
//   - The child workflow has its own WorkflowID, making it independently
//     queryable and cancellable.
func InsuranceVerificationChildWorkflow(ctx workflow.Context, insurance InsuranceInfo) (*InsuranceVerificationResult, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("starting insurance verification", "provider", insurance.ProviderName)

	// Call the insurance verification API.
	verifyCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 60 * time.Second,
		HeartbeatTimeout:    15 * time.Second, // activity must heartbeat every 15s
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    2 * time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    60 * time.Second,
			MaximumAttempts:    5,
		},
	})

	var activities *PatientActivities
	var initialResult InsuranceVerificationResult
	err := workflow.ExecuteActivity(verifyCtx, activities.VerifyInsurance, insurance).Get(ctx, &initialResult)
	if err != nil {
		return nil, fmt.Errorf("insurance verification activity failed: %w", err)
	}

	// If the API returned an immediate result, we're done.
	if initialResult.Verified {
		return &initialResult, nil
	}

	// Otherwise, wait for an async response signal from the insurance company.
	// Use a Selector to race the signal against a 48-hour timeout.
	//
	// workflow.Selector is like Go's select statement but for Temporal:
	// it multiplexes futures (timers, activity results) and channels (signals).
	logger.Info("waiting for insurance response signal (48h timeout)")

	signalCh := workflow.GetSignalChannel(ctx, SignalInsuranceResponse)
	timer := workflow.NewTimer(ctx, 48*time.Hour)

	var result InsuranceVerificationResult
	timedOut := false

	selector := workflow.NewSelector(ctx)

	// Branch 1: insurance response signal arrives.
	selector.AddReceive(signalCh, func(c workflow.ReceiveChannel, more bool) {
		var signal InsuranceResponseSignal
		c.Receive(ctx, &signal)

		result = InsuranceVerificationResult{
			Verified:     signal.Approved,
			EligibleFrom: signal.EligibleFrom,
			EligibleTo:   signal.EligibleTo,
			Copay:        signal.Copay,
			Deductible:   signal.Deductible,
			Notes:        signal.DenialReason,
		}
	})

	// Branch 2: 48-hour timeout fires.
	selector.AddFuture(timer, func(f workflow.Future) {
		timedOut = true
	})

	// Block until one branch fires.
	selector.Select(ctx)

	if timedOut {
		return nil, temporal.NewNonRetryableApplicationError(
			"insurance verification timed out after 48 hours",
			"INSURANCE_TIMEOUT",
			nil,
		)
	}

	return &result, nil
}

// countRequiredDocs counts documents marked as required.
func countRequiredDocs(docs []DocumentRequirement) int {
	count := 0
	for _, d := range docs {
		if d.Required {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("=== MSO Scheduler — Workflow Orchestration ===\n")

	fmt.Println("PatientOnboardingWorkflow execution plan:")
	fmt.Println("  1. ValidatePatientData       — 10s timeout, no retries")
	fmt.Println("  2. CreatePatientRecord        — 30s timeout, 5 retries, exponential backoff")
	fmt.Println("  3. InsuranceVerification      — child workflow, 72h run timeout")
	fmt.Println("     └─ VerifyInsurance         — 60s timeout, 15s heartbeat, 5 retries")
	fmt.Println("     └─ Wait for signal         — 48h timeout via workflow.Selector")
	fmt.Println("  4. ScheduleAppointment        — 30s timeout, 5 retries")
	fmt.Println("  5. RoutePrescription (per rx)  — 30s timeout, 5 retries")
	fmt.Println("  6. CollectDocuments           — 30s timeout, then wait for signals (7d)")
	fmt.Println("  7. SendNotification           — 30s timeout, 3 retries (non-critical)")

	fmt.Println("\nSaga compensations (reverse order on failure):")
	fmt.Println("  CancelPrescription (per rx)")
	fmt.Println("  CancelAppointment")
	fmt.Println("  DeletePatientRecord")

	fmt.Println("\nSignal channels:")
	fmt.Printf("  %s — patient uploaded a document\n", SignalDocumentUploaded)
	fmt.Printf("  %s — insurance company responded\n", SignalInsuranceResponse)

	fmt.Println("\nQuery handlers:")
	fmt.Printf("  %s — returns current OnboardingState\n", QueryOnboardingStatus)

	fmt.Println("\nTemporal determinism rules applied:")
	fmt.Println("  - workflow.Now(ctx) instead of time.Now()")
	fmt.Println("  - workflow.NewTimer() instead of time.Sleep()")
	fmt.Println("  - workflow.GetSignalChannel() for async events")
	fmt.Println("  - workflow.Selector for multiplexing futures and channels")
	fmt.Println("  - No goroutines — use workflow.Go() if needed")
	fmt.Println("  - No direct I/O — all side effects happen in activities")
}
