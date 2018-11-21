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