package main

import "fmt"
import "time"

func main() {
  // https://golang.org/ref/spec#Switch_statements
  // the below are all "expression switches"

  i := 2
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  default: // default clause is not necessary
    fmt.Println(i)
  }
  // TODO: turn this into a function that takes in a numeric value

  switch time.Now().Weekday() {
    // multiple expressions in a clause
  case time.Saturday, time.Sunday:
    fmt.Println("play hard")
  default:
    fmt.Println("work hard")
  }

  t := time.Now()
  // the switch statement without an expression means "true". achieves the same results as using if/else
  switch {
  case t.Hour() < 12:
    fmt.Println("joh-eun achim!")
  default:
    fmt.Println("boa tarde :)")
  }
}
