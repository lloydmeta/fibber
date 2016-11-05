package fib

import (
	"fmt"
	"math/big"
	"testing"
)

var fibSeq = []int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811}
var fibBigInts = make([]*big.Int, len(fibSeq))

func init() {
	for idx, i := range fibSeq {
		fibBigInts[idx] = big.NewInt(i)
	}
}

func TestFibOf(t *testing.T) {
	checkFibAt(t, func(idx int) *big.Int { return Of(uint64(idx)) })
}

func TestFibForEach(t *testing.T) {
	// Create an empty array that we can fill-in in our callback
	forEachedGenerated := make([]*big.Int, len(fibSeq))
	currentIdx := 0
	ForEach(uint64(len(fibSeq)-1), func(i *big.Int) {
		forEachedGenerated[currentIdx] = i
		currentIdx++
	})
	checkFibAt(t, func(idx int) *big.Int { return forEachedGenerated[idx] })
}

func TestCachedFilb(t *testing.T) {
	fibGen := NewCachedFib()
	checkFibAt(t, func(idx int) *big.Int { return fibGen.Of(uint64(idx)) })
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
	ForEach(uint64(len(fibSeq)-1), func(i *big.Int) {
		forEachedGenerated[currentIdx] = i
		currentIdx++
	})
	fmt.Printf("%v", forEachedGenerated)
}

func BenchmarkFib30(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Of(30)
	}
}
