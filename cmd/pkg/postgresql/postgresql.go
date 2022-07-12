package postgresql

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// )

// var db *sql.DB
// var DBdns *string

// func GetDatabaseDNS(a *string) {
// 	DBdns = a
// }
// func Init() {

// 	db, err := sql.Open("postgres", "DBdns")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	err = db.Ping()
// 	if err != nil {

// 		log.Fatal(err)
// 		fmt.Print("no connected")

// 	}
// 	fmt.Print("connected")

// }
