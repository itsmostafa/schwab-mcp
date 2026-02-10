# GORM

## Key Concepts

### What GORM gives you (and what it does not)

- GORM is a productive ORM for CRUD-heavy services: model mapping, query builder, associations, hooks, migrations.
- It does not remove SQL thinking. You still need to reason about indexes, joins, query plans, lock scope, and transaction boundaries.
- In interviews, a strong answer is: use GORM for speed and consistency, then drop to raw SQL for hot paths or complex reporting.

### Model definition and conventions

- Convention defaults:
  - Table name from struct name (`User` -> `users`).
  - Field names mapped to snake_case columns.
  - `ID` as primary key by default.
  - `CreatedAt`, `UpdatedAt` auto-managed; `DeletedAt` enables soft delete.
- Common tags:
  - `primaryKey`, `uniqueIndex`, `index`
  - `column:<name>`, `type:<sql type>`
  - `not null`, `default:<value>`
  - association tags: `foreignKey`, `references`, `many2many`

```go
type User struct {
    ID        uint           `gorm:"primaryKey"`
    Email     string         `gorm:"uniqueIndex;not null"`
    Name      string         `gorm:"size:120;not null"`
    CompanyID uint
    Company   Company        `gorm:"foreignKey:CompanyID;references:ID"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

### Migrations: `AutoMigrate` plus manual strategy

- `AutoMigrate` is useful for local dev and additive changes:
  - creates tables
  - adds missing columns/indexes/constraints
- Important caveat: it does not safely handle destructive changes (for example dropping columns), and behavior can vary by dialect.
- Production pattern:
  - use versioned SQL migrations (for example `goose`, `atlas`, `dbmate`)
  - run migration checks in CI
  - treat schema changes as explicit, reviewed code

### Querying essentials and common gotchas

- `First`, `Take`, `Last`:
  - single row; return `ErrRecordNotFound` when no row matches.
  - `First`/`Last` apply ordering by primary key; `Take` does not.
- `Find`:
  - for multiple rows; does not return `ErrRecordNotFound` for empty result sets.
- `Save` vs `Updates`:
  - `Save` writes all fields (upsert-like behavior based on PK presence).
  - `Updates(struct)` ignores zero-value fields by default.
  - `Updates(map[string]any)` includes provided zero values.
- `Select` and `Omit`:
  - use to explicitly control read/write field sets and avoid accidental updates.
- Scopes:
  - reusable query fragments for consistent filtering/pagination/tenancy.

```go
// Struct update ignores zero values unless selected explicitly.
db.Model(&user).Select("Name", "Age").Updates(User{Name: "", Age: 0})

// Map update includes zero values.
db.Model(&user).Updates(map[string]any{"name": "", "age": 0})
```

### Associations and N+1 prevention

- Association types:
  - `BelongsTo`, `HasOne`, `HasMany`, `Many2Many`
- N+1 problem:
  - fetching parents, then one query per parent for children.
- Preferred fix:
  - use `Preload` to eager load associations in batched queries.
- `Preload` vs `Joins`:
  - `Preload`: cleaner for relationship loading and avoiding row explosion.
  - `Joins`: better when filtering by related-table fields in one SQL query.
- Practical pattern:
  - combine `Joins` for filtering + `Preload` for complete graph loading.

```go
var users []User
err := db.
    Joins("Company").
    Where("companies.name = ?", "Acme").
    Preload("Orders").
    Preload("Orders.Items").
    Find(&users).Error
```

### Transactions and consistency boundaries

- Use `db.Transaction(func(tx *gorm.DB) error { ... })` for unit-of-work safety.
- Return error to rollback; return `nil` to commit.
- Keep transactions short: avoid slow network calls inside a transaction.
- Nested transactions are savepoint-based in GORM; useful for partial rollback logic.
- GORM wraps single writes in default transactions for safety. Disabling (`SkipDefaultTransaction`) can improve performance but increases correctness risk.

```go
err := db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&order).Error; err != nil {
        return err
    }
    if err := tx.Model(&inventory).
        Where("sku = ? AND qty >= ?", sku, order.Qty).
        Update("qty", gorm.Expr("qty - ?", order.Qty)).Error; err != nil {
        return err
    }
    return nil
})
```

### Hooks and lifecycle callbacks

- Common hooks: `BeforeCreate`, `AfterCreate`, `BeforeUpdate`, `AfterUpdate`, `BeforeDelete`, `AfterFind`, etc.
- Good use cases:
  - normalization/validation close to persistence
  - lightweight audit fields
- Risks:
  - hidden side effects and extra queries
  - performance surprises on bulk operations
- Interview note: keep business workflows in service layer; keep hooks small and predictable.

### Soft delete behavior

- With `DeletedAt`, `Delete` marks rows as deleted instead of hard-deleting.
- Default queries exclude soft-deleted rows.
- Use `Unscoped()` to include or permanently delete them.

```go
db.Delete(&user)                 // soft delete
db.Unscoped().Delete(&user)      // hard delete
db.Unscoped().Where("id = ?", 1).First(&user) // include deleted rows
```

### Performance and observability checklist

- Always propagate context: `db.WithContext(ctx)` to honor deadlines/cancellation.
- Tune `database/sql` pool via `sqlDB, _ := db.DB()`:
  - `SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime`
- Use targeted `Select` fields and pagination.
- Use batch inserts for large writes (`CreateInBatches`).
- Enable SQL logging and dry-run during debugging:
  - `db.Debug()`
  - `db.Session(&gorm.Session{DryRun: true})`

## Interview Questions

1. **When would you choose GORM over raw SQL in a Go service?**

   Use GORM for typical CRUD-heavy services where productivity matters: it handles model mapping, query building, associations, hooks, migrations, and soft deletes out of the box, reducing boilerplate. Use raw SQL (or `db.Raw`) for performance-critical hot paths, complex reporting queries with multiple joins/aggregations/window functions, or when you need precise control over the query plan. A strong production approach is GORM by default, raw SQL where profiling shows it's needed.

2. **What does `AutoMigrate` do, and what does it intentionally not do?**

   `AutoMigrate` creates tables, adds missing columns, adds missing indexes, and adds missing constraints. It intentionally does **not** drop columns, change column types destructively, or remove indexes/constraints. Behavior on some edge cases can vary by database dialect. For production, use versioned migration tools like `goose`, `atlas`, or `dbmate` where every schema change is explicit, reviewed, and reversible.

3. **Explain `First` vs `Take` vs `Find`, including not-found behavior.**

   - `First`: returns the first record ordered by primary key ascending. Returns `gorm.ErrRecordNotFound` if no match.
   - `Take`: returns one record with no ordering guarantee. Also returns `gorm.ErrRecordNotFound` if no match.
   - `Find`: returns multiple records into a slice. Does **not** return `ErrRecordNotFound` for an empty result — you get an empty slice and a nil error.

   A common mistake is using `Find` for a single record and expecting a not-found error; it won't come. Use `First` or `Take` when you need that signal.

4. **Why can `Updates(struct)` skip fields unexpectedly?**

   Go has zero values for every type (`0` for int, `""` for string, `false` for bool). When you pass a struct to `Updates`, GORM ignores any field whose value is the zero value for its type, because it cannot distinguish "intentionally set to zero" from "not set." This means `Updates(User{Age: 0})` will not update `age` to 0.

5. **How do you force zero-value updates safely?**

   Two approaches:
   - **Use a map**: `db.Model(&user).Updates(map[string]any{"age": 0, "name": ""})` — map values are always included regardless of zero value.
   - **Use `Select`**: `db.Model(&user).Select("Age", "Name").Updates(User{Age: 0, Name: ""})` — `Select` explicitly tells GORM which fields to write, overriding the zero-value filter.

6. **What is the N+1 query problem, and how does `Preload` help?**

   N+1 occurs when you fetch N parent records and then execute one additional query per parent to load its children — resulting in N+1 total queries. `Preload` solves this by issuing a separate batched query for the association (e.g., `SELECT * FROM orders WHERE user_id IN (1,2,3,...)`), reducing N+1 queries down to 2 queries regardless of N. You can chain `Preload` for nested associations (`Preload("Orders.Items")`).

7. **When is `Joins` better than `Preload`?**

   Use `Joins` when you need to **filter** parent records based on related-table columns (e.g., `WHERE companies.name = 'Acme'`) because it produces a single SQL `JOIN` query. `Preload` cannot filter the parent set by association fields — it only controls what gets loaded. The practical pattern is to combine both: `Joins` for filtering, `Preload` for loading the full association graph.

   `Joins` can also cause row explosion with `HasMany`/`Many2Many` (one parent row repeated per child row), so `Preload` is better for clean graph loading.

8. **How would you model `BelongsTo`, `HasMany`, and `Many2Many` in GORM?**

   - **BelongsTo**: The child struct holds the foreign key. E.g., `User` belongs to `Company` — `User` has `CompanyID` and a `Company` field with `gorm:"foreignKey:CompanyID"`.
   - **HasMany**: The parent declares a slice of children. E.g., `Company` has many `Users` — `Company` has `Users []User` and GORM infers the FK from the parent's name (`CompanyID` on `User`).
   - **Many2Many**: Both sides reference each other through a join table. E.g., `User` has `Languages []Language \`gorm:"many2many:user_languages"\`` — GORM creates/uses a `user_languages` join table with `user_id` and `language_id` columns.

9. **What are the tradeoffs of soft deletes?**

   Advantages: data recovery, audit trails, referential safety (no cascading hard deletes), and the ability to "undo" operations.

   Downsides:
   - Every query must include `WHERE deleted_at IS NULL`, which GORM handles automatically but adds overhead and index considerations.
   - Tables grow unbounded with deleted rows, hurting performance over time.
   - Unique constraints become tricky — a soft-deleted record still occupies the unique slot unless you use partial/filtered indexes.
   - Developers must remember to use `Unscoped()` when they genuinely need all rows, including deleted ones.
   - Data retention and compliance (GDPR) may require actual deletion, making soft delete insufficient.

10. **How does `db.Transaction` handle commit and rollback?**

    `db.Transaction(func(tx *gorm.DB) error { ... })` begins a transaction and passes the transactional `tx` handle to your function. If the function returns `nil`, GORM commits. If it returns an error (or panics), GORM rolls back. This ensures you cannot forget to commit or rollback. Nested calls to `tx.Transaction` create savepoints — rolling back an inner transaction rolls back to the savepoint, not the entire outer transaction.

11. **What are the risks of doing network I/O inside a DB transaction?**

    A database transaction holds a connection from the pool and potentially row/table locks for its entire duration. If you make a slow network call (HTTP request, RPC, message publish) inside the transaction:
    - Lock duration increases, blocking other queries and risking deadlocks.
    - Connection pool exhaustion if many requests stall on external I/O.
    - Timeouts on the external call can leave the transaction open longer than expected.
    - If the network call succeeds but the commit fails, you have an inconsistency (the external side-effect happened but the DB change didn't).

    Better pattern: do the DB work in a transaction, then perform the network call after commit (use the outbox pattern if you need atomicity across both).

12. **Why might disabling `SkipDefaultTransaction` be risky?**

    (Note: the question means "enabling `SkipDefaultTransaction: true`" — i.e., skipping default transactions.)

    By default, GORM wraps every single `Create`, `Update`, and `Delete` in an implicit transaction. Setting `SkipDefaultTransaction: true` removes that wrapper for a performance gain (~30%). The risk is that operations involving hooks (`BeforeCreate`, `AfterCreate`, etc.) are no longer atomic with the main operation. If a hook fails partway through, the database may be left in an inconsistent state because there's no transaction to roll back. Only disable this if you're confident your hooks don't need transactional guarantees, or you always use explicit `db.Transaction` blocks.

13. **Where should you put business logic: hooks or service layer?**

    Service layer. Hooks are best for small, model-scoped concerns: field normalization (e.g., lowercasing email), lightweight validation, or setting audit timestamps. Business workflows — orchestrating multiple models, calling external services, publishing events — belong in the service layer where they're explicit, testable, and visible in code review. Hooks hide side effects, making debugging harder and creating performance surprises on bulk operations (hooks run per-row).

14. **How do you tune GORM/database performance in production?**

    - **Connection pool tuning**: via `db.DB()` → `SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime` to match your DB's connection limits and avoid churn.
    - **Context propagation**: always use `db.WithContext(ctx)` so queries respect request deadlines and cancellation.
    - **Selective fields**: use `Select("id", "name")` instead of `SELECT *` to reduce I/O.
    - **Batch writes**: use `CreateInBatches` for bulk inserts.
    - **Pagination**: use `Limit`/`Offset` or cursor-based pagination to avoid full table scans.
    - **Preload strategically**: avoid loading associations you don't need.
    - **Indexing**: ensure query patterns have appropriate DB indexes; use `EXPLAIN` to verify.
    - **`SkipDefaultTransaction`**: enable selectively for read-heavy or hookless write paths.

15. **How do you structure repository methods to keep query logic testable?**

    - Define a repository interface (e.g., `UserRepository`) with methods like `FindByID`, `ListByTenant`, `Create`.
    - Implement it with a struct that holds `*gorm.DB`.
    - Accept `context.Context` in every method and use `db.WithContext(ctx)`.
    - Use scopes to compose reusable query fragments (tenant filter, soft-delete override, pagination).
    - For unit tests, either use an in-memory SQLite database (fast but dialect differences exist), or mock the interface at the service layer. Integration tests should use a real Postgres instance (e.g., via testcontainers).
    - Keep raw business logic out of queries — the repository returns data, the service layer makes decisions.

16. **How would you enforce tenant isolation with scopes?**

    Define a scope that always filters by `tenant_id`:

    ```go
    func TenantScope(tenantID uint) func(db *gorm.DB) *gorm.DB {
        return func(db *gorm.DB) *gorm.DB {
            return db.Where("tenant_id = ?", tenantID)
        }
    }
    ```

    Apply it to every query: `db.Scopes(TenantScope(tenantID)).Find(&records)`. To make it automatic, extract `tenantID` from the request context in middleware and attach the scoped `*gorm.DB` to the context so repository methods always receive a tenant-filtered DB instance. This prevents accidental cross-tenant data leaks.

17. **What metrics/logs would you use to detect slow GORM queries?**

    - **GORM logger**: configure `logger.Default.LogMode(logger.Warn)` to log queries exceeding a slow threshold. Custom loggers can emit structured logs with query duration, SQL, and affected rows.
    - **OpenTelemetry integration**: use a GORM plugin (e.g., `otelgorm`) to emit spans per query with duration, statement, and error attributes — visible in tracing backends like Jaeger or Tempo.
    - **Connection pool metrics**: expose `db.DB().Stats()` (open connections, in-use, idle, wait count, wait duration) as Prometheus metrics to detect pool exhaustion.
    - **Application-level metrics**: track p50/p95/p99 query latency per repository method. Alert on sustained increases.
    - **Database-side**: enable `pg_stat_statements` (Postgres) or slow query log (MySQL) for the DB's perspective on query performance.

Practice prompts:
- Build a repository layer with scoped queries (tenant + status + pagination) and unit tests.
- Implement an order creation flow that uses transaction + inventory decrement + outbox insert.
- Refactor a naive endpoint with N+1 into `Preload`/`Joins` and compare generated SQL.

## Resources

- GORM docs (home): https://gorm.io/docs/
- Models: https://gorm.io/docs/models.html
- Conventions: https://gorm.io/docs/conventions.html
- Create/Query/Update/Delete: https://gorm.io/docs/create.html
- Query details (`First`, `Take`, `Find`, conditions): https://gorm.io/docs/query.html
- Update behavior (`Save`, `Updates`, `Select`, `Omit`): https://gorm.io/docs/update.html
- Associations: https://gorm.io/docs/associations.html
- Preloading: https://gorm.io/docs/preload.html
- Transactions: https://gorm.io/docs/transactions.html
- Hooks: https://gorm.io/docs/hooks.html
- Migration: https://gorm.io/docs/migration.html
- DeepWiki (`go-gorm/gorm` overview): https://deepwiki.com/go-gorm/gorm
- DeepWiki page map used for this note: https://deepwiki.com/go-gorm/gorm/wiki
