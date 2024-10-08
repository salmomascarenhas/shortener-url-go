package api

import (
	"encoding/json"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func sendJSON(w http.ResponseWriter, resp ShortenResponse, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal response: ", "error", err)
		sendJSON(w, ShortenResponse{
			Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
		return
	}
}

func NewHandler(db map[string]string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/shorten", handlePost(db))
	r.Get("/{code}", handleGet(db))

	return r
}

type ShortenBody struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func handlePost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body ShortenBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, ShortenResponse{
				Error: "invalid request body"},
				http.StatusBadRequest,
			)
			return
		}

		if _, err := url.Parse(body.URL); err != nil {
			sendJSON(w, ShortenResponse{Error: "invalid url"}, http.StatusBadRequest)
			return
		}

		code := genCode()
		db[code] = body.URL
		sendJSON(w, ShortenResponse{Data: code}, http.StatusCreated)
	}

}

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func genCode() string {
	const n = 8
	byts := make([]byte, n)
	for i := range n {
		byts[i] = characters[rand.IntN(len(characters))]
	}
	return string(byts)
}

func handleGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}

}
