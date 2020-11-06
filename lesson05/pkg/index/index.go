// Package index реализует обратный индекс для документов
// Пакет позволяет добавлять документы в индекс и искать документы по индексу
package index

import (
	"fmt"
	"lesson05/pkg/bst"
	"strings"
)

// Index - структура, которая содержит данные и индекс для поиска
type Index struct {
	Data      bst.Tree
	Service   map[string][]int
	CurrentID int
}

// Page - стуруктура документа
type Page struct {
	id    int
	url   string
	title string
}

// New - создаёт новую структуру с индексом и данными.
func New() *Index {
	var t bst.Tree
	ind := Index{
		Data:      t,
		Service:   make(map[string][]int),
		CurrentID: 0,
	}
	return &ind
}

// Add - добавляет документ и обновляет индекс.
func (ind *Index) Add(url string, title string) {
	nextID := ind.CurrentID + 1
	ind.CurrentID = nextID

	p := Page{id: nextID, url: url, title: title}
	e := bst.Element{
		Value: p,
		ID:    nextID,
	}
	ind.Data.Insert(&e)

	words := strings.Fields(title)
	for _, w := range words {
		w = strings.ToLower(w)
		value := ind.Service[w]

		if !contains(value, nextID) {
			value = append(value, nextID)
			ind.Service[w] = value
		}
	}
}

// IDs - возвращает ID документа по индексу.
func (ind Index) IDs(w string) []int {
	ids := ind.Service[w]
	return ids
}

// Pages - возвращает документы по массиву ID.
func (ind Index) Pages(ids []int) []Page {
	var r []Page

	for _, id := range ids {
		e, found := ind.Data.Search(id)

		if found {
			p := e.Value.(Page) // Это вообще методом тыка и гугла смог подобрать
			r = append(r, p)
		}
	}

	return r
}

func (ind Index) String() string {
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

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
