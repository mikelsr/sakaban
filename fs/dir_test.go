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
	rootpath := filepath.Join(testDir, "maketree", "a")
	os.MkdirAll(filepath.Join(rootpath, "b"), 0770)
	os.MkdirAll(filepath.Join(rootpath, "c"), 0770)
	fpath := filepath.Join(rootpath, "c", "x")
	ioutil.WriteFile(fpath, []byte{2, 1, 1, 3}, permissionFile)

	// // FIXME: why?
	// if _, err := MakeTree(testFailDir); err == nil {
	// 	t.Fatalf("could read from failDir: '%s'", testFailDir)
	// }

	if _, err := MakeTree(filepath.Join(testDir, "forbidden")); err == nil {
		t.Fatalf("could read from forbidden dir")
	}

	dirTree, err := MakeTree(rootpath)
	if err != nil {
		t.Fatal(err)
	}

	content, _ := ioutil.ReadFile(fpath)
	expectedHash := Hash(content)
	actualHash := dirTree.Subnodes()[1].Subnodes()[0].PreHash()

	if !bytes.Equal(expectedHash, actualHash) {
		t.Fatalf("mismatched hashes '%x' and '%x' (expected)",
			actualHash, expectedHash)
	}

	// create unreadable file
	ioutil.WriteFile(filepath.Join(rootpath, "c", "y"), []byte{0}, 000)
	if _, err = MakeTree(rootpath); err == nil {
		t.Fatal("read file without permissions\n")
	}

}
