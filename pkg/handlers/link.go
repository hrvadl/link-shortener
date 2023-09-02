package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hrvadl/link-shortener/pkg/services"
)

type LinkShortener struct {
	srv services.Shortener
}

func NewLinkShortener(srv services.Shortener) *LinkShortener {
	return &LinkShortener{
		srv: srv,
	}
}

func (s *LinkShortener) HandleGetURL(*fiber.Ctx) error {
	return nil
}

func (s *LinkShortener) HandleShortenURL(*fiber.Ctx) error {
	return nil
}
