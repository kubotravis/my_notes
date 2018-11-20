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

// 
  if num := 9 ; num < 0 {
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits")
  }
}