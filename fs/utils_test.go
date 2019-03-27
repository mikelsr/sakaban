package fs

import (
	"bytes"
	"crypto/sha256"
	"log"
	"testing"
)

func TestMultiHash(t *testing.T) {
	data := []byte("let's run some tests")
	expectedHash := "QmfEJKzW6rH4AfuN1VKuPLeKi4nBnWsFg2qFzf6X4RR4fG"
	sha256Hash := sha256.Sum256(data)
	hash := MultiHash(sha256Hash[: /*[32]byte to []byte*/])
	if hash != expectedHash {
		t.Fatalf("mimsmatched hashes '%s' and '%s' (expected))", hash, expectedHash)
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
}
