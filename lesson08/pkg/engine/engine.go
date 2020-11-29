package engine

import (
	"lesson08/pkg/crawler"
	"lesson08/pkg/index"
)

// Engine - поисковый движок.
// Его задача - обслуживание поисковых запросов.
// функциональность:
// - обработка поискового запроса;
// - поиск документов в индексе;
// - возврат посиковой выдачи.

// Service - поисковый движок.
type Service struct {
	index index.Interface
}

// New - конструктор.
func New(index index.Interface) *Service {
	s := Service{
		index: index,
	}
	return &s
}

// Search ищет документы, соответствующие поисковому запросу.
func (s *Service) Search(query string) []crawler.Document {
	if query == "" {
		return nil
	}

	ids := s.index.IDs(query)
	docs := s.index.Docs(ids)
	return docs
}
