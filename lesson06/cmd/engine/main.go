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

type fileError struct {
	s string
}

func (e *fileError) Error() string {
	return e.s
}

const dataFile = "./data.json"

func main() {
	ind := index.New()
	data, err := readFromFile(dataFile)
	if err != nil {
		if err, ok := err.(*fileError); ok {
		} else {
			log.Fatal(err)
		}
	}
	if err == nil {
		for url, title := range data {
			ind.Add(url, title)
		}
	}

	const url = "https://habr.com"
	go scanSite(ind, url, 1)

	for {
		fmt.Print("Enter word or quit for exit: ")

		var w string
		fmt.Scanln(&w)
		w = strings.ToLower(w)

		if w == "quit" {
			allData := ind.AllData()
			err := writeToFile(dataFile, allData)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Exit!")
			return
		}

		ids := ind.IDs(w)
		fmt.Printf("Found IDs: %v\n", ids)

		pages := ind.Pages(ids)
		fmt.Printf("%+v\n", pages)
	}
}

func readFromFile(filename string) (data map[string]string, err error) {
	if !fileExists(filename) {
		return nil, &fileError{"not found"}
	}

	byteValue, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data = make(map[string]string)
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func scanSite(ind *index.Index, url string, depth int) {
	data, err := spider.Scan(url, depth)
	if err != nil {
		log.Fatal(err)
	}

	for url, title := range data {
		ind.Add(url, title)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func writeToFile(filename string, data map[string]string) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
