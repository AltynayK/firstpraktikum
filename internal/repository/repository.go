package repository

import (
	"database/sql"
)

var DB *sql.DB

type Repo interface {
	InsertData(string, string, string) bool
	InsertMultipleData(string, string, string, string) bool
	GetLongURLByID(int) string
}

var DBdns *string
var fileStoragePath *string

func GetDataConfig(a *string, b *string) {
	DBdns = a
	fileStoragePath = b
}

func New() Repo {
	switch {
	case *DBdns != "":
		NewPostgresDB(DBdns)
		return NewDataBase()
	default:
		//service.ReadFile(fileStoragePath)
		return NewFile()
	}
}
