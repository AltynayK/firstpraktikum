package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/AltynayK/firstpraktikum/internal/service"
	"github.com/gorilla/mux"
)

type Url struct {
	LongURL string `json:"url" json:"URL"`
}
type PostResponse struct {
	Result string `json:"result"`
}

func PostJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var url Url
	var jsonRes []byte

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ShortURL := "http://localhost:8080/api/shorten/" + strconv.Itoa(service.WriteURLByID(url.LongURL))

	okRes := PostResponse{
		Result: ShortURL,
	}

	if jsonRes, err = json.Marshal(okRes); err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "response json marshal err")

		return
	}

	w.Header().Set("Location", ShortURL)

	w.WriteHeader(201)
	fmt.Fprint(w, string(jsonRes))

}

func PostText(w http.ResponseWriter, r *http.Request) {
	// var (
	// 	url *domain.Url
	// )

	w.Header().Set("content-type", "plain/text")

	url, err := io.ReadAll(r.Body)
	// обрабатываем ошибку
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	longURL := string(url)

	shortURL := "http://localhost:8080/" + (strconv.Itoa(service.WriteURLByID(longURL)))
	//shortURL := "http://" + r.Host + r.URL.String() + (strconv.Itoa(service.WriteURLByID(longURL)))
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
	longURL := service.GetURLFromID(b)
	//
	w.Header().Set("Location", longURL)

	w.WriteHeader(307)
	fmt.Fprint(w)
	//w.Header.WriteSubset(w io.Writer, app.LongUrl(IdList[id]))

	//log.Print(app.LongURL(IDList[b]))

	//w.Write([]byte(app.LongURL(IDList[b])))
	//w.Write([]byte(app.LongURL(IDList[id])))
}
