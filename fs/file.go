package fs

import "crypto/sha256"
import "bitbucket.org/mikelsr/sakaban/fs/tree"

// File represents a file as an implementation of tree.Node
type File struct {
	content []byte // content of the file
	prehash []byte // prehash of the file
	name    string
}

// Name returns the basename of the file
func (f File) Name() string {
	return f.name
}

// PreHash returns the hash of the content of the file
func (f File) PreHash() []byte {
	return f.prehash
}

// Hash returns the hash of combination of content and name of the file
func (f File) Hash() []byte {
	hasher := sha256.New()
	hasher.Write(f.PreHash())
	hasher.Write([]byte(f.Name()))
	return hasher.Sum(nil)
}

// IsDir retuns false
func (f File) IsDir() bool {
	return false
}

// Subnodes returns a nil pointer
func (f File) Subnodes() []tree.Node {
	return nil
}
