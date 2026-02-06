package models

import "database/sql"

type Exercise struct {
	ID          int
	Name        string
	Link        string
	Description string
	Duration    float64
}

type Exercise_musclegroup struct {
	Exercise_id    int
	Musclegroup_id int
	Isprimary      bool
}

type Musclegroup struct {
	ID   int
	Name string
	Slug string
}

type ExerciseRepo struct {
	db *sql.DB
}

func NewExerciseRepo(db *sql.DB) *ExerciseRepo {
	return &ExerciseRepo{
		db: db,
	}
}
