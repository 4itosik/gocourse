// Package binary - реалазиция поиска используя бинарный поиск
package binary

import (
	"sort"
)

// Index - хранилище данных
type Index struct {
	Data []*Element
}

// Element - элемент массив
type Element struct {
	Value interface{}
	ID    int
}

// New - конструктор.
func New() *Index {
	var data []*Element
	ind := Index{
		Data: data,
	}
	return &ind
}

// Insert - добавлеяет елемент.
func (ind *Index) Insert(e *Element) {
	data := ind.Data
	data = append(data, e)
	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})
	ind.Data = data
}

// Search - поиск значения, выдаёт найденный.
func (ind *Index) Search(id int) (*Element, bool) {
	i := sort.Search(len(ind.Data), func(i int) bool { return ind.Data[i].ID >= id })
	if i < len(ind.Data) && ind.Data[i].ID == id {
		return ind.Data[i], true
	}

	return nil, false
}
