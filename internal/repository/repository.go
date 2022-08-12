package repository

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func CreateTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS data (id serial primary key, short_url varchar, original_url varchar UNIQUE, user_id varchar, correlation_id varchar)")
	if err != nil {
		panic(err)
	}
	DB = db

}

func InsertDataToDB(shortURL string, originalURL string, userID string) bool {

	sqlStatement := `INSERT INTO data (short_url, original_url, user_id) VALUES ($1, $2, $3)`
	_, err := DB.Exec(sqlStatement, shortURL, originalURL, userID)
	return err == nil
}

func InsertDataToDBCor(shortURL string, originalURL string, userID string, correlationID string) bool {

	sqlStatementt := `INSERT INTO data (short_url, original_url, user_id, correlation_id) VALUES ($1, $2, $3, $4)`
	_, err := DB.Exec(sqlStatementt, shortURL, originalURL, userID, correlationID)
	return err == nil
}

type DBUrls struct {
	id          int
	shorturl    string
	originalurl string
	userid      string
}

func ReturnShortURL(LongURL string) string {
	row := DB.QueryRow("SELECT short_url FROM data WHERE original_url = $1", LongURL)
	alb := DBUrls{}
	if err := row.Scan(&alb.shorturl); err != nil {
		log.Fatal(err)
	}
	return alb.shorturl
}
func Ping() bool {
	err := DB.Ping()
	return err == nil
}
