package bst

import (
	"reflect"
	"testing"
)

func TestBst_Insert(t *testing.T) {
	var tree Tree
	addItem(&tree)

	var got []int
	tree.IDs(&got)

	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("получили %v, ожидалось %v", got, want)
	}
}

func TestBst_Search(t *testing.T) {
	var tree Tree
	addItem(&tree)

	el, find := tree.Search(3)
	if !find {
		t.Errorf("bst.Search(); got not found; want found")
		return
	}

	got := el.Value
	want := 3
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}

func addItem(tree *Tree) {
	e1 := Element{
		Value: 1,
		ID:    1,
	}
	tree.Insert(&e1)

	e2 := Element{
		Value: 2,
		ID:    2,
	}
	tree.Insert(&e2)

	e3 := Element{
		Value: 3,
		ID:    3,
	}
	tree.Insert(&e3)
}
