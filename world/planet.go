package world

//Planet implements graph of connections between Lands
type Planet struct {
  lands [5]Land
}

// Land contains elements
type Land struct {
  elements       []elementptr
  elementsLookup map[elementptr]int
}

func NewLand() *Land {
  return &Land{
    elements:       []elementptr{},
    elementsLookup: map[elementptr]int{},
  }
}

// Add element to the land
func (l *Land) Add(e elementptr) {
  l.elements = append(l.elements, e)
  l.elementsLookup[e] = len(l.elements) - 1
}

// Check if an element is present in the land
func (l *Land) Check(e elementptr) bool {
  _, ok := l.elementsLookup[e]
  return ok

}
