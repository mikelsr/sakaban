package merkletree

// Node represents a Node of a MerkleTree, be it a normal node or a leaf node
type Node interface {
	Hash() []byte
	Name() string
	Subnodes() []Node
}

// Path to a file, the first node is the root node and the last one the file node
type Path []Node

// Hash returns the hash of the path, recursively
func (p Path) Hash() []byte {
	return p.Last().Hash()
}

// Last returns the last node of a path
func (p Path) Last() Node {
	return p[len(p)-1]
}

// Name returns the name of the path, recursively
func (p Path) Name() string {
	return p.Last().Name()
}

// Subnodes returns the subnodes of the final node of the path
func (p Path) Subnodes() []Node {
	return p.Last().Subnodes()
}
