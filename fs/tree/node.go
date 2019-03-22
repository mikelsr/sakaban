package tree

// Node represents a Node of the merkletree
// a directory is a branch node while a file is a leaf node
type Node interface {
	Name() string     // name of the folder or file
	PreHash() []byte  // hash of the content of the folder or file
	Hash() []byte     // combined hash of the name and prehash of the node
	IsDir() bool      // wether the node is a directory
	Subnodes() []Node // children nodes of a folder node
}

// SerializableNode is used to serialize/deserialize a node
type SerializableNode struct {
	Name     string             `json:"name"`
	PreHash  string             `json:"prehash"`
	Hash     string             `json:"hash"`
	IsDir    string             `json:"isdir"`
	Subnodes []SerializableNode `json:"subnodes"`
}

// FromNode creates a SerializableNode from a Node
func FromNode() SerializableNode {
	// TODO
}

// ToNode creates a Node from a SerializableNode
func ToNode() Node {
	// TODO
}
