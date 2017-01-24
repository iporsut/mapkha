package mapkha

type DictEdgeBuilder struct {
	dict     PrefixTree
	pointers []dictBuilderPointer
}

type dictBuilderPointer struct {
	NodeID  int
	S       int
	Offset  int
	IsFinal bool
}

func NewDictEdgeBuilder(dict PrefixTree) *DictEdgeBuilder {
	return &DictEdgeBuilder{dict: dict}
}

// Build - build new edge from dictionary
func (builder *DictEdgeBuilder) Build(context *EdgeBuildingContext) (edge Edge, ok bool) {
	if isSpace(context.Ch) || isLatin(context.Ch) {
		return
	}

	builder.pointers = append(builder.pointers, dictBuilderPointer{S: context.I})

	newIndex := 0
	for i, _ := range builder.pointers {
		p := builder.pointers[i]
		childNode, found := builder.dict[PrefixTreeNode{p.NodeID, p.Offset, context.Ch}]
		if !found {
			continue
		}
		p.Offset++
		p.NodeID = childNode.ChildID
		p.IsFinal = childNode.IsFinal
		builder.pointers[newIndex] = p
		newIndex++
	}

	builder.pointers = builder.pointers[:newIndex]

	for _, pointer := range builder.pointers {
		if pointer.IsFinal {
			s := 1 + context.I - pointer.Offset
			source := context.Path[s]
			e := Edge{
				S:         s,
				EdgeType:  DICT,
				WordCount: source.WordCount + 1,
				UnkCount:  source.UnkCount}
			if !ok {
				edge = e
				ok = true
			} else if !edge.IsBetterThan(e) {
				edge = e
			}
		}
	}

	return
}

func (builder *DictEdgeBuilder) Reset() {
	builder.pointers = builder.pointers[:0]
}
