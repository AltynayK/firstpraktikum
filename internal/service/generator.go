package service

import (
	"bufio"
	"log"
	"os"
)

var (
	IDList map[int]string
	id     int = 0
)
var FilePath *string

func ReadFile(a *string) {
	FilePath = a
	//FileStorage := flag.String("f", "texts.txt", "FILE_STORAGE_PATH - путь до файла с сокращёнными URL")
	//flag.Parse()
	// if u, flg := os.LookupEnv("FILE_STORAGE_PATH"); flg {
	// 	*FileStorage = u
	// }
	// FilePath = *FileStorage
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
