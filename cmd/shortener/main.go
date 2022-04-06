package main

import (
	"io"
	"net/http"
	"strconv"

	"log"

	"github.com/AltynayK/firstpraktikum/internal/app"
)

var (
	IdList map[int]string
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

	url := app.ShortUrl(string(b))
	if IdList == nil {
		IdList = make(map[int]string)
	}

	if IdList[id] != url {
		id++
		IdList[id] = url
	}

	log.Print(id)

	//log.Print(IdList)
	w.WriteHeader(201)

	w.Write([]byte(url))

}

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		log.Print("upal")
		return
	}
	w.WriteHeader(307)

	r.Header.Set("Location", app.LongUrl(IdList[id]))

	// remove row below
	log.Print(IdList[id])
	w.Write([]byte(app.LongUrl(IdList[id])))
	//w.Write([]byte(GetLongURLFromID(string(q))))
	//http.Get()
}

func main() {

	http.HandleFunc("/", Post)
	http.HandleFunc("/{id}", Get)

	server := &http.Server{
		Addr: "localhost:8080",
	}

	server.ListenAndServe()

}

// func WriteShortURLByID(url string) int {

// 	return id

// }

// func GetLongURLFromID(id string) string {

// }
