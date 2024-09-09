package link

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrEmptyURL        = errors.New("url can not be empty")
	ErrFailedToShorten = errors.New("failed to shorten URL")
)
