package merkletree

import "path/filepath"

// Path to a file, the first node is the root node and the last one the file node
type Path []Node

// Hash returns the hash of the path, recursively
func (p Path) Hash() []byte {
	return p.Last().Hash()
}

// IsDir returns true if the last node is a directory
func (p Path) IsDir() bool {
	return p.Last().IsDir()
}

// Last returns the last node of a path
func (p Path) Last() Node {
	return p[len(p)-1]
}

// Name returns the name of the path, recursively
func (p Path) Name() string {
	return p.Last().Name()
}

// Path retuns a string with the path to a node
func (p Path) Path() string {
	names := make([]string, len(p))
	for i, node := range p {
		names[i] = node.Name()
	}
	return filepath.Join(names...)
}

// Subnodes returns the subnodes of the final node of the path
func (p Path) Subnodes() []Node {
	return p.Last().Subnodes()
}
