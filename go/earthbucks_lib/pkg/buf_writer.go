package earthbucks

import (
	"bytes"
	"encoding/binary"
)

type BufWriter struct {
	bufs [][]byte
}

func NewBufWriter(bufs ...[]byte) *BufWriter {
	w := &BufWriter{}
	for _, buf := range bufs {
		w.Write(&buf)
	}
	return w
}

func (w *BufWriter) GetLength() int {
	length := 0
	for _, buf := range w.bufs {
		length += len(buf)
	}
	return length
}

func (w *BufWriter) ToBuf() []byte {
	return bytes.Join(w.bufs, nil)
}

func (w *BufWriter) Write(buf *[]byte) *BufWriter {
	w.bufs = append(w.bufs, *buf)
	return w
}

func (w *BufWriter) WriteU8(u8 *U8) *BufWriter {
	buf := u8.ToBEBuf()
	w.Write(&buf)
	return w
}

func (w *BufWriter) WriteU16BE(u16 *U16) *BufWriter {
	buf := u16.ToBEBuf()
	w.Write(&buf)
	return w
}

func (w *BufWriter) WriteU32BE(u32 *U32) *BufWriter {
	buf := u32.ToBEBuf()
	w.Write(&buf)
	return w
}

func (w *BufWriter) WriteU64BE(u64 *U64) *BufWriter {
	buf := u64.ToBEBuf()
	w.Write(&buf)
	return w
}

func (w *BufWriter) WriteU128BE(u128 *U128) *BufWriter {
	buf := u128.ToBEBuf()
	w.Write(&buf)
	return w
}

func (w *BufWriter) Write256BE(u256 *U256) *BufWriter {
	buf := u256.ToBEBuf()
	w.Write(&buf)
	return w
}

func (w *BufWriter) WriteVarInt(u64 *U64) *BufWriter {
	buf, _ := VarIntBuf(u64)
	w.Write(&buf)
	return w
}

func VarIntBuf(bn *U64) ([]byte, error) {
	n,_ := bn.N()
	var buf []byte

	if n < 253 {
		buf = make([]byte, 1)
		buf[0] = byte(n)
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
		binary.BigEndian.PutUint64(buf[1:], bn.value.Uint64())
	}

	return buf, nil
}

