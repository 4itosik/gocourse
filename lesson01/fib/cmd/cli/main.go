package main

import (
	"flag"
	"fmt"
	"lesson01/fib/pkg/fib"
	"os"
)

const minNumber int = 0
const maxNumber int = 20

func main() {
	number := flag.Int("n", -1, "Fib number")
	flag.Parse()

	if *number <= minNumber {
		fmt.Println("Ivalid number", *number, "Number must be greater than or equal 1")
		os.Exit(2)
	} else if *number > maxNumber {
		fmt.Println("Ivalid number", *number, "Number should be less 20")
		os.Exit(2)
	}

	result := fib.Calc(*number)
	fmt.Printf("Fib %v is eq: %v", *number, result)
	return
}
