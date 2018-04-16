package db

// Store ...
type Store interface {
	Add(string, Aggragate) error
	Get(string) (Aggragate, error)
	Update(Aggragate) error
	Delete(string) error
}

// Aggragate ...
type Aggragate struct {
	ID string		
}
