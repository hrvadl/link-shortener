package services

import (
	"github.com/google/uuid"
	"github.com/hrvadl/link-shortener/pkg/repository"
)

type Shortener interface {
	Shorten(URL string) (string, error)
	Get(uuid string) (string, error)
}

const TTL = 0

type URLShortener struct {
	linkRepo repository.LinkRepository
}

func NewURLShortener(linkRepo repository.LinkRepository) *URLShortener {
	return &URLShortener{
		linkRepo: linkRepo,
	}
}

func (e *URLShortener) Shorten(URL string) (string, error) {
	uuid := uuid.New().String()

	if err := e.linkRepo.Set(uuid, URL, TTL); err != nil {
		return "", err
	}

	return uuid, nil
}

func (e *URLShortener) Get(uuid string) (string, error) {
	return e.linkRepo.Get(uuid)
}
