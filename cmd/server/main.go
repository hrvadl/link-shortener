package main

import (
	"log/slog"
	"os"

	"github.com/hrvadl/link-shortener/internal/app"
	"github.com/hrvadl/link-shortener/internal/config"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	cfg, err := config.NewFromEnv()
	if err != nil {
		log.Error("Failed to parse config", slog.Any("err", err))
		os.Exit(1)
	}

	app := app.New(cfg, log)

	go app.MustRun()
	if err := app.StopGracefully(); err != nil {
		log.Error("Failed to stop server gracefully", slog.Any("err", err))
		os.Exit(1)
	}
}
