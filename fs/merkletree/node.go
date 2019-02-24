package merkletree

// Node represents a Node of a MerkleTree, be it a normal node or a leaf node
type Node interface {
	Hash() []byte
	IsDir() bool
	Name() string
	Subnodes() []Node
}
