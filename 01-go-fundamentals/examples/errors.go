package main

import (
	"errors"
	"fmt"
)

// Errors: sentinel errors, wrapping with %w, errors.Is, errors.As.
//
// Key interview points:
//   - Go errors are values — the error interface is just: Error() string.
//   - Wrapping with fmt.Errorf("...: %w", err) creates an error chain.
//     Use %w (not %v) to preserve the chain for unwrapping.
//   - errors.Is(err, target) walks the entire chain looking for a match.
//     This replaces err == ErrSomething for wrapped errors.
//   - errors.As(err, &target) walks the chain and finds the first error
//     matching the target type, then assigns it. This replaces type assertions
//     on wrapped errors.
//   - Since Go 1.20, errors.Join can combine multiple errors.
//   - Since Go 1.13, fmt.Errorf with %w and errors.Is/As are the standard
//     pattern — prior to that, the community used pkg/errors.

// --- Sentinel errors ---
// Sentinel errors are package-level variables used for well-known error
// conditions. Callers compare against them with errors.Is().
// Convention: name them ErrXxx.
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
)

// --- Custom error type ---
// A custom error type carries structured data beyond a message string.
// It satisfies the error interface by implementing Error() string.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: field %q %s", e.Field, e.Message)
}

func main() {
	sentinelDemo()
	wrappingDemo()
	errorsIsDemo()
	errorsAsDemo()
	multiWrapDemo()
}

func sentinelDemo() {
	fmt.Println("=== Sentinel Errors ===")

	err := findUser(0)
	if err != nil {
		fmt.Println("  findUser(0):", err)
	}

	// Direct comparison works for unwrapped sentinels.
	if err == ErrNotFound {
		fmt.Println("  direct comparison: err == ErrNotFound (true)")
	}
	fmt.Println()
}

func wrappingDemo() {
	fmt.Println("=== Error Wrapping with %w ===")

	// Wrap a sentinel with additional context.
	// %w preserves the error chain so errors.Is can find the original.
	inner := ErrNotFound
	wrapped := fmt.Errorf("repository.GetUser: %w", inner)
	doubleWrapped := fmt.Errorf("service.HandleRequest: %w", wrapped)

	fmt.Println("  inner:         ", inner)
	fmt.Println("  wrapped:       ", wrapped)
	fmt.Println("  double-wrapped:", doubleWrapped)

	// The chain: doubleWrapped -> wrapped -> ErrNotFound
	// Each layer adds context while preserving the root cause.
	fmt.Println()
}

func errorsIsDemo() {
	fmt.Println("=== errors.Is (walks the chain) ===")

	// Simulate a layered error: handler -> service -> repo -> sentinel.
	err := handleRequest()

	// errors.Is walks the entire chain to find ErrNotFound.
	// This works even though err itself is a wrapped error, not ErrNotFound.
	if errors.Is(err, ErrNotFound) {
		fmt.Println("  errors.Is(err, ErrNotFound) = true")
		fmt.Println("  full error:", err)
	}

	// Negative case: ErrUnauthorized is not in this chain.
	if !errors.Is(err, ErrUnauthorized) {
		fmt.Println("  errors.Is(err, ErrUnauthorized) = false")
	}
	fmt.Println()
}

func errorsAsDemo() {
	fmt.Println("=== errors.As (extract typed error) ===")

	err := validateAndProcess()

	// errors.As finds the first error in the chain matching the target type
	// and assigns it, giving us access to the structured fields.
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		fmt.Println("  found ValidationError in chain:")
		fmt.Printf("    field:   %s\n", valErr.Field)
		fmt.Printf("    message: %s\n", valErr.Message)
		fmt.Println("  full error:", err)
	}
	fmt.Println()
}

func multiWrapDemo() {
	fmt.Println("=== Why %w Matters (vs %v) ===")

	base := ErrNotFound

	// %w: wraps and preserves the chain — errors.Is can find the original.
	withW := fmt.Errorf("wrapped with %%w: %w", base)

	// %v: formats the error as a string — the chain is broken.
	withV := fmt.Errorf("wrapped with %%v: %v", base)

	fmt.Printf("  %%w chain intact: errors.Is = %t\n", errors.Is(withW, ErrNotFound))
	fmt.Printf("  %%v chain broken: errors.Is = %t\n", errors.Is(withV, ErrNotFound))
	fmt.Println("  Always use %w when you want callers to inspect the cause.")
	fmt.Println()
}

// --- Helper functions that build error chains ---

func findUser(id int) error {
	if id == 0 {
		return ErrNotFound
	}
	return nil
}

func handleRequest() error {
	// Each layer wraps the error with its own context.
	err := findUser(0)
	if err != nil {
		repoErr := fmt.Errorf("repo.FindUser(id=0): %w", err)
		serviceErr := fmt.Errorf("service.GetUser: %w", repoErr)
		return fmt.Errorf("handler.HandleRequest: %w", serviceErr)
	}
	return nil
}

func validateAndProcess() error {
	// Return a custom error type wrapped with context.
	valErr := &ValidationError{Field: "email", Message: "must not be empty"}
	return fmt.Errorf("service.CreateUser: %w", valErr)
}
