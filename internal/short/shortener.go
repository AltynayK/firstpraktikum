package short

import (
	"strconv"

	"github.com/AltynayK/firstpraktikum/internal/service"
)

var Init *string

func GetBaseURL(a *string) {
	Init = a
}
func WriteShortURL(url string) string {

	return *Init + "/" + strconv.Itoa(service.WriteURLByID(url))

}
