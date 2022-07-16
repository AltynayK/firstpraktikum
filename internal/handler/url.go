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

	"github.com/AltynayK/firstpraktikum/internal/service"
	"github.com/AltynayK/firstpraktikum/internal/short"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type URL struct {
	LongURL       string `json:"url"`
	Result        string `json:"result"`
	CorrelationID string `json:"correlation_id"`
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
	a := r.Context().Value(userCtxKey).(string)
	service.MakeData(url.LongURL, ShortURL, a)
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

	url, err := io.ReadAll(r.Body)
	// обрабатываем ошибку
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	longURL := string(url)

	shortURL := short.WriteShortURL(longURL)
	//

	a := r.Context().Value(userCtxKey).(string)
	service.MakeData(longURL, shortURL, a)

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

type Posts []URLs

func PostMultipleUrls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var url Posts

	content, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(content, &url)
	if err != nil {
		//log.Fatal(err)
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
		service.MakeDataForMultipleCase(value.LongURL, ShortURL, a, okRes.CorrelationID)
		//jsonRes, _ = json.Marshal(okRes)
		// if jsonRes, err = json.Marshal(okRes); err != nil {
		// 	w.WriteHeader(500)
		// 	fmt.Fprintf(w, "response json marshal err")
		// 	return
		// }
		JsonArray = append(JsonArray, okRes)
		//fmt.Print("`correlation_id:`", value.CorrelationID, value.LongURL)
	}
	if jsonRes, err = json.Marshal(JsonArray); err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "response json marshal err")
		return
	}
	w.WriteHeader(201)
	fmt.Fprint(w, string(jsonRes))
	//fmt.Fprint(w, res)

}
