package models

import "database/sql"

type exercise struct {
	id   int
	name string
}

type ExerciseRepo struct {
	db *sql.DB
}

type workout struct {
	exercises []exercise
	duration  float64
}

type WorkoutRepo struct {
	db *sql.DB
}
