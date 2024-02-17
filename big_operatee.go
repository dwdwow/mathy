package mathy

type BNOperatee interface {
	BN() *Big
}

type Float float64

func (f Float) BN() *Big {
	return newBNWithFloat64(float64(f))
}

type Int int64

func (i Int) BN() *Big {
	return newBNWithInt64(int64(i))
}

type String string

func (s String) BN() *Big {
	return newBNWithStrIgnoreErr(string(s))
}

type Uint uint64

func (u Uint) BN() *Big {
	return newBNWithUint64(uint64(u))
}
