package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/models"
	"github.com/AltynayK/firstpraktikum/internal/service"
)

type File struct {
}

func NewFile() Repo {
	return &File{}
}
func (f *File) InsertData(shortURL string, longURL string, userID string) bool {
	service.WriteToFile(longURL)
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
	return err == nil
}
func (f *File) InsertMultipleData(shortURL string, longURL string, userID string, correlationID string) bool {
	service.WriteToFile(longURL)
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
	return err == nil
}
func InsertData(longURL string) int {
	if service.IDList == nil {
		service.IDList = make(map[int]string)
	}
	service.ID++
	service.IDList[service.ID] = longURL
	return service.ID
}

func (f *File) GetLongURLByID(id int) string {
	//fmt.Print(service.IDList[id])
	return service.IDList[id]

}

func (f *File) MakeShortURL(url string) string {

	return app.NewConfig().BaseURL + "/" + strconv.Itoa(InsertData(url))
}
func (f *File) ReturnShortURL(url string) string {
	return url
}
func (f *File) CheckDeletion(id int) bool {
	return true
}
func (f *File) DeleteMultiple([]int) {

}
