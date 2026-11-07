# Methods on Types in Go: Summary & Examples

In Go, you can **only** add methods to custom types that you define within the **same package**.

There are no classes in GoLang we create custom type by usin type
and structs group variables and functions together mimicing class behaviour

---

## 🚫 What You CANNOT Do

- **No Built-in Types:** You cannot add methods directly to native types like `int`, `string`, `bool`, or `map`.
- **No External Packages:** You cannot add methods to types imported from other packages (e.g., `time.Time`).
- _Compiler Error:_ Attempting this results in: `cannot define new methods on non-local type`.

---

## ✅ What You CAN Do (Examples)

### 1. Adding a Method to a Custom Struct

```go
package main

import "fmt"

// Define a custom struct
type Rectangle struct {
    Width, Height int
}

// Attach a method to the struct type
func (r Rectangle) Area() int {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println(rect.Area()) // Outputs: 50
}
```

### 2. The Workaround for Built-in Types (Custom Type Definition)

If you need custom methods on a built-in type like a `string`, redefine it as a custom type in your own package:

```go
package main

import "fmt"

// Define a custom type based on a built-in string
type MyString string

// Attach the method to your custom type
func (ms MyString) Shout() string {
    return string(ms) + "!"
}

func main() {
    var greeting MyString = "hello"
    fmt.Println(greeting.Shout()) // Outputs: hello!
}
```

---

## 💡 Key Summary Rules

- **The "Same Package" Rule:** The receiver type and the method definition must live in the exact same package.
- **Methods vs. Functions:** A method is just a function with a receiver parameter written between the `func` keyword and the function name (e.g., `func (r Rectangle) Area()`).
