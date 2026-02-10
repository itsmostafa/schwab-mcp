// Run: go mod init example && go get gorm.io/gorm gorm.io/driver/sqlite && go run queries.go
//
// This file demonstrates First/Take/Find query methods, Save vs Updates behavior,
// the zero-value gotcha, and how to work around it with map and Select.

package main

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string `gorm:"size:100;not null"`
	Email  string `gorm:"uniqueIndex;not null"`
	Age    int
	Active bool
}

func main() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect:", err)
	}
	db.AutoMigrate(&User{})

	// Seed data for querying.
	users := []User{
		{Name: "Alice", Email: "alice@example.com", Age: 30, Active: true},
		{Name: "Bob", Email: "bob@example.com", Age: 25, Active: true},
		{Name: "Charlie", Email: "charlie@example.com", Age: 0, Active: false},
	}
	db.Create(&users)
	fmt.Println("=== Seeded 3 users ===")

	// --- First: ordered by primary key, returns ErrRecordNotFound if none ---
	fmt.Println("\n=== First ===")
	var first User
	result := db.First(&first)
	fmt.Printf("First user: %s (ID=%d) | Rows: %d\n", first.Name, first.ID, result.RowsAffected)
	// Interview note: First adds "ORDER BY id ASC LIMIT 1".

	// First with a condition that matches nothing.
	var missing User
	result = db.Where("name = ?", "Nobody").First(&missing)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Interview note: always use errors.Is, not == comparison, because GORM
		// may wrap errors. This is the idiomatic Go error checking pattern.
		fmt.Println("First (no match): correctly got ErrRecordNotFound")
	}

	// --- Take: no ordering, returns ErrRecordNotFound if none ---
	fmt.Println("\n=== Take ===")
	var taken User
	result = db.Take(&taken)
	fmt.Printf("Take user: %s (ID=%d) | Rows: %d\n", taken.Name, taken.ID, result.RowsAffected)
	// Interview note: Take has no ORDER BY, so the row returned depends on the DB engine.

	var takenMissing User
	result = db.Where("age > ?", 999).Take(&takenMissing)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("Take (no match): correctly got ErrRecordNotFound")
	}

	// --- Find: returns a slice, does NOT return ErrRecordNotFound ---
	fmt.Println("\n=== Find ===")
	var allUsers []User
	result = db.Find(&allUsers)
	fmt.Printf("Find all: %d users found | Error: %v\n", result.RowsAffected, result.Error)

	// Find with a condition that matches nothing — no error, just empty slice.
	var noUsers []User
	result = db.Where("age > ?", 999).Find(&noUsers)
	fmt.Printf("Find (no match): %d users | Error: %v\n", len(noUsers), result.Error)
	// Interview note: Find NEVER returns ErrRecordNotFound. This is a common
	// interview question — don't use Find when you need not-found detection.

	// --- Save: updates ALL fields, including zero values ---
	fmt.Println("\n=== Save (updates ALL fields) ===")
	var alice User
	db.First(&alice, "name = ?", "Alice")
	fmt.Printf("Before Save: Name=%s, Age=%d, Active=%v\n", alice.Name, alice.Age, alice.Active)

	alice.Age = 0      // intentionally setting to zero
	alice.Active = false // intentionally setting to false (zero value for bool)
	db.Save(&alice)

	var aliceAfter User
	db.First(&aliceAfter, alice.ID)
	fmt.Printf("After Save:  Name=%s, Age=%d, Active=%v\n", aliceAfter.Name, aliceAfter.Age, aliceAfter.Active)
	fmt.Println("Save wrote Age=0 and Active=false — it updates ALL fields.")

	// Reset alice for the next demo.
	db.Model(&alice).Updates(map[string]any{"age": 30, "active": true})

	// --- Updates with struct: SKIPS zero-value fields (THE GOTCHA) ---
	fmt.Println("\n=== Updates with Struct (zero-value gotcha) ===")
	var bob User
	db.First(&bob, "name = ?", "Bob")
	fmt.Printf("Before: Name=%s, Age=%d, Active=%v\n", bob.Name, bob.Age, bob.Active)

	// Attempting to set Age=0 and Active=false using a struct — these will be SKIPPED.
	// Interview note: THIS IS THE #1 GORM GOTCHA. Go's zero values (0, "", false)
	// are indistinguishable from "field not set" when using struct updates. GORM
	// skips them to avoid accidentally blanking fields you didn't intend to change.
	db.Model(&bob).Updates(User{Age: 0, Active: false})

	var bobAfter User
	db.First(&bobAfter, bob.ID)
	fmt.Printf("After struct Updates: Age=%d, Active=%v  <-- NOT updated! Zero values skipped.\n",
		bobAfter.Age, bobAfter.Active)

	// --- Updates with map: includes zero values ---
	fmt.Println("\n=== Updates with Map (zero values included) ===")
	// Interview note: use map[string]any when you need to set fields to their
	// zero value. Map keys are treated as column names; all values are written.
	db.Model(&bob).Updates(map[string]any{"age": 0, "active": false})

	var bobMap User
	db.First(&bobMap, bob.ID)
	fmt.Printf("After map Updates: Age=%d, Active=%v  <-- Updated correctly!\n",
		bobMap.Age, bobMap.Active)

	// Reset bob for the next demo.
	db.Model(&bob).Updates(map[string]any{"age": 25, "active": true})

	// --- Select + Updates: explicitly include zero-value fields ---
	fmt.Println("\n=== Select + Updates (explicit zero-value fields) ===")
	// Interview note: Select tells GORM exactly which fields to write, bypassing
	// the zero-value filter. This is the type-safe alternative to using a map.
	db.Model(&bob).Select("Age", "Active").Updates(User{Age: 0, Active: false})

	var bobSelect User
	db.First(&bobSelect, bob.ID)
	fmt.Printf("After Select+Updates: Age=%d, Active=%v  <-- Updated correctly!\n",
		bobSelect.Age, bobSelect.Active)

	// --- Always check RowsAffected ---
	fmt.Println("\n=== RowsAffected Check ===")
	result = db.Model(&User{}).Where("id = ?", 9999).Updates(map[string]any{"name": "Ghost"})
	fmt.Printf("Update nonexistent ID: RowsAffected=%d, Error=%v\n", result.RowsAffected, result.Error)
	// Interview note: a zero RowsAffected with nil error means the query ran fine
	// but matched nothing. Always check RowsAffected when the caller needs to know
	// whether the target row actually existed.

	fmt.Println("\n=== Summary ===")
	fmt.Println("  - First/Take: return ErrRecordNotFound when no match")
	fmt.Println("  - Find: returns empty slice, NO ErrRecordNotFound")
	fmt.Println("  - Save: writes ALL fields including zero values")
	fmt.Println("  - Updates(struct): SKIPS zero-value fields (the gotcha!)")
	fmt.Println("  - Updates(map): includes zero values")
	fmt.Println("  - Select + Updates: type-safe way to include zero values")
	fmt.Println("  - Always check RowsAffected for update/delete operations")
}
