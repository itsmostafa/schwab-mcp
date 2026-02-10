package main

import (
	"fmt"
	"strings"
)

// Interfaces: implicit satisfaction, type assertions, type switches, composition.
//
// Key interview points:
//   - Implicit satisfaction: a type satisfies an interface by implementing its
//     methods — no "implements" keyword. This enables decoupling and is central
//     to Go's design philosophy.
//   - interface{} (or "any" since Go 1.18) is the empty interface satisfied by
//     every type. Useful for generic containers, but lose type safety.
//   - Nil interface pitfall: an interface value is nil only when both its type
//     AND value are nil. A non-nil concrete type stored in an interface makes
//     the interface non-nil, even if the concrete value is nil.
//   - Method sets:
//       * Value receiver methods -> in method set of both T and *T.
//       * Pointer receiver methods -> in method set of *T only.
//     This matters for interface satisfaction: if an interface requires a
//     pointer-receiver method, only *T (not T) satisfies it.

func main() {
	implicitSatisfactionDemo()
	typeAssertionDemo()
	typeSwitchDemo()
	interfaceCompositionDemo()
	nilInterfacePitfallDemo()
}

// --- Interface definitions ---

// Speaker is a simple interface. Any type with a Speak() string method
// satisfies it — no declaration of intent required.
type Speaker interface {
	Speak() string
}

// Mover is another simple interface.
type Mover interface {
	Move() string
}

// Animal composes Speaker and Mover. A type must implement both
// Speak() and Move() to satisfy Animal. This is interface composition,
// analogous to io.ReadWriter = io.Reader + io.Writer.
type Animal interface {
	Speaker
	Mover
}

// --- Concrete types ---

type Dog struct{ Name string }

func (d Dog) Speak() string { return d.Name + " says: Woof!" }
func (d Dog) Move() string  { return d.Name + " runs on four legs" }

type Cat struct{ Name string }

func (c Cat) Speak() string { return c.Name + " says: Meow!" }
func (c Cat) Move() string  { return c.Name + " slinks silently" }

type Robot struct{ Model string }

// Robot only satisfies Speaker, not Mover — it cannot be used as an Animal.
func (r Robot) Speak() string { return r.Model + " says: Beep boop!" }

// --- Demos ---

func implicitSatisfactionDemo() {
	fmt.Println("=== Implicit Interface Satisfaction ===")

	// Dog and Cat satisfy Speaker without any explicit declaration.
	// The compiler checks structural compatibility at assignment time.
	var s Speaker

	s = Dog{Name: "Rex"}
	fmt.Println(" ", s.Speak())

	s = Cat{Name: "Whiskers"}
	fmt.Println(" ", s.Speak())

	s = Robot{Model: "T-800"}
	fmt.Println(" ", s.Speak())
	fmt.Println()
}

func typeAssertionDemo() {
	fmt.Println("=== Type Assertion (value, ok) ===")

	var s Speaker = Dog{Name: "Buddy"}

	// Two-value form: safe — does not panic if the assertion fails.
	if dog, ok := s.(Dog); ok {
		fmt.Println("  it's a dog:", dog.Name)
	}

	// This assertion will fail because s holds a Dog, not a Cat.
	if _, ok := s.(Cat); !ok {
		fmt.Println("  not a cat (ok=false, no panic)")
	}

	// Single-value form: s.(Cat) would panic here. Always prefer the
	// two-value form unless you're certain of the underlying type.
	fmt.Println()
}

func typeSwitchDemo() {
	fmt.Println("=== Type Switch ===")

	speakers := []Speaker{
		Dog{Name: "Rex"},
		Cat{Name: "Whiskers"},
		Robot{Model: "R2D2"},
	}

	for _, s := range speakers {
		// Type switch extracts the concrete type from an interface value.
		// Each case binds `v` to the concrete type — no explicit cast needed.
		switch v := s.(type) {
		case Dog:
			fmt.Printf("  Dog: %s (struct field access: %s)\n", v.Speak(), v.Name)
		case Cat:
			fmt.Printf("  Cat: %s (struct field access: %s)\n", v.Speak(), v.Name)
		default:
			// v is still Speaker here; we can call Speak() but not
			// access type-specific fields without another assertion.
			fmt.Printf("  Unknown: %s\n", v.Speak())
		}
	}
	fmt.Println()
}

func interfaceCompositionDemo() {
	fmt.Println("=== Interface Composition ===")

	// Dog satisfies both Speaker and Mover, so it satisfies Animal.
	var a Animal = Dog{Name: "Fido"}
	fmt.Println(" ", a.Speak())
	fmt.Println(" ", a.Move())

	// Robot only has Speak() — assigning it to Animal would be a compile error:
	// var a2 Animal = Robot{Model: "T-800"} // won't compile

	// Real-world example: io.ReadWriter composes io.Reader + io.Writer.
	// This pattern lets you build powerful abstractions from small interfaces.
	// The standard library uses this everywhere (io.ReadCloser, io.ReadWriteCloser, etc.).
	fmt.Println()

	// Demonstrate with a stdlib-like example using strings.
	fmt.Println("  Composed interface with strings.Reader:")
	r := strings.NewReader("hello interfaces")
	buf := make([]byte, 5)
	n, _ := r.Read(buf) // strings.Reader satisfies io.Reader.
	fmt.Printf("  read %d bytes: %s\n", n, buf[:n])
	fmt.Println()
}

func nilInterfacePitfallDemo() {
	fmt.Println("=== Nil Interface Pitfall ===")

	// A nil interface has no type and no value.
	var s Speaker
	fmt.Printf("  nil interface: value=%v, isNil=%t\n", s, s == nil)

	// A non-nil interface with a nil concrete value.
	// This is the classic gotcha: the interface is NOT nil because it
	// carries type information (*Dog), even though the pointer itself is nil.
	var d *Dog     // nil pointer
	s = d          // interface now has type=*Dog, value=nil
	fmt.Printf("  interface holding nil *Dog: value=%v, isNil=%t\n", s, s == nil)
	fmt.Println("  (the interface is non-nil because it carries type info)")
	fmt.Println()
}
