package earthbucks

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestBufWriter(t *testing.T) {
	

	t.Run("writeUInt8", func(t *testing.T) {
        bufferWriter := NewBufWriter()
		u8 := U8{value: 123}
		bufferWriter.WriteU8BE(&u8)
		result := bufferWriter.ToBuf()
		if result[0] != u8.value {
			t.Errorf("Expected %d, got %d", u8.value, result[0])
		}
	})

	t.Run("writeUInt16BE", func(t *testing.T) {
        bufferWriter := NewBufWriter()
		u16 := U16{value: 12345}
		bufferWriter.WriteU16BE(&u16)
		result := bufferWriter.ToBuf()
		if got := binary.BigEndian.Uint16(result); got != u16.value {
			t.Errorf("Expected %d, got %d", u16.value, got)
		}
	})

	t.Run("writeUInt32BE", func(t *testing.T) {
        bufferWriter := NewBufWriter()
		u32 := U32{value: 1234567890}
		bufferWriter.WriteU32BE(&u32)
		result := bufferWriter.ToBuf()
		if got := binary.BigEndian.Uint32(result); got != u32.value {
			t.Errorf("Expected %d, got %d", u32.value, got)
		}
	})

	t.Run("writeUInt64BE", func(t *testing.T) {
        bufferWriter := NewBufWriter()
		u64 := U64{value: 1234567890123456789}
		bufferWriter.WriteU64BE(&u64)
		result := bufferWriter.ToBuf()
		if got := binary.BigEndian.Uint64(result); got != u64.value {
			t.Errorf("Expected %d, got %d", u64.value, got)
		}
	})

	t.Run("writeVarInt", func(t *testing.T) {
        bufferWriter := NewBufWriter()
		bn := U64{value: 1234567890123456789}
		WriteVarInt(&bn, bufferWriter)
		result := bufferWriter.ToBuf()
		expectedHex := "ff112210f47de98115"
		if hexString := bytesToHex(result); hexString != expectedHex {
			t.Errorf("Expected %s, got %s", expectedHex, hexString)
		}
	})

	t.Run("varIntBufBigInt", func(t *testing.T) {
		t.Run("should write a bigint less than 253 as a single byte", func(t *testing.T) {
			bn := U64{value: 252}
			result,_ := VarIntBuf(&bn)
			if result[0] != byte(bn.value) {
				t.Errorf("Expected %d, got %d", bn.value, result[0])
			}
		})

		t.Run("should write a bigint less than 0x10000 as a 3-byte integer", func(t *testing.T) {
			bn := U64{value: 0xffff}
			result,_ := VarIntBuf(&bn)
			if result[0] != 253 {
				t.Errorf("Expected 253, got %d", result[0])
			}
			if (result[1]<<8|result[2]) != byte(bn.value) {
				t.Errorf("Expected %d, got %d", bn.value, (result[1]<<8|result[2]))
			}
		})

		t.Run("should write a bigint less than 0x100000000 as a 5-byte integer", func(t *testing.T) {
			bn := U64{value: 0xffffffff}
			result,_ := VarIntBuf(&bn)
			if result[0] != 254 {
				t.Errorf("Expected 254, got %d", result[0])
			}
			expectedHex := "feffffffff"
			if hexString := bytesToHex(result); hexString != expectedHex {
				t.Errorf("Expected %s, got %s", expectedHex, hexString)
			}
		})

		t.Run("should write a bigint greater than or equal to 0x100000000 as a 9-byte integer", func(t *testing.T) {
			u64 := U64{value: 0x100000000}
			result,_ := VarIntBuf(&u64)
			if result[0] != 255 {
				t.Errorf("Expected 255, got %d", result[0])
			}
			readBn := (uint64(result[1]) << 56) |
				(uint64(result[2]) << 48) |
				(uint64(result[3]) << 40) |
				(uint64(result[4]) << 32) |
				(uint64(result[5]) << 24) |
				(uint64(result[6]) << 16) |
				(uint64(result[7]) << 8) |
				uint64(result[8])
			if readBn != u64.value {
				t.Errorf("Expected %d, got %d", u64.value, readBn)
			}
		})
	})
}

func bytesToHex(b []byte) string {
	return fmt.Sprintf("%x", b)
}

