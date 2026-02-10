// Run: go mod init example && go get gorm.io/gorm gorm.io/driver/sqlite && go run transactions.go
//
// This file demonstrates GORM's transaction callback pattern, automatic rollback
// on error, automatic commit on nil return, and nested savepoints.

package main

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name    string  `gorm:"size:100;not null"`
	Balance float64 `gorm:"not null;default:0"`
}

type TransferLog struct {
	gorm.Model
	FromAccountID uint    `gorm:"not null"`
	ToAccountID   uint    `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
	Note          string  `gorm:"size:200"`
}

// ErrInsufficientFunds is a domain error used to trigger rollback.
var ErrInsufficientFunds = errors.New("insufficient funds")

func main() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect:", err)
	}
	db.AutoMigrate(&Account{}, &TransferLog{})

	// Seed two accounts.
	alice := Account{Name: "Alice", Balance: 100.00}
	bob := Account{Name: "Bob", Balance: 50.00}
	db.Create(&alice)
	db.Create(&bob)

	printBalances := func(label string) {
		var a, b Account
		db.First(&a, alice.ID)
		db.First(&b, bob.ID)
		fmt.Printf("  %s — Alice: $%.2f, Bob: $%.2f\n", label, a.Balance, b.Balance)
	}

	// --- Successful Transaction (commit on nil return) ---
	fmt.Println("=== Successful Transaction ===")
	printBalances("Before")

	// Interview note: db.Transaction takes a callback. Use the `tx` parameter
	// (NOT the outer `db`) for all operations inside the callback. Using `db`
	// instead of `tx` would bypass the transaction entirely — a common mistake.
	err = db.Transaction(func(tx *gorm.DB) error {
		// Debit Alice.
		if err := tx.Model(&Account{}).Where("id = ?", alice.ID).
			Update("balance", gorm.Expr("balance - ?", 30.00)).Error; err != nil {
			return err // returning error triggers automatic ROLLBACK
		}

		// Credit Bob.
		if err := tx.Model(&Account{}).Where("id = ?", bob.ID).
			Update("balance", gorm.Expr("balance + ?", 30.00)).Error; err != nil {
			return err
		}

		// Log the transfer.
		if err := tx.Create(&TransferLog{
			FromAccountID: alice.ID,
			ToAccountID:   bob.ID,
			Amount:        30.00,
			Note:          "payment for services",
		}).Error; err != nil {
			return err
		}

		// Interview note: returning nil triggers automatic COMMIT.
		// All three operations (debit, credit, log) commit atomically.
		return nil
	})
	if err != nil {
		fmt.Printf("  Transaction failed: %v\n", err)
	} else {
		fmt.Println("  Transaction committed successfully.")
	}
	printBalances("After")

	// --- Failed Transaction (rollback on error return) ---
	fmt.Println("\n=== Failed Transaction (rollback) ===")
	printBalances("Before")

	err = db.Transaction(func(tx *gorm.DB) error {
		// Debit Alice — attempting more than her balance.
		var currentAlice Account
		if err := tx.First(&currentAlice, alice.ID).Error; err != nil {
			return err
		}

		amount := 999.00
		if currentAlice.Balance < amount {
			// Interview note: returning ANY non-nil error causes GORM to ROLLBACK
			// the entire transaction. No partial changes are persisted.
			return fmt.Errorf("%w: tried $%.2f but only $%.2f available",
				ErrInsufficientFunds, amount, currentAlice.Balance)
		}

		// This line is never reached — but if it were, it would also be rolled back.
		tx.Model(&currentAlice).Update("balance", gorm.Expr("balance - ?", amount))
		return nil
	})
	if err != nil {
		fmt.Printf("  Transaction rolled back: %v\n", err)
	}
	printBalances("After (unchanged)")

	// --- Nested Transaction (Savepoints) ---
	fmt.Println("\n=== Nested Transaction (Savepoints) ===")
	printBalances("Before")

	// Interview note: nested tx.Transaction() calls create SAVEPOINTs in the DB.
	// Rolling back an inner transaction only undoes work back to the savepoint,
	// NOT the entire outer transaction. The outer transaction can still commit.
	err = db.Transaction(func(tx *gorm.DB) error {
		// Outer: debit Alice $10 (this will persist).
		if err := tx.Model(&Account{}).Where("id = ?", alice.ID).
			Update("balance", gorm.Expr("balance - ?", 10.00)).Error; err != nil {
			return err
		}
		fmt.Println("  Outer: debited Alice $10")

		// Inner (nested): try to credit Bob, but simulate a failure.
		innerErr := tx.Transaction(func(tx2 *gorm.DB) error {
			// Credit Bob $10.
			if err := tx2.Model(&Account{}).Where("id = ?", bob.ID).
				Update("balance", gorm.Expr("balance + ?", 10.00)).Error; err != nil {
				return err
			}
			fmt.Println("  Inner: credited Bob $10 (will be rolled back)")

			// Simulate a failure in the inner transaction.
			// This rolls back to the SAVEPOINT, undoing only the inner credit.
			return errors.New("inner operation failed")
		})

		if innerErr != nil {
			// Interview note: the outer transaction can inspect the inner error
			// and decide whether to continue or also fail. Here we continue,
			// so Alice's debit persists but Bob's credit is rolled back.
			fmt.Printf("  Inner transaction rolled back: %v\n", innerErr)
			fmt.Println("  Outer transaction continues...")
		}

		// Log that a partial transfer happened.
		tx.Create(&TransferLog{
			FromAccountID: alice.ID,
			ToAccountID:   bob.ID,
			Amount:        10.00,
			Note:          "partial: inner savepoint rolled back",
		})

		// Outer returns nil → COMMIT.
		return nil
	})
	if err != nil {
		fmt.Printf("  Outer transaction failed: %v\n", err)
	} else {
		fmt.Println("  Outer transaction committed.")
	}
	printBalances("After")
	fmt.Println("  Alice lost $10 (outer commit), Bob unchanged (inner rollback).")

	// --- Verify transfer logs ---
	fmt.Println("\n=== Transfer Logs ===")
	var logs []TransferLog
	db.Find(&logs)
	for _, l := range logs {
		fmt.Printf("  Log #%d: Account %d → Account %d, $%.2f (%s)\n",
			l.ID, l.FromAccountID, l.ToAccountID, l.Amount, l.Note)
	}

	fmt.Println("\n=== Summary ===")
	fmt.Println("  - db.Transaction(func(tx *gorm.DB) error {...}): callback pattern")
	fmt.Println("  - Return nil → COMMIT, return error → ROLLBACK")
	fmt.Println("  - ALWAYS use tx (not db) inside the callback")
	fmt.Println("  - Nested tx.Transaction() creates SAVEPOINTs")
	fmt.Println("  - Inner rollback undoes only inner work; outer can still commit")
	fmt.Println("  - Manual Begin/Commit/Rollback available but callback is preferred")
}
