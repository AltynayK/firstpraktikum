package handler

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/AltynayK/firstpraktikum/internal/models"
	"github.com/AltynayK/firstpraktikum/internal/repository"
	"github.com/AltynayK/firstpraktikum/internal/short"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB
var DBdns *string

func GetDatabaseDNS(a *string) {
	DBdns = a
}

var d = repository.DataBase{}
var f = repository.File{}

//increment#2
func PostJSON(w http.ResponseWriter, r *http.Request) {
	var ShortURL string
	w.Header().Set("content-type", "application/json")
	var url models.URL
	var jsonRes []byte
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	a := r.Context().Value(userCtxKey).(string)
	if *DBdns != "" {
		ShortURL = short.Hash(url.LongURL)
		if !d.InsertData(ShortURL, url.LongURL, a) {
			ShortURL = repository.ReturnShortURL(url.LongURL)
			w.WriteHeader(409)
		} else {
			w.WriteHeader(201)
		}
	} else {
		ShortURL = short.WriteShortURL(url.LongURL)
		f.InsertData(url.LongURL, ShortURL, a)
		w.WriteHeader(201)
	}
	okRes := models.URL{
		Result: ShortURL,
	}
	if jsonRes, err = json.Marshal(okRes); err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "response json marshal err")
		return
	}
	w.Header().Set("Location", ShortURL)
	fmt.Fprint(w, string(jsonRes))
}

//increment#1
func PostText(w http.ResponseWriter, r *http.Request) {
	var shortURL string
	w.Header().Set("content-type", "plain/text")
	url, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	longURL := string(url)
	a := r.Context().Value(userCtxKey).(string)
	if *DBdns != "" {
		shortURL = short.Hash(longURL)
		if !d.InsertData(shortURL, longURL, a) {
			shortURL = repository.ReturnShortURL(longURL)
			w.WriteHeader(409)
		} else {
			w.WriteHeader(201)
		}
	} else {
		shortURL = short.WriteShortURL(longURL)
		f.InsertData(longURL, shortURL, a)
		w.WriteHeader(201)
	}
	w.Header().Set("Location", shortURL)
	w.Write([]byte(shortURL))
}

//increment#1
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var longURL string
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(400)
		return
	}
	b, err := strconv.Atoi(id)
	if err != nil && b < 1 {
		w.WriteHeader(400)
		return
	}
	if *DBdns != "" {
		longURL = d.GetLongURLByID(b)
	} else {
		longURL = f.GetLongURLByID(b)
	}
	w.Header().Set("Location", longURL)
	w.WriteHeader(307)
	fmt.Fprint(w)
}

//increment#10
func CheckConnection(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	db, err := sql.Open("postgres", *DBdns)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

//increment#9
func GetAllUrls(w http.ResponseWriter, r *http.Request) {
	var x []*models.URLStruct
	var jsonRes []byte
	var result string
	w.Header().Set("content-type", "application/json")
	file, err := os.OpenFile("./output.txt", os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNoContent)
		}
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if result == "" {
			result = line
		}
		if result != "" && line != "\n" {
			result = result + "," + line
		}
	}
	a := "[" + result + "]"
	jsonRes = []byte(a)

	json.Unmarshal(jsonRes, &x)
	var x2 []*models.URLStruct
	for _, v := range x {
		if v.Userid == r.Context().Value(userCtxKey) {
			x2 = append(x2, v)
		}
	}
	if x2 == nil {
		w.WriteHeader(http.StatusNoContent)
	}
	data, _ := json.MarshalIndent(x2, " ", " ")
	w.Write(data)
}

type Posts []models.URLs

//increment#12
func PostMultipleUrls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var url Posts
	content, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(content, &url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var okRes models.MultURL
	var ShortURL string
	var jsonRes []byte
	var JSONArray []models.MultURL
	for _, value := range url {
		a := r.Context().Value(userCtxKey).(string)
		if repository.Ping() {
			ShortURL = short.MakeShortURLToDB(value.LongURL)
			d.InsertMultipleData(ShortURL, value.LongURL, a, okRes.CorrelationID)
			ShortURL = repository.ReturnShortURL(value.LongURL)
		} else {
			ShortURL = short.WriteShortURL(value.LongURL)
			f.InsertMultipleData(ShortURL, value.LongURL, a, okRes.CorrelationID)
		}
		okRes = models.MultURL{
			CorrelationID: value.CorrelationID,
			Result:        ShortURL,
		}
		JSONArray = append(JSONArray, okRes)
	}
	if jsonRes, err = json.Marshal(JSONArray); err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "response json marshal err")
		return
	}
	w.WriteHeader(201)
	fmt.Fprint(w, string(jsonRes))

}
