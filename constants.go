package main

import "fmt"
import "math" // https://golang.org/pkg/math/

const yoda string = "a constant, this is"

func main() {
  fmt.Println(yoda)

  const n = 1000

  const d = 3e20 / n // the "e" stands for exponent
  fmt.Println(d)

  fmt.Println(int64(d)) // explicitly cast type for the numeric constant - golang.org/ref/spec#Lexical_elements under numberic types

  fmt.Println(math.Sin(n)) // in the contexts where type is required, numeric values will be given one according to the requirements. (i.e. math.Sin expects float64)

  // on line 13 the resulting value is arbitrarily determined, whereas the math package used in line 18 provides functions for precise arithmetic

  // TODO: how do you check for a value's type in Go?
}
