package fs

import (
	"testing"
)

func TestHash(t *testing.T) {
	data := []byte("let's run some tests")
	expectedHash := "QmfEJKzW6rH4AfuN1VKuPLeKi4nBnWsFg2qFzf6X4RR4fG"
	hash, err := Hash(data)
	if err != nil {
		t.Fatal(err)
	}
	if hash != expectedHash {
		t.Fatalf("mimsmatched hashes '%s' and '%s' (expected))", hash, expectedHash)
	}
}
