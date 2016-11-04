package fib

// Of returns the Fibonacci number at a given index
func Of(to uint64) uint64 {
	if to == 0 {
		return 0
	} else if to == 1 {
		return 1
	} else {
		var prev uint64
		var current uint64 = 1
		var i uint64 = 1
		for ; i < to; i++ {
			temp := current + prev
			prev = current
			current = temp
		}
		return current
	}
}
