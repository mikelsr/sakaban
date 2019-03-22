package tree

import (
	"path/filepath"
)

// Branch is an implementation of node
// the first Node of the slide is the root tree, the last one the current
// folder/directory
type Branch []Node

func (b Branch) last() Node {
	return b[len(b)-1]
}

// Name of the last node of the branch
func (b Branch) Name() string {
	return b.last().Name()
}

// PreHash of the last node of the branch
func (b Branch) PreHash() []byte {
	return b.last().PreHash()
}

// Hash of the last node of the branch
func (b Branch) Hash() []byte {
	return b.last().Hash()
}

// IsDir returns true if the last node of the branch is a directory node
func (b Branch) IsDir() bool {
	return b.last().IsDir()
}

// Subnodes of the last node of the branch
func (b Branch) Subnodes() []Node {
	return b.last().Subnodes()
}

// Path of the last node of the branch
func (b Branch) Path() string {
	filenames := make([]string, len(b))
	for i, node := range b {
		filenames[i] = node.Name()
	}
	return filepath.Join(filenames...)
}
