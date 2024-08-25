package earthbucks

import (
	"github.com/zeebo/blake3"
)

// HashFunction defines a function that hashes input data.
type HashFunction func(input []byte) (*FixedBuf, error)

// MacFunction defines a function that performs a keyed hash.
type MacFunction func(key *FixedBuf, data []byte) (*FixedBuf, error)

var (
	Blake3Hash         HashFunction
	DoubleBlake3Hash  HashFunction
	Blake3Mac          MacFunction
)

func init() {
	Blake3Hash = func(data []byte) (*FixedBuf, error) {
		hash := blake3.New()
		hash.Write(data)
		return NewFixedBuf(32, hash.Sum(nil))
	}

	DoubleBlake3Hash = func(data []byte) (*FixedBuf, error) {
		hash1, err := Blake3Hash(data)
		if err != nil {
			return nil, err
		}
		return Blake3Hash(hash1.buf)
	}

	Blake3Mac = func(key *FixedBuf, data []byte) (*FixedBuf, error) {
		hasher,err := blake3.NewKeyed(key.buf)
		if err != nil {
			return nil, err
		}
		hasher.Write(data)
		return NewFixedBuf(32, hasher.Sum(nil))
	}

}