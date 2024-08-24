package earthbucks

import (
	"github.com/zeebo/blake3"
)

// HashFunction defines a function that hashes input data.
type HashFunction func(input []byte) (*FixedBuf, error)

// MacFunction defines a function that performs a keyed hash.
type MacFunction func(key *FixedBuf, data []byte) (*FixedBuf, error)

var (
	blake3Hash         HashFunction
	doubleBlake3Hash  HashFunction
	blake3Mac          MacFunction
)

func init() {
	blake3Hash = func(data []byte) (*FixedBuf, error) {
		hash := blake3.New()
		hash.Write(data)
		return NewFixedBuf(32, hash.Sum(nil))
	}

	doubleBlake3Hash = func(data []byte) (*FixedBuf, error) {
		hash1, err := blake3Hash(data)
		if err != nil {
			return nil, err
		}
		return blake3Hash(hash1.buf)
	}

	blake3Mac = func(key *FixedBuf, data []byte) (*FixedBuf, error) {
		hasher,err := blake3.NewKeyed(key.buf)
		if err != nil {
			return nil, err
		}
		hasher.Write(data)
		return NewFixedBuf(32, hasher.Sum(nil))
	}

}