package world

//Planet implements graph of connections between Lands
type Planet struct {
  lands [5]Land
}

// Land contains elements
type Land struct {
  ElList []string
  elHash map[string]int
}

func NewLand() *Land {
  return &Land{
    ElList: []string{},
    elHash: map[string]int{},
  }
}

// Add element to the land
func (l *Land) Add(e string) {
  l.ElList = append(l.ElList, e)
  l.elHash[e] = len(l.ElList) - 1
}

// Check if an element is present in the land
func (l *Land) Check(e string) bool {
  _, ok := l.elHash[e]
  return ok
}
