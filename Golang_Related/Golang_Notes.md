# 1. Hello World

We are gonna see code for classic hellow world program and how we can execute them,

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