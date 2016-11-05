package fib

import (
	"math/big"
	"sync"
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

// Memoed is a thread-safe Fibonacci number provider that has a cache
type Memoed struct {
	lock  sync.Mutex
	cache []*big.Int
}

// NewMemoed creates and returns a new Memoed instance
func NewMemoed() *Memoed {
	return &Memoed{sync.Mutex{}, make([]*big.Int, 1000)}
}

// Of returns the Fibonacci number at a given index
func (self *Memoed) Of(to uint64) *big.Int {
	toInt := int(to)
	self.lock.Lock()
	defer self.lock.Unlock()
	// ensure that `to` is not bigger than current cache
	if len(self.cache) < int(to) {
		newSlice := make([]*big.Int, toInt+1)
		self.cache = append(newSlice, self.cache...)
	}
	if cached := self.cache[toInt]; cached != nil {
		return cached
	}
	currentIdx := toInt - 1
	stack := &intList{toInt, nil} // start off at ToInt
	for ; currentIdx >= 0 && self.cache[currentIdx] == nil; currentIdx-- {
		stack = stack.Prepend(currentIdx)
	}
	// unwind the stack
	for stack != nil {
		idx, tail := stack.Pop()
		stack = tail
		if idx <= 1 {
			self.cache[idx] = big.NewInt(int64(idx))
			continue
		}
		self.cache[idx] = big.NewInt(0).Add(self.cache[idx-1], self.cache[idx-2])
	}
	return self.cache[to]
}
