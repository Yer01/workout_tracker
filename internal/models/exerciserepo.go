package models

import "database/sql"

type exercise struct {
	id   int
	name string
}

type ExerciseRepo struct {
	db *sql.DB
}
