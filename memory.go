package db

import (
	"sync"
	"errors"
)

type inMemoryStore struct {
	sync.Map
}

func (db *inMemoryStore) Add(id string, agr Aggragate) error {
	db.Map.Store(id, agr)
	return nil
}

func (db *inMemoryStore) Get(id string) (Aggragate, error) {
	if val, ok := db.Load(id); ok {
		return val.(Aggragate), nil
	}

	return Aggragate{}, errors.New("not found")
}

func (db *inMemoryStore) Update(agr Aggragate) error {
	return db.Add(agr.ID, agr)
}

func (db *inMemoryStore) Delete(id string) error {
	db.Map.Delete(id)
	return nil
}

func NewInMemoryStore() Store {
	return &inMemoryStore{
	}
}
