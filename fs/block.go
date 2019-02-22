package fs

// Block contains some bytes of a file.
// 	Content: bytes of the file
type Block struct {
	Content []byte
}

// Equals compares the hashes of two blocks
func (b *Block) Equals(b1 *Block) bool {
	hash1, err := Hash(b.Content)
	if err != nil {
		return false
	}
	hash2, err := Hash(b1.Content)
	if err != nil || hash1 != hash2 {
		return false
	}
	return true
}

// Size returns the lenght of the block in bytes
func (b *Block) Size() int {
	return len(b.Content)
}
