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

func (wr *WorkoutRepo) Get(id int) (Workout_plan, error) {
	quer := `SELECT * FROM workout_plans WHERE id = $1`

	data := wr.db.QueryRow(quer, id)

	var workout Workout_plan
	err := data.Scan(&workout.ID, &workout.Total_duration, &workout.Name, &workout.User_id, &workout.Created_at, &workout.Updated_at, &workout.Description)

	if err != nil {
		return Workout_plan{}, err
	}

	return workout, nil
}

func (wr *WorkoutRepo) GetAll() ([]Workout_plan, error) {
	quer := `SELECT * FROM workout_plans`

	data, err := wr.db.Query(quer)

	if err != nil {
		return []Workout_plan{}, err
	}

	var workouts []Workout_plan

	for data.Next() {
		var workout Workout_plan

		err := data.Scan(&workout.ID, &workout.Total_duration, &workout.Name, &workout.User_id, &workout.Created_at, &workout.Updated_at, &workout.Description)
		if err != nil {
			return []Workout_plan{}, err
		}

		workouts = append(workouts, workout)
	}

	if data.Err() != nil {
		return []Workout_plan{}, data.Err()
	}

	return workouts, nil
}

func (wr *WorkoutRepo) Insert(name string, content string, exercises string, duration float64) (int64, error) {
	quer := `INSERT INTO workout_plans 
	(name, content, exercises, duration)
	VALUES($1, $2, $3, $4) RETURNING workout_id`

	data := wr.db.QueryRow(quer, name, content, exercises, duration)

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
