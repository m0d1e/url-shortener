package service

type Repository interface {
	SaveUrl(originalURL, alias string) error
	GetUrl(alias string) (string, error)
	DeleteUrl(alias string) error
}
