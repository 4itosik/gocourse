package storage

// Хранилище документов на диске.

import (
	"lesson07/pkg/crawler"
)

// Interface определяет контракт хранилища данных.
type Interface interface {
	ReadData() ([]crawler.Document, error)
	WriteData([]crawler.Document) error
}
