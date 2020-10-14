package fib

import "errors"

// Calc - culate fibonacci count
func Calc(n int) (int, error) {
	if n > 20 {
		return -1, errors.New("Number should be more 20")
	}

	a := 0
	b := 1
	// Iterate until desired position in sequence.
	for i := 0; i < n; i++ {
		// Use temporary variable to swap values.
		temp := a
		a = b
		b = temp + a
	}
	return a, nil
}
