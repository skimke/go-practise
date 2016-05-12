package main

import "fmt"

func main() {
  // https://golang.org/ref/spec#Lexical_elements

  fmt.Println("go" + "lang")

  fmt.Println("1+1 =", 1+1)
  fmt.Println("7.0/3.0 =", 7.0/3.0) // floating-point literals
  fmt.Println("7/3 =", 7/3) // integer literals

  fmt.Println(true || false)
  fmt.Println(true && false)
  fmt.Println(!true)
}
