package link

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func NewHandler(shortener Shortener) *Handler {
	return &Handler{
		shortener: shortener,
	}
}

type Shortener interface {
	Get(ctx context.Context, uuid string) (string, error)
	Shorten(ctx context.Context, url string) (string, error)
}

type Handler struct {
	shortener Shortener
}

func (s *Handler) HandleGetURL(ctx *fiber.Ctx) error {
	uuid := ctx.Params("id")
	url, err := s.shortener.Get(ctx.Context(), uuid)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return ctx.Redirect(url)
}

type postBody struct {
	URL string `json:"url"`
}

func (s *Handler) HandleShortenURL(ctx *fiber.Ctx) error {
	var body postBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	shortened, err := s.shortener.Shorten(ctx.Context(), body.URL)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendString(shortened)
}
