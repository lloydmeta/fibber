package fib

import (
	"math/big"
)

// Of returns the Fibonacci number at a given index
func Of(to uint64) *big.Int {
	if to <= 1 {
		return big.NewInt(int64(to))
	}
	prev := big.NewInt(0)
	current := big.NewInt(1)
	var i uint64 = 2
	for ; i <= to; i++ {
		sum := big.NewInt(0)
		sum.Add(prev, current)
		prev = current
		current = sum
	}
	return current
}
