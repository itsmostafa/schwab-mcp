package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
)

// HTTP API for starting, querying, and signaling patient onboarding workflows.
//
// Key interview concepts:
//   - WorkflowID design: hash of patient email ensures one onboarding per
//     patient. If the same patient is submitted twice, Temporal returns
//     WorkflowExecutionAlreadyStartedError instead of creating a duplicate.
//   - client.QueryWorkflow: synchronous read of workflow state via the
//     registered query handler. Returns the current OnboardingState without
//     blocking the workflow.
//   - client.SignalWorkflow: sends an async event to a running workflow.
//     Used for document uploads and insurance responses that arrive
//     externally.
//   - The API server and worker can be separate processes. They share the
//     Temporal client but serve different roles: the API accepts HTTP requests,
//     the worker executes workflow/activity code.

// OnboardingAPI holds the Temporal client and task queue name.
// Injected at startup — no globals.
type OnboardingAPI struct {
	TemporalClient client.Client
	TaskQueue      string
}

// StartOnboardingRequest is the JSON body for POST /api/v1/patients/onboard.
type StartOnboardingRequest struct {
	Patient       Patient             `json:"patient" binding:"required"`
	Insurance     InsuranceInfo       `json:"insurance" binding:"required"`
	Intake        IntakeQuestionnaire `json:"intake"`
	Prescriptions []Prescription      `json:"prescriptions"`
	Appointment   Appointment         `json:"appointment"`
	Documents     []DocumentRequirement `json:"documents"`
}

// HandleStartOnboarding starts a new patient onboarding workflow.
//
// POST /api/v1/patients/onboard
//
// WorkflowID = sha256(patient.email) — ensures exactly one onboarding
// workflow per patient email. If the workflow already exists, Temporal
// returns an error and we return 409 Conflict.
func (api *OnboardingAPI) HandleStartOnboarding(c *gin.Context) {
	var req StartOnboardingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Deterministic workflow ID from patient email.
	// This prevents duplicate onboarding workflows for the same patient.
	workflowID := fmt.Sprintf("patient-onboard-%x",
		sha256.Sum256([]byte(req.Patient.Email)))

	// Build the workflow input.
	onboardingReq := OnboardingRequest{
		Patient:       req.Patient,
		Insurance:     req.Insurance,
		Intake:        req.Intake,
		Prescriptions: req.Prescriptions,
		Appointment:   req.Appointment,
		Documents:     req.Documents,
	}

	// Start the workflow.
	// WorkflowIDReusePolicy is not set — default is AllowDuplicateFailedOnly,
	// meaning a new run is only allowed if the previous one failed.
	// If a workflow with this ID is already running, ExecuteWorkflow returns
	// a WorkflowExecutionAlreadyStartedError.
	opts := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: api.TaskQueue,
	}

	run, err := api.TemporalClient.ExecuteWorkflow(c.Request.Context(), opts,
		PatientOnboardingWorkflow, onboardingReq)
	if err != nil {
		// Check if workflow already exists.
		if isAlreadyStartedError(err) {
			c.JSON(http.StatusConflict, gin.H{
				"error":       "onboarding already in progress for this patient",
				"workflow_id": workflowID,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"workflow_id": run.GetID(),
		"run_id":      run.GetRunID(),
		"message":     "onboarding workflow started",
	})
}

// HandleGetStatus queries the current onboarding state of a running workflow.
//
// GET /api/v1/patients/:workflowID/status
//
// Uses client.QueryWorkflow to invoke the "onboarding_status" query handler
// registered in the workflow. This is a synchronous read — it returns the
// current state immediately without blocking the workflow execution.
func (api *OnboardingAPI) HandleGetStatus(c *gin.Context) {
	workflowID := c.Param("workflowID")

	// Query the workflow's registered query handler.
	// The empty string for runID means "latest run of this workflow ID".
	resp, err := api.TemporalClient.QueryWorkflow(c.Request.Context(),
		workflowID, "", QueryOnboardingStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var state OnboardingState
	if err := resp.Get(&state); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, state)
}

// HandleDocumentUploaded signals a running workflow that a document was uploaded.
//
// POST /api/v1/patients/:workflowID/documents
//
// Uses client.SignalWorkflow to send a DocumentUploadedSignal to the workflow's
// "document_uploaded" channel. The workflow is blocked on this channel during
// the document collection phase.
func (api *OnboardingAPI) HandleDocumentUploaded(c *gin.Context) {
	workflowID := c.Param("workflowID")

	var signal DocumentUploadedSignal
	if err := c.ShouldBindJSON(&signal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// SignalWorkflow sends the signal asynchronously. The workflow receives
	// it on its signal channel. Empty runID = latest run.
	err := api.TemporalClient.SignalWorkflow(c.Request.Context(),
		workflowID, "", SignalDocumentUploaded, signal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "document upload signal sent",
		"document_type": signal.DocumentType,
	})
}

// HandleInsuranceResponse signals a running workflow with the insurance
// company's verification response.
//
// POST /api/v1/patients/:workflowID/insurance-response
//
// This signal is received by the InsuranceVerificationChildWorkflow,
// which is waiting on the "insurance_response" channel with a 48-hour timeout.
func (api *OnboardingAPI) HandleInsuranceResponse(c *gin.Context) {
	workflowID := c.Param("workflowID")

	var signal InsuranceResponseSignal
	if err := c.ShouldBindJSON(&signal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := api.TemporalClient.SignalWorkflow(c.Request.Context(),
		workflowID, "", SignalInsuranceResponse, signal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "insurance response signal sent",
		"approved": signal.Approved,
	})
}

// isAlreadyStartedError checks if the error indicates a workflow with
// this ID is already running. In production, use the Temporal SDK's
// typed error checking.
func isAlreadyStartedError(err error) bool {
	// The Temporal SDK returns *serviceerror.WorkflowExecutionAlreadyStarted.
	// For simplicity, we check the error string here. In production, use
	// type assertion against the specific error type.
	return err != nil && (fmt.Sprintf("%v", err) != "" &&
		contains(err.Error(), "already started"))
}

// contains checks if s contains substr (simple helper to avoid importing strings).
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("=== MSO Scheduler — HTTP API Server ===\n")

	// --- Create Temporal client ---
	c, err := client.Dial(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "default",
	})
	if err != nil {
		log.Fatal("unable to create Temporal client:", err)
	}
	defer c.Close()

	// --- Create API handler ---
	api := &OnboardingAPI{
		TemporalClient: c,
		TaskQueue:      TaskQueuePatientOnboarding,
	}

	// --- Gin router setup ---
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	// Patient onboarding endpoints.
	v1 := r.Group("/api/v1/patients")
	{
		v1.POST("/onboard", api.HandleStartOnboarding)
		v1.GET("/:workflowID/status", api.HandleGetStatus)
		v1.POST("/:workflowID/documents", api.HandleDocumentUploaded)
		v1.POST("/:workflowID/insurance-response", api.HandleInsuranceResponse)
	}

	// Health check.
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	fmt.Println("API endpoints:")
	fmt.Println("  POST /api/v1/patients/onboard              — start onboarding workflow")
	fmt.Println("  GET  /api/v1/patients/:workflowID/status   — query onboarding status")
	fmt.Println("  POST /api/v1/patients/:workflowID/documents — signal document uploaded")
	fmt.Println("  POST /api/v1/patients/:workflowID/insurance-response — signal insurance response")
	fmt.Println("  GET  /health                                — health check")

	fmt.Println("\nStarting API server on :8080")

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("server error:", err)
	}
}
