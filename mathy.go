package mathy

import (
	"errors"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// ======================================================
// Trim decimals
// ------------------------------------------------------

type trimDecimalsType int

const (
	trimDecimalsTypeRound trimDecimalsType = iota
	trimDecimalsTypeCeil
	trimDecimalsTypeFloor
)

func Round(input float64, precision int) float64 {
	trimmed, _ := trimDecimals(input, precision, trimDecimalsTypeRound)
	return trimmed
}

func Ceil(input float64, precision int) float64 {
	trimmed, _ := trimDecimals(input, precision, trimDecimalsTypeCeil)
	return trimmed
}

func Floor(input float64, precision int) float64 {
	trimmed, _ := trimDecimals(input, precision, trimDecimalsTypeFloor)
	return trimmed
}

func trimDecimalsNotPrecise(input float64, precision int, trimType trimDecimalsType) (trimmed float64, isNaN bool) {

	if math.IsNaN(input) {
		return input, true
	}

	sign := 1.0
	if input < 0 {
		sign = -1
		input *= -1
	}

	prec := math.Pow(10, float64(precision))

	digit := input * prec

	_, decimal := math.Modf(digit)

	switch {
	case (trimType == trimDecimalsTypeRound) && (decimal >= 0.5),
		(trimType == trimDecimalsTypeCeil) && (sign == 1),
		(trimType == trimDecimalsTypeFloor) && (sign == -1):
		trimmed = math.Ceil(digit)
	default: // ignore unknown trimDecimalTypes
		trimmed = math.Floor(digit)
	}

	return trimmed / prec * sign, false
}

func trimDecimals(input float64, precision int, trimType trimDecimalsType) (float64, error) {

	if math.IsNaN(input) {
		return input, errors.New("input is NaN")
	}

	digit := input
	sign := 1.0
	if input < 0 {
		sign = -1
		digit *= -1
	}

	prec := math.Pow(10, float64(precision))

	sdigit := strconv.FormatFloat(digit*prec, 'f', -1, 64)

	digit, err := strconv.ParseFloat(sdigit, 64)

	if err != nil {
		return input, err
	}

	var isDecimalGtePoint5, hasDecimals bool

	// if string contains decimal point, the string must be a float
	if strings.Contains(sdigit, ".") {
		hasDecimals = true
		pointIndex := strings.IndexByte(sdigit, '.')
		if len(sdigit) > pointIndex+1 {
			keynum := sdigit[pointIndex+1]
			switch keynum {
			case '5', '6', '7', '8', '9':
				isDecimalGtePoint5 = true
			}
		}
		sdigit := sdigit[:pointIndex]
		digit, err = strconv.ParseFloat(sdigit, 64)
		if err != nil {
			return input, err
		}
	}

	// ignore unknown trimDecimalTypes
	switch {
	case (trimType == trimDecimalsTypeRound) && isDecimalGtePoint5,
		(trimType == trimDecimalsTypeCeil) && (sign == 1) && hasDecimals,
		(trimType == trimDecimalsTypeFloor) && (sign == -1) && hasDecimals:
		digit += 1
	}

	if precision >= 0 {
		strimmed := strconv.FormatFloat(digit/prec*sign, 'f', -1, 64)
		return strconv.ParseFloat(strimmed, 64)
	}

	return sign * digit * math.Pow(10, float64(-precision)), nil
}

// ------------------------------------------------------
// Trim decimals
// ======================================================

// ======================================================
// Random
// ------------------------------------------------------

// RandFloat returns a float in the half-open interval [min, max)
func RandFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandFloats returns a list of float in the half-open interval [min, max)
func RandFloats(min, max float64, n int) []float64 {
	var res []float64
	for i := 0; i < n; i++ {
		res = append(res, RandFloat(min, max))
	}
	return res
}

// ------------------------------------------------------
// Random
// ======================================================
