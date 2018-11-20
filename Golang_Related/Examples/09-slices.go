package main

import "fmt"

func main() {

// Create an empty slice with non-zero length, use the builtin make. A slice of strings of length 3 (initially zero-valued).
  s := make([]string, 3)
  fmt.Println("emp:", s)

// We can set and get just like with arrays
  s[0] = "a"
  s[1] = "b"
  s[3] = "c"
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
  fmt.Println("sl1:", l)

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