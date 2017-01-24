package mapkha

var globalContext = &EdgeBuildingContext{}

func buildPath(textRunes []rune) []Edge {
	path := make([]Edge, len(textRunes)+1)
	path[0] = Edge{S: 0, EdgeType: INIT, WordCount: 0, UnkCount: 0}
	leftBoundary := 0
	for i, ch := range textRunes {
		var bestEdge Edge
		var found bool
		for _, edgeBuilder := range edgeBuilders {
			globalContext.runes = textRunes
			globalContext.Path = path
			globalContext.I = i
			globalContext.Ch = ch
			globalContext.LeftBoundary = leftBoundary
			globalContext.BestEdge = bestEdge
			globalContext.Found = found

			edge, ok := edgeBuilder.Build(globalContext)
			if !found && ok {
				found = true
				bestEdge = edge
			} else if found && ok && edge.IsBetterThan(bestEdge) {
				bestEdge = edge
			}

		}

		if !found {
			panic("bestEdge not found")
		}

		if bestEdge.EdgeType != UNK {
			leftBoundary = i + 1
		}

		path[i+1] = bestEdge
	}
	return path
}
