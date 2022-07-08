package postgresql

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// var db *sql.DB

// func init() {
// 	db, err := sql.Open("postgres", "postgres://altynay:password@localhost/somedb?sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	err = db.Ping()
// 	if err != nil {
// 		//w.WriteHeader(http.StatusInternalServerError)
// 		log.Fatal(err)

// 	}
// 	fmt.Print("connected")
// }
