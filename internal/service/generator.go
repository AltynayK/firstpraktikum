package service

import (
	"bufio"
	"fmt"
	"os"
)

var (
	IDList map[int]string
	ID     int = 0
)

var filePath *string

//increment#6
func ReadFile(a *string) {
	filePath = a
	file, err := os.OpenFile(*filePath, os.O_RDONLY|os.O_CREATE, 0664)
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
		ID++
		IDList[ID] = line
	}

}

func WriteURLByID(url string) int {
	if IDList == nil {
		IDList = make(map[int]string)
	}
	ID++
	IDList[ID] = url
	return ID
}

func WriteToFile(LongURL string) {
	f, err := os.OpenFile(*filePath, os.O_APPEND|os.O_WRONLY, 0600)
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

func ScanFile(file *os.File) string {
	var result string
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
	return result
}
