package db

import (
	"os"
	"log"
	"encoding/json"
)
type Config struct {
	Datastore struct {
		ProjectID string `json:"projectId"`
	} `json:"datastore"`
	Mongo struct {
		URL string `json:"url"`
	} `json:"mongo"`
	DynamoDb struct {
		Region   string `json:"region"`
		FileName string `json:"fileName"`
		Profile  string `json:"profile"`
	} `json:"dynamoDb"`
	Memcached struct {
		Server string `json:"server"`
	} `json:"memcached"`
	Redis struct {
		Addr     string `json:"Addr"`
		Password string `json:"Password"`
		DB       int    `json:"DB"`
	} `json:"redis"`
	MySql struct {
		DbSourceName string `json:"dbSourceName"`
	} `json:"mySql"`
	PostGres struct {
		DbSourceName string `json:"dbSourceName"`
	} `json:"postGres"`
	EventStore struct {
		URL      string `json:"url"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"eventStore"`
	Kafka struct {
		Brokers string `json:"brokers"`
	} `json:"kafka"`
	NATS struct {
		ClusterID string `json:"clusterId"`
		ClientID  string `json:"clientId"`
	} `json:"NATS"`
}

// Store ...
type Store interface {
	Add(string, User) error
	Get(string) (User, error)
	Update(User) error
	Delete(string) error
}

type User struct {
	Name    string `datastore:"name" json:"name" bson:"name"`
	Address string `datastore:"address" json:"address" bson:"address"`
	Contact string `datastore:"contact" json:"contact" bson:"contact"`
}

func getCredentials() *Config {
	f, err := os.Open("./credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var config Config

	if err := json.NewDecoder(f).Decode(&config); err != nil {
		log.Fatalln(err)
	}

	return &config
}
