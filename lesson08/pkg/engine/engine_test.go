package engine

import (
	"fmt"
	"lesson08/pkg/index/memory"
	"testing"
)

func TestEngine_Search(t *testing.T) {
	ind := memory.New()
	e := New(ind)

	data := e.Search("query")

	var got string
	for _, doc := range data {
		got += fmt.Sprintf("%s ", doc.Title)
	}
	if len(got) > 0 {
		got = got[:len(got)-1]
	}

	want := "Яндекс"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
