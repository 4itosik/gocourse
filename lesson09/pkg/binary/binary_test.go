package binary

import (
	"math/rand"
	"testing"
)

func BenchmarkBinary(b *testing.B) {
	bin := New()
	sampleData(bin)
	b.Run("binary search", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := rand.Intn(1000)
			res, _ := bin.Search(n)
			_ = res
		}
	})
}

func sampleData(ind *Index) {
	for i := 0; i < 100_000; i++ {
		el := Element{
			Value: i,
			ID:    i,
		}
		ind.Insert(&el)
	}
}
