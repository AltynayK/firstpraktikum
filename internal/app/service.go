package app

import (
	"encoding/base64"
	"fmt"
)

func ShortURL(url string) string {
	ShortURL := base64.StdEncoding.EncodeToString([]byte(url))
	//fmt.Println(ShortUrl)
	return ShortURL
}

func LongURL(url string) string {
	LongURL, err := base64.StdEncoding.DecodeString(url)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(LongUrl))
	return string(LongURL)

}
