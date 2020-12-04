package one

import "testing"

func Test_maxYears(t *testing.T) {
	e1 := employee{age: 21}
	e2 := employee{age: 25}
	c1 := customer{age: 31}
	c2 := customer{age: 33}

	got := maxYears(e1, e2, c1, c2)
	want := 33
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}
