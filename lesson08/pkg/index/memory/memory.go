package memory

import (
	"lesson08/pkg/crawler"
	"lesson08/pkg/index"
)

var docs = []crawler.Document{
	{
		URL:   "https://yandex.ru",
		Title: "Яндекс",
	},
}

// Index - структура, которая содержит данные и индекс для теста
type Index struct {
	index.Interface
}

// New - конструктор.
func New() *Index {
	s := Index{}
	return &s
}

// IDs - возвращает всегда одинаковые ID документа по индексу для теста.
func (ind *Index) IDs(w string) []int {
	ids := []int{1}
	return ids
}

// Docs - возвращает всегда один и тот же документ для теста.
func (ind *Index) Docs(ids []int) []crawler.Document {
	return docs
}
