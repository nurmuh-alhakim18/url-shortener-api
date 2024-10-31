package url

import "time"

type URL struct {
	ID             int       `json:"id"`
	CustomAlias    string    `json:"custom_alias"`
	OriginalUrl    string    `json:"original_url"`
	CreatedAt      time.Time `json:"created_at"`
	ExpirationDate time.Time `json:"expiration_date"`
}

type ShortenURLReq struct {
	OriginalUrl string `json:"original_url"`
	CustomAlias string `json:"custom_alias"`
}

type ShortenURLResp struct {
	GeneratedLink string `json:"generated_link"`
}
