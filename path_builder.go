package mapkha

var context = EdgeBuildingContext{}

func buildPath(textRunes []rune) []Edge {
	path := make([]Edge, len(textRunes)+1)
	path[0] = Edge{S: 0, EdgeType: INIT, WordCount: 0, UnkCount: 0}
	context.LeftBoundary = 0
	for i, ch := range textRunes {
		var bestEdge Edge
		var found bool
		for _, edgeBuilder := range edgeBuilders {
			context.runes = textRunes
			context.Path = path
			context.I = i
			context.Ch = ch
			context.BestEdge = bestEdge
			context.Found = found

			edge, ok := edgeBuilder.Build()
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
			context.LeftBoundary = i + 1
		}

		path[i+1] = bestEdge
	}
	return path
}
