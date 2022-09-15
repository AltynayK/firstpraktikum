package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/AltynayK/firstpraktikum/internal/models"
	"github.com/AltynayK/firstpraktikum/internal/service"
)

var DB *sql.DB

type Repo interface {
	InsertData(string, string, string)
	InsertMultipleData(string, string, string, string)
	GetLongURLByID(int)
}

type File struct{}
type DataBase struct{}
type Lists struct{}

func (d *DataBase) InsertData(shortURL string, originalURL string, userID string) bool {
	sqlStatement := `INSERT INTO data (short_url, original_url, user_id) VALUES ($1, $2, $3)`
	_, err := DB.Exec(sqlStatement, shortURL, originalURL, userID)
	return err == nil
}
func (f *File) InsertData(longURL string, shortURL string, userID string) {
	rankings := models.Abs{
		LongURL:  longURL,
		ShortURL: shortURL,
		UserID:   userID,
	}
	rankingsJSON, _ := json.Marshal(rankings)
	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Print("Folder does not exist.")
		}
	}
	file.Write(rankingsJSON)
	file.WriteString("\n")
}

//var id = len(service.IDList) - 1

func (l *Lists) InsertData(longURL string) {
	if service.IDList == nil {
		service.IDList = make(map[int]string)
	}
	//id++
	service.IDList[len(service.IDList)] = longURL
}

func (d *DataBase) InsertMultipleData(shortURL string, originalURL string, userID string, correlationID string) bool {
	sqlStatementt := `INSERT INTO data (short_url, original_url, user_id, correlation_id) VALUES ($1, $2, $3, $4)`
	_, err := DB.Exec(sqlStatementt, shortURL, originalURL, userID, correlationID)
	return err == nil
}

func (f *File) InsertMultipleData(shortURL string, longURL string, userID string, correlationID string) {
	rankings := models.Abs{
		CorrelationID: correlationID,
		LongURL:       longURL,
		ShortURL:      shortURL,
		UserID:        userID,
	}
	rankingsJSON, _ := json.Marshal(rankings)
	file, err := os.OpenFile("outputMultiple.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Print("Folder does not exist.")
		}
	}
	file.Write(rankingsJSON)
	file.WriteString("\n")
}

func (d *DataBase) GetLongURLByID(id int) string {
	row := DB.QueryRow("SELECT original_url FROM data WHERE id = $1", id)
	alb := models.DBUrl{}
	if err := row.Scan(&alb.Originalurl); err != nil {
		fmt.Print("Error.")
	}
	return alb.Originalurl
}

func (f *File) GetLongURLByID(id int) string {
	return service.IDList[id]
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
