package main

import (
	"fmt"
	"lesson02/crawler/pkg/spider"
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

	var array []page
	for url, title := range data {
		p := page{url: url, title: title}
		array = append(array, p)

		fmt.Printf("Страница %s имеет адрес: %s\n", url, title)
	}

	// w := flag.String("w", "", "Target word")
	// flag.Parse()
	// word := strings.ToLower(*w) // Обязательная часть

	for {
		fmt.Printf("Enter word or quit for exit: ")

		var w string
		fmt.Scanln(&w)
		word := strings.ToLower(w)

		if word == "quit" {
			fmt.Printf("Exit!\n")
			return
		}

		for _, p := range array {
			url := strings.ToLower(p.url)
			title := strings.ToLower(p.title)

			if strings.Contains(url, word) || strings.Contains(title, word) {
				fmt.Printf("%s - %s\n", p.url, p.title)
			}
		}

		fmt.Printf("\n")
	}

}
