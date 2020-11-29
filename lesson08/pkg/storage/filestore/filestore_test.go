package filestore

import (
	"flag"
	"fmt"
	"lesson08/pkg/crawler"
	"os"
	"testing"
)

var filename = "./fakeData.json"
var docs = []crawler.Document{
	{
		URL:   "https://yandex.ru",
		Title: "Яндекс",
	},
	{
		URL:   "https://google.ru",
		Title: "Google",
	},
}

func TestMain(m *testing.M) {
	flag.Parse()
	if !testing.Short() {
		fmt.Println("skip test. Run with -short")
		return
	}

	_ = os.Remove(filename)
	exitVal := m.Run()
	_ = os.Remove(filename)
	os.Exit(exitVal)
}

func TestFilestore_WriteAndReadData(t *testing.T) {
	if !testing.Short() {
		t.Skip("skip test. Run with -short")
	}

	f := New(filename)
	err := f.WriteData(docs)
	if err != nil {
		t.Errorf("c.WriteData(); err = %s; want nil", err)
		return
	}

	data, err := f.ReadData()
	if err != nil {
		t.Errorf("c.ReadData(); err = %s; want nil", err)
		return
	}

	var got string
	for _, doc := range data {
		got += fmt.Sprintf("%s ", doc.Title)
	}
	if len(got) > 0 {
		got = got[:len(got)-1]
	}

	want := "Яндекс Google"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
