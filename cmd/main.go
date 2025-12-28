package main

import (
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	// logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	h := api.mount()
	api.run(h)

	if err := api.run(h); err != nil {
		slog.Error("Sever failed to start,", "error", err)
		os.Exit(1)
	}
}
