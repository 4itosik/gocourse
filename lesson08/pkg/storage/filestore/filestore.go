package filestore

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"lesson08/pkg/crawler"
	"os"
)

// ErrFileNotFound - ошибка, если отсутствует файл, откуда надо прочитать данные
var ErrFileNotFound = errors.New("Файл не найден")

// DB - хранилище данных
type DB struct {
	filepath string
	docs     []crawler.Document
}

// New - конструктор.
func New(filepath string) *DB {
	db := DB{
		filepath: filepath,
	}
	return &db
}

// ReadData - читает данные из файла и возврашает crawler.Document.
func (db *DB) ReadData() ([]crawler.Document, error) {
	if !fileExists(db.filepath) {
		return nil, ErrFileNotFound
	}

	bytes, err := ioutil.ReadFile(db.filepath)
	if err != nil {
		return nil, err
	}

	data := []crawler.Document{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// WriteData - записывает данные в файл.
func (db *DB) WriteData(data []crawler.Document) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(db.filepath, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
