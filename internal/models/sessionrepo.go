package models

import "time"

type WorkoutSession struct {
	ID           int
	User_id      int
	Workout_id   int
	Started_at   time.Time
	Completed_at time.Time
	Notes        string
}

type Session_set struct {
	Set_num     int
	Session_id  int
	Exercise_id int
	Reps        int
	Weight      float64
}
