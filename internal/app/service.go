package app

import (
	"encoding/base64"
	"fmt"
)

func ShortURL(url string) string {
	ShortUrl := base64.StdEncoding.EncodeToString([]byte(url))
	//fmt.Println(ShortUrl)
	return ShortUrl
}

func LongURL(url string) string {
	LongUrl, err := base64.StdEncoding.DecodeString(url)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(LongUrl))
	return string(LongUrl)

}
