package api

import (
	"10_rest/hw/entity"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var (
	tasks = []*entity.Task{
		{ID: 1, Title: "Купити хліб", Done: false},
		{ID: 2, Title: "Прочитати книгу", Done: true},
	}
	nextID = tasks[len(tasks)-1].ID + 1
)

func GetTaskList(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Printf("Не вдалося повернути список задач: %v", err)
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, task := range tasks {
		if task.ID == id {
			if err := json.NewEncoder(w).Encode(task); err != nil {
				log.Printf("Не вдалося повернути задачу: %v", err)
			}
			return
		}
	}

	http.NotFound(w, r)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.ID = nextID
	tasks = append(tasks, &task)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Printf("Не вдалося створити задачу: %v", err)
		return
	}

	nextID++
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedTask entity.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Done = updatedTask.Done
			if err := json.NewEncoder(w).Encode(tasks[i]); err != nil {
				log.Printf("Не вдалося оновити задачу: %v", err)
				return
			}
			return
		}
	}

	http.NotFound(w, r)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}

	http.NotFound(w, r)
}

func InitRoutes(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/tasks", GetTaskList).Methods(http.MethodGet)
	api.HandleFunc("/tasks/{id}", GetTask).Methods(http.MethodGet)
	api.HandleFunc("/tasks", CreateTask).Methods(http.MethodPost)
	api.HandleFunc("/tasks/{id}", UpdateTask).Methods(http.MethodPut)
	api.HandleFunc("/tasks/{id}", DeleteTask).Methods(http.MethodDelete)
}
