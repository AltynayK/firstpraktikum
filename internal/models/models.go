package models

type URL struct {
	LongURL string `json:"url"`
	Result  string `json:"result"`
	//CorrelationID string `json:"correlation_id"`
}
type MultURL struct {
	CorrelationID string `json:"correlation_id"`
	Result        string `json:"short_url"`
}
type URLs struct {
	ShortURL      string `json:"short_url"`
	LongURL       string `json:"original_url"`
	CorrelationID string `json:"correlation_id"`
}
type DBUrl struct {
	id          int
	shorturl    string
	Originalurl string
	userid      string
}
type Abs struct {
	ShortURL      string `json:"short_url"`
	LongURL       string `json:"original_url"`
	UserID        string `json:"userID"`
	CorrelationID string `json:"correlation_id"`
}

type Dbid struct {
	Maxid int
}

type DBUrls struct {
	id          int
	Shorturl    string
	originalurl string
	userid      string
}

//type DBsql struct{}

type URLStruct struct {
	Shorturl    string `json:"short_url"`
	Originalurl string `json:"original_url"`
	Userid      string `json:"userID"`
}
