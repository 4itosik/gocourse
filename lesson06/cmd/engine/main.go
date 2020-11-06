package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lesson06/pkg/index"
	"lesson06/pkg/spider"
	"log"
	"os"
	"strings"
)

const dataFile = "./data.json"

func main() {
	ind := index.New()
	readFromFile(ind)

	const url = "https://habr.com"
	go crawlerSite(ind, url, 2)

	for {
		fmt.Print("Enter word or quit for exit: ")

		var w string
		fmt.Scanln(&w)
		w = strings.ToLower(w)

		if w == "quit" {
			fmt.Println("Exit!")
			return
		}

		ids := ind.IDs(w)
		fmt.Printf("Found IDs: %v\n", ids)

		pages := ind.Pages(ids)
		fmt.Printf("%+v\n", pages)
	}
}

func readFromFile(ind *index.Index) {
	if !fileExists(dataFile) {
		return
	}

	byteValue, err := ioutil.ReadFile(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]string)
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatal(err)
	}

	for url, title := range data {
		ind.Add(url, title)
	}
}

func crawlerSite(ind *index.Index, url string, depth int) {
	data, err := spider.Scan(url, depth)
	if err != nil {
		fmt.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
	}

	ind.Clear()
	for url, title := range data {
		ind.Add(url, title)
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(dataFile, bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
