package tree

// File represents a File node
type File struct {
	name    string
	prehash []byte
	// TODO: remove hash attribute, obtain in Hash() method
	hash []byte
}

// Name of the file
func (f File) Name() string {
	return f.name
}

// PreHash returns the hash of the content of the file
func (f File) PreHash() []byte {
	return f.prehash
}

// Hash of the content and name of the file
func (f File) Hash() []byte {
	return f.hash
}

// IsDir returns false
func (f File) IsDir() bool {
	return false
}

// Subnodes returns a nil pointer
func (f File) Subnodes() []Node {
	return nil
}
