package fib

import (
	"fmt"
	"math/big"
	"testing"
)

var fibSeq = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811}
var fibBigInts = make([]*big.Int, len(fibSeq))

func init() {
	for idx, i := range fibSeq {
		fibBigInts[idx] = big.NewInt(int64(i))
	}
}

func TestFibOf(t *testing.T) {
	checkFibAt(t, func(idx int) *big.Int { return Of(uint(idx)) })
}

func TestFibForEach(t *testing.T) {
	// Create an empty array that we can fill-in in our callback
	forEachedGenerated := make([]*big.Int, len(fibSeq))
	currentIdx := 0
	ForEach(uint(len(fibSeq)-1), func(i *big.Int) {
		forEachedGenerated[currentIdx] = i
		currentIdx++
	})
	checkFibAt(t, func(idx int) *big.Int { return forEachedGenerated[idx] })
}

func TestMemoed(t *testing.T) {
	fibGen := NewMemoed()
	checkFibAt(t, func(idx int) *big.Int { return fibGen.Of(uint(idx)) })
	if fib50 := fibGen.Of(50); fib50.Cmp(big.NewInt(12586269025)) != 0 {
		err := fmt.Sprintf("Fib of 50 was expected to be %v but got %v", 12586269025, fibGen.Of(50))
		t.Error(err)
	}
}

func TestMemoedbBig(t *testing.T) {
	fibGen := NewMemoed()
	if fibGen.Of(1000).Cmp(Of(1000)) != 0 {
		err := fmt.Sprintf("Fib of 1000 was expected to be %v but got %v", Of(1000), fibGen.Of(1000))
		t.Error(err)
	}
}

func TestMemoedConcurrent(t *testing.T) {
	numRoutines := 50
	fibTo := 100
	fibGen := NewMemoed()
	type pair struct {
		idx uint
		fib *big.Int
	}
	fibChannel := make(chan pair)
	for i := 0; i < numRoutines; i++ {
		go func(c chan<- pair) {
			for j := 0; j < fibTo; j++ {
				u := uint(j)
				g := fibGen.Of(u)
				c <- pair{uint(u), g}
			}
		}(fibChannel)
	}

	// numRoutines*fibTo times from the fibChannel and compare observed w/ expected each time
	for i := 0; i < numRoutines*fibTo; i++ {
		thePair := <-fibChannel
		if observed, expected := thePair.fib, Of(thePair.idx); observed.Cmp(expected) != 0 {
			err := fmt.Sprintf("Expected %d for Fib(%d) but got %d", expected, thePair.idx, observed)
			t.Error(err)
		}
	}

}

// Helper function to DRY up testing
func checkFibAt(t *testing.T, getFib func(int) *big.Int) {
	for idx, expected := range fibBigInts {
		observed := getFib(idx)
		if expected.Cmp(observed) != 0 {
			err := fmt.Sprintf("Expected %d for Fib(%d) but got %d", expected, idx, observed)
			t.Error(err)
		}
	}
}

func ExampleOf() {
	fib100 := Of(100)
	fmt.Printf("Fib 100 is %d", fib100)
}

func ExampleForEach() {
	forEachedGenerated := make([]*big.Int, len(fibSeq))
	currentIdx := 0
	ForEach(uint(len(fibSeq)-1), func(i *big.Int) {
		forEachedGenerated[currentIdx] = i
		currentIdx++
	})
	fmt.Printf("%v", forEachedGenerated)
}

func ExampleMemoed() {
	memoed := NewMemoed()
	fmt.Printf("Fib 100 is %d", memoed.Of(100))
}

func BenchmarkFib30(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Of(30)
	}
}

func BenchmarkMemoedFib30(b *testing.B) {
	memoed := NewMemoed()
	for n := 0; n < b.N; n++ {
		memoed.Of(30)
	}
}
