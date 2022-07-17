package short

import (
	"log"
	"strconv"

	"github.com/AltynayK/firstpraktikum/internal/repository"
	"github.com/AltynayK/firstpraktikum/internal/service"
)

var Init *string

func GetBaseURL(a *string) {
	Init = a
}
func WriteShortURL(url string) string {

	return *Init + "/" + strconv.Itoa(service.WriteURLByID(url))

}

type Dbid struct {
	maxid int
}

func MakeShortURLToDB(url string) string {
	db := repository.DB
	id := db.QueryRow("SELECT id FROM data ORDER BY id DESC LIMIT 1")
	alb := Dbid{}
	if err := id.Scan(&alb.maxid); err != nil {
		log.Fatal(err)
	}
	nextid := alb.maxid + 1
	return *Init + "/" + strconv.Itoa(nextid)
}
