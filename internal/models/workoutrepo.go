package models

import "database/sql"

type Workout struct {
	Workout_id int
	Name       string
	Content    string
	Exercises  string //would be array of exercises later
	Duration   float64
}

type WorkoutRepo struct {
	DB *sql.DB
}

func (wr *WorkoutRepo) Get(id int) (Workout, error) {
	quer := `SELECT * FROM workouts WHERE workout_id = $1`

	data := wr.DB.QueryRow(quer, id)

	var workout Workout
	err := data.Scan(&workout.Workout_id, &workout.Name, &workout.Content, &workout.Exercises, &workout.Duration)

	if err != nil {
		return Workout{}, err
	}

	return workout, nil
}

func (wr *WorkoutRepo) GetAll() ([]Workout, error) {
	quer := `SELECT * FROM workouts`

	data, err := wr.DB.Query(quer)

	if err != nil {
		return []Workout{}, err
	}

	var workouts []Workout

	for data.Next() {
		var workout Workout

		err = data.Scan(&workout.Workout_id, &workout.Name, &workout.Content, &workout.Exercises, &workout.Duration)
		if err != nil {
			return []Workout{}, err
		}

		workouts = append(workouts, workout)
	}

	if data.Err() != nil {
		return []Workout{}, data.Err()
	}

	return workouts, nil
}

func (wr *WorkoutRepo) Insert(name string, content string, exercises string, duration float64) (int64, error) {
	quer := `INSERT INTO workouts 
	(name, content, exercises, duration)
	VALUES($1, $2, $3, $4) RETURNING workout_id`

	data := wr.DB.QueryRow(quer, name, content, exercises, duration)

	if data.Err() != nil {
		return -1, data.Err()
	}

	var id int64

	err := data.Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil

}
