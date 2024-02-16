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

type Big interface {
	Add(Big) Big
	Sub(Big) Big
	Mul(Big) Big
	Div(Big) Big
	Pow(Big) Big

	Cmp(Big) int
	Equal(Big) bool
	Gt(Big) bool
	Gte(Big) bool
	Lt(Big) bool
	Lte(Big) bool

	Sqrt() Big
	Abs() Big

	Round(int32) Big
	RoundCeil(int32) Big
	RoundFloor(int32) Big
	RoundInfinity(int32) Big
	RoundZero(int32) Big

	String() string
	Float64() float64
	BigFloat() *big.Float
}

type BigBaseNumber interface {
	uint64 | int64 | float64 | string
	//*Big[uint64] | *Big[int64] | *Big[float64] | *Big[string]
}

type BigNumber struct {
	d decimal.Decimal
}

func BN[N BigBaseNumber](n N) *BigNumber {
	return newBN(n)
}

func newBN[N BigBaseNumber](n N) *BigNumber {
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

func newBNWithInt64(n int64) *BigNumber {
	return &BigNumber{d: decimal.NewFromInt(n)}
}

func newBNWithFloat64(n float64) *BigNumber {
	return &BigNumber{d: decimal.NewFromFloat(n)}
}

func newDecimalWithStrIgnoreErr(n string) decimal.Decimal {
	if n == "" {
		n = "0"
	}
	d, err := decimal.NewFromString(n)
	if err != nil {
		panic("mathy: can not new BigNumber with string " + n)
	}
	return d
}

func newBigWithStr(n string) (*BigNumber, error) {
	if n == "" {
		n = "0"
	}
	d, err := decimal.NewFromString(n)
	if err != nil {
		return nil, err
	}
	return &BigNumber{d: d}, nil
}

func newBNWithStrIgnoreErr(n string) *BigNumber {
	b, err := newBigWithStr(n)
	if err != nil {
		panic("mathy: newBigWithStr, " + err.Error())
	}
	return b
}

func newBNWithUint64(n uint64) *BigNumber {
	return newBNWithStrIgnoreErr(strconv.FormatUint(n, 10))
}

func (b *BigNumber) Copy() *BigNumber {
	return b.copy()
}

func (b *BigNumber) copy() *BigNumber {
	return &BigNumber{d: b.d.Copy()}
}

func (b *BigNumber) operateBig(n Big, f func(decimal.Decimal) decimal.Decimal) Big {
	if n == nil {
		return b.copy()
	}
	s := n.String()
	if s == "" {
		s = "0"
	}
	return &BigNumber{d: f(newDecimalWithStrIgnoreErr(n.String()))}
}

func (b *BigNumber) Add(n Big) Big {
	return b.operateBig(n, b.d.Add)
}

func (b *BigNumber) Sub(n Big) Big {
	return b.operateBig(n, b.d.Sub)
}

func (b *BigNumber) Mul(n Big) Big {
	return b.operateBig(n, b.d.Mul)
}

func (b *BigNumber) Div(n Big) Big {
	return b.operateBig(n, b.d.Div)
}

func (b *BigNumber) Pow(n Big) Big {
	return b.operateBig(n, b.d.Pow)
}

func (b *BigNumber) Cmp(n Big) int {
	return b.d.Cmp(newDecimalWithStrIgnoreErr(n.String()))
}

func (b *BigNumber) Equal(n Big) bool {
	return b.Cmp(n) == 0
}

func (b *BigNumber) Gt(n Big) bool {
	return b.Cmp(n) == 1
}

func (b *BigNumber) Gte(n Big) bool {
	return b.Cmp(n) > -1
}

func (b *BigNumber) Lt(n Big) bool {
	return b.Cmp(n) == -1
}

func (b *BigNumber) Lte(n Big) bool {
	return b.Cmp(n) < 1
}

func (b *BigNumber) Round(places int32) Big {
	return &BigNumber{d: b.d.Round(places)}
}

func (b *BigNumber) RoundCeil(places int32) Big {
	return &BigNumber{d: b.d.RoundCeil(places)}
}

func (b *BigNumber) RoundFloor(places int32) Big {
	return &BigNumber{d: b.d.RoundFloor(places)}
}

func (b *BigNumber) RoundInfinity(places int32) Big {
	return &BigNumber{d: b.d.RoundUp(places)}
}

func (b *BigNumber) RoundZero(places int32) Big {
	return &BigNumber{d: b.d.RoundDown(places)}
}

func (b *BigNumber) String() string {
	return b.d.String()
}

func (b *BigNumber) Float64() float64 {
	f, _ := b.d.Float64()
	return f
}

func (b *BigNumber) Sqrt() Big {
	bf := new(big.Float).
		SetPrec(512).
		Sqrt(b.d.BigFloat()).
		Text('f', -1)
	d := newDecimalWithStrIgnoreErr(bf)
	return &BigNumber{d: d}
}

func (b *BigNumber) Abs() Big {
	return &BigNumber{d: b.d.Abs()}
}

func (b *BigNumber) BigFloat() *big.Float {
	return b.d.BigFloat()
}
