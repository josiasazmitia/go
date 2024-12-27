# Go Lang

## Introduction to Go Lang Programming

Go, also known as Golang, is an open-source programming language developed by Google. It is designed to be simple, efficient, and reliable, making it ideal for building scalable and high-performance applications. Go is statically typed and compiled, with a syntax similar to C, but with memory safety, garbage collection, and structural typing.

### Basics of Go Lang

#### Variables
In Go, variables are explicitly declared and used by the compiler to check the type-correctness of function calls.

```go
var a string = "initial"
var b, c int = 1, 2
d := true
```

#### Functions
Functions in Go are defined using the `func` keyword. They can accept parameters and return values.

```go
func add(x int, y int) int {
  return x + y
}
```

#### Control Structures
Go supports standard control structures like `if`, `for`, and `switch`.

```go
if x > 10 {
  fmt.Println("x is greater than 10")
}

for i := 0; i < 10; i++ {
  fmt.Println(i)
}

switch day {
case "Monday":
  fmt.Println("Start of the work week")
default:
  fmt.Println("Another day")
}
```

#### Arrays and Slices
Arrays in Go have a fixed size, while slices are dynamically-sized, more flexible views into arrays.

```go
var a [5]int
a[2] = 7

s := []int{2, 3, 5, 7, 11, 13}
```

#### Maps
Maps are Go's built-in associative data type (sometimes called hashes or dicts in other languages).

```go
m := make(map[string]int)
m["key"] = 42
```

### Hello World
```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}
```

### Run Go Program
```bash
go run hello.go
```

### Debug Go Program
```bash
go build hello.go
./hello
```