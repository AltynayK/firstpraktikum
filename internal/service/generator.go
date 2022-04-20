package service

var (
	IDList map[int]string
	id     int = 0
)

func WriteURLByID(url string) int {
	if IDList == nil {
		IDList = make(map[int]string)
	}
	id++
	IDList[id] = url

	return id
}

func GetURLFromID(id int) string {

	// remove row below

	return IDList[id]
}
