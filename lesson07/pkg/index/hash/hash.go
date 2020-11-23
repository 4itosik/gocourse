// Package hash реализует обратный индекс для документов
// Пакет позволяет добавлять документы в индекс и искать документы по индексу
package hash

import (
	"fmt"
	"lesson07/pkg/crawler"
	"lesson07/pkg/index/hash/bst"
	"strings"
)

// Index - структура, которая содержит данные и индекс для поиска
type Index struct {
	Data      bst.Tree
	Service   map[string][]int
	CurrentID int
	urls      []string
}

// New - создаёт новую структуру с индексом и данными.
func New() *Index {
	var t bst.Tree
	ind := Index{
		Data:      t,
		Service:   make(map[string][]int),
		CurrentID: 0,
		urls:      []string{},
	}
	return &ind
}

// Add - добавляет документ и обновляет индекс.
func (ind *Index) Add(doc crawler.Document) {
	if containsURL(ind.urls, doc.URL) {
		return
	}

	nextID := ind.CurrentID + 1
	ind.CurrentID = nextID

	e := bst.Element{
		Value: doc,
		ID:    nextID,
	}
	ind.Data.Insert(&e)

	urls := append(ind.urls, doc.URL)
	ind.urls = urls

	words := strings.Fields(doc.Title)
	for _, w := range words {
		w = strings.ToLower(w)
		value := ind.Service[w]

		if !containsID(value, nextID) {
			value = append(value, nextID)
			ind.Service[w] = value
		}
	}
}

// IDs - возвращает ID документа по индексу.
func (ind *Index) IDs(w string) []int {
	ids := ind.Service[w]
	return ids
}

// Docs - возвращает документы по массиву ID.
func (ind *Index) Docs(ids []int) []crawler.Document {
	var r []crawler.Document

	for _, id := range ids {
		e, found := ind.Data.Search(id)

		if found {
			d := e.Value.(crawler.Document)
			r = append(r, d)
		}
	}

	return r
}

// AllDocs - возвращает все документы
func (ind *Index) AllDocs() []crawler.Document {
	var arr []bst.Element
	ind.Data.Elements(&arr)

	var r []crawler.Document
	for _, el := range arr {
		d := el.Value.(crawler.Document)
		r = append(r, d)
	}
	return r
}

func (ind *Index) String() string {
	var ids []int
	ind.Data.IDs(&ids)

	var s string
	for _, id := range ids {
		s += fmt.Sprintf("%v ", id)
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

func containsID(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsURL(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
