package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/temporal"
)

// Activity implementations for patient onboarding.
//
// Key interview concepts:
//   - Activities are the only place where side effects (HTTP calls, DB writes) happen.
//     Workflows must be deterministic; activities don't need to be.
//   - Each activity uses an idempotency key (hash of stable inputs) so that retries
//     from Temporal don't create duplicate records in downstream services.
//   - activity.RecordHeartbeat signals to the Temporal server that a long-running
//     activity is still alive. If heartbeats stop, the server considers the activity
//     failed and schedules a retry.
//   - NonRetryableApplicationError tells Temporal to stop retrying — the failure
//     is permanent (e.g., invalid input, denied insurance).
//   - Struct-based activity registration (registering PatientActivities as a struct)
//     lets Temporal inject dependencies via the receiver, avoiding globals.

// PatientActivities holds downstream service URLs and an HTTP client.
// Registered as a struct with the Temporal worker so that all methods become
// activities with injected dependencies — no globals needed.
type PatientActivities struct {
	PatientServiceURL    string
	InsuranceAPIURL      string
	SchedulingServiceURL string
	PharmacyServiceURL   string
	DocumentServiceURL   string
	NotificationServiceURL string
	HTTPClient           *http.Client
}

// ValidatePatientData performs local validation of the intake form.
// This is a pure function — no external calls — so it's fast and does not
// need heartbeats or idempotency.
//
// Returns a NonRetryableApplicationError for invalid data because retrying
// with the same bad input will always fail.
func (a *PatientActivities) ValidatePatientData(ctx context.Context, req OnboardingRequest) error {
	log.Printf("[activity] ValidatePatientData: %s %s", req.Patient.FirstName, req.Patient.LastName)

	if req.Patient.Email == "" {
		return temporal.NewNonRetryableApplicationError(
			"patient email is required",
			"VALIDATION_ERROR",
			nil,
		)
	}
	if req.Patient.FirstName == "" || req.Patient.LastName == "" {
		return temporal.NewNonRetryableApplicationError(
			"patient first and last name are required",
			"VALIDATION_ERROR",
			nil,
		)
	}
	if req.Insurance.MemberID == "" {
		return temporal.NewNonRetryableApplicationError(
			"insurance member ID is required",
			"VALIDATION_ERROR",
			nil,
		)
	}
	if req.Patient.DateOfBirth.After(time.Now()) {
		return temporal.NewNonRetryableApplicationError(
			"date of birth cannot be in the future",
			"VALIDATION_ERROR",
			nil,
		)
	}

	log.Printf("[activity] ValidatePatientData: passed")
	return nil
}

// CreatePatientRecord sends a POST to the Patient DB Service to create
// the patient record.
//
// Idempotency: uses a hash of email + DOB as the Idempotency-Key header.
// If the downstream service receives the same key twice (due to Temporal
// retry), it returns the existing record instead of creating a duplicate.
func (a *PatientActivities) CreatePatientRecord(ctx context.Context, patient Patient) (string, error) {
	log.Printf("[activity] CreatePatientRecord: %s %s", patient.FirstName, patient.LastName)

	// Compute idempotency key from stable patient identifiers.
	// email + DOB uniquely identifies a patient for onboarding purposes.
	idempotencyKey := computeIdempotencyKey(patient.Email, patient.DateOfBirth.Format("2006-01-02"))

	body, _ := json.Marshal(patient)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		a.PatientServiceURL+"/api/v1/patients", strings.NewReader(string(body)))
	if err != nil {
		return "", fmt.Errorf("building request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", idempotencyKey)

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		// Network error — retryable by default.
		return "", fmt.Errorf("calling patient service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusConflict {
		// 409 means the idempotency key matched an existing request.
		// Extract the patient ID from the response and return it.
		var result struct{ PatientID string `json:"patient_id"` }
		json.NewDecoder(resp.Body).Decode(&result)
		log.Printf("[activity] CreatePatientRecord: idempotent hit, existing patient %s", result.PatientID)
		return result.PatientID, nil
	}

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("patient service returned %d", resp.StatusCode)
	}

	var result struct{ PatientID string `json:"patient_id"` }
	json.NewDecoder(resp.Body).Decode(&result)
	log.Printf("[activity] CreatePatientRecord: created %s", result.PatientID)
	return result.PatientID, nil
}

// VerifyInsurance sends the insurance details to the eligibility API.
//
// This activity uses RecordHeartbeat because the insurance API can be slow
// (up to several seconds per request). Without heartbeats, the Temporal
// server might assume the activity is dead and schedule a retry, leading
// to duplicate verification requests.
//
// The activity distinguishes between transient errors (network issues,
// 5xx responses) which Temporal should retry, and permanent errors
// (invalid member ID, denied coverage) which should not be retried.
func (a *PatientActivities) VerifyInsurance(ctx context.Context, insurance InsuranceInfo) (*InsuranceVerificationResult, error) {
	log.Printf("[activity] VerifyInsurance: %s (member: %s)", insurance.ProviderName, insurance.MemberID)

	// Heartbeat before the potentially slow API call.
	// The Temporal server uses heartbeats to detect stuck activities.
	// If the server doesn't receive a heartbeat within the HeartbeatTimeout
	// configured in ActivityOptions, it marks the activity as failed.
	activity.RecordHeartbeat(ctx, "sending verification request")

	body, _ := json.Marshal(insurance)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		a.InsuranceAPIURL+"/api/v1/verify", strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		// Network error — transient, Temporal will retry.
		return nil, fmt.Errorf("calling insurance API: %w", err)
	}
	defer resp.Body.Close()

	activity.RecordHeartbeat(ctx, "processing verification response")

	// Distinguish transient vs permanent failures.
	switch {
	case resp.StatusCode == http.StatusOK:
		var result InsuranceVerificationResult
		json.NewDecoder(resp.Body).Decode(&result)
		return &result, nil

	case resp.StatusCode == http.StatusUnprocessableEntity:
		// 422 = invalid member ID or plan. Permanent failure — don't retry.
		return nil, temporal.NewNonRetryableApplicationError(
			"insurance verification failed: invalid member ID or plan",
			"INSURANCE_INVALID",
			nil,
		)

	case resp.StatusCode >= 500:
		// 5xx = server-side issue. Transient — Temporal will retry with backoff.
		return nil, fmt.Errorf("insurance API server error: %d", resp.StatusCode)

	default:
		return nil, fmt.Errorf("insurance API unexpected status: %d", resp.StatusCode)
	}
}

// ScheduleAppointment creates an appointment in the scheduling service.
// Idempotency key: hash of patientID + appointment type.
func (a *PatientActivities) ScheduleAppointment(ctx context.Context, patientID string, appt Appointment) (string, error) {
	log.Printf("[activity] ScheduleAppointment: type=%s provider=%s", appt.Type, appt.ProviderID)

	idempotencyKey := computeIdempotencyKey(patientID, appt.Type)

	payload := map[string]interface{}{
		"patient_id":   patientID,
		"type":         appt.Type,
		"provider_id":  appt.ProviderID,
		"scheduled_at": appt.ScheduledAt,
		"location":     appt.Location,
		"notes":        appt.Notes,
	}
	body, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		a.SchedulingServiceURL+"/api/v1/appointments", strings.NewReader(string(body)))
	if err != nil {
		return "", fmt.Errorf("building request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", idempotencyKey)

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("calling scheduling service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusConflict {
		return "", fmt.Errorf("scheduling service returned %d", resp.StatusCode)
	}

	var result struct{ AppointmentID string `json:"appointment_id"` }
	json.NewDecoder(resp.Body).Decode(&result)
	log.Printf("[activity] ScheduleAppointment: %s", result.AppointmentID)
	return result.AppointmentID, nil
}

// RoutePrescription sends a single prescription to the pharmacy service.
// Called once per prescription in the onboarding request.
// Idempotency key: hash of patientID + medication name.
func (a *PatientActivities) RoutePrescription(ctx context.Context, patientID string, rx Prescription) (string, error) {
	log.Printf("[activity] RoutePrescription: %s %s", rx.MedicationName, rx.Dosage)

	idempotencyKey := computeIdempotencyKey(patientID, rx.MedicationName)

	payload := map[string]interface{}{
		"patient_id":      patientID,
		"medication_name": rx.MedicationName,
		"dosage":          rx.Dosage,
		"frequency":       rx.Frequency,
		"prescriber":      rx.Prescriber,
		"pharmacy_id":     rx.PharmacyID,
	}
	body, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		a.PharmacyServiceURL+"/api/v1/prescriptions", strings.NewReader(string(body)))
	if err != nil {
		return "", fmt.Errorf("building request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", idempotencyKey)

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("calling pharmacy service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusConflict {
		return "", fmt.Errorf("pharmacy service returned %d", resp.StatusCode)
	}

	var result struct{ PrescriptionID string `json:"prescription_id"` }
	json.NewDecoder(resp.Body).Decode(&result)
	log.Printf("[activity] RoutePrescription: %s", result.PrescriptionID)
	return result.PrescriptionID, nil
}

// CollectDocuments initializes the document requirements list in the
// document management service. The patient will upload documents via
// the portal, which sends signals to the workflow.
func (a *PatientActivities) CollectDocuments(ctx context.Context, patientID string, docs []DocumentRequirement) error {
	log.Printf("[activity] CollectDocuments: %d documents for patient %s", len(docs), patientID)

	payload := map[string]interface{}{
		"patient_id": patientID,
		"documents":  docs,
	}
	body, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		a.DocumentServiceURL+"/api/v1/document-requirements", strings.NewReader(string(body)))
	if err != nil {
		return fmt.Errorf("building request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("calling document service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("document service returned %d", resp.StatusCode)
	}

	log.Printf("[activity] CollectDocuments: initialized")
	return nil
}

// SendNotification sends a notification (email/SMS) to the patient.
// Idempotency key: hash of notification type + patientID + time bucket.
// The time bucket (rounded to the hour) prevents the same notification
// from being sent twice within the same hour during retries, while still
// allowing re-sends in later hours if needed.
func (a *PatientActivities) SendNotification(ctx context.Context, patientID, notificationType, message string) error {
	log.Printf("[activity] SendNotification: type=%s patient=%s", notificationType, patientID)

	// Time bucket: round to the current hour.
	timeBucket := time.Now().UTC().Truncate(time.Hour).Format("2006-01-02T15")
	idempotencyKey := computeIdempotencyKey(notificationType, patientID, timeBucket)

	payload := map[string]interface{}{
		"patient_id": patientID,
		"type":       notificationType,
		"message":    message,
		"channels":   []string{"email", "sms"},
	}
	body, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		a.NotificationServiceURL+"/api/v1/notifications", strings.NewReader(string(body)))
	if err != nil {
		return fmt.Errorf("building request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", idempotencyKey)

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("calling notification service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusConflict {
		return fmt.Errorf("notification service returned %d", resp.StatusCode)
	}

	log.Printf("[activity] SendNotification: sent")
	return nil
}

// --- Compensation activities ---
// These are called during saga rollback when a later step fails.
// Each undoes the side effect of a forward activity.

// DeletePatientRecord removes a patient record created by CreatePatientRecord.
func (a *PatientActivities) DeletePatientRecord(ctx context.Context, patientID string) error {
	log.Printf("[activity] DeletePatientRecord (compensation): %s", patientID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete,
		fmt.Sprintf("%s/api/v1/patients/%s", a.PatientServiceURL, patientID), nil)
	if err != nil {
		return fmt.Errorf("building request: %w", err)
	}

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("calling patient service: %w", err)
	}
	defer resp.Body.Close()

	// 404 is acceptable — the record may not have been created if the
	// original activity failed after the network call but before returning.
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound {
		return fmt.Errorf("patient service returned %d during compensation", resp.StatusCode)
	}
	return nil
}

// CancelAppointment cancels an appointment created by ScheduleAppointment.
func (a *PatientActivities) CancelAppointment(ctx context.Context, appointmentID string) error {
	log.Printf("[activity] CancelAppointment (compensation): %s", appointmentID)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		fmt.Sprintf("%s/api/v1/appointments/%s/cancel", a.SchedulingServiceURL, appointmentID), nil)
	if err != nil {
		return fmt.Errorf("building request: %w", err)
	}

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("calling scheduling service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound {
		return fmt.Errorf("scheduling service returned %d during compensation", resp.StatusCode)
	}
	return nil
}

// CancelPrescription cancels a prescription routed by RoutePrescription.
func (a *PatientActivities) CancelPrescription(ctx context.Context, prescriptionID string) error {
	log.Printf("[activity] CancelPrescription (compensation): %s", prescriptionID)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		fmt.Sprintf("%s/api/v1/prescriptions/%s/cancel", a.PharmacyServiceURL, prescriptionID), nil)
	if err != nil {
		return fmt.Errorf("building request: %w", err)
	}

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("calling pharmacy service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound {
		return fmt.Errorf("pharmacy service returned %d during compensation", resp.StatusCode)
	}
	return nil
}

// --- Helpers ---

// computeIdempotencyKey produces a deterministic SHA-256 hash from the given
// parts. Used as the Idempotency-Key header in downstream HTTP calls.
//
// Why SHA-256 and not UUID?
//   - UUIDs are random — a new one each retry would defeat idempotency.
//   - A hash of stable inputs (email, DOB, medication name) produces the
//     same key on every retry, which is exactly what idempotent APIs need.
func computeIdempotencyKey(parts ...string) string {
	h := sha256.New()
	for _, p := range parts {
		h.Write([]byte(p))
		h.Write([]byte("|")) // separator to avoid collisions like "ab"+"c" vs "a"+"bc"
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	fmt.Println("=== MSO Scheduler — Activity Implementations ===\n")

	// Create a PatientActivities struct with downstream service URLs.
	// In production, these come from config or environment variables.
	activities := &PatientActivities{
		PatientServiceURL:      "http://patient-service:8080",
		InsuranceAPIURL:        "http://insurance-api:8080",
		SchedulingServiceURL:   "http://scheduling-service:8080",
		PharmacyServiceURL:     "http://pharmacy-service:8080",
		DocumentServiceURL:     "http://document-service:8080",
		NotificationServiceURL: "http://notification-service:8080",
		HTTPClient:             &http.Client{Timeout: 30 * time.Second},
	}

	fmt.Println("Downstream services configured:")
	fmt.Printf("  Patient:      %s\n", activities.PatientServiceURL)
	fmt.Printf("  Insurance:    %s\n", activities.InsuranceAPIURL)
	fmt.Printf("  Scheduling:   %s\n", activities.SchedulingServiceURL)
	fmt.Printf("  Pharmacy:     %s\n", activities.PharmacyServiceURL)
	fmt.Printf("  Documents:    %s\n", activities.DocumentServiceURL)
	fmt.Printf("  Notification: %s\n", activities.NotificationServiceURL)

	// Demonstrate idempotency key computation.
	fmt.Println("\n--- Idempotency Keys ---")
	key1 := computeIdempotencyKey("jane@example.com", "1990-03-15")
	key2 := computeIdempotencyKey("jane@example.com", "1990-03-15")
	key3 := computeIdempotencyKey("john@example.com", "1985-07-20")
	fmt.Printf("  email+DOB key (same input):  %s\n", key1[:16]+"...")
	fmt.Printf("  email+DOB key (same again):  %s\n", key2[:16]+"...")
	fmt.Printf("  email+DOB key (diff input):  %s\n", key3[:16]+"...")
	fmt.Printf("  Same input produces same key: %v\n", key1 == key2)
	fmt.Printf("  Diff input produces diff key: %v\n", key1 != key3)

	// Show which activities are forward vs compensation.
	fmt.Println("\n--- Activity Registry ---")
	fmt.Println("Forward activities:")
	fmt.Println("  ValidatePatientData     — pure validation, no external calls")
	fmt.Println("  CreatePatientRecord     — POST to Patient DB Service")
	fmt.Println("  VerifyInsurance         — POST to Insurance Eligibility API")
	fmt.Println("  ScheduleAppointment     — POST to Scheduling Service")
	fmt.Println("  RoutePrescription       — POST to Pharmacy Service (per rx)")
	fmt.Println("  CollectDocuments        — POST to Document Management Service")
	fmt.Println("  SendNotification        — POST to Notification Service")
	fmt.Println("\nCompensation activities (saga rollback):")
	fmt.Println("  DeletePatientRecord     — DELETE patient on rollback")
	fmt.Println("  CancelAppointment       — POST cancel on rollback")
	fmt.Println("  CancelPrescription      — POST cancel on rollback (per rx)")
}
