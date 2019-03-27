package fs

import (
	"bytes"
	"crypto/sha256"
	"io/ioutil"
	"testing"
)

func TestFile_PreHash(t *testing.T) {
	content, err := ioutil.ReadFile(muffinPath)
	if err != nil {
		t.Fatal(err)
	}
	f := File{name: muffinName, content: content}
	expectedHash := sha256.Sum256(content)
	preHash := f.PreHash()
	if !bytes.Equal(preHash, expectedHash[:]) {
		t.Fatalf("prehash doesn't match expected prehash: %x vs %x",
			preHash, expectedHash)
	}
}

func TestFile_Hash(t *testing.T) {
	content, err := ioutil.ReadFile(muffinPath)
	if err != nil {
		t.Fatal(err)
	}
	f := File{name: muffinName, content: content}
	expectedHash := sha256.Sum256(append(content, []byte(muffinName)...))
	hash := f.Hash()
	if !bytes.Equal(hash, expectedHash[:]) {
		t.Fatalf("hash doesn't match expected hash: %x vs %x",
			hash, expectedHash)
	}
}

func TestFile_IsDir(t *testing.T) {
	if new(File).IsDir() {
		t.Fatal("File.IsDir() returned true")
	}
}

func TestFile_Subnodes(t *testing.T) {
	if new(File).Subnodes() != nil {
		t.Fatal("File.Subnodes() didn't return nil")
	}
}
