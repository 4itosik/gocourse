package main

import (
	"fmt"
	"lesson03/crawler/pkg/fakespider"
	"log"
)

type Scanner interface {
	Scan(string, int) (map[string]string, error)
}

type crawler struct{}

func (c crawler) Scan(url string, depth int) (data map[string]string, err error) {
	return fakespider.Scan(url, depth)
}

func main() {
	const url = "https://habr.com"

	var sc Scanner
	sc = crawler{}
	data, err := sc.Scan(url, 1)
	if err != nil {
		log.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
	}

	for k, v := range data {
		fmt.Printf("Страница %s имеет адрес: %s\n", v, k)
	}
}
