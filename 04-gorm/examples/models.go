// Run: go mod init example && go get gorm.io/gorm gorm.io/driver/sqlite && go run models.go
//
// This file demonstrates GORM model definitions, conventions, association tags,
// custom table names, and AutoMigrate behavior.

package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// --- Model Definitions ---

// User demonstrates gorm.Model embedding and HasMany / HasOne associations.
// gorm.Model provides: ID (uint, primary key), CreatedAt, UpdatedAt, DeletedAt (soft delete).
// Interview note: gorm.Model gives you soft delete for free via DeletedAt (gorm.DeletedAt).
type User struct {
	gorm.Model
	Name  string `gorm:"size:100;not null"`
	Email string `gorm:"uniqueIndex;not null"`

	// HasMany: one user has many posts.
	// GORM infers foreign key as "UserID" on the Post table by convention.
	// Explicit tags shown here for clarity — useful when FK name differs from convention.
	Posts []Post `gorm:"foreignKey:UserID;references:ID"`

	// HasOne: one user has one profile.
	Profile Profile `gorm:"foreignKey:UserID;references:ID"`
}

// Post demonstrates BelongsTo (User) and HasMany (Comments).
// Convention: GORM pluralizes the struct name for the table → "posts".
// Convention: field names map to snake_case columns → UserID becomes "user_id".
type Post struct {
	gorm.Model
	Title     string `gorm:"size:200;not null"`
	Body      string `gorm:"type:text"`
	Published bool   `gorm:"default:false"`

	// BelongsTo: post belongs to a user.
	// The FK field (UserID) lives on this struct — that's what makes it BelongsTo.
	// We omit the User back-reference here to avoid a recursive type cycle.
	// In real code, you can use a pointer (*User) to break the cycle if needed.
	UserID uint

	// HasMany: one post has many comments.
	Comments []Comment `gorm:"foreignKey:PostID;references:ID"`
}

// Comment demonstrates a simple BelongsTo association back to Post.
type Comment struct {
	gorm.Model
	Body string `gorm:"type:text;not null"`

	// BelongsTo: comment belongs to a post.
	// The FK field (PostID) establishes the BelongsTo relationship.
	PostID uint
}

// Profile demonstrates HasOne (owned by User) with a custom table name.
// Note: we omit the back-reference to User here to avoid a recursive type.
// In practice you can break the cycle with a pointer (*User) or simply
// omit the inverse side when you don't need bidirectional navigation.
type Profile struct {
	gorm.Model
	Bio    string `gorm:"type:text"`
	Avatar string `gorm:"size:255"`

	// BelongsTo: profile belongs to a user.
	// The FK field (UserID) is what makes this a BelongsTo relationship.
	UserID uint
}

// TableName overrides GORM's default table name convention.
// Without this, GORM would use "profiles" (pluralized struct name).
// Interview note: use TableName() when you need to match an existing schema
// or follow a naming convention different from GORM's default pluralization.
func (Profile) TableName() string {
	return "user_profiles"
}

func main() {
	// Open an in-memory SQLite database — no file needed, perfect for examples and tests.
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	fmt.Println("=== AutoMigrate ===")
	// AutoMigrate creates tables, adds missing columns and indexes.
	// Interview note: AutoMigrate is ADDITIVE ONLY — it will NOT drop columns,
	// change column types, or remove indexes. For production, use versioned
	// migration tools (goose, atlas, dbmate).
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}, &Profile{}); err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}
	fmt.Println("AutoMigrate completed successfully.")
	fmt.Println("Tables created: users, posts, comments, user_profiles (custom name)")

	// Verify tables exist by querying SQLite's metadata.
	fmt.Println("\n=== Created Tables ===")
	var tables []struct{ Name string }
	db.Raw("SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%' ORDER BY name").Scan(&tables)
	for _, t := range tables {
		fmt.Printf("  - %s\n", t.Name)
	}

	// Show column info for the users table to demonstrate convention mappings.
	fmt.Println("\n=== Column Info: users ===")
	type ColumnInfo struct {
		CID       int
		Name      string
		Type      string
		NotNull   int
		DfltValue *string
		PK        int
	}
	var columns []ColumnInfo
	db.Raw("PRAGMA table_info(users)").Scan(&columns)
	for _, c := range columns {
		pk := ""
		if c.PK == 1 {
			pk = " (PRIMARY KEY)"
		}
		fmt.Printf("  %-15s %-15s%s\n", c.Name, c.Type, pk)
	}
	// Interview note: notice snake_case columns (created_at, deleted_at),
	// auto-generated "id" primary key, and "deleted_at" for soft delete support.

	// Show the custom table name for Profile.
	fmt.Println("\n=== Column Info: user_profiles (custom TableName) ===")
	var profileCols []ColumnInfo
	db.Raw("PRAGMA table_info(user_profiles)").Scan(&profileCols)
	for _, c := range profileCols {
		fmt.Printf("  %-15s %-15s\n", c.Name, c.Type)
	}

	// Demonstrate that AutoMigrate is safe to run multiple times (idempotent).
	fmt.Println("\n=== Re-running AutoMigrate (idempotent) ===")
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}, &Profile{}); err != nil {
		log.Fatal("second AutoMigrate failed:", err)
	}
	fmt.Println("Second AutoMigrate completed — no errors, no duplicate tables.")

	fmt.Println("\n=== Summary of GORM Conventions ===")
	fmt.Println("  - Struct name → pluralized, snake_case table (User → users)")
	fmt.Println("  - Field name → snake_case column (UserID → user_id)")
	fmt.Println("  - gorm.Model → id, created_at, updated_at, deleted_at")
	fmt.Println("  - AutoMigrate: additive only, safe to re-run, no column drops")
	fmt.Println("  - TableName() method overrides default table naming")
	fmt.Println("  - Association tags: foreignKey sets FK column, references sets target column")
}
