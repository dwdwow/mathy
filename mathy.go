package mathy

import (
	"math/rand"
)

func Round(f float64, places int32) float64 {
	return BN(f).Round(places).Float64()
}

func RoundCeil(f float64, places int32) float64 {
	return BN(f).RoundCeil(places).Float64()
}
func RoundFloor(f float64, places int32) float64 {
	return BN(f).RoundFloor(places).Float64()
}
func RoundInfinity(f float64, places int32) float64 {
	return BN(f).RoundInfinity(places).Float64()
}
func RoundZero(f float64, places int32) float64 {
	return BN(f).RoundZero(places).Float64()
}

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
