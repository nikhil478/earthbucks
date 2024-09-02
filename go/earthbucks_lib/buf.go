package earthbucks

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"

	"github.com/mr-tron/base58"
)

var (
    ErrInvalidSize       = errors.New("invalid size")
    ErrInvalidHex        = errors.New("invalid hex string")
    ErrInvalidEncoding   = errors.New("invalid encoding")
)

// isValidHex checks if a string is a valid hex string
func isValidHex(hexStr string) bool {
    _, err := hex.DecodeString(hexStr)
    return err == nil && len(hexStr)%2 == 0
}

// encodeHex encodes a byte slice to a hex string
func encodeHex(data []byte) string {
    return hex.EncodeToString(data)
}

// decodeHex decodes a hex string to a byte slice
func decodeHex(hexStr string) ([]byte, error) {
    if !isValidHex(hexStr) {
        return nil, ErrInvalidHex
    }
    return hex.DecodeString(hexStr)
}

// EbxBuf represents a buffer with various encoding/decoding methods
type EbxBuf struct {
    buf []byte
}

// NewEbxBuf creates a new EbxBuf with a specific size and content
func NewEbxBuf(size int, data []byte) (*EbxBuf, error) {
    if len(data) != size {
        return nil, ErrInvalidSize
    }
    return &EbxBuf{buf: data}, nil
}

// FromBuf creates an EbxBuf from a buffer
func EbxBufFromBuf(size int, buf []byte) (*EbxBuf, error) {
    return NewEbxBuf(size, buf)
}

// Alloc allocates a new EbxBuf with a specific size and optional fill byte
func EbxBufAlloc(size int, fill byte) (*EbxBuf, error) {
    buf := make([]byte, size)
    if fill != 0 {
        for i := range buf {
            buf[i] = fill
        }
    }
    return NewEbxBuf(size, buf)
}

// FromHex creates an EbxBuf from a hex string
func EbxBufFromHex(size int, hexStr string) (*EbxBuf, error) {
    data, err := hex.DecodeString(hexStr)
    if err != nil {
        return nil, ErrInvalidHex
    }
    return NewEbxBuf(size, data)
}

// ToHex converts the EbxBuf to a hex string
func (e *EbxBuf) ToHex() string {
    return hex.EncodeToString(e.buf)
}

// FromBase64 creates an EbxBuf from a base64 string
func EbxBufFromBase64(size int, base64Str string) (*EbxBuf, error) {
    data, err := base64.StdEncoding.DecodeString(base64Str)
    if err != nil {
        return nil, ErrInvalidEncoding
    }
    return NewEbxBuf(size, data)
}

// ToBase64 converts the EbxBuf to a base64 string
func (e *EbxBuf) ToBase64() string {
    return base64.StdEncoding.EncodeToString(e.buf)
}

// FromBase58 creates an EbxBuf from a base58 string
func EbxBufFromBase58(size int, base58Str string) (*EbxBuf, error) {
    data, err := base58.Decode(base58Str)
    if len(data) == 0 || err != nil{
        return nil, ErrInvalidEncoding
    }
    return NewEbxBuf(size, data)
}

// ToBase58 converts the EbxBuf to a base58 string
func (e *EbxBuf) ToBase58() string {
    return base58.Encode(e.buf)
}

// FromRandom creates an EbxBuf with random content of a specific size
func EbxBufFromRandom(size int) (*EbxBuf, error) {
    buf := make([]byte, size)
    _, err := rand.Read(buf)
    if err != nil {
        return nil, err
    }
    return NewEbxBuf(size, buf)
}

// FixedBuf is a specialized buffer with a fixed size
type FixedBuf struct {
	EbxBuf
	size int
}

// NewFixedBuf creates a new FixedBuf with a specific size and content
func NewFixedBuf(size int, data []byte) (*FixedBuf, error) {
	buf, err := NewEbxBuf(size, data)
	if err != nil {
		return nil, err
	}
	return &FixedBuf{EbxBuf: *buf, size: size}, nil
}

// FromBuf creates a FixedBuf from a buffer
func FixedBufFromBuf(size int, buf []byte) (*FixedBuf, error) {
	return NewFixedBuf(size, buf)
}

// Alloc allocates a new FixedBuf with a specific size and optional fill byte
func FixedBufAlloc(size int, fill byte) (*FixedBuf, error) {
	buf := make([]byte, size)
	for i := range buf {
		buf[i] = fill
	}
	return NewFixedBuf(size, buf)
}

// FromHex creates a FixedBuf from a hex string
func FixedBufFromHex(size int, hexStr string) (*FixedBuf, error) {
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, ErrInvalidHex
	}
	return NewFixedBuf(size, data)
}

// FromBase58 creates a FixedBuf from a base58 string
func FixedBufFromBase58(size int, base58Str string) (*FixedBuf, error) {
	data, err := base58.Decode(base58Str)
	if len(data) == 0 || err != nil {
		return nil, ErrInvalidEncoding
	}
	return NewFixedBuf(size, data)
}

// FromRandom creates a FixedBuf with random content of a specific size
func FixedBufFromRandom(size int) (*FixedBuf, error) {
	buf := make([]byte, size)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}
	return NewFixedBuf(size, buf)
}