package application

import (
	"database/sql"

	"github.com/Yer01/workout_tracker/internal/api/handlers"
	"github.com/Yer01/workout_tracker/internal/models"
	"github.com/Yer01/workout_tracker/internal/services"
)

type Container struct {
	WorkoutHandler *handlers.WorkoutHandler
}

func NewContainer(db *sql.DB) *Container {
	wrepo := &models.WorkoutRepo{
		DB: db,
	}
	wservice := services.NewWorkoutService(wrepo)

	whandler := handlers.NewWorkoutHandler(wservice)

	return &Container{
		WorkoutHandler: whandler,
	}
}
