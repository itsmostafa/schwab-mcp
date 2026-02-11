package main

import (
	"fmt"
	"time"
)

// Data models for an MSO (Medical Service Organization) patient onboarding
// scheduler built on Temporal.
//
// Design decisions:
//   - The scheduler doesn't own patient data — it orchestrates onboarding steps
//     across downstream services (patient DB, insurance, scheduling, pharmacy,
//     document management, notifications).
//   - OnboardingState is the workflow's queryable state, updated as each step
//     completes. Temporal persists this automatically via workflow history.
//   - Signal payloads (DocumentUploadedSignal, InsuranceResponseSignal) model
//     async events that arrive from external systems mid-workflow.
//   - The state machine enforces valid status transitions in application code.
//     Temporal's workflow history provides the audit trail.

// --- Patient & related structs ---

// Patient holds demographic info submitted in the intake form.
type Patient struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     Address   `json:"address"`
}

// Address is a sub-document embedded in Patient.
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

// InsuranceInfo holds the patient's insurance details for eligibility verification.
type InsuranceInfo struct {
	ProviderName string `json:"provider_name"`
	MemberID     string `json:"member_id"`
	GroupNumber  string `json:"group_number"`
	PlanType     string `json:"plan_type"` // HMO, PPO, EPO
}

// InsuranceVerificationResult is returned by the insurance verification child workflow.
type InsuranceVerificationResult struct {
	Verified     bool   `json:"verified"`
	EligibleFrom string `json:"eligible_from"` // ISO date
	EligibleTo   string `json:"eligible_to"`   // ISO date
	Copay        int    `json:"copay"`         // cents
	Deductible   int    `json:"deductible"`    // cents
	Notes        string `json:"notes"`
}

// IntakeQuestionnaire holds patient-submitted health intake data.
type IntakeQuestionnaire struct {
	PrimaryConcern    string   `json:"primary_concern"`
	CurrentMedications []string `json:"current_medications"`
	Allergies         []string `json:"allergies"`
	PastSurgeries     []string `json:"past_surgeries"`
}

// Prescription represents a medication to be routed to a pharmacy.
type Prescription struct {
	MedicationName string `json:"medication_name"`
	Dosage         string `json:"dosage"`
	Frequency      string `json:"frequency"`
	Prescriber     string `json:"prescriber"`
	PharmacyID     string `json:"pharmacy_id"`
}

// Appointment represents a scheduled visit.
type Appointment struct {
	Type       string    `json:"type"` // initial_consultation, follow_up, lab_work
	ProviderID string    `json:"provider_id"`
	ScheduledAt time.Time `json:"scheduled_at"`
	Location   string    `json:"location"`
	Notes      string    `json:"notes"`
}

// DocumentRequirement represents a document the patient must upload.
type DocumentRequirement struct {
	DocumentType string `json:"document_type"` // photo_id, insurance_card, consent_form, medical_records
	Required     bool   `json:"required"`
	Uploaded     bool   `json:"uploaded"`
	UploadedAt   *time.Time `json:"uploaded_at,omitempty"`
}

// --- Workflow input/state ---

// OnboardingRequest is the input to the PatientOnboardingWorkflow.
// Contains everything needed to orchestrate the full onboarding flow.
type OnboardingRequest struct {
	Patient       Patient             `json:"patient"`
	Insurance     InsuranceInfo       `json:"insurance"`
	Intake        IntakeQuestionnaire `json:"intake"`
	Prescriptions []Prescription      `json:"prescriptions"`
	Appointment   Appointment         `json:"appointment"`
	Documents     []DocumentRequirement `json:"documents"`
}

// OnboardingState is the queryable workflow state. Updated after each step
// completes. Temporal persists this as part of the workflow execution — no
// external database needed for tracking progress.
type OnboardingState struct {
	Status              OnboardingStatus            `json:"status"`
	PatientID           string                      `json:"patient_id,omitempty"`
	InsuranceResult     *InsuranceVerificationResult `json:"insurance_result,omitempty"`
	AppointmentID       string                      `json:"appointment_id,omitempty"`
	PrescriptionIDs     []string                    `json:"prescription_ids,omitempty"`
	DocumentsCollected  int                         `json:"documents_collected"`
	DocumentsRequired   int                         `json:"documents_required"`
	CurrentStep         string                      `json:"current_step"`
	FailureReason       string                      `json:"failure_reason,omitempty"`
	StartedAt           time.Time                   `json:"started_at"`
	CompletedAt         *time.Time                  `json:"completed_at,omitempty"`
}

// --- Signal payloads ---

// DocumentUploadedSignal is sent to the workflow when a patient uploads a
// required document via the portal. The workflow waits for all required
// documents before proceeding.
type DocumentUploadedSignal struct {
	DocumentType string `json:"document_type"`
	StorageURL   string `json:"storage_url"`
}

// InsuranceResponseSignal is sent when the insurance company responds to
// the verification request. This may arrive hours or days after the initial
// request, so it's modeled as a signal rather than a synchronous response.
type InsuranceResponseSignal struct {
	Approved     bool   `json:"approved"`
	EligibleFrom string `json:"eligible_from"`
	EligibleTo   string `json:"eligible_to"`
	Copay        int    `json:"copay"`
	Deductible   int    `json:"deductible"`
	DenialReason string `json:"denial_reason,omitempty"`
}

// --- State machine ---

// OnboardingStatus represents a step in the onboarding lifecycle.
type OnboardingStatus string

const (
	StatusPending              OnboardingStatus = "pending"
	StatusValidating           OnboardingStatus = "validating"
	StatusCreatingRecord       OnboardingStatus = "creating_record"
	StatusVerifyingInsurance   OnboardingStatus = "verifying_insurance"
	StatusSchedulingAppointment OnboardingStatus = "scheduling_appointment"
	StatusRoutingPrescription  OnboardingStatus = "routing_prescription"
	StatusCollectingDocuments  OnboardingStatus = "collecting_documents"
	StatusSendingNotification  OnboardingStatus = "sending_notification"
	StatusCompleted            OnboardingStatus = "completed"
	StatusFailed               OnboardingStatus = "failed"
)

// validOnboardingTransitions defines the state machine. Each status maps to
// the set of statuses it can transition to. Any status can transition to
// "failed" — this is enforced in TransitionTo rather than duplicating it
// in every entry.
//
// Happy path: pending → validating → creating_record → verifying_insurance →
//   scheduling_appointment → routing_prescription → collecting_documents →
//   sending_notification → completed
var validOnboardingTransitions = map[OnboardingStatus][]OnboardingStatus{
	StatusPending:              {StatusValidating},
	StatusValidating:           {StatusCreatingRecord},
	StatusCreatingRecord:       {StatusVerifyingInsurance},
	StatusVerifyingInsurance:   {StatusSchedulingAppointment},
	StatusSchedulingAppointment: {StatusRoutingPrescription},
	StatusRoutingPrescription:  {StatusCollectingDocuments},
	StatusCollectingDocuments:  {StatusSendingNotification},
	StatusSendingNotification:  {StatusCompleted},
}

// TransitionTo validates and applies a status transition on OnboardingState.
// Any state can transition to StatusFailed (represents unrecoverable errors).
// Returns an error if the transition is not allowed by the state machine.
func (s *OnboardingState) TransitionTo(next OnboardingStatus) error {
	// Any state can fail.
	if next == StatusFailed {
		s.Status = StatusFailed
		return nil
	}

	allowed := validOnboardingTransitions[s.Status]
	for _, valid := range allowed {
		if valid == next {
			s.Status = next
			return nil
		}
	}
	return fmt.Errorf("invalid onboarding transition: %s -> %s", s.Status, next)
}

func main() {
	fmt.Println("=== MSO Scheduler — Data Models & State Machine ===\n")

	// Create a sample onboarding request.
	req := OnboardingRequest{
		Patient: Patient{
			FirstName:   "Jane",
			LastName:    "Smith",
			DateOfBirth: time.Date(1990, 3, 15, 0, 0, 0, 0, time.UTC),
			Email:       "jane.smith@example.com",
			Phone:       "+15551234567",
			Address: Address{
				Street: "123 Main St",
				City:   "Austin",
				State:  "TX",
				Zip:    "78701",
			},
		},
		Insurance: InsuranceInfo{
			ProviderName: "Blue Cross Blue Shield",
			MemberID:     "BCBS-789012",
			GroupNumber:  "GRP-456",
			PlanType:     "PPO",
		},
		Intake: IntakeQuestionnaire{
			PrimaryConcern:     "Annual physical examination",
			CurrentMedications: []string{"Lisinopril 10mg"},
			Allergies:          []string{"Penicillin"},
		},
		Prescriptions: []Prescription{
			{
				MedicationName: "Lisinopril",
				Dosage:         "10mg",
				Frequency:      "once daily",
				Prescriber:     "Dr. Johnson",
				PharmacyID:     "pharm-001",
			},
		},
		Appointment: Appointment{
			Type:        "initial_consultation",
			ProviderID:  "prov-dr-johnson",
			ScheduledAt: time.Now().Add(7 * 24 * time.Hour),
			Location:    "Main Clinic - Room 204",
		},
		Documents: []DocumentRequirement{
			{DocumentType: "photo_id", Required: true},
			{DocumentType: "insurance_card", Required: true},
			{DocumentType: "consent_form", Required: true},
			{DocumentType: "medical_records", Required: false},
		},
	}

	fmt.Printf("Patient: %s %s (%s)\n", req.Patient.FirstName, req.Patient.LastName, req.Patient.Email)
	fmt.Printf("Insurance: %s (Member: %s)\n", req.Insurance.ProviderName, req.Insurance.MemberID)
	fmt.Printf("Prescriptions: %d\n", len(req.Prescriptions))
	fmt.Printf("Required documents: ")
	count := 0
	for _, d := range req.Documents {
		if d.Required {
			count++
		}
	}
	fmt.Printf("%d\n\n", count)

	// Demonstrate the state machine.
	state := OnboardingState{
		Status:    StatusPending,
		StartedAt: time.Now(),
	}

	// Walk through the happy path.
	fmt.Println("--- State Machine Transitions ---")
	happyPath := []OnboardingStatus{
		StatusValidating,
		StatusCreatingRecord,
		StatusVerifyingInsurance,
		StatusSchedulingAppointment,
		StatusRoutingPrescription,
		StatusCollectingDocuments,
		StatusSendingNotification,
		StatusCompleted,
	}

	for _, next := range happyPath {
		prev := state.Status
		if err := state.TransitionTo(next); err != nil {
			fmt.Printf("  ERROR: %v\n", err)
		} else {
			fmt.Printf("  %s -> %s\n", prev, state.Status)
		}
	}

	// Demonstrate invalid transition.
	fmt.Println("\n--- Invalid Transition ---")
	state2 := OnboardingState{Status: StatusPending}
	if err := state2.TransitionTo(StatusCompleted); err != nil {
		fmt.Printf("  Expected error: %v\n", err)
	}

	// Demonstrate failure from any state.
	fmt.Println("\n--- Failure Transition ---")
	state3 := OnboardingState{Status: StatusVerifyingInsurance}
	if err := state3.TransitionTo(StatusFailed); err != nil {
		fmt.Printf("  Unexpected error: %v\n", err)
	} else {
		fmt.Printf("  %s -> %s (any state can fail)\n", StatusVerifyingInsurance, state3.Status)
	}
}
