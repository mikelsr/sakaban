package snap

import (
	"bytes"
	"crypto/rand"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"gitlab.com/mikelsr/sakaban/hash"
)

func TestBlobWrite(t *testing.T) {
	// create dir to test on
	d := filepath.Join(testDir, "testblobwrite")
	os.Mkdir(d, testPermDir)
	// create random content
	content := make([]byte, 8)
	rand.Read(content)
	// calculate hash of generated content
	expectedHash := hash.Hash(content)
	blob := Blob{hash: expectedHash, content: content}

	// write hash
	blob.Write(d)

	// check name of file and content
	expectedMH := hash.MultiHash(expectedHash)
	// check hash by opening the file, which is at ${d}/${expectedMH}
	actualContent, err := ioutil.ReadFile(filepath.Join(d, expectedMH))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(content, actualContent) {
		t.Fatalf("mismatched contents: actual (%x) vs expected (%x)",
			actualContent, content)
	}

	// write hash to invalid directory
	if err := blob.Write(testFailDir); err == nil {
		t.Fatalf("wrote blob to invalid directory: '%s'", testFailDir)
	}
}

func TestReadBlob(t *testing.T) {
	// create dir to test on
	d := filepath.Join(testDir, "testreadblob")
	os.Mkdir(d, testPermDir)
	// create random content
	content := make([]byte, 8)
	rand.Read(content)
	// create and write original blob
	expectedBlob := Blob{hash: hash.Hash(content), content: content}
	expectedBlob.Write(d)
	// read blob
	actualBlob, err := ReadBlob(filepath.Join(d, hash.MultiHash(expectedBlob.hash)))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(actualBlob.hash, expectedBlob.hash) {
		t.Fatalf("mismatched hashes: expected (%x) vs actual (%s)",
			expectedBlob.hash, actualBlob.hash)
	}
	if !bytes.Equal(actualBlob.content, expectedBlob.content) {
		t.Fatalf("mismatched contents: expected (%x) vs actual (%s)",
			expectedBlob.content, actualBlob.content)
	}

	// read from directory, not file
	if _, err = ReadBlob(testDir); err == nil {
		t.Fatalf("read file blob from directory: '%s'", testDir)
	}
	// read from missing file
	failPath := filepath.Join(testFailDir, hash.MultiHash(expectedBlob.hash))
	if _, err = ReadBlob(failPath); err == nil {
		t.Fatalf("read file blob from directory: '%s'", testDir)
	}
	// read blob with mismatched hash
	failHash := expectedBlob.hash
	// increase hash by 1 in the MSb
	failHash[0]++
	Blob{hash: failHash, content: content}.Write(d)
	if _, err := ReadBlob(filepath.Join(d, hash.MultiHash(failHash))); err == nil {
		t.Fatalf("ignored mismatched hash")
	}
}
