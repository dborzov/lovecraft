package world

import (
  "fmt"

  "github.com/BurntSushi/toml"
)

// Element is a thing
type element struct {
  Name        string
  Description string
}

type World struct {
  Elements map[string]element
}

func NewWorld() *World {
  return new(World)
}

func (w *World) Load(dirname string) error {
  _, err := toml.DecodeFile(dirname+"/universe.toml", w)
  if err != nil {
    return fmt.Errorf("cant load universe toml: %s", err)
  }
  return nil
}
