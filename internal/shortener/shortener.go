package shortener

import (
	"os"
	"strconv"

	"github.com/AltynayK/firstpraktikum/internal/service"
)

func WriteShortURL(url string) string {
	base := os.Getenv("BASE_URL")
	return base + "/" + strconv.Itoa(service.WriteURLByID(url))

}
