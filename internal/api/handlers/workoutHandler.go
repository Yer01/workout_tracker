package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Yer01/workout_tracker/internal/services"
	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
	workout_service services.WorkoutService
}

func NewWorkoutHandler(service services.WorkoutService) *WorkoutHandler {
	return &WorkoutHandler{
		workout_service: service,
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Workout tracker app")
}

func (wh *WorkoutHandler) ShowSingle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("ID is wrong type or format: %v", err)
		return
	}
	workout, err := wh.workout_service.GetSingle(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Problems with fetching data from database: %v", err)
		return
	}

	if err = json.NewEncoder(w).Encode(workout); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Problems with enoding data : %v", err)
		return
	}
}

func (wh *WorkoutHandler) ShowAll(w http.ResponseWriter, r *http.Request) {
	workouts, err := wh.workout_service.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Problems with fetching all data from database: %v", err)
		return
	}

	if err = json.NewEncoder(w).Encode(workouts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Can't encode all data : %v", err)
		return
	}
}

func (wh *WorkoutHandler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updated existing workout")
}

func (wh *WorkoutHandler) Create(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	content := r.FormValue("content")
	exercises := r.FormValue("exercises")
	duration, err := strconv.ParseFloat(r.FormValue("duration"), 64)
	if err != nil {
		http.Error(w, fmt.Errorf("Duration is wrong type or format: %v", err).Error(), http.StatusBadRequest)
		log.Printf("Duration is wrong type or format: %v", err)
		return
	}
	create_id, err := wh.workout_service.Create(name, content, exercises, duration)

	if err != nil || create_id == -1 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Problems with creating new workout: %v", err)
		return
	}
	resp := fmt.Sprintf("Workout #%d was succesfully created!", create_id)
	log.Println(resp)
	w.Write([]byte(resp))
}

func (wh *WorkoutHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
