package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"

	redis_rate "github.com/go-redis/redis_rate/v10"
)

// Redis analytics patterns for the phishing platform.
//
// Key interview concepts:
//   - HyperLogLog: probabilistic unique counting with 12 KB per key and 0.81% error.
//   - Pipelining: batch multiple commands in one round-trip for throughput.
//   - Rate limiting: sliding window limiter using sorted sets.
//   - Distributed locks: prevent concurrent campaign scheduling.
//   - Caching: template cache with TTL to avoid repeated MongoDB reads.
//   - Hash counters: HSET/HINCRBY for real-time dashboard state.

// AnalyticsService provides Redis-backed analytics operations.
type AnalyticsService struct {
	rdb     *redis.Client
	limiter *redis_rate.Limiter
	locker  *redislock.Client
}

// NewAnalyticsService creates the service with all Redis dependencies.
func NewAnalyticsService(rdb *redis.Client) *AnalyticsService {
	return &AnalyticsService{
		rdb:     rdb,
		limiter: redis_rate.NewLimiter(rdb),
		locker:  redislock.New(rdb),
	}
}

// --- Pattern 1: HyperLogLog + Pipelined Counters ---

// RecordOpen updates both unique and total open counters using a pipeline.
//
// HyperLogLog (PFADD):
//   - Counts unique elements probabilistically.
//   - Uses only 12 KB per key regardless of cardinality.
//   - 0.81% standard error — acceptable for analytics dashboards.
//   - Adding the same employeeID multiple times doesn't increase the count.
//
// Pipeline:
//   - Sends both PFADD and INCR in a single Redis round-trip.
//   - Reduces latency from 2 RTTs to 1 RTT.
//   - Commands are executed atomically on the server side.
func (s *AnalyticsService) RecordOpen(ctx context.Context, campaignID, employeeID string) error {
	pipe := s.rdb.Pipeline()
	pipe.PFAdd(ctx, fmt.Sprintf("campaign:%s:unique_opens", campaignID), employeeID)
	pipe.Incr(ctx, fmt.Sprintf("campaign:%s:total_opens", campaignID))
	_, err := pipe.Exec(ctx)
	return err
}

// RecordClick updates both unique and total click counters.
func (s *AnalyticsService) RecordClick(ctx context.Context, campaignID, employeeID string) error {
	pipe := s.rdb.Pipeline()
	pipe.PFAdd(ctx, fmt.Sprintf("campaign:%s:unique_clicks", campaignID), employeeID)
	pipe.Incr(ctx, fmt.Sprintf("campaign:%s:total_clicks", campaignID))
	_, err := pipe.Exec(ctx)
	return err
}

// CampaignStats holds the analytics counters for a campaign.
type CampaignStats struct {
	UniqueOpens  int64
	TotalOpens   int64
	UniqueClicks int64
	TotalClicks  int64
}

// GetCampaignStats retrieves all counters for a campaign in a single pipeline.
// Uses PFCOUNT for HyperLogLog cardinality and GET for simple counters.
func (s *AnalyticsService) GetCampaignStats(ctx context.Context, campaignID string) (*CampaignStats, error) {
	pipe := s.rdb.Pipeline()

	uniqueOpens := pipe.PFCount(ctx, fmt.Sprintf("campaign:%s:unique_opens", campaignID))
	totalOpens := pipe.Get(ctx, fmt.Sprintf("campaign:%s:total_opens", campaignID))
	uniqueClicks := pipe.PFCount(ctx, fmt.Sprintf("campaign:%s:unique_clicks", campaignID))
	totalClicks := pipe.Get(ctx, fmt.Sprintf("campaign:%s:total_clicks", campaignID))

	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("pipeline exec: %w", err)
	}

	stats := &CampaignStats{}
	stats.UniqueOpens, _ = uniqueOpens.Result()
	stats.TotalOpens, _ = totalOpens.Int64()
	stats.UniqueClicks, _ = uniqueClicks.Result()
	stats.TotalClicks, _ = totalClicks.Int64()

	return stats, nil
}

// --- Pattern 2: Sliding Window Rate Limiting ---

// CheckRateLimit verifies that an organization hasn't exceeded its hourly send limit.
//
// redis_rate implements a sliding window algorithm using sorted sets:
//   - Each allowed request adds a timestamped entry to a sorted set.
//   - The limiter counts entries within the window to determine remaining capacity.
//   - Old entries outside the window are automatically removed.
//
// Returns:
//   - allowed: whether the request is within the limit
//   - retryAfter: how long to wait before retrying (if rate limited)
func (s *AnalyticsService) CheckRateLimit(ctx context.Context, orgID string, perHourLimit int) (allowed bool, retryAfter time.Duration, err error) {
	res, err := s.limiter.Allow(ctx,
		fmt.Sprintf("rate:org:%s", orgID),
		redis_rate.PerHour(perHourLimit),
	)
	if err != nil {
		return false, 0, fmt.Errorf("rate limit check: %w", err)
	}
	return res.Remaining > 0, res.RetryAfter, nil
}

// --- Pattern 3: Distributed Lock ---

// WithCampaignLock acquires a distributed lock for a campaign and executes fn.
// Only one process/goroutine can hold the lock at a time.
//
// Use case: prevent two scheduler instances from scheduling the same campaign
// concurrently, which would create duplicate tasks.
//
// How it works:
//   - Uses SET NX (set if not exists) with a TTL.
//   - The lock auto-expires after the TTL even if the holder crashes.
//   - On success, fn runs with the lock held.
//   - The lock is released immediately after fn completes (or on error).
//
// The TTL should be longer than the expected operation duration but short
// enough that a crashed holder doesn't block others for too long.
func (s *AnalyticsService) WithCampaignLock(ctx context.Context, campaignID string, ttl time.Duration, fn func() error) error {
	lock, err := s.locker.Obtain(ctx,
		fmt.Sprintf("lock:campaign:%s", campaignID),
		ttl,
		nil, // default retry strategy (no retry — fail immediately if locked)
	)
	if err == redislock.ErrNotObtained {
		return fmt.Errorf("campaign %s is locked by another process", campaignID)
	}
	if err != nil {
		return fmt.Errorf("obtaining lock: %w", err)
	}
	defer lock.Release(ctx)

	return fn()
}

// --- Pattern 4: Template Cache ---

// GetCachedTemplate retrieves a rendered email template from Redis cache.
// Falls back to the provided loader function on cache miss.
//
// Cache-aside pattern:
//  1. Check Redis for the cached value.
//  2. On miss, call the loader to get the value from the primary store (MongoDB).
//  3. Store the result in Redis with a TTL for future requests.
//
// TTL of 15 minutes balances freshness with read reduction. Templates change
// infrequently, so stale data for 15 minutes is acceptable.
func (s *AnalyticsService) GetCachedTemplate(ctx context.Context, templateID string, loader func() (string, error)) (string, error) {
	key := fmt.Sprintf("tmpl:%s", templateID)

	// Try cache first.
	cached, err := s.rdb.Get(ctx, key).Result()
	if err == nil {
		return cached, nil // cache hit
	}
	if err != redis.Nil {
		return "", fmt.Errorf("redis get: %w", err)
	}

	// Cache miss — load from primary store.
	value, err := loader()
	if err != nil {
		return "", fmt.Errorf("template loader: %w", err)
	}

	// Store in cache with TTL (best-effort — don't fail the request if caching fails).
	if err := s.rdb.Set(ctx, key, value, 15*time.Minute).Err(); err != nil {
		log.Printf("[warn] failed to cache template %s: %v", templateID, err)
	}

	return value, nil
}

// --- Pattern 5: Real-Time Dashboard Counters (Hash) ---

// UpdateProgress updates campaign progress counters in a Redis hash.
//
// Redis hashes store multiple field-value pairs under a single key.
// HINCRBY atomically increments a field, making it safe for concurrent updates
// from multiple workers.
//
// The dashboard reads the full hash with HGETALL to display:
//   - sent: emails successfully sent
//   - delivered: delivery confirmations received
//   - opened: tracking pixel hits
//   - clicked: link click events
func (s *AnalyticsService) UpdateProgress(ctx context.Context, campaignID, field string, delta int64) error {
	key := fmt.Sprintf("progress:%s", campaignID)
	return s.rdb.HIncrBy(ctx, key, field, delta).Err()
}

// GetProgress retrieves all progress counters for a campaign.
func (s *AnalyticsService) GetProgress(ctx context.Context, campaignID string) (map[string]string, error) {
	return s.rdb.HGetAll(ctx, fmt.Sprintf("progress:%s", campaignID)).Result()
}

func main() {
	fmt.Println("=== Redis Analytics Patterns ===\n")

	ctx := context.Background()

	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	defer rdb.Close()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal("redis ping:", err)
	}
	fmt.Println("Connected to Redis\n")

	svc := NewAnalyticsService(rdb)
	campaignID := "campaign_demo_001"

	// --- Demo: HyperLogLog + Counters ---
	fmt.Println("--- HyperLogLog + Pipelined Counters ---")

	// Simulate opens from 5 employees (some duplicates).
	employees := []string{"emp_1", "emp_2", "emp_3", "emp_2", "emp_1"}
	for _, emp := range employees {
		if err := svc.RecordOpen(ctx, campaignID, emp); err != nil {
			log.Fatal("record open:", err)
		}
	}
	fmt.Printf("Recorded %d open events (3 unique employees)\n", len(employees))

	// Simulate clicks from 2 employees.
	for _, emp := range []string{"emp_1", "emp_3"} {
		if err := svc.RecordClick(ctx, campaignID, emp); err != nil {
			log.Fatal("record click:", err)
		}
	}
	fmt.Println("Recorded 2 click events (2 unique employees)")

	// Read stats.
	stats, err := svc.GetCampaignStats(ctx, campaignID)
	if err != nil {
		log.Fatal("get stats:", err)
	}
	fmt.Printf("\nCampaign stats:\n")
	fmt.Printf("  Unique opens:  %d (expected: 3)\n", stats.UniqueOpens)
	fmt.Printf("  Total opens:   %d (expected: 5)\n", stats.TotalOpens)
	fmt.Printf("  Unique clicks: %d (expected: 2)\n", stats.UniqueClicks)
	fmt.Printf("  Total clicks:  %d (expected: 2)\n", stats.TotalClicks)

	// --- Demo: Rate Limiting ---
	fmt.Println("\n--- Sliding Window Rate Limiting ---")

	orgID := "org_demo_001"
	for i := 0; i < 5; i++ {
		allowed, retryAfter, err := svc.CheckRateLimit(ctx, orgID, 3) // limit: 3 per hour
		if err != nil {
			log.Fatal("rate limit:", err)
		}
		if allowed {
			fmt.Printf("  Request %d: allowed\n", i+1)
		} else {
			fmt.Printf("  Request %d: rate limited (retry after %v)\n", i+1, retryAfter)
		}
	}

	// --- Demo: Distributed Lock ---
	fmt.Println("\n--- Distributed Lock ---")

	err = svc.WithCampaignLock(ctx, campaignID, 10*time.Second, func() error {
		fmt.Println("  Lock acquired — scheduling campaign...")
		time.Sleep(100 * time.Millisecond) // simulate work
		fmt.Println("  Scheduling complete.")
		return nil
	})
	if err != nil {
		fmt.Printf("  Lock error: %v\n", err)
	}

	// --- Demo: Template Cache ---
	fmt.Println("\n--- Template Cache ---")

	loadCount := 0
	loader := func() (string, error) {
		loadCount++
		return "<html><body>Phishing template content</body></html>", nil
	}

	// First call: cache miss, calls loader.
	tmpl, _ := svc.GetCachedTemplate(ctx, "tmpl_001", loader)
	fmt.Printf("  Call 1 (miss): loaded from source, length=%d\n", len(tmpl))

	// Second call: cache hit, skips loader.
	tmpl, _ = svc.GetCachedTemplate(ctx, "tmpl_001", loader)
	fmt.Printf("  Call 2 (hit):  from cache, length=%d\n", len(tmpl))
	fmt.Printf("  Loader called %d time(s)\n", loadCount)

	// --- Demo: Dashboard Counters ---
	fmt.Println("\n--- Real-Time Dashboard Counters ---")

	svc.UpdateProgress(ctx, campaignID, "sent", 42)
	svc.UpdateProgress(ctx, campaignID, "delivered", 40)
	svc.UpdateProgress(ctx, campaignID, "opened", 25)
	svc.UpdateProgress(ctx, campaignID, "clicked", 8)

	progress, _ := svc.GetProgress(ctx, campaignID)
	fmt.Println("  Campaign progress:")
	for field, val := range progress {
		fmt.Printf("    %s: %s\n", field, val)
	}

	// Cleanup demo keys.
	rdb.Del(ctx,
		fmt.Sprintf("campaign:%s:unique_opens", campaignID),
		fmt.Sprintf("campaign:%s:total_opens", campaignID),
		fmt.Sprintf("campaign:%s:unique_clicks", campaignID),
		fmt.Sprintf("campaign:%s:total_clicks", campaignID),
		fmt.Sprintf("rate:org:%s", orgID),
		fmt.Sprintf("tmpl:tmpl_001"),
		fmt.Sprintf("progress:%s", campaignID),
	)
	fmt.Println("\nDemo keys cleaned up.")
}
