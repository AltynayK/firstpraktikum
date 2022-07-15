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
	"time"

	"github.com/AltynayK/firstpraktikum/internal/service"
	"github.com/AltynayK/firstpraktikum/internal/short"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
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
	cookie, err := r.Cookie("session")
	Id = uuid.NewV4()
	if err != nil {

		cookie = &http.Cookie{
			Name:       "session",
			Value:      Id.String(),
			Path:       "",
			Domain:     "",
			Expires:    time.Time{},
			RawExpires: "",
			MaxAge:     0,
			Secure:     false,
			HttpOnly:   true,
			SameSite:   0,
			Raw:        "",
			Unparsed:   []string{},
		}

	}
	ShortURL := short.WriteShortURL(url.LongURL)
	okRes := URL{
		Result: ShortURL,
	}
	service.MakeData(url.LongURL, ShortURL, cookie.Value)
	if jsonRes, err = json.Marshal(okRes); err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "response json marshal err")

		return
	}
	//postgresql.InsertDataToDB(ShortURL, url.LongURL, cookie.Value)
	w.Header().Set("Location", ShortURL)

	w.WriteHeader(201)
	fmt.Fprint(w, string(jsonRes))

}

func PostText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "plain/text")
	cookie, err := r.Cookie("session")
	Id = uuid.NewV4()
	if err != nil {

		cookie = &http.Cookie{
			Name:       "session",
			Value:      Id.String(),
			Path:       "",
			Domain:     "",
			Expires:    time.Time{},
			RawExpires: "",
			MaxAge:     0,
			Secure:     false,
			HttpOnly:   true,
			SameSite:   0,
			Raw:        "",
			Unparsed:   []string{},
		}

	}
	url, err := io.ReadAll(r.Body)
	// обрабатываем ошибку
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	longURL := string(url)

	shortURL := short.WriteShortURL(longURL)
	//
	service.MakeData(longURL, shortURL, cookie.Value)

	w.Header().Set("Location", shortURL)
	w.WriteHeader(201)

	//postgresql.InsertDataToDB(shortURL, longURL, cookie.Value)

	w.Write([]byte(shortURL))
	//db.Exec("insert into urls (short_url, original_url) values (shortURL, longURL)")
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
	longURL := service.GetURLFromID(b)

	w.Header().Set("Location", longURL)

	w.WriteHeader(307)

	fmt.Fprint(w)

}

// func GetAllUrls(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")

// 	var jsonRes []byte
// 	var result []string

// 	file, err := os.OpenFile("./output.json", os.O_RDONLY|os.O_CREATE, 0777)

// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			log.Fatal("Folder does not exist.")
// 			w.WriteHeader(http.StatusNoContent)
// 		}
// 	}
// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		result = append(result, line)

// 	}
// 	jsonRes, _ = json.Marshal(result)
// 	w.Write(jsonRes)
// 	return
// }

var db *sql.DB
var DBdns *string

func GetDatabaseDNS(a *string) {
	DBdns = a
}
func CheckConnection(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")

	db, err := sql.Open("postgres", *DBdns)
	//fmt.Println(*DBdns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		//fmt.Print("no connected")

	}
	//fmt.Print("connected")
	w.WriteHeader(http.StatusOK)

}

// type MultipleUrl []struct {
// 	CorrelationID string `json:"correlation_id"`
// 	OriginalURL   string `json:"original_url"`
// }

// func PostMultipleUrls(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")

// 	var url []MultipleUrl
// 	//var jsonRes []byte

// 	err := json.NewDecoder(r.Body).Decode(&url)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	cookie, err := r.Cookie("session")
// 	Id = uuid.NewV4()
// 	if err != nil {

// 		cookie = &http.Cookie{
// 			Name:       "session",
// 			Value:      Id.String(),
// 			Path:       "",
// 			Domain:     "",
// 			Expires:    time.Time{},
// 			RawExpires: "",
// 			MaxAge:     0,
// 			Secure:     false,
// 			HttpOnly:   true,
// 			SameSite:   0,
// 			Raw:        "",
// 			Unparsed:   []string{},
// 		}

// 	}
// 	ShortURL := short.WriteShortURL(url.OriginalURL)
// 	okRes := MultipleUrl{
// 		CorrelationID: CorrelationID,
// 		short_url:     ShortURL,
// 	}
// 	service.MakeData(url.LongURL, ShortURL, cookie.Value)
// 	if jsonRes, err = json.Marshal(okRes); err != nil {
// 		w.WriteHeader(500)
// 		fmt.Fprintf(w, "response json marshal err")

// 		return
// 	}
// 	//postgresql.InsertDataToDB(ShortURL, url.LongURL, cookie.Value)
// 	w.Header().Set("Location", ShortURL)

// 	w.WriteHeader(201)
// 	fmt.Fprint(w, string(jsonRes))

// }

func GetAllUrls(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	Id = uuid.NewV4()
	if err != nil {

		cookie = &http.Cookie{
			Name:       "session",
			Value:      Id.String(),
			Path:       "",
			Domain:     "",
			Expires:    time.Time{},
			RawExpires: "",
			MaxAge:     0,
			Secure:     false,
			HttpOnly:   true,
			SameSite:   0,
			Raw:        "",
			Unparsed:   []string{},
		}

	}
	type languageStruct struct {
		Shorturl    string `json:"short_url"`
		Originalurl string `json:"original_url"`
		Userid      string `json:"userID"`
	}
	var x []*languageStruct
	var jsonRes []byte
	var result string
	w.Header().Set("content-type", "application/json")
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
		//result = append(result, line)
		if result == "" {
			result = line
		}
		if result != "" && line != "\n" {
			result = result + "," + line
		}

	}
	a := "[" + result + "]"
	//jsonRes, _ = json.Marshal(result)
	//fmt.Print(a)
	jsonRes = []byte(a)
	err = json.Unmarshal(jsonRes, &x)
	var x2 []*languageStruct

	for _, v := range x {
		//fmt.Print(v)
		if v.Userid == cookie.Value {
			x2 = append(x2, v)
		}

	}
	if x2 == nil {
		w.WriteHeader(http.StatusNoContent)
	}
	data, err := json.MarshalIndent(x2, " ", " ")

	//fmt.Println(string(data))
	//fmt.Print(x)
	w.Write(data)

	return
}
