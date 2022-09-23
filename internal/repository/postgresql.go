package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/models"
	_ "github.com/lib/pq"
)

type DataBase struct {
	config *app.Config
	db     *sql.DB
}

func NewDataBase(config *app.Config) Repo {
	return &DataBase{
		config: config,
	}
}
func NewPostgresDB(cfg *string) (*sql.DB, error) {
	db, err := sql.Open("postgres", *cfg)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS data (id serial primary key, short_url varchar, original_url varchar UNIQUE, user_id varchar, correlation_id varchar)")
	if err != nil {
		panic(err)
	}

	return db, nil

}

// func CreateTable(db *sql.DB) {

// 	DB = db

// }
func (d *DataBase) InsertData(shortURL string, originalURL string, userID string) bool {
	sqlStatement := `INSERT INTO data (short_url, original_url, user_id) VALUES ($1, $2, $3)`
	_, err := d.db.Exec(sqlStatement, shortURL, originalURL, userID)
	return err == nil
}

func (d *DataBase) InsertMultipleData(shortURL string, originalURL string, userID string, correlationID string) bool {

	sqlStatementt := `INSERT INTO data (short_url, original_url, user_id, correlation_id) VALUES ($1, $2, $3, $4)`
	_, err := d.db.Exec(sqlStatementt, shortURL, originalURL, userID, correlationID)
	return err == nil
}

func (d *DataBase) GetLongURLByID(id int) string {
	row := d.db.QueryRow("SELECT original_url FROM data WHERE id = $1", id)
	alb := models.DBUrl{}
	if err := row.Scan(&alb.Originalurl); err != nil {
		fmt.Print("Error.")
	}
	return alb.Originalurl
}
func (d *DataBase) ReturnShortURL(LongURL string) string {
	row := d.db.QueryRow("SELECT short_url FROM data WHERE original_url = $1", LongURL)
	alb := models.DBUrls{}
	if err := row.Scan(&alb.Shorturl); err != nil {
		fmt.Print("Error.")
	}
	return alb.Shorturl
}

func Ping() bool {
	db, err := sql.Open("postgres", app.NewConfig().DatabaseDNS)

	if err != nil {
		return err == nil
	}
	err = db.Ping()
	return err == nil
}

func (d *DataBase) MakeShortURL(url string) string {
	id := d.db.QueryRow("SELECT id FROM data ORDER BY id DESC LIMIT 1")
	alb := models.Dbid{}
	if err := id.Scan(&alb.Maxid); err != nil {
		alb.Maxid = 0
	}
	nextid := alb.Maxid + 1
	return app.NewConfig().BaseURL + "/" + strconv.Itoa(nextid)
}
