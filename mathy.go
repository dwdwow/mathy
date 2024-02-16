package mathy

import (
	"math"
	"math/rand"
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

func Round(input float64, places int) float64 {
	trimmed, _ := trimDecimals(input, places, trimDecimalsTypeRound)
	return trimmed
}

func Ceil(input float64, places int) float64 {
	trimmed, _ := trimDecimals(input, places, trimDecimalsTypeCeil)
	return trimmed
}

func Floor(input float64, places int) float64 {
	trimmed, _ := trimDecimals(input, places, trimDecimalsTypeFloor)
	return trimmed
}

// Borrowed from github.com/montanaflynn/stats
func trimDecimals(input float64, places int, trimType trimDecimalsType) (trimmed float64, isNaN bool) {

	// If the float is not a number
	if math.IsNaN(input) {
		return input, true
	}

	// Find out the actual sign and correct the input for later
	sign := 1.0
	if input < 0 {
		sign = -1
		input *= -1
	}

	// Use the places arg to get the amount of precision wanted
	precision := math.Pow(10, float64(places))

	// Find the decimal place we are looking to round
	digit := input * precision

	// Get the actual decimal number as a fraction to be compared
	_, decimal := math.Modf(digit)

	switch {
	case (trimType == trimDecimalsTypeRound) == (decimal >= 0.5),
		(trimType == trimDecimalsTypeCeil) == (sign == 1),
		(trimType == trimDecimalsTypeFloor) == (sign == -1):
		trimmed = math.Ceil(digit)
	default: // ignore unknown trimDecimalTypes
		trimmed = math.Floor(digit)
	}

	// Finally we do the math to actually create a rounded number
	return trimmed / precision * sign, false
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
