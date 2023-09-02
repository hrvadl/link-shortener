package app

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/hrvadl/link-shortener/pkg/handlers"
	"github.com/hrvadl/link-shortener/pkg/repository"
	"github.com/hrvadl/link-shortener/pkg/services"
)

type Server struct {
	app *fiber.App
	log *slog.Logger
}

func New(l *slog.Logger) (*Server, error) {
	app := fiber.New()

	repo := repository.NewLinkRepo()
	service := services.NewURLShortener(repo)
	handler := handlers.NewLinkShortener(service)

	app.Get("/short/:hash", handler.HandleGetURL)
	app.Post("/", handler.HandleGetURL)

	server := &Server{
		app: app,
		log: l,
	}

	return server, nil
}

func (s *Server) Serve(port string) error {
	s.log.Info("Server has been started on port %v", port)
	return s.app.Listen(port)
}
