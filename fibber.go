package fib

import (
	"math/big"
	"sync"
)

// Of returns the Fibonacci number at a given index
func Of(to uint) *big.Int {
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
func ForEach(to uint, callback func(*big.Int)) {
	prev := big.NewInt(0)
	current := big.NewInt(1)
	var i uint
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
	lock              sync.RWMutex
	cache             []*big.Int
	cacheGrowthFactor int
}

// NewMemoed creates and returns a new Memoed instance with sensible defaults
//
// The cache length is 100 and its growth factor is 3
func NewMemoed() *Memoed {
	return &Memoed{sync.RWMutex{}, make([]*big.Int, 100), 3}
}

// Of returns the Fibonacci number at a given index
//
// Internally uses the cache Memoed's private cache and is stack-safe and thread-safe
func (self *Memoed) Of(to uint) *big.Int {
	toInt := int(to)

	// First try a read with just a plain Read lock
	self.lock.RLock()
	// Note that this is somewhat repeated below but because of the use of RWLock,
	// trying to DRY this out doesn't buy much.
	if len(self.cache) > toInt && self.cache[toInt] != nil {
		// we put it in a temporary variable instead of returning directly from
		// the array to avoid having to use defer (to avoid raciness) because defer adds ~70ns
		existing := self.cache[toInt]
		self.lock.RUnlock()
		return existing
	}
	self.lock.RUnlock() // Plain read failed, so unlock

	// Lock for writing
	self.lock.Lock()
	// Try another read in case another thread wrote into the cache whilst we
	// were acquiring the lock.
	//
	// Note there is no need to acquire a RLock again because we have a WLock.
	// In fact, Golang's RLock blocks until all WLocks are released, even within the same
	// goroutine, so trying to get a RLock here will deadlock.
	if len(self.cache) > toInt && self.cache[toInt] != nil {
		// See above note on avoiding defer
		existing := self.cache[toInt]
		self.lock.Unlock()
		return existing
	}

	// Ensure that `to` is not bigger than or equal to  current cache
	// because we want to access the ith index of the cache (len of [] must be i + 1
	// if we want to access or set [i])
	//
	// If it is though, create a new cache slice based on the growth factor and
	//  copy the old members to the new cache
	if len(self.cache) <= toInt {
		newSlice := make([]*big.Int, toInt*self.cacheGrowthFactor)
		copy(newSlice, self.cache)
		self.cache = newSlice
	}

	// Next, we walk down the cache until we find a cached item, or we reach index 0
	// For every index that does not have a cached item, we add it into a stack to
	// so we can fill in the cache later.
	stack := &intList{toInt, nil} // Start off our stack at ToInt
	currentIdx := toInt - 1
	for ; currentIdx >= 0 && self.cache[currentIdx] == nil; currentIdx-- {
		stack = stack.Prepend(currentIdx)
	}

	// Unwind the stack by popping until tail is nil.
	for stack != nil {
		idx, tail := stack.Pop()
		stack = tail
		if idx <= 1 {
			self.cache[idx] = big.NewInt(int64(idx))
			continue
		}
		self.cache[idx] = big.NewInt(0).Add(self.cache[idx-1], self.cache[idx-2])
	}
	// See above note on avoiding defer
	existing := self.cache[toInt]
	self.lock.Unlock()
	return existing
}
