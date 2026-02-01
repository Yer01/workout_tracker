package services

import (
	"github.com/Yer01/workout_tracker/internal/models"
)

type WorkoutService interface {
	GetSingle(int) (models.Workout, error)
}

type workoutservice struct {
	wr *models.WorkoutRepo
}

func NewWorkoutService(wr *models.WorkoutRepo) WorkoutService {
	return &workoutservice{
		wr: wr,
	}
}

func (ws *workoutservice) GetSingle(id int) (models.Workout, error) {
	workout, err := ws.wr.Get(id)
	if err != nil {
		return models.Workout{}, err
	}

	return workout, nil
}
