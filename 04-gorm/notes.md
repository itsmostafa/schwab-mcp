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

1. When would you choose GORM over raw SQL in a Go service?
2. What does `AutoMigrate` do, and what does it intentionally not do?
3. Explain `First` vs `Take` vs `Find`, including not-found behavior.
4. Why can `Updates(struct)` skip fields unexpectedly?
5. How do you force zero-value updates safely?
6. What is the N+1 query problem, and how does `Preload` help?
7. When is `Joins` better than `Preload`?
8. How would you model `BelongsTo`, `HasMany`, and `Many2Many` in GORM?
9. What are the tradeoffs of soft deletes?
10. How does `db.Transaction` handle commit and rollback?
11. What are the risks of doing network I/O inside a DB transaction?
12. Why might disabling `SkipDefaultTransaction` be risky?
13. Where should you put business logic: hooks or service layer?
14. How do you tune GORM/database performance in production?
15. How do you structure repository methods to keep query logic testable?
16. How would you enforce tenant isolation with scopes?
17. What metrics/logs would you use to detect slow GORM queries?

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
