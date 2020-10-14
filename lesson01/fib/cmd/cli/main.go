package main

import (
	"flag"
	"fmt"
	"lesson01/fib/pkg/fib"
	"os"
)

const minNumber int = 0

func main() {
	number := flag.Int("n", -1, "Fib number")
	flag.Parse()

	if *number <= minNumber {
		fmt.Println("Ivalid number", *number, "Number must be greater than or equal 1")
		os.Exit(2)
	}

	result, err := fib.Calc(*number)

	if err != nil {
		panic(err)
	}

	fmt.Println("Fib", *number, "is eq:", result)
	return
}
