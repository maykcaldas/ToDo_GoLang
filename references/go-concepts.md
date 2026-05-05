# Go Concepts

This file maps Go concepts to the projects and sections where they are applied.

## Concepts

### Basic Syntax and Structure
- Project 1, Section 1
- Concept: Go programs start with a `main` package and a `main()` function. Variables, constants, and control structures like `if`, `for`, and `switch` are fundamental.
- Additional Notes:
  - **Packages**: Every Go file belongs to a package. The `main` package is special and required for executable programs.
    ```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, Go!")
    }
    ```
  - **Imports**: Use the `import` keyword to include standard or third-party libraries.
    ```go
    import (
        "fmt"
        "math"
    )
    ```
  - **Variables**: Declared with `var` or `:=` for short declarations. Constants use `const`.
    ```go
    var name string = "Go"
    age := 10
    const version = 1.0
    ```
  - **Control Structures**: 
    - `if` statements don’t require parentheses.
      ```go
      if age > 18 {
          fmt.Println("Adult")
      } else {
          fmt.Println("Minor")
      }
      ```
    - `for` is the only loop construct in Go (used for traditional loops, ranges, and infinite loops).
      ```go
      for i := 0; i < 5; i++ {
          fmt.Println(i)
      }
      ```
    - `switch` supports multiple cases and doesn’t require `break` statements.
      ```go
      switch day := "Monday"; day {
      case "Monday":
          fmt.Println("Start of the week")
      default:
          fmt.Println("Another day")
      }
      ```
  - **Functions**: Declared with `func`, and parameters and return types are explicitly typed.
    ```go
    func add(a int, b int) int {
        return a + b
    }

    func main() {
        result := add(3, 4)
        fmt.Println("Result:", result)
    }
    ```

### Functions and Methods
- Project 1, Section 3
- Concept: Methods in Go are functions attached to a type using a receiver. The receiver appears between `func` and the method name.
- Additional Notes:
  - **Value Receiver**: Operates on a copy of the type. Use when the method only reads from the receiver.
    ```go
    func (t Task) String() string {
        return fmt.Sprintf("ID: %d, Title: %s", t.id, t.title)
    }
    ```
  - **Pointer Receiver**: Operates on the original object. Use when the method modifies the receiver.
    ```go
    func (t *Task) markComplete() {
        t.status = "complete"
    }
    ```
  - **Rule of thumb**: If any method on a type uses a pointer receiver, make all methods use pointer receivers for consistency.
  - **Calling methods**: Same syntax for both — `task.String()` or `task.markComplete()`. Go handles the dereferencing automatically.

### Structs and Interfaces
- Project 1, Section 3
- Concept: Interfaces in Go are satisfied implicitly — any type that implements the required methods satisfies the interface, with no `implements` keyword needed.
- Additional Notes:
  - **fmt.Stringer Interface**: If a type implements `String() string`, the `fmt` package calls it automatically.
    ```go
    type Stringer interface {
        String() string
    }
    // fmt.Println(task) will call task.String() if it exists
    ```
  - **Implicit satisfaction**: Unlike Python's ABCs or Java's `implements`, Go interfaces are duck-typed at compile time. If your type has the right methods, it satisfies the interface — no declaration needed.

### Error Handling
- Project 1, Section 2
- Concept: Error handling in Go is done using the `error` type. Functions often return an `error` as the second value to indicate failure.
- Additional Notes:
  - **Checking for Errors**:
    ```go
    id, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("Invalid task ID")
        return
    }
    ```
  - **Nil Errors**:
    - A `nil` error means the operation was successful.
    - Always check for `err != nil` to handle errors gracefully.
  - **Common Practices**:
    - Return early when an error occurs to keep code clean and readable.
    ```go
    if err != nil {
        return err
    }
    ```

### File I/O
- Project 1, Section 4

### Command-line Arguments
- Project 1, Section 2
- Concept: The `os.Args` slice is used to parse command-line arguments. The first element is the program name, and subsequent elements are the arguments.
- Additional Notes:
  - **Accessing Arguments**:
    ```go
    args := os.Args
    fmt.Println("Program Name:", args[0])
    fmt.Println("Arguments:", args[1:])
    ```
  - **Handling Commands**:
    Use a `switch` statement to process commands:
    ```go
    switch command {
    case "add":
        fmt.Println("Add command")
    case "list":
        fmt.Println("List command")
    default:
        fmt.Println("Unknown command")
    }
    ```

### Compilation and Execution
- Project 1, Section 1
- Concept: `go run` compiles code to a temporary binary for quick testing, while `go build` creates a persistent binary for distribution.

### Slices
- Project 1, Section 2
- Concept: A slice is a dynamically-sized, flexible view into the elements of an array. Slices are used to manage collections of data in Go.
- Additional Notes:
  - **Creating Slices**:
    ```go
    fruits := []string{"apple", "banana", "cherry"}
    numbers := make([]int, 3) // Creates a slice of length 3
    ```
  - **Accessing Elements**:
    ```go
    fmt.Println(fruits[0]) // Output: apple
    ```
  - **Modifying Elements**:
    ```go
    fruits[1] = "blueberry"
    fmt.Println(fruits) // Output: [apple blueberry cherry]
    ```
  - **Appending Elements**:
    ```go
    fruits = append(fruits, "date")
    fmt.Println(fruits) // Output: [apple blueberry cherry date]
    ```
  - **Slicing a Slice**:
    ```go
    sub := fruits[1:3]
    fmt.Println(sub) // Output: [blueberry cherry]
    ```
  - **Iterating Over a Slice**:
    ```go
    for i, fruit := range fruits {
        fmt.Println(i, fruit)
    }
    ```
  - **Length and Capacity**:
    ```go
    fmt.Println(len(fruits)) // Length
    fmt.Println(cap(fruits)) // Capacity
    ```

---

(Additional concepts will be added as we progress.)