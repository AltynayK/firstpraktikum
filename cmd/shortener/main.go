package main

import (
	"net/http"
	"os"

	"log"

	"github.com/AltynayK/firstpraktikum/internal/handler"
	"github.com/gorilla/mux"
)

func main() {
	mux := initHandlers()
	//IDList = make(map[int]string)
	port := ":8080"
	os.Setenv("SERVER_ADDRESS", "localhost"+port)
	os.Setenv("BASE_URL", "http://"+os.Getenv("SERVER_ADDRESS")+"/")
	srv := http.Server{
		Addr:    port,
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
