package application

import (
	"context"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	ctx := context.Background()
	mux.Get("/home", homeHandler(ctx))

	return mux
}
