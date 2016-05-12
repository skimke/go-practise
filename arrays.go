package main

import "fmt"

func main() {
  // https://golang.org/ref/spec#Array_types

  // array type includes length and type of elements within
  var array1 [5]int
  fmt.Println(array1)

  array1[4] = 100 // value assigned on an index
  fmt.Println(array1)
  fmt.Println(len(array1)) // built-in func len

  // single line assignment syntax
  array2 := [3]int{27, 99, 23}
  var array3 = [...]int{27, 99, 23} // long form, with length equal to the maximum element index plus one 
  fmt.Println(array2)
  fmt.Println(array3)

  // two dimensional array
  var array2D [2][3]int
  for i := 0; i < 2; i++ {
    for j := 1; j < 3; j++ {
      array2D[i][j] = i + j
    }
  }
  fmt.Println(array2D)
}
