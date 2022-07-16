package postgresql

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB
var DBdns *string

func GetDatabaseDNSS(a *string) {
	DBdns = a
}

func InsertDataToDB(short_url string, original_url string, user_id string) {
	db, err := sql.Open("postgres", *DBdns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStatement := `INSERT INTO data (short_url, original_url, user_id) VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, short_url, original_url, user_id)
	if err != nil {
		panic(err)
	}
}
