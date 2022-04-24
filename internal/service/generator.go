package service

import (
	"bufio"
	"flag"
	"os"
)

var (
	IDList map[int]string
	id     int = 0
)
var FilePath string

func ReadFile() {
	FileStorage := flag.String("f", "text.txt", "FILE_STORAGE_PATH - путь до файла с сокращёнными URL")
	flag.Parse()
	if u, flg := os.LookupEnv("FILE_STORAGE_PATH"); flg {
		*FileStorage = u
	}
	FilePath = *FileStorage
	file, err := os.Open(FilePath)
	if err != nil {
		panic(err)
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
	WriteToFile(url)
	return id
}

func GetURLFromID(id int) string {

	// remove row below

	return IDList[id]
}

func WriteToFile(LongURL string) {
	f, err := os.OpenFile(FilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(LongURL + "\n"); err != nil {
		panic(err)
	}
}
