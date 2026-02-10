package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Tracking service: handles email open tracking (pixel) and link click tracking.
//
// Key interview concepts:
//   - Tracking pixel: 1x1 transparent GIF embedded in email HTML. When the email
//     client renders the image, it hits our endpoint, recording an "opened" event.
//   - Link click: all URLs in the email are rewritten to pass through our endpoint,
//     which records a "clicked" event and redirects to the educational landing page.
//   - Tokens are generated with crypto/rand — unpredictable, preventing enumeration.
//   - Redirect URL validation prevents open redirect vulnerabilities.
//   - Cache-Control: no-store prevents caching from suppressing open tracking.

// transparentGIF is the smallest valid 1x1 transparent GIF (43 bytes).
// This is returned by the tracking pixel endpoint.
// Using a hardcoded byte slice avoids file I/O on every request.
var transparentGIF = []byte{
	0x47, 0x49, 0x46, 0x38, 0x39, 0x61, // GIF89a
	0x01, 0x00, 0x01, 0x00, // width=1, height=1
	0x80, 0x00, 0x00, // GCT flag, 2 colors
	0xff, 0xff, 0xff, // color 0: white
	0x00, 0x00, 0x00, // color 1: black
	0x21, 0xf9, 0x04, // graphic control extension
	0x01, 0x00, 0x00, 0x00, 0x00, // transparent index = 0
	0x2c, 0x00, 0x00, 0x00, 0x00, // image descriptor
	0x01, 0x00, 0x01, 0x00, 0x00, // 1x1
	0x02, 0x02, 0x44, 0x01, 0x00, // LZW min code size, data
	0x3b, // trailer
}

// TrackerService handles tracking pixel and link click endpoints.
type TrackerService struct {
	events *mongo.Collection
	rdb    *redis.Client
}

// HandleOpen serves the tracking pixel and records an "opened" event.
//
// GET /t/open/:token
//
// Flow:
//  1. Look up the token in email_events to find the campaign/employee.
//  2. Insert an "opened" event with IP and User-Agent metadata.
//  3. Update Redis HyperLogLog for unique opens and increment total counter.
//  4. Return the 1x1 transparent GIF.
//
// The endpoint always returns 200 + GIF regardless of errors. Tracking is
// best-effort — a failed DB write should not break the user's email client.
func (s *TrackerService) HandleOpen(c *gin.Context) {
	token := c.Param("token")
	ctx := c.Request.Context()

	// Look up the original sent event to get campaign and employee IDs.
	var sentEvent struct {
		CampaignID primitive.ObjectID `bson:"campaign_id"`
		EmployeeID primitive.ObjectID `bson:"employee_id"`
	}
	err := s.events.FindOne(ctx, bson.M{
		"token":      token,
		"event_type": "sent",
	}).Decode(&sentEvent)

	if err == nil {
		// Record the open event. Use InsertOne — duplicates are fine for opens
		// (the user may open the email multiple times).
		_, _ = s.events.InsertOne(ctx, bson.M{
			"campaign_id": sentEvent.CampaignID,
			"employee_id": sentEvent.EmployeeID,
			"token":       token,
			"event_type":  "opened",
			"metadata": bson.M{
				"ip":         c.ClientIP(),
				"user_agent": c.Request.UserAgent(),
			},
			"occurred_at": time.Now(),
		})

		// Update Redis counters (best-effort, ignore errors).
		// Pipeline sends both commands in one round-trip.
		cid := sentEvent.CampaignID.Hex()
		eid := sentEvent.EmployeeID.Hex()
		pipe := s.rdb.Pipeline()
		pipe.PFAdd(ctx, fmt.Sprintf("campaign:%s:unique_opens", cid), eid)
		pipe.Incr(ctx, fmt.Sprintf("campaign:%s:total_opens", cid))
		pipe.Exec(ctx)
	}

	// Always return the GIF. Tracking is best-effort.
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Data(http.StatusOK, "image/gif", transparentGIF)
}

// HandleClick records a "clicked" event and redirects to the landing page.
//
// GET /t/click/:token?url=<encoded_original_url>
//
// Flow:
//  1. Validate the URL parameter against an allowlist (prevent open redirects).
//  2. Look up the token to find campaign/employee.
//  3. Insert a "clicked" event.
//  4. Redirect to the educational landing page (not the original phishing URL).
//
// Security: The redirect always goes to the landing page, never to the
// attacker-controlled URL. The url parameter is logged for analytics but
// not used as the redirect target.
func (s *TrackerService) HandleClick(c *gin.Context) {
	token := c.Param("token")
	rawURL := c.Query("url")
	ctx := c.Request.Context()

	// Validate the URL parameter. In production, this would check against
	// the campaign's list of tracked URLs to prevent abuse.
	if !isValidTrackingURL(rawURL) {
		c.String(http.StatusBadRequest, "invalid url")
		return
	}

	// Look up the original sent event.
	var sentEvent struct {
		CampaignID primitive.ObjectID `bson:"campaign_id"`
		EmployeeID primitive.ObjectID `bson:"employee_id"`
	}
	err := s.events.FindOne(ctx, bson.M{
		"token":      token,
		"event_type": "sent",
	}).Decode(&sentEvent)

	if err == nil {
		// Record the click event.
		_, _ = s.events.InsertOne(ctx, bson.M{
			"campaign_id": sentEvent.CampaignID,
			"employee_id": sentEvent.EmployeeID,
			"token":       token,
			"event_type":  "clicked",
			"metadata": bson.M{
				"ip":         c.ClientIP(),
				"user_agent": c.Request.UserAgent(),
				"link_url":   rawURL,
			},
			"occurred_at": time.Now(),
		})

		// Update Redis counters.
		cid := sentEvent.CampaignID.Hex()
		eid := sentEvent.EmployeeID.Hex()
		pipe := s.rdb.Pipeline()
		pipe.PFAdd(ctx, fmt.Sprintf("campaign:%s:unique_clicks", cid), eid)
		pipe.Incr(ctx, fmt.Sprintf("campaign:%s:total_clicks", cid))
		pipe.Exec(ctx)
	}

	// Redirect to the educational landing page.
	// This page explains that the email was a phishing simulation and provides
	// security awareness training tips.
	c.Redirect(http.StatusFound, "/education/phishing-awareness")
}

// isValidTrackingURL validates that the URL is well-formed and uses HTTPS.
// In production, this would also check against the campaign's URL allowlist
// to prevent open redirect attacks.
func isValidTrackingURL(rawURL string) bool {
	if rawURL == "" {
		return false
	}
	u, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	// Only allow http/https schemes.
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	// Must have a host.
	if u.Host == "" {
		return false
	}
	return true
}

// generateToken creates a 32-byte URL-safe random token using crypto/rand.
// Tokens must be unpredictable to prevent:
//   - Enumeration attacks (guessing tokens to view other users' tracking data)
//   - Forgery (crafting fake open/click events)
func generateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("reading random bytes: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func main() {
	fmt.Println("=== Tracking Service ===\n")

	// Connect to MongoDB.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("mongo connect:", err)
	}
	defer mongoClient.Disconnect(ctx)

	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	defer rdb.Close()

	tracker := &TrackerService{
		events: mongoClient.Database("phishing_platform").Collection("email_events"),
		rdb:    rdb,
	}

	// Set up Gin router.
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	// Tracking endpoints — no auth middleware (must be accessible from email clients).
	// Rate limiting should be applied here in production to prevent abuse.
	tracking := r.Group("/t")
	{
		tracking.GET("/open/:token", tracker.HandleOpen)
		tracking.GET("/click/:token", tracker.HandleClick)
	}

	// Educational landing page (placeholder).
	r.GET("/education/phishing-awareness", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title":   "Security Awareness Training",
			"message": "This was a simulated phishing email. No credentials were captured.",
			"tips": []string{
				"Always verify the sender's email address",
				"Hover over links before clicking to check the URL",
				"Report suspicious emails to your IT security team",
				"Never enter credentials on unfamiliar websites",
			},
		})
	})

	// Demonstrate token generation.
	token, err := generateToken()
	if err != nil {
		log.Fatal("generating token:", err)
	}
	fmt.Printf("Sample token: %s\n", token)
	fmt.Printf("Tracking pixel URL: /t/open/%s\n", token)
	fmt.Printf("Click tracking URL: /t/click/%s?url=%s\n\n",
		token, url.QueryEscape("https://example.com/target-page"))

	fmt.Println("Starting tracking server on :8080")
	fmt.Println("  GET /t/open/:token    — tracking pixel")
	fmt.Println("  GET /t/click/:token   — link click tracker")
	fmt.Println("  GET /education/...    — landing page")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("server error:", err)
	}
}
