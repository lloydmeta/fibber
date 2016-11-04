package fib

import (
	"math/big"
)

// Of returns the Fibonacci number at a given index
func Of(to uint64) *big.Int {
	if to == 0 {
		return big.NewInt(0)
	} else if to == 1 {
		return big.NewInt(1)
	} else {
		prev := big.NewInt(0)
		current := big.NewInt(1)
		var i uint64 = 1
		for ; i < to; i++ {
			temp := big.NewInt(0)
			temp.Add(prev, current)
			prev = current
			current = temp
		}
		return current
	}
}
