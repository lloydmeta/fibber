package main

import (
	"flag"
	"fmt"
	"math/big"
	"strconv"

	"github.com/lloydmeta/fibber"
)

func main() {
	flag.Parse()

	directIdx := flag.Arg(0)

	if directIdx == "" {
		interactiveMode()
	} else {
		number := parseToUInt(directIdx)
		fib := fib.Of(number)
		printFib(number, fib)
	}
}

func parseToUInt(s string) uint {
	number, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Invalid number: %s\n", s))
	} else {
		return uint(number)
	}
}

func interactiveMode() {
	fmt.Println("\n**************** Interactive Mode ********************")
	fmt.Println(" ctrl+c or ctrl+d or enter any invalid number to exit ")
	fmt.Println("******************************************************")
	memoed := fib.NewMemoed()
	for true {
		fmt.Println("\nWhich Fibonacci number would you like to know? ")
		var input string
		fmt.Scanln(&input)
		number := parseToUInt(input)
		fib := memoed.Of(number)
		printFib(number, fib)
	}
}

func printFib(of uint, fib *big.Int) {
	fmt.Printf("Fibonacci[%d] is %d\n", of, fib)
}
