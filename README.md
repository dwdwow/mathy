# mathy

A comprehensive Go library providing high-precision decimal arithmetic and mathematical utilities that extend beyond Go's standard `math` package.

## Features

- **High-Precision Decimal Arithmetic**: Built on `github.com/shopspring/decimal` for accurate decimal calculations
- **Big Number Operations**: Complete set of arithmetic, comparison, and conversion operations
- **Rounding Functions**: Multiple rounding modes (round, ceil, floor, infinity, zero)
- **Random Number Generation**: Float random number utilities
- **Type-Safe API**: Generic functions with compile-time type checking

## Installation

```bash
go get github.com/dwdwow/mathy
```

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/dwdwow/mathy"
)

func main() {
    // Create a Big number
    a := mathy.BN(3.14159)
    b := mathy.BN("2.5")
    
    // Arithmetic operations
    sum := a.Add(b)           // 5.64159
    product := a.Mul(b)        // 7.853975
    quotient := a.Div(b)       // 1.256636
    
    // Comparison
    if a.Gt(b) {
        fmt.Println("a is greater than b")
    }
    
    // Rounding
    rounded := a.Round(2)     // 3.14
    ceiling := a.RoundCeil(1)  // 3.2
    
    fmt.Println(sum.String())
}
```

### Rounding Functions

```go
// Direct rounding functions for float64
result := mathy.Round(3.14159, 2)        // 3.14
result = mathy.RoundCeil(3.1, 0)          // 4.0
result = mathy.RoundFloor(3.9, 0)        // 3.0
result = mathy.RoundInfinity(3.1, 0)     // 4.0 (away from zero)
result = mathy.RoundZero(3.9, 0)         // 3.0 (towards zero)
```

### Random Number Generation

```go
// Generate random float in range [min, max)
random := mathy.RandFloat(0.0, 10.0)

// Generate multiple random floats
randoms := mathy.RandFloats(1.0, 5.0, 10) // 10 floats in [1.0, 5.0)
```

## API Reference

### Big Number Operations

#### Construction

- `BN(value)` - Create a Big number from int, int64, float64, string, or *Big
- `NumberLikeToBN(value)` - Convert any NumberLike to *Big

#### Arithmetic

- `Add(value)` - Addition
- `Sub(value)` - Subtraction  
- `Mul(value)` - Multiplication
- `Div(value)` - Division
- `Pow(exponent)` - Power (exponent must be integer)
- `Sqrt()` - Square root (approximate)

#### Comparison

- `Cmp(value)` - Compare (returns -1, 0, or 1)
- `Equal(value)` - Equality check
- `Gt(value)` - Greater than
- `Gte(value)` - Greater than or equal
- `Lt(value)` - Less than
- `Lte(value)` - Less than or equal

#### Rounding

- `Round(places)` - Standard rounding
- `RoundCeil(places)` - Round up
- `RoundFloor(places)` - Round down
- `RoundInfinity(places)` - Round away from zero
- `RoundZero(places)` - Round towards zero

#### Conversion

- `String()` - Convert to string
- `Float64()` - Convert to float64
- `BigFloat()` - Convert to *big.Float
- `BigInt()` - Convert to *big.Int
- `BigIntRound()` - Convert with rounding
- `BigIntCeil()` - Convert with ceiling
- `BigIntFloor()` - Convert with floor
- `BigIntInfinity()` - Convert away from zero
- `BigIntZero()` - Convert towards zero

#### Utility

- `Copy()` - Create independent copy
- `Abs()` - Absolute value
- `BN()` - Return self (for chaining)

### Global Functions

#### Rounding

- `Round(f, places)` - Round float64
- `RoundCeil(f, places)` - Ceiling for float64
- `RoundFloor(f, places)` - Floor for float64
- `RoundInfinity(f, places)` - Away from zero for float64
- `RoundZero(f, places)` - Towards zero for float64

#### Random

- `RandFloat(min, max)` - Random float in [min, max)
- `RandFloats(min, max, n)` - n random floats in [min, max)

#### Utility

- `MaxBN(a, b)` - Maximum of two Big numbers
- `MinBN(a, b)` - Minimum of two Big numbers

## Examples

### Financial Calculations

```go
// Calculate compound interest with high precision
principal := mathy.BN(1000)
rate := mathy.BN(0.05)  // 5%
years := mathy.BN(10)

// A = P(1 + r)^t
amount := principal.Mul(rate.Add(1).Pow(years))
fmt.Println(amount.String()) // "1628.894626777442"
```

### Statistical Operations

```go
// Calculate mean with precision
numbers := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
sum := mathy.BN(0)
for _, n := range numbers {
    sum = sum.Add(n)
}
mean := sum.Div(len(numbers))
fmt.Println(mean.String()) // "3.3"
```

### Rounding Strategies

```go
price := mathy.BN(19.99)

// Different rounding for different contexts
tax := price.Mul(0.08).RoundCeil(2)        // Tax: round up
discount := price.Mul(0.1).RoundFloor(2)  // Discount: round down
final := price.Add(tax).Sub(discount)     // Final price
```

## Testing

The library includes comprehensive tests with 84.5% code coverage:

```bash
go test -v
go test -cover
```

## Performance

- Uses `github.com/shopspring/decimal` for high-precision arithmetic
- Division precision set to 100 decimal places by default
- Optimized for financial and scientific calculations

## License

MIT License - see LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
