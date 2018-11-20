package main

import "fmt"

func main() {

// The most basic type, with a single condition
  fmt.Println("Basic loop")
  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }


// A classic initial/condition/after for loop
  fmt.Println("Second Example")
  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }

// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function
  fmt.Println("Breaking the loop")
  for {
    fmt.Println("loop")
    break
  }

// You can also continue to the next iteration of the loop
  fmt.Println("Continue basis")
  for n := 0; n <= 5; n++ {
    if n%2 == 0 {
      continue
    }
    fmt.Println(n)
  }

}