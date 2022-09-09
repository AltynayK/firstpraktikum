package short

import (
	"strconv"

	"github.com/AltynayK/firstpraktikum/internal/models"
	"github.com/AltynayK/firstpraktikum/internal/repository"
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
	return *Init + "/" + strconv.Itoa(service.WriteURLByID(url))

}

//
func Hash(url string) string {
	hd := hashids.NewData()
	hd.Salt = url
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{45, 434, 1313, 99})
	return *Init + "/" + e
}

//increment#13
func MakeShortURLToDB(url string) string {
	db := repository.DB
	id := db.QueryRow("SELECT id FROM data ORDER BY id DESC LIMIT 1")
	alb := models.Dbid{}
	if err := id.Scan(&alb.Maxid); err != nil {
		alb.Maxid = 0
	}
	nextid := alb.Maxid + 1

	return *Init + "/" + strconv.Itoa(nextid)
	//return *Init + "/" + Hash(url)
}
