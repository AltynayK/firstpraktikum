package service

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

var (
	IDList map[int]string
	id     int = 0
)
var FilePath *string

type abs struct {
	ShortURL string `json:"short_url"`
	LongURL  string `json:"original_url"`
}

func ReadFile(a *string) {
	FilePath = a

	file, err := os.OpenFile(*FilePath, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Folder does not exist.")
		}
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if IDList == nil {
			IDList = make(map[int]string)
		}
		id++
		IDList[id] = line
	}
	return
}
func WriteURLByID(url string) int {
	if IDList == nil {
		IDList = make(map[int]string)
	}
	id++
	IDList[id] = url
	WriteToFile(url)
	return id
}

func GetURLFromID(id int) string {

	// remove row below

	return IDList[id]
}

func WriteToFile(LongURL string) {
	f, err := os.OpenFile(*FilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Folder does not exist.")
		}
	}
	defer f.Close()

	if _, err = f.WriteString(LongURL + "\n"); err != nil {
		log.Fatal("Folder does not exist.")
	}
}

// func WriteToFileAll(LongURL string) {
// 	f, err := os.OpenFile(*FilePath, os.O_APPEND|os.O_WRONLY, 0600)
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			log.Fatal("Folder does not exist.")
// 		}
// 	}
// 	defer f.Close()

// 	if _, err = f.WriteString(LongURL + "\n"); err != nil {
// 		log.Fatal("Folder does not exist.")
// 	}
// }

func MakeData(longURL string, shortURL string) {

	// var result []byte

	// okRes := abs{
	// 	ShortURL: longURL,
	// 	LongURL:  shortURL,
	// }

	// result = append(result, okRes)
	// fmt.Print(result)

	// ioutil.WriteFile("output.json", u1, 0644)
	var jsonBlob = []byte(`{ShortURL: shortURL, LongURL: longURL,}`)
	rankings := abs{
		LongURL:  longURL,
		ShortURL: shortURL,
	}
	err := json.Unmarshal(jsonBlob, &rankings)
	if err != nil {
		// nozzle.printError("opening config file", err.Error())
	}
	rankingsJson, _ := json.Marshal(rankings)

	file, err := os.OpenFile("output.json", os.O_APPEND|os.O_WRONLY, 0600)
	file.WriteString("\n")
	file.Write(rankingsJson)

}
