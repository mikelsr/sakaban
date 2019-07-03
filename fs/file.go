package fs

import (
	"gitlab.com/mikelsr/sakaban/fs/tree"
)

// File is the implementation of the file node
type File struct {
	name     string      // name of the directory
	hash     []byte      // precalculated hash of the file
	subnodes []tree.Node // content blocks of the file
}

// Name of the file
func (f File) Name() string {
	return f.name
}

// Hash of the content of the file
func (f File) Hash() []byte {
	return f.hash
}

// IsDir returns false
func (f File) IsDir() bool {
	return false
}

// IsFile always returns true
func (f File) IsFile() bool {
	return true
}

// Subnodes returns the list of blocks of the file
func (f File) Subnodes() []tree.Node {
	return f.subnodes
}
