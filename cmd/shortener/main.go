package main

import (
	"io"
	"net/http"

	"log"

	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/gorilla/mux"
)

var (
	IDList map[int]string
	id     int
)

func Post(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "plain/text")

	b, err := io.ReadAll(r.Body)
	// обрабатываем ошибку
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	url := app.ShortURL(string(b))
	if IDList == nil {
		IDList = make(map[int]string)
	}

	if IDList[id] != url {
		id++
		IDList[id] = url
	}

	//log.Print(id)

	//log.Print(IdList)
	w.WriteHeader(201)

	w.Write([]byte(url))
	return

}

func Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(307)

	r, err := http.NewRequest("GET", r.URL.Query().Get("id"), nil)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Location", app.LongURL(IDList[id]))

	//

	//w.Header.WriteSubset(w io.Writer, app.LongUrl(IdList[id]))

	//log.Print(IDList[id])

	//w.Write([]byte(app.LongURL(IDList[id])))
	//w.Write([]byte(app.LongURL(IDList[id])))
	return
}

const port = ":8080"

func main() {
	mux := initHandlers()
	IDList = make(map[int]string)

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
	router.HandleFunc("/", Post).Methods("POST")
	router.HandleFunc("/{id}", Get).Methods("GET")

	return router
}

// func WriteShortURLByID(url string) int {

// 	return id

// }

// func GetLongURLFromID(id string) string {

// }
