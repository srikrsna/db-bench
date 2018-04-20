package db

import (
	"errors"
	"log"
	"github.com/nats-io/go-nats-streaming"
	"sync"
	"bytes"
	"encoding/gob"
)

type nats struct {
	stan.Conn
	bytePool sync.Pool

}

func NewNats() Store {
	settings := getCredentials()

	conn, err := stan.Connect(settings.NATS.ClusterID, settings.NATS.ClientID)
	if err != nil {
		log.Print(err)

	}
	return &nats{
		conn,
		sync.Pool{New: func() interface{} {
			return new(bytes.Buffer)
		},
		},
	}
}

func (db *nats) Add(id string, user User) error {
	data := db.bytePool.Get().(*bytes.Buffer)
	data.Reset()

	enc := gob.NewEncoder(data)
	if err := enc.Encode(&user); err != nil {
		return err
	}

	db.Publish("test",data.Bytes())


	return nil
}

func (db *nats) Get(id string) (User, error) {

	return User{}, errors.New("not found")
}

func (db *nats) Update(user User) error {
	return nil
}

func (db *nats) Delete(id string) error {
	return nil
}
