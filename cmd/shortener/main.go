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

	"github.com/gorilla/mux"
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

var Checkvar bool

func main() {

	mux := initHandlers()

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

	srv := http.Server{
		Addr:    ServerAddress,
		Handler: mux,
	}

	service.ReadFile(&FileStoragePath)

	repository.NewPostgresDB(repository.Config{
		DBdns: &DatabaseDNS,
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
