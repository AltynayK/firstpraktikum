package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/AltynayK/firstpraktikum/internal/handler"
	"github.com/AltynayK/firstpraktikum/internal/models"
	"github.com/AltynayK/firstpraktikum/internal/repository"
	"github.com/AltynayK/firstpraktikum/internal/service"
	"github.com/AltynayK/firstpraktikum/internal/short"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var conf models.Config

func init() {
	//increment#5
	flag.StringVar(&conf.ServerAddress, "a", "127.0.0.1:8080", "ServerAddress - адрес запуска HTTP-сервера")
	flag.StringVar(&conf.BaseURL, "b", "http://"+conf.ServerAddress, "BaseURL")
	//increment#
	flag.StringVar(&conf.FileStoragePath, "f", "texts.txt", "FileStoragePath - путь до файла LongURL")
	flag.StringVar(&conf.DatabaseDNS, "d", "host=localhost port=5432 user=altynay password=password dbname=somedb sslmode=disable", "DatabaseDNS")
	//flag.StringVar(&conf.DatabaseDNS, "d", "", "DatabaseDNS")
}

func main() {

	mux := initHandlers()

	if u, f := os.LookupEnv("SERVER_ADDRESS"); f {
		conf.ServerAddress = u
	}
	if u, f := os.LookupEnv("BASE_URL"); f {
		conf.BaseURL = u
	}
	if u, flg := os.LookupEnv("FILE_STORAGE_PATH"); flg {
		conf.FileStoragePath = u
	}
	if u, f := os.LookupEnv("DatabaseDNS"); f {
		conf.DatabaseDNS = u
	}

	flag.Parse()
	short.GetBaseURL(&conf.BaseURL)
	handler.GetDatabaseDNS(&conf.DatabaseDNS)
	repository.GetDatabaseDNSs(&conf.DatabaseDNS)
	repository.GetBaseURLL(&conf.BaseURL)
	srv := http.Server{
		Addr:    conf.ServerAddress,
		Handler: mux,
	}

	service.ReadFile(&conf.FileStoragePath)

	repository.NewPostgresDB(repository.Config{
		DBdns: &conf.DatabaseDNS,
	})

	fmt.Print(srv.ListenAndServe())

}

func initHandlers() *mux.Router {

	router := mux.NewRouter()
	router.Use(handler.Decompress)
	router.Use(handler.GzipHandler)
	router.Use(handler.SetCookie)
	//router.Use(handler.CheckCookie)

	router.HandleFunc("/", handler.PostText).Methods("POST")
	router.HandleFunc("/api/shorten", handler.PostJSON).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", handler.Get).Methods("GET")
	router.HandleFunc("/api/user/urls", handler.GetAllUrls).Methods("GET")
	router.HandleFunc("/ping", handler.CheckConnection).Methods("GET")
	router.HandleFunc("/api/shorten/batch", handler.PostMultipleUrls).Methods("POST")
	return router
}
