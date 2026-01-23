package application

import (
	"github.com/Yer01/workout_tracker/internal/api/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	mux.Get("/", handlers.HomeHandler)
	mux.Route("/workouts", loadWorkoutRouter)
	return mux
}

func loadWorkoutRouter(mux chi.Router) {
	handler := &handlers.WorkoutHandler{}

	mux.Get("/{id}", handler.Get)
	mux.Post("/", handler.Create)
	mux.Put("/{id}", handler.Update)
}
