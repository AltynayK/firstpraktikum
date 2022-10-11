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
	dB     *sql.DB
}

func NewDataBase(config *app.Config) Repo {
	return &DataBase{
		config: config,
		dB:     NewPostgresDB(&config.DatabaseDNS),
	}
}
func NewPostgresDB(cfg *string) *sql.DB {
	db, err := sql.Open("postgres", *cfg)
	if err != nil {
		return nil
	}
	err = db.Ping()
	if err != nil {
		return nil
	}

	CreateTable(db)
	return db

}

func CreateTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS data (id serial primary key, short_url varchar, original_url varchar UNIQUE, user_id varchar, correlation_id varchar, active bool  DEFAULT true)")
	if err != nil {
		panic(err)
	}

}
func (d *DataBase) InsertData(shortURL string, originalURL string, userID string) bool {
	sqlStatement := `INSERT INTO data (short_url, original_url, user_id) VALUES ($1, $2, $3)`
	_, err := d.dB.Exec(sqlStatement, shortURL, originalURL, userID)
	return err == nil
}

func (d *DataBase) InsertMultipleData(shortURL string, originalURL string, userID string, correlationID string) bool {

	sqlStatementt := `INSERT INTO data (short_url, original_url, user_id, correlation_id) VALUES ($1, $2, $3, $4)`
	_, err := d.dB.Exec(sqlStatementt, shortURL, originalURL, userID, correlationID)
	return err == nil
}

func (d *DataBase) GetLongURLByID(id int) string {
	row := d.dB.QueryRow("SELECT original_url FROM data WHERE id = $1", id)
	alb := models.DBUrl{}
	if err := row.Scan(&alb.Originalurl); err != nil {
		fmt.Print("Error GetLongURLByID.")
	}
	return alb.Originalurl
}
func (d *DataBase) ReturnShortURL(LongURL string) string {
	row := d.dB.QueryRow("SELECT short_url FROM data WHERE original_url = $1", LongURL)
	alb := models.DBUrls{}
	if err := row.Scan(&alb.Shorturl); err != nil {
		fmt.Print("Error ReturnShortURL.")
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
	id := d.dB.QueryRow("SELECT id FROM data ORDER BY id DESC LIMIT 1")
	alb := models.Dbid{}
	if err := id.Scan(&alb.Maxid); err != nil {
		alb.Maxid = 0
	}
	nextid := alb.Maxid + 1
	return d.config.BaseURL + "/" + strconv.Itoa(nextid)
}

func (d *DataBase) CheckDeletion(id int) bool {
	row := d.dB.QueryRow("SELECT active FROM data WHERE id = $1", id)
	alb := models.DBUrl{}
	if err := row.Scan(&alb.Active); err != nil {
		fmt.Print("Error CheckStatus.")
	}
	return alb.Active
}

// func (d *DataBase) Delete(id int, err error) {
// 	d.dB.QueryRow("UPDATE data SET active=false WHERE id = $1", id)
// }
func (d *DataBase) DeleteMultiple(id []int) error {

	for _, i := range id {
		d.dB.QueryRow("UPDATE data SET active=false WHERE id = $1", i)
	}
	return nil
}
