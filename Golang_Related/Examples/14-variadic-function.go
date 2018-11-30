package main

import "fmt"

// Function takes arbitary no.og arguments
func sum(nums...int) {
  fmt.Println(nums, " ")
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}
func main() {
  
  // Calling the variadic function
  sum(1, 2)
  sum(1, 2, 3)

  // If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this
  nums := []int{1, 2, 3, 4}
  sum(nums...)
}
