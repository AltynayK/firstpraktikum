package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Config struct {
	DBdns *string
}

func NewPostgresDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", *cfg.DBdns)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	CreateTable(db)
	return db, nil

}
func CreateTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS data (id serial primary key, short_url varchar, original_url varchar UNIQUE, user_id varchar, correlation_id varchar)")
	if err != nil {
		panic(err)
	}
	DB = db

}
