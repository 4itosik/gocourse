package main

import (
	"fmt"
	"lesson04/crawler/pkg/index"
	"lesson04/crawler/pkg/spider"
	"strings"
)

type page struct {
	url   string
	title string
}

func main() {
	const url = "https://habr.com"
	data, err := spider.Scan(url, 2)
	if err != nil {
		fmt.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
	}

	for url, title := range data {
		index.Add(url, title)
	}

	for {
		fmt.Print("Enter word or quit for exit: ")

		var w string
		fmt.Scanln(&w)
		w = strings.ToLower(w)

		if w == "quit" {
			fmt.Println("Exit!")
			return
		}

		ids := index.IDs(w)
		fmt.Printf("Found IDs: %v\n", ids)

		pages := index.Pages(ids)
		fmt.Printf("%+v\n", pages)
	}
}
