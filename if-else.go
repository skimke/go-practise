package main

import "fmt"

func main() {
  // https://golang.org/ref/spec#If_statements
  // note there are no ternaries in go
  if 5%2 == 0 {
    fmt.Println("5 is even, derp")
  } else {
    fmt.Println("5 is odd")
  }

  if 10%5 == 0 {
    fmt.Println("10 is divisible by 5")
  }

  // a statement before the conditional is available in all the branches
  if num := 7; num < 0 {
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "is all grown up now")
  }
}
