package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"url_shortener/internal/http/dto"
	//"url_shortener/internal/lib/slErr"
	rand "url_shortener/service/generateAlias"
)

type URLService struct {
	repo Repository
	log  *slog.Logger
}

func NewURLService(repo Repository, logger *slog.Logger) *URLService {
	return &URLService{repo: repo, log: logger}
}

func (s *URLService) Save(req dto.Request) (string, error) {
	err := validator.New().Struct(req)
	if err != nil {
		ve := err.(validator.ValidationErrors)
		for _, fieldErr := range ve {
			s.log.Warn("validation failed", slog.String("field", fieldErr.Field()), slog.String("rule", fieldErr.Tag()))
		}
		return "", fmt.Errorf("validation error: %w", err)
	}
	
	alias := req.Alias
	if alias == "" {
		alias = rand.GenerateAlias()
	}
	
	err = s.repo.SaveUrl(req.URL, alias)
	if err != nil {
		return "", err
	}
	
	return alias, nil
}

func (s *URLService) Get(alias string) (string, error) {
	return s.repo.GetUrl(alias)
}

func (s *URLService) Delete(alias string) error {
	return s.repo.DeleteUrl(alias)
}
