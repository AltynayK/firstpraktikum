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
