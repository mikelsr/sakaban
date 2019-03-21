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
	hash, err := MultiHash(data)
	if err != nil {
		t.Fatal(err)
	}
	if hash != expectedHash {
		t.Fatalf("mimsmatched hashes '%s' and '%s' (expected))", hash, expectedHash)
	}
}

func TestSimpleHash(t *testing.T) {
	data := []byte("let's run some tests")
	expectedHash := sha256.Sum256(data)
	hash := SimpleHash(data)
	if !bytes.Equal(hash, expectedHash[:]) {
		log.Fatalf("expected '%x' and calculated '%x' hashes do not match",
			expectedHash, hash)
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
