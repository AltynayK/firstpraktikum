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
