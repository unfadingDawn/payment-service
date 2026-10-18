package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) error {
	var err error
	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		_, err = w.Write([]byte("cool"))
	})
	return err
}
