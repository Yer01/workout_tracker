package models

import (
	"database/sql"
	"time"
)

type Workout_plan struct {
	ID             int
	Total_duration float64
	Name           string
	User_id        int
	Created_at     time.Time
	Updated_at     time.Time
	Description    string
}

type Workout_plan_user struct {
	User_id    int
	Workout_id int
	Addet_at   time.Time
}

type Workout_plan_exercise struct {
	Workout_id  int
	Exercise_id int
	Order_index int
	Sets        int
	Reps        int
	Weight      float64
	Notes       string
}

type WorkoutRepo struct {
	db *sql.DB
}

func NewWorkoutRepo(db *sql.DB) *WorkoutRepo {
	return &WorkoutRepo{
		db: db,
	}
}
