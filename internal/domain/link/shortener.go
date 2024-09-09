package link

import (
	"context"

	"github.com/google/uuid"
)

func NewShortener(links LinksSource) *Shortener {
	return &Shortener{
		links: links,
	}
}

type LinksSource interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

type Shortener struct {
	links LinksSource
}

func (e *Shortener) Shorten(ctx context.Context, url string) (string, error) {
	uuid := uuid.New().String()
	if err := e.links.Set(ctx, uuid, url); err != nil {
		return "", err
	}

	return uuid, nil
}

func (e *Shortener) Get(ctx context.Context, uuid string) (string, error) {
	return e.links.Get(ctx, uuid)
}
