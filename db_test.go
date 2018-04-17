package db_test

import (
	"testing"
	"github.com/srikrsna/db-bench"
	"fmt"
)

var user = db.User{Name:"test", Address: "India", Contact: "9988776651",}

func BenchmarkNewInMemoryStore(b *testing.B) {
	store := db.NewInMemoryStore()
	for i := 0; i < b.N; i++ {
		store.Add("id", user)
	}
}

func BenchmarkNewInMemoryStorePar(b *testing.B) {
	store := db.NewInMemoryStore()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Add("id", user)
		}
	})
}

func BenchmarkNewDataStore(b *testing.B) {
	store := db.NewDataStore()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add("id", user)
	}
}


func BenchmarkNewredis(b *testing.B) {
	store := db.Newredis()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add(fmt.Sprint("%d", i), user)
	}
}

func BenchmarkNewMongo(b *testing.B) {
	store := db.NewMongo()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add(fmt.Sprint("%d", i), user)
	}
}


func BenchmarkNewPostGres(b *testing.B) {
	store := db.NewPostGres()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add(fmt.Sprint("%d", i), user)
	}
}


func BenchmarkNewMySql(b *testing.B) {
	store := db.NewMySql()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add(fmt.Sprint("%d", i), user)
	}
}

func BenchmarkNewDynamoDB(b *testing.B) {
	store := db.NewDynamoDB()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add(fmt.Sprint("%d", i), user)
	}
}


func BenchmarkNewMemcached(b *testing.B) {
	store := db.NewMemcached()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add(fmt.Sprint("%d", i), user)
	}
}

