package main

import (
	"fmt"
	"lesson08/pkg/crawler"
	"lesson08/pkg/crawler/spider"
	"lesson08/pkg/engine"
	"lesson08/pkg/index"
	"lesson08/pkg/index/hash"
	"lesson08/pkg/storage"
	"lesson08/pkg/storage/filestore"
	"log"
	"strings"
)

const dataFile = "./data.json"

type gosearch struct {
	scanner crawler.Interface
	storage storage.Interface
	index   index.Interface
	engine  *engine.Service
	url     string
	depth   int
}

func main() {
	server := new()
	server.init()
	server.run()
}

func new() *gosearch {
	gs := gosearch{}
	gs.scanner = spider.New()
	gs.storage = filestore.New(dataFile)
	gs.index = hash.New()
	gs.engine = engine.New(gs.index)
	gs.url = "https://go.dev"
	gs.depth = 2
	return &gs
}

func (gs *gosearch) init() {
	fileData, err := gs.storage.ReadData()
	if err != nil {
		gs.scanSite()
		return
	}
	gs.addToIndex(fileData)
	go gs.scanSite()
}

func (gs *gosearch) run() {
	for {
		fmt.Print("Enter word or quit for exit: ")

		var w string
		fmt.Scanln(&w)
		w = strings.ToLower(w)

		if w == "quit" {
			allDocs := gs.index.AllDocs()
			err := gs.storage.WriteData(allDocs)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Exit!")
			return
		}

		docs := gs.engine.Search(w)
		fmt.Printf("%+v\n", docs)
	}
}

func (gs *gosearch) addToIndex(data []crawler.Document) {
	for _, doc := range data {
		gs.index.Add(doc)
	}
}

func (gs *gosearch) scanSite() {
	data, err := gs.scanner.Scan(gs.url, gs.depth)
	if err != nil {
		log.Fatal(err)
	}
	gs.addToIndex(data)
}
