package two

import "testing"

func Test_maxAgePerson(t *testing.T) {
	e1 := employee{age: 21}
	e2 := employee{age: 25}
	c1 := customer{age: 31}
	c2 := customer{age: 33}

	got := maxAgePerson(e1, e2, c1, c2)
	want := c2
	if got != want {
		t.Fatalf("получили %v, ожидалось %v", got, want)
	}
}
