package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB document models for a simulated phishing platform.
//
// Design decisions:
//   - Campaign embeds EmailTemplate and CampaignStats because they are always
//     read together and bounded in size.
//   - EmailEvent is a separate collection because events grow unboundedly
//     (one per recipient per event type) and need independent querying.
//   - bson tags map Go field names to MongoDB field names. Use snake_case
//     to follow MongoDB conventions.
//   - primitive.ObjectID is the Go type for MongoDB's ObjectId (_id field).

// Organization represents an enrolled company.
type Organization struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name"          json:"name"`
	Domain    string             `bson:"domain"        json:"domain"`
	PlanTier  string             `bson:"plan_tier"     json:"plan_tier"` // free, pro, enterprise
	CreatedAt time.Time          `bson:"created_at"    json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"    json:"updated_at"`
}

// Employee represents a person in an organization's roster.
type Employee struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrgID      primitive.ObjectID `bson:"org_id"        json:"org_id"`
	Email      string             `bson:"email"         json:"email"`
	FirstName  string             `bson:"first_name"    json:"first_name"`
	LastName   string             `bson:"last_name"     json:"last_name"`
	Department string             `bson:"department"    json:"department"`
	ImportedAt time.Time          `bson:"imported_at"   json:"imported_at"`
}

// Campaign represents a phishing simulation campaign.
//
// Status state machine:
//   draft → scheduled → sending → completed
//   draft → cancelled
//   scheduled → cancelled
//   sending → cancelled (already-sent emails cannot be recalled)
type Campaign struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	OrgID     primitive.ObjectID   `bson:"org_id"        json:"org_id"`
	Name      string               `bson:"name"          json:"name"`
	Status    string               `bson:"status"        json:"status"`
	Template  EmailTemplate        `bson:"template"      json:"template"`
	Schedule  Schedule             `bson:"schedule"      json:"schedule"`
	TargetIDs []primitive.ObjectID `bson:"target_ids"    json:"target_ids"`
	Stats     CampaignStats        `bson:"stats"         json:"stats"`
	CreatedAt time.Time            `bson:"created_at"    json:"created_at"`
	UpdatedAt time.Time            `bson:"updated_at"    json:"updated_at"`
}

// EmailTemplate is embedded in Campaign. Bounded size, always read together.
type EmailTemplate struct {
	Subject  string `bson:"subject"   json:"subject"`
	HTMLBody string `bson:"html_body" json:"html_body"`
	FromName string `bson:"from_name" json:"from_name"`
	FromAddr string `bson:"from_addr" json:"from_addr"`
}

// Schedule defines the send window for a campaign.
type Schedule struct {
	StartAt time.Time `bson:"start_at" json:"start_at"`
	EndAt   time.Time `bson:"end_at"   json:"end_at"`
}

// CampaignStats holds denormalized counters. Updated by workers after each send.
// Denormalized to avoid aggregation queries for the campaign list view.
type CampaignStats struct {
	TotalTargets int `bson:"total_targets" json:"total_targets"`
	Sent         int `bson:"sent"          json:"sent"`
	Delivered    int `bson:"delivered"     json:"delivered"`
	Opened       int `bson:"opened"        json:"opened"`
	Clicked      int `bson:"clicked"       json:"clicked"`
}

// EmailEvent records a single tracking event (event-sourcing style).
// Each event type carries different metadata — this schema flexibility
// is one reason MongoDB is chosen over a relational database.
type EmailEvent struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CampaignID primitive.ObjectID `bson:"campaign_id"   json:"campaign_id"`
	EmployeeID primitive.ObjectID `bson:"employee_id"   json:"employee_id"`
	Token      string             `bson:"token"         json:"token"`
	EventType  string             `bson:"event_type"    json:"event_type"` // sent, delivered, opened, clicked
	Metadata   EventMetadata      `bson:"metadata,omitempty" json:"metadata,omitempty"`
	OccurredAt time.Time          `bson:"occurred_at"   json:"occurred_at"`
}

// EventMetadata holds optional fields that vary by event type.
// "opened" events have IP and UserAgent; "clicked" events have LinkURL.
type EventMetadata struct {
	IP        string `bson:"ip,omitempty"         json:"ip,omitempty"`
	UserAgent string `bson:"user_agent,omitempty" json:"user_agent,omitempty"`
	LinkURL   string `bson:"link_url,omitempty"   json:"link_url,omitempty"`
}

// Campaign status constants.
const (
	StatusDraft     = "draft"
	StatusScheduled = "scheduled"
	StatusSending   = "sending"
	StatusCompleted = "completed"
	StatusCancelled = "cancelled"
)

// validTransitions defines the campaign state machine.
// Used by TransitionTo to enforce valid state changes.
var validTransitions = map[string][]string{
	StatusDraft:     {StatusScheduled, StatusCancelled},
	StatusScheduled: {StatusSending, StatusCancelled},
	StatusSending:   {StatusCompleted, StatusCancelled},
}

// TransitionTo validates and applies a status transition.
// Returns an error if the transition is not allowed.
func (c *Campaign) TransitionTo(next string) error {
	allowed := validTransitions[c.Status]
	for _, s := range allowed {
		if s == next {
			c.Status = next
			c.UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("invalid campaign transition: %s → %s", c.Status, next)
}

// createIndexes sets up all required indexes for the four collections.
//
// Index strategy:
//   organizations:  unique on domain (prevent duplicate orgs)
//   employees:      unique compound on {org_id, email} (prevent duplicate employees per org)
//   campaigns:      compound on {org_id, status} (list by org and filter by status)
//   email_events:   unique on token (tracking lookups)
//                   compound on {campaign_id, event_type} (aggregation queries)
//                   TTL on occurred_at (auto-delete after 90 days for GDPR)
func createIndexes(ctx context.Context, db *mongo.Database) error {
	// --- organizations ---
	_, err := db.Collection("organizations").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "domain", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return fmt.Errorf("creating organizations index: %w", err)
	}

	// --- employees ---
	_, err = db.Collection("employees").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "org_id", Value: 1}, {Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return fmt.Errorf("creating employees index: %w", err)
	}

	// --- campaigns ---
	_, err = db.Collection("campaigns").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "org_id", Value: 1}, {Key: "status", Value: 1}},
	})
	if err != nil {
		return fmt.Errorf("creating campaigns index: %w", err)
	}

	// --- email_events ---
	eventIndexes := db.Collection("email_events").Indexes()

	// Unique token index — used by tracking endpoints and idempotency checks.
	_, err = eventIndexes.CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "token", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return fmt.Errorf("creating email_events token index: %w", err)
	}

	// Compound index for aggregation queries (campaign analytics).
	_, err = eventIndexes.CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "campaign_id", Value: 1}, {Key: "event_type", Value: 1}},
	})
	if err != nil {
		return fmt.Errorf("creating email_events compound index: %w", err)
	}

	// TTL index — MongoDB automatically deletes documents 90 days after occurred_at.
	// This handles GDPR data retention without custom batch cleanup jobs.
	ninetyDays := int32(90 * 24 * 60 * 60)
	_, err = eventIndexes.CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "occurred_at", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(ninetyDays),
	})
	if err != nil {
		return fmt.Errorf("creating email_events TTL index: %w", err)
	}

	return nil
}

func main() {
	fmt.Println("=== MongoDB Models & Index Creation ===\n")

	// Connect to MongoDB.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("mongo connect:", err)
	}
	defer client.Disconnect(ctx)

	// Ping to verify connection.
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("mongo ping:", err)
	}
	fmt.Println("Connected to MongoDB")

	db := client.Database("phishing_platform")

	// Create all indexes.
	if err := createIndexes(ctx, db); err != nil {
		log.Fatal("creating indexes:", err)
	}
	fmt.Println("All indexes created successfully")

	// Demonstrate inserting an organization.
	org := Organization{
		Name:      "Acme Corp",
		Domain:    "acme.com",
		PlanTier:  "pro",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := db.Collection("organizations").InsertOne(ctx, org)
	if err != nil {
		log.Fatal("insert org:", err)
	}
	fmt.Printf("Inserted organization: %s\n", result.InsertedID)

	// Demonstrate the state machine.
	campaign := Campaign{Status: StatusDraft}
	fmt.Printf("\nCampaign status: %s\n", campaign.Status)

	if err := campaign.TransitionTo(StatusScheduled); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After transition: %s\n", campaign.Status)

	// Invalid transition: scheduled → completed (must go through sending first).
	if err := campaign.TransitionTo(StatusCompleted); err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	// Valid: scheduled → sending → completed.
	campaign.TransitionTo(StatusSending)
	campaign.TransitionTo(StatusCompleted)
	fmt.Printf("Final status: %s\n", campaign.Status)
}
