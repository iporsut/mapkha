package mapkha

import "bufio"

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

var edgeBuilders [4]EdgeBuilder

func MakeEdgeBuilders(dict PrefixTree) {
	edgeBuilders = [4]EdgeBuilder{
		&PatEdgeBuilder{foundS: false,
			foundE:   false,
			edgeType: SPACE,
			isPat:    isSpace,
		},
		&PatEdgeBuilder{foundS: false,
			foundE:   false,
			edgeType: LATIN,
			isPat:    isLatin,
		},
		NewDictEdgeBuilder(dict),
		&UnkEdgeBuilder{},
	}
}

func Reset() {
	for _, builder := range edgeBuilders {
		builder.Reset()
	}
}

func Segment(bw *bufio.Writer, text string) {
	Reset()
	textRunes := []rune(text)
	path := buildPath(textRunes)
	ranges := pathToRanges(path)
	l := len(ranges)
	if l > 0 {
		r := ranges[0]
		for j := r.s; j < r.e; j++ {
			//_ = textRunes[j]
			bw.WriteRune(textRunes[j])
		}
	}
	for i := 1; i < l; i++ {
		bw.WriteRune('|')
		r := ranges[i]
		for j := r.s; j < r.e; j++ {
			bw.WriteRune(textRunes[j])
			//_ = textRunes[j]
		}
	}
	bw.WriteRune('\n')
}
