package shortener

import (
	"strconv"

	"github.com/AltynayK/firstpraktikum/internal/service"
)

func WriteShortURL(url string) string {
	return "http://localhost:8080/" + strconv.Itoa(service.WriteURLByID(url))

}
