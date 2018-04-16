package db_test

import (
	"testing"
	"github.com/srikrsna/db-bench"
)

func BenchmarkNewInMemoryStore(b *testing.B) {
	store := db.NewInMemoryStore()

	for i := 0; i< b.N; i++ {
		store.Add("id", db.Aggragate{})
	}
}

func BenchmarkNewInMemoryStorePar(b *testing.B) {
	store := db.NewInMemoryStore()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Add("id", db.Aggragate{})
		}
	})
}