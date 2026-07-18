package application

import (
	"database/sql"

	"github.com/Yer01/workout_tracker/internal/api/handlers"
	"github.com/Yer01/workout_tracker/internal/models"
	"github.com/Yer01/workout_tracker/internal/services"
)

type Container struct {
	WorkoutHandler *handlers.WorkoutHandler
	AuthHandler    *handlers.AuthHandler
}

func NewContainer(db *sql.DB) *Container {
	wrepo := models.NewWorkoutRepo(db)
	wservice := services.NewWorkoutService(wrepo)
	whandler := handlers.NewWorkoutHandler(wservice)

	urepo := models.NewUserRepo(db)
	authservice := services.NewAuthService(urepo)
	authhandler := handlers.NewAuthHandler(authservice)

	return &Container{
		WorkoutHandler: whandler,
		AuthHandler:    authhandler,
	}
}
