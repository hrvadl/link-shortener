package link

import (
	"context"
	"encoding/base64"
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
	id := base64.StdEncoding.EncodeToString([]byte(url))
	if err := e.links.Set(ctx, id, url); err != nil {
		return "", err
	}

	return id, nil
}

func (e *Shortener) Get(ctx context.Context, uuid string) (string, error) {
	return e.links.Get(ctx, uuid)
}
