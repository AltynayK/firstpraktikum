package repository

import (
	"database/sql"

	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/service"
)

var DB *sql.DB

type Repo interface {
	InsertData(string, string, string) bool
	InsertMultipleData(string, string, string, string) bool
	GetLongURLByID(int) string
}

func New() Repo {
	conf := app.NewConfig()
	switch {
	case conf.DatabaseDNS != "":
		NewPostgresDB(&conf.DatabaseDNS)
		return NewDataBase()
	default:
		service.ReadFile(&conf.FileStoragePath)
		return NewFile()
	}
}
