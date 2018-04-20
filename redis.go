package db

import (
	"errors"
	"log"
	"github.com/go-redis/redis"
	"bytes"
	"encoding/gob"
	"sync"
)

type redisDB struct {
	*redis.Client
	bytePool sync.Pool
}

func Newredis() Store {
	settings := getCredentials()

	client := redis.NewClient(&redis.Options{
		Addr:     settings.Redis.Addr,
		Password: settings.Redis.Password,
		DB:       settings.Redis.DB,
	})

	return &redisDB{
		client,
		sync.Pool{New: func() interface{} {
			return new(bytes.Buffer)
		},
		},
	}
}

func (db *redisDB) Add(id string, user User) error {

	data := db.bytePool.Get().(*bytes.Buffer)
	data.Reset()

	enc := gob.NewEncoder(data)
	if err := enc.Encode(&user); err != nil {
		return err
	}

	err := db.Client.Set(id, data.Bytes(), 0).Err()
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
