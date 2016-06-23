package world

//Planet implements graph of connections between Lands
type Planet struct {
  lands [5]Land
}

// Land contains elements
type Land struct {
  elements       []*element
  elementsLookup map[string]int
}

func NewLand() *Land {
  return &Land{
    elements:       []*element{},
    elementsLookup: map[string]int{},
  }
}

// Add element to the land
func (l *Land) Add(e *element) {
  l.elements = append(l.elements, e)
  l.elementsLookup[e.Name] = len(l.elements)
}
