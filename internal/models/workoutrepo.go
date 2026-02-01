package models

import "database/sql"

type Workout struct {
	workout_id int
	name       string
	content    string
	exercises  string //would be array of exercises later
	duration   float64
}

type WorkoutRepo struct {
	DB *sql.DB
}

func (wr *WorkoutRepo) Get(id int) (Workout, error) {
	quer := `SELECT * FROM workouts WHERE workout_id = $1`

	data := wr.DB.QueryRow(quer, id)

	var workout Workout
	err := data.Scan(&workout.workout_id, &workout.name, &workout.content, &workout.exercises, &workout.duration)

	if err != nil {
		return Workout{}, err
	}

	return workout, nil
}
