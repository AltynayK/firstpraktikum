package domain

type Url struct {
	LongURL  string `json:"LongUrl"`
	ShortURL string `json:"short_url"`
}

type PostResponse struct {
	Result string `json:"result"`
}
