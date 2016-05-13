package main

import "fmt"

func main() {
  // https://golang.org/ref/spec#Slice_types

  // unlike arrays, typed just by the elements within - no length specified
  s := make([]string, 3) // built-in func make takes in the length - https://golang.org/ref/spec#Making_slices_maps_and_channels
  fmt.Println("empty:", s)

  s[0] = "one"
  s[1] = "two"
  s[2] = "three"
  fmt.Println("full:", s)
  fmt.Println("length:", len(s)) // built-in func len

  // https://golang.org/ref/spec#Appending_and_copying_slices
  s = append(s, "four")
  fmt.Println("appended:", s)
  fmt.Println("s[3]:", s[3])

  c := make([]string, len(s))
  copy(c, s) // func copy
  fmt.Println("copied:", c)

  // https://golang.org/ref/spec#Slice_expressions
  l := s[2:4]
  fmt.Println("s[2:4]:", l)
  l = s[2:3]
  fmt.Println("s[2:3]:", l)
  l = s[2:]
  fmt.Println("s[2:]:", l)

  l = s[:2]
  fmt.Println("s[:2]:", l)

  // single line variable and assignment
  t := []string{"g", "h", "i"}
  fmt.Println("declared:", t)

  // two dimensional slices
  // with arrays, length of inner dimensions could not vary. slices do allow this
  s2D := make([][]int, 3)
  fmt.Println("2D:", s2D)

  for i := 0; i < 3; i++ {
    len := i + 1 // length of each inner dimension will be incrementally larger than the index
    s2D[i] = make([]int, len) // create the inner slice

    for j := 0; j < len; j++ {
      s2D[i][j] = i + j
    }
  }
  fmt.Println("2D assigned:", s2D)
}
