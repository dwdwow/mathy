package mathy

import (
	"math/big"

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
