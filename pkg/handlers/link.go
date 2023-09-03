package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hrvadl/link-shortener/pkg/services"
)

type postBody struct {
	URL string `json:"url"`
}

type LinkShortener struct {
	srv services.Shortener
}

func NewLinkShortener(srv services.Shortener) *LinkShortener {
	return &LinkShortener{
		srv: srv,
	}
}

func (s *LinkShortener) HandleGetURL(ctx *fiber.Ctx) error {
	uuid := ctx.Params("id")
	url, err := s.srv.Get(uuid)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Redirect(url)
}

func (s *LinkShortener) HandleShortenURL(ctx *fiber.Ctx) error {
	var body postBody

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	shortened, err := s.srv.Shorten(body.URL)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendString(shortened)
}
