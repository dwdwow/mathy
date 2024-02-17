package mathy

type BNOperatee interface {
	BN() *Big
}

type OperateeFloat64 float64

func (f OperateeFloat64) BN() *Big {
	return newBNWithFloat64(float64(f))
}

type OperateeInt64 int64

func (i OperateeInt64) BN() *Big {
	return newBNWithInt64(int64(i))
}

type OperateeString string

func (s OperateeString) BN() *Big {
	return newBNWithStrIgnoreErr(string(s))
}

type OperateeUint64 uint64

func (u OperateeUint64) BN() *Big {
	return newBNWithUint64(uint64(u))
}
