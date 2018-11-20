# Golang by Example

## Table of Contents

[1. Hello World](https://github.com/kubotravis/my_notes/blob/master/Golang_Related/Golang_Notes.md#1-hello-world)
[2. Values](https://github.com/kubotravis/my_notes/blob/master/Golang_Related/Golang_Notes.md#1-values)
[3](https://github.com/kubotravis/my_notes/blob/master/Golang_Related/Golang_Notes.md#3-variables)
[4](https://github.com/kubotravis/my_notes/blob/master/Golang_Related/Golang_Notes.md#4-constants)
[5](https://github.com/kubotravis/my_notes/blob/master/Golang_Related/Golang_Notes.md#5-for)

# 1. Hello World

We are gonna see code for classic hello world program and how we can execute them,

Below is the typical code hello world, save the below content into `hello-world.go`

```go
package main

import "fmt"

func main(){
  fmt.Println("Hello World")
}
```

Run the code with compiling (on demand with go bin)

```
$ go run hello-world.go
hello world
```

Build the binary from the source

```
$ go build hello-world.go
$ ls 
hello-world hello-world.go 
```

Now run the binary

```
$ ./hello-world
```

# 2 . Values

Its has various type of value types and they are,
- Strings, Intergers, Booleans, Floats and etc

```go
package main

import "fmt"

func main() {

// Strings can be added together with +
  fmt.Println("go" + "lang")

// Integer and Float
  fmt.Println("2+2", 3+3)
  fmt.Println("9.0/4.0", 9/4)

// Boolean operator
  fmt.Println(true && false)
  fmt.Println(true || false)
  fmt.Println(!true)
}
```

Lets execute the same,

```
$ go run 02-values.go  
golang
2+2 6
9.0/4.0 2
false
true
false
```

# 3. Variables

In go, Variables are explicitly declared and used by compiles, e.g - check the type-correctness of the function calls

```go
package main

import "fmt"

func main() {

// var declares 1 or more variables.
  var a = "this"
  fmt.Println(a)

// multiple var can be declared at once
  var b, c int = 1, 2
  fmt.Println(b, c)

// Go will infer the type of initialized variables
  var d = true
  fmt.Println(d)

// variabled declared with any value initialized with 0 (its int 0)
  var e int
  fmt.Println(e)

  f := "shortdeclaration"
  fmt.Println(f)
}
```

**Note**
The `:=` syntax is shorthand for declaring and initializing a variable, e.g. for `var f string = "short"` in this case.

```
$ go run 03-variables.go  
this
1 2
true
0
shortdeclaration
```

# 04. Constants

Go supports the `constants` of character, string, boolean and numeric values.

```go
package main

import "fmt"
import "math"

// const declares a contanst value
const s string = "constant"

func main() {
  fmt.Println(s)

  // It can appear anywhere a var statement can
  const n = 500000000

  // Constant expressions perform arithmetic with arbitrary precision
  const d = 3e20 / n
  fmt.Println(d)

  // A numeric constant has no type until it’s given one, such as by an explicit cast
  fmt.Println(int64(d))

  // A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64
  fmt.Println(math.Sin(n))
}
```

```
$ go run 04-constants.go  
constant
6e+11
600000000000
-0.28470407323754404
```

# 05. For

`for` is the only Go's looping construct.

```go
package main

import "fmt"

func main() {

// The most basic type, with a single condition
  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

// A classic initial/condition/after for loop
  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }

// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function
  for {
    fmt.Println("loop")
    break
  }

// You can also continue to the next iteration of the loop
  for n := 0; n <= 5; n++ {
    if n%2 == 0 {
      continue
    }
    fmt.Println(n)
  }

}
```

```
$ go run 05-for.go  
Basic loop
1
2
3
Second Example
7
8
9
Breaking the loop
loop
Continue basis
1
3
5
```