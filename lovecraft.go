package main

import (
  "fmt"
  "log"
  "os"

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
  var cmd string
  for true {
    fmt.Printf("\n>")
    fmt.Scanf("%s\n", &cmd)
    command(cmd)
  }
}

func command(cmd string) {
  if cmd == "exit" {
    os.Exit(0)
  }
  fmt.Printf("command \"%s\" is not recognized", cmd)
}
