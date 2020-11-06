// Package bst - реализация структуры данных "Двоичное дерево поиска"
// Пример для простоты приведён для дерева, содержащего в качетве значений целые числа.
// Можно по аналогии со стандартной библиотекой с помощью интерфейса обобщить на произвольные типы данных.
// Для этого потребуется контракт на функцию сравнения элементов.
//
// Википедия: Двоичное дерево поиска — это двоичное дерево, для которого выполняются следующие дополнительные условия (свойства дерева поиска):
// - Оба поддерева — левое и правое — являются двоичными деревьями поиска.
// - У всех узлов левого поддерева произвольного узла X значения ключей данных меньше, нежели значение ключа данных самого узла X.
// - У всех узлов правого поддерева произвольного узла X значения ключей данных больше либо равны, нежели значение ключа данных самого узла X.
package bst

// Tree - Двоичное дерево поиска
type Tree struct {
	root *Element
}

// Element - элемент дерева
type Element struct {
	left, right *Element
	Value       interface{}
	ID          int
}

// Insert - вставка элемента в дерево
func (t *Tree) Insert(e *Element) {
	if t.root == nil {
		t.root = e
		return
	}
	insert(t.root, e)
}

// inset рекурсивно вставляет элемент в нужный уровень дерева.
func insert(node, new *Element) {
	if new.ID < node.ID {
		if node.left == nil {
			node.left = new
			return
		}
		insert(node.left, new)
	}
	if new.ID >= node.ID {
		if node.right == nil {
			node.right = new
			return
		}
		insert(node.right, new)
	}
}

// Search - поиск значения в дереве, выдаёт найденный.
func (t *Tree) Search(id int) (*Element, bool) {
	return search(t.root, id)
}
func search(el *Element, id int) (*Element, bool) {
	if el == nil {
		return nil, false
	}
	if el.ID == id {
		return el, true
	}
	if el.ID < id {
		return search(el.right, id)
	}
	return search(el.left, id)
}

// IDs - собирает ID елементов в дереве
func (t *Tree) IDs(s *[]int) {
	t.root.ids(s)
}
func (el *Element) ids(s *[]int) {
	if el == nil {
		return
	}

	*s = append(*s, el.ID)

	el.left.ids(s)
	el.right.ids(s)
}
