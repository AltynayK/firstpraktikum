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
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS data (id INTEGER PRIMARY KEY, short_url varchar, original_url varchar, user_id varchar")
	statement.Exec()
	statement, _ = db.Prepare("INSERT INTO data (short_url, original_url, user_id) VALUES (?, ?, ?)")
	statement.Exec(short_url, original_url, user_id)
	//statement.Exec()

}
