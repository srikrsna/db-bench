package db

import (
	"database/sql"
	"log"
	"errors"
)


type mySql struct {
	*sql.DB
}

func NewMySql() Store {
	settings := getCredentials()
	db, err := sql.Open("mysql", settings.MySql.DbSourceName)
	if err != nil {
	log.Fatal(err)
	}
	defer db.Close()

	return &mySql{
		db,
	}
}

func (db *mySql) Add(id string, user User) error {
	sqlStatement := `INSERT INTO users (name, address, contact)VALUES (?, ?, ?)`
	_, err := db.Exec(sqlStatement, user.Name, user.Address, user.Contact)
	if err != nil {
		panic(err)
	}
	return nil
}

func (db *mySql) Get(id string) (User, error) {
	sqlStatement := `SELECT * FROM users WHERE id = ?`
	_, err := db.Exec(sqlStatement,id)
	if err != nil {
		panic(err)
	}
	return User{}, errors.New("not found")
}


func (db *mySql) Update(user User) error {
	return nil
}

func (db *mySql) Delete(id string) error {
	return nil
}
