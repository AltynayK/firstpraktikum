package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/models"
	"github.com/AltynayK/firstpraktikum/internal/repository"
	"github.com/AltynayK/firstpraktikum/internal/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Handler struct {
	config *app.Config
	repo   repository.Repo
	ch     chan int
}

func NewHandler(config *app.Config) *Handler {

	return &Handler{
		config: config,
		repo:   repository.New(config),
	}
}
func (s *Handler) Run(config *app.Config) {
	mux := s.InitHandlers()

	srv := http.Server{
		Addr:    config.ServerAddress,
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Print(err)
	}

}

func (s *Handler) InitHandlers() *mux.Router {

	router := mux.NewRouter()
	router.Use(Decompress)
	router.Use(GzipHandler)
	router.Use(SetCookie)
	//router.Use(handler.CheckCookie)

	router.HandleFunc("/", s.PostText).Methods("POST")
	router.HandleFunc("/api/shorten", s.PostJSON).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", s.Get).Methods("GET")
	router.HandleFunc("/api/user/urls", s.GetAllUrls).Methods("GET")
	router.HandleFunc("/ping", s.CheckConnection).Methods("GET")
	router.HandleFunc("/api/shorten/batch", s.PostMultipleUrls).Methods("POST")
	router.HandleFunc("/api/user/urls", s.DeleteUrls).Methods("DELETE")
	return router
}

func (s *Handler) PostJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var url models.URL
	var jsonRes []byte
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	a := r.Context().Value(userCtxKey).(string)
	shortURL := s.repo.MakeShortURL(url.LongURL)
	ok := s.repo.InsertData(shortURL, url.LongURL, a)
	if !ok {
		shortURL = s.repo.ReturnShortURL(url.LongURL)
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	okRes := models.URL{
		Result: shortURL,
	}
	if jsonRes, err = json.Marshal(okRes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "response json marshal err")
		return
	}
	w.Header().Set("Location", shortURL)
	fmt.Fprint(w, string(jsonRes))
}

//increment#1
func (s *Handler) PostText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "plain/text")
	url, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	longURL := string(url)
	shortURL := s.repo.MakeShortURL(longURL)
	a := r.Context().Value(userCtxKey).(string)
	//repo := repository.New()
	ok := s.repo.InsertData(shortURL, longURL, a)
	if !ok {
		shortURL = s.repo.ReturnShortURL(longURL)
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	w.Header().Set("Location", shortURL)
	w.Write([]byte(shortURL))
}

//increment#1
func (s *Handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var longURL string
	status := true
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	b, err := strconv.Atoi(id)
	if err != nil && b < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//repo := repository.New()
	longURL = s.repo.GetLongURLByID(b)
	status = s.repo.CheckDeletion(b)
	if !status {
		w.WriteHeader(http.StatusGone)
	}
	w.Header().Set("Location", longURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
	fmt.Fprint(w)
}

//increment#10
func (s *Handler) CheckConnection(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")

	ok := repository.Ping()
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//increment#9
func (s *Handler) GetAllUrls(w http.ResponseWriter, r *http.Request) {
	var x []*models.URLStruct
	var jsonRes []byte
	w.Header().Set("content-type", "application/json")
	file, err := os.OpenFile("./output.txt", os.O_RDONLY|os.O_CREATE, 0664)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNoContent)
		}
	}
	a := "[" + service.ScanFile(file) + "]"
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

//increment#12
func (s *Handler) PostMultipleUrls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var url []models.URLs
	content, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(content, &url)
	if err != nil {
		fmt.Print(err)
		return
	}
	var okRes models.MultURL
	var jsonRes []byte
	var JSONArray []models.MultURL
	//repo := repository.New()
	for _, value := range url {
		a := r.Context().Value(userCtxKey).(string)
		shortURL := s.repo.MakeShortURL(value.LongURL)

		ok := s.repo.InsertMultipleData(shortURL, value.LongURL, a, okRes.CorrelationID)
		if !ok {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
		shortURL = s.repo.ReturnShortURL(value.LongURL)
		okRes = models.MultURL{
			CorrelationID: value.CorrelationID,
			Result:        shortURL,
		}
		JSONArray = append(JSONArray, okRes)
	}
	if jsonRes, err = json.Marshal(JSONArray); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "response json marshal err")
		return
	}

	fmt.Fprint(w, string(jsonRes))
}

func (s *Handler) DeleteUrls(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	var url []string
	var slice []int
	content, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(content, &url)
	if err != nil {
		fmt.Print(err)
		return
	}
	for _, value := range url {
		a, err := strconv.Atoi(value)
		if err != nil {
			fmt.Print("error")
		}
		slice = append(slice, a)
	}
	s.repo.DeleteMultiple(slice)
	w.WriteHeader(http.StatusAccepted)
}
