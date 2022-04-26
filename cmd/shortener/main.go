package main

import (
	"flag"
	"net/http"
	"os"

	"log"

	"github.com/AltynayK/firstpraktikum/internal/handler"
	"github.com/AltynayK/firstpraktikum/internal/service"
	"github.com/AltynayK/firstpraktikum/internal/short"

	"github.com/gorilla/mux"
)

var (
	SERVER_ADDRESS    *string
	BASE_URL          *string
	FILE_STORAGE_PATH *string
)

func init() {
	SERVER_ADDRESS = flag.String("a", "127.0.0.1:8080", "SERVER_ADDRESS - адрес запуска HTTP-сервера")
	BASE_URL = flag.String("b", "http://"+*SERVER_ADDRESS, "BASE_URL")
	FILE_STORAGE_PATH = flag.String("f", "texts.txt", "FILE_STORAGE_PATH - путь до файла с сокращёнными URL")
}

func main() {
	//mw := handler.CompressGzip{}
	mux := initHandlers()

	flag.Parse()
	if u, f := os.LookupEnv("SERVER_ADDRESS"); f {
		*SERVER_ADDRESS = u
	}
	if u, f := os.LookupEnv("BASE_URL"); f {
		*BASE_URL = u
	}
	//fmt.Print(os.Getenv("BASE_URL"))
	if u, flg := os.LookupEnv("FILE_STORAGE_PATH"); flg {
		*FILE_STORAGE_PATH = u
	}
	short.GetBaseUrl(BASE_URL)
	//FilePath: = *FILE_STORAGE_PATH
	//os.Setenv("SERVER_ADDRESS", "127.0.0.1:8080")
	//os.Setenv("BASE_URL", *BaseUrl)
	srv := http.Server{
		Addr:    *SERVER_ADDRESS,
		Handler: mux,
	}
	service.ReadFile(FILE_STORAGE_PATH)

	log.Fatal(srv.ListenAndServe())

}

func initHandlers() *mux.Router {

	router := mux.NewRouter()
	router.Use(handler.CompressGzip)

	router.HandleFunc("/", handler.PostText).Methods("POST")
	router.HandleFunc("/api/shorten", handler.PostJSON).Methods("POST")
	router.HandleFunc("/{id}", handler.Get).Methods("GET")
	//router.Use(mw)
	return router
}
