package services

import "database/sql"

type WorkoutService interface {
}

type workoutservice struct {
	db *sql.DB
}

func NewWorkoutService(db *sql.DB) WorkoutService {
	return workoutservice{
		db: db,
	}
}
