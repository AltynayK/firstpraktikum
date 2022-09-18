package repository

import (
	"database/sql"

	"github.com/AltynayK/firstpraktikum/internal/service"
)

var DB *sql.DB

type Repo interface {
	InsertData(string, string, string) bool
	InsertMultipleData(string, string, string, string) bool
	GetLongURLByID(int) string
}

var DBdns *string
var fileStoragePath *string
var baseURL *string

func GetDataConfig(a *string, b *string, c *string) {
	DBdns = a
	fileStoragePath = b
	baseURL = c

}
func New() Repo {
	switch {
	case *DBdns != "":
		NewPostgresDB(DBdns)
		return NewDataBase()
	default:
		service.ReadFile(fileStoragePath)
		return NewFile()
	}
}
