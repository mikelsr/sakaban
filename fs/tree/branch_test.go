package tree

import (
	"bytes"
	"crypto/rand"
	"path/filepath"
	"testing"
)

func TestBranch(t *testing.T) {
	// "/"
	// |-"a/"
	// |  | - "a1.raw"
	// |  | - "a2.raw"
	leafContent1 := make([]byte, 1024)
	leafContent2 := make([]byte, 1024)
	rand.Read(leafContent1)
	rand.Read(leafContent2)

	a1 := leafNode{
		name:    "a1.raw",
		prehash: hashFn(leafContent1),
	}
	a2 := leafNode{
		name:    "a2.raw",
		prehash: hashFn(leafContent1),
	}
	a := treeNode{
		name:     "a",
		subnodes: []Node{a1, a2},
	}
	root := treeNode{
		name:     "",
		subnodes: []Node{a},
	}

	broot := []Node{root}
	ba := []Node{root, a}
	ba1 := []Node{root, a, a1}
	ba2 := []Node{root, a, a2}

	expected := []Node{root, a, a1, a2}
	branches := []Branch{broot, ba, ba1, ba2}

	for i := 0; i < len(expected); i++ {
		if branches[i].Name() != expected[i].Name() {
			t.Fatalf("Mismatched names %s, %s",
				branches[i].Name(), expected[i].Name())
		}
		if !bytes.Equal(branches[i].PreHash(), expected[i].PreHash()) {
			t.Fatalf("Mismatched prehashes %x, %x",
				branches[i].PreHash(), expected[i].PreHash())
		}
		if !bytes.Equal(branches[i].Hash(), expected[i].Hash()) {
			t.Fatalf("Mismatched hashes %x, %x",
				branches[i].Hash(), expected[i].Hash())
		}
		if branches[i].IsDir() != expected[i].IsDir() {
			t.Fatalf("Mismatched directory types %t, %t",
				branches[i].IsDir(), expected[i].IsDir())
		}
		if len(branches[i].Subnodes()) != len(expected[i].Subnodes()) {
			t.Fatalf("Different subnode amount %d, %d",
				len(branches[i].Subnodes()),
				len(expected[i].Subnodes()))
		}
	}
	expectedPaths := []string{
		filepath.Join(""),
		filepath.Join("", "a"),
		filepath.Join("", "a", "a1"),
		filepath.Join("", "a", "a2"),
	}
	for i, b := range branches {
		if b.Path() != expectedPaths[i] {
			t.Fatalf("Mismatched paths %s, %s",
				b.Path(), expectedPaths[i])
		}
	}
}
