package main

import (
	"flag"
	"fmt"
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
		printFib(number)
	}
}

func parseToUInt(s string) uint64 {
	number, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Invalid number: %s\n", s))
	} else {
		return number
	}
}

func interactiveMode() {
	fmt.Println("Fib to which number? ")
	var input string
	fmt.Scanln(&input)
	number := parseToUInt(input)
	printFib(number)
}

func printFib(fibOf uint64) {
	fib := fib.Of(fibOf)
	fmt.Printf("Fib of %d is %d\n", fibOf, fib)
}
