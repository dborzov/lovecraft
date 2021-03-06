package world

import (
  "fmt"
  "sort"
  "strings"

  "github.com/BurntSushi/toml"
)

// Size defines
const Size = 5

// Element is a thing
type element struct {
  Name        string
  Description string
  Combos      []string
  Size        int
}

type World struct {
  Elements        map[string]*element
  Lands           [Size]*Land
  Starter         []string
  Transformations map[string]string
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
    elptr, err := w.validate(elKey)
    if err != nil {
      return fmt.Errorf("starter %s", err)
    }
    w.Lands[i].Add(elptr)
  }

  w.Transformations = make(map[string]string)
  for key, el := range w.Elements {
    for _, combo := range el.Combos {
      parents := strings.Split(combo, "+")
      sort.Strings(parents)
      w.Transformations[strings.Join(parents, "+")] = key
    }
  }
  return nil
}

// Action applies a reaction among elements if there is one
func (w *World) Action(pos int, elKeys []string) (result string, success bool) {
  sort.Strings(elKeys)
  pos = pos % Size
  for _, val := range elKeys {
    if !w.Lands[pos].Check(val) {
      return val, false
    }
  }
  result, ok := w.Transformations[strings.Join(elKeys, "+")]
  if !ok {
    return strings.Join(elKeys, "+"), false
  }

  if w.Lands[pos].Check(result) {
    return result, false
  }
  el := w.Elements[result]
  width := el.Size / 2
  for cur := pos - width; cur <= pos+width; cur++ {
    ind := cur % Size
    if cur < 0 {
      ind = Size + ind
    }

    w.Lands[ind].Add(result)
  }
  return result, true
}

// validate checks if the string is a valid elementptr (element pointer)
func (w *World) validate(elptr string) (string, error) {
  _, ok := w.Elements[elptr]
  if !ok {
    return "", fmt.Errorf("element ptr not recognized: %s", elptr)
  }
  return elptr, nil
}
