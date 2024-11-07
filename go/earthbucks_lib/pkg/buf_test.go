package earthbucks

import (
    "encoding/hex"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestEbxBuf(t *testing.T) {
    t.Run("to/from buf", func(t *testing.T) {
        buf, _ := hex.DecodeString("deadbeef")
        n := len(buf)
        isoBuf,_ := EbxBufFromBuf(&n, &buf)
        assert.IsType(t, &EbxBuf{}, isoBuf)
        assert.Equal(t, "deadbeef", hex.EncodeToString(*isoBuf.buf))
    })

    t.Run("to/from base58", func(t *testing.T) {
        buf, _ := hex.DecodeString("deadbeef")
        n := len(buf)
        isoBuf,_ := EbxBufFromBuf(&n, &buf)        
        assert.IsType(t, &EbxBuf{}, isoBuf)
        base58 := isoBuf.ToBase58()
        isoBuf2,_ := EbxBufFromBase58(&n, &base58)
        assert.Equal(t, "deadbeef", hex.EncodeToString(*isoBuf2.buf))
    })
}

func TestFixedEbxBuf(t *testing.T) {
    t.Run("to/from buf", func(t *testing.T) {
        buf, _ := hex.DecodeString("deadbeef")
        n := 4
        fixedEbxBuf, _ := FixedBufFromBuf(&n, &buf)
        assert.IsType(t, &FixedBuf{}, fixedEbxBuf)
        assert.Equal(t, "deadbeef", hex.EncodeToString(*fixedEbxBuf.buf))
    })

    t.Run("to/from base58", func(t *testing.T) {
        buf, _ := hex.DecodeString("deadbeef")
        n := 4
        fixedEbxBuf, _ := NewFixedBuf(&n, &buf)
        assert.IsType(t, &FixedBuf{}, fixedEbxBuf)
        base58 := fixedEbxBuf.ToBase58()
        fixedEbxBuf2, _ := FixedBufFromBase58(&n, &base58)
        assert.Equal(t, "deadbeef", hex.EncodeToString(*fixedEbxBuf2.buf))
    })
}
