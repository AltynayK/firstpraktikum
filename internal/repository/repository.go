package repository

import (
	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/service"
)

type Repo interface {
	InsertData(string, string, string) bool
	InsertMultipleData(string, string, string, string) bool
	GetLongURLByID(int) string
	MakeShortURL(string) string
	ReturnShortURL(string) string
	CheckDeletion(int) bool
	DeleteMultiple([]int) error
}

func New(config *app.Config) Repo {

	switch {
	case config.DatabaseDNS != "":
		NewPostgresDB(&config.DatabaseDNS)
		return NewDataBase(app.NewConfig())
	default:
		service.ReadFile(&config.FileStoragePath)
		return NewFile()
	}
}
