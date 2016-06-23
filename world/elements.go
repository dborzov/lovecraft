package world

import (
  "fmt"

  "github.com/BurntSushi/toml"
)

const Size = 5

// Element is a thing
type element struct {
  Name        string
  Description string
}

type World struct {
  Elements map[string]*element
  Lands    [Size]*Land
}

func NewWorld() *World {
  w := new(World)
  for i := 0; i < Size; i++ {
    w.Lands[i] = NewLand()
  }
  return w
}

func (w *World) Load(dirname string) error {
  _, err := toml.DecodeFile(dirname+"/universe.toml", w)
  if err != nil {
    return fmt.Errorf("cant load universe toml: %s", err)
  }
  return nil
}
