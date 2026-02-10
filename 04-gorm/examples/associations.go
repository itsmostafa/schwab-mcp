// Run: go mod init example && go get gorm.io/gorm gorm.io/driver/sqlite && go run associations.go
//
// This file demonstrates Preload vs Joins, the N+1 problem, nested preload,
// and conditional preloading for GORM associations.

package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Name  string `gorm:"size:100;not null"`
	Posts []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	Title     string    `gorm:"size:200;not null"`
	Published bool      `gorm:"default:false"`
	UserID    uint      `gorm:"not null"`
	Comments  []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	gorm.Model
	Body   string `gorm:"type:text;not null"`
	PostID uint   `gorm:"not null"`
}

func seedData(db *gorm.DB) {
	// Create users with nested associations in one call.
	// GORM automatically creates associated records when the parent is created.
	users := []User{
		{
			Name: "Alice",
			Posts: []Post{
				{
					Title:     "Go Concurrency",
					Published: true,
					Comments: []Comment{
						{Body: "Great article!"},
						{Body: "Very helpful."},
					},
				},
				{
					Title:     "Draft Post",
					Published: false,
					Comments: []Comment{
						{Body: "Needs editing."},
					},
				},
			},
		},
		{
			Name: "Bob",
			Posts: []Post{
				{
					Title:     "GORM Tips",
					Published: true,
					Comments: []Comment{
						{Body: "Nice tips!"},
						{Body: "Thanks for sharing."},
						{Body: "Bookmarked."},
					},
				},
			},
		},
		{
			Name: "Charlie",
			// Charlie has no posts — useful for testing outer-join behavior.
		},
	}
	db.Create(&users)
}

func main() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		// Use Warn level so we can see SQL queries when we enable Debug().
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatal("failed to connect:", err)
	}
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	seedData(db)

	fmt.Println("=== Seeded: 3 users, 3 posts, 6 comments ===")

	// --- N+1 Problem (BAD) ---
	fmt.Println("\n=== N+1 Problem (BAD) ===")
	// Interview note: this is the classic N+1 problem. We issue 1 query to get
	// all users, then N additional queries (one per user) to get their posts.
	// With 1000 users, that's 1001 queries — a performance disaster.
	var users []User
	db.Find(&users) // Query 1: SELECT * FROM users
	queryCount := 1
	for _, u := range users {
		var posts []Post
		db.Where("user_id = ?", u.ID).Find(&posts) // Query 2, 3, ... N+1
		queryCount++
		fmt.Printf("  User %s has %d posts (separate query #%d)\n", u.Name, len(posts), queryCount)
	}
	fmt.Printf("  Total queries: %d (1 for users + %d for posts)\n", queryCount, queryCount-1)

	// --- Preload Solution (GOOD) ---
	fmt.Println("\n=== Preload Solution (GOOD) ===")
	// Interview note: Preload issues a SEPARATE batched query for the association:
	//   1. SELECT * FROM users
	//   2. SELECT * FROM posts WHERE user_id IN (1, 2, 3)
	// Only 2 queries regardless of how many users — O(1) instead of O(N).
	var preloadedUsers []User
	db.Preload("Posts").Find(&preloadedUsers)
	for _, u := range preloadedUsers {
		fmt.Printf("  User %s has %d posts\n", u.Name, len(u.Posts))
	}
	fmt.Println("  Total queries: 2 (1 for users + 1 batched for posts)")

	// --- Nested Preload ---
	fmt.Println("\n=== Nested Preload (Posts.Comments) ===")
	// Interview note: dot notation loads nested associations. This adds one more
	// batched query for comments, so 3 queries total — still O(1) per level.
	var nestedUsers []User
	db.Preload("Posts.Comments").Find(&nestedUsers)
	for _, u := range nestedUsers {
		fmt.Printf("  User %s:\n", u.Name)
		for _, p := range u.Posts {
			fmt.Printf("    Post '%s' has %d comments\n", p.Title, len(p.Comments))
		}
	}
	fmt.Println("  Total queries: 3 (users + posts + comments)")

	// --- Preload with Conditions ---
	fmt.Println("\n=== Preload with Conditions ===")
	// Interview note: you can filter what gets preloaded. This does NOT filter
	// the parent (users) — it only filters which associated records are loaded.
	// All users are returned; only published posts are included in the slice.
	var condUsers []User
	db.Preload("Posts", "published = ?", true).Find(&condUsers)
	for _, u := range condUsers {
		fmt.Printf("  User %s: %d published posts loaded\n", u.Name, len(u.Posts))
	}
	fmt.Println("  Note: ALL users returned; only published posts are preloaded.")

	// --- Joins for Filtering ---
	fmt.Println("\n=== Joins for Filtering ===")
	// Interview note: use Joins when you need to FILTER PARENTS based on
	// associated data. Preload cannot do this — it only controls what gets loaded.
	// Joins produces a SQL JOIN, which lets you use WHERE on joined columns.
	var activeAuthors []User
	db.Joins("JOIN posts ON posts.user_id = users.id AND posts.deleted_at IS NULL").
		Where("posts.published = ?", true).
		Group("users.id"). // avoid duplicate users from the join
		Find(&activeAuthors)
	fmt.Println("  Users with at least one published post:")
	for _, u := range activeAuthors {
		fmt.Printf("    - %s (ID=%d)\n", u.Name, u.ID)
	}
	fmt.Println("  Charlie excluded — has no published posts.")

	// --- Combining Joins + Preload ---
	fmt.Println("\n=== Joins + Preload Combined ===")
	// Interview note: the practical pattern is Joins for filtering parents,
	// then Preload for loading the full association graph. This gives you
	// efficient filtering AND clean data loading in minimal queries.
	var filteredWithPosts []User
	db.Joins("JOIN posts ON posts.user_id = users.id AND posts.deleted_at IS NULL").
		Where("posts.published = ?", true).
		Group("users.id").
		Preload("Posts").         // load ALL posts for matched users
		Preload("Posts.Comments"). // also load comments
		Find(&filteredWithPosts)
	fmt.Println("  Users with published posts (full graph loaded):")
	for _, u := range filteredWithPosts {
		fmt.Printf("    %s: %d total posts\n", u.Name, len(u.Posts))
		for _, p := range u.Posts {
			fmt.Printf("      '%s' (published=%v) — %d comments\n", p.Title, p.Published, len(p.Comments))
		}
	}

	fmt.Println("\n=== Summary ===")
	fmt.Println("  - N+1: 1 query for parents + N queries for children = bad")
	fmt.Println("  - Preload: 1 query for parents + 1 batched query per association level = good")
	fmt.Println("  - Nested Preload: dot notation (\"Posts.Comments\") for deep loading")
	fmt.Println("  - Preload with conditions: filters LOADED associations, not parents")
	fmt.Println("  - Joins: filters PARENTS by association data (SQL JOIN)")
	fmt.Println("  - Combine Joins (filter) + Preload (load) for best of both")
}
