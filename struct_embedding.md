# Struct Embedding in Go

**Struct embedding** is a core feature in Go (Golang) that allows you to include one struct type inside another without specifying a field name. It enables **composition over inheritance**, letting the outer struct automatically access the fields and methods of the embedded struct.

---

## What Struct Embedding IS

### 1. Composition (Has-A Relationship)

Struct embedding builds complex types by combining smaller, reusable structs. The outer struct _contains_ the inner struct as a field whose name defaults to the type name itself.

### 2. Field and Method Promotion

When you embed a struct, its exported fields and methods are **promoted** to the outer struct. You can call them directly on the outer struct as if they belonged to it, though you can still access them through the embedded type name.

```go
package main

import "fmt"

type Engine struct {
	Horsepower int
}

func (e Engine) Start() {
	fmt.Printf("Engine with %d HP started.\n", e.Horsepower)
}

type Car struct {
	Engine // Embedded struct (no field name)
	Model  string
}

func main() {
	myCar := Car{
		Engine: Engine{Horsepower: 300},
		Model:  "Sedan",
	}

	// Promoted field and method accessed directly on Car
	fmt.Println("Horsepower:", myCar.Horsepower)
	myCar.Start()

	// Explicit access still works
	fmt.Println("Explicit HP:", myCar.Engine.Horsepower)
}
```

### 3. Implicit Interface Satisfaction

If an embedded struct implements an interface, the outer struct automatically satisfies that interface as well. This makes it easy to wrap or extend existing functionality without writing boilerplate wrapper methods.

```go
package main

import "fmt"

type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hello, I am " + p.Name
}

// Employee automatically implements Greeter because Person is embedded
type Employee struct {
	Person
	Role string
}

func SayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	emp := Employee{
		Person: Person{Name: "Alice"},
		Role:   "Developer",
	}
	SayHello(emp) // Works seamlessly
}
```

### 4. Method Shadowing (Overriding Behavior)

You can define a method on the outer struct with the same name as a method on the embedded struct. This **shadows** (hides) the inner method when called on the outer struct, while keeping the original method accessible via the embedded type name.

```go
package main

import "fmt"

type BaseLogger struct{}

func (BaseLogger) Log(msg string) {
	fmt.Println("[BASE]:", msg)
}

type CustomLogger struct {
	BaseLogger
}

// Shadows BaseLogger.Log
func (c CustomLogger) Log(msg string) {
	fmt.Println("[CUSTOM]:", msg)
	// Can still delegate to the embedded method
	c.BaseLogger.Log(msg)
}

func main() {
	logger := CustomLogger{}
	logger.Log("System boot")
}
```

---

## What Struct Embedding is NOT

### 1. It is NOT Inheritance (No Subtyping)

Go does not have classes or inheritance. Embedding `Animal` into `Dog` does **not** make `Dog` a subtype of `Animal`. You cannot pass a `Dog` to a function that expects an `Animal`.

```go
package main

type Animal struct {
	Name string
}

type Dog struct {
	Animal
	Breed string
}

func Feed(a Animal) {}

func main() {
	d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}

	// COMPILE ERROR: cannot use d (variable of type Dog) as Animal value
	// Feed(d)

	// You must explicitly pass the embedded struct:
	Feed(d.Animal)
}
```

### 2. It is NOT Dynamic Dispatch (No True Polymorphism)

When an embedded method calls another method on itself, the receiver is always the **inner embedded struct**, never the outer struct. Shadowing a method on the outer struct does not change what the inner struct calls internally.

```go
package main

import "fmt"

type Parent struct{}

func (p Parent) Identify() {
	fmt.Println("I am Parent")
}

func (p Parent) Run() {
	// Always calls Parent.Identify(), even if called via Child
	p.Identify()
}

type Child struct {
	Parent
}

func (c Child) Identify() {
	fmt.Println("I am Child")
}

func main() {
	c := Child{}
	c.Identify() // Outputs: "I am Child"
	c.Run()      // Outputs: "I am Parent" (NOT "I am Child")
}
```

### 3. It Does NOT Allow Ambiguous Access Silently

If you embed two structs that share the exact same field or method name at the same depth, accessing that name directly on the outer struct results in a **compile-time ambiguity error**.

```go
package main

type A struct {
	ID int
}

type B struct {
	ID string
}

type Combined struct {
	A
	B
}

func main() {
	c := Combined{A: A{ID: 1}, B: B{ID: "abc"}}

	// COMPILE ERROR: ambiguous selector c.ID
	// fmt.Println(c.ID)

	// Must resolve explicitly:
	println(c.A.ID)
	println(c.B.ID)
}
```

---

## Summary Comparison

| Feature                | Struct Embedding (Go)                               | Class Inheritance (OOP)                              |
| :--------------------- | :-------------------------------------------------- | :--------------------------------------------------- |
| **Relationship**       | Has-A (Composition)                                 | Is-A (Subtyping)                                     |
| **Type Compatibility** | Outer type ≠ Inner type                             | Child type can be used as Parent type                |
| **Method Resolution**  | Static (receiver is always the inner struct)        | Dynamic (virtual method dispatch / `this` / `super`) |
| **Name Collisions**    | Compile error on direct access; resolved explicitly | Overridden based on inheritance hierarchy            |
