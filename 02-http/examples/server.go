// server.go — ServeMux, Handler/HandlerFunc, JSON CRUD, MaxBytesReader, server timeouts
//
// Run:
//   go run server.go
//
// Test with curl:
//   curl -s localhost:8080/health
//   curl -s localhost:8080/tasks | jq
//   curl -s -X POST localhost:8080/tasks -d '{"title":"Buy milk","done":false}' | jq
//   curl -s localhost:8080/tasks/1 | jq
//   curl -s -X PUT localhost:8080/tasks/1 -d '{"title":"Buy oat milk","done":true}' | jq
//   curl -s -X DELETE localhost:8080/tasks/1 | jq
//   curl -s localhost:8080/tasks | jq

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// --- Domain types ---

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// TaskStore is an in-memory store protected by a mutex.
// In production you'd use a database, but this shows the concurrency concern:
// http.Server handles each request in its own goroutine, so shared state
// must be synchronized.
type TaskStore struct {
	mu     sync.RWMutex
	tasks  map[int]Task
	nextID int
}

func NewTaskStore() *TaskStore {
	return &TaskStore{tasks: make(map[int]Task), nextID: 1}
}

func (s *TaskStore) All() []Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		out = append(out, t)
	}
	return out
}

func (s *TaskStore) Get(id int) (Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.tasks[id]
	return t, ok
}

func (s *TaskStore) Create(t Task) Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	t.ID = s.nextID
	s.nextID++
	s.tasks[t.ID] = t
	return t
}

func (s *TaskStore) Update(id int, t Task) (Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tasks[id]; !ok {
		return Task{}, false
	}
	t.ID = id
	s.tasks[id] = t
	return t, true
}

func (s *TaskStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tasks[id]; !ok {
		return false
	}
	delete(s.tasks, id)
	return true
}

// --- JSON helpers ---

// writeJSON encodes v as JSON with the given status code.
// Using json.NewEncoder writes directly to w without an intermediate buffer,
// which is more efficient than json.Marshal + w.Write for large payloads.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		// At this point headers are already sent, so we can only log.
		log.Printf("json encode error: %v", err)
	}
}

// writeError sends a structured JSON error response.
// Consistent error shapes make life easier for API consumers.
func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

// --- http.Handler (struct) pattern ---

// HealthHandler demonstrates implementing http.Handler via a struct.
// Use the struct approach when the handler needs injected dependencies
// (DB pool, logger, config) as fields — it makes testing easy because
// you can swap implementations.
type HealthHandler struct {
	startedAt time.Time
}

// ServeHTTP satisfies the http.Handler interface.
// Interview note: the Handler interface is a single-method interface, which
// aligns with Go's preference for small interfaces.
func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"uptime": time.Since(h.startedAt).String(),
	})
}

// --- http.HandlerFunc pattern ---

// makeTaskHandler returns an http.HandlerFunc — a plain function that satisfies
// http.Handler via the adapter type. This is the most common pattern for
// lightweight handlers.
//
// Interview note: http.HandlerFunc is a type conversion, not a wrapper struct.
// The type itself has a ServeHTTP method that calls the underlying function.
// This is the adapter pattern in the standard library.
func makeTaskHandler(store *TaskStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Routing: Go 1.22+ ServeMux supports method and path-value patterns
		// like "GET /tasks/{id}", but for broad compatibility this example
		// does manual dispatch. Framework routers (Gin, Chi) handle this more
		// ergonomically.
		path := strings.TrimPrefix(r.URL.Path, "/tasks")
		path = strings.TrimPrefix(path, "/")

		if path == "" {
			// /tasks — collection endpoint
			switch r.Method {
			case http.MethodGet:
				handleListTasks(w, r, store)
			case http.MethodPost:
				handleCreateTask(w, r, store)
			default:
				// 405 tells the client which methods are valid.
				w.Header().Set("Allow", "GET, POST")
				writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			}
			return
		}

		// /tasks/{id} — single resource endpoint
		id, err := strconv.Atoi(path)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid task id")
			return
		}

		switch r.Method {
		case http.MethodGet:
			handleGetTask(w, r, store, id)
		case http.MethodPut:
			handleUpdateTask(w, r, store, id)
		case http.MethodDelete:
			handleDeleteTask(w, r, store, id)
		default:
			w.Header().Set("Allow", "GET, PUT, DELETE")
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	}
}

func handleListTasks(w http.ResponseWriter, _ *http.Request, store *TaskStore) {
	tasks := store.All()
	writeJSON(w, http.StatusOK, tasks)
}

func handleGetTask(w http.ResponseWriter, _ *http.Request, store *TaskStore, id int) {
	task, ok := store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "task not found")
		return
	}
	writeJSON(w, http.StatusOK, task)
}

func handleCreateTask(w http.ResponseWriter, r *http.Request, store *TaskStore) {
	// MaxBytesReader wraps r.Body so that reading beyond the limit returns an error.
	// This prevents a malicious client from sending a huge body that exhausts memory.
	// The limit is checked during reads, not upfront, so it works with streaming decoders.
	const maxBodySize = 1 << 20 // 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)

	var input struct {
		Title string `json:"title"`
		Done  bool   `json:"done"`
	}

	// json.NewDecoder reads from the stream and decodes in one pass.
	// For request bodies this is preferred over json.Unmarshal(io.ReadAll(...))
	// because it avoids buffering the entire body in memory first.
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&input); err != nil {
		// If MaxBytesReader triggered, err will mention "http: request body too large".
		writeError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}

	if input.Title == "" {
		// 400 for missing required fields — some APIs use 422 for semantic
		// validation errors, but 400 is simpler and widely understood.
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	task := store.Create(Task{Title: input.Title, Done: input.Done})

	// 201 Created is the correct status for a successful resource creation.
	// Including the created resource in the response body lets clients
	// use the server-assigned ID without a follow-up GET.
	writeJSON(w, http.StatusCreated, task)
}

func handleUpdateTask(w http.ResponseWriter, r *http.Request, store *TaskStore, id int) {
	const maxBodySize = 1 << 20
	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)

	var input struct {
		Title string `json:"title"`
		Done  bool   `json:"done"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	if input.Title == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	task, ok := store.Update(id, Task{Title: input.Title, Done: input.Done})
	if !ok {
		writeError(w, http.StatusNotFound, "task not found")
		return
	}
	writeJSON(w, http.StatusOK, task)
}

func handleDeleteTask(w http.ResponseWriter, _ *http.Request, store *TaskStore, id int) {
	if !store.Delete(id) {
		writeError(w, http.StatusNotFound, "task not found")
		return
	}
	// 200 with a confirmation body is clearer than 204 No Content for JSON APIs,
	// because some HTTP clients treat empty 204 bodies as errors.
	writeJSON(w, http.StatusOK, map[string]string{"deleted": strconv.Itoa(id)})
}

// --- Main ---

func main() {
	store := NewTaskStore()

	// Interview note: always create your own ServeMux instead of using
	// http.DefaultServeMux (which http.Handle / http.HandleFunc register on).
	// DefaultServeMux is a global, so any imported package can silently register
	// routes on it (e.g., net/http/pprof does this via init()). That's a
	// security risk in production. An explicit mux keeps routing under your control.
	mux := http.NewServeMux()

	// Register the struct-based handler (Handler interface).
	mux.Handle("/health", &HealthHandler{startedAt: time.Now()})

	// Register the function-based handler (HandlerFunc adapter).
	// Note: the trailing slash makes this a subtree pattern — it matches
	// /tasks, /tasks/, /tasks/1, etc.
	mux.HandleFunc("/tasks/", makeTaskHandler(store))
	// Also match the exact "/tasks" path without trailing slash.
	mux.HandleFunc("/tasks", makeTaskHandler(store))

	// Interview note on request lifecycle in net/http:
	// 1. Listener accepts TCP connection.
	// 2. Server spawns a goroutine per connection.
	// 3. Connection goroutine reads HTTP request, selects handler via ServeMux.
	// 4. Handler runs, writes response to ResponseWriter.
	// 5. If keep-alive, connection goroutine loops back to step 3.
	// 6. Server timeouts (below) bound each phase to prevent resource leaks.

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,

		// ReadTimeout covers the time from connection accept to the end of
		// reading the request body. Protects against slow clients that trickle
		// data (slowloris attacks).
		ReadTimeout: 5 * time.Second,

		// WriteTimeout covers the time from the end of request read to the end
		// of response write. If a handler takes too long, the connection is
		// closed. Note: this is a hard deadline — for per-handler timeouts with
		// a proper HTTP response, use http.TimeoutHandler instead.
		WriteTimeout: 10 * time.Second,

		// IdleTimeout controls how long keep-alive connections stay open
		// between requests. Lower values free resources faster; higher values
		// reduce TCP/TLS handshake overhead for repeat clients.
		IdleTimeout: 60 * time.Second,
	}

	fmt.Println("server listening on :8080")
	fmt.Println("try: curl localhost:8080/health")
	fmt.Println("try: curl -X POST localhost:8080/tasks -d '{\"title\":\"Learn Go\"}'")

	// ListenAndServe blocks until the server stops. In production, run this in
	// a goroutine and use graceful shutdown (see graceful_shutdown.go).
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
