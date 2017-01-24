package mapkha

type UnkEdgeBuilder struct {
}

// Build - build dummy edge when there is no edge created.
func (builder *UnkEdgeBuilder) Build() (edge Edge, ok bool) {
	if context.Found {
		return
	}

	source := context.Path[context.LeftBoundary]

	return Edge{S: context.LeftBoundary,
		EdgeType:  UNK,
		WordCount: source.WordCount + 1,
		UnkCount:  source.UnkCount + 1}, true
}

func (builder *UnkEdgeBuilder) Reset() {}
