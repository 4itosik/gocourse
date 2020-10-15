package fib

import (
	"testing"
)

func Test_Calc(t *testing.T) {
	want := 13
	got := Calc(7)

	if want != got {
		t.Errorf("Got = %d; want %d", got, want)
	}
}
