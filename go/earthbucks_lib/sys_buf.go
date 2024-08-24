package earthbucks

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"

	"github.com/mr-tron/base58"
)

// SysBuf is a wrapper around a byte slice with additional methods
type SysBuf struct {
    buf []byte
}

// NewSysBuf creates a new SysBuf with the given byte slice
func NewSysBuf(data []byte) *SysBuf {
    return &SysBuf{buf: data}
}

// Alloc allocates a new SysBuf with the given size and optional fill byte
func (SysBuf) Alloc(size int, fill byte) *SysBuf {
    buf := make([]byte, size)
    if fill != 0 {
        for i := range buf {
            buf[i] = fill
        }
    }
    return NewSysBuf(buf)
}

// FromHex creates a SysBuf from a hex string
func (SysBuf) FromHex(hexStr string) (*SysBuf, error) {
    data, err := hex.DecodeString(hexStr)
    if err != nil {
        return nil, ErrInvalidHex
    }
    return NewSysBuf(data), nil
}

// ToHex converts the SysBuf to a hex string
func (s *SysBuf) ToHex() string {
    return hex.EncodeToString(s.buf)
}

// FromBase64 creates a SysBuf from a base64 string
func (SysBuf) FromBase64(base64Str string) (*SysBuf, error) {
    data, err := base64.StdEncoding.DecodeString(base64Str)
    if err != nil {
        return nil, ErrInvalidEncoding
    }
    return NewSysBuf(data), nil
}

// ToBase64 converts the SysBuf to a base64 string
func (s *SysBuf) ToBase64() string {
    return base64.StdEncoding.EncodeToString(s.buf)
}

// FromBase58 creates a SysBuf from a base58 string
func (SysBuf) FromBase58(base58Str string) (*SysBuf, error) {
    data, err := base58.Decode(base58Str)
    if len(data) == 0 || err != nil {
        return nil, ErrInvalidEncoding
    }
    return NewSysBuf(data), nil
}

// ToBase58 converts the SysBuf to a base58 string
func (s *SysBuf) ToBase58() string {
    return base58.Encode(s.buf)
}

// FromRandom creates a SysBuf with random content of a specific size
func (SysBuf) FromRandom(size int) (*SysBuf, error) {
    buf := make([]byte, size)
    _, err := rand.Read(buf)
    if err != nil {
        return nil, err
    }
    return NewSysBuf(buf), nil
}