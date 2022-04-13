package handler

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

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

	longURL := string(b)
	if IDList == nil {
		IDList = make(map[int]string)
	}

	id++
	IDList[id] = longURL

	shortURL := "http://" + r.Host + r.URL.String() + (strconv.Itoa(id))
	//log.Print(id)

	//log.Print(IDList)
	w.Header().Set("Location", shortURL)
	w.WriteHeader(201)

	w.Write([]byte(shortURL))

}

func Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {

		w.WriteHeader(400)
		return
	}
	//fmt.Println(`id := `, id)

	b, err := strconv.Atoi(id)
	if err != nil && b < 1 {
		w.WriteHeader(400)
		return
	}
	longURL := IDList[b]
	//
	w.Header().Set("Location", longURL)

	w.WriteHeader(307)
	fmt.Fprint(w)
	//w.Header.WriteSubset(w io.Writer, app.LongUrl(IdList[id]))

	//log.Print(app.LongURL(IDList[b]))

	//w.Write([]byte(app.LongURL(IDList[b])))
	//w.Write([]byte(app.LongURL(IDList[id])))
}
