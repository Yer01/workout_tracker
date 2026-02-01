package services

import (
	"github.com/Yer01/workout_tracker/internal/models"
)

type WorkoutService interface {
	GetSingle(int) (models.Workout, error)
	GetAll() ([]models.Workout, error)
	Create(string, string, string, float64) (int64, error)
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

func (ws *workoutservice) GetAll() ([]models.Workout, error) {
	workouts, err := ws.wr.GetAll()

	if err != nil {
		return []models.Workout{}, err
	}

	return workouts, nil
}

func (ws *workoutservice) Create(name string, content string, exercises string, duration float64) (int64, error) {
	create_id, err := ws.wr.Insert(name, content, exercises, duration)

	if err != nil {
		return -1, err
	}

	return create_id, nil
}
