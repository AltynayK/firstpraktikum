package repository

import (
	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/service"
)

type Repo interface {
	InsertData(string, string, string) bool
	InsertMultipleData(string, string, string, string) bool
	GetLongURLByID(int) string
}

func New(config *app.Config) Repo {

	switch {
	case config.DatabaseDNS != "":
		NewPostgresDB(&config.DatabaseDNS)
		return NewDataBase()
	default:
		service.ReadFile(&config.FileStoragePath)
		return NewFile()
	}
}
