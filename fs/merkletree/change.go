package merkletree

// Change is a change made to a file
type Change struct {
	// From is the origin file, nil if created
	From Path
	// To is file under change, nil if deleted
	To Path
}

// Action returns the type of change
func (c Change) Action() Action {
	if c.From == nil {
		return Create
	}
	if c.To == nil {
		return Delete
	}
	return Modify
}
