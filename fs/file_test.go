package fs

import (
	"encoding/json"
	"path/filepath"
	"testing"
)

func TestNewFileFromPath(t *testing.T) {
	// valid file
	f, err := NewFileFromPath(muffinPath)
	if err != nil {
		t.Fatal(err)
	}
	if f.Hash != muffinHash {
		t.Fatalf("mismatched muffin hashes '%s' and '%s' (expected)",
			f.Hash, muffinHash)
	}
	for i, b := range f.Blocks {
		if b.Hash != muffinBlockHashes[i] {
			t.Fatalf("mismatched muffin block hashes '%s' and '%s' (expected)",
				b.Hash, muffinBlockHashes[i])
		}
	}

	// directory
	if _, err = NewFileFromPath(testFailDir); err == nil {
		t.Fatalf("created valid file from directory")
	}

	// unreadable file
	if _, err = NewFileFromPath(testFailFile); err == nil {
		t.Fatalf("created valid file from unreadable file")
	}
}

func TestFile_Name(t *testing.T) {
	f := File{Path: muffinPath}
	expected := filepath.Base(muffinPath)
	if f.Name() != expected {
		t.Fatalf("mismatched file names '%s' and '%s' (expected)",
			f.Name(), expected)
	}
}

func TestFile_String(t *testing.T) {
	f, _ := NewFileFromPath(muffinPath)
	if err := json.Unmarshal([]byte(f.String()), f); err != nil {
		t.FailNow()
	}
}

func TestFile_Write(t *testing.T) {
	// valid write
	f, _ := NewFileFromPath(muffinPath)
	f.Path = testFile
	if err := f.Write(); err != nil {
		t.Fatalf("could write file to '%s'", testFile)
	}
	if !isFile(testFile) {
		t.Fatalf("written file isn't a file '%s'", testFile)
	}
	f1, err := NewFileFromPath(testFile)
	if err != nil {
		t.Fatal(err)
	}
	if f.Hash != f1.Hash {
		t.Fatalf("original and written file do not match")
	}

	// invalid write
	f.Path = testFailFile
	if err = f.Write(); err == nil {
		t.Fatalf("successfully wrote file to protected directory")
	}
}
