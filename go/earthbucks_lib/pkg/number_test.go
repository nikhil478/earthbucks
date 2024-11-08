package earthbucks

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumbers(t *testing.T) {
    t.Run("U8", func(t *testing.T) {
        a, err := NewU8(*big.NewInt(10))
        if err != nil {
            t.Fatalf("error creating U8: %v", err)
        }
        b, err := NewU8(*big.NewInt(20))
        if err != nil {
            t.Fatalf("error creating U8: %v", err)
        }
        sum, err := a.Add(b)
        if err != nil {
            t.Fatalf("error adding U8: %v", err)
        }
        mul, err := a.Mul(b)
        if err != nil {
            t.Fatalf("error multiplying U8: %v", err)
        }
        assert.Equal(t, big.NewInt(30).String(), sum.Bn().String())
        assert.Equal(t, big.NewInt(200).String(), mul.Bn().String())
    })

    t.Run("U16", func(t *testing.T) {
        a, err := NewU16(*big.NewInt(10))
        if err != nil {
            t.Fatalf("error creating U16: %v", err)
        }
        b, err := NewU16(*big.NewInt(20))
        if err != nil {
            t.Fatalf("error creating U16: %v", err)
        }
        sum, err := a.Add(b)
        if err != nil {
            t.Fatalf("error adding U16: %v", err)
        }
        mul, err := a.Mul(b)
        if err != nil {
            t.Fatalf("error multiplying U16: %v", err)
        }
        assert.Equal(t, big.NewInt(30).String(), sum.Bn().String())
        assert.Equal(t, big.NewInt(200).String(), mul.Bn().String())
    })

    t.Run("U32", func(t *testing.T) {
        a, err := NewU32(*big.NewInt(10))
        if err != nil {
            t.Fatalf("error creating U32: %v", err)
        }
        b, err := NewU32(*big.NewInt(20))
        if err != nil {
            t.Fatalf("error creating U32: %v", err)
        }
        sum, err := a.Add(b)
        if err != nil {
            t.Fatalf("error adding U32: %v", err)
        }
        mul, err := a.Mul(b)
        if err != nil {
            t.Fatalf("error multiplying U32: %v", err)
        }
        assert.Equal(t, big.NewInt(30).String(), sum.Bn().String())
        assert.Equal(t, big.NewInt(200).String(), mul.Bn().String())
    })

    t.Run("U64", func(t *testing.T) {
        a, err := NewU64(*big.NewInt(10))
        if err != nil {
            t.Fatalf("error creating U64: %v", err)
        }
        b, err := NewU64(*big.NewInt(20))
        if err != nil {
            t.Fatalf("error creating U64: %v", err)
        }
        sum, err := a.Add(b)
        if err != nil {
            t.Fatalf("error adding U64: %v", err)
        }
        mul, err := a.Mul(b)
        if err != nil {
            t.Fatalf("error multiplying U64: %v", err)
        }
        assert.Equal(t, big.NewInt(30).String(), sum.Bn().String())
        assert.Equal(t, big.NewInt(200).String(), mul.Bn().String())
    })

    t.Run("U128", func(t *testing.T) {
        a, err := NewU128(*big.NewInt(10))
        if err != nil {
            t.Fatalf("error creating U128: %v", err)
        }
        b, err := NewU128(*big.NewInt(20))
        if err != nil {
            t.Fatalf("error creating U128: %v", err)
        }
        sum, err := a.Add(b)
        if err != nil {
            t.Fatalf("error adding U128: %v", err)
        }
        mul, err := a.Mul(b)
        if err != nil {
            t.Fatalf("error multiplying U128: %v", err)
        }
        assert.Equal(t, big.NewInt(30).String(), sum.value.String())
        assert.Equal(t, big.NewInt(200).String(), mul.value.String())
    })

    t.Run("U256", func(t *testing.T) {
        a, _ := NewU256(*big.NewInt(10))
        b, _ := NewU256(*big.NewInt(20))
        sum, _ := a.Add(b)
        mul, _:= a.Mul(b)
        assert.Equal(t, big.NewInt(30).String(), sum.value.String())
        assert.Equal(t, big.NewInt(200).String(), mul.value.String())
    })
}