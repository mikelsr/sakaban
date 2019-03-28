package fs

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
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
		t.Fatalf("mismatched prehashes '%x' and '%x' (expected)",
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
		t.Fatalf("mismatched hashes '%x' and '%x' (expected)",
			actualHash, expectedHash)
	}
}

func TestMakeTree(t *testing.T) {
	// a (dir)
	// |-b (dir)
	// |-c (dir)
	//   |- x (file)
	os.MkdirAll(filepath.Join(testDir, "maketree", "a", "b"), 0770)
	os.MkdirAll(filepath.Join(testDir, "maketree", "a", "c"), 0770)
	ioutil.WriteFile(filepath.Join(testDir, "maketree", "a", "c", "x"),
		[]byte{2, 1, 1, 3}, permissionFile)

	// // FIXME: why?
	// if _, err := MakeTree(testFailDir); err == nil {
	// 	t.Fatalf("could read from failDir: '%s'", testFailDir)
	// }

	if _, err := MakeTree(filepath.Join(testDir, "forbidden")); err == nil {
		t.Fatalf("could read from forbidden dir")
	}

	dirTree, err := MakeTree(filepath.Join(testDir, "maketree", "a"))
	if err != nil {
		t.Fatal(err)
	}

	// FIXME
	expectedHash, _ := UnHash("QmWCXYxnBMHXU3zzRAw3XxJHXjvCQqz6g2qLgnaoQpA9Hb")
	actualHash := dirTree.Hash()

	if !bytes.Equal(expectedHash, actualHash) {
		t.Fatalf("mismatched hashes '%x' and '%x' (expected)",
			actualHash, expectedHash)
	}

	if len(dirTree.Subnodes()) != 2 {
		t.Fatalf("unexpected subnode amount: got %d expected %d",
			len(dirTree.Subnodes()), 2)
	}

	// // FIXME
	// expectedHash, _ = UnHash("QmZEnzvbX2BMG6CbdXuqN75cxARPVfQSt1ZZTQLQvqDza9")
	// actualHash = dirTree.Subnodes()[1].Subnodes()[0].Hash()
	// if !bytes.Equal(expectedHash, actualHash) {
	// 	t.Fatalf("mismatched file hashes '%x' and '%x' (expected)",
	// 		actualHash, expectedHash)
	// }

}
