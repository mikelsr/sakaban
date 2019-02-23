package merkletree

// Action is a type of change made to a file
type Action int

const (
	_ Action = iota
	// Create the file
	Create
	// Delete the file
	Delete
	// Modify the content of the file
	Modify
)
