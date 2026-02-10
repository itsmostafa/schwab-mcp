# Go Fundamentals

## Key Concepts

### Goroutines and the GMP Scheduler

Go's runtime scheduler uses the **GMP model** — three entities that work together to execute goroutines:

| Component | What it is | Key detail |
|-----------|-----------|------------|
| **G** (Goroutine) | Lightweight, user-space thread (~2–8 KB initial stack) | Created with the `go` keyword; `g` structs are pooled and reused |
| **M** (Machine) | OS thread that executes Go code | Many can exist; they block on syscalls |
| **P** (Processor) | Logical processor holding scheduler state, run queues, and mcache | Count set by `GOMAXPROCS` (defaults to number of CPU cores) |

**How scheduling works:** An M must acquire a P to run a G. Each P has a local run queue; there is also a global run queue. When a goroutine blocks (channel op, syscall, mutex), the M releases its P so another M can pick it up.

**Preemption** (since Go 1.14 — asynchronous preemption):

- **Synchronous**: runtime poisons the goroutine's stack bound so the next stack-growth check triggers a yield.
- **Asynchronous**: runtime sends an OS signal (SIGURG on Linux/macOS) to stop a goroutine at a safe-point, preventing tight loops from monopolizing a P.

**Concurrency vs Parallelism:**

- *Concurrency* — multiple goroutines are in progress (may share a single core via rapid context switching).
- *Parallelism* — multiple goroutines execute simultaneously on different cores (requires `GOMAXPROCS > 1`).

---

### Channels — Buffered vs Unbuffered, Select

Channels are typed, FIFO conduits for communicating between goroutines.

```go
ch := make(chan int)      // unbuffered
ch := make(chan int, 10)  // buffered, capacity 10
```

| | Unbuffered | Buffered |
|-|-----------|----------|
| **Send blocks until** | A receiver is ready | Buffer is full |
| **Receive blocks until** | A sender is ready | Buffer is empty |
| **Synchronization** | Send happens-before receive completes | k-th receive happens-before (k+C)-th send |

**Channel directions** restrict usage in function signatures:

```go
func produce(out chan<- int) { out <- 42 }   // send-only
func consume(in <-chan int)  { v := <-in }   // receive-only
```

**`nil` channels** block forever on both send and receive — useful to disable a `select` case dynamically.

**Closing:** `close(ch)` signals no more values. Receives on a closed channel return the zero value immediately. Sending on a closed channel panics.

**`select` statement** — waits on multiple channel operations:

```go
select {
case v := <-ch1:
    // handle ch1
case ch2 <- val:
    // handle ch2
default:
    // non-blocking fallthrough
}
```

- If multiple cases are ready, one is chosen **uniformly at random**.
- A `default` case makes the select non-blocking.
- Operands are evaluated once, in source order, before any case is selected.

---

### `context.Context` — Propagation, Cancellation, Timeouts, Values

`context.Context` carries deadlines, cancellation signals, and request-scoped values across API boundaries.

**Creating root contexts:**

| Function | Use when |
|----------|----------|
| `context.Background()` | Main function, initialization, top-level request handler |
| `context.TODO()` | Unsure which context to use yet (placeholder) |

**Deriving contexts:**

```go
ctx, cancel := context.WithCancel(parent)        // manual cancellation
ctx, cancel := context.WithTimeout(parent, 5*time.Second) // auto-cancel after duration
ctx, cancel := context.WithDeadline(parent, deadline)     // auto-cancel at wall-clock time
ctx         := context.WithValue(parent, key, val)        // attach request-scoped data
```

**Cancellation propagation:** Contexts form a tree. When a parent is cancelled, all derived children are cancelled too. The runtime traverses the child list and closes each child's `Done()` channel.

**Best practices:**

- Pass `ctx` as the **first parameter** of every function that needs it — never store it in a struct.
- Always `defer cancel()` to release resources.
- Never pass `nil` — use `context.TODO()` instead.
- Use `context.WithValue` only for request-scoped data (trace IDs, auth tokens), **not** for optional function parameters.
- Use unexported key types to avoid collisions: `type ctxKey struct{}`.

---

### Interfaces — Implicit Satisfaction, Common Patterns

Go interfaces are satisfied **implicitly** — a type implements an interface by having the right methods, with no `implements` keyword.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
// *os.File implements io.Reader because it has a matching Read method.
```

**The empty interface (`any` / `interface{}`):** Every type satisfies it. Used when you need to accept arbitrary values (e.g., `json.Unmarshal`, `fmt.Println`).

**Type assertions** extract the concrete value:

```go
r, ok := i.(io.Reader)   // safe — ok is false if i doesn't implement Reader
r := i.(io.Reader)        // unsafe — panics if i doesn't implement Reader
```

**Type switches** branch on the dynamic type:

```go
switch v := i.(type) {
case string:
    fmt.Println("string:", v)
case int:
    fmt.Println("int:", v)
default:
    fmt.Println("unknown")
}
```

**Common standard library interfaces:**

| Interface | Methods | Used for |
|-----------|---------|----------|
| `io.Reader` | `Read([]byte) (int, error)` | Reading bytes from a source |
| `io.Writer` | `Write([]byte) (int, error)` | Writing bytes to a destination |
| `io.Closer` | `Close() error` | Releasing resources |
| `error` | `Error() string` | All error values in Go |
| `fmt.Stringer` | `String() string` | Custom string representation |
| `sort.Interface` | `Len()`, `Less(i,j)`, `Swap(i,j)` | Custom sort ordering |

**Interface composition** — embed interfaces to build larger ones:

```go
type ReadWriter interface {
    io.Reader
    io.Writer
}
```

---

### Error Handling — Sentinel Errors, Wrapping, `errors.Is` / `errors.As`

Go uses the `error` interface (`Error() string`) for error values — no exceptions.

**Sentinel errors** — package-level variables for well-known error conditions:

```go
var ErrNotFound = errors.New("not found")
```

**Error wrapping** with `fmt.Errorf` and `%w` creates an error chain:

```go
return fmt.Errorf("fetching user %d: %w", id, err)
```

The `%w` verb stores a reference to the original error, accessible via `Unwrap()`.

**`errors.Is` vs `==`:**

```go
// BAD — fails if err wraps the sentinel
if err == io.EOF { ... }

// GOOD — walks the entire error chain
if errors.Is(err, io.EOF) { ... }
```

`errors.Is` also calls an optional `Is(target error) bool` method on each error in the chain, allowing custom matching logic.

**`errors.As`** — finds the first error in the chain matching a target type:

```go
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println("path:", pathErr.Path)
}
```

Target must be a non-nil pointer to an error type or interface type.

**Error handling patterns:**

- Return early on error; happy path stays un-indented.
- Wrap errors with context at each call site to build a useful chain.
- Use sentinel errors for conditions callers need to branch on.
- Use custom error types when callers need structured data about the failure.

---

### Structs, Embedding, and Composition

**Structs** are named collections of typed fields:

```go
type User struct {
    ID   int
    Name string
}
```

**Embedding** promotes an inner type's fields and methods to the outer struct:

```go
type Admin struct {
    User          // embedded — Admin.Name works directly
    Permissions []string
}
```

- Embedding is **composition, not inheritance** — there is no "is-a" relationship.
- Promoted fields cannot be used in composite literals: `Admin{User: User{ID: 1}, Permissions: nil}`.
- If two embedded types define the same field/method, accessing it is ambiguous and requires explicit qualification.

**Method sets and pointer receivers:**

| Struct `S` embeds | Methods added to `S` | Methods added to `*S` |
|-------------------|---------------------|-----------------------|
| `T` (value) | `T`'s value-receiver methods | `T` and `*T` methods |
| `*T` (pointer) | `T` and `*T` methods | `T` and `*T` methods |

Rule of thumb: if any method on `T` mutates state, all methods should use `*T` receiver, and you likely want to embed `*T`.

---

### Packages and Module Structure

**Packages** group related `.go` files in a single directory. All files in a directory must share the same `package` declaration.

**Modules** are the unit of versioning and dependency management (Go 1.11+):

| File | Purpose |
|------|---------|
| `go.mod` | Declares module path, Go version, and dependencies (`require`, `replace`, `exclude`) |
| `go.sum` | Contains cryptographic checksums of dependencies for integrity verification |

**Module path** = import path prefix for all packages in the module:

```
module github.com/yourname/project

go 1.22

require (
    github.com/gin-gonic/gin v1.9.1
)
```

**Import paths** for sub-packages are `module_path + subdirectory`:

```go
import "github.com/yourname/project/internal/repo"
```

**`internal/` packages** can only be imported by code rooted at the parent of `internal/`. This enforces encapsulation within a module.

**`init()` functions:**

- Special function executed automatically at package initialization (before `main()`).
- Each file can have multiple `init()` functions; they run in source-file order.
- Use sparingly — for registering drivers, codecs, or initializing package-level state.
- Initialization order: imported packages first (depth-first), then the importing package's `init()` functions, then `main()`.

---

## Interview Questions

**Goroutines & Scheduling**

**Q: What happens when a goroutine makes a blocking syscall?**
A: The M (OS thread) running the goroutine enters the syscall and releases its P. The runtime can then assign that P to another M (or spin up a new M) so other goroutines keep running. When the syscall returns, the M tries to reacquire a P; if none are available, the goroutine is placed on the global run queue.

**Q: How does Go prevent a CPU-bound goroutine from starving others?**
A: Since Go 1.14, the runtime uses asynchronous preemption — it sends a signal (SIGURG) to the thread running the goroutine, forcing it to yield at a safe-point. Before 1.14, preemption only happened at function prologues (synchronous), so a tight loop with no function calls could starve others.

**Q: What is the difference between concurrency and parallelism?**
A: Concurrency is about dealing with multiple things at once (structure); parallelism is about doing multiple things at once (execution). A single-core machine can be concurrent but not parallel. Go enables both: goroutines provide concurrency, and `GOMAXPROCS > 1` enables parallelism.

**Channels**

**Q: What happens if you send on a closed channel? Receive from a closed channel?**
A: Sending on a closed channel panics. Receiving from a closed channel returns the zero value immediately (and `ok` is `false` in the two-value form `v, ok := <-ch`). Buffered channels drain remaining values before returning zero values.

**Q: When would you use a `nil` channel?**
A: To dynamically disable a `select` case. Setting a channel variable to `nil` makes its case block forever, effectively removing it from selection without restructuring the `select`.

**Q: What's the difference between a buffered channel of size 1 and an unbuffered channel?**
A: An unbuffered channel requires both sender and receiver to be ready simultaneously (rendezvous). A buffered channel of size 1 allows one send to proceed without a receiver, decoupling the sender momentarily. The synchronization guarantees also differ: unbuffered guarantees send-before-receive-completes; buffered provides a weaker ordering.

**Context**

**Q: Why should you pass `context.Context` as the first function parameter instead of storing it in a struct?**
A: Contexts are request-scoped and should flow through the call chain explicitly. Storing them in a struct obscures the lifecycle, makes it easy to accidentally share a context across unrelated requests, and violates the convention the entire ecosystem relies on.

**Q: What happens if you forget to call `cancel()` on a derived context?**
A: The resources associated with the context (goroutines watching the `Done` channel, timer entries for deadlines) leak until the parent is cancelled or the program exits. Always `defer cancel()`.

**Q: How do you propagate a timeout across an HTTP call?**
A: Create a context with `context.WithTimeout`, pass it to `http.NewRequestWithContext`, and the HTTP client will respect the deadline — cancelling the request if the timeout expires.

**Interfaces**

**Q: Why does Go use implicit interface satisfaction?**
A: It enables decoupling — a package can define an interface and accept any type that happens to have the right methods, even types from other packages that have no knowledge of the interface. This supports the dependency inversion principle without import cycles.

**Q: Can a concrete type satisfy an interface with pointer receivers if you have a value (not a pointer)?**
A: No. If the method set of `T` only includes value-receiver methods, but the interface requires a pointer-receiver method, then `T` does not satisfy the interface — only `*T` does. The compiler will report an error.

**Q: What is the purpose of the empty interface (`any`)?**
A: It accepts any value because every type has at least zero methods. Used for generic containers before Go generics, JSON unmarshaling, and `fmt.Println`-style functions. With Go 1.18+ generics, prefer type parameters over `any` when the set of types is constrained.

**Error Handling**

**Q: When should you use `errors.Is` vs `errors.As`?**
A: Use `errors.Is` when you want to check if a specific sentinel error exists anywhere in the chain (e.g., `errors.Is(err, sql.ErrNoRows)`). Use `errors.As` when you need to extract a specific error type to access its fields (e.g., `*os.PathError` to read `.Path`).

**Q: What is the difference between `%v` and `%w` in `fmt.Errorf`?**
A: `%w` wraps the error — the original error is preserved in the chain and discoverable by `errors.Is`/`errors.As`. `%v` formats the error as a string only — the original error is lost and cannot be unwrapped.

**Q: How should you structure error handling in a layered application (handler → service → repo)?**
A: Each layer wraps the error with its own context using `fmt.Errorf("...: %w", err)` and returns it. The top layer (handler) decides how to present the error. Sentinel errors or custom types are defined where the condition originates. Never log-and-return — either handle it (log/respond) or wrap-and-propagate.

**Structs & Composition**

**Q: What is the difference between embedding and having a named field?**
A: Embedding promotes the inner type's methods and fields so they can be accessed directly on the outer type (e.g., `admin.Name` instead of `admin.User.Name`). A named field requires explicit access through the field name. Embedding also affects the method set of the outer type.

**Q: If two embedded types have a method with the same name, what happens?**
A: The method is ambiguous. Calling it without qualification results in a compile error. You must access it explicitly through the embedded field name (e.g., `s.TypeA.Method()` or `s.TypeB.Method()`).

**Packages & Modules**

**Q: What is the purpose of the `internal/` directory convention?**
A: Packages under `internal/` can only be imported by code within the parent tree. This lets you share code within your module without exposing it as public API. The compiler enforces this restriction.

**Q: When should you use `init()` functions, and what are the risks?**
A: Use `init()` for side-effect registrations (e.g., database/sql drivers, image codecs). Risks: implicit execution order can cause surprises, makes testing harder (no way to skip init), and can hide dependencies. Prefer explicit initialization in `main()` when possible.

---

## Resources

- [A Tour of Go](https://go.dev/tour/) — interactive introduction to all fundamentals
- [Effective Go](https://go.dev/doc/effective_go) — idiomatic patterns and conventions
- [The Go Programming Language Specification](https://go.dev/ref/spec) — authoritative reference
- [Go Blog: Share Memory By Communicating](https://go.dev/blog/codelab-share) — goroutines and channels philosophy
- [Go Blog: Go Concurrency Patterns](https://go.dev/blog/pipelines) — pipelines, fan-out/fan-in
- [Go Blog: Context](https://go.dev/blog/context) — motivation and usage of `context.Context`
- [Go Blog: Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors) — `errors.Is`, `errors.As`, wrapping
- [Go Blog: Using Go Modules](https://go.dev/blog/using-go-modules) — modules, `go.mod`, dependency management
- [Go Wiki: LearnConcurrency](https://go.dev/wiki/LearnConcurrency) — curated concurrency learning resources
- [Ardan Labs: Scheduling in Go (series)](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html) — deep dive into the GMP scheduler
- [Dave Cheney: Practical Go — Error Handling](https://dave.cheney.net/practical-go/presentations/qcon-china.html) — error handling patterns and strategies
