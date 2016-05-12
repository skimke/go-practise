package main

import "fmt"

func main() {
  // https://golang.org/ref/spec#For_statements
  i := 1
  for i <= 3 {
    fmt.Println("single condition", i)
    i = i + 1
  }

  for j := 1; j <= 3; j++ {
    fmt.Println("init/condition/after", j)
  }

  for {
    fmt.Println("no condition, get out!")
    break
  }
}
