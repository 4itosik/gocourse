package main

import (
	"fmt"
	"lesson05/pkg/index"
	"lesson05/pkg/spider"
	"strings"
)

func main() {
	const url = "https://habr.com"
	data, err := spider.Scan(url, 1)
	if err != nil {
		fmt.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
	}

	index := index.New()
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
