package main

import (
	"io"
	"net/http"
	"strconv"

	"log"

	"github.com/AltynayK/firstpraktikum/internal/app"
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

	id++
	IDList[id] = url

	//log.Print(id)

	//log.Print(IDList)
	w.WriteHeader(201)

	w.Write([]byte(url))

}

func Get(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "The query parameter is missing", http.StatusBadRequest)
		return
	}
	b, err := strconv.Atoi(id)
	if err != nil && b < 1 {
		w.WriteHeader(400)
		return
		//
	}
	//fmt.Fprintf(w, "Welcome to the home page!")
	w.Header().Set("Location", app.LongURL(IDList[b]))

	// Первый путь
	// id := query["id"][0]

	// Второй способ

	w.WriteHeader(307)

	//w.Header.WriteSubset(w io.Writer, app.LongUrl(IdList[id]))

	//log.Print(app.LongURL(IDList[b]))

	//w.Write([]byte(app.LongURL(IDList[b])))
	//w.Write([]byte(app.LongURL(IDList[id])))
}

// const port = ":8080"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Post)
	mux.HandleFunc("/{id}", Get)
	log.Fatal(http.ListenAndServe(":8080", mux))
	// mux := initHandlers()
	//IDList = make(map[int]string)

	// srv := http.Server{
	// 	Addr:    port,
	// 	Handler: mux,
	// }

	//log.Printf("App listening port: %s", port)
	// log.Fatal(srv.ListenAndServe())

}

// func initHandlers() *mux.Router {
// TODO: how handler 404 (if not found some url, example: /not_exist_url)
// TODO: handle "Not Allowed Method" example: DELETE method request to /

// 	router := mux.NewRouter()
// 	router.HandleFunc("/", Post).Methods("POST")
// 	router.HandleFunc("/id", Get).Methods("GET")

// 	return router
// }

// func WriteShortURLByID(url string) int {

// 	return id

// }

// func GetLongURLFromID(id string) string {

// }
