package fs

import (
	"gitlab.com/mikelsr/sakaban/fs/tree"
)

// Block is the implementation of the block node
type Block struct {
	index   string // index of the block in the file
	content []byte // content of the block
	hash    []byte // precalculated hash used for optimization
}

// Name returns the index of the block
func (b Block) Name() string {
	return b.index
}

// Hash of the content of the block
func (b Block) Hash() []byte {
	if b.hash == nil {
		// TODO
	}
	return b.hash
}

// IsDir returns false
func (b Block) IsDir() bool {
	return false
}

// IsFile always returns false
func (b Block) IsFile() bool {
	return false
}

// Subnodes returns nil
func (b Block) Subnodes() []tree.Node {
	return nil
}
