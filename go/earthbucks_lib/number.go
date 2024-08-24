package earthbucks

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
)

// Error definitions
var (
	ErrInvalidValue   = errors.New("value is not valid")
)

// BasicNumber defines an abstract base type for basic numbers.
type BasicNumber interface {
	Add(other BasicNumber) (BasicNumber, error)
	Sub(other BasicNumber) (BasicNumber, error)
	Mul(other BasicNumber) (BasicNumber, error)
	Div(other BasicNumber) (BasicNumber, error)
	Bn() uint64
	N() uint32
	ToBEBuf() []byte
	ToHex() string
}

// U8 represents an 8-bit unsigned integer.
type U8 struct {
	value uint8
}

func NewU8(value uint8) (*U8, error) {
	if value < 0x00 || value > 0xff {
		return nil, ErrInvalidValue
	}
	return &U8{value: value}, nil
}

func (u *U8) Add(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U8)
	if !ok {
		return nil, fmt.Errorf("invalid type for addition")
	}
	result := u.value + o.value
	return NewU8(result)
}

func (u *U8) Sub(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U8)
	if !ok {
		return nil, fmt.Errorf("invalid type for subtraction")
	}
	result := u.value - o.value
	return NewU8(result)
}

func (u *U8) Mul(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U8)
	if !ok {
		return nil, fmt.Errorf("invalid type for multiplication")
	}
	result := u.value * o.value
	return NewU8(result)
}

func (u *U8) Div(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U8)
	if !ok {
		return nil, fmt.Errorf("invalid type for division")
	}
	if o.value == 0 {
		return nil, fmt.Errorf("division by zero")
	}
	result := u.value / o.value
	return NewU8(result)
}

func (u *U8) Bn() uint64 {
	return uint64(u.value)
}

func (u *U8) N() uint32 {
	return uint32(u.value)
}

func (u *U8) ToBEBuf() []byte {
	buf := make([]byte, 1)
	buf[0] = u.value
	return buf
}

func (u *U8) ToHex() string {
	return hex.EncodeToString(u.ToBEBuf())
}

func FromBEBufU8(buf []byte) (*U8, error) {
	if len(buf) != 1 {
		return nil, ErrInvalidSize
	}
	return NewU8(buf[0])
}

func FromHexU8(hexStr string) (*U8, error) {
	buf, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return FromBEBufU8(buf)
}

// U16 represents a 16-bit unsigned integer.
type U16 struct {
	value uint16
}

func NewU16(value uint16) (*U16, error) {
	if value < 0x0000 || value > 0xffff {
		return nil, ErrInvalidValue
	}
	return &U16{value: value}, nil
}

func (u *U16) Add(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U16)
	if !ok {
		return nil, fmt.Errorf("invalid type for addition")
	}
	result := u.value + o.value
	return NewU16(result)
}

func (u *U16) Sub(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U16)
	if !ok {
		return nil, fmt.Errorf("invalid type for subtraction")
	}
	result := u.value - o.value
	return NewU16(result)
}

func (u *U16) Mul(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U16)
	if !ok {
		return nil, fmt.Errorf("invalid type for multiplication")
	}
	result := u.value * o.value
	return NewU16(result)
}

func (u *U16) Div(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U16)
	if !ok {
		return nil, fmt.Errorf("invalid type for division")
	}
	if o.value == 0 {
		return nil, fmt.Errorf("division by zero")
	}
	result := u.value / o.value
	return NewU16(result)
}

func (u *U16) Bn() uint64 {
	return uint64(u.value)
}

func (u *U16) N() uint32 {
	return uint32(u.value)
}

func (u *U16) ToBEBuf() []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, u.value)
	return buf
}

func (u *U16) ToHex() string {
	return hex.EncodeToString(u.ToBEBuf())
}

func FromBEBufU16(buf []byte) (*U16, error) {
	if len(buf) != 2 {
		return nil, ErrInvalidSize
	}
	return NewU16(binary.BigEndian.Uint16(buf))
}

func FromHexU16(hexStr string) (*U16, error) {
	buf, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return FromBEBufU16(buf)
}

// U32 represents a 32-bit unsigned integer.
type U32 struct {
	value uint32
}

func NewU32(value uint32) (*U32, error) {
	if value < 0x00000000 || value > 0xffffffff {
		return nil, ErrInvalidValue
	}
	return &U32{value: value}, nil
}

func (u *U32) Add(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U32)
	if !ok {
		return nil, fmt.Errorf("invalid type for addition")
	}
	result := u.value + o.value
	return NewU32(result)
}

func (u *U32) Sub(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U32)
	if !ok {
		return nil, fmt.Errorf("invalid type for subtraction")
	}
	result := u.value - o.value
	return NewU32(result)
}

func (u *U32) Mul(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U32)
	if !ok {
		return nil, fmt.Errorf("invalid type for multiplication")
	}
	result := u.value * o.value
	return NewU32(result)
}

func (u *U32) Div(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U32)
	if !ok {
		return nil, fmt.Errorf("invalid type for division")
	}
	if o.value == 0 {
		return nil, fmt.Errorf("division by zero")
	}
	result := u.value / o.value
	return NewU32(result)
}

func (u *U32) Bn() uint64 {
	return uint64(u.value)
}

func (u *U32) N() uint32 {
	return u.value
}

func (u *U32) ToBEBuf() []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, u.value)
	return buf
}

func (u *U32) ToHex() string {
	return hex.EncodeToString(u.ToBEBuf())
}

func FromBEBufU32(buf []byte) (*U32, error) {
	if len(buf) != 4 {
		return nil, ErrInvalidSize
	}
	return NewU32(binary.BigEndian.Uint32(buf))
}

func FromHexU32(hexStr string) (*U32, error) {
	buf, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return FromBEBufU32(buf)
}

// U64 represents a 64-bit unsigned integer.
type U64 struct {
	value uint64
}

func NewU64(value uint64) (*U64, error) {
	if value < 0x0000000000000000 || value > 0xffffffffffffffff {
		return nil, ErrInvalidValue
	}
	return &U64{value: value}, nil
}

func (u *U64) Add(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U64)
	if !ok {
		return nil, fmt.Errorf("invalid type for addition")
	}
	result := u.value + o.value
	return NewU64(result)
}

func (u *U64) Sub(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U64)
	if !ok {
		return nil, fmt.Errorf("invalid type for subtraction")
	}
	result := u.value - o.value
	return NewU64(result)
}

func (u *U64) Mul(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U64)
	if !ok {
		return nil, fmt.Errorf("invalid type for multiplication")
	}
	result := u.value * o.value
	return NewU64(result)
}

func (u *U64) Div(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U64)
	if !ok {
		return nil, fmt.Errorf("invalid type for division")
	}
	if o.value == 0 {
		return nil, fmt.Errorf("division by zero")
	}
	result := u.value / o.value
	return NewU64(result)
}

func (u *U64) Bn() uint64 {
	return u.value
}

func (u *U64) N() uint32 {
	return uint32(u.value)
}

func (u *U64) ToBEBuf() []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, u.value)
	return buf
}

func (u *U64) ToHex() string {
	return hex.EncodeToString(u.ToBEBuf())
}

func FromBEBufU64(buf []byte) (*U64, error) {
	if len(buf) != 8 {
		return nil, ErrInvalidSize
	}
	return NewU64(binary.BigEndian.Uint64(buf))
}

func FromHexU64(hexStr string) (*U64, error) {
	buf, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return FromBEBufU64(buf)
}

// U128 represents a 128-bit unsigned integer.
type U128 struct {
	value uint128
}

type uint128 struct {
	high, low uint64
}

func NewU128(value uint128) (*U128, error) {
	// Since Go's uint128 is not built-in, we simulate it
	if value.high > 0xffffffffffffffff || value.low > 0xffffffffffffffff {
		return nil, ErrInvalidValue
	}
	return &U128{value: value}, nil
}

func (u *U128) Add(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U128)
	if !ok {
		return nil, fmt.Errorf("invalid type for addition")
	}
	// Implement addition for uint128
	resultHigh, resultLow := u.value.high+o.value.high, u.value.low+o.value.low
	if resultLow < u.value.low {
		resultHigh++
	}
	return NewU128(uint128{high: resultHigh, low: resultLow})
}

func (u *U128) Sub(other BasicNumber) (BasicNumber, error) {
	o, ok := other.(*U128)
	if !ok {
		return nil, fmt.Errorf("invalid type for subtraction")
	}
	// Implement subtraction for uint128
	resultHigh, resultLow := u.value.high-o.value.high, u.value.low-o.value.low
	if resultLow > u.value.low {
		resultHigh--
	}
	return NewU128(uint128{high: resultHigh, low: resultLow})
}

func (u *U128) Mul(other BasicNumber) (BasicNumber, error) {
	_, ok := other.(*U128)
	if !ok {
		return nil, fmt.Errorf("invalid type for multiplication")
	}
	// Implement multiplication for uint128
	// This will need a more complex implementation in a real scenario
	return nil, fmt.Errorf("multiplication not implemented for uint128")
}

func (u *U128) Div(other BasicNumber) (BasicNumber, error) {
	_, ok := other.(*U128)
	if !ok {
		return nil, fmt.Errorf("invalid type for division")
	}
	// Implement division for uint128
	// This will need a more complex implementation in a real scenario
	return nil, fmt.Errorf("division not implemented for uint128")
}

func (u *U128) Bn() uint64 {
	return u.value.low
}

func (u *U128) N() uint32 {
	return uint32(u.value.low)
}

func (u *U128) ToBEBuf() []byte {
	buf := make([]byte, 16)
	binary.BigEndian.PutUint64(buf[:8], u.value.high)
	binary.BigEndian.PutUint64(buf[8:], u.value.low)
	return buf
}

func (u *U128) ToHex() string {
	return hex.EncodeToString(u.ToBEBuf())
}

func FromBEBufU128(buf []byte) (*U128, error) {
	if len(buf) != 16 {
		return nil, ErrInvalidSize
	}
	high := binary.BigEndian.Uint64(buf[:8])
	low := binary.BigEndian.Uint64(buf[8:])
	return NewU128(uint128{high: high, low: low})
}

func FromHexU128(hexStr string) (*U128, error) {
	buf, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return FromBEBufU128(buf)
}

// U256 represents a 256-bit unsigned integer.
type U256 struct {
	value *big.Int
}

// NewU256 creates a new U256 instance from a *big.Int value.
func NewU256(value *big.Int) *U256 {
	return &U256{value: value}
}

// NewU256FromUint64 creates a U256 instance from a uint64.
func NewU256FromUint64(value uint64) *U256 {
	return &U256{value: new(big.Int).SetUint64(value)}
}

// Add adds another U256 to the current U256 and returns the result.
func (u *U256) Add(other *U256) *U256 {
	result := new(big.Int).Add(u.value, other.value)
	return &U256{value: result}
}

// Sub subtracts another U256 from the current U256 and returns the result.
func (u *U256) Sub(other *U256) *U256 {
	result := new(big.Int).Sub(u.value, other.value)
	return &U256{value: result}
}

// Mul multiplies another U256 with the current U256 and returns the result.
func (u *U256) Mul(other *U256) *U256 {
	result := new(big.Int).Mul(u.value, other.value)
	return &U256{value: result}
}

// Div divides the current U256 by another U256 and returns the result.
func (u *U256) Div(other *U256) (*U256, error) {
	if other.value.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("division by zero")
	}
	result := new(big.Int).Div(u.value, other.value)
	return &U256{value: result}, nil
}

// ToBEBuf converts the U256 value to a big-endian byte buffer.
func (u *U256) ToBEBuf() []byte {
	buf := u.value.Bytes()
	// Ensure buffer length is 32 bytes
	paddedBuf := make([]byte, 32)
	copy(paddedBuf[32-len(buf):], buf)
	return paddedBuf
}

// ToHex converts the U256 value to a hexadecimal string.
func (u *U256) ToHex() string {
	return hex.EncodeToString(u.ToBEBuf())
}

// FromBEBuf creates a U256 from a big-endian byte buffer.
func FromBEBufU256(buf []byte) (*U256, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("buffer length must be 32 bytes")
	}
	value := new(big.Int).SetBytes(buf)
	return &U256{value: value}, nil
}

// FromHex creates a U256 from a hexadecimal string.
func FromHexU256(hexStr string) (*U256, error) {
	buf, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid hex string length for U256")
	}
	return FromBEBufU256(buf)
}