package fs

import (
	"bitbucket.org/mikelsr/sakaban/fs/tree"
)

// Dir represents a dir node implemented as a tree.Node
type Dir struct {
	name     string
	subnodes []tree.Node
}

// Name returns the name of the directory
func (d Dir) Name() string {
	return d.name
}

// PreHash returns the combined hash of all the subnodes of the directory
func (d Dir) PreHash() []byte {
	raw := []byte{}
	for _, subnode := range d.Subnodes() {
		raw = append(raw, subnode.Hash()...)
	}
	return Hash(raw)
}

// Hash returns the hash of combination of prehash and name of the directory
func (d Dir) Hash() []byte {
	return Hash(append(d.PreHash(), []byte(d.Name())...))
}

// IsDir always returns true
func (d Dir) IsDir() bool {
	return true
}

// Subnodes returns the subdirs and files of the directory
func (d Dir) Subnodes() []tree.Node {
	return d.subnodes
}
