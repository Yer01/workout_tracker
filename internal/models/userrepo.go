package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id       int
	Username string
	Email    string
}

type Workout_schedule struct {
	Workout_id    int
	User_id       int
	Schedule_date time.Time
	Status        string
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}
