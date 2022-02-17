package storage

import "github.com/jmoiron/sqlx"

type Storage interface {
	Get() string
}

func (d DbStorage) Get() string {
	return ""
}

// DbStorage implements the Storage methods in memory as golang maps
type DbStorage struct {
	db *sqlx.DB
}

// NewDbStorage returns a NewDbStorage with internal maps initialized
func NewDbStorage(db *sqlx.DB) *DbStorage {
	return &DbStorage{db: db}
}
