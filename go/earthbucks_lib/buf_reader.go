package earthbucks

import (
	"encoding/binary"
	"errors"
	"math/big"
)


var (
	ErrNotEnoughData       = errors.New("not enough data")
	ErrNonMinimalEncoding  = errors.New("non-minimal encoding")
	ErrInsufficientPrecision = errors.New("insufficient precision")
	ErrValueExceeds128Bits = errors.New("err value exceeds 128 bits") 
)

// BufReader reads from a byte buffer
type BufReader struct {
	buf []byte
	pos int
}

// NewBufReader creates a new BufReader
func NewBufReader(buf []byte) *BufReader {
	return &BufReader{buf: buf, pos: 0}
}

// EOF checks if the end of the buffer has been reached
func (r *BufReader) EOF() bool {
	return r.pos >= len(r.buf)
}

// Read reads a specific length from the buffer
func (r *BufReader) Read(length int) ([]byte, error) {
	if r.pos+length > len(r.buf) { // len() used correctly on the slice
		return nil, ErrNotEnoughData
	}
	data := r.buf[r.pos : r.pos+length] // Slicing the buffer correctly
	r.pos += length
	return data, nil
}

// ReadFixed reads a fixed-size buffer
func (r *BufReader) ReadFixed(len int) (*FixedBuf, error) {
	data, err := r.Read(len)
	if err != nil {
		return nil, err
	}
	return FixedBuf.FromBuf(FixedBuf{}, len, data)
}

// ReadRemainder reads the remainder of the buffer
func (r *BufReader) ReadRemainder() ([]byte, error) {
	return r.Read(len(r.buf) - r.pos)
}

// ReadU8 reads an 8-bit unsigned integer
func (r *BufReader) ReadU8() (*U8, error) {
	data, err := r.Read(1)
	if err != nil {
		return nil, err
	}
	return NewU8(data[0])
}

// ReadU16BE reads a 16-bit unsigned integer in big-endian
func (r *BufReader) ReadU16BE() (*U16, error) {
	data, err := r.Read(2)
	if err != nil {
		return nil, err
	}
	return NewU16(binary.BigEndian.Uint16(data))
}

// ReadU32BE reads a 32-bit unsigned integer in big-endian
func (r *BufReader) ReadU32BE() (*U32, error) {
	data, err := r.Read(4)
	if err != nil {
		return nil, err
	}
	return NewU32(binary.BigEndian.Uint32(data))
}

// ReadU64BE reads a 64-bit unsigned integer in big-endian
func (r *BufReader) ReadU64BE() (*U64, error) {
	data, err := r.Read(8)
	if err != nil {
		return nil, err
	}
	return NewU64(binary.BigEndian.Uint64(data))
}

// ReadU128BE reads a 128-bit unsigned integer in big-endian
func (r *BufReader) ReadU128BE() (*U128, error) {
    // Read 16 bytes from the buffer
    data, err := r.Read(16)
    if err != nil {
        return nil, err
    }

    // Convert the byte slice to a big.Int
    val := new(big.Int).SetBytes(data)

    // Ensure that the value fits within 128 bits
    if val.BitLen() > 128 {
        return nil, ErrValueExceeds128Bits
    }

    // Extract high and low 64 bits from the big.Int
    high := val.Rsh(val, 64).Uint64()
    low := val.Uint64()

    // Create a uint128 value
    u128 := uint128{
        high: high,
        low:  low,
    }

    // Create a U128 from the uint128 value
    return NewU128(u128)
}




// ReadU256BE reads a 256-bit unsigned integer in big-endian
func (r *BufReader) ReadU256BE() (*U256, error) {
	data, err := r.Read(32)
	if err != nil {
		return nil, err
	}
	return NewU256(new(big.Int).SetBytes(data)), nil
}

// ReadVarIntBuf reads a variable-length integer buffer
func (r *BufReader) ReadVarIntBuf() ([]byte, error) {
	first, err := r.ReadU8()
	if err != nil {
		return nil, err
	}
	switch first.value {
	case 0xfd:
		buf, err := r.Read(2)
		if err != nil {
			return nil, err
		}
		if binary.BigEndian.Uint16(buf) < 0xfd {
			return nil, ErrNonMinimalEncoding
		}
		return append([]byte{first.value}, buf...), nil
	case 0xfe:
		buf, err := r.Read(4)
		if err != nil {
			return nil, err
		}
		if binary.BigEndian.Uint32(buf) < 0x10000 {
			return nil, ErrNonMinimalEncoding
		}
		return append([]byte{first.value}, buf...), nil
	case 0xff:
		buf, err := r.Read(8)
		if err != nil {
			return nil, err
		}
		if binary.BigEndian.Uint64(buf) < 0x100000000 {
			return nil, ErrNonMinimalEncoding
		}
		return append([]byte{first.value}, buf...), nil
	default:
		return []byte{first.value}, nil
	}
}

// ReadVarInt reads a variable-length integer
func (r *BufReader) ReadVarInt() (*U64, error) {
	buf, err := r.ReadVarIntBuf()
	if err != nil {
		return nil, err
	}
	first := buf[0]
	var value *big.Int
	switch first {
	case 0xfd:
		value = new(big.Int).SetUint64(uint64(binary.BigEndian.Uint16(buf[1:])))
	case 0xfe:
		value = new(big.Int).SetUint64(uint64(binary.BigEndian.Uint32(buf[1:])))
	case 0xff:
		value = new(big.Int).SetUint64(binary.BigEndian.Uint64(buf[1:]))
	default:
		value = new(big.Int).SetUint64(uint64(first))
	}
	return NewU64(value.Uint64())
}
