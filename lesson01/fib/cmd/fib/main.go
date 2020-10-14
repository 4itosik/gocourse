package main

import (
	"fmt"
	"lesson01/fib/pkg/fib"
)

func main() {
	result1 := fib.Calc(3)
	fmt.Println("Fib 3 is eq:", result1)

	result2 := fib.Calc(7)
	fmt.Println("Fib 7 is eq:", result2)

	result3 := fib.Calc(15)
	fmt.Println("Fib 15 is eq:", result3)
}
