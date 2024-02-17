package mathy

import (
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/shopspring/decimal"
)

func TestBN(t *testing.T) {
	times := 1_000_000_000
	groupNum := times / 10

	// test float64
	t.Logf("test BN with float64")
	fs := RandFloats(-1_000_000_000_000_000, 1_000_000_000_000_000, times)
	for i, f := range fs {
		if (i+1)%groupNum == 0 {
			t.Log((i+1)/groupNum, "finished")
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

func TestBig_BN(t *testing.T) {
	t.Log("test Big_BN")
	times := 1_000
	groupNum := times / 10
	for i := 0; i < times; i++ {
		f := RandFloat(-1_000_000_000_000, 1_000_000_000_000)
		bn := BN(f)
		if bn.BN() != bn {
			t.Error("BN", bn.String(), "!= Big.BN()", bn.BN().String())
			t.FailNow()
		}
		j := i + 1
		if j%groupNum == 0 {
			t.Log(j/groupNum, "finished")
		}
	}
}

func TestBig_Copy(t *testing.T) {
	t.Log("test Big_BN")
	times := 1_000
	groupNum := times / 10
	for i := 0; i < times; i++ {
		f := RandFloat(-1_000_000_000_000, 1_000_000_000_000)
		bn := BN(f)
		copied := bn.Copy()
		if copied == bn {
			t.Error("BN", bn.String(), "copy != own copy", bn.BN().String())
			t.FailNow()
		}
		if copied.String() != bn.String() {
			t.Error("BN copied", copied.String(), "!= BN", bn.String())
			t.FailNow()
		}
		j := i + 1
		if j%groupNum == 0 {
			t.Log(j/groupNum, "finished")
		}
	}
}

func testOperation(t *testing.T, fy float64, bigFunc func(BNOperatee) *Big, deciFunc func(decimal.Decimal) decimal.Decimal) {
	sfy := strconv.FormatFloat(fy, 'f', -1, 64)
	iy := rand.Int63()
	uy := uint64(rand.Int63())
	// add Big
	deciRes := deciFunc(decimal.NewFromFloat(fy))
	bigRes := bigFunc(BN(fy))
	if bigRes.String() != deciRes.String() {
		t.Error("add Big not equal", bigRes.String(), deciRes.String())
		t.FailNow()
	}

	// operate float64
	bigRes = bigFunc(Float(fy))
	if bigRes.String() != deciRes.String() {
		t.Error("add Float not equal", bigRes.String(), deciRes.String())
		t.FailNow()
	}

	// operate string
	deciStr, err := decimal.NewFromString(sfy)
	if err != nil {
		t.Error("new decimal from string", sfy, "err", err)
		t.FailNow()
	}
	deciRes = deciFunc(deciStr)
	bigRes = bigFunc(String(sfy))
	if bigRes.String() != deciRes.String() {
		t.Error("add String not equal", bigRes.String(), deciRes.String())
		t.FailNow()
	}

	// operate int64
	deciRes = deciFunc(decimal.NewFromInt(iy))
	bigRes = bigFunc(Int(iy))
	if bigRes.String() != deciRes.String() {
		t.Error("add Int not equal", bigRes.String(), deciRes.String())
		t.FailNow()
	}

	// operate uint64
	deciRes = deciFunc(decimal.NewFromInt(int64(uy)))
	bigRes = bigFunc(Uint(uy))
	if bigRes.String() != deciRes.String() {
		t.Error("add Int not equal", bigRes.String(), deciRes.String())
		t.FailNow()
	}

}

func TestBig_Operate(t *testing.T) {
	times := 100_000_000
	groupNum := 100_000
	groups := times / groupNum
	for i := 0; i < times; i++ {
		fx := RandFloat(-1_000_000_000_000, 1_000_000_000_000)
		fy := RandFloat(-1_000_000_000_000, 1_000_000_000_000)
		x := BN(fx)
		dx := decimal.NewFromFloat(fx)

		testOperation(t, fy, x.Add, dx.Add)
		testOperation(t, fy, x.Sub, dx.Sub)
		testOperation(t, fy, x.Mul, dx.Mul)
		testOperation(t, fy, x.Div, dx.Div)
		//fx = RandFloat(-1_000_000, 1_000_000)
		//fy = RandFloat(1, 10)
		//x = BN(fx)
		//dx = decimal.NewFromFloat(fx)
		//testOperation(t, fy, x.Pow, dx.Pow)

		j := i + 1
		if j%groupNum == 0 {
			t.Log("group", j/groupNum, "/", groups, "finished", j)
		}
	}
}

func TestBig_Sqrt(t *testing.T) {
	times := 1_000_000
	groupNum := 100_000
	groups := times / groupNum

	for i := 0; i < times; i++ {
		fx := RandFloat(0, 1_000_000_000_000)
		x := BN(fx)
		dx := decimal.NewFromFloat(fx)

		xSqrted := x.Sqrt()
		sx2 := xSqrted.Pow(Int(4))
		dx2 := dx.Pow(decimal.NewFromFloat(2))

		if sx2.Div(String(dx2.String())).Sub(Int(1)).Abs().Gt(Float(0.0000000000000001)) {
			t.Error(sx2.String(), dx2.String())
			t.FailNow()
		}

		j := i + 1
		if j%groupNum == 0 {
			t.Log("group", j/groupNum, "/", groups, "finished", j)
		}
	}
}

func TestBig_Abs(t *testing.T) {
	times := 1_000_000_000
	groupNum := 100_000
	groups := times / groupNum

	for i := 0; i < times; i++ {
		fx := RandFloat(-1_000_000_000_000, 1_000_000_000_000)
		x := BN(fx)

		sx := x.String()
		absx := x.Abs().String()

		if strings.ReplaceAll(sx, "-", "") != absx {
			t.Error("abs not equal", sx, absx)
		}

		j := i + 1
		if j%groupNum == 0 {
			t.Log("group", j/groupNum, "/", groups, "finished", j)
		}
	}
}

func TestBig_Cmp(t *testing.T) {
	times := 1_000_000_000
	groupNum := 100_000
	groups := times / groupNum

	for i := 0; i < times; i++ {
		fx := RandFloat(-1_000_000_000_000, 1_000_000_000_000)
		fy := RandFloat(-1_000_000_000_000, 1_000_000_000_000)
		x := BN(fx)
		y := BN(fy)

		sy := strconv.FormatFloat(fy, 'f', -1, 64)
		iy := rand.Int63()
		//uy := uint64(rand.Int63())

		var fcmp int
		if fx < fy {
			fcmp = -1
		} else if fx == fy {
			fcmp = 0
		} else {
			fcmp = 1
		}

		var icmp int
		if fx < float64(iy) {
			icmp = -1
		} else if fx == float64(iy) {
			icmp = 0
		} else {
			icmp = 1
		}

		if x.Cmp(y) != fcmp {
			t.Error("fcmp", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}

		if x.Cmp(String(sy)) != fcmp {
			t.Error("fcmp", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}

		if x.Gte(Float(fy)) && fcmp == -1 {
			t.Error("gte", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}

		if x.Gt(Int(iy)) && icmp != 1 {
			t.Error("gt", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}

		//if x.Gte(Uint(uy)) && fcmp == -1 {
		//	t.Error("gte", "fx", fx, "fy", fy, "x", x, "y", y)
		//	t.FailNow()
		//}

		if x.Lt(y) && fcmp != -1 {
			t.Error("lt", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}
		if x.Lt(String(sy)) && fcmp != -1 {
			t.Error("slt", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}
		if x.Lt(Int(iy)) && icmp != -1 {
			t.Error("ilt", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}
		if x.Lt(Float(fy)) && fcmp != -1 {
			t.Error("flt", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}

		//if x.Lt(Uint(uy)) && fcmp != -1 {
		//	t.Error("lt", "fx", fx, "fy", fy, "x", x, "y", y)
		//	t.FailNow()
		//}

		if x.Lte(y) && fcmp == 1 {
			t.Error("lte", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}
		if x.Lte(String(sy)) && fcmp == 1 {
			t.Error("lte", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}
		if x.Lte(Float(fy)) && fcmp == 1 {
			t.Error("lte", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}
		if x.Lte(Int(iy)) && icmp == 1 {
			t.Error("lte", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}
		//if x.Lte(Uint(uy)) && fcmp == 1 {
		//	t.Error("lte", "fx", fx, "fy", fy, "x", x, "y", y)
		//	t.FailNow()
		//}

		if x.Equal(y) && fcmp != 0 {
			t.Error("eq", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}
		if x.Equal(String(sy)) && fcmp != 0 {
			t.Error("eq", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}
		if x.Equal(Float(fy)) && fcmp != 0 {
			t.Error("eq", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}
		if x.Equal(Int(iy)) && icmp != 0 {
			t.Error("eq", "fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}
		//if x.Equal(Uint(uy)) && fcmp != 0 {
		//	t.Error("eq", "fx", fx, "fy", fy, "x", x, "y", y)
		//	t.FailNow()
		//}

		if x.Cmp(BN(fx)) != 0 {
			t.Error("x != x")
			t.FailNow()
		}

		j := i + 1
		if j%groupNum == 0 {
			t.Log("group", j/groupNum, "/", groups, "finished", j)
		}
	}
}

func TestBig_Round(t *testing.T) {
	testRound := func(places int32, bigFunc func(int32) *Big, deciFunc func(int32) decimal.Decimal) {
		sbig := bigFunc(places).String()
		sdeci := deciFunc(places).String()
		if sbig != sdeci {
			t.Error(sbig, sdeci)
			t.FailNow()
		}
	}

	times := 1_000_000
	groupNum := 100_000
	groups := times / groupNum

	places := int32(300)

	for i := 0; i < times; i++ {
		fx := RandFloat(-1_000_000_000, 1_000_000_000)
		sfx := strconv.FormatFloat(fx, 'f', -1, 64)
		is := strconv.FormatInt(rand.Int63n(1_000_000_000_000), 10)
		ds := strconv.FormatInt(rand.Int63n(1_000_000_000_000), 10)
		sfx = is + sfx + ds
		if strings.Contains(sfx, "-") {
			sfx = "-" + strings.ReplaceAll(sfx, "-", "")
		}
		x := BN(sfx)
		d, err := decimal.NewFromString(sfx)
		if err != nil {
			t.Error("new decimal from string", sfx, "err", err)
			t.FailNow()
		}
		for k := -places; k <= places; k++ {
			testRound(k, x.Round, d.Round)
			testRound(k, x.RoundCeil, d.RoundCeil)
			testRound(k, x.RoundFloor, d.RoundFloor)
			testRound(k, x.RoundInfinity, d.RoundUp)
			testRound(k, x.RoundZero, d.RoundDown)
		}

		j := i + 1
		if j%groupNum == 0 {
			t.Log("group", j/groupNum, "/", groups, "finished", j)
		}
	}
}

func TestBig_BigFloat(t *testing.T) {
	times := 1_000_000_000
	groupNum := 100_000
	groups := times / groupNum

	for i := 0; i < times; i++ {
		fx := RandFloat(-1_000_000_000_000, 1_000_000_000_000)
		x := BN(fx)
		//if math.Abs(x.Float64()/fx-1) > 0.000000000000001 {
		if x.Float64() != fx {
			t.Error(fx, x.String())
			t.FailNow()
		}

		j := i + 1
		if j%groupNum == 0 {
			t.Log("group", j/groupNum, "/", groups, "finished", j)
		}
	}
}

func TestBig_BigIntRound(t *testing.T) {
	times := 1_000_000_000
	groupNum := 100_000
	groups := times / groupNum

	for i := 0; i < times; i++ {
		fx := RandFloat(-1_000_000_000_000, 1_000_000_000_000)
		x := BN(fx)

		sfx := strconv.FormatFloat(fx, 'f', -1, 64)

		if !strings.Contains(sfx, ".") {
			sfx += ".0"
		}

		neg := strings.Contains(sfx, "-")

		sfxs := strings.Split(sfx, ".")

		si := sfxs[0]

		sd := sfxs[1]

		lenSdNon0 := len(strings.ReplaceAll(sd, "0", ""))

		keynum := sd[0]

		// for BigInt
		if x.BigInt().String() != si {
			t.Error("Big.BigInt test", x, x.BigInt(), si)
			t.FailNow()
		}

		// for round
		switch keynum {
		case '0', '1', '2', '3', '4':
			if x.BigIntRound().String() != si {
				t.Error("Big.BigIntRound", x, x.BigIntRound(), si)
				t.FailNow()
			}
		default:
			one := int64(-1)
			if neg {
				one = 1
			}
			if big.NewInt(0).Add(x.BigIntRound(), big.NewInt(one)).String() != si {
				t.Error("Big.BigIntRound", x, x.BigIntRound(), si)
				t.FailNow()
			}
		}

		// for ceil
		if lenSdNon0 == 0 {
			if x.BigIntCeil().String() != si {
				t.Error("Big.BigIntCeil", neg, x, x.BigIntCeil(), si)
				t.FailNow()
			}
		} else {
			if neg && x.BigIntCeil().String() != si {
				t.Error("Big.BigIntCeil", neg, x, x.BigIntCeil(), si)
				t.FailNow()
			}
			if !neg && big.NewInt(0).Sub(x.BigIntCeil(), big.NewInt(1)).String() != si {
				t.Error("Big.BigIntCeil", neg, x, x.BigIntCeil(), si, big.NewInt(0).Sub(x.BigIntCeil(), big.NewInt(1)))
				t.FailNow()
			}
		}

		// for floor
		if lenSdNon0 == 0 {
			if x.BigIntFloor().String() != si {
				t.Error("Big.BigIntFloor", x, x.BigIntFloor(), si)
				t.FailNow()
			}
		} else {
			if !neg && x.BigIntFloor().String() != si {
				t.Error("Big.BigIntFloor", x, x.BigIntFloor(), si)
				t.FailNow()
			}
			if neg && big.NewInt(0).Add(x.BigIntFloor(), big.NewInt(1)).String() != si {
				t.Error("Big.BigIntFloor", x, x.BigIntFloor(), si, big.NewInt(0).Add(x.BigIntFloor(), big.NewInt(1)).String())
				t.FailNow()
			}
		}

		// for infinity
		if lenSdNon0 == 0 {
			if x.BigIntInfinity().String() != si {
				t.Error("Big.BigIntInfinity", x, x.BigIntInfinity(), si)
				t.FailNow()
			}
		} else {
			one := int64(-1)
			if neg {
				one = 1
			}
			if big.NewInt(0).Add(x.BigIntInfinity(), big.NewInt(one)).String() != si {
				t.Error("Big.BigIntInfinity", x, x.BigIntInfinity(), si)
				t.FailNow()
			}
		}

		// for zero
		if x.BigIntZero().String() != si {
			t.Error("Big.BigIntZero", x, x.BigIntZero(), si)
			t.FailNow()
		}

		j := i + 1
		if j%groupNum == 0 {
			t.Log("group", j/groupNum, "/", groups, "finished", j)
		}
	}
}

func TestMaxBN(t *testing.T) {
	times := 1_000_000_000
	groupNum := 100_000
	groups := times / groupNum

	for i := 0; i < times; i++ {
		fx := RandFloat(-1_000_000_000_000, 1_000_000_000_000)
		fy := RandFloat(-1_000_000_000_000, 1_000_000_000_000)
		x := BN(fx)
		y := BN(fy)

		if (fx == fy && (MaxBN(x, y) != x || MaxBN(y, x) != y || MinBN(x, y) != x || MinBN(y, x) != y)) ||
			(fx > fy && (MaxBN(x, y) != x || MinBN(x, y) != y)) ||
			(fx < fy && (MinBN(x, y) != x || MaxBN(x, y) != y)) {
			t.Error("fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}

		j := i + 1
		if j%groupNum == 0 {
			t.Log("group", j/groupNum, "/", groups, "finished", j)
		}
	}
}

func TestMinBN(t *testing.T) {
	times := 1_000_000_000
	groupNum := 100_000
	groups := times / groupNum

	for i := 0; i < times; i++ {
		fx := RandFloat(-1_000_000_000_000, 1_000_000_000_000)
		fy := RandFloat(-1_000_000_000_000, 1_000_000_000_000)
		x := BN(fx)
		y := BN(fy)

		if (fx == fy && (MaxBN(x, y) != x || MaxBN(y, x) != y || MinBN(x, y) != x || MinBN(y, x) != y)) ||
			(fx > fy && (MaxBN(x, y) != x || MinBN(x, y) != y)) ||
			(fx < fy && (MinBN(x, y) != x || MaxBN(x, y) != y)) {
			t.Error("fx", fx, "fy", fy, "x", x, "y", y)
			t.FailNow()
		}

		j := i + 1
		if j%groupNum == 0 {
			t.Log("group", j/groupNum, "/", groups, "finished", j)
		}
	}
}
