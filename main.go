package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/salmomascarenhas/shortener-url-go/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute application: ", "error", err)
		return
	}
	slog.Info("All done")
}

func run() error {
	handler := api.NewHandler()

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
