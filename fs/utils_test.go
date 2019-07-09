package fs

import (
	"bytes"
	"crypto/sha256"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"gitlab.com/mikelsr/sakaban/fs/tree"
)

func TestHash(t *testing.T) {
	data := []byte("let's run some tests")
	expectedHash := sha256.Sum256(data)
	hash := Hash(data)
	if !bytes.Equal(hash, expectedHash[:]) {
		t.Fatalf("mimsmatched hashes '%s' and '%s' (expected))",
			hash, expectedHash)
	}
}

func TestMultiHash(t *testing.T) {
	data := []byte("let's run some tests")
	expectedHash := "QmfEJKzW6rH4AfuN1VKuPLeKi4nBnWsFg2qFzf6X4RR4fG"
	sha256Hash := sha256.Sum256(data)
	hash := MultiHash(sha256Hash[: /*[32]byte to []byte*/])
	if hash != expectedHash {
		t.Fatalf("mimsmatched hashes '%s' and '%s' (expected))",
			hash, expectedHash)
	}
}

func TestUnHash(t *testing.T) {
	data := []byte("let's run some tests")
	encodedHash := "QmfEJKzW6rH4AfuN1VKuPLeKi4nBnWsFg2qFzf6X4RR4fG"
	expectedHash := sha256.Sum256(data)
	hash, err := UnHash(encodedHash)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(hash, expectedHash[:]) {
		log.Fatalf("expected '%x' and unhashed '%x' hashes do not match",
			expectedHash, hash)
	}
	if _, err = UnHash("_"); err == nil {
		t.Fatal("unhashed invalid hash")
	}
}

func TestBlocksInFile(t *testing.T) {
	contentSize := 1001 // kB / 500kB per block -> 3 blocks
	expected := 3
	actual := blocksInFile(contentSize)
	if actual != expected {
		t.Fatalf("expected '%d' got '%d'", expected, actual)
	}
}

func TestSprintTree(t *testing.T) {
	content := []byte{2, 1, 1, 3}
	x := File{name: "x", hash: Hash(content)}
	a := Dir{name: "a", subnodes: []tree.Node{x}}
	expected := string(readGolden(t, "testsprinttree.golden.tree"))
	actual := sprintTree(a, 0)
	if actual != expected {
		t.Fatalf("mismatched strings\nExpected:\n%s\nGot:\n%s",
			expected, actual)
	}
}

func TestMakeFile(t *testing.T) {
	// file containing two blocks
	n := int(float64(blockSize) * 1.5)
	fileContent := make([]byte, n)
	// bits of first block to 0, second block to 1
	for i := blockSize; i < n; i++ {
		fileContent[i] = 0xFF
	}

	// call makeFile
	f, err := makeFile("x", fileContent)
	if err != nil {
		t.Fatal(err)
	}
	expected := string(readGolden(t, "testmakefile.golden.tree"))
	actual := sprintTree(f, 0)
	if actual != expected {
		t.Fatalf("mismatched trees:\nactual:\n%s\nexpected:\n%s",
			actual, expected)
	}
}

func TestMakeDir(t *testing.T) {
	dirName := "testmakedir"
	subDirName := "subdir"
	fileName := "testfile"
	n := int(float64(blockSize) * 1.5)
	fileContent := make([]byte, n)

	os.Mkdir(filepath.Join(testDir, dirName), permDir)
	os.Mkdir(filepath.Join(testDir, dirName, subDirName), permDir)
	ioutil.WriteFile(filepath.Join(testDir, dirName, fileName),
		fileContent, permFile)

	d, err := makeDir(filepath.Join(testDir, dirName))
	if err != nil {
		t.Fatal(err)
	}
	expected := string(readGolden(t, "testmakedir.golden.tree"))
	actual := sprintTree(d, 0)

	if actual != expected {
		t.Fatalf("mismatched trees:\nactual:\n%s\nexpected:\n%s",
			actual, expected)
	}
}
