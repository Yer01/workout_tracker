package application

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Yer01/workout_tracker/internal/api/handlers"
	"github.com/Yer01/workout_tracker/internal/models"
	"github.com/Yer01/workout_tracker/internal/services"
)

type App struct {
	router http.Handler
	DB     *sql.DB
}

func New() *App {
	return &App{
		router: routes(),
	}
}

func (a *App) Start(ctx context.Context) error {

	wrep := &models.WorkoutRepo{
		DB: a.DB,
	}
	wservice := services.NewWorkoutService(wrep)
	whandler := handlers.NewWorkoutHandler(wservice)

	log.Println("Starting server on port 8081...")
	if err := http.ListenAndServe("localhost:8081", a.router); err != nil {
		return fmt.Errorf("Error in setting up a server: %w", err)
	}

	return nil
}
