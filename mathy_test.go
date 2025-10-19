package mathy

import (
	"math"
	"testing"
)

// Test rounding functions
func TestRoundingFunctions(t *testing.T) {
	t.Run("Round function", func(t *testing.T) {
		// Test basic rounding
		result := Round(3.14159, 2)
		expected := 3.14
		if math.Abs(result-expected) > 0.0001 {
			t.Errorf("Expected %f, got %f", expected, result)
		}

		// Test rounding to integer
		result = Round(3.7, 0)
		expected = 4.0
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}

		// Test negative numbers
		result = Round(-3.14159, 2)
		expected = -3.14
		if math.Abs(result-expected) > 0.0001 {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("RoundCeil function", func(t *testing.T) {
		// Test ceiling rounding
		result := RoundCeil(3.1, 0)
		expected := 4.0
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}

		// Test negative numbers
		result = RoundCeil(-3.1, 0)
		expected = -3.0
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}

		// Test with decimal places
		result = RoundCeil(3.14159, 2)
		expected = 3.15
		if math.Abs(result-expected) > 0.0001 {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("RoundFloor function", func(t *testing.T) {
		// Test floor rounding
		result := RoundFloor(3.9, 0)
		expected := 3.0
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}

		// Test negative numbers
		result = RoundFloor(-3.9, 0)
		expected = -4.0
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}

		// Test with decimal places
		result = RoundFloor(3.14159, 2)
		expected = 3.14
		if math.Abs(result-expected) > 0.0001 {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("RoundInfinity function", func(t *testing.T) {
		// Test infinity rounding (away from zero)
		result := RoundInfinity(3.1, 0)
		expected := 4.0
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}

		// Test negative numbers
		result = RoundInfinity(-3.1, 0)
		expected = -4.0
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}

		// Test with decimal places
		result = RoundInfinity(3.14159, 2)
		expected = 3.15
		if math.Abs(result-expected) > 0.0001 {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("RoundZero function", func(t *testing.T) {
		// Test zero rounding (towards zero)
		result := RoundZero(3.9, 0)
		expected := 3.0
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}

		// Test negative numbers
		result = RoundZero(-3.9, 0)
		expected = -3.0
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}

		// Test with decimal places
		result = RoundZero(3.14159, 2)
		expected = 3.14
		if math.Abs(result-expected) > 0.0001 {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})
}

// Test random functions
func TestRandomFunctions(t *testing.T) {
	t.Run("RandFloat function", func(t *testing.T) {
		// Test basic range
		min, max := 0.0, 10.0
		for i := 0; i < 100; i++ {
			result := RandFloat(min, max)
			if result < min || result >= max {
				t.Errorf("RandFloat(%f, %f) = %f, expected value in range [%f, %f)", min, max, result, min, max)
			}
		}

		// Test negative range
		min, max = -5.0, 5.0
		for i := 0; i < 100; i++ {
			result := RandFloat(min, max)
			if result < min || result >= max {
				t.Errorf("RandFloat(%f, %f) = %f, expected value in range [%f, %f)", min, max, result, min, max)
			}
		}

		// Test same min and max (should return min)
		result := RandFloat(5.0, 5.0)
		if result != 5.0 {
			t.Errorf("RandFloat(5.0, 5.0) = %f, expected 5.0", result)
		}

		// Test decimal range
		min, max = 1.5, 2.5
		for i := 0; i < 100; i++ {
			result := RandFloat(min, max)
			if result < min || result >= max {
				t.Errorf("RandFloat(%f, %f) = %f, expected value in range [%f, %f)", min, max, result, min, max)
			}
		}
	})

	t.Run("RandFloats function", func(t *testing.T) {
		// Test basic functionality
		min, max := 0.0, 10.0
		n := 50
		results := RandFloats(min, max, n)

		// Check length
		if len(results) != n {
			t.Errorf("RandFloats(%f, %f, %d) returned %d values, expected %d", min, max, n, len(results), n)
		}

		// Check all values are in range
		for i, result := range results {
			if result < min || result >= max {
				t.Errorf("RandFloats[%d] = %f, expected value in range [%f, %f)", i, result, min, max)
			}
		}

		// Test with zero count
		results = RandFloats(min, max, 0)
		if len(results) != 0 {
			t.Errorf("RandFloats(%f, %f, 0) returned %d values, expected 0", min, max, len(results))
		}

		// Test with negative range
		min, max = -10.0, -5.0
		n = 30
		results = RandFloats(min, max, n)

		if len(results) != n {
			t.Errorf("RandFloats(%f, %f, %d) returned %d values, expected %d", min, max, n, len(results), n)
		}

		for i, result := range results {
			if result < min || result >= max {
				t.Errorf("RandFloats[%d] = %f, expected value in range [%f, %f)", i, result, min, max)
			}
		}

		// Test that results are different (very high probability)
		min, max = 0.0, 100.0
		n = 1000
		results = RandFloats(min, max, n)

		// Check for uniqueness (should be very likely with large range and many values)
		uniqueCount := 0
		seen := make(map[float64]bool)
		for _, result := range results {
			if !seen[result] {
				seen[result] = true
				uniqueCount++
			}
		}

		// With 1000 values in range [0, 100), we should have many unique values
		if uniqueCount < 500 {
			t.Errorf("RandFloats produced only %d unique values out of %d, expected more variety", uniqueCount, n)
		}
	})
}
