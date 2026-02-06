package application

import (
	"github.com/Yer01/workout_tracker/internal/api/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routes(cont *Container) *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)

	mux.Get("/", handlers.HomeHandler)
	mux.Route("/workouts", loadWorkoutRouter(cont))
	return mux
}

func loadWorkoutRouter(cont *Container) func(chi.Router) {
	return func(mux chi.Router) {
		mux.Get("/{id}", cont.WorkoutHandler.ShowSingle)
		mux.Get("/", cont.WorkoutHandler.ShowAll)
		mux.Post("/new", cont.WorkoutHandler.Create)
		mux.Put("/{id}", cont.WorkoutHandler.Update)
		mux.Delete("/{id}", cont.WorkoutHandler.Delete)
	}
}
