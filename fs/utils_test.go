package fs

import (
	"bytes"
	"crypto/sha256"
	"log"
	"testing"

	"bitbucket.org/mikelsr/sakaban/fs/tree"
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

func TestSprintTree(t *testing.T) {
	content := []byte{2, 1, 1, 3}
	x := File{name: "x", content: content, prehash: Hash(content)}
	a := Dir{name: "a", subnodes: []tree.Node{x}}
	expected := `a: QmbfzSMMYoNCZFrNnRWxqpepoJw31GLgjtkm55nKA2nA7a
	x: QmdXZVajHMALk3Y6xQ4tzYivFtTnHPnNZnHbVBLQL1yxeh` + "\n"
	actual := SprintTree(a, 0)
	if actual != expected {
		t.Fatalf("mismatched strings")
	}
}
