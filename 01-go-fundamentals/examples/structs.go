package main

import "fmt"

// Structs: embedding, field promotion, method sets, pointer receivers.
//
// Key interview points:
//   - Go has no inheritance. Embedding is composition, not subclassing.
//     The embedded type's methods are "promoted" but there is no polymorphism
//     between the outer and inner types.
//   - Field promotion: fields and methods of an embedded struct can be accessed
//     directly on the outer struct (syntactic sugar — the compiler rewrites it).
//   - Method sets and interface satisfaction:
//       * A value of type T has methods with value receivers only.
//       * A pointer *T has methods with both value AND pointer receivers.
//     This means: if an interface requires a pointer-receiver method,
//     only *T satisfies it, not T. This is a very common interview question.
//   - Why the distinction? A value receiver gets a copy, so it can't modify
//     the original. A pointer receiver can. If you could call a pointer-receiver
//     method on an unaddressable value, modifications would silently disappear.

func main() {
	embeddingDemo()
	fieldPromotionDemo()
	methodPromotionDemo()
	methodSetDemo()
}

// --- Types ---

// Base holds common fields. It will be embedded in other structs.
type Base struct {
	ID   int
	Name string
}

// Describe is a value receiver method on Base.
func (b Base) Describe() string {
	return fmt.Sprintf("Base{ID: %d, Name: %q}", b.ID, b.Name)
}

// Employee embeds Base. This is composition, not inheritance.
// Employee "has a" Base, and Base's fields and methods are promoted.
type Employee struct {
	Base               // Embedded (anonymous) field — enables promotion.
	Role     string
	Salary   float64
}

// --- Pointer vs Value Receiver ---

type Counter struct {
	count int
}

// Increment uses a pointer receiver because it mutates the struct.
// This method is in the method set of *Counter, but NOT Counter.
func (c *Counter) Increment() {
	c.count++
}

// Value uses a value receiver — it only reads the struct.
// This method is in the method set of both Counter and *Counter.
func (c Counter) Value() int {
	return c.count
}

// Incrementer is an interface requiring a pointer-receiver method.
type Incrementer interface {
	Increment()
}

// Valuer is an interface requiring only a value-receiver method.
type Valuer interface {
	Value() int
}

// --- Demos ---

func embeddingDemo() {
	fmt.Println("=== Struct Embedding (Composition) ===")

	emp := Employee{
		Base:   Base{ID: 1, Name: "Alice"},
		Role:   "Engineer",
		Salary: 120000,
	}

	// The embedded Base is accessible as a field named "Base".
	fmt.Printf("  employee: %+v\n", emp)
	fmt.Printf("  embedded base: %+v\n", emp.Base)
	fmt.Println()
}

func fieldPromotionDemo() {
	fmt.Println("=== Field Promotion ===")

	emp := Employee{
		Base: Base{ID: 2, Name: "Bob"},
		Role: "Manager",
	}

	// Promoted fields: access Base's fields directly on Employee.
	// emp.Name is syntactic sugar for emp.Base.Name.
	fmt.Println("  emp.Name (promoted):", emp.Name)
	fmt.Println("  emp.ID   (promoted):", emp.ID)
	fmt.Println("  emp.Role (own field):", emp.Role)

	// If Employee had its own Name field, it would shadow Base.Name.
	// You'd then need emp.Base.Name to reach the embedded one.
	fmt.Println()
}

func methodPromotionDemo() {
	fmt.Println("=== Method Promotion ===")

	emp := Employee{
		Base: Base{ID: 3, Name: "Carol"},
		Role: "SRE",
	}

	// Describe() is defined on Base but promoted to Employee.
	// emp.Describe() is rewritten by the compiler to emp.Base.Describe().
	fmt.Println("  emp.Describe():", emp.Describe())

	// Note: if Employee defined its own Describe(), it would shadow Base's.
	// This is NOT method overriding (no vtable, no polymorphism).
	fmt.Println()
}

func methodSetDemo() {
	fmt.Println("=== Method Sets: Pointer vs Value Receiver ===")

	c := Counter{count: 0}

	// On a variable (addressable), Go auto-takes the address, so both
	// value and pointer receiver methods work. This is compiler convenience.
	c.Increment() // Compiler rewrites to (&c).Increment()
	c.Increment()
	fmt.Println("  counter value after 2 increments:", c.Value())

	// But for INTERFACE SATISFACTION, method sets are strict:

	// *Counter satisfies Incrementer (pointer receiver method).
	var inc Incrementer = &c // Works: *Counter has Increment() in its method set.
	inc.Increment()
	fmt.Println("  counter value after Incrementer.Increment():", c.Value())

	// Counter (value) does NOT satisfy Incrementer:
	// var inc2 Incrementer = c  // COMPILE ERROR: Counter does not implement Incrementer
	// Why? Because you can't reliably take the address of all values
	// (e.g., map values, return values). So Go enforces the rule strictly.

	// Both Counter and *Counter satisfy Valuer (value receiver method).
	var v1 Valuer = c  // Value type works.
	var v2 Valuer = &c // Pointer type also works.
	fmt.Printf("  Valuer from value: %d, from pointer: %d\n", v1.Value(), v2.Value())

	fmt.Println()
	fmt.Println("  Summary:")
	fmt.Println("    Value receiver  -> method set of T and *T")
	fmt.Println("    Pointer receiver -> method set of *T only")
	fmt.Println("    Interface check uses method sets strictly (no auto-addressing)")
	fmt.Println()
}
