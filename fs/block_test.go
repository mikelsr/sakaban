package fs

import (
	"crypto/rand"
	"testing"
)

func TestBlock_Equals(t *testing.T) {
	content := make([]byte, 20)
	content1 := make([]byte, 20)
	rand.Read(content)
	rand.Read(content)

	b := &Block{Content: content}
	b1 := &Block{Content: content1}

	if b.Equals(b1) {
		t.Fatal("blocks shouldn't be equal")
	}
	b.Content = b1.Content
	if !b.Equals(b1) {
		t.Fatal("blocks should be equal")
	}
}

func TestBlock_Size(t *testing.T) {
	b := Block{Content: []byte{0, 1}}
	if b.Size() != 2 {
		t.Fatal("incorrect block size")
	}
}
