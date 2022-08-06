package handler

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/AltynayK/firstpraktikum/internal/repository"
	"github.com/AltynayK/firstpraktikum/internal/service"
	"github.com/AltynayK/firstpraktikum/internal/short"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type URL struct {
	LongURL string `json:"url"`
	Result  string `json:"result"`
	//CorrelationID string `json:"correlation_id"`
}
type MultURL struct {
	CorrelationID string `json:"correlation_id"`
	Result        string `json:"short_url"`
}
type URLs struct {
	ShortURL      string `json:"short_url"`
	LongURL       string `json:"original_url"`
	CorrelationID string `json:"correlation_id"`
}
type DbUrl struct {
	id          int
	shorturl    string
	originalurl string
	userid      string
}

var db *sql.DB
var DBdns *string

func GetDatabaseDNS(a *string) {
	DBdns = a
}

func PostJSON(w http.ResponseWriter, r *http.Request) {
	var ShortURL string
	w.Header().Set("content-type", "application/json")
	var url URL
	var jsonRes []byte
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a := r.Context().Value(userCtxKey).(string)

	//_, exists := os.LookupEnv("DatabaseDNS")
	if repository.Ping() == true {
		//if exists {
		ShortURL = short.WriteShortURL(url.LongURL)
		//service.WriteURLByID(url.LongURL)

		if repository.InsertDataToDB(ShortURL, url.LongURL, a) == false {
			ShortURL = repository.ReturnShortURL(url.LongURL)
			w.WriteHeader(409)
		} else {
			//service.MakeData(url.LongURL, ShortURL, a)
			w.WriteHeader(201)
		}

	} else {
		ShortURL = short.WriteShortURL(url.LongURL)

		w.WriteHeader(201)
	}
	service.MakeData(url.LongURL, ShortURL, a)
	okRes := URL{
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
	if repository.Ping() == true {
		//_, exists := os.LookupEnv("DatabaseDNS")
		//if exists {
		shortURL = short.WriteShortURL(longURL)
		//service.WriteURLByID(longURL)
		if repository.InsertDataToDB(shortURL, longURL, a) == false {
			shortURL = repository.ReturnShortURL(longURL)
			w.WriteHeader(409)
		} else {
			//service.MakeData(longURL, shortURL, a)
			w.WriteHeader(201)
		}

	} else {
		shortURL = short.WriteShortURL(longURL)

		w.WriteHeader(201)

	}
	service.MakeData(longURL, shortURL, a)
	w.Header().Set("Location", shortURL)
	w.Write([]byte(shortURL))
}

func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
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

	_, exists := os.LookupEnv("DatabaseDNS")
	if exists {
		db = repository.DB

		row := db.QueryRow("SELECT original_url FROM data WHERE id = $1", b)
		alb := DbUrl{}
		if err := row.Scan(&alb.originalurl); err != nil {
			log.Fatal(err)
		}
		longURLL := alb.originalurl
		w.Header().Set("Location", longURLL)
		w.WriteHeader(307)
		fmt.Fprint(w)
	} else {
		longURL := service.GetURLFromID(b)
		w.Header().Set("Location", longURL)
		w.WriteHeader(307)
		fmt.Fprint(w)
	}

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
	w.WriteHeader(http.StatusOK)
}

func GetAllUrls(w http.ResponseWriter, r *http.Request) {
	type URLStruct struct {
		Shorturl    string `json:"short_url"`
		Originalurl string `json:"original_url"`
		Userid      string `json:"userID"`
	}
	var x []*URLStruct
	var jsonRes []byte
	var result string
	w.Header().Set("content-type", "application/json")
	_, exists := os.LookupEnv("DatabaseDNS")
	if exists {
		db = repository.DB

		rows, _ := db.Query("SELECT short_url, original_url, user_id FROM data WHERE user_id = $1", r.Context().Value(userCtxKey))

		defer rows.Close()
		var albums []URLStruct

		for rows.Next() {
			alb := URLStruct{}
			if err := rows.Scan(&alb.Shorturl, &alb.Originalurl, &alb.Userid); err != nil {
				log.Fatal(err)
			}
			albums = append(albums, alb)
		}
		//fmt.Print(albums)
		data, _ := json.MarshalIndent(albums, " ", " ")
		w.Write(data)
		return

	} else {
		file, err := os.OpenFile("./output.txt", os.O_RDONLY|os.O_CREATE, 0777)
		if err != nil {
			if os.IsNotExist(err) {
				log.Fatal("Folder does not exist.")
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

		err = json.Unmarshal(jsonRes, &x)
		var x2 []*URLStruct
		for _, v := range x {
			if v.Userid == r.Context().Value(userCtxKey) {
				x2 = append(x2, v)
			}
		}
		if x2 == nil {
			w.WriteHeader(http.StatusNoContent)
		}
		data, err := json.MarshalIndent(x2, " ", " ")
		w.Write(data)
		return
	}
}

type Posts []URLs

func PostMultipleUrls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var url Posts
	content, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(content, &url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var okRes MultURL
	var ShortURL string
	var jsonRes []byte
	var JsonArray []MultURL
	for _, value := range url {
		ShortURL = short.WriteShortURL(value.LongURL)
		okRes = MultURL{
			CorrelationID: value.CorrelationID,
			Result:        ShortURL,
		}
		a := r.Context().Value(userCtxKey).(string)
		_, exists := os.LookupEnv("DatabaseDNS")
		if exists {
			repository.InsertDataToDBCor(ShortURL, value.LongURL, a, okRes.CorrelationID)
		} else {
			service.MakeDataForMultipleCase(ShortURL, value.LongURL, a, okRes.CorrelationID)

		}
		JsonArray = append(JsonArray, okRes)
	}

	if jsonRes, err = json.Marshal(JsonArray); err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "response json marshal err")
		return
	}
	w.WriteHeader(201)
	fmt.Fprint(w, string(jsonRes))

}
