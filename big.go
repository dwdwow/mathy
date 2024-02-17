package mathy

import (
	"math/big"
	"strconv"

	"github.com/shopspring/decimal"
)

func init() {
	// must set DivisionPrecision, default is 16
	decimal.DivisionPrecision = 100
}

type BigBaseNumber interface {
	uint64 | int64 | float64 | string
}

// Big is a wrapper of github.com/shopspring/decimal,
// has more DIY api.
type Big struct {
	d decimal.Decimal
}

func BN[N BigBaseNumber](n N) *Big {
	return newBN(n)
}

func newBN[N BigBaseNumber](n N) *Big {
	switch n := any(n).(type) {
	case float64:
		return newBNWithFloat64(n)
	case string:
		return newBNWithStrIgnoreErr(n)
	case int64:
		return newBNWithInt64(n)
	case uint64:
		return newBNWithUint64(n)
	default:
		return nil
	}
}

func newBNWithInt64(n int64) *Big {
	return &Big{d: decimal.NewFromInt(n)}
}

func newBNWithFloat64(n float64) *Big {
	return &Big{d: decimal.NewFromFloat(n)}
}

func newDecimalWithStrIgnoreErr(n string) decimal.Decimal {
	if n == "" {
		n = "0"
	}
	d, err := decimal.NewFromString(n)
	if err != nil {
		panic("mathy: can not new Big with string " + n)
	}
	return d
}

func newBigWithStr(n string) (*Big, error) {
	if n == "" {
		n = "0"
	}
	d, err := decimal.NewFromString(n)
	if err != nil {
		return nil, err
	}
	return &Big{d: d}, nil
}

func newBNWithStrIgnoreErr(n string) *Big {
	b, err := newBigWithStr(n)
	if err != nil {
		panic("mathy: newBigWithStr, " + err.Error())
	}
	return b
}

func newBNWithUint64(n uint64) *Big {
	return newBNWithStrIgnoreErr(strconv.FormatUint(n, 10))
}

func (b *Big) BN() *Big {
	if b == nil {
		return &Big{d: decimal.NewFromInt(0)}
	}
	return b
}

func (b *Big) Copy() *Big {
	return b.copy()
}

func (b *Big) copy() *Big {
	return &Big{d: b.d.Copy()}
}

func (b *Big) operateBig(n BNOperatee, f func(decimal.Decimal) decimal.Decimal) *Big {
	if n == nil {
		return b.copy()
	}
	return &Big{d: f(n.BN().d)}
}

func (b *Big) Add(n BNOperatee) *Big {
	return b.operateBig(n, b.d.Add)
}

func (b *Big) Sub(n BNOperatee) *Big {
	return b.operateBig(n, b.d.Sub)
}

func (b *Big) Mul(n BNOperatee) *Big {
	return b.operateBig(n, b.d.Mul)
}

func (b *Big) Div(n BNOperatee) *Big {
	return b.operateBig(n, b.d.Div)
}

func (b *Big) Pow(n BNOperatee) *Big {
	return b.operateBig(n, b.d.Pow)
}

func (b *Big) Cmp(n BNOperatee) int {
	return b.d.Cmp(n.BN().d)
}

func (b *Big) Equal(n BNOperatee) bool {
	return b.Cmp(n) == 0
}

func (b *Big) Gt(n BNOperatee) bool {
	return b.Cmp(n) == 1
}

func (b *Big) Gte(n BNOperatee) bool {
	return b.Cmp(n) > -1
}

func (b *Big) Lt(n BNOperatee) bool {
	return b.Cmp(n) == -1
}

func (b *Big) Lte(n BNOperatee) bool {
	return b.Cmp(n) < 1
}

func (b *Big) Round(places int32) *Big {
	return &Big{d: b.d.Round(places)}
}

func (b *Big) RoundCeil(places int32) *Big {
	return &Big{d: b.d.RoundCeil(places)}
}

func (b *Big) RoundFloor(places int32) *Big {
	return &Big{d: b.d.RoundFloor(places)}
}

func (b *Big) RoundInfinity(places int32) *Big {
	return &Big{d: b.d.RoundUp(places)}
}

func (b *Big) RoundZero(places int32) *Big {
	return &Big{d: b.d.RoundDown(places)}
}

func (b *Big) String() string {
	return b.d.String()
}

func (b *Big) Float64() float64 {
	f, _ := b.d.Float64()
	return f
}

func (b *Big) Sqrt() *Big {
	bf := new(big.Float).
		SetPrec(512).
		Sqrt(b.d.BigFloat()).
		Text('f', -1)
	d := newDecimalWithStrIgnoreErr(bf)
	return &Big{d: d}
}

func (b *Big) Abs() *Big {
	return &Big{d: b.d.Abs()}
}

// BigFloat returns decimal as BigFloat.
// Be aware that casting decimal to BigFloat might cause a loss of precision.
func (b *Big) BigFloat() *big.Float {
	return b.d.BigFloat()
}

func (b *Big) BigIntRound() *big.Int {
	return b.d.BigInt()
}

func (b *Big) BigIntCeil() *big.Int {
	return b.d.RoundCeil(0).BigInt()
}

func (b *Big) BigIntFloor() *big.Int {
	return b.d.RoundFloor(0).BigInt()
}

func (b *Big) BigIntInfinity() *big.Int {
	return b.d.RoundUp(0).BigInt()
}

func (b *Big) BigIntZero() *big.Int {
	return b.d.RoundDown(0).BigInt()
}

func MaxBN(x, y *Big) *Big {
	if x.Gte(y) {
		return x
	}
	return y
}

func MinBN(x, y *Big) *Big {
	if x.Lte(y) {
		return x
	}
	return y
}
