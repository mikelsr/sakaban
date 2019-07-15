package snap

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gitlab.com/mikelsr/sakaban/hash"
)

// Blob maps a hash to the binary content used to create the hash
type Blob struct {
	hash    []byte // hash of the blob
	content []byte // content of the blob
}

// Write the content of a Blob to ${path}/Blob.hash
func (b Blob) Write(dpath string) error {
	fp := filepath.Join(dpath, hash.MultiHash(b.hash))
	return ioutil.WriteFile(fp, b.content, permFile)
}

// ReadBlob creates a Blob struct from one stored in ${fpath}
func ReadBlob(fpath string) (Blob, error) {
	mh := filepath.Base(fpath)
	content, err := ioutil.ReadFile(fpath)
	if err != nil {
		return Blob{}, err
	}
	// compare hash in path to actual hash of the content
	mhContent := hash.MultiHash(hash.Hash(content))
	if mh != mhContent {
		return Blob{}, fmt.Errorf("hash provided in path (%s) doesn't match hash calculated from content (%s)",
			mh, mhContent)
	}

	h, err := hash.UnHash(mh)
	if err != nil {
		return Blob{}, err
	}

	return Blob{hash: h, content: content}, nil
}
