package earthbucks

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/mr-tron/base58"
)

// Error definitions
var (
	ErrInvalidSize     = errors.New("invalid size error")
	ErrInvalidHex      = errors.New("invalid hex error")
	ErrInvalidEncoding = errors.New("invalid encoding error")
)

// EbxBuf represents a buffer with additional methods for encoding/decoding.
type EbxBuf struct {
	buf []byte
}

// NewEbxBuf creates a new EbxBuf instance.
func NewEbxBuf(size int, data []byte) (*EbxBuf, error) {
	if len(data) != size {
		return nil, ErrInvalidSize
	}
	return &EbxBuf{buf: data}, nil
}

// FromBuf creates an EbxBuf from a byte slice.
func FromBuf(size int, data []byte) (*EbxBuf, error) {
	return NewEbxBuf(size, data)
}

// AllocateBuffer creates an EbxBuf with allocated size and optional fill value.
func AllocateBuffer(size int, fill byte) (*EbxBuf, error) {
	data := make([]byte, size)
	if fill != 0 {
		for i := range data {
			data[i] = fill
		}
	}
	return NewEbxBuf(size, data)
}

// FromHex creates an EbxBuf from a hex string.
func FromHex(size int, hexStr string) (*EbxBuf, error) {
	data, err := DecodeHex(hexStr)
	if err != nil {
		return nil, err
	}
	return NewEbxBuf(size, data)
}

// ToHex encodes the buffer into a hex string.
func (b *EbxBuf) ToHex() string {
	return EncodeHex(b.buf)
}

// FromBase64 creates an EbxBuf from a base64 string.
func FromBase64(size int, base64Str string) (*EbxBuf, error) {
	// Base64 decoding should be implemented if needed
	return nil, errors.New("base64 decoding is not yet implemented")
}

// ToBase64 encodes the buffer into a base64 string.
func (b *EbxBuf) ToBase64() string {
	return base64.StdEncoding.EncodeToString(b.buf)
}

// FromBase58 creates an EbxBuf from a base58 string.
func FromBase58(size int, base58Str string) (*EbxBuf, error) {
	data, err := DecodeBase58(base58Str)
	if err != nil {
		return nil, err
	}
	return NewEbxBuf(size, data)
}

// ToBase58 encodes the buffer into a base58 string.
func (b *EbxBuf) ToBase58() string {
	return EncodeBase58(b.buf)
}

// FromRandom creates an EbxBuf with random data.
func FromRandom(size int) (*EbxBuf, error) {
	data := make([]byte, size)
	_, err := rand.Read(data)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random data: %w", err)
	}
	return NewEbxBuf(size, data)
}

// FixedBuf represents a fixed-size buffer.
type FixedBuf struct {
	EbxBuf
	size int
}

// NewFixedBuf creates a new FixedBuf instance.
func NewFixedBuf(size int, data []byte) (*FixedBuf, error) {
	buf, err := NewEbxBuf(size, data)
	if err != nil {
		return nil, err
	}
	return &FixedBuf{EbxBuf: *buf, size: size}, nil
}

// AllocateFixedBuffer creates a FixedBuf with allocated size and optional fill value.
func AllocateFixedBuffer(size int, fill byte) (*FixedBuf, error) {
	data := make([]byte, size)
	if fill != 0 {
		for i := range data {
			data[i] = fill
		}
	}
	return NewFixedBuf(size, data)
}

// FromHexFixed creates a FixedBuf from a hex string.
func FromHexFixed(size int, hexStr string) (*FixedBuf, error) {
	data, err := DecodeHex(hexStr)
	if err != nil {
		return nil, err
	}
	return NewFixedBuf(size, data)
}

// FromBase58Fixed creates a FixedBuf from a base58 string.
func FromBase58Fixed(size int, base58Str string) (*FixedBuf, error) {
	data, err := DecodeBase58(base58Str)
	if err != nil {
		return nil, err
	}
	return NewFixedBuf(size, data)
}

// FromRandomFixed creates a FixedBuf with random data.
func FromRandomFixed(size int) (*FixedBuf, error) {
	data := make([]byte, size)
	_, err := rand.Read(data)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random data: %w", err)
	}
	return NewFixedBuf(size, data)
}


// FromBuf initializes a FixedBuf from a given buffer
func FromBufFixed(size int, buf []byte) (fb *FixedBuf, err error) {
    if len(buf) != size {
        return nil, nil
    }
    fb, err = NewFixedBuf(size, buf)
    copy(fb.buf, buf)
    return fb , err
}

// IsValidHex checks if a string is a valid hex value.
func IsValidHex(hexStr string) bool {
	_, err := hex.DecodeString(hexStr)
	return err == nil && len(hexStr)%2 == 0
}

// DecodeHex decodes a hex string into a byte slice.
func DecodeHex(hexStr string) ([]byte, error) {
	if !IsValidHex(hexStr) {
		return nil, ErrInvalidHex
	}
	return hex.DecodeString(hexStr)
}

// EncodeHex encodes a byte slice into a hex string.
func EncodeHex(data []byte) string {
	return hex.EncodeToString(data)
}

// DecodeBase58 decodes a base58 string into a byte slice.
func DecodeBase58(base58Str string) ([]byte, error) {
	data, err := base58.Decode(base58Str)
	if err != nil {
		return nil, ErrInvalidEncoding
	}
	return data, nil
}

// EncodeBase58 encodes a byte slice into a base58 string.
func EncodeBase58(data []byte) string {
	return base58.Encode(data)
}