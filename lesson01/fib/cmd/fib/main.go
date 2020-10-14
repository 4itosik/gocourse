package main

import (
	"fmt"
	"lesson01/fib/pkg/fib"
)

func main() {
	result1, err := fib.Calc(25)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Fib 3 is eq:", result1)
	}

	result2, err := fib.Calc(7)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Fib 7 is eq:", result2)
	}

	result3, err := fib.Calc(15)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Fib 15 is eq:", result3)
	}
}
