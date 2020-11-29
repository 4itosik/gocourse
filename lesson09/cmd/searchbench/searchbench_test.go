package main

import (
	"lesson09/pkg/binary"
	"lesson09/pkg/bst"
	"math/rand"
	"testing"
)

func BenchmarkBinary(b *testing.B) {
	bin := binary.New()
	sampleDataForBin(bin)
	b.Run("binary search", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := rand.Intn(1000)
			res, _ := bin.Search(n)
			_ = res
		}
	})
}

func BenchmarkBst(b *testing.B) {
	var tree bst.Tree
	sampleDataForBst(&tree)
	b.Run("bst search", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := rand.Intn(1000)
			res, _ := tree.Search(n)
			_ = res
		}
	})
}

func sampleDataForBin(ind *binary.Index) {
	for i := 0; i < 100_000; i++ {
		el := binary.Element{
			Value: i,
			ID:    i,
		}
		ind.Insert(&el)
	}
}

func sampleDataForBst(t *bst.Tree) {
	for i := 0; i < 100_000; i++ {
		el := bst.Element{
			Value: i,
			ID:    i,
		}
		t.Insert(&el)
	}
}
