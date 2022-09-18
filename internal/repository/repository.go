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

func GetDataConfig(a *string) {
	DBdns = a
}
func New() Repo {
	switch {
	case *DBdns != "":
		NewPostgresDB(DBdns)
		return NewDataBase()
	default:
		return NewFile()
	}
}
