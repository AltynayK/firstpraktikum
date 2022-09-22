package main

import (
	"fmt"
	"net/http"

	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/handler"
	"github.com/AltynayK/firstpraktikum/internal/short"

	_ "github.com/lib/pq"
)

// var (
// 	ServerAddress   string
// 	BaseURL         string
// 	FileStoragePath string
// 	DatabaseDNS     string
// )

// func init() {
// 	//increment#5
// 	flag.StringVar(&ServerAddress, "a", "127.0.0.1:8080", "ServerAddress - адрес запуска HTTP-сервера")
// 	flag.StringVar(&BaseURL, "b", "http://"+ServerAddress, "BaseURL")
// 	//increment#
// 	flag.StringVar(&FileStoragePath, "f", "texts.txt", "FileStoragePath - путь до файла LongURL")
// 	//flag.StringVar(&DatabaseDNS, "d", "host=localhost port=5432 user=altynay password=password dbname=somedb sslmode=disable", "DatabaseDNS")
// 	flag.StringVar(&DatabaseDNS, "d", "", "DatabaseDNS")
// }

func main() {
	config := app.NewConfig()
	mux := handler.InitHandlers()
	// if u, f := os.LookupEnv("SERVER_ADDRESS"); f {
	// 	config.ServerAddress = u
	// }
	// if u, f := os.LookupEnv("BASE_URL"); f {
	// 	config.BaseURL = u
	// }
	// if u, flg := os.LookupEnv("FILE_STORAGE_PATH"); flg {
	// 	config.FileStoragePath = u
	// }
	// if u, f := os.LookupEnv("DatabaseDNS"); f {

	// 	config.DatabaseDNS = u
	// }

	short.GetBaseURL(&config.BaseURL)

	srv := http.Server{
		Addr:    config.ServerAddress,
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Print(err)
	}
}
