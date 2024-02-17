package mathy

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func TestBN(t *testing.T) {
	times := 1_000_000_000
	groupNum := times / 10

	// test float64
	t.Logf("test BN with float64")
	fs := RandFloats(-1_000_000_000_000_000, 1_000_000_000_000_000, times)
	for i, f := range fs {
		if (i+1)%groupNum == 0 {
			t.Log((1+1)/groupNum, "finished")
		}
		sf := strconv.FormatFloat(f, 'f', -1, 64)
		bn := BN(f)
		if bn.String() != sf {
			t.Error("BN", bn.String(), "!=", sf)
			t.FailNow()
		}
	}

	// test int64
	t.Logf("test BN with int64")
	for i := 0; i < times; i++ {
		if (i+1)%groupNum == 0 {
			t.Log((i+1)/groupNum, "finished")
		}
		num := rand.Int63()
		snum := strconv.FormatInt(num, 10)
		bn := BN(num)
		sbn := bn.String()
		if sbn != snum {
			t.Error("BN", bn.String(), "!=", snum)
			t.FailNow()
		}
	}

	// test uint64
	t.Logf("test BN with uint64")
	for i := 0; i < times; i++ {
		if (i+1)%groupNum == 0 {
			t.Log((i+1)/groupNum, "finished")
		}
		num := rand.Int63()
		snum := strconv.FormatInt(num, 10)
		bn := BN(uint64(num))
		sbn := bn.String()
		if sbn != snum {
			t.Error("BN", bn.String(), "!=", snum)
			t.FailNow()
		}
	}

	// test string
	t.Logf("test BN with string")
	ss := RandFloats(-1_000_000_000_000_000, 1_000_000_000_000_000, times)
	for i, f := range ss {
		if (i+1)%groupNum == 0 {
			t.Log((i+1)/groupNum, "finished")
		}
		sf := strconv.FormatFloat(f, 'f', -1, 64)
		ip := strconv.FormatInt(rand.Int63(), 10)
		dp := strconv.FormatInt(rand.Int63(), 10)
		sf = ip + sf + dp
		if strings.Contains(sf, "-") {
			sf = strings.ReplaceAll(sf, "-", "")
			sf = "-" + sf
		}
		sf = strings.TrimRight(sf, "0")
		bn := BN(sf)
		if bn.String() != sf {
			t.Error("BN", bn.String(), "!=", sf)
			t.FailNow()
		}
	}
}
