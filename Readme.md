# GoLang

- Go Lang is fast , performant and efficient programming language.
- backed by google and open source community.
- directly compiled to machine code, which makes it faster than interpreted languages.
- concurrency model beats java due to light weight threads (goroutines) powerful multithreaded applications
- All devops tools are written in GoLang (Docker, Kubernetes, Terraform, etc)

![alt text](image.png)

Build time : CI-CD pipeline less cost as faster go program builds

Fast Startup : GoLang is compiled language, so it starts faster than interpreted languages like python, ruby, etc

(this is important for microservices architecture where we have multiple services running in parallel) we want new instance to spin up fast and handle requests as soon as possible. (Horizontal scaling)

Efficient usage of memory and cpu (via goroutines) : GoLang is designed to be efficient in terms of memory and CPU usage. It uses goroutines, which are lightweight threads that allow for concurrent execution of code. This means that GoLang can handle a large number of concurrent requests without consuming too much memory or CPU resources.

efficient usage of cpu (concurrency model) : GoLang's concurrency model is based on goroutines and channels, which allows for efficient communication between concurrent processes. This means that GoLang can handle multiple requests simultaneously without blocking or slowing down the application.

also when we build the compiled binary is very smaller in size

## Installation

Download GO onto your system from the official website:

`    https://go.dev/dl/`

Install this extension in vscode for go synatx highlighting and suggestions

![alt text](image-1.png)

Build go program to generate executable binary file using the following command

` go build <path_to_your_go_file>`

To run directly use

`go run <path_to_your_go_file>`

main is the entry point of a go application same like c/c++/java

# Go Pointers and "Reference Types"

**Yes, pointers work on "reference types" in Go, but you almost never need them.** 

Technically speaking, Go is strictly a **pass-by-value language**. What developers commonly call "reference types" (like **slices, maps, and channels**) are actually small struct headers that internally contain a pointer to an underlying data structure. 

Because they already hold internal pointers, passing a map or a slice to a function copies the lightweight header, but still points to the exact same underlying data. 

---

### When pointers to reference types are actually useful

You only use a pointer to a reference type if you need to **modify the header itself** (such as changing its memory address, length, or capacity).

#### 1. Slices (`*[]T`)
If you pass a slice directly to a function and use `append()`, a re-allocation might trigger. This creates a new underlying array and updates the header's length/capacity, but that update is lost to the caller. You use a pointer to ensure the caller sees the assignment.

```go
// Without a pointer: The caller's slice won't see the new element if it reallocates
func failAppend(s []int) {
    s = append(s, 42) 
}

// With a pointer: Modifies the caller's slice header directly
func successAppend(s *[]int) {
    *s = append(*s, 42) 
}
```

#### 2. Maps (`*map[K]V`)
You almost **never** need a pointer to a map (`*map[string]int`) because modifying, adding, or deleting keys alters the underlying data structure directly. The only exception is if you want to initialize a `nil` map inside a function and have that initialization reflect back to the caller:

```go
func initializeMap(m *map[string]int) {
    if *m == nil {
        *m = make(map[string]int) // Instantiates the map for the caller
    }
    (*m)["status"] = 200
}
```

---

### Summary Reference Guide

| Type | Do you need a pointer (`*`) to mutate data? | Reason |
| :--- | :--- | :--- |
| **Slices** | **Only** if changing length/capacity (`append`) | Internal header has pointer, length, and capacity. |
| **Maps** | **No** (Except to assign a brand new map) | Maps are already pointers to runtime descriptors under the hood. |
| **Channels**| **No** | Channels are handles to a shared queue pointer. |


Best docs for learning goLang are:

https://go.dev/tour/basics/2

https://gobyexample.com/

https://go.dev/doc/effective_go

https://go.dev/ref/spec


Official standard library of go package 

https://pkg.go.dev/std

https://pkg.go.dev/strings@go1.27.2#ToUpper

https://pkg.go.dev/os@go1.27.2