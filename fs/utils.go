package fs

import (
	mhopts "github.com/multiformats/go-multihash/opts"
)

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
