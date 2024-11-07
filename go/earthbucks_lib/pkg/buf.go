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

func encodeHex(data []byte) string {
    return hex.EncodeToString(data)
}

func decodeHex(hexStr string) ([]byte, error) {
    return hex.DecodeString(hexStr)
}

type EbxBuf struct {
    buf *[]byte
}

func NewEbxBuf(size *int, data *[]byte) (*EbxBuf, error) {
    if len(*data) != *size {
        return nil, ErrInvalidSize
    }
    return &EbxBuf{buf: data}, nil
}

func EbxBufFromBuf(size *int, buf *[]byte) (*EbxBuf, error) {
    return NewEbxBuf(size, buf)
}

func EbxBufAlloc(size *int, fill byte) (*EbxBuf, error) {
    buf := make([]byte, *size)
    // utf-8 encoded as alloc func in js bydefault encode that to utf-8 if buf is string
    // TODO: @ryanxcharles 
    if fill != 0 {
        for i := range buf {
            buf[i] = fill
        }
    }
    return NewEbxBuf(size, &buf)

}

func (ebxBuf *EbxBuf) ToBuf() *[]byte {
    return ebxBuf.buf
}

func EbxBufFromHex(size *int, hexStr *string) (*EbxBuf, error) {
    data, err := hex.DecodeString(*hexStr)
    if err != nil {
        return nil, ErrInvalidHex
    }
    return NewEbxBuf(size, &data)
}

func (ebxBuf *EbxBuf) ToHex() string {
    return encodeHex(*ebxBuf.buf)
}

func EbxBufFromBase64(size *int, base64Str *string) (*EbxBuf, error) {
    data, err := base64.StdEncoding.DecodeString(*base64Str)
    if err != nil {
        return nil, ErrInvalidEncoding
    }
    return NewEbxBuf(size, &data)
}

func (ebxBuf *EbxBuf) ToBase64() string {
    return base64.StdEncoding.EncodeToString(*ebxBuf.buf)
}

func EbxBufFromBase58(size *int, base58Str *string) (*EbxBuf, error) {
    data, err := base58.Decode(*base58Str)
    if len(data) == 0 || err != nil{
        return nil, ErrInvalidEncoding
    }
    return NewEbxBuf(size, &data)
}

func (ebxBuf *EbxBuf) ToBase58() string {
    return base58.Encode(*ebxBuf.buf)
}

func EbxBufFromRandom(size *int) (*EbxBuf, error) {
    buf := make([]byte, *size)
    _, err := rand.Read(buf)
    if err != nil {
        return nil, err
    }
    return NewEbxBuf(size, &buf)
}

type FixedBuf struct {
	*EbxBuf
	size *int
}

func NewFixedBuf(size *int, data *[]byte) (*FixedBuf, error) {
	buf, err := NewEbxBuf(size, data)
	if err != nil {
		return nil, err
	}
	return &FixedBuf{EbxBuf: buf, size: size}, nil
}

func FixedBufFromBuf(size *int, buf *[]byte) (*FixedBuf, error) {
	return NewFixedBuf(size, buf)
}

func FixedBufAlloc(size *int, fill byte) (*FixedBuf, error) {
	buf := make([]byte, *size)
	for i := range buf {
		buf[i] = fill
	}
	return NewFixedBuf(size, &buf)
}

func FixedBufFromHex(size *int, hexStr *string) (*FixedBuf, error) {
	data, err := hex.DecodeString(*hexStr)
	if err != nil {
		return nil, ErrInvalidHex
	}
	return NewFixedBuf(size, &data)
}

func FixedBufFromBase58(size *int, base58Str *string) (*FixedBuf, error) {
	data, err := base58.Decode(*base58Str)
	if err != nil {
		return nil, ErrInvalidEncoding
	}
	return NewFixedBuf(size, &data)
}

func FixedBufFromRandom(size *int) (*FixedBuf, error) {
	buf := make([]byte, *size)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}
	return NewFixedBuf(size, &buf)
}