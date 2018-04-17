package db

import (
	"errors"
	"log"
	"gopkg.in/mgo.v2"
)

type mongoDB struct {
	*mgo.Session
}

func NewMongo() Store {
	session, err := mgo.Dial("server1.example.com,server2.example.com")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return &mongoDB{
		session,
	}
}

func (db *mongoDB) Add(id string, user User) error {
	c := db.DB("test").C("people")
	err := c.Insert(&user)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (db *mongoDB) Get(id string) (User, error) {
	_ ,err := db.Get(id)
	if err != nil {
		log.Fatal(err)
	}
	return User{}, errors.New("not found")
}


func (db *mongoDB) Update(user User) error {
	return nil
}

func (db *mongoDB) Delete(id string) error {
	return nil
}
