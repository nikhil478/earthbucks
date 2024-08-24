package earthbucks

import (
	"bytes"
	"encoding/binary"
	"math/big"
)

// BufWriter is a buffer writer that accumulates slices of bytes.
type BufWriter struct {
	bufs [][]byte
}

// NewBufWriter creates a new BufWriter with an optional initial set of buffers.
func NewBufWriter(bufs ...[]byte) *BufWriter {
	w := &BufWriter{}
	for _, buf := range bufs {
		w.Write(buf)
	}
	return w
}

// GetLength returns the total length of all accumulated buffers.
func (w *BufWriter) GetLength() int {
	length := 0
	for _, buf := range w.bufs {
		length += len(buf)
	}
	return length
}

// ToBuf concatenates all accumulated buffers into one single byte slice.
func (w *BufWriter) ToBuf() []byte {
	return bytes.Join(w.bufs, nil)
}

// Write appends a new buffer to the accumulated buffers.
func (w *BufWriter) Write(buf []byte) *BufWriter {
	w.bufs = append(w.bufs, buf)
	return w
}

// WriteU8 writes a single byte (uint8) to the buffer.
func (w *BufWriter) WriteU8(u8 uint8) *BufWriter {
	buf := []byte{u8}
	w.Write(buf)
	return w
}

// WriteU16BE writes a 16-bit unsigned integer in big-endian order.
func (w *BufWriter) WriteU16BE(u16 uint16) *BufWriter {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, u16)
	w.Write(buf)
	return w
}

// WriteU32BE writes a 32-bit unsigned integer in big-endian order.
func (w *BufWriter) WriteU32BE(u32 uint32) *BufWriter {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, u32)
	w.Write(buf)
	return w
}

// WriteU64BE writes a 64-bit unsigned integer in big-endian order.
func (w *BufWriter) WriteU64BE(u64 uint64) *BufWriter {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, u64)
	w.Write(buf)
	return w
}

// WriteU128BE writes a 128-bit unsigned integer in big-endian order.
func (w *BufWriter) WriteU128BE(u128 *big.Int) *BufWriter {
	buf := make([]byte, 16)
	u128Bytes := u128.Bytes()
	copy(buf[16-len(u128Bytes):], u128Bytes)
	w.Write(buf)
	return w
}

// WriteVarInt writes a variable-length integer to the buffer.
func (w *BufWriter) WriteVarInt(u64 uint64) *BufWriter {
	buf := w.VarIntBuf(u64)
	w.Write(buf)
	return w
}

// VarIntBuf creates a buffer for a variable-length integer.
func (w *BufWriter) VarIntBuf(n uint64) []byte {
	var buf []byte
	if n < 253 {
		buf = []byte{byte(n)}
	} else if n < 0x10000 {
		buf = make([]byte, 3)
		buf[0] = 253
		binary.BigEndian.PutUint16(buf[1:], uint16(n))
	} else if n < 0x100000000 {
		buf = make([]byte, 5)
		buf[0] = 254
		binary.BigEndian.PutUint32(buf[1:], uint32(n))
	} else {
		buf = make([]byte, 9)
		buf[0] = 255
		binary.BigEndian.PutUint64(buf[1:], n)
	}
	return buf
}