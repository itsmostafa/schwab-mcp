// Run: go mod init example && go get github.com/gin-gonic/gin && go run crud_api.go
//
// This example demonstrates Gin route groups, API versioning, path parameters,
// ShouldBindJSON, and in-memory CRUD with thread-safe map access.
//
// Test with curl:
//   curl -s localhost:8080/api/v1/users | jq
//   curl -s -X POST localhost:8080/api/v1/users -H 'Content-Type: application/json' -d '{"name":"Alice","email":"alice@example.com"}' | jq
//   curl -s localhost:8080/api/v1/users/1 | jq
//   curl -s -X PUT localhost:8080/api/v1/users/1 -H 'Content-Type: application/json' -d '{"name":"Alice Updated","email":"alice-new@example.com"}' | jq
//   curl -s -X DELETE localhost:8080/api/v1/users/1 | jq
//   curl -s localhost:8080/api/v1/users/999 | jq   # 404

package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

// User is our domain model. JSON tags control serialization; binding tags
// control validation. Gin delegates validation to go-playground/validator/v10.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"  binding:"required,min=2,max=50"`
	Email string `json:"email" binding:"required,email"`
}

// userStore is an in-memory store protected by a mutex.
// Interview note: sync.Mutex is needed because Gin handles each request in its
// own goroutine, so concurrent map reads/writes would race without it.
type userStore struct {
	mu     sync.Mutex
	users  map[int]User
	nextID int
}

func newUserStore() *userStore {
	return &userStore{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func main() {
	// gin.Default() returns an Engine with Logger and Recovery middleware
	// already attached. The Engine implements http.Handler, so it can be
	// passed to http.Server for graceful shutdown or mounted as a sub-handler.
	r := gin.Default()
	store := newUserStore()

	// Route groups let you scope a common path prefix and middleware.
	// This is the standard pattern for API versioning: each version gets
	// its own group, so you can evolve middleware per version independently.
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", store.listUsers)
		v1.GET("/users/:id", store.getUser) // :id is a path param parsed by the radix tree router
		v1.POST("/users", store.createUser)
		v1.PUT("/users/:id", store.updateUser)
		v1.DELETE("/users/:id", store.deleteUser)
	}

	// Interview note: Gin's router uses a radix tree (compressed trie) — one
	// tree per HTTP method. Route lookup is O(path length), not O(number of
	// routes). Path params like :id are extracted during tree traversal with
	// zero allocation.

	fmt.Println("Server running on :8080")
	r.Run(":8080")
}

// listUsers returns all users as a JSON array.
func (s *userStore) listUsers(c *gin.Context) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Collect values into a slice so the JSON output is an array, not a map.
	users := make([]User, 0, len(s.users))
	for _, u := range s.users {
		users = append(users, u)
	}

	// c.JSON serializes to compact JSON and sets Content-Type.
	// c.IndentedJSON is useful during development for readable output,
	// but adds overhead in production from the extra whitespace.
	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"count": len(users),
	})
}

// getUser retrieves a single user by path parameter :id.
func (s *userStore) getUser(c *gin.Context) {
	// c.Param reads path parameters set by the router.
	// The value is always a string — you must parse it yourself.
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
		return
	}

	s.mu.Lock()
	user, exists := s.users[id]
	s.mu.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user %d not found", id)})
		return
	}

	c.JSON(http.StatusOK, user)
}

// createUser decodes a JSON body into a User struct and stores it.
func (s *userStore) createUser(c *gin.Context) {
	var req User

	// ShouldBindJSON decodes the body and runs validation from binding tags.
	// Interview note: prefer ShouldBind* over Bind* (MustBind). Bind* auto-
	// writes a 400 and aborts on error, coupling binding to the response.
	// ShouldBind* returns the error so you control the response format — this
	// is critical for returning field-level validation details.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.mu.Lock()
	req.ID = s.nextID
	s.nextID++
	s.users[req.ID] = req
	s.mu.Unlock()

	// 201 Created is the correct status for successful resource creation.
	c.JSON(http.StatusCreated, req)
}

// updateUser replaces the user at the given ID with the request body.
func (s *userStore) updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
		return
	}

	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user %d not found", id)})
		return
	}

	// Preserve the ID from the path param, not the request body.
	req.ID = id
	s.users[id] = req

	c.JSON(http.StatusOK, req)
}

// deleteUser removes the user at the given ID.
func (s *userStore) deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user %d not found", id)})
		return
	}

	delete(s.users, id)

	// 200 with a confirmation message. Some APIs prefer 204 No Content
	// for deletes (no body), but 200 with a message is more informative
	// for debugging.
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("user %d deleted", id)})
}
