package main

import (
  "fmt"

  "github.com/dborzov/lovecraft/world"
)

var w *world.World

func init() {
  w = world.NewWorld()
  w.Load("world")
}
func main() {
  fmt.Printf("yes this works: %#v \n", w)
}
