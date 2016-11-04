package fib

import (
	"math/big"
)

// Of returns the Fibonacci number at a given index
func Of(to uint64) *big.Int {
	var toReturn *big.Int
	ForEach(to, func(i *big.Int) { toReturn = i })
	return toReturn
}

// ForEach allows you to pass a callback function that gets
// invoked for every number in the fibonacci sequence up to
// the "to" you specify
//
// Useful for instance when you want to efficiently create a
// memoised of fibonacci numbers
func ForEach(to uint64, callback func(*big.Int)) {
	prev := big.NewInt(0)
	current := big.NewInt(1)
	var i uint64
	for ; i <= to; i++ {
		callback(prev)
		sum := big.NewInt(0)
		sum.Add(prev, current)
		prev = current
		current = sum
	}
}
