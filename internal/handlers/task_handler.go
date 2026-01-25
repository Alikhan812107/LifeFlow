package handlers

import (
	"Assignment3/internal/models"
	"Assignment3/internal/service"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	result := h.service.Create(task)
	json.NewEncoder(w).Encode(result)
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.service.GetAll())
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	task, ok := h.service.GetByID(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	result, ok := h.service.Update(id, task)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if !h.service.Delete(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) ViewHTML(w http.ResponseWriter, r *http.Request) {
	tasks := h.service.GetAll()

	tmpl := template.Must(template.ParseFiles("templates/tasks.html"))
	tmpl.Execute(w, tasks)
}

func (h *TaskHandler) CreateFromHTML(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	title := r.FormValue("title")

	h.service.Create(models.Task{
		Title:  title,
		Done:   false,
		UserID: 1,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
