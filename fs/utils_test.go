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

func TestMarshalTree(t *testing.T) {
	d := Dir{
		name: "x",
		subnodes: []tree.Node{
			File{
				name: "y",
				subnodes: []tree.Node{
					Block{
						index: "0",
						hash:  Hash([]byte{0x00}),
					},
					Block{
						index: "1",
						hash:  Hash([]byte{0xFF}),
					},
				},
			},
		},
	}
	expected := readGolden(t, "testmarshaltree.golden.json")
	actual := MarshalTree(d)
	if !bytes.Equal(actual, expected) {
		t.Fatal("mismatched marshalled trees")
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
	expected := readGolden(t, "testmakefile.golden.json")
	actual := tree.FromNode(f, MultiHash).JSON()
	if !bytes.Equal(actual, expected) {
		t.Fatalf("mismatched trees:\nactual:\n%s\nexpected:\n%s",
			actual, expected)
	}
}

func testMakeDir(t *testing.T) (string, []byte) {
	// invalid directories
	// missing
	if _, err := makeDir("/ _"); err == nil {
		t.Fatalf("created Dir from invalid direcroty '/ _'")
	}
	// no perm
	if _, err := makeDir(testFailDir); err == nil {
		t.Fatalf("created Dir from non-readable directory")
	}
	// testDir contains the testFailDir subdirectory
	if _, err := makeDir(testDir); err == nil {
		t.FailNow()
	}

	// create and test valid directories
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
	expected := readGolden(t, "testmakedir.golden.json")
	actual := tree.FromNode(d, MultiHash).JSON()

	if !bytes.Equal(actual, expected) {
		t.Fatalf("mismatched trees:\nactual:\n%s\nexpected:\n%s",
			actual, expected)
		return "", []byte{}
	}
	return filepath.Join(testDir, dirName), actual
}

func TestMakeTree(t *testing.T) {
	if _, err := MakeTree("/ _"); err == nil {
		t.Fatalf("created tree from missing dir")
	}

	os.Mkdir(filepath.Join(testDir, "testmaketree"), permDir)
	ioutil.WriteFile(filepath.Join(testDir, "testmaketree", "x"),
		[]byte{}, permFile)
	if _, err := MakeTree(filepath.Join(testDir, "testmaketree", "x")); err == nil {
		t.Fatalf("created tree from file")
	}

	path, expected := testMakeDir(t)
	d, err := MakeTree(path)
	if err != nil {
		t.Fatal(err)
	}
	actual := tree.FromNode(d, MultiHash).JSON()
	if !bytes.Equal(actual, expected) {
		t.Fatalf("mismatched trees:\nactual:\n%s\nexpected:\n%s",
			actual, expected)
	}
}
