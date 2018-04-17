package db

import (
	"database/sql"
	"log"
	"errors"
)


type postGres struct {
	*sql.DB
}

func NewPostGres() Store {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return &postGres{
		db,
	}
}

func (db *postGres) Add(id string, user User) error {
	sqlStatement := `INSERT INTO users (name, address, contact)VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, user.Name, user.Address, user.Contact)
	if err != nil {
		panic(err)
	}
	return nil
}

func (db *postGres) Get(id string) (User, error) {
	sqlStatement := `SELECT * FROM users WHERE id = $1`
	_, err := db.Exec(sqlStatement,id)
	if err != nil {
		panic(err)
	}
	return User{}, errors.New("not found")
}


func (db *postGres) Update(user User) error {
	return nil
}

func (db *postGres) Delete(id string) error {
	return nil
}
