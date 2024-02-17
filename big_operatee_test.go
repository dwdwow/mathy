package mathy

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestFloat_BN(t *testing.T) {
	times := 100_000_000
	t.Log("testing Float.BN", "total times", times)
	for i := 0; i < times; i++ {
		f := RandFloat(-1_000_000_000_000_000, 1_000_000_000_000_000)
		bn := Float(f).BN()
		if bn.Float64() != f {
			t.Error("Float", f, bn)
			t.FailNow()
		}
	}
}

func TestString_BN(t *testing.T) {
	times := 100_000_000
	t.Log("testing String.BN", "total times", times)
	for i := 0; i < times; i++ {
		s := strconv.FormatFloat(RandFloat(-1_000_000_000_000_000, 1_000_000_000_000_000), 'f', -1, 64)
		bn := String(s).BN()
		if bn.String() != s {
			t.Error("String", s, bn)
			t.FailNow()
		}
	}
}

func TestInt_BN(t *testing.T) {
	times := 100_000_000
	t.Log("testing Int.BN", "total times", times)
	for i := 0; i < times; i++ {
		integer := rand.Int63()
		bn := Int(integer).BN()
		if bn.String() != strconv.FormatInt(integer, 10) {
			t.Error("Int", integer, bn)
			t.FailNow()
		}
	}
}

func TestUint_BN(t *testing.T) {
	times := 100_000_000
	t.Log("testing Uint.BN", "total times", times)
	for i := 0; i < times; i++ {
		integer := rand.Uint64()
		bn := Uint(integer).BN()
		if bn.String() != strconv.FormatUint(integer, 10) {
			t.Error("Uint", integer, bn)
			t.FailNow()
		}
	}
}
