package fs

const (
	//BlockSize is the size of every block (excluding final ones)
	BlockSize int = 1e6

	hashAlg = 0x12     // Algorithm, Multihash SHA2_256
	hashEnc = "base58" // Encoding, base58
	hashLen = 0x20     // Lenght, 32 bytes
)
