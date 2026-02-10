// Run: go mod init example && go get gorm.io/gorm gorm.io/driver/sqlite && go run scopes.go
//
// This file demonstrates GORM scopes (reusable query fragments), scope chaining,
// soft delete behavior, and Unscoped() for bypassing the soft delete filter.

package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name     string `gorm:"size:100;not null"`
	Status   string `gorm:"size:20;not null;default:'active'"` // active, archived, draft
	TenantID uint   `gorm:"not null;index"`
	Price    float64
}

// --- Scope Definitions ---

// Interview note: a GORM scope is any function with the signature
// func(*gorm.DB) *gorm.DB. Scopes are reusable query fragments that
// you can compose together. They promote DRY, consistent query logic.

// TenantScope filters records by tenant. In a multi-tenant system, this
// prevents cross-tenant data leaks. Apply it to every query via middleware
// or a scoped DB instance extracted from request context.
func TenantScope(tenantID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("tenant_id = ?", tenantID)
	}
}

// PaginationScope adds LIMIT and OFFSET for paginated queries.
// Interview note: offset-based pagination is simple but degrades on large
// offsets. For production, consider cursor-based pagination instead.
func PaginationScope(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// StatusScope filters records by status field.
func StatusScope(status string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", status)
	}
}

// PriceRangeScope filters records within a price range.
// Demonstrates that scopes can accept any parameters — they're just closures.
func PriceRangeScope(min, max float64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price BETWEEN ? AND ?", min, max)
	}
}

func main() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect:", err)
	}
	db.AutoMigrate(&Item{})

	// Seed data across two tenants.
	items := []Item{
		// Tenant 1 items
		{Name: "Widget A", Status: "active", TenantID: 1, Price: 10.00},
		{Name: "Widget B", Status: "active", TenantID: 1, Price: 25.00},
		{Name: "Widget C", Status: "archived", TenantID: 1, Price: 15.00},
		{Name: "Widget D", Status: "active", TenantID: 1, Price: 50.00},
		{Name: "Widget E", Status: "draft", TenantID: 1, Price: 30.00},
		// Tenant 2 items
		{Name: "Gadget X", Status: "active", TenantID: 2, Price: 20.00},
		{Name: "Gadget Y", Status: "active", TenantID: 2, Price: 45.00},
		{Name: "Gadget Z", Status: "archived", TenantID: 2, Price: 35.00},
	}
	db.Create(&items)
	fmt.Println("=== Seeded 8 items across 2 tenants ===")

	// --- Single Scope ---
	fmt.Println("\n=== TenantScope(1) ===")
	var tenant1Items []Item
	db.Scopes(TenantScope(1)).Find(&tenant1Items)
	fmt.Printf("  Tenant 1 items: %d\n", len(tenant1Items))
	for _, item := range tenant1Items {
		fmt.Printf("    - %s (status=%s, price=$%.2f)\n", item.Name, item.Status, item.Price)
	}

	// --- Chaining Multiple Scopes ---
	fmt.Println("\n=== Chained Scopes: TenantScope(1) + StatusScope(\"active\") ===")
	// Interview note: scopes compose cleanly. Each scope adds its WHERE clause.
	// GORM ANDs them together. This is how you build complex queries from
	// simple, reusable building blocks.
	var activeT1 []Item
	db.Scopes(TenantScope(1), StatusScope("active")).Find(&activeT1)
	fmt.Printf("  Tenant 1 active items: %d\n", len(activeT1))
	for _, item := range activeT1 {
		fmt.Printf("    - %s ($%.2f)\n", item.Name, item.Price)
	}

	// --- Pagination Scope ---
	fmt.Println("\n=== Pagination: TenantScope(1) + StatusScope(\"active\") + Page 1, Size 2 ===")
	var page1 []Item
	db.Scopes(TenantScope(1), StatusScope("active"), PaginationScope(1, 2)).Find(&page1)
	fmt.Printf("  Page 1 (size=2): %d items\n", len(page1))
	for _, item := range page1 {
		fmt.Printf("    - %s\n", item.Name)
	}

	fmt.Println("\n=== Pagination: Page 2, Size 2 ===")
	var page2 []Item
	db.Scopes(TenantScope(1), StatusScope("active"), PaginationScope(2, 2)).Find(&page2)
	fmt.Printf("  Page 2 (size=2): %d items\n", len(page2))
	for _, item := range page2 {
		fmt.Printf("    - %s\n", item.Name)
	}

	// --- Price Range Scope ---
	fmt.Println("\n=== PriceRangeScope: Tenant 1, $10-$30 ===")
	var priced []Item
	db.Scopes(TenantScope(1), PriceRangeScope(10, 30)).Find(&priced)
	fmt.Printf("  Items in $10-$30 range: %d\n", len(priced))
	for _, item := range priced {
		fmt.Printf("    - %s ($%.2f, status=%s)\n", item.Name, item.Price, item.Status)
	}

	// --- Soft Delete ---
	fmt.Println("\n=== Soft Delete ===")
	// Delete Widget A (soft delete because Item embeds gorm.Model with DeletedAt).
	var widgetA Item
	db.Where("name = ?", "Widget A").First(&widgetA)
	db.Delete(&widgetA)
	fmt.Printf("  Soft-deleted: %s (ID=%d)\n", widgetA.Name, widgetA.ID)

	// Normal query — soft-deleted records are automatically excluded.
	// Interview note: GORM adds "WHERE deleted_at IS NULL" to every query
	// when the model has a DeletedAt field. This is invisible but always active.
	var remaining []Item
	db.Scopes(TenantScope(1)).Find(&remaining)
	fmt.Printf("\n  Normal query (Tenant 1): %d items (Widget A hidden)\n", len(remaining))
	for _, item := range remaining {
		fmt.Printf("    - %s\n", item.Name)
	}

	// --- Unscoped: Bypass Soft Delete ---
	fmt.Println("\n=== Unscoped() — Reveals Soft-Deleted Records ===")
	// Interview note: Unscoped() removes the automatic "deleted_at IS NULL" filter.
	// Use it when you need to see or recover deleted records, or for admin/audit views.
	var allIncluding []Item
	db.Unscoped().Scopes(TenantScope(1)).Find(&allIncluding)
	fmt.Printf("  Unscoped query (Tenant 1): %d items (Widget A visible)\n", len(allIncluding))
	for _, item := range allIncluding {
		deleted := ""
		if item.DeletedAt.Valid {
			deleted = " [SOFT DELETED]"
		}
		fmt.Printf("    - %s%s\n", item.Name, deleted)
	}

	// --- Permanent Delete with Unscoped ---
	fmt.Println("\n=== Permanent (Hard) Delete with Unscoped ===")
	// Interview note: to truly remove a soft-deleted record from the database,
	// use Unscoped().Delete(). Without Unscoped, Delete on an already-deleted
	// record is a no-op because the WHERE clause excludes it.
	db.Unscoped().Where("name = ?", "Widget A").Delete(&Item{})
	var afterHardDelete []Item
	db.Unscoped().Scopes(TenantScope(1)).Find(&afterHardDelete)
	fmt.Printf("  After hard delete (Tenant 1, unscoped): %d items\n", len(afterHardDelete))
	for _, item := range afterHardDelete {
		fmt.Printf("    - %s\n", item.Name)
	}
	fmt.Println("  Widget A permanently removed from database.")

	fmt.Println("\n=== Summary ===")
	fmt.Println("  - Scopes: func(*gorm.DB) *gorm.DB — reusable query fragments")
	fmt.Println("  - db.Scopes(s1, s2, s3): composes scopes with AND logic")
	fmt.Println("  - TenantScope: essential for multi-tenant data isolation")
	fmt.Println("  - PaginationScope: offset/limit (consider cursor-based for large datasets)")
	fmt.Println("  - Soft delete: automatic WHERE deleted_at IS NULL on all queries")
	fmt.Println("  - Unscoped(): bypasses soft delete filter to see/recover deleted records")
	fmt.Println("  - Unscoped().Delete(): permanently removes records from the database")
}
