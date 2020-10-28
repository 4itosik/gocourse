package main

import (
	"lesson03/crawler/pkg/fakespider"
	"testing"
)

func Test_scan(t *testing.T) {
	const url = "https://habr.com"
	c := fakespider.New()
	data, _ := scan(c, url, 1)

	want := 3
	got := len(data)

	if want != got {
		t.Errorf("Got = %d; want %d", got, want)
	}
}
