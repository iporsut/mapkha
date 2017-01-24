package mapkha

import (
	"reflect"
	"testing"
)

func lookup(t PrefixTree, i int, j int, ch rune) (PrefixTreePointer, bool) {
	p, f := t[PrefixTreeNode{i, j, ch}]
	return p, f
}

func TestOneCharPrefixTree(t *testing.T) {
	words := []string{"A"}
	prefixTree := MakePrefixTree(words)
	expect := PrefixTreePointer{0, true}
	testLookup(t, expect, "Expect to find 0, 0, A")(lookup(prefixTree, 0, 0, 'A'))
}

func TestOneWordPrefixTree(t *testing.T) {
	words := []string{"AB"}
	prefixTree := MakePrefixTree(words)

	var expect PrefixTreePointer

	expect = PrefixTreePointer{0, false}
	testLookup(t, expect, "Expect to find 0, 0, A")(lookup(prefixTree, 0, 0, 'A'))

	expect = PrefixTreePointer{0, true}
	testLookup(t, expect, "Expect to find 0, 1, B")(lookup(prefixTree, 0, 1, 'B'))
}

func TestTwoWordsPrefixTree(t *testing.T) {
	words := []string{"AB", "AC", "D"}
	prefixTree := MakePrefixTree(words)

	var expect PrefixTreePointer

	expect = PrefixTreePointer{0, false}
	testLookup(t, expect, "Expect to find 0, 0, A")(lookup(prefixTree, 0, 0, 'A'))

	expect = PrefixTreePointer{0, true}
	testLookup(t, expect, "Expect to find 0, 1, B")(lookup(prefixTree, 0, 1, 'B'))

	expect = PrefixTreePointer{1, true}
	testLookup(t, expect, "Expect to find 0, 1, C")(lookup(prefixTree, 0, 1, 'C'))

	expect = PrefixTreePointer{2, true}
	testLookup(t, expect, "Expect to find 0, 0, D")(lookup(prefixTree, 0, 0, 'D'))
}

func TestKaPrefixTree(t *testing.T) {
	words := []string{"กา"}
	prefixTree := MakePrefixTree(words)
	var expect PrefixTreePointer

	expect = PrefixTreePointer{0, false}
	testLookup(t, expect, "Expect to find 0, 0, ก")(lookup(prefixTree, 0, 0, 'ก'))

	expect = PrefixTreePointer{0, true}
	testLookup(t, expect, "Expect to find 0, 1, า")(lookup(prefixTree, 0, 1, 'า'))
}

func TestViaDict(t *testing.T) {
	dict, _ := LoadDefaultDict()
	var child PrefixTreePointer
	var found bool

	child, found = lookup(dict, 0, 0, 'ม')
	if !found {
		t.Errorf("Expect to find ม")
	}

	child, found = lookup(dict, child.ChildID, 1, 'า')
	if !found {
		t.Errorf("Expect to find า")
	}

	child, found = lookup(dict, child.ChildID, 2, 'ต')
	if !found {
		t.Errorf("Expect to find ต")
	}

	child, found = lookup(dict, child.ChildID, 3, 'ร')
	if !found {
		t.Errorf("Expect to find ร")
	}

	child, found = lookup(dict, child.ChildID, 4, 'า')
	if !found {
		t.Errorf("Expect to find last า")
	}

	if !child.IsFinal {
		t.Errorf("Expect last า to be final")
	}

}

func TestViaDictNotFinal(t *testing.T) {
	dict, _ := LoadDefaultDict()
	var child PrefixTreePointer

	child, _ = lookup(dict, 0, 0, 'ต')

	child, _ = lookup(dict, child.ChildID, 1, 'ร')

	if child.IsFinal {
		t.Errorf("Expect last ร not to be final")
	}

}

func testLookup(t *testing.T, expect PrefixTreePointer, msg string) func(PrefixTreePointer, bool) {
	return func(child PrefixTreePointer, found bool) {
		if !found {
			t.Errorf(msg)
		}

		if !reflect.DeepEqual(expect, child) {
			t.Errorf("Expect %q got %q", expect, child)
		}
	}
}
