package tree

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"testing"
)

// hashFn is called to create test hashes
func hashFn(content []byte) []byte {
	hasher := sha256.New()
	hasher.Write(content)
	return hasher.Sum(nil)
}

// hashToString converts test hashes to string
func hashToString(hash []byte) string {
	return base64.StdEncoding.EncodeToString(hash)
}

// treeNode represents a test tree node
type treeNode struct {
	name     string
	subnodes []Node
}

// leafNode represents a test leaf node
type leafNode struct {
	name     string
	hash     []byte
	subnodes []Node
}

// Name of the tree
func (t treeNode) Name() string {
	return t.name
}

// Name of the leaf
func (l leafNode) Name() string {
	return l.name
}

// PreHash returns the hash of the content of the tree
func (t treeNode) Hash() []byte {
	hashBytes := []byte{}
	for _, subnode := range t.Subnodes() {
		hashBytes = append(hashBytes, subnode.Hash()...)
	}
	return hashFn(hashBytes)
}

// Hash returns the hash of the content of the leaf
func (l leafNode) Hash() []byte {
	return l.hash
}

// IsDir returns true
func (t treeNode) IsDir() bool {
	return true
}

// IsDir returns false
func (l leafNode) IsDir() bool {
	return false
}

// IsFile returns false
func (t treeNode) IsFile() bool {
	return false
}

// IsFile returns false
func (l leafNode) IsFile() bool {
	return true
}

// Subnodes returns the subnodes of the tree
func (t treeNode) Subnodes() []Node {
	return t.subnodes
}

// Subnodes returns a nil pointer
func (l leafNode) Subnodes() []Node {
	return nil
}

func testSerializableNodeJSON(marshalledNode []byte, t *testing.T) {
	sn := new(SerializableNode)
	if err := json.Unmarshal(marshalledNode, sn); err != nil {
		t.Fatal(err)
	}
}

func TestFromNode(t *testing.T) {
	// "/"
	// |-"a/"
	// |  | - "a1.raw"
	// |  | - "a2.raw"
	// |-"b/"
	leafContent1 := make([]byte, 1024)
	leafContent2 := make([]byte, 1024)
	rand.Read(leafContent1)
	rand.Read(leafContent2)

	a1 := leafNode{
		name: "a1.raw",
		hash: hashFn(leafContent1),
	}
	a2 := leafNode{
		name: "a2.raw",
		hash: hashFn(leafContent1),
	}
	a := treeNode{
		name:     "a",
		subnodes: []Node{a1, a2},
	}
	b := treeNode{
		name:     "b",
		subnodes: []Node{},
	}
	root := treeNode{
		name:     "",
		subnodes: []Node{a, b},
	}

	sn := FromNode(root, hashToString)
	if sn.Hash != hashToString(root.Hash()) {
		t.Fatalf("Serialized and original hashes do not match %s %s",
			sn.Hash, string(hashToString(root.Hash())))
	}

	// test SerializableNode.JSON()
	marshalled := sn.JSON()
	testSerializableNodeJSON(marshalled, t)
}
