package world

import "testing"

func TestAddElementToLand(t *testing.T) {
  w := getFixture(t)
  if w == nil {
    return
  }

  sun, ok := w.Elements["sun"]
  if !ok {
    t.Error("Sun element not loaded")
  }
  w.Lands[0].Add(sun)

}

func getFixture(t *testing.T) *World {
  w := NewWorld()
  err := w.Load(".")
  if err != nil {
    t.Error("Unable to load world as fixture: ", err)
    return nil
  }
  return w
}
