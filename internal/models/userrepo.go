package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id       int
	Username string
	Email    string
	Password string
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

func (ur *UserRepo) Select(email string) (int, error) {
	quer := `SELECT id FROM users WHERE email = $1`

	res := ur.db.QueryRow(quer, email)

	var id int = 0

	if err := res.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (ur *UserRepo) Insert(email string, password string, name string) (int, error) {
	quer := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`

	res := ur.db.QueryRow(quer, name, email, password)

	id := 0

	if err := res.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}
