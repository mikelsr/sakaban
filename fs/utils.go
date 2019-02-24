package fs

import (
	"crypto/sha256"
	"os"

	mhopts "github.com/multiformats/go-multihash/opts"
)

// BlockAmount returns the number of blocks corresponding to a file
// given the file and block sizes
func BlockAmount(fileSize int64, blockSize int64) int {
	n := int(fileSize / blockSize)
	if fileSize%blockSize == 0 {
		return n
	}
	return n + 1
}

// isFile checks if a path exists and contains a file, not a directory.
func isFile(path string) bool {
	file, err := os.Stat(path)
	// Check that file exists
	if os.IsNotExist(err) {
		return false
	}
	// Check that file is not a directory
	defer func() {
		// file.Mode().IsRegular() panicked
		if r := recover(); r != nil {
		}
	}()
	return file.Mode().IsRegular()
}

// MultiHash creates a multihash using the sha256 algorithm,
// then encodes it in base58.
func MultiHash(b []byte) (string, error) {
	shaHash := SimpleHash(b)
	preHash := []byte{hashAlg, hashLen}
	hash := append(preHash, shaHash[:]...)
	return mhopts.Encode(hashEnc, hash)
}

// SimpleHash creates a sha256 hash of b
func SimpleHash(b []byte) []byte {
	hasher := sha256.New()
	hasher.Write(b)
	return hasher.Sum(nil)
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
