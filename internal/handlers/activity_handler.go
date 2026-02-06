package handlers

import (
	"Assignment3/internal/models"
	"Assignment3/internal/service"
	"encoding/json"
	"html/template"
	"net/http"
	"time"
)

type ActivityHandler struct {
	service *service.ActivityService
}

func NewActivityHandler(service *service.ActivityService) *ActivityHandler {
	return &ActivityHandler{service: service}
}

func (h *ActivityHandler) ViewHTML(w http.ResponseWriter, r *http.Request) {
	activities, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/activity.html"))
	tmpl.Execute(w, activities)
}

func (h *ActivityHandler) CreateFromHTML(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}
	
	description := r.FormValue("description")
	if description == "" {
		http.Error(w, "need description", http.StatusBadRequest)
		return
	}
	
	activity := models.Activity{
		Description: description,
		UserID:      "user1",
		Timestamp:   time.Now(),
	}
	
	_, err := h.service.Create(activity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	http.Redirect(w, r, "/activity", http.StatusSeeOther)
}

func (h *ActivityHandler) GetAllJSON(w http.ResponseWriter, r *http.Request) {
	activities, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}
