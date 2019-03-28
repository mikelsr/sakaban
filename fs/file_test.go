package fs

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"

	"reflect"
)

// TestMain will create and delete the testing directory
// before and after running tests, respectively
func TestMain(m *testing.M) {
	rand.Seed(time.Now().UTC().UnixNano())
	os.MkdirAll(testDir, permissionDir)
	os.Mkdir(testFailDir, 000)
	m.Run()
	// cleanup
	os.RemoveAll(testDir)
}

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

func Test_MakeFile(t *testing.T) {
	content, err := ioutil.ReadFile(muffinPath)
	if err != nil {
		t.Fatal(err)
	}
	expectedFile := File{name: muffinName, content: content, prehash: Hash(content)}
	if _, err := MakeFile(""); err == nil {
		t.Fatal("err is nil")
	}
	file, err := MakeFile(muffinPath)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(expectedFile, *file) {
		t.Fatalf("mismatched files")
	}
}
