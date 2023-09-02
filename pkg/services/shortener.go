package services

import (
	"github.com/google/uuid"
	"github.com/hrvadl/link-shortener/pkg/repository"
)

type Shortener interface {
	Shorten(URL string) (string, error)
	GetDesiredURL(uuid string) (string, error)
}

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

}

func (e *URLShortener) GetDesiredURL(uuid string) (string, error) {

}
