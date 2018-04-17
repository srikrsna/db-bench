package db

import (
	"errors"
	"context"
	"cloud.google.com/go/datastore"
	"log"
)

type dataStore struct {
	*datastore.Client
}

func NewDataStore() Store {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, "db-test-201306")
	if err != nil {
		log.Fatal(err)
	}
	return &dataStore{
		dsClient,
	}
}

func (db *dataStore) Add(id string, user User) error {
	userKey := datastore.NameKey("users", "", nil)
	if _, err := db.Put(context.Background(), userKey, &user); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}
	return nil
}

func (db *dataStore) Get(id string) (User, error) {
	return User{}, errors.New("not found")
}

func (db *dataStore) Update(user User) error {
	return nil
}

func (db *dataStore) Delete(id string) error {

	return nil
}
