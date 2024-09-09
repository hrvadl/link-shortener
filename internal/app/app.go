package app

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	"github.com/hrvadl/link-shortener/internal/config"
	linkdomain "github.com/hrvadl/link-shortener/internal/domain/link"
	"github.com/hrvadl/link-shortener/internal/storage/db"
	linkstorage "github.com/hrvadl/link-shortener/internal/storage/link"
	"github.com/hrvadl/link-shortener/internal/transport/link"
)

func New(cfg *config.Config, log *slog.Logger) *App {
	fiberSrv := fiber.New()

	redis := db.NewRedis(cfg.RedisAddr, cfg.RedisPassword)
	repo := linkstorage.NewRepo(redis)
	shortener := linkdomain.NewShortener(repo)
	handler := link.NewHandler(shortener)

	fiberSrv.Get("/:id", handler.HandleGetURL)
	fiberSrv.Post("/", handler.HandleShortenURL)

	app := &App{
		app: fiberSrv,
		log: log,
		cfg: cfg,
	}

	return app
}

type App struct {
	app *fiber.App
	log *slog.Logger
	cfg *config.Config
}

func (s *App) MustRun() {
	if err := s.Run(s.cfg.Port); err != nil {
		panic(err)
	}
}

func (s *App) Run(port string) error {
	s.log.Info("Server has been started", slog.String("port", port))
	return s.app.Listen(port)
}

func (s *App) StopGracefully() error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	return s.app.Shutdown()
}
