package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB
var DBdns *string

func GetDatabaseDNS(a *string) {
	DBdns = a
}

func Init() {
	db, err := sql.Open("postgres", "DBdns")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		//w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		//fmt.Print("no connected")

	}
	fmt.Print("connected")
}
