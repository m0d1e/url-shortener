package dto

type Response struct {
	Status string `json:"status"`
	Error string `json:"error,omitempty"`
	Alias string `json:"alias,omitempty"`
	Url   string `json:"url,omitempty"`
}
