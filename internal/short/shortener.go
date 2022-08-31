package short

import (
	"github.com/AltynayK/firstpraktikum/internal/service"
	"github.com/speps/go-hashids"
)

var Init *string

func GetBaseURL(a *string) {
	Init = a
}

//increment#1
func WriteShortURL(url string) string {
	service.WriteToFile(url)
	service.WriteURLByID(url)
	return Hash(url)

}

//сокращает url, используется при сохранении в бд
func Hash(url string) string {
	hd := hashids.NewData()
	hd.Salt = url
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{45, 434, 1313, 99})

	return *Init + "/" + e
}
