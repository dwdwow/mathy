package mathy

import (
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestRandFloat(t *testing.T) {
	var _r float64
	for i := 0; i < 1_000_000_000; i++ {
		r := RandFloat(10, 100)
		if r < 10 || r >= 100 {
			panic(r)
		}
		if r == _r {
			panic(r)
		}
		_r = r
	}
}

func TestRandFloats(t *testing.T) {
	var _r float64
	for i := 0; i < 1000; i++ {
		rs := RandFloats(10, 100, 1_000_000)
		for _, r := range rs {
			if r < 10 || r >= 100 {
				panic(r)
			}
			if r == _r {
				panic(r)
			}
			_r = r
		}
	}
}

type trimDecimalsType int

const (
	trimTypeRound trimDecimalsType = iota
	trimTypeCeil
	trimTypeFloor
	trimTypeInfinity
	trimTypeZero
)

func testTrim(t *testing.T, trim func(float64, int32) float64, decimalsType trimDecimalsType) {
	times := 10_000_000
	groupNum := 100_000
	groups := times / groupNum
	placesSection := 60
	var placesList []int32
	for i := -placesSection; i <= placesSection; i++ {
		placesList = append(placesList, int32(i))
	}
	for _, places := range placesList {
		for i := 0; i < times; i++ {
			origin := RandFloat(-100_000_000, 100_000_000)
			trimmed := trim(origin, places)

			sign := 1.0
			if origin < 0 {
				sign = -1
				origin *= -1
			}

			sorigin := strconv.FormatFloat(origin, 'f', -1, 64)

			origin *= sign

			sorigin = strings.Repeat("0", 100) + sorigin

			if !strings.Contains(sorigin, ".") {
				sorigin += "."
			}

			sorigin += strings.Repeat("0", 100)

			pointIndex := strings.IndexByte(sorigin, '.')

			pointShiftedIndex := pointIndex + int(places)

			sorigin = strings.ReplaceAll(sorigin, ".", "")
			shiftedIntegers := sorigin[:pointShiftedIndex]

			_trimmed, err := strconv.ParseFloat(shiftedIntegers, 64)
			if err != nil {
				t.Error("places", places, "parse shiftedIntegers", "err", err, "origin", origin, "trimmed", trimmed, "_trimmed", _trimmed)
				t.FailNow()
			}

			switch {
			case (decimalsType == trimTypeCeil) && sign > 0,
				(decimalsType == trimTypeFloor) && sign < 0,
				decimalsType == trimTypeInfinity:
				shiftedDecimals := sorigin[pointShiftedIndex:]
				no0decimals := strings.ReplaceAll(shiftedDecimals, "0", "")
				if len(no0decimals) != 0 {
					_trimmed++
				}
			case decimalsType == trimTypeRound:
				switch sorigin[pointShiftedIndex] {
				case '5', '6', '7', '8', '9':
					_trimmed++
				}
			}

			_strimmed := strconv.FormatFloat(_trimmed, 'f', -1, 64)

			if places < 0 {
				_strimmed += strings.Repeat("0", int(-places))
			} else if places > 0 {
				_sceiledList := strings.Split(_strimmed, "")
				_sceiledList = slices.Insert(_sceiledList, len(_sceiledList)-int(places), ".")
				_strimmed = strings.Join(_sceiledList, "")
			}

			_trimmed, err = strconv.ParseFloat(_strimmed, 64)
			if err != nil {
				t.Error("places", places, "parse _strimmed", "err", err, "origin", origin, "trimmed", trimmed, "_trimmed", _trimmed)
				t.FailNow()
			}

			_trimmed *= sign

			if trimmed != _trimmed {
				t.Error("places", places, "origin", strconv.FormatFloat(origin, 'f', -1, 64), origin, "trimmed", strconv.FormatFloat(trimmed, 'f', -1, 64), trimmed, "_trimmed", strconv.FormatFloat(_trimmed, 'f', -1, 64))
				t.FailNow()
			}

			j := i + 1
			if j%groupNum == 0 {
				t.Log("places", places, "group", j/groupNum, "/", groups, "tested", j)
			}
		}
	}
}

func TestRound(t *testing.T) {
	testTrim(t, Round, trimTypeRound)
}

func TestRoundCeil(t *testing.T) {
	testTrim(t, RoundCeil, trimTypeCeil)
}

func TestRoundFloor(t *testing.T) {
	testTrim(t, RoundFloor, trimTypeFloor)
}

func TestRoundInfinity(t *testing.T) {
	testTrim(t, RoundInfinity, trimTypeInfinity)
}

func TestRoundZero(t *testing.T) {
	testTrim(t, RoundZero, trimTypeZero)
}
