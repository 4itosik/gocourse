package hash

import (
	"fmt"
	"lesson07/pkg/crawler"
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

func TestIndex_Docs(t *testing.T) {
	i := New()
	addItem(i)

	ids := []int{1}
	p := i.Docs(ids)[0]
	got := p.Title
	want := "Страница 1"
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}
}

func TestIndex_AllDocs(t *testing.T) {
	i := New()
	addItem(i)

	data := i.AllDocs()
	var got string
	for _, doc := range data {
		got += fmt.Sprintf("%s ", doc.URL)
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
	doc1 := crawler.Document{
		URL:   "http://page1.ru",
		Title: "Страница 1",
	}
	i.Add(doc1)

	doc2 := crawler.Document{
		URL:   "http://page2.ru",
		Title: "Страница 2",
	}
	i.Add(doc2)

	doc3 := crawler.Document{
		URL:   "http://page3.ru",
		Title: "Страница 3",
	}
	i.Add(doc3)
}
