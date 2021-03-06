package index

import (
	"fmt"
	"testing"
)

func TestIndex_Add(t *testing.T) {
	i := New()
	addItem(i)

	got := i.String()
	want := "1 2 3"
	if got != want {
		t.Fatalf("получили %v, ожидалось %v", got, want)
	}
}

func TestIndex_IDs(t *testing.T) {
	i := New()
	addItem(i)

	got := i.IDs("1")[0]
	want := 1
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}

func TestIndex_Pages(t *testing.T) {
	i := New()
	addItem(i)

	ids := []int{1}
	p := i.Pages(ids)[0]
	got := p.title
	want := "Страница 1"
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}
}

func TestIndex_AllData(t *testing.T) {
	i := New()
	addItem(i)

	data := i.AllData()
	var got string
	for url := range data {
		got += fmt.Sprintf("%s ", url)
	}
	if len(got) > 0 {
		got = got[:len(got)-1]
	}

	want := "http://page1.ru http://page2.ru http://page3.ru"
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}
}

func addItem(i *Index) {
	i.Add("http://page1.ru", "Страница 1")
	i.Add("http://page2.ru", "Страница 2")
	i.Add("http://page3.ru", "Страница 3")
}
