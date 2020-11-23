package membot

import (
	"fmt"
	"testing"
)

func TestService_Scan(t *testing.T) {
	m := New()
	data, _ := m.Scan("https://yandex.ru", 1)

	var got string

	for _, document := range data {
		got += fmt.Sprintf("%s ", document.Title)
	}

	if len(got) > 0 {
		got = got[:len(got)-1]
	}

	want := "Яндекс Google"
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}
}
