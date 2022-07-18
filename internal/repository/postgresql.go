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

// func CreateTable() {

// 	_, err: = db.Exec("CREATE TABLE IF NOT EXISTS data (id serial primary key, short_url varchar, original_url varchar, user_id varchar)")
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func InsertDataToDB(short_url string, original_url string, user_id string) {

// 	sqlStatement := `INSERT INTO data (short_url, original_url, user_id) VALUES ($1, $2, $3)`
// 	_, err := db.Exec(sqlStatement, short_url, original_url, user_id)
// 	if err != nil {
// 		panic(err)
// 	}
// }
