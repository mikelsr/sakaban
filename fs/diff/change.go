package diff

import "bitbucket.org/mikelsr/sakaban/fs/tree"

// Change represents
type Change struct {
	From *tree.Branch
	To   *tree.Branch
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
