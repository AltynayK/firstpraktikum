package handler

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/AltynayK/firstpraktikum/internal/service"
	"github.com/AltynayK/firstpraktikum/internal/short"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type URL struct {
	LongURL string `json:"url"`
	Result  string `json:"result"`
}

type URLs struct {
	ShortURL string `json:"short_url"`
	LongURL  string `json:"original_url"`
}

// type PostResponse struct {

// }

func PostJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var url URL
	var jsonRes []byte

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ShortURL := short.WriteShortURL(url.LongURL)
	okRes := URL{
		Result: ShortURL,
	}
	service.MakeData(url.LongURL, ShortURL)
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

	w.Header().Set("content-type", "plain/text")

	url, err := io.ReadAll(r.Body)
	// обрабатываем ошибку
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	longURL := string(url)

	shortURL := short.WriteShortURL(longURL)
	service.MakeData(longURL, shortURL)

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

func GetAllUrls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var jsonRes []byte
	var result []string
	file, err := os.OpenFile("./output.json", os.O_RDONLY|os.O_CREATE, 0777)

	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Folder does not exist.")
			w.WriteHeader(http.StatusNoContent)
		}
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		result = append(result, line)

	}
	if len(result) == 0 {
		w.WriteHeader(http.StatusNoContent)
	}
	jsonRes, _ = json.Marshal(result)

	w.Write(jsonRes)
	return
}

var db *sql.DB
var DBdns *string

func GetDatabaseDNS(a *string) {
	DBdns = a
}

func CheckConnection(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	db, err := sql.Open("postgres", *DBdns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)

	}
	w.WriteHeader(http.StatusAccepted)

}
