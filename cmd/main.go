package main

import (
	"log/slog"
	"os"

	"github.com/hrvadl/link-shortener/pkg/app"
)

const port = ":3000"

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app, err := app.New(log)

	if err != nil {
		log.Error("server could not be initialized: %v", err)
		return
	}

	if err := app.Serve(port); err != nil {
		log.Error("server couldn't listen on port %v: %v", port, err)
		return
	}
}
