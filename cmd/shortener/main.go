package main

import (
	"fmt"
	"net/http"
	"os"

	"log"

	"github.com/AltynayK/firstpraktikum/internal/handler"
	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
)

type config struct {
	//SERVER_ADDRESS string `opts:"help=listening interface, default=127.0.0.1"`
	Port string `env:"Port" envDefault:":8080"`
}

func main() {
	mux := initHandlers()
	//IDList = make(map[int]string)
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	os.Setenv("SERVER_ADDRESS", "localhost")
	os.Setenv("BASE_URL", "http://"+os.Getenv("SERVER_ADDRESS")+cfg.Port)
	srv := http.Server{
		Addr:    cfg.Port,
		Handler: mux,
	}

	//log.Printf("App listening port: %s", port)
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
