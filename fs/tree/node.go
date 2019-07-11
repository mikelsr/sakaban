package tree

import "encoding/json"

// Node represents a Node of the merkletree
// a directory is a branch node while a file is a leaf node
type Node interface {
	Name() string     // name of the folder or file
	Hash() []byte     // hash of the content of the node or file
	IsDir() bool      // wether the node is a directory
	IsFile() bool     // wether the node is a file node
	Subnodes() []Node // children nodes of a folder node
}

// SerializableNode is used to serialize/deserialize a node
type SerializableNode struct {
	Name     string             `json:"name"`
	Hash     string             `json:"hash"`
	IsDir    bool               `json:"isdir"`
	IsFile   bool               `json:"isfile"`
	Subnodes []SerializableNode `json:"subnodes,omitempty"`
}

// JSON converts the serializeable node into a Json
func (sn SerializableNode) JSON() []byte {
	content, _ := json.Marshal(sn)
	return content
}

// FromNode recursively creates a SerializableNode from a Node
// hashToStringFn: funcion that converts a hash to a string
func FromNode(node Node, hashToStringFn func(hash []byte) string) SerializableNode {
	sn := SerializableNode{}
	sn.Name = node.Name()
	sn.IsDir = node.IsDir()
	sn.Hash = hashToStringFn(node.Hash())

	sn.Subnodes = make([]SerializableNode, len(node.Subnodes()))
	for i, subnode := range node.Subnodes() {
		sn.Subnodes[i] = FromNode(subnode, hashToStringFn)
	}
	return sn
}
