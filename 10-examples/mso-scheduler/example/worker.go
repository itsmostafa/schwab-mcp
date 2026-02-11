package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

// Temporal worker setup for the patient onboarding scheduler.
//
// Key interview concepts:
//   - client.Dial creates a gRPC connection to the Temporal server. The
//     client is shared across the worker and API server.
//   - worker.New creates a worker that polls a specific task queue for
//     both workflow and activity tasks.
//   - Struct-based activity registration: registering &PatientActivities{}
//     makes all exported methods on the struct available as activities.
//     Dependencies (HTTP client, service URLs) are injected via the struct
//     fields — no globals.
//   - MaxConcurrentActivityExecutionSize limits parallel activity goroutines
//     to prevent overwhelming downstream services.
//   - MaxConcurrentWorkflowTaskExecutionSize limits parallel workflow task
//     processing. Workflow tasks are CPU-bound (replay), so this should be
//     tuned to available cores.
//   - Graceful shutdown: the worker's InterruptCh receives OS signals. On
//     SIGTERM/SIGINT, the worker stops accepting new tasks, finishes in-progress
//     ones, then exits cleanly.

func main() {
	fmt.Println("=== MSO Scheduler — Temporal Worker ===\n")

	// --- Create Temporal client ---
	// client.Dial establishes a gRPC connection to the Temporal server.
	// Options:
	//   HostPort:  Temporal server address (default: localhost:7233)
	//   Namespace: logical isolation within a Temporal cluster ("default" is fine for dev)
	//   Logger:    plug in a structured logger (zap, zerolog) for production
	c, err := client.Dial(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "default",
	})
	if err != nil {
		log.Fatal("unable to create Temporal client:", err)
	}
	defer c.Close()
	fmt.Println("Connected to Temporal server at localhost:7233")

	// --- Create worker ---
	// The worker polls the "patient-onboarding" task queue for both workflow
	// tasks (workflow step execution / replay) and activity tasks (side-effect
	// execution like HTTP calls).
	//
	// Concurrency tuning:
	//   MaxConcurrentActivityExecutionSize: 10 — at most 10 activities run in
	//     parallel on this worker. Prevents overwhelming downstream services.
	//     In production, tune this based on downstream rate limits and worker
	//     resources (CPU, memory, file descriptors).
	//   MaxConcurrentWorkflowTaskExecutionSize: 5 — at most 5 workflow tasks
	//     replay simultaneously. Workflow tasks are CPU-bound (event replay),
	//     so this should roughly match available cores.
	w := worker.New(c, TaskQueuePatientOnboarding, worker.Options{
		MaxConcurrentActivityExecutionSize:     10,
		MaxConcurrentWorkflowTaskExecutionSize: 5,
	})

	// --- Register workflows ---
	// Both the main workflow and child workflow are registered on the same
	// worker since they share the same task queue.
	w.RegisterWorkflow(PatientOnboardingWorkflow)
	w.RegisterWorkflow(InsuranceVerificationChildWorkflow)

	// --- Register activities (struct-based) ---
	// Registering a struct makes all its exported methods available as activities.
	// The struct fields (service URLs, HTTP client) serve as dependency injection.
	//
	// Advantages over function-based registration:
	//   - Dependencies are injected once at worker startup, not per-activity call.
	//   - No global variables needed.
	//   - Easy to swap implementations in tests (inject a mock struct).
	activities := &PatientActivities{
		PatientServiceURL:      getEnvOrDefault("PATIENT_SERVICE_URL", "http://localhost:8081"),
		InsuranceAPIURL:        getEnvOrDefault("INSURANCE_API_URL", "http://localhost:8082"),
		SchedulingServiceURL:   getEnvOrDefault("SCHEDULING_SERVICE_URL", "http://localhost:8083"),
		PharmacyServiceURL:     getEnvOrDefault("PHARMACY_SERVICE_URL", "http://localhost:8084"),
		DocumentServiceURL:     getEnvOrDefault("DOCUMENT_SERVICE_URL", "http://localhost:8085"),
		NotificationServiceURL: getEnvOrDefault("NOTIFICATION_SERVICE_URL", "http://localhost:8086"),
	}
	w.RegisterActivity(activities)

	fmt.Println("Registered workflows:")
	fmt.Println("  PatientOnboardingWorkflow")
	fmt.Println("  InsuranceVerificationChildWorkflow")
	fmt.Println("Registered activities:")
	fmt.Println("  PatientActivities (struct-based: all exported methods)")
	fmt.Printf("Task queue: %s\n", TaskQueuePatientOnboarding)
	fmt.Println("Activity concurrency: 10")
	fmt.Println("Workflow task concurrency: 5")

	// --- Graceful shutdown ---
	// worker.InterruptCh() returns a channel that receives OS signals
	// (SIGINT, SIGTERM). When a signal arrives:
	//   1. The worker stops polling for new tasks.
	//   2. In-progress workflow tasks and activities continue to completion.
	//   3. Once all in-progress tasks finish, the worker exits.
	//
	// This prevents:
	//   - Partially executed activities (e.g., half-sent HTTP request).
	//   - Unnecessary activity retries on clean deployments.
	//   - Lost workflow task progress (the task goes back to the queue on
	//     ungraceful shutdown, causing redundant replay).
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sig := <-quit
		fmt.Printf("\nReceived %s, initiating graceful shutdown...\n", sig)
		// worker.Stop is called by the Run method when it detects the interrupt.
	}()

	fmt.Println("\nWorker started. Press Ctrl+C for graceful shutdown.")

	// worker.Run blocks until interrupted (via InterruptCh or direct Stop call).
	// It handles the graceful shutdown lifecycle internally.
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatal("worker failed:", err)
	}

	fmt.Println("Worker stopped gracefully.")
}

// getEnvOrDefault reads an environment variable or returns a default value.
func getEnvOrDefault(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
