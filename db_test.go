package db_test

import (
	"testing"
	"github.com/srikrsna/db-bench"
)

var user = db.User{Name:"test", Address: "India", Contact: "9988776651",}

func BenchmarkNewInMemoryStore(b *testing.B) {
	b.ReportAllocs()
	store := db.NewInMemoryStore()
	for i := 0; i < b.N; i++ {
		store.Add("id", user)
	}
}

func BenchmarkNewInMemoryStorePar(b *testing.B) {
	b.ReportAllocs()
	store := db.NewInMemoryStore()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Add("id", user)
		}
	})
}

func BenchmarkNewDataStore(b *testing.B) {
	b.ReportAllocs()
	store := db.NewDataStore()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add("id", user)
	}
}

