package mathy

import (
	"fmt"
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

func testTrim(t *testing.T, trim func(float64, int) float64, decimalsType trimDecimalsType) {
	loopTimes := 100_000_000
	placesSection := 6
	totalTimes := float64(loopTimes) * float64(placesSection*2+1)
	failTimes := 0.0
	var placesList []int
	for i := -placesSection; i <= placesSection; i++ {
		placesList = append(placesList, i)
	}
	for _, places := range placesList {
		for i := 0; i < loopTimes; i++ {
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

			pointShiftedIndex := pointIndex + places

			sorigin = strings.ReplaceAll(sorigin, ".", "")
			shiftedIntegers := sorigin[:pointShiftedIndex]

			_trimmed, err := strconv.ParseFloat(shiftedIntegers, 64)
			if err != nil {
				t.Error("places", places, "parse shiftedIntegers", "err", err, "origin", origin, "trimmed", trimmed, "_trimmed", _trimmed)
				t.FailNow()
			}

			switch {
			case (decimalsType == trimDecimalsTypeCeil) && sign > 0,
				(decimalsType == trimDecimalsTypeFloor) && sign < 0:
				shiftedDecimals := sorigin[pointShiftedIndex:]
				no0decimals := strings.ReplaceAll(shiftedDecimals, "0", "")
				if len(no0decimals) != 0 {
					_trimmed += 1
				}
			case decimalsType == trimDecimalsTypeRound:
				switch sorigin[pointShiftedIndex] {
				case '5', '6', '7', '8', '9':
					_trimmed += 1
				}
			}

			_strimmed := strconv.FormatFloat(_trimmed, 'f', -1, 64)

			if places < 0 {
				_strimmed += strings.Repeat("0", -places)
			} else if places > 0 {
				_sceiledList := strings.Split(_strimmed, "")
				_sceiledList = slices.Insert(_sceiledList, len(_sceiledList)-places, ".")
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
				failTimes++
			}
		}
	}
	fmt.Println("fail rate", failTimes/totalTimes)
}

func TestRound(t *testing.T) {
	testTrim(t, Round, trimDecimalsTypeRound)
}

func TestCeil(t *testing.T) {
	testTrim(t, Ceil, trimDecimalsTypeCeil)
}

func TestFloor(t *testing.T) {
	testTrim(t, Floor, trimDecimalsTypeFloor)
}
