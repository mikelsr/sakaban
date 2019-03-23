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
	name    string
	prehash []byte
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
func (t treeNode) PreHash() []byte {
	hashBytes := []byte{}
	for _, subnode := range t.Subnodes() {
		hashBytes = append(hashBytes, subnode.Hash()...)
	}
	return hashFn(hashBytes)
}

// PreHash returns the hash of the content of the leaf
func (l leafNode) PreHash() []byte {
	return l.prehash
}

// Hash of the content and name of the tree
func (t treeNode) Hash() []byte {
	return hashFn(append(t.PreHash(), []byte(t.name)...))
}

// Hash of the content and name of the leaf
func (l leafNode) Hash() []byte {
	return hashFn(append(l.prehash, []byte(l.name)...))
}

// IsDir returns true
func (t treeNode) IsDir() bool {
	return true
}

// IsDir returns false
func (l leafNode) IsDir() bool {
	return false
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
		name:    "a1.raw",
		prehash: hashFn(leafContent1),
	}
	a2 := leafNode{
		name:    "a2.raw",
		prehash: hashFn(leafContent1),
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
			sn.PreHash, string(hashToString(root.PreHash())))
	}

	if sn.Subnodes[0].Subnodes[0].Hash != hashToString(a1.Hash()) {
		t.Fatalf("Serialized and original hashes do not match %s %s",
			sn.Subnodes[0].Subnodes[0].Hash, hashToString(a1.Hash()))
	}

	// test SerializableNode.JSON()
	marshalled := sn.JSON()
	testSerializableNodeJSON(marshalled, t)
}
