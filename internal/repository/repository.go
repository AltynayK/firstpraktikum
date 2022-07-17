package repository

import (
	"database/sql"
)

var DB *sql.DB

func CreateTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS data (id serial primary key, short_url varchar, original_url varchar UNIQUE, user_id varchar, correlation_id varchar)")
	if err != nil {
		panic(err)
	}
	DB = db

}

func InsertDataToDB(short_url string, original_url string, user_id string) bool {

	sqlStatement := `INSERT INTO data (short_url, original_url, user_id) VALUES ($1, $2, $3)`
	_, err := DB.Exec(sqlStatement, short_url, original_url, user_id)
	if err != nil {
		return false
	}
	return true
}
func InsertDataToDBCor(short_url string, original_url string, user_id string, correlation_id string) bool {

	sqlStatement := `INSERT INTO data (short_url, original_url, user_id, correlation_id) VALUES ($1, $2, $3, $4)`
	_, err := DB.Exec(sqlStatement, short_url, original_url, user_id, correlation_id)
	if err != nil {
		return false
	}
	return true
}

func Ping() bool {
	err := DB.Ping()
	if err != nil {
		return false
	}
	return true
}
