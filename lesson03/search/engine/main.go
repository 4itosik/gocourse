package main

import (
	"fmt"
	"lesson03/crawler/pkg/spider"
	"strings"
)

type page struct {
	url   string
	title string
}

type scanner interface {
	Scan(string, int) (map[string]string, error)
}

func main() {
	const url = "https://habr.com"
	c := spider.New()
	data, err := scan(c, url, 1)
	if err != nil {
		fmt.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
	}

	var pages []page
	for url, title := range data {
		p := page{url: url, title: title}
		pages = append(pages, p)

		fmt.Printf("Страница %s имеет адрес: %s\n", url, title)
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

		for _, p := range pages {
			url := strings.ToLower(p.url)
			title := strings.ToLower(p.title)

			if strings.Contains(url, w) || strings.Contains(title, w) {
				fmt.Printf("%s - %s\n", p.url, p.title)
			}
		}

		fmt.Print("\n")
	}

}

func scan(s scanner, url string, depth int) (data map[string]string, err error) {
	return s.Scan(url, depth)
}
