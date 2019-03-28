package fs

import (
	"bytes"
	"testing"

	"bitbucket.org/mikelsr/sakaban/fs/tree"
)

func TestDir_PreHash(t *testing.T) {
	// a (dir)
	// |-b (dir)
	// |-c (dir)
	//   |- x (file)
	x := File{name: "x", content: []byte{2, 1, 1, 3}}
	x.prehash = Hash(x.content)
	c := Dir{name: "c", subnodes: []tree.Node{x}}
	b := Dir{name: "d", subnodes: []tree.Node{}}
	a := Dir{name: "a", subnodes: []tree.Node{b, c}}
	expectedHash, _ := UnHash("QmUr22JDQWnHMCp8gqLU2BnQyBxY16ueWgz8MoT8weHdEM")
	actualHash := a.PreHash()
	if !bytes.Equal(expectedHash, actualHash) {
		t.Fatalf("Mismatched prehashes '%x' and '%x' (expected)",
			actualHash, expectedHash)
	}
}

func TestDir_Hash(t *testing.T) {
	// a (dir)
	// |-b (dir)
	// |-c (dir)
	//   |- x (file)
	x := File{name: "x", content: []byte{2, 1, 1, 3}}
	x.prehash = Hash(x.content)
	c := Dir{name: "c", subnodes: []tree.Node{x}}
	b := Dir{name: "d", subnodes: []tree.Node{}}
	a := Dir{name: "a", subnodes: []tree.Node{b, c}}
	expectedHash, _ := UnHash("QmWCXYxnBMHXU3zzRAw3XxJHXjvCQqz6g2qLgnaoQpA9Hb")
	actualHash := a.Hash()
	if !bytes.Equal(expectedHash, actualHash) {
		t.Fatalf("Mismatched hashes '%x' and '%x' (expected)",
			actualHash, expectedHash)
	}
}
