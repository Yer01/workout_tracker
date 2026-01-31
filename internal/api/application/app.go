package application

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	router    http.Handler
	DB        *sql.DB
	container *Container
}

func New(db *sql.DB) *App {
	container := NewContainer(db)
	return &App{
		router:    routes(container),
		DB:        db,
		container: container,
	}
}

func (a *App) Start(ctx context.Context) error {

	log.Println("Starting server on port 8081...")
	if err := http.ListenAndServe("localhost:8081", a.router); err != nil {
		return fmt.Errorf("Error in setting up a server: %w", err)
	}

	return nil
}
