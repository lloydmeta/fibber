package main

import (
	"fmt"
	"strconv"

	"github.com/lloydmeta/fib"
)

func main() {
	fmt.Println("Fib to which number? ")
	var input string
	fmt.Scanln(&input)

	number, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		fmt.Printf("Invalid number: %s\n", input)
	} else {
		fib := fib.Of(number)
		fmt.Printf("Fib of %d is %d\n", number, fib)
	}
}
