package mapkha

// Edge - edge of word graph
type Edge struct {
	S         int
	EdgeType  Etype
	WordCount int
	UnkCount  int
}

// IsBetterThan - comparing this edge to another edge
func (edge *Edge) IsBetterThan(another *Edge) bool {
	if edge == nil {
		return false
	}

	if another == nil {
		return true
	}

	if (edge.UnkCount < another.UnkCount) || ((edge.UnkCount == another.UnkCount) && (edge.WordCount < another.WordCount)) {
		return true
	}

	return false
}
