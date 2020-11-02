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

// NewIndex создаёт новую структуру с индексом и данными
func NewIndex() *Index {
	var t bst.Tree
	i := Index{
		Data:      t,
		Service:   make(map[string][]int),
		CurrentID: 0,
	}
	return &i
}

// Add - добавляет документ и обновляет индекс.
func (i *Index) Add(url string, title string) {
	nextID := i.CurrentID + 1
	i.CurrentID = nextID

	p := Page{id: nextID, url: url, title: title}
	e := bst.Element{
		Value: p,
		ID:    nextID,
	}
	i.Data.Insert(&e)

	words := strings.Fields(title)
	for _, w := range words {
		w = strings.ToLower(w)
		value := i.Service[w]

		if !contains(value, nextID) {
			value = append(value, nextID)
			i.Service[w] = value
		}
	}
}

// IDs - возвращает ID документа по индексу.
func (i Index) IDs(w string) []int {
	ids := i.Service[w]
	return ids
}

// Pages - возвращает документы по массиву ID.
func (i Index) Pages(ids []int) []Page {
	var result []Page

	for _, id := range ids {
		search(i, id, &result)
	}

	return result
}
func search(i Index, id int, r *[]Page) {
	e, found := i.Data.Search(id)

	if found {
		p := e.Value.(Page) // Это вообще методом тыка и гугла смог подобрать
		*r = append(*r, p)
	}
}

func (i Index) String() string {
	var ids []int
	i.Data.IDs(&ids)

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
