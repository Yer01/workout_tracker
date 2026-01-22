package handlers

import (
	"context"
	"fmt"

	"github.com/Yer01/workout_tracker/internal/services"
)

type WorkoutHandler struct {
	service services.WorkoutService
}

func NewWorkoutHandler(service services.WorkoutService) *WorkoutHandler {
	return &WorkoutHandler{
		service: service,
	}
}

func (wh *WorkoutHandler) homeHandler(ctx *context.Context) {
	fmt.Println("Workout tracker app")
}
