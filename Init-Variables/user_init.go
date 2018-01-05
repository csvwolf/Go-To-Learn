package main

import (
  "fmt"
  "./trans"
)

var twoPi = 2 * trans.Pi

func main() {
  fmt.Printf("2*Pi = %g\n", twoPi)
  fmt.Printf("%d %d %d\n", trans.Zero, trans.First, trans.Second)
}
