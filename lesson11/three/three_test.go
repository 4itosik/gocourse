package three

import (
	"bytes"
	"testing"
)

func Test_printAll(t *testing.T) {
	a := "str1"
	i := 3
	b := false
	c := "str2"
	d := "str3"

	var buf bytes.Buffer
	printAll(&buf, a, i, b, c, d)
	got := buf.String()
	want := "str1str2str3"
	if got != want {
		t.Fatalf("получили %v, ожидалось %v", got, want)
	}
}
