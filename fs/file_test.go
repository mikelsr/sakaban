package fs

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestFile_PreHash(t *testing.T) {
	content, err := ioutil.ReadFile(muffinPath)
	if err != nil {
		t.Fatal(err)
	}
	f := File{name: muffinName, content: content}
	expectedHash := Hash(content)
	_preHash := Hash(f.content)
	f.prehash = _preHash
	preHash := f.PreHash()
	if !bytes.Equal(preHash, expectedHash) {
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
	contentHash := Hash(content)
	f.prehash = contentHash
	expectedHash := Hash(append(contentHash, []byte(muffinName)...))
	hash := f.Hash()
	if !bytes.Equal(hash, expectedHash) {
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
