package earthbucks

import (
    "encoding/hex"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestEbxBuf(t *testing.T) {
    t.Run("to/from buf", func(t *testing.T) {
        buf, _ := hex.DecodeString("deadbeef")
        isoBuf,_ := EbxBufFromBuf(len(buf), buf)

        assert.IsType(t, &EbxBuf{}, isoBuf)
        assert.Equal(t, "deadbeef", hex.EncodeToString(isoBuf.buf))
    })

    t.Run("to/from base58", func(t *testing.T) {
        buf, _ := hex.DecodeString("deadbeef")
        isoBuf,_ := EbxBufFromBuf(len(buf), buf)
        
        assert.IsType(t, &EbxBuf{}, isoBuf)
        
        base58 := isoBuf.ToBase58()
        isoBuf2,_ := EbxBufFromBase58(len(buf), base58)
        assert.Equal(t, "deadbeef", hex.EncodeToString(isoBuf2.buf))
    })
}

func TestFixedEbxBuf(t *testing.T) {
    t.Run("to/from buf", func(t *testing.T) {
        buf, _ := hex.DecodeString("deadbeef")
        fixedEbxBuf, _ := FixedBufFromBuf(4, buf)
        assert.IsType(t, &FixedBuf{}, fixedEbxBuf)
        assert.Equal(t, "deadbeef", hex.EncodeToString(fixedEbxBuf.buf))
    })

    t.Run("to/from base58", func(t *testing.T) {
        buf, _ := hex.DecodeString("deadbeef")
        fixedEbxBuf, _ := NewFixedBuf(4, buf)

        assert.IsType(t, &FixedBuf{}, fixedEbxBuf)
        
        base58 := fixedEbxBuf.ToBase58()
        fixedEbxBuf2, _ := FixedBufFromBase58(4, base58)
        assert.Equal(t, "deadbeef", hex.EncodeToString(fixedEbxBuf2.buf))
    })
}
