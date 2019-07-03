package fs

import (
	"crypto/sha256"
	"fmt"
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

// sprintTree is used to recursively print a tree
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
