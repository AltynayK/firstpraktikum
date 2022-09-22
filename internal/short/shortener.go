package short

import (
	"math/rand"
	"strconv"

	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/repository"
)

//increment#1
func WriteShortURL(url string) string {

	return app.NewConfig().BaseURL + "/" + strconv.Itoa(repository.InsertData(url))
}

// func Hash(url string) string {
// 	hd := hashids.NewData()
// 	hd.Salt = url
// 	hd.MinLength = 30
// 	h, _ := hashids.NewWithData(hd)
// 	e, _ := h.Encode([]int{45, 434, 1313, 99})
// 	return *Init + "/" + e
// }

//increment#13

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func shorting() string {
	b := make([]byte, 5)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
