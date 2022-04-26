package short

import (
	"strconv"

	"github.com/AltynayK/firstpraktikum/internal/service"
)

var Base *string

func GetBaseURL(a *string) {
	Base = a
}
func WriteShortURL(url string) string {

	return *Base + "/" + strconv.Itoa(service.WriteURLByID(url))

}
