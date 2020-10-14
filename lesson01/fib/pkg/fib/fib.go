package fib

// Calc - culate fibonacci count
func Calc(n int) int {
	a := 0
	b := 1
	// Iterate until desired position in sequence.
	for i := 0; i < n; i++ {
		// Use temporary variable to swap values.
		temp := a
		a = b
		b = temp + a
	}
	return a
}
