package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler() http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/shorten", HandlePost)
	r.Get("/{code}", HandleGet)

	return r
}

type ShortenBoy struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func HandlePost(w http.ResponseWriter, r *http.Request) {

}

func HandleGet(w http.ResponseWriter, r *http.Request) {

}
