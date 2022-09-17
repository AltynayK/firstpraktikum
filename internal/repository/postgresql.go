package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/AltynayK/firstpraktikum/internal/models"
	_ "github.com/lib/pq"
)

type Config struct {
	DBdns *string
}
type DataBase struct{}

func NewDataBase() Repo {
	return &DataBase{}
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
func CreateTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS data (id serial primary key, short_url varchar, original_url varchar UNIQUE, user_id varchar, correlation_id varchar)")
	if err != nil {
		panic(err)
	}
	DB = db

}
func (d *DataBase) InsertData(shortURL string, originalURL string, userID string) bool {
	sqlStatement := `INSERT INTO data (short_url, original_url, user_id) VALUES ($1, $2, $3)`
	_, err := DB.Exec(sqlStatement, shortURL, originalURL, userID)
	return err == nil
}

func (d *DataBase) InsertMultipleData(shortURL string, originalURL string, userID string, correlationID string) bool {
	sshortURL := MakeShortURLToDB(originalURL)
	shortURL = sshortURL
	sqlStatementt := `INSERT INTO data (short_url, original_url, user_id, correlation_id) VALUES ($1, $2, $3, $4)`
	_, err := DB.Exec(sqlStatementt, shortURL, originalURL, userID, correlationID)
	return err == nil
}

func (d *DataBase) GetLongURLByID(id int) string {
	row := DB.QueryRow("SELECT original_url FROM data WHERE id = $1", id)
	alb := models.DBUrl{}
	if err := row.Scan(&alb.Originalurl); err != nil {
		fmt.Print("Error.")
	}
	return alb.Originalurl
}
func ReturnShortURL(LongURL string) string {
	row := DB.QueryRow("SELECT short_url FROM data WHERE original_url = $1", LongURL)
	alb := models.DBUrls{}
	if err := row.Scan(&alb.Shorturl); err != nil {
		fmt.Print("Error.")
	}
	return alb.Shorturl
}

func Ping() bool {
	err := DB.Ping()
	return err == nil
}

var Init *string

func GetBaseURLL(a *string) {
	Init = a
}

func MakeShortURLToDB(url string) string {
	db := DB
	id := db.QueryRow("SELECT id FROM data ORDER BY id DESC LIMIT 1")
	alb := models.Dbid{}
	if err := id.Scan(&alb.Maxid); err != nil {
		alb.Maxid = 0
	}
	nextid := alb.Maxid + 1
	return *Init + "/" + strconv.Itoa(nextid)
}
