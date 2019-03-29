package fs

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"bitbucket.org/mikelsr/sakaban/fs/tree"
)

// Dir represents a dir node implemented as a tree.Node
type Dir struct {
	name     string
	subnodes []tree.Node
}

// Name returns the name of the directory
func (d Dir) Name() string {
	return d.name
}

// PreHash returns the combined hash of all the subnodes of the directory
func (d Dir) PreHash() []byte {
	raw := []byte{}
	for _, subnode := range d.Subnodes() {
		raw = append(raw, subnode.Hash()...)
	}
	return Hash(raw)
}

// Hash returns the hash of combination of prehash and name of the directory
func (d Dir) Hash() []byte {
	return Hash(append(d.PreHash(), []byte(d.Name())...))
}

// IsDir always returns true
func (d Dir) IsDir() bool {
	return true
}

// Subnodes returns the subdirs and files of the directory
func (d Dir) Subnodes() []tree.Node {
	return d.subnodes
}

// Sort alphabetically sorts the subnodes of d
func (d *Dir) Sort() {
	sort.Slice(d.subnodes, func(i, j int) bool {
		return d.subnodes[i].Name() < d.subnodes[j].Name()
	})
}

// MakeTree recursively creates a tree given a valid path
func MakeTree(path string) (tree.Node, error) {
	// get information about current path
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	// create file node
	if !info.IsDir() {
		return MakeFile(path)
	}

	// scan directory
	dir, _ := ioutil.ReadDir(path)
	// recursively create nodes
	subnodes := make([]tree.Node, len(dir))
	for i, f := range dir {
		t, err := MakeTree(filepath.Join(path, f.Name()))
		if err != nil {
			return nil, err
		}
		subnodes[i] = t
	}
	t := Dir{name: info.Name(), subnodes: subnodes}
	t.Sort()
	return &t, nil
}
