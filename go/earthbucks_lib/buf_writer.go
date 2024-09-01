package earthbucks

import (
	"bytes"
	"encoding/binary"
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

func (w *BufWriter) WriteU8BE(u8 *U8) *BufWriter {
	w.Write(u8.ToBEBuf())
	return w
}

func (w *BufWriter) WriteU16BE(u16 *U16) *BufWriter {
	w.Write(u16.ToBEBuf())
	return w
}

func (w *BufWriter) WriteU32BE(u32 *U32) *BufWriter {
	w.Write(u32.ToBEBuf())
	return w
}


func (w *BufWriter) WriteU64BE(u64 *U64) *BufWriter {
	w.Write(u64.ToBEBuf())
	return w
}

func (w *BufWriter) WriteU128BE(u128 *U128) *BufWriter {
	w.Write(u128.ToBEBuf())
	return w
}

func (w *BufWriter) Write256BE(u256 *U256) *BufWriter {
	w.Write(u256.ToBEBuf())
	return w
}

// WriteVarInt writes a variable-length integer to the buffer.
func (w *BufWriter) WriteVarInt(u64 *U64) *BufWriter {
    buf := w.VarIntBuf(u64)
    w.Write(buf)
    return w
}

// VarIntBuf creates a buffer for a variable-length integer.
func (w *BufWriter) VarIntBuf(bn *U64) []byte {
    var buf bytes.Buffer
    n := bn.value

    if n < 253 {
        buf.WriteByte(byte(n))
		binary.Write(&buf, binary.BigEndian, uint16(n))
    } else if n < 0x10000 {
        buf.WriteByte(253)
        binary.Write(&buf, binary.BigEndian, uint16(n))
    } else if n < 0x100000000 {
        buf.WriteByte(254)
        binary.Write(&buf, binary.BigEndian, uint32(n))
    } else {
        buf.WriteByte(255)
        // For large values, write 8 bytes
        binary.Write(&buf, binary.BigEndian, uint64(n))
    }

    return buf.Bytes()
}
