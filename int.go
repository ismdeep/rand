package rand

import (
	"crypto/rand"
	"math"
	"math/big"
)

// Int returns, as an int
func Int() int {
	return Intn(math.MaxInt)
}

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Intn(n int) int {
	maxVal := big.NewInt(int64(n))
	v, _ := rand.Int(rand.Reader, maxVal)
	return int(v.Int64())
}

// IntBetween returns, as an int, a random number in [a,b]
func IntBetween(left int, right int) int {
	v := Intn(right - left + 1)
	return v + left
}
