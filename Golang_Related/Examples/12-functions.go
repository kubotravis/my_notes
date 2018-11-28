package main

import "fmt"


// Here’s a function that takes two ints and returns their sum as an int.

func plus(a int, b int) int {
  // Go requires explicit returns, It won’t automatically return the value of the last expression.
  return a + b
}

// Multiple consecutive parameters of the same type declaration

func plusPlus(a, b, c int) int {
  return a + b + c
}

func main() {
  // Call a function just as with name(args)
  res := plus(1, 2)
  fmt.Println("1 + 2 = ", res)

  res = plusPlus(1, 2, 3)
  fmt.Println("1 + 2 + 3 =", res)
}
