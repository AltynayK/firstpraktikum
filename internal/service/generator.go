package service

import (
	"bufio"
	"fmt"
	"os"
)

var (
	IDList map[int]string
	id     int = 0
)
var FilePath *string

//increment#6
func ReadFile(a *string) {
	FilePath = a
	file, err := os.OpenFile(*FilePath, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Print("Folder does not exist.")
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

}

func WriteURLByID(url string) int {
	if IDList == nil {
		IDList = make(map[int]string)
	}
	id++
	IDList[id] = url

	return id
}

// func GetURLFromID(id int) string {
// 	return IDList[id]
// }

//increment#6
func WriteToFile(LongURL string) {
	f, err := os.OpenFile(*FilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Print("Folder does not exist.")
		}
	}
	defer f.Close()
	if _, err = f.WriteString(LongURL + "\n"); err != nil {
		fmt.Print("Folder does not exist.")
	}
}

// func MakeData(longURL string, shortURL string, userID string) {
// 	rankings := models.Abs{
// 		LongURL:  longURL,
// 		ShortURL: shortURL,
// 		UserID:   userID,
// 	}
// 	rankingsJSON, _ := json.Marshal(rankings)
// 	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			fmt.Print("Folder does not exist.")
// 		}
// 	}
// 	file.Write(rankingsJSON)
// 	file.WriteString("\n")
// }

// func MakeDataForMultipleCase(shortURL string, longURL string, userID string, correlationID string) {
// 	rankings := models.Abs{
// 		CorrelationID: correlationID,
// 		LongURL:       longURL,
// 		ShortURL:      shortURL,
// 		UserID:        userID,
// 	}
// 	rankingsJSON, _ := json.Marshal(rankings)
// 	file, err := os.OpenFile("outputMultiple.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			fmt.Print("Folder does not exist.")
// 		}
// 	}
// 	file.Write(rankingsJSON)
// 	file.WriteString("\n")
// }
