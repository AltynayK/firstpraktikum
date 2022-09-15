package repository

import (
	"database/sql"

	"github.com/AltynayK/firstpraktikum/internal/models"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg models.Config) (*sql.DB, error) {
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
