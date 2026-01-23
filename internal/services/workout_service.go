package services

import (
	"github.com/Yer01/workout_tracker/internal/models"
)

type WorkoutService interface {
}

type workoutservice struct {
	wr *models.WorkoutRepo
}

func NewWorkoutService(wr *models.WorkoutRepo) WorkoutService {
	return workoutservice{
		wr: wr,
	}
}
