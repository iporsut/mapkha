package mapkha

type edgeBuilderFactory func() EdgeBuilder

type Wordcut struct {
	edgeBuilders []EdgeBuilder
}

func isSpace(ch rune) bool {
	return ch == ' ' ||
		ch == '\n' ||
		ch == '\t' ||
		ch == '"' ||
		ch == '(' ||
		ch == ')' ||
		ch == '“' ||
		ch == '”'
}

func isLatin(ch rune) bool {

	return (ch >= 'A' && ch <= 'Z') ||
		(ch >= 'a' && ch <= 'z')

}

func NewWordcut(dict PrefixTree) *Wordcut {
	factories := []edgeBuilderFactory{
		func() EdgeBuilder {
			return &PatEdgeBuilder{foundS: false,
				foundE:   false,
				edgeType: SPACE,
				isPat:    isSpace,
			}
		},
		func() EdgeBuilder {
			return &PatEdgeBuilder{foundS: false,
				foundE:   false,
				edgeType: LATIN,
				isPat:    isLatin,
			}
		},
		func() EdgeBuilder {
			return NewDictEdgeBuilder(dict)
		},
		func() EdgeBuilder {
			return &UnkEdgeBuilder{}
		}}

	w := &Wordcut{make([]EdgeBuilder, 0, 4)}
	for _, factory := range factories {
		w.edgeBuilders = append(w.edgeBuilders, factory())
	}

	return w
}

func (w *Wordcut) Reset() {
	for _, builder := range w.edgeBuilders {
		builder.Reset()
	}
}

func (w *Wordcut) Segment(text string) []string {
	w.Reset()
	textRunes := []rune(text)
	path := buildPath(textRunes, w.edgeBuilders)
	ranges := pathToRanges(path)
	tokens := make([]string, len(ranges))
	for i, r := range ranges {
		tokens[i] = string(textRunes[r.s:r.e])
	}
	return tokens
}
