package fs

import (
	"crypto/sha256"

	mhopts "github.com/multiformats/go-multihash/opts"
)

// Hash creates a multihash using the sha256 algorithm,
// then encodes it in base58
func Hash(b []byte) (string, error) {
	shaHash := sha256.Sum256(b)
	preHash := []byte{hashAlg, hashLen}
	hash := append(preHash, shaHash[:]...)
	return mhopts.Encode("base58", hash)
}
