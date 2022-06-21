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
	ServerAddress   *string
	BaseURL         *string
	FileStoragePath *string
)

func init() {
	ServerAddress = flag.String("a", "127.0.0.1:8080", "ServerAddress - адрес запуска HTTP-сервера")
	BaseURL = flag.String("b", "http://"+*ServerAddress, "BaseURL")
	FileStoragePath = flag.String("f", "texts.txt", "FileStoragePath - путь до файла LongURL")
	//DatabaseDNS = flag.String("-d", , "")
}

func main() {
	//mw := handler.CompressGzip{}
	mux := initHandlers()

	flag.Parse()
	if u, f := os.LookupEnv("SERVER_ADDRESS"); f {
		*ServerAddress = u
	}
	if u, f := os.LookupEnv("BASE_URL"); f {
		*BaseURL = u
	}
	//fmt.Print(os.Getenv("BASE_URL"))
	if u, flg := os.LookupEnv("FILE_STORAGE_PATH"); flg {
		*FileStoragePath = u
	}
	short.GetBaseURL(BaseURL)
	//FilePath: = *FILE_STORAGE_PATH
	//os.Setenv("SERVER_ADDRESS", "127.0.0.1:8080")
	//os.Setenv("BASE_URL", *BaseUrl)
	srv := http.Server{
		Addr:    *ServerAddress,
		Handler: mux,
	}
	service.ReadFile(FileStoragePath)

	log.Fatal(srv.ListenAndServe())

}

func initHandlers() *mux.Router {

	router := mux.NewRouter()
	router.Use(handler.CompressGzip)
	//router.Use(handler.GzipHandle)
	router.Use(handler.Cookie)

	router.HandleFunc("/", handler.PostText).Methods("POST")
	router.HandleFunc("/api/shorten", handler.PostJSON).Methods("POST")
	router.HandleFunc("/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/api/user/urls", handler.GetAllUrls).Methods("GET")

	//router.HandleFunc("/ping", ).Methods("GET")
	//router.Use(mw)
	return router
}
