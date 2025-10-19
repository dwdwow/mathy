package mathy

import (
	"testing"
)

// Test constructors and basic creation
func TestBigConstructors(t *testing.T) {
	// Test BN function with different types
	t.Run("BN with int", func(t *testing.T) {
		b := BN(42)
		if b.String() != "42" {
			t.Errorf("Expected '42', got '%s'", b.String())
		}
	})

	t.Run("BN with int64", func(t *testing.T) {
		b := BN(int64(123))
		if b.String() != "123" {
			t.Errorf("Expected '123', got '%s'", b.String())
		}
	})

	t.Run("BN with float64", func(t *testing.T) {
		b := BN(3.14)
		if b.String() != "3.14" {
			t.Errorf("Expected '3.14', got '%s'", b.String())
		}
	})

	t.Run("BN with string", func(t *testing.T) {
		b := BN("999")
		if b.String() != "999" {
			t.Errorf("Expected '999', got '%s'", b.String())
		}
	})

	t.Run("BN with hex string", func(t *testing.T) {
		b := BN("0xFF")
		if b.String() != "255" {
			t.Errorf("Expected '255', got '%s'", b.String())
		}
	})

	t.Run("BN with binary string", func(t *testing.T) {
		b := BN("0b1010")
		if b.String() != "10" {
			t.Errorf("Expected '10', got '%s'", b.String())
		}
	})
}

// Test NumberLikeToBN function
func TestNumberLikeToBN(t *testing.T) {
	t.Run("Convert int", func(t *testing.T) {
		b := NumberLikeToBN(42)
		if b.String() != "42" {
			t.Errorf("Expected '42', got '%s'", b.String())
		}
	})

	t.Run("Convert float64", func(t *testing.T) {
		b := NumberLikeToBN(3.14)
		if b.String() != "3.14" {
			t.Errorf("Expected '3.14', got '%s'", b.String())
		}
	})

	t.Run("Convert string", func(t *testing.T) {
		b := NumberLikeToBN("123")
		if b.String() != "123" {
			t.Errorf("Expected '123', got '%s'", b.String())
		}
	})

	t.Run("Convert Big", func(t *testing.T) {
		original := BN(456)
		b := NumberLikeToBN(original)
		if b.String() != "456" {
			t.Errorf("Expected '456', got '%s'", b.String())
		}
	})
}

// Test arithmetic operations
func TestBigArithmetic(t *testing.T) {
	t.Run("Add operations", func(t *testing.T) {
		a := BN(10)

		// Test with different types
		result1 := a.Add(5)
		if result1.String() != "15" {
			t.Errorf("Expected '15', got '%s'", result1.String())
		}

		result2 := a.Add(3.5)
		if result2.String() != "13.5" {
			t.Errorf("Expected '13.5', got '%s'", result2.String())
		}

		result3 := a.Add("7")
		if result3.String() != "17" {
			t.Errorf("Expected '17', got '%s'", result3.String())
		}

		result4 := a.Add(BN(2))
		if result4.String() != "12" {
			t.Errorf("Expected '12', got '%s'", result4.String())
		}
	})

	t.Run("Sub operations", func(t *testing.T) {
		a := BN(20)

		result1 := a.Sub(5)
		if result1.String() != "15" {
			t.Errorf("Expected '15', got '%s'", result1.String())
		}

		result2 := a.Sub(3.5)
		if result2.String() != "16.5" {
			t.Errorf("Expected '16.5', got '%s'", result2.String())
		}

		result3 := a.Sub("7")
		if result3.String() != "13" {
			t.Errorf("Expected '13', got '%s'", result3.String())
		}
	})

	t.Run("Mul operations", func(t *testing.T) {
		a := BN(4)

		result1 := a.Mul(3)
		if result1.String() != "12" {
			t.Errorf("Expected '12', got '%s'", result1.String())
		}

		result2 := a.Mul(2.5)
		if result2.String() != "10" {
			t.Errorf("Expected '10', got '%s'", result2.String())
		}

		result3 := a.Mul("5")
		if result3.String() != "20" {
			t.Errorf("Expected '20', got '%s'", result3.String())
		}
	})

	t.Run("Div operations", func(t *testing.T) {
		a := BN(15)

		result1 := a.Div(3)
		if result1.String() != "5" {
			t.Errorf("Expected '5', got '%s'", result1.String())
		}

		result2 := a.Div(2.5)
		if result2.String() != "6" {
			t.Errorf("Expected '6', got '%s'", result2.String())
		}

		result3 := a.Div("5")
		if result3.String() != "3" {
			t.Errorf("Expected '3', got '%s'", result3.String())
		}
	})

	t.Run("Pow operations", func(t *testing.T) {
		a := BN(2)

		result1 := a.Pow(3)
		if result1.String() != "8" {
			t.Errorf("Expected '8', got '%s'", result1.String())
		}

		result2 := a.Pow(10)
		if result2.String() != "1024" {
			t.Errorf("Expected '1024', got '%s'", result2.String())
		}

		// Test negative exponent
		result3 := a.Pow(-2)
		if result3.String() != "0.25" {
			t.Errorf("Expected '0.25', got '%s'", result3.String())
		}
	})

	t.Run("Sqrt operations", func(t *testing.T) {
		a := BN(16)
		result := a.Sqrt()
		// Sqrt is not precise, so we check it's close to 4
		if result.String() != "4" {
			t.Errorf("Expected '4', got '%s'", result.String())
		}

		b := BN(9)
		result2 := b.Sqrt()
		if result2.String() != "3" {
			t.Errorf("Expected '3', got '%s'", result2.String())
		}
	})
}

// Test comparison methods
func TestBigComparison(t *testing.T) {
	t.Run("Cmp operations", func(t *testing.T) {
		a := BN(10)
		b := BN(5)
		c := BN(10)

		// Test Cmp with different types
		if a.Cmp(b) != 1 {
			t.Errorf("Expected 1 (a > b), got %d", a.Cmp(b))
		}

		if b.Cmp(a) != -1 {
			t.Errorf("Expected -1 (b < a), got %d", b.Cmp(a))
		}

		if a.Cmp(c) != 0 {
			t.Errorf("Expected 0 (a == c), got %d", a.Cmp(c))
		}

		// Test with different number types
		if a.Cmp(5) != 1 {
			t.Errorf("Expected 1 (a > 5), got %d", a.Cmp(5))
		}

		if a.Cmp(10.0) != 0 {
			t.Errorf("Expected 0 (a == 10.0), got %d", a.Cmp(10.0))
		}

		if a.Cmp("15") != -1 {
			t.Errorf("Expected -1 (a < '15'), got %d", a.Cmp("15"))
		}
	})

	t.Run("Equal operations", func(t *testing.T) {
		a := BN(10)

		if !a.Equal(10) {
			t.Error("Expected a.Equal(10) to be true")
		}

		if !a.Equal(10.0) {
			t.Error("Expected a.Equal(10.0) to be true")
		}

		if !a.Equal("10") {
			t.Error("Expected a.Equal('10') to be true")
		}

		if !a.Equal(BN(10)) {
			t.Error("Expected a.Equal(BN(10)) to be true")
		}

		if a.Equal(5) {
			t.Error("Expected a.Equal(5) to be false")
		}
	})

	t.Run("Gt operations", func(t *testing.T) {
		a := BN(10)

		if !a.Gt(5) {
			t.Error("Expected a.Gt(5) to be true")
		}

		if !a.Gt(9.9) {
			t.Error("Expected a.Gt(9.9) to be true")
		}

		if !a.Gt("9") {
			t.Error("Expected a.Gt('9') to be true")
		}

		if a.Gt(10) {
			t.Error("Expected a.Gt(10) to be false")
		}

		if a.Gt(15) {
			t.Error("Expected a.Gt(15) to be false")
		}
	})

	t.Run("Gte operations", func(t *testing.T) {
		a := BN(10)

		if !a.Gte(5) {
			t.Error("Expected a.Gte(5) to be true")
		}

		if !a.Gte(10) {
			t.Error("Expected a.Gte(10) to be true")
		}

		if !a.Gte(9.9) {
			t.Error("Expected a.Gte(9.9) to be true")
		}

		if a.Gte(15) {
			t.Error("Expected a.Gte(15) to be false")
		}
	})

	t.Run("Lt operations", func(t *testing.T) {
		a := BN(10)

		if !a.Lt(15) {
			t.Error("Expected a.Lt(15) to be true")
		}

		if !a.Lt(10.1) {
			t.Error("Expected a.Lt(10.1) to be true")
		}

		if !a.Lt("11") {
			t.Error("Expected a.Lt('11') to be true")
		}

		if a.Lt(10) {
			t.Error("Expected a.Lt(10) to be false")
		}

		if a.Lt(5) {
			t.Error("Expected a.Lt(5) to be false")
		}
	})

	t.Run("Lte operations", func(t *testing.T) {
		a := BN(10)

		if !a.Lte(15) {
			t.Error("Expected a.Lte(15) to be true")
		}

		if !a.Lte(10) {
			t.Error("Expected a.Lte(10) to be true")
		}

		if !a.Lte(10.1) {
			t.Error("Expected a.Lte(10.1) to be true")
		}

		if a.Lte(5) {
			t.Error("Expected a.Lte(5) to be false")
		}
	})
}

// Test rounding methods
func TestBigRounding(t *testing.T) {
	t.Run("Round operations", func(t *testing.T) {
		a := BN(3.14159)

		// Test Round with different decimal places
		result1 := a.Round(2)
		if result1.String() != "3.14" {
			t.Errorf("Expected '3.14', got '%s'", result1.String())
		}

		result2 := a.Round(3)
		if result2.String() != "3.142" {
			t.Errorf("Expected '3.142', got '%s'", result2.String())
		}

		result3 := a.Round(0)
		if result3.String() != "3" {
			t.Errorf("Expected '3', got '%s'", result3.String())
		}
	})

	t.Run("RoundCeil operations", func(t *testing.T) {
		a := BN(3.1)
		b := BN(-3.1)

		result1 := a.RoundCeil(0)
		if result1.String() != "4" {
			t.Errorf("Expected '4', got '%s'", result1.String())
		}

		result2 := b.RoundCeil(0)
		if result2.String() != "-3" {
			t.Errorf("Expected '-3', got '%s'", result2.String())
		}

		result3 := a.RoundCeil(1)
		if result3.String() != "3.1" {
			t.Errorf("Expected '3.1', got '%s'", result3.String())
		}
	})

	t.Run("RoundFloor operations", func(t *testing.T) {
		a := BN(3.9)
		b := BN(-3.9)

		result1 := a.RoundFloor(0)
		if result1.String() != "3" {
			t.Errorf("Expected '3', got '%s'", result1.String())
		}

		result2 := b.RoundFloor(0)
		if result2.String() != "-4" {
			t.Errorf("Expected '-4', got '%s'", result2.String())
		}

		result3 := a.RoundFloor(1)
		if result3.String() != "3.9" {
			t.Errorf("Expected '3.9', got '%s'", result3.String())
		}
	})

	t.Run("RoundInfinity operations", func(t *testing.T) {
		a := BN(3.1)
		b := BN(-3.1)

		result1 := a.RoundInfinity(0)
		if result1.String() != "4" {
			t.Errorf("Expected '4', got '%s'", result1.String())
		}

		result2 := b.RoundInfinity(0)
		if result2.String() != "-4" {
			t.Errorf("Expected '-4', got '%s'", result2.String())
		}
	})

	t.Run("RoundZero operations", func(t *testing.T) {
		a := BN(3.9)
		b := BN(-3.9)

		result1 := a.RoundZero(0)
		if result1.String() != "3" {
			t.Errorf("Expected '3', got '%s'", result1.String())
		}

		result2 := b.RoundZero(0)
		if result2.String() != "-3" {
			t.Errorf("Expected '-3', got '%s'", result2.String())
		}
	})
}

// Test conversion methods
func TestBigConversion(t *testing.T) {
	t.Run("String conversion", func(t *testing.T) {
		a := BN(123.456)
		if a.String() != "123.456" {
			t.Errorf("Expected '123.456', got '%s'", a.String())
		}

		b := BN(-42)
		if b.String() != "-42" {
			t.Errorf("Expected '-42', got '%s'", b.String())
		}

		// Test nil case
		var c *Big
		if c.String() != "" {
			t.Errorf("Expected empty string for nil Big, got '%s'", c.String())
		}
	})

	t.Run("Float64 conversion", func(t *testing.T) {
		a := BN(3.14159)
		f := a.Float64()
		if f != 3.14159 {
			t.Errorf("Expected 3.14159, got %f", f)
		}

		b := BN(42)
		f2 := b.Float64()
		if f2 != 42.0 {
			t.Errorf("Expected 42.0, got %f", f2)
		}
	})

	t.Run("BigFloat conversion", func(t *testing.T) {
		a := BN(3.14159)
		bf := a.BigFloat()
		if bf == nil {
			t.Error("Expected non-nil BigFloat")
		}

		// Convert back to string to verify
		str := bf.Text('f', -1)
		if str != "3.14159" {
			t.Errorf("Expected '3.14159', got '%s'", str)
		}
	})

	t.Run("BigInt conversion", func(t *testing.T) {
		a := BN(42)
		bi := a.BigInt()
		if bi.String() != "42" {
			t.Errorf("Expected '42', got '%s'", bi.String())
		}

		b := BN(3.7)
		bi2 := b.BigInt()
		if bi2.String() != "3" {
			t.Errorf("Expected '3', got '%s'", bi2.String())
		}
	})

	t.Run("BigInt rounding variants", func(t *testing.T) {
		a := BN(3.7)
		b := BN(-3.7)

		// Test BigIntRound
		if a.BigIntRound().String() != "4" {
			t.Errorf("Expected '4', got '%s'", a.BigIntRound().String())
		}

		// Test BigIntCeil
		if a.BigIntCeil().String() != "4" {
			t.Errorf("Expected '4', got '%s'", a.BigIntCeil().String())
		}
		if b.BigIntCeil().String() != "-3" {
			t.Errorf("Expected '-3', got '%s'", b.BigIntCeil().String())
		}

		// Test BigIntFloor
		if a.BigIntFloor().String() != "3" {
			t.Errorf("Expected '3', got '%s'", a.BigIntFloor().String())
		}
		if b.BigIntFloor().String() != "-4" {
			t.Errorf("Expected '-4', got '%s'", b.BigIntFloor().String())
		}

		// Test BigIntInfinity
		if a.BigIntInfinity().String() != "4" {
			t.Errorf("Expected '4', got '%s'", a.BigIntInfinity().String())
		}
		if b.BigIntInfinity().String() != "-4" {
			t.Errorf("Expected '-4', got '%s'", b.BigIntInfinity().String())
		}

		// Test BigIntZero
		if a.BigIntZero().String() != "3" {
			t.Errorf("Expected '3', got '%s'", a.BigIntZero().String())
		}
		if b.BigIntZero().String() != "-3" {
			t.Errorf("Expected '-3', got '%s'", b.BigIntZero().String())
		}
	})
}

// Test utility methods
func TestBigUtility(t *testing.T) {
	t.Run("Copy operations", func(t *testing.T) {
		original := BN(42.5)
		copied := original.Copy()

		// Test that copy is equal to original
		if !original.Equal(copied) {
			t.Error("Expected copied to equal original")
		}

		// Test that copy is independent
		original = original.Add(1) // Create new Big with Add
		if original.Equal(copied) {
			t.Error("Expected copy to be independent of original")
		}

		// Test BN() method
		bn := original.BN()
		if !original.Equal(bn) {
			t.Error("Expected BN() to return same value")
		}
	})

	t.Run("Abs operations", func(t *testing.T) {
		positive := BN(42)
		negative := BN(-42)
		zero := BN(0)

		// Test positive number
		abs1 := positive.Abs()
		if !abs1.Equal(42) {
			t.Errorf("Expected 42, got %s", abs1.String())
		}

		// Test negative number
		abs2 := negative.Abs()
		if !abs2.Equal(42) {
			t.Errorf("Expected 42, got %s", abs2.String())
		}

		// Test zero
		abs3 := zero.Abs()
		if !abs3.Equal(0) {
			t.Errorf("Expected 0, got %s", abs3.String())
		}
	})

	t.Run("MaxBN operations", func(t *testing.T) {
		a := BN(10)
		b := BN(20)
		c := BN(15)

		// Test MaxBN with different values
		max1 := MaxBN(a, b)
		if !max1.Equal(20) {
			t.Errorf("Expected 20, got %s", max1.String())
		}

		max2 := MaxBN(b, c)
		if !max2.Equal(20) {
			t.Errorf("Expected 20, got %s", max2.String())
		}

		max3 := MaxBN(a, c)
		if !max3.Equal(15) {
			t.Errorf("Expected 15, got %s", max3.String())
		}

		// Test with equal values
		max4 := MaxBN(a, a)
		if !max4.Equal(10) {
			t.Errorf("Expected 10, got %s", max4.String())
		}
	})

	t.Run("MinBN operations", func(t *testing.T) {
		a := BN(10)
		b := BN(20)
		c := BN(15)

		// Test MinBN with different values
		min1 := MinBN(a, b)
		if !min1.Equal(10) {
			t.Errorf("Expected 10, got %s", min1.String())
		}

		min2 := MinBN(b, c)
		if !min2.Equal(15) {
			t.Errorf("Expected 15, got %s", min2.String())
		}

		min3 := MinBN(a, c)
		if !min3.Equal(10) {
			t.Errorf("Expected 10, got %s", min3.String())
		}

		// Test with equal values
		min4 := MinBN(a, a)
		if !min4.Equal(10) {
			t.Errorf("Expected 10, got %s", min4.String())
		}
	})

	t.Run("BN0 constant", func(t *testing.T) {
		if !BN0.Equal(0) {
			t.Errorf("Expected BN0 to equal 0, got %s", BN0.String())
		}
	})
}
