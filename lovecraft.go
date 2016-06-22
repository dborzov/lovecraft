package main

import (
  "fmt"
  "os"

  "github.com/BurntSushi/toml"
)

// Element is a thing
type element struct {
  Name        string
  Description string
}

// World contains EVERYTHING
type World struct {
  Elements map[string]element
}

var world World

func init() {
  _, err := toml.DecodeFile("assets/elements.toml", &world)
  if err != nil {
    fmt.Printf("failed to load toml \n")
    os.Exit(1)
  }
}

func main() {
  fmt.Printf("yes this works: %#v \n", world)
}
