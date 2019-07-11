package fs

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	mhopts "github.com/multiformats/go-multihash/opts"
	"gitlab.com/mikelsr/sakaban/fs/tree"
)

// Hash creates a sha256 hash of data
func Hash(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}

// MultiHash creates a multihash using the sha256 algorithm,
// then encodes it in base58
func MultiHash(sha256Hash []byte) string {
	preHash := []byte{hashAlg, hashLen}
	hash := append(preHash, sha256Hash[:]...)
	mh, err := mhopts.Encode(hashEnc, hash)
	if err != nil {
		// TODO
	}
	return mh
}

// UnHash retuns the original hash given a base58-encoded multihash
func UnHash(hash string) ([]byte, error) {
	multihash, err := mhopts.Decode(hashEnc, hash)
	if err != nil {
		return nil, err
	}
	// extract algorithm and key lenght
	return multihash[2:], nil
}

// blockNumber calculates the number of block a file is split to
func blocksInFile(contentSize int) int {
	return int(math.Ceil(float64(contentSize) / float64(blockSize)))
}

// sprintTree is used to recursively print a tree in a readable format
func sprintTree(t tree.Node, tabLvl int) string {
	var str strings.Builder
	tab := strings.Repeat("\t", tabLvl)
	str.WriteString(
		fmt.Sprintf("%s%s: %s\n", tab, t.Name(), MultiHash(t.Hash())))

	for _, subnode := range t.Subnodes() {
		str.WriteString(sprintTree(subnode, tabLvl+1))
	}
	return str.String()
}

// MarshalTree creates a JSON from a tree.Node
func MarshalTree(t tree.Node) []byte {
	return tree.FromNode(t, MultiHash).JSON()
}

func makeFile(name string, content []byte) (File, error) {
	// split the file into blocks
	n := blocksInFile(len(content))
	blocks := make([]Block, n)
	for i := 0; i < n-1; i++ {
		blocks[i] = Block{
			index: strconv.Itoa(i),
			hash:  Hash(content[i*blockSize : (i+1)*blockSize]),
		}
	}
	// the size of the last block is variable
	blocks[n-1] = Block{
		index: strconv.Itoa(n - 1),
		hash:  Hash(content[(n-1)*blockSize:]),
	}

	// type []Block needs to be explicitely converted to type []tree.Node
	subnodes := make([]tree.Node, len(blocks))
	for i, block := range blocks {
		subnodes[i] = block
	}

	f := File{
		name:     name,
		hash:     Hash(content),
		subnodes: subnodes,
	}
	return f, nil
}

// makeTree creates a tree from a path to a directory
func makeDir(path string) (Dir, error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return Dir{}, err
	}
	// recursively create subnodes
	subnodes := make([]tree.Node, len(dir))
	for i, subnode := range dir {
		snPath := filepath.Join(path, subnode.Name())
		snInfo, err := os.Stat(snPath)
		if err != nil {
			return Dir{}, nil
		}
		var sn tree.Node
		if snInfo.IsDir() {
			sn, err = makeDir(snPath)
		} else {
			var content []byte
			content, err = ioutil.ReadFile(snPath)
			if err != nil {
				return Dir{}, err
			}
			sn, err = makeFile(subnode.Name(), content)
		}
		if err != nil {
			return Dir{}, err
		}
		subnodes[i] = sn
	}
	root := Dir{name: filepath.Base(path), subnodes: subnodes}
	return root, nil
}

// MakeTree creates the tree of a directory given a path to it
func MakeTree(path string) (Dir, error) {
	info, err := os.Stat(path)
	// check path
	if err != nil {
		return Dir{}, err
	}
	// verify that the path points to a directory
	if !info.IsDir() {
		return Dir{}, fmt.Errorf("'%s' is not a directory", path)
	}
	return makeDir(path)
}
