package repository

import "database/sql"

type WorkService struct {
	db *sql.DB
}

func NewWorkService(db *sql.DB) *WorkService {
	return &WorkService{db: db}
}

func (db *sql.DB) CreateTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS data (id serial primary key, short_url varchar, original_url varchar, user_id varchar)")
	if err != nil {
		panic(err)
	}
}
