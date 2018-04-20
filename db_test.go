package db_test

import (
	"testing"
	"github.com/srikrsna/db-bench"


)

var user = db.User{Name: "test", Address: "India", Contact: "9988776651",}

func BenchmarkInMemoryStore_Add(b *testing.B) {
	b.ReportAllocs()
	store := db.NewInMemoryStore()
	for i := 0; i < b.N; i++ {
		store.Add("id", user)
	}
}

func BenchmarkInMemoryStorePar_Add(b *testing.B) {
	b.ReportAllocs()
	store := db.NewInMemoryStore()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Add("id", user)
		}
	})
}

func BenchmarkDataStore_Add(b *testing.B) {
	b.ReportAllocs()
	store := db.NewDataStore()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add("id", user)
	}
}

func BenchmarkRedis_Add(b *testing.B) {
	b.ReportAllocs()
	store := db.Newredis()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add(fmt.Sprintf("%d", i), user)
	}
}

func BenchmarkMongo_Add(b *testing.B) {
	b.ReportAllocs()
	store := db.NewMongo()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add(fmt.Sprintf("%d", i), user)
	}
}

func BenchmarkPostGres_Add(b *testing.B) {
	b.ReportAllocs()
	store := db.NewPostGres()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add(fmt.Sprintf("%d", i), user)
	}
}

func BenchmarkMySql_Add(b *testing.B) {
	b.ReportAllocs()
	store := db.NewMySql()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add(fmt.Sprintf("%d", i), user)
	}
}

func BenchmarkDynamoDB_Add(b *testing.B) {
	b.ReportAllocs()
	store := db.NewDynamoDB()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add(fmt.Sprintf("%d", i), user)
	}
}

func BenchmarkMemcached_Add(b *testing.B) {
	b.ReportAllocs()
	store := db.NewMemcached()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add(fmt.Sprintf("%d", i), user)
	}
}

func BenchmarkEventStore_Add(b *testing.B) {
	b.ReportAllocs()
	store := db.NewEventStore()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Add(fmt.Sprintf("%d", i), user)
	}
}

func BenchmarkEventStorePar_Add(b *testing.B) {
	b.ReportAllocs()
	store := db.NewEventStore()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Add("id", user)
		}
	})

}

func BenchmarkKafkaProducerPar_Add(b *testing.B) {
	b.ReportAllocs()
	store := db.NewKafka()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Add("id", user)
		}
	})

}


func BenchmarkNewNats_Add(b *testing.B) {
	var store = db.NewNats()

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Add("id", user)
		}
	})

}