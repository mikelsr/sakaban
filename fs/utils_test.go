package fs

import (
	"bytes"
	"crypto/sha256"
	"log"
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
	expected := `a: QmYxH366TNoD7rwiSujCfrT648MoXhUtWf3wh9TrouWAZx
	x: QmWCMyGmbnwnEREqsvS924UuHVimBdPzTpCEdUdP8ecrkV` + "\n"
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

	fileName := "x"
	blockN := 2

	// call makeFile
	f, err := makeFile(fileName, fileContent)
	if err != nil {
		t.Fatal(err)
	}
	// check number of blocks
	if len(f.Subnodes()) != blockN {
		t.Fatalf("mismatched subnode len: expected '%d' got '%d'",
			blockN, len(f.Subnodes()))
	}

	block1 := f.Subnodes()[0]
	block2 := f.Subnodes()[1]

	// check names and hashes of file and blocks
	// calculate hashes of both bloks and of file
	if fileName != f.Name() {
		t.Fatalf("mismatched names: expected '%s' for '%s'",
			fileName, f.Name())
	}
	fileHash := Hash(fileContent)
	if !bytes.Equal(fileHash, f.Hash()) {
		t.Fatalf("mismatched file hashes: expected '%x' got '%x'",
			fileHash, f.Hash())
	}

	if "0" != block1.Name() {
		t.Fatalf("mismatched names: expected '%s' got '%s'",
			"0", block1.Name())
	}
	block1Hash := Hash(fileContent[:blockSize])
	if !bytes.Equal(block1Hash, block1.Hash()) {
		t.Fatalf("mismatched file hashes: expected '%x' got '%x'",
			block1Hash, block1.Hash())
	}

	if "1" != block2.Name() {
		t.Fatalf("mismatched names: expected '%s' got '%s'",
			"1", block2.Name())
	}
	block2Hash := Hash(fileContent[blockSize:])
	if !bytes.Equal(block2Hash, block2.Hash()) {
		t.Fatalf("mismatched file hashes: expected '%x' got '%x'",
			block2Hash, block2.Hash())
	}

}
