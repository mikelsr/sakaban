package fs

import (
	"gitlab.com/mikelsr/sakaban/fs/tree"
	"gitlab.com/mikelsr/sakaban/hash"
)

// Dir is the implementation of the directory node
type Dir struct {
	name     string      // name of the directory
	subnodes []tree.Node // files and subdirectories
}

// Name of the directory
func (d Dir) Name() string {
	return d.name
}

// Hash of the hashes of the subnodes of the directory
func (d Dir) Hash() []byte {
	raw := []byte{}
	for _, subnode := range d.Subnodes() {
		raw = append(raw, subnode.Hash()...)
	}
	return hash.Hash(raw)
}

// IsDir returns true
func (d Dir) IsDir() bool {
	return true
}

// IsFile always returns false
func (d Dir) IsFile() bool {
	return false
}

// Subnodes returns a list of file and dir nodes
func (d Dir) Subnodes() []tree.Node {
	return d.subnodes
}
