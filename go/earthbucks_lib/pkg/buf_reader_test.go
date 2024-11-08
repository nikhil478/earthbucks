package earthbucks

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestVectorEbxBufReader struct {
	Read           TestVectorReadEbxBuf        `json:"read"`
	ReadU8         TestVectorReadErrors       `json:"read_u8"`
	ReadU16BE      TestVectorReadErrors       `json:"read_u16_be"`
	ReadU32BE      TestVectorReadErrors       `json:"read_u32_be"`
	ReadU64BE      TestVectorReadErrors       `json:"read_u64_be"`
	ReadVarIntBuf  TestVectorReadErrors       `json:"read_var_int_buf"`
	ReadVarInt     TestVectorReadErrors       `json:"read_var_int"`
}

type TestVectorReadEbxBuf struct {
	Errors []TestVectorReadEbxBufError `json:"errors"`
}

type TestVectorReadEbxBufError struct {
	Hex   string `json:"hex"`
	Len   int    `json:"len"`
	Error string `json:"error"`
}

type TestVectorReadErrors struct {
	Errors []TestVectorReadError `json:"errors"`
}

type TestVectorReadError struct {
	Hex   string `json:"hex"`
	Error string `json:"error"`
}

// TestBufReader tests basic functionality of the BufReader
func TestBufReader(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	bufferReader := NewBufReader(&data)

	t.Run("Constructor", func(t *testing.T) {
		assert.True(t, bytes.Equal(*bufferReader.buf, data), "buffer mismatch")
		assert.Equal(t, 0, bufferReader.pos, "initial position should be 0")
	})

	t.Run("Read", func(t *testing.T) {
		num := 4
		subarray, err := bufferReader.Read(&num)
		assert.NoError(t, err, "Read should not return an error")
		assert.True(t, bytes.Equal(*subarray, data[:4]), "subarray mismatch")
	})

	t.Run("PositionUpdate", func(t *testing.T) {
		assert.Equal(t, 4, bufferReader.pos, "position should be updated to 4")
	})

	t.Run("ReadU8", func(t *testing.T) {
		val := data[4:]
		bufferReader = NewBufReader(&val)
		value, err := bufferReader.ReadU8()
		assert.NoError(t, err, "ReadU8 should not return an error")
		assert.Equal(t, uint8(5), value.value, "ReadU8 value mismatch")
	})

	t.Run("ReadU16BE", func(t *testing.T) {
		bufferReader = NewBufReader(&[]byte{0x01, 0x02})
		value, err := bufferReader.ReadU16BE()
		assert.NoError(t, err, "ReadU16BE should not return an error")
		assert.Equal(t, uint16(0x0102), value.value, "ReadU16BE value mismatch")
	})

	t.Run("ReadU32BE", func(t *testing.T) {
		bufferReader = NewBufReader(&[]byte{0x01, 0x02, 0x03, 0x04})
		value, err := bufferReader.ReadU32BE()
		assert.NoError(t, err, "ReadU32BE should not return an error")
		assert.Equal(t, uint32(0x01020304), value.value, "ReadU32BE value mismatch")
	})

	t.Run("ReadU64BE", func(t *testing.T) {
		bufferReader = NewBufReader(&[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef})
		value, err := bufferReader.ReadU64BE()
		assert.NoError(t, err, "ReadU64BE should not return an error")
		assert.Equal(t, uint64(0x0123456789ABCDEF), value.value, "ReadU64BE value mismatch")
	})

	t.Run("ReadVarIntBuf", func(t *testing.T) {
		bufferReader = NewBufReader(&[]byte{0x01})
		buf, err := bufferReader.ReadVarIntBuf()
		assert.NoError(t, err, "ReadVarIntBuf should not return an error")
		assert.True(t, bytes.Equal(buf, []byte{0x01}), "ReadVarIntBuf value mismatch")
	})

	t.Run("ReadVarInt", func(t *testing.T) {
		bufferReader = NewBufReader(&[]byte{0x01})
		value, err := bufferReader.ReadVarInt()
		assert.NoError(t, err, "ReadVarInt should not return an error")
		assert.Equal(t, uint64(1), value.value, "ReadVarInt value mismatch")
	})
}


func TestBufReaderTestVectors(t *testing.T) {
	filePath := "../test-vectors/buf_reader.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read test vector file: %v", err)
	}

	var testVector TestVectorEbxBufReader
	if err := json.Unmarshal(data, &testVector); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Helper function to parse hex string to bytes
	parseHex := func(hexStr string) []byte {
		data := make([]byte, len(hexStr)/2)
		for i := 0; i < len(hexStr); i += 2 {
			fmt.Sscanf(hexStr[i:i+2], "%2x", &data[i/2])
		}
		return data
	}

	// Helper function to map error strings to errors
	mapErrorType := func(errorStr string) error {
		switch errorStr {
		case "non-minimal encoding":
			return ErrNonMinimalEncoding
		case "not enough bytes in the buffer to read":
			return ErrNotEnoughData
		default:
			return errors.New("generic error")
		}
	}

	t.Run("test vectors: read", func(t *testing.T) {
		for _, test := range testVector.Read.Errors {
			buf := parseHex(test.Hex)
			bufferReader := NewBufReader(&buf)
			expectedErr := mapErrorType(test.Error)
			_, err := bufferReader.Read(&test.Len)
			assert.True(t, errors.Is(err, expectedErr), "error mismatch")
		}
	})

	t.Run("test vectors: read_u8", func(t *testing.T) {
		for _, test := range testVector.ReadU8.Errors {
			buf := parseHex(test.Hex)
			bufferReader := NewBufReader(&buf)
			expectedErr := mapErrorType(test.Error)
			_, err := bufferReader.ReadU8()
			assert.True(t, errors.Is(err, expectedErr), "error mismatch")
		}
	})

	t.Run("test vectors: read_u16_be", func(t *testing.T) {
		for _, test := range testVector.ReadU16BE.Errors {
			buf := parseHex(test.Hex)
			bufferReader := NewBufReader(&buf)
			expectedErr := mapErrorType(test.Error)
			_, err := bufferReader.ReadU16BE()
			assert.True(t, errors.Is(err, expectedErr), "error mismatch")
		}
	})

	t.Run("test vectors: read_u32_be", func(t *testing.T) {
		for _, test := range testVector.ReadU32BE.Errors {
			buf := parseHex(test.Hex)
			bufferReader := NewBufReader(&buf)
			expectedErr := mapErrorType(test.Error)
			_, err := bufferReader.ReadU32BE()
			assert.True(t, errors.Is(err, expectedErr), "error mismatch")
		}
	})

	t.Run("test vectors: read_u64_be", func(t *testing.T) {
		for _, test := range testVector.ReadU64BE.Errors {
			buf := parseHex(test.Hex)
			bufferReader := NewBufReader(&buf)
			expectedErr := mapErrorType(test.Error)
			_, err := bufferReader.ReadU64BE()
			assert.True(t, errors.Is(err, expectedErr), "error mismatch")
		}
	})

	t.Run("test vectors: read_var_int_buf", func(t *testing.T) {
		for _, test := range testVector.ReadVarIntBuf.Errors {
			buf := parseHex(test.Hex)
			bufferReader := NewBufReader(&buf)
			expectedErr := mapErrorType(test.Error)
			_, err := bufferReader.ReadVarIntBuf()
			assert.True(t, errors.Is(err, expectedErr), "error mismatch")
		}
	})

	t.Run("test vectors: read_var_int", func(t *testing.T) {
		for _, test := range testVector.ReadVarInt.Errors {
			buf := parseHex(test.Hex)
			bufferReader := NewBufReader(&buf)
			expectedErr := mapErrorType(test.Error)
			_, err := bufferReader.ReadVarInt()
			assert.True(t, errors.Is(err, expectedErr), "error mismatch")
		}
	})
}

