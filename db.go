package db

import (
	"os"
	"fmt"
	"io/ioutil"
	"github.com/Jeffail/gabs"
)

// Store ...
type Store interface {
	Add(string, User) error
	Get(string) (User, error)
	Update(User) error
	Delete(string) error
}

type User struct {
	Name    string `datastore:"name"`
	Address string `datastore:"address"`
	Contact string `datastore:"contact"`
}


func getCredentials ()interface{} {
	raw, err := ioutil.ReadFile("./credentials.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	settings, err := gabs.ParseJSON([]byte(raw))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return settings
}
