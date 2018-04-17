package db

import (
	"errors"
	"log"
	"github.com/go-redis/redis"
)


type redisDB struct {
	*redis.Client
}

func Newredis() Store {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &redisDB{
		client,
	}
}

func (db *redisDB) Add(id string, user User) error {

	err := db.Client.Set(id, &user, 0).Err()
	if err != nil {
		log.Fatalf("Failed to save %v", err)

	}

	return nil
}

func (db *redisDB) Get(id string) (User, error) {
	_, err := db.Client.Get(id).Result()
	if err != nil {
		panic(err)
	}
	return User{}, errors.New("not found")
}


func (db *redisDB) Update(user User) error {
	return nil
}

func (db *redisDB) Delete(id string) error {
	return nil
}
