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