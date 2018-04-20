package db

import (
	"errors"
	"log"
	"gopkg.in/mgo.v2"
	"crypto/tls"
	"net"
)

type mongoDB struct {
	*mgo.Session
}

func NewMongo() Store {
	settings := getCredentials()

	dialInfo, err := mgo.ParseURL(settings.Mongo.URL)

	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}

	return &mongoDB{
		session,
	}
}

func (db *mongoDB) Add(id string, user User) error {
	c := db.DB("test").C("users")
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if err := c.Insert(&user); err != nil {
		panic(err)
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
