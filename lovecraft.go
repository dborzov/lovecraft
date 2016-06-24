package main

import (
  "fmt"
  "log"

  "github.com/dborzov/lovecraft/world"
)

var w *world.World

func init() {
  w = world.NewWorld()
  err := w.Load("world")
  if err != nil {
    log.Fatalf("Failure loading universe: %s", err)
  }
}
func main() {
  fmt.Printf("yes this works: %#v \n", w)
}
