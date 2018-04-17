package db

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
