package models

import "database/sql"

type workout struct {
	exercises []exercise
	duration  float64
}

type WorkoutRepo struct {
	DB *sql.DB
}
