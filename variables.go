package main

import "fmt"

func main() {
  var a string = "this be a string"
  fmt.Println(a)

  var b, c int = 1, 100
  fmt.Println(b, c)

  var d = 26 // inferred type int
  fmt.Println(d)

  var e = true
  fmt.Println(e)

  var f int // zero-valued without corresponding init
  fmt.Println(f)

  g := "hullo" // shorthand
  var h = "anyone out there?" // inferred type string
  fmt.Println(g)
  fmt.Println(h)
}
