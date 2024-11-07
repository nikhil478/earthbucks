package earthbucks

type VarInt struct {
   buf *[]byte
} 

func NewVarInt(buf *[]byte) *VarInt {
	if buf == nil {
		defaultBuf := make([]byte, 0)
		buf = &defaultBuf             
	}
	return &VarInt{
		buf: buf,
	}
}

func (vi *VarInt) ToBuf() *[]byte {
	return vi.buf
} 

func (vi *VarInt) ToU64() (*U64, error) {
	return NewBufReader(*vi.buf).ReadVarInt()
}

func (vi *VarInt) ToU32() (*U32, error) {
	u64, err := NewBufReader(*vi.buf).ReadVarInt()
	if err != nil {
		return nil, err
	}
	return &U32{
		value: uint32(u64.value),
	}, nil
}

func (vi *VarInt) isMinimal() bool {

	u64, err := vi.ToU64()
	if err != nil {
		return false
	}

	// Create a new VarInt from the uint64
	varInt := VarIntFromU64(u64)

	// Compare the original buffer and the buffer from varInt
	return equal(u64, varInt.ToBuf()) == 0
}


func VarIntFromU64(u64 *U64) *VarInt {
	buf := NewBufWriter().WriteVarInt(u64).ToBuf();
	return NewVarInt(buf)
}

func VarIntFromU32(u32 *U32) *VarInt {
	buf := NewBufWriter().WriteVarInt(NewU64(uint64(u32.value))).ToBuf()
}

func VarIntFromBufReader(br *BufReader) (*VarInt, error) {
	buf, err := br.ReadVarIntBuf()
	if err != nil {
		return nil, err
	}
	return NewVarInt(&buf), nil
}

