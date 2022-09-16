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

func GetDatabaseDNSs(a *string) {
	DBdns = a
}
func New() Repo {
	switch {
	case *DBdns != "":
		return NewDataBase()
	default:
		return NewFile()
	}
}
