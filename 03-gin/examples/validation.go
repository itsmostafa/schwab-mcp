// Run: go mod init example && go get github.com/gin-gonic/gin && go run validation.go
//
// This example demonstrates how to extract and format validation errors from
// Gin's ShouldBindJSON into user-friendly, field-level JSON responses.
// Gin uses go-playground/validator/v10 under the hood.
//
// Test with curl:
//   # Valid request — passes all validation
//   curl -s -X POST localhost:8080/register -H 'Content-Type: application/json' \
//     -d '{"name":"Alice","email":"alice@example.com","age":30,"role":"admin","password":"secret123"}' | jq
//
//   # Missing required fields — triggers "required" validation
//   curl -s -X POST localhost:8080/register -H 'Content-Type: application/json' \
//     -d '{}' | jq
//
//   # Invalid email and out-of-range age — triggers "email" and "min"/"max" rules
//   curl -s -X POST localhost:8080/register -H 'Content-Type: application/json' \
//     -d '{"name":"A","email":"not-an-email","age":200,"role":"superuser","password":"ab"}' | jq
//
//   # Malformed JSON — triggers a non-validation parse error
//   curl -s -X POST localhost:8080/register -H 'Content-Type: application/json' \
//     -d '{bad json}' | jq

package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// RegisterRequest defines the expected shape of a registration payload.
//
// Interview note: Gin uses the `binding` tag (not `validate`) to drive
// validation via go-playground/validator/v10. The `binding` tag is Gin's
// convention; if you use the validator library directly outside of Gin,
// you would use the `validate` tag instead. Both use the same syntax.
type RegisterRequest struct {
	// `required` — field must be present and non-zero-value.
	// `min`/`max` on strings validates length (rune count).
	Name string `json:"name" binding:"required,min=2,max=50"`

	// `email` — must match a valid email format (RFC 5322).
	Email string `json:"email" binding:"required,email"`

	// `min`/`max` on numeric types validates the value itself, not length.
	Age int `json:"age" binding:"required,min=1,max=150"`

	// `oneof` restricts the value to a predefined set.
	Role string `json:"role" binding:"required,oneof=admin user"`

	// Combine multiple rules — at least 6 characters long.
	Password string `json:"password" binding:"required,min=6"`
}

// FieldError represents a single validation error in the API response.
// Using a struct (rather than a raw string) gives clients machine-readable
// fields they can use to display targeted error messages in a UI.
type FieldError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   any    `json:"value"`
	Message string `json:"message"`
}

func main() {
	r := gin.Default()

	r.POST("/register", registerHandler)

	fmt.Println("Server running on :8080")
	r.Run(":8080")
}

func registerHandler(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// ShouldBindJSON can return two kinds of errors:
		// 1. validator.ValidationErrors — struct field validation failures
		// 2. Other errors — malformed JSON, type mismatches, EOF, etc.
		//
		// Interview note: use errors.As (not a type assertion) to unwrap
		// the error. errors.As handles wrapped errors correctly, which is
		// important if Gin or middleware wraps the original error.
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			fieldErrors := formatValidationErrors(ve)
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "validation failed",
				"fields": fieldErrors,
			})
			return
		}

		// Non-validation error: the JSON itself is malformed.
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid request body: %s", err.Error()),
		})
		return
	}

	// If we reach here, all validation passed.
	c.JSON(http.StatusCreated, gin.H{
		"message": "registration successful",
		"user": gin.H{
			"name":  req.Name,
			"email": req.Email,
			"age":   req.Age,
			"role":  req.Role,
		},
	})
}

// formatValidationErrors converts validator.ValidationErrors into a slice of
// structured FieldError values with human-readable messages.
//
// Interview note: validator.ValidationErrors is a slice of validator.FieldError
// interfaces. Each FieldError exposes:
//   - Field()     — struct field name (Go name, not JSON name)
//   - Tag()       — which validation rule failed (e.g., "required", "email")
//   - Param()     — rule parameter (e.g., "50" for max=50, "admin user" for oneof)
//   - Value()     — the actual value that failed validation
//   - Namespace() — full path including nested structs (e.g., "RegisterRequest.Name")
func formatValidationErrors(ve validator.ValidationErrors) []FieldError {
	fieldErrors := make([]FieldError, 0, len(ve))

	for _, fe := range ve {
		fieldErrors = append(fieldErrors, FieldError{
			Field:   fe.Field(),
			Tag:     fe.Tag(),
			Value:   fe.Value(),
			Message: buildMessage(fe),
		})
	}

	return fieldErrors
}

// buildMessage translates a FieldError into a human-readable string.
// In production, you might use a map or i18n library for localization.
func buildMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fe.Field())
	case "min":
		// Param() gives the rule argument — e.g., "2" for min=2.
		return fmt.Sprintf("%s must be at least %s", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s", fe.Field(), fe.Param())
	case "oneof":
		// Param() for oneof returns the allowed values, e.g., "admin user".
		return fmt.Sprintf("%s must be one of: %s", fe.Field(), fe.Param())
	default:
		return fmt.Sprintf("%s failed on '%s' validation", fe.Field(), fe.Tag())
	}
}
