package tree

// Dir represents a Directory node
type Dir struct {
	name    string
	prehash []byte
	// TODO: remove hash attribute, obtain in Hash() method
	hash     []byte
	subnodes []Node
}

// Name of the directory
func (d Dir) Name() string {
	return d.name
}

// PreHash returns the hash of the content of the directory
func (d Dir) PreHash() []byte {
	return d.prehash
}

// Hash of the content and name of the directory
func (d Dir) Hash() []byte {
	return d.hash
}

// IsDir returns true
func (d Dir) IsDir() bool {
	return false
}

// Subnodes returns the children of the folder node
func (d Dir) Subnodes() []Node {
	return d.subnodes
}
