package db

import (
	"sync"
	"errors"
)

type inMemoryStore struct {
	sync.Map
}

func (db *inMemoryStore) Add(id string, user User) error {
	db.Map.Store(id, user)
	return nil
}

func (db *inMemoryStore) Get(id string) (User, error) {
	if val, ok := db.Load(id); ok {
		return val.(User), nil
	}
	return User{}, errors.New("not found")
}

func (db *inMemoryStore) Update(user User) error {
	return db.Add(user.Name, user)
}

func (db *inMemoryStore) Delete(id string) error {
	db.Map.Delete(id)
	return nil
}

func NewInMemoryStore() Store {
	return &inMemoryStore{
	}
}
