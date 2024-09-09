package link

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"errors"
)

func NewShortener(links LinksSource) *Shortener {
	return &Shortener{
		links: links,
	}
}

//go:generate mockgen -destination=./mocks/mock_source.go -package=mocks . LinksSource
type LinksSource interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

type Shortener struct {
	links LinksSource
}

func (e *Shortener) Shorten(ctx context.Context, url string) (string, error) {
	if url == "" {
		return "", ErrEmptyURL
	}

	hash := md5.Sum([]byte(url))
	hashedURL := base64.StdEncoding.EncodeToString(hash[:])
	if err := e.links.Set(ctx, hashedURL, url); err != nil {
		return "", errors.Join(ErrFailedToShorten, err)
	}

	return hashedURL, nil
}

func (e *Shortener) Get(ctx context.Context, hash string) (string, error) {
	shortened, err := e.links.Get(ctx, hash)
	if err != nil {
		return "", errors.Join(ErrNotFound, err)
	}

	return shortened, nil
}
