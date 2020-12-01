package bst

import (
	"math/rand"
	"testing"
)

func BenchmarkBst(b *testing.B) {
	var tree Tree
	sampleData(&tree)
	b.Run("bst search", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := rand.Intn(1000)
			res, _ := tree.Search(n)
			_ = res
		}
	})
}

func sampleData(t *Tree) {
	for i := 0; i < 100_000; i++ {
		el := Element{
			Value: i,
			ID:    i,
		}
		t.Insert(&el)
	}
}
