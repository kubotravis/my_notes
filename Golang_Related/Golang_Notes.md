# Golang by Example

## Table of Contents

[1. Hello World](https://github.com/kubotravis/my_notes/blob/master/Golang_Related/Golang_Notes.md#1-hello-world)

[2. Values](https://github.com/kubotravis/my_notes/blob/master/Golang_Related/Golang_Notes.md#2-values)

[3. Variables](https://github.com/kubotravis/my_notes/blob/master/Golang_Related/Golang_Notes.md#3-variables)

[4. Constants](https://github.com/kubotravis/my_notes/blob/master/Golang_Related/Golang_Notes.md#4-constants)

[5. For](https://github.com/kubotravis/my_notes/blob/master/Golang_Related/Golang_Notes.md#5-for)

[6. If/Else](https://github.com/kubotravis/my_notes/blob/master/Golang_Related/Golang_Notes.md#6-ifelse)

[7.Switch](https://github.com/kubotravis/my_notes/blob/master/Golang_Related/Golang_Notes.md#7-switch)

[8.Arrays](https://github.com/kubotravis/my_notes/blob/master/Golang_Related/Golang_Notes.md#8-arrays)

[9.Slices](https://github.com/kubotravis/my_notes/blob/master/Golang_Related/Golang_Notes.md#9-slices)

[10.Maps](https://github.com/kubotravis/my_notes/blob/master/Golang_Related/Golang_Notes.md#10-maps)


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

# 2. Values

Its has various type of value types and they are,
- Strings, Intergers, Booleans, Floats and etc

```go
package main

import "fmt"

func main() {

// Strings can be added together with +
  fmt.Println("go" + "lang")

// Integer and Float
  fmt.Println("3+3 =", 3+3)
  fmt.Println("9.0/4.0 =", 9/4)

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

In go, Variables are explicitly declared and used by compilers, e.g - check the type-correctness of the function calls

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

// variable declared with any value initialized with 0 (its int 0)
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

# 4. Constants

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

# 5. For

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
**Note**
Other for forms later when we look at `range` statements, channels, and other data structures.

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

# 6. If/Else

`if` and `else` in go very straight forward

```go
package main

import "fmt"


func main() {

// Here’s a basic example.
  if 7%2 == 0 {
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is odd")
  }

// You can have an if statement without an else
  if 8%2 == 0 {
    fmt.Println("8 is divided by 4")
  }

// A statement can precede conditionals; any variables declared in this statement are available in all branches.
  if num := 9 ; num < 0 {
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits")
  }
}
```

```
$ go run 06-if-else.go  
7 is odd
8 is divided by 4
9 has 1 digit
```
**Note**
- Note that you don’t need parentheses around conditions in Go, but that the braces are required.
- There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.

# 7. Switch

`Switch` statements express conditionals across many branches.

```go
package main

import "fmt"
import "time"

func main() {

// Here’s a basic switch.
  i := 2
  fmt.Println("Write ", i, " as ")
  switch i {
    case 1:
      fmt.Println("one")
    case 2:
      fmt.Println("two")
    case 3:
      fmt.Println("three")
  }

// commas to separate multiple expressions in the same case statement, the optional default case also used here
  switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
      fmt.Println("Its the Weekend")
    default:
      fmt.Println("Its the Weekday")
  }

// switch without an expression is an alternate way to express if/else logic, the case expressions can be non-constants as below
  t := time.Now()
  switch {
    case t.Hour() < 12:
      fmt.Println("Its before noon")
    default:
      fmt.Print("Its after noon")
  }

// A type switch compares types instead of values
// This to discover the type of an interface value, Below the variable t will have the type corresponding to its clause
  whatAmI := func(i interface{}) {
    switch t := i.(type) {
      case bool:
        fmt.Println("Im a Bool")
      case int:
        fmt.Println("Im a Int")
      default:
        fmt.Println("Dont know type %T\n", t)
    }
  }
  whatAmI(true)
  whatAmI(1)
  whatAmI("hey")
}
```

**Note**
Output format little different than expected

```
$ go run 07-switch.go  
Write  2  as 
two
Its the Weekday
Its before noon
Im a Bool
Im a Int
Dont know type %T
 hey
```

# 8. Arrays

In Go, an array is a numbered sequence of elements of a specific length.

```go
package main

import "fmt"

func main() {

// An array a that will hold exactly 5 ints.
// The type of elements and length are both part of the array’s type.
// By default an array is zero-valued, which for ints means 0s.
  var a [5]int
  fmt.Println("emp:", a)

// We can set a value at an index using the array[index] = value syntax, and get a value with array[index].
  a[4] = 100
  fmt.Println("set:", a)
  fmt.Println("get:", a[4])

// The builtin len returns the length of an array
  fmt.Println("len:", len(a))

// Syntax to declare and initialize an array in one line
  b := [5]int{1, 2, 3, 4, 5}
  fmt.Println("dcl:", b)

// Array types are one-dimensional, but we can compose types to build multi-dimensional data structures.
  var twoD [2][3]int
  for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2D: ", twoD)
}
```

```
$ go run 08-arrays.go  
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2D:  [[0 1 2] [1 2 3]]
```

# 9. Slices
- `Slices` are a key data type in Go, giving a more powerful interface to sequences than arrays.

- A slice is a convenient, flexible and powerful wrapper on top of an array. Slices do not own any data on their own. They are the just references to existing arrays.

- Unlike arrays, slices are typed only by the elements they contain (not the number of elements).

- In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values.

- Slices can also be copy’d.

- Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4].

- We can declare and initialize a variable for slice in a single line as well.

- Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
```go
package main

import "fmt"

func main() {

// Create an empty slice with non-zero length, use the builtin make. A slice of strings of length 3 (initially zero-valued).
  s := make([]string, 3)
  fmt.Println("emp:", s)

// We can set and get just like with arrays
  s[0] = "a"
  s[1] = "b"
  s[2] = "c"
  fmt.Println("set:", s)
  fmt.Println("get:", s[2])

// len returns the length of the slice as expected
  fmt.Println("len:", len(s))

// Note that we need to accept a return value from append as we may get a new slice value.
  s = append(s, "d")
  s = append(s, "e", "f")
  fmt.Println("apnd:", s)

//Here we create an empty slice c of the same length as s and copy into c from s
  c := make([]string, len(s))
  copy(c, s)
  fmt.Println("cpy:", c)

// gets a slice of the elements s[2], s[3], and s[4]
  l := s[2:5]
  fmt.Println("sl1", l)

// This slices up to (but excluding) s[5].
  l = s[:5]
  fmt.Println("sl2:", l)

// And this slices up from (and including) s[2].
  l = s[2:]
  fmt.Println("sl3:", l)

// Declare and initialize a variable for slice in a single line
  t := []string{"g", "h", "i"}
  fmt.Println("dcl:", t)

// multi-dimensional data
  twoD := make([][]int,3)
  for i := 0; i < 3; i++ {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := 0; j < innerLen; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d: ", twoD)
}
```

```
$ go run 09-slices.go
emp: [  ]
set: [a b c]
get: c
len: 3
apnd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]
```

# 10. Maps
Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).

- To create an empty map, use the builtin make: `make(map[key-type]val-type)`

```go
package main

import "fmt"

func main() {

// create an empty map, use the builtin make
  m := make(map[string]int)

// Set key/value pairs
  m["k1"] = 7
  m["k2"] = 13

// Printing a map
  fmt.Println("map:", m)

//Get a value for a key with name[key]
  v1 := m["k1"]
  fmt.Println("v1:", v1)

// len returns the number of key/value
  fmt.Println("len:", len(m))

// delete removes key/value pairs
  delete(m, "k2")
  fmt.Println("map:", m)

// The optional second return value when getting a value from a map indicates if the key was present in the map.
// This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
// Here we didn’t need the value itself, so we ignored it with the blank identifier _.
  _, prs := m["k2"]
  fmt.Println("prs:", prs)

// Declare and initialize a new map
  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", n)
}

```

```
$ go run 10-maps.go  
map: map[k1:7 k2:13]
v1: 7
len: 2
map: map[k1:7]
prs: false
map: map[bar:2 foo:1]
```

# 11. Range

- `Range` is the iteration over the variety of data structures

- Lets go the sample code

```go
package main

import "fmt"

func main() {

// Here we use range to sum the numbers in a slice. Arrays work like this too
  nums := []int{2, 3, 4}
  sum := 0

  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum)
// range on arrays and slices provides both the index and value for each entry
// Above we didn’t need the index, so we ignored it with the blank identifier _.
// Sometimes we actually want the indexes though.
  for i, num := range nums {
    if num == 3 {
      fmt.Println("index:", i)
    }
  }

// range on map iterates over key/value pairs
  kvs := map[string]string{"a": "apple", "b" : "banana", "c" : "cat"}
  for k, v := range kvs {
    fmt.Println("%s -> %s\n", k ,v)
  }
// range can also iterate over just the keys of a map.
  for k := range kvs {
    fmt.Println("Key :", k)
  }

// range on strings iterates over Unicode code points
// The first value is the starting byte index of the rune and the second the rune itself.
  for i, c := range "go" {
    fmt.Println(i, c)
  }
}
```

```
$ go run Examples/11-range.go  
sum: 9
index: 1
%s -> %s
 a apple
%s -> %s
 b banana
%s -> %s
 c cat
Key : a
Key : b
Key : c
0 103
1 111
```
