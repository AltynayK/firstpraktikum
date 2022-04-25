package main

import (
	"flag"
	"net/http"
	"os"

	"log"

	"github.com/AltynayK/firstpraktikum/internal/handler"
	"github.com/AltynayK/firstpraktikum/internal/service"

	"github.com/gorilla/mux"
)

// type Path struct {
// 	FilePath string
// }

// type Config struct {
// 	FileStoragePath string `env:"FILE_STORAGE_PATH"`
// }

func main() {
	//	var path Path
	mux := initHandlers()
	//IDList = make(map[int]string)

	ServerAddress := flag.String("-a", "127.0.0.1:8080", "SERVER_ADDRESS - адрес запуска HTTP-сервера")
	BaseUrl := flag.String("-b", "http://"+*ServerAddress, "BASE_URL")
	flag.Parse()
	if u, f := os.LookupEnv("SERVER_ADDRESS"); f {
		*ServerAddress = u
	}
	if u, f := os.LookupEnv("BASE_URL"); f {
		*BaseUrl = u
	}
	//os.Setenv("SERVER_ADDRESS", "127.0.0.1:8080")
	os.Setenv("BASE_URL", *BaseUrl)
	srv := http.Server{
		Addr:    *ServerAddress,
		Handler: mux,
	}
	service.ReadFile()
	// cfg := Config{}
	// if err := env.Parse(&cfg); err != nil {
	// 	fmt.Printf("%+v\n", err)
	// }

	log.Fatal(srv.ListenAndServe())

}

func initHandlers() *mux.Router {
	// TODO: how handler 404 (if not found some url, example: /not_exist_url)
	// TODO: handle "Not Allowed Method" example: DELETE method request to /

	router := mux.NewRouter()
	router.HandleFunc("/", handler.PostText).Methods("POST")
	router.HandleFunc("/api/shorten", handler.PostJSON).Methods("POST")
	router.HandleFunc("/{id}", handler.Get).Methods("GET")

	return router
}

// func WriteShortURLByID(url string) int {

// 	return id

// }

// func GetLongURLFromID(id string) string {

// }
