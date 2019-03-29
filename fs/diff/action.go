package diff

// Action represents a type of change made to a tree.Node
type Action int

const (
	_ Action = iota
	// Create a new file
	Create
	// Delete a file
	Delete
	// Modify the content of a file
	Modify
)
