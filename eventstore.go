package db

import (
	"errors"
	"log"
	"github.com/jetbasrawi/go.geteventstore"
)

type eventStore struct {
	*goes.Client
}

func NewEventStore() Store {
	settings := getCredentials()

	client, err := goes.NewClient(nil, settings.EventStore.URL)
	if err != nil {
		log.Fatal(err)
	}
	client.SetBasicAuth(settings.EventStore.Username, settings.EventStore.Password)

	return &eventStore{
		client,
	}
}

func (db *eventStore) Add(id string, user User) error {
	eventMeta := make(map[string]string)
	eventMeta["Id"] = id

	goesEvent := goes.NewEvent(goes.NewUUID(), "User", user, eventMeta)

	writer := db.NewStreamWriter("UserStream")

	err := writer.Append(nil, goesEvent)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (db *eventStore) Get(id string) (User, error) {

	reader := db.NewStreamReader("UserStream")
	for reader.Next() {
		if reader.Err() != nil {
		}
		user := User{}
		fooMeta := make(map[string]string)
		err := reader.Scan(&user, &fooMeta)
		if err != nil {
			log.Fatal(err)
		}
	}

	return User{}, errors.New("not found")
}

func (db *eventStore) Update(user User) error {
	return nil
}

func (db *eventStore) Delete(id string) error {
	return nil
}
