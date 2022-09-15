package sqlite

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("postgres", path)
	if err != nil {
		return nil, fmt.Errorf("can't open database", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to database", err)
	}
	return &Storage{db: db}, nil
}
