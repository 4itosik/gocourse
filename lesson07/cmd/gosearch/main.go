package main

import (
	"fmt"
	"lesson07/pkg/crawler"
	"lesson07/pkg/crawler/spider"
	"lesson07/pkg/engine"
	"lesson07/pkg/index"
	"lesson07/pkg/index/hash"
	"lesson07/pkg/storage"
	"lesson07/pkg/storage/filestore"
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
	if err == nil {
		addToIndex(gs, fileData)
		go scanSite(gs)
	}

	if err != nil {
		scanSite(gs)
	}
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

func addToIndex(gs *gosearch, data []crawler.Document) {
	for _, doc := range data {
		gs.index.Add(doc)
	}
}

func scanSite(gs *gosearch) {
	data, err := gs.scanner.Scan(gs.url, gs.depth)
	if err != nil {
		log.Fatal(err)
	}
	addToIndex(gs, data)
}
