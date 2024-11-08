package earthbucks

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
)


var (
	ErrNotEnoughData       = errors.New("not enough bytes in the buffer to read")
	ErrNonMinimalEncoding  = errors.New("non-minimal encoding")
	ErrInsufficientPrecision = errors.New("insufficient precision")
	ErrValueExceeds128Bits = errors.New("err value exceeds 128 bits") 
)

type BufReader struct {
	buf *[]byte
	pos *int
}

func NewBufReader(buf *[]byte) *BufReader {
	n := 0
	return &BufReader{buf: buf, pos: &n}
}

func (r *BufReader) EOF() bool {
	return *r.pos >= len(*r.buf)
}

func (r *BufReader) Read(length *int) (*[]byte, error) {
	if *r.pos+*length > len(*r.buf) {
		return nil, ErrNotEnoughData
	}
	buf := *r.buf
	data := buf[*r.pos : *r.pos+ *length]
	newBuf := make([]byte, *length)
	copy(newBuf, data)
	*r.pos += *length

	return &newBuf, nil
}


func (r *BufReader) ReadFixed(len *int) (*FixedBuf, error) {
	data, err := r.Read(len)
	if err != nil {
		return nil, err
	}
	return FixedBufFromBuf(len, data)
}

func (r *BufReader) ReadRemainder() (*[]byte, error) {
	result := len(*r.buf) - *r.pos
	return r.Read(&result)
}


func (r *BufReader) ReadU8() (*U8, error) {
	buf := *r.buf
	sliceBuf := buf[*r.pos:*r.pos+1]
	val, err := U8FromBEBuf(&sliceBuf)
	if err != nil {
		return nil, err
	}
	*r.pos += 1
	return val, err
}

func (r *BufReader) ReadU16BE() (*U16, error) {
	buf := *r.buf
	sliceBuf := buf[*r.pos:*r.pos+1]
	val, err := U16FromBEBuf(&sliceBuf)
	if err != nil {
		return nil, err
	}
	*r.pos += 2
	return val, err
}

func (r *BufReader) ReadU32BE() (*U32, error) {
	buf := *r.buf
	sliceBuf := buf[*r.pos:*r.pos+1]
	val, err := U32FromBEBuf(&sliceBuf)
	if err != nil {
		return nil, err
	}
	*r.pos += 4
	return val, err
}

func (r *BufReader) ReadU64BE() (*U64, error) {
	buf := *r.buf
	sliceBuf := buf[*r.pos:*r.pos+1]
	val, err := U64FromBEBuf(&sliceBuf)
	if err != nil {
		return nil, err
	}
	*r.pos += 8
	return val, err
}

func (r *BufReader) ReadU128BE() (*U128, error) {
	buf := *r.buf
	sliceBuf := buf[*r.pos:*r.pos+1]
	val, err := U128FromBEBuf(&sliceBuf)
	if err != nil {
		return nil, err
	}
	*r.pos += 16
	return val, err
}

func (r *BufReader) ReadU256BE() (*U256, error) {
	buf := *r.buf
	sliceBuf := buf[*r.pos:*r.pos+1]
	val, err := U256FromBEBuf(&sliceBuf)
	if err != nil {
		return nil, err
	}
	*r.pos += 32
	return val, err
}


func (r *BufReader) ReadVarIntBuf() ([]byte, error) {
	first, err := r.ReadU8()
	if err != nil {
		return nil, err
	} 
	var firstBuffer bytes.Buffer
	val , _ := first.N()
	err = binary.Write(&firstBuffer, binary.BigEndian, val)
	if err != nil {
		return nil , err
	}
	switch val {
	case 0xfd:
		n := 2
		buf, err := r.Read(&n)
		if err != nil {
			return nil, err
		}
		if binary.BigEndian.Uint16(*buf) < 0xfd {
			return nil, ErrNonMinimalEncoding
		}
		return append(firstBuffer.Bytes(), *buf...), nil
	case 0xfe:
		n := 4
		buf, err := r.Read(&n)
		if err != nil {
			return nil, err
		}
		if binary.BigEndian.Uint32(*buf) < 0x10000 {
			return nil, ErrNonMinimalEncoding
		}
		return append(firstBuffer.Bytes(), *buf...), nil
	case 0xff:
		n := 8
		buf, err := r.Read(&n)
		if err != nil {
			return nil, err
		}
		if binary.BigEndian.Uint64(*buf) < 0x100000000 {
			return nil, ErrNonMinimalEncoding
		}
		return append(firstBuffer.Bytes(), *buf...), nil
	default:
		return firstBuffer.Bytes(), nil
	}
}


func (r *BufReader) ReadVarInt() (*U64, error) {
	buf, err := r.ReadVarIntBuf()
	if err != nil {
		return nil, err
	}
	first := buf[0]
	var value *big.Int
	switch first {
	case 0xfd:
    	if len(buf) < 3 {
            return nil, fmt.Errorf("not enough data to read 0xfd varint")
        }
        value = new(big.Int).SetUint64(uint64(binary.BigEndian.Uint16(buf[1:3])))
    case 0xfe:
        if len(buf) < 5 {
            return nil, fmt.Errorf("not enough data to read 0xfe varint")
        }
        value = new(big.Int).SetUint64(uint64(binary.BigEndian.Uint32(buf[1:5])))
    case 0xff:
        if len(buf) < 9 {
            return nil, fmt.Errorf("not enough data to read 0xff varint")
        }
        value = new(big.Int).SetUint64(binary.BigEndian.Uint64(buf[1:9]))
    default:
       value = new(big.Int).SetUint64(uint64(first))
    }
	return NewU64(*value)
}
