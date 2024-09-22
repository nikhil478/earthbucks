package earthbucks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBufWriter(t *testing.T) {
    var bufferWriter *BufWriter

    // Before each test, initialize a new BufWriter
    beforeEach := func() {
        bufferWriter = NewBufWriter()
    }

    t.Run("writeUInt8", func(t *testing.T) {
        beforeEach()
        u8, _ := NewU8(123)
        bufferWriter.WriteU8BE(u8)
        result := bufferWriter.ToBuf()
        assert.Equal(t, u8.value, result[0])
    })

    t.Run("writeUInt16BE", func(t *testing.T) {
        beforeEach()
        u16,_ := NewU16(12345)
        bufferWriter.WriteU16BE(u16)
        result := bufferWriter.ToBuf()
        assert.Equal(t, u16.value, uint16(result[0])<<8|uint16(result[1]))
    })

    t.Run("writeUInt32BE", func(t *testing.T) {
        beforeEach()
        u32,_ := NewU32(1234567890)
        bufferWriter.WriteU32BE(u32)
        result := bufferWriter.ToBuf()
        assert.Equal(t, u32.value, uint32(result[0])<<24|uint32(result[1])<<16|uint32(result[2])<<8|uint32(result[3]))
    })

    t.Run("writeUInt64BEBn", func(t *testing.T) {
        beforeEach()
        u64,_ := NewU64(1234567890123456789)
        bufferWriter.WriteU64BE(u64)
        result := bufferWriter.ToBuf()
        assert.Equal(t, u64.value, uint64(result[0])<<56|uint64(result[1])<<48|uint64(result[2])<<40|uint64(result[3])<<32|uint64(result[4])<<24|uint64(result[5])<<16|uint64(result[6])<<8|uint64(result[7]))
    })

    t.Run("writeVarInt", func(t *testing.T) {
        beforeEach()
        bn,_ := NewU64(1234567890123456789)
        bufferWriter.WriteVarInt(bn)
        result := bufferWriter.ToBuf()
        assert.Equal(t, "ff112210f47de98115", bytesToHex(result))
    })

	t.Run("varIntBufBigInt", func(t *testing.T) {
		testCases := []struct {
			input    *U64
			expected string
		}{
			{func() *U64 { v, _ := NewU64(252); return v }(), "fcfc"},
			{func() *U64 { v, _ := NewU64(0xffff); return v }(), "fdffff"},
			{func() *U64 { v, _ := NewU64(0xffffffff); return v }(), "feffffffff"},
			{func() *U64 { v, _ := NewU64(0x100000000); return v }(), "ffff0000000000000000"},
		}
	
		for _, tc := range testCases {
			t.Run(tc.expected, func(t *testing.T) {
				beforeEach()
				result := bufferWriter.VarIntBuf(tc.input)
				fmt.Print(bytesToHex(result))
				assert.Equal(t, tc.expected, bytesToHex(result))
			})
		}
	})
}

func bytesToHex(b []byte) string {
    hexStr := make([]byte, len(b)*2)
    hexChars := "0123456789abcdef"
    
    for i, v := range b {
        hexStr[i*2] = hexChars[v>>4]
        hexStr[i*2+1] = hexChars[v&0xf]
    }
    
    return string(hexStr)
}

