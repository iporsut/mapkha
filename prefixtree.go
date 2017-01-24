package mapkha

import "sort"

// PrefixTreeNode represents node in a prefix tree
type PrefixTreeNode struct {
	NodeID int
	Offset int
	Ch     rune
}

// PrefixTreePointer is partial information of edge
type PrefixTreePointer struct {
	ChildID int
	IsFinal bool
}

// PrefixTree is a Hash-based Prefix Tree for searching words
type PrefixTree map[PrefixTreeNode]PrefixTreePointer

// MakePrefixTree is for constructing prefix tree for word with payload list
func MakePrefixTree(wordsWithPayload []string) PrefixTree {
	sort.Strings(wordsWithPayload)
	tab := make(map[PrefixTreeNode]PrefixTreePointer)

	for i, wordWithPayload := range wordsWithPayload {
		word := wordWithPayload
		rowNo := 0

		runes := []rune(word)
		for j, ch := range runes {
			isFinal := ((j + 1) == len(runes))
			node := PrefixTreeNode{rowNo, j, ch}
			child, found := tab[node]

			if !found {
				tab[node] = PrefixTreePointer{i, isFinal}
				rowNo = i
			} else {
				rowNo = child.ChildID
			}
		}
	}
	return PrefixTree(tab)
}
