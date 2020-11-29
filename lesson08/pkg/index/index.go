package index

// Обратный индекс отсканированных документов.

import "lesson08/pkg/crawler"

// Interface определяет контракт службы индексирования документов.
type Interface interface {
	Add(crawler.Document)
	IDs(string) []int
	Docs([]int) []crawler.Document
	AllDocs() []crawler.Document
}
