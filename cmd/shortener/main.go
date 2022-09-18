package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/AltynayK/firstpraktikum/internal/handler"
	"github.com/AltynayK/firstpraktikum/internal/repository"
	"github.com/AltynayK/firstpraktikum/internal/service"
	"github.com/AltynayK/firstpraktikum/internal/short"

	_ "github.com/lib/pq"
)

var (
	ServerAddress   string
	BaseURL         string
	FileStoragePath string
	DatabaseDNS     string
)

func init() {
	//increment#5
	flag.StringVar(&ServerAddress, "a", "127.0.0.1:8080", "ServerAddress - адрес запуска HTTP-сервера")
	flag.StringVar(&BaseURL, "b", "http://"+ServerAddress, "BaseURL")
	//increment#
	flag.StringVar(&FileStoragePath, "f", "texts.txt", "FileStoragePath - путь до файла LongURL")
	//flag.StringVar(&DatabaseDNS, "d", "host=localhost port=5432 user=altynay password=password dbname=somedb sslmode=disable", "DatabaseDNS")
	flag.StringVar(&DatabaseDNS, "d", "", "DatabaseDNS")
}

func main() {
	mux := handler.InitHandlers()
	if u, f := os.LookupEnv("SERVER_ADDRESS"); f {
		ServerAddress = u
	}
	if u, f := os.LookupEnv("BASE_URL"); f {
		BaseURL = u
	}
	if u, flg := os.LookupEnv("FILE_STORAGE_PATH"); flg {
		FileStoragePath = u
	}
	if u, f := os.LookupEnv("DatabaseDNS"); f {
		DatabaseDNS = u
	}

	flag.Parse()
	short.GetBaseURL(&BaseURL)
	handler.GetDatabaseDNS(&DatabaseDNS)
	repository.GetDatabaseDNSs(&DatabaseDNS)
	repository.GetBaseURLL(&BaseURL)
	srv := http.Server{
		Addr:    ServerAddress,
		Handler: mux,
	}
	service.ReadFile(&FileStoragePath)

	repository.NewPostgresDB(repository.Config{
		DBdns: &DatabaseDNS,
	})

	//fmt.Print(srv.ListenAndServe())
	if err := srv.ListenAndServe(); err != nil {
		fmt.Print(err)
	}
}
