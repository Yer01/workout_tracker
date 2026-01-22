package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App {
	return &App{
		router: routes(),
	}
}

func (a *App) Start(ctx context.Context) error {
	log.Println("Starting server on port 8081...")
	if err := http.ListenAndServe("localhost:8081", a.router); err != nil {
		return fmt.Errorf("Error in setting up a server: %w", err)
	}

	return nil
}
