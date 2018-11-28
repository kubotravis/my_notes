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
