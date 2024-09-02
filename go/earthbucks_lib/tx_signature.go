package earthbucks

import "errors"

// Constants for hash types and size
const (
	SIGHASH_ALL          = 0x00000001
	SIGHASH_NONE         = 0x00000002
	SIGHASH_SINGLE       = 0x00000003
	SIGHASH_ANYONECANPAY = 0x00000080
	TxSignatureSize      = 65
)

// TxSignature represents a transaction signature
type TxSignature struct {
	hashType *U8
	sigBuf   *FixedBuf
}

// NewTxSignature creates a new TxSignature
func NewTxSignature(hashType *U8, sigBuf *FixedBuf) *TxSignature {
	return &TxSignature{hashType: hashType, sigBuf: sigBuf}
}

// ToBuf serializes the TxSignature into a byte buffer
func (tx *TxSignature) ToBuf() []byte {
	buf := make([]byte, TxSignatureSize)
	buf[0] = tx.hashType.value
	copy(buf[1:], tx.sigBuf.buf)
	return buf
}

// FromBuf deserializes a byte buffer into a TxSignature
func TxSignatureFromBuf(buf []byte) (*TxSignature, error) {
	if len(buf) != TxSignatureSize {
		return nil, errors.New("invalid TxSignature length")
	}
	hashType, err := NewU8(buf[0])
    if  err != nil {
        return nil, err
    }
	sigBuf, err := NewFixedBuf(64, buf[1:])
	if err != nil {
		return nil, err
	}
	return NewTxSignature(hashType, sigBuf), nil
}