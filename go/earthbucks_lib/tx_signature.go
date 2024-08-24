package earthbucks

import (
    "errors"
)

// Value returns the underlying uint8 value
func (u U8) Value() uint8 {
    return u.value
}


// Define SysBuf type representing a general buffer (similar to Node.js's Buffer)
type SysBuf []byte

// Alloc allocates a new SysBuf of the given size
func Alloc(size int) SysBuf {
    return make(SysBuf, size)
}

// Concat concatenates multiple SysBuf into a single SysBuf
func Concat(bufs []SysBuf) SysBuf {
    totalLen := 0
    for _, buf := range bufs {
        totalLen += len(buf)
    }
    result := make(SysBuf, totalLen)
    offset := 0
    for _, buf := range bufs {
        copy(result[offset:], buf)
        offset += len(buf)
    }
    return result
}

// WriteUInt8 writes a uint8 value to SysBuf at the specified index
func (buf SysBuf) WriteUInt8(index int, value uint8) {
    buf[index] = value
}

// ReadUInt8 reads a uint8 value from SysBuf at the specified index
func (buf SysBuf) ReadUInt8(index int) uint8 {
    return buf[index]
}

// TxSignature represents a transaction signature
type TxSignature struct {
    HashType U8
    SigBuf   *FixedBuf
}

// Define constants
const (
    SIGHASH_ALL             = 0x00000001
    SIGHASH_NONE            = 0x00000002
    SIGHASH_SINGLE          = 0x00000003
    SIGHASH_ANYONECANPAY    = 0x00000080
    TXSIGNATURE_SIZE        = 65
)

// NewTxSignature creates a new TxSignature instance
func NewTxSignature(hashType U8, sigBuf *FixedBuf) *TxSignature {
    return &TxSignature{
        HashType: hashType,
        SigBuf:   sigBuf,
    }
}

// ToBuf converts the TxSignature to a buffer
func (ts *TxSignature) ToBuf() SysBuf {
    hashTypeBuf := Alloc(1)
    hashTypeBuf.WriteUInt8(0, ts.HashType.Value())
    return Concat([]SysBuf{hashTypeBuf, ts.SigBuf.buf})
}

// FromBuf creates a TxSignature from a buffer
func (ts *TxSignature) FromBuf(buf SysBuf) (tx *TxSignature, err error) {
    if len(buf) != TXSIGNATURE_SIZE {
        return nil, errors.New("invalid TxSignature length")
    }

    hashType, err := NewU8(buf.ReadUInt8(0))
    if err != nil {
        return nil, err
    }
    sigBuf := buf[1:]
    fixedBuf,err := FromBufFixed(64, sigBuf)

    if err != nil {
        return nil, err
    }

    if fixedBuf == nil {
        return nil, errors.New("failed to create FixedBuf")
    }

    return NewTxSignature(*hashType, fixedBuf), nil
}
