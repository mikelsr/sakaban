package fs

// Block contains some bytes of a file.
//	Hash: base58-encoded multihash of the block
// 	Content: bytes of the file
type Block struct {
	Content []byte `json:"-"`
	Hash    string `json:"hash"`
}

// Equals compares the hashes of two blocks.
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

// NewBlockFromBytes creates a Block given it's content in bytes
func NewBlockFromBytes(content []byte) (*Block, error) {
	hash, err := Hash(content)
	if err != nil {
		return nil, err
	}
	return &Block{Content: content, Hash: hash}, nil
}

// Size returns the lenght of the block in bytes.
func (b *Block) Size() int {
	return len(b.Content)
}
