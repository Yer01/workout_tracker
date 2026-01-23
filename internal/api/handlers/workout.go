package handlers

import (
	"fmt"
	"net/http"

	"github.com/Yer01/workout_tracker/internal/services"
)

type WorkoutHandler struct {
	service services.WorkoutService
}

func NewWorkoutHandler(service services.WorkoutService) *WorkoutHandler {
	return &WorkoutHandler{
		service: service,
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Workout tracker app")
}

func (wh *WorkoutHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a workout from db")
}

func (wh *WorkoutHandler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updated existing workout")
}

func (wh *WorkoutHandler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Created new workout")
}
