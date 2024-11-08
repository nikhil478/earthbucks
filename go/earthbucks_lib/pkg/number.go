package earthbucks

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
)

var (
	ErrInvalidValue   = errors.New("value is not valid")
	ErrDivisionByZero = errors.New("division by zero")
)

type IBasicNumber interface {
	Add(other BasicNumber) (BasicNumber, error)
	Sub(other BasicNumber) (BasicNumber, error)
	Mul(other BasicNumber) (BasicNumber, error)
	Div(other BasicNumber) (BasicNumber, error)
	Bn() *big.Int
	N() float64
	ToBEBuf() []byte
	ToHex() string
}

type BasicNumber struct {
	value *big.Int
	min   *big.Int
	max   *big.Int
}

func NewBasicNumber(value, min, max *big.Int) (*BasicNumber, error) {
	if value.Cmp(min) < 0 || value.Cmp(max) > 0 {
		return nil, fmt.Errorf("Value %s is not a valid number", value)
	}

	return &BasicNumber{
		value: new(big.Int).Set(value),
		min:   new(big.Int).Set(min),
		max:   new(big.Int).Set(max),
	}, nil
}

type U8 struct {
	*BasicNumber
}

func NewU8(value big.Int) (*U8, error) {
	min := new(big.Int)
	min.SetString("0x00", 0) 
	max := new(big.Int)
	max.SetString("0xff", 0) 
	basicNumber, err := NewBasicNumber(&value, min, max)
	if err != nil {
		return nil, err
	}
	return &U8{BasicNumber: basicNumber}, nil
}

func (u *U8) Add(other *U8) (*U8, error) {
	result := new(big.Int).Add(u.value, other.value)
	return NewU8(*result)
}

func (u *U8) Sub(other *U8) (*U8, error) {
	result := new(big.Int).Sub(u.value, other.value)
	return NewU8(*result)
}

func (u *U8) Mul(other *U8) (*U8, error) {
	result := new(big.Int).Mul(u.value,other.value)
	return NewU8(*result)
}

func (u *U8) Div(other *U8) (*U8, error) {
	result := new(big.Int).Div(u.value, other.value)
	return NewU8(*result)
}

func (u *U8) Bn() *big.Int {
	return u.value
}

func (u *U8) N() (float64, big.Accuracy) {
	return u.value.Float64()
}

func (u *U8) ToBEBuf() []byte {
	return u.value.Bytes()
}

func (u *U8) ToHex() string {
	return hex.EncodeToString(u.ToBEBuf())
}

func U8FromBEBuf(buf *[]byte) (*U8, error) {
	bigInt := new(big.Int).SetBytes(*buf)
	return NewU8(*bigInt)
}

func U8FromHex(hex *string) (*U8, error) {
	n := 1 
	ebxBuf, err := EbxBufFromHex(&n, hex)
	if err != nil {
		return nil, err
	}
	return U8FromBEBuf(ebxBuf.buf)
}

type U16 struct {
	*BasicNumber
}

func NewU16(value big.Int) (*U16, error) {
	min := new(big.Int)
	min.SetString("0x0000", 0) 
	max := new(big.Int)
	max.SetString("0xffff", 0) 
	basicNumber, err := NewBasicNumber(&value, min, max)
	if err != nil {
		return nil, err
	}
	return &U16{BasicNumber: basicNumber}, nil
}

func (u *U16) Add(other *U16) (*U16, error) {
	result := new(big.Int).Add(u.value, other.value)
	return NewU16(*result)
}

func (u *U16) Sub(other *U16) (*U16, error) {
	result := new(big.Int).Sub(u.value, other.value)
	return NewU16(*result)
}

func (u *U16) Mul(other *U16) (*U16, error) {
	result := new(big.Int).Mul(u.value,other.value)
	return NewU16(*result)
}

func (u *U16) Div(other *U16) (*U16, error) {
	result := new(big.Int).Div(u.value, other.value)
	return NewU16(*result)
}

func (u *U16) Bn() *big.Int {
	return u.value
}

func (u *U16) N() (float64, big.Accuracy) {
	return u.value.Float64()
}

func (u *U16) ToBEBuf() []byte {
	return u.value.Bytes()
}

func (u *U16) ToHex() string {
	return hex.EncodeToString(u.ToBEBuf())
}

func U16FromBEBuf(buf *[]byte) (*U16, error) {
	bigInt := new(big.Int).SetBytes(*buf)
	return NewU16(*bigInt)
}

func U16FromHex(hex *string) (*U16, error) {
	n := 2
	ebxBuf, err := EbxBufFromHex(&n, hex)
	if err != nil {
		return nil, err
	}
	return U16FromBEBuf(ebxBuf.buf)
}

type U32 struct {
	*BasicNumber
}

func NewU32(value big.Int) (*U32, error) {
	min := new(big.Int)
	min.SetString("0x00000000", 0) 
	max := new(big.Int)
	max.SetString("0xffffffff", 0) 
	basicNumber, err := NewBasicNumber(&value, min, max)
	if err != nil {
		return nil, err
	}
	return &U32{BasicNumber: basicNumber}, nil
}


func (u *U32) Add(other *U32) (*U32, error) {
	result := new(big.Int).Add(u.value, other.value)
	return NewU32(*result)
}

func (u *U32) Sub(other *U32) (*U32, error) {
	result := new(big.Int).Sub(u.value, other.value)
	return NewU32(*result)
}

func (u *U32) Mul(other *U32) (*U32, error) {
	result := new(big.Int).Mul(u.value,other.value)
	return NewU32(*result)
}

func (u *U32) Div(other *U32) (*U32, error) {
	result := new(big.Int).Div(u.value, other.value)
	return NewU32(*result)
}

func (u *U32) Bn() *big.Int {
	return u.value
}

func (u *U32) N() (float64, big.Accuracy) {
	return u.value.Float64()
}

func (u *U32) ToBEBuf() []byte {
	return u.value.Bytes()
}

func (u *U32) ToHex() string {
	return hex.EncodeToString(u.ToBEBuf())
}

func U32FromBEBuf(buf *[]byte) (*U32, error) {
	bigInt := new(big.Int).SetBytes(*buf)
	return NewU32(*bigInt)
}

func U32FromHex(hex *string) (*U32, error) {
	n := 4
	ebxBuf, err := EbxBufFromHex(&n, hex)
	if err != nil {
		return nil, err
	}
	return U32FromBEBuf(ebxBuf.buf)
}

type U64 struct {
	*BasicNumber
}


func NewU64(value big.Int) (*U64, error) {
	min := new(big.Int)
	min.SetString("0x0000000000000000", 0) 
	max := new(big.Int)
	max.SetString("0xffffffffffffffff", 0) 
	basicNumber, err := NewBasicNumber(&value, min, max)
	if err != nil {
		return nil, err
	}
	return &U64{BasicNumber: basicNumber}, nil
}


func (u *U64) Add(other *U64) (*U64, error) {
	result := new(big.Int).Add(u.value, other.value)
	return NewU64(*result)
}

func (u *U64) Sub(other *U64) (*U64, error) {
	result := new(big.Int).Sub(u.value, other.value)
	return NewU64(*result)
}

func (u *U64) Mul(other *U64) (*U64, error) {
	result := new(big.Int).Mul(u.value,other.value)
	return NewU64(*result)
}

func (u *U64) Div(other *U64) (*U64, error) {
	result := new(big.Int).Div(u.value, other.value)
	return NewU64(*result)
}

func (u *U64) Bn() *big.Int {
	return u.value
}

func (u *U64) N() (float64, big.Accuracy) {
	return u.value.Float64()
}

func (u *U64) ToBEBuf() []byte {
	return u.value.Bytes()
}

func (u *U64) ToHex() string {
	return hex.EncodeToString(u.ToBEBuf())
}

func U64FromBEBuf(buf *[]byte) (*U64, error) {
	bigInt := new(big.Int).SetBytes(*buf)
	return NewU64(*bigInt)
}

func U64FromHex(hex *string) (*U64, error) {
	n := 8
	ebxBuf, err := EbxBufFromHex(&n, hex)
	if err != nil {
		return nil, err
	}
	return U64FromBEBuf(ebxBuf.buf)
}

type U128 struct {
	*BasicNumber
}

func NewU128(value big.Int) (*U128, error) {
	min := new(big.Int)
	min.SetString("0x00000000000000000000000000000000", 0) 
	max := new(big.Int)
	max.SetString("0xffffffffffffffffffffffffffffffff", 0) 
	basicNumber, err := NewBasicNumber(&value, min, max)
	if err != nil {
		return nil, err
	}
	return &U128{BasicNumber: basicNumber}, nil
}


func (u *U128) Add(other *U128) (*U128, error) {
	result := new(big.Int).Add(u.value, other.value)
	return NewU128(*result)
}

func (u *U128) Sub(other *U128) (*U128, error) {
	result := new(big.Int).Sub(u.value, other.value)
	return NewU128(*result)
}

func (u *U128) Mul(other *U128) (*U128, error) {
	result := new(big.Int).Mul(u.value,other.value)
	return NewU128(*result)
}

func (u *U128) Div(other *U128) (*U128, error) {
	result := new(big.Int).Div(u.value, other.value)
	return NewU128(*result)
}

func (u *U128) Bn() *big.Int {
	return u.value
}

func (u *U128) N() (float64, big.Accuracy) {
	return u.value.Float64()
}

func (u *U128) ToBEBuf() []byte {
	return u.value.Bytes()
}

func (u *U128) ToHex() string {
	return hex.EncodeToString(u.ToBEBuf())
}

func U128FromBEBuf(buf *[]byte) (*U128, error) {
	bigInt := new(big.Int).SetBytes(*buf)
	return NewU128(*bigInt)
}

func U128FromHex(hex *string) (*U128, error) {
	n := 16
	ebxBuf, err := EbxBufFromHex(&n, hex)
	if err != nil {
		return nil, err
	}
	return U128FromBEBuf(ebxBuf.buf)
}

type U256 struct {
	*BasicNumber
}

func NewU256(value big.Int) (*U256, error) {
	min := new(big.Int)
	min.SetString("0x0000000000000000000000000000000000000000000000000000000000000000", 0) 
	max := new(big.Int)
	max.SetString("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 0) 
	basicNumber, err := NewBasicNumber(&value, min, max)
	if err != nil {
		return nil, err
	}
	return &U256{BasicNumber: basicNumber}, nil
}

func (u *U256) Add(other *U256) (*U256, error) {
	result := new(big.Int).Add(u.value, other.value)
	return NewU256(*result)
}

func (u *U256) Sub(other *U256) (*U256, error) {
	result := new(big.Int).Sub(u.value, other.value)
	return NewU256(*result)
}

func (u *U256) Mul(other *U256) (*U256, error) {
	result := new(big.Int).Mul(u.value,other.value)
	return NewU256(*result)
}

func (u *U256) Div(other *U256) (*U256, error) {
	result := new(big.Int).Div(u.value, other.value)
	return NewU256(*result)
}

func (u *U256) Bn() *big.Int {
	return u.value
}

func (u *U256) N() (float64, big.Accuracy) {
	return u.value.Float64()
}

func (u *U256) ToBEBuf() []byte {
	return u.value.Bytes()
}

func (u *U256) ToHex() string {
	return hex.EncodeToString(u.ToBEBuf())
}

func U256FromBEBuf(buf *[]byte) (*U256, error) {
	bigInt := new(big.Int).SetBytes(*buf)
	return NewU256(*bigInt)
}

func U256FromHex(hex *string) (*U256, error) {
	n := 16
	ebxBuf, err := EbxBufFromHex(&n, hex)
	if err != nil {
		return nil, err
	}
	return U256FromBEBuf(ebxBuf.buf)
}