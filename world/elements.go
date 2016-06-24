package world

import (
  "fmt"

  "github.com/BurntSushi/toml"
)

// Size defines
const Size = 5

// Element is a thing
type element struct {
  Name        string
  Description string
}

type World struct {
  Elements map[string]*element
  Lands    [Size]*Land
  Starter  []string
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

  if len(w.Starter) != len(w.Lands) {
    return fmt.Errorf("universe toml broken: the number of starter does not match land length")
  }

  for i, elKey := range w.Starter {
    el, ok := w.Elements[elKey]
    if !ok {
      return fmt.Errorf("a starter element \"%s\" not recognized", elKey)
    }
    w.Lands[i].Add(el)
  }
  return nil
}
