package fs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// File represents a file in a Unix system.
//	Hash:	base58-encoded multihash of the block
//	Path:	path of the file
//	Perm:	permissions of the file
//	Blocks:	blocks that form the file
type File struct {
	Blocks []*Block    `json:"blocks"`
	Hash   string      `json:"hash"`
	Path   string      `json:"path"`
	Perm   os.FileMode `json:"permissions"`
}

// NewFileFromPath creates a File given a path to a file.
func NewFileFromPath(path string) (*File, error) {
	// attribute: path
	if !isFile(path) {
		return nil, fmt.Errorf("'%s' is not a valid file", path)
	}
	// attribute: permissions
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	// attribute: hash
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	hash, err := Hash(content)
	if err != nil {
		return nil, err
	}

	// attr: blocks
	n := BlockAmount(info.Size(), int64(BlockSize))
	blocks := make([]*Block, n)
	for i := 0; i < n; i++ {
		var blockContent []byte
		if i < n-1 {
			blockContent = content[i*BlockSize : (i+1)*BlockSize]
		} else { // last block may be smaller
			blockContent = content[i*BlockSize:]
		}
		block, err := NewBlockFromBytes(blockContent)
		if err != nil {
			return nil, err
		}
		blocks[i] = block
	}

	f := File{Blocks: blocks, Hash: hash, Path: path, Perm: info.Mode()}
	return &f, nil
}

// Name returns the filename of the file
func (f *File) Name() string {
	return filepath.Base(f.Path)
}

func (f *File) String() string {
	s, _ := json.Marshal(f)
	return string(s)
}

// Write creates a file at f.Path with the content of f.Blocks
func (f *File) Write() error {
	fi, err := os.OpenFile(f.Path, os.O_CREATE|os.O_WRONLY, f.Perm)
	if err != nil {
		return err
	}
	for _, b := range f.Blocks {
		fi.Write(b.Content)
	}
	return nil
}
