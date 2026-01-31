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
	mux.Route("/workouts", func(r chi.Router) {
		mux.Get("/{id}", cont.WorkoutHandler.Get)
		mux.Post("/", cont.WorkoutHandler.Create)
		mux.Put("/{id}", cont.WorkoutHandler.Update)
	})
	return mux
}

/*func loadWorkoutRouter(mux chi.Router) {

	mux.Get("/{id}", cont.WorkoutHandler.Get)
	mux.Post("/", handler.Create)
	mux.Put("/{id}", handler.Update)
}*/
