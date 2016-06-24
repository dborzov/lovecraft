package world

//Planet implements graph of connections between Lands
type Planet struct {
  lands [5]Land
}

// Land contains elements
type Land struct {
  ElList []elementptr
  elHash map[elementptr]int
}

func NewLand() *Land {
  return &Land{
    ElList: []elementptr{},
    elHash: map[elementptr]int{},
  }
}

// Add element to the land
func (l *Land) Add(e elementptr) {
  l.ElList = append(l.ElList, e)
  l.elHash[e] = len(l.ElList) - 1
}

// Check if an element is present in the land
func (l *Land) Check(e elementptr) bool {
  _, ok := l.elHash[e]
  return ok
}
