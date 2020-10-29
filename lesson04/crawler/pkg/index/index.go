// Package index реализует обратный индекс для документов
// Пакет позволяет добавлять документы в индекс и искать документы по индексу
package index

import (
	"sort"
	"strings"
)

var index = make(map[string][]int)

// Page - стуруктура документа
type Page struct {
	id    int
	url   string
	title string
}

var pages []Page

// Add - добавляет документ и обновляет индекс.
func Add(url string, title string) {
	nextID := nextID()

	p := Page{id: nextID, url: url, title: title}
	pages = append(pages, p)
	sort.Slice(pages, func(i, j int) bool {
		return pages[i].id < pages[j].id
	})

	words := strings.Fields(title)
	for _, w := range words {
		w = strings.ToLower(w)
		value, _ := index[w]

		if !contains(value, nextID) {
			value = append(value, nextID)
			index[w] = value
		}
	}
}

// IDs - возвращает ID документа по индексу.
func IDs(w string) []int {
	ids := index[w]
	return ids
}

// Pages - возвращает документы по массиву ID.
func Pages(ids []int) []Page {
	var result []Page

	for _, id := range ids {
		search(id, &result)
	}

	return result
}

func search(id int, r *[]Page) {
	i := sort.Search(len(pages), func(i int) bool { return pages[i].id >= id })
	if i < len(pages) && pages[i].id == id {
		*r = append(*r, pages[i])
	}
}

func nextID() int {
	maxNumber := 0
	for _, p := range pages {
		if p.id > maxNumber {
			maxNumber = p.id
		}
	}
	return maxNumber + 1
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
