package fib

import (
	"fmt"
	"testing"
)

var fibSeq = []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811}

func TestFibOf(t *testing.T) {
	for idx, expected := range fibSeq {
		observed := Of(uint64(idx))
		if observed != expected {
			err := fmt.Sprintf("Expected %d for Fib(%d) but got %d", expected, idx, observed)
			t.Error(err)
		}
	}
}

func ExampleOf() {
	Of(100)
}

func BenchmarkFib30(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Of(30)
	}
}
