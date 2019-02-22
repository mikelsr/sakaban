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

// Hash creates a multihash using the sha256 algorithm,
// then encodes it in base58.
func Hash(b []byte) (string, error) {
	hasher := sha256.New()
	hasher.Write(b)
	shaHash := hasher.Sum(nil)
	preHash := []byte{hashAlg, hashLen}
	hash := append(preHash, shaHash[:]...)
	return mhopts.Encode("base58", hash)
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
