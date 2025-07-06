package dto

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty" validate:"omitempty,alphanum"`
}
