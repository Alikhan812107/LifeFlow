package handlers

import (
	"Assignment3/internal/models"
	"Assignment3/internal/service"
	"encoding/json"
	"html/template"
	"net/http"
	"time"
)

type SleepHandler struct {
	service *service.SleepService
}

func NewSleepHandler(service *service.SleepService) *SleepHandler {
	return &SleepHandler{service: service}
}

func (h *SleepHandler) ViewHTML(w http.ResponseWriter, r *http.Request) {
	sleeps, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/sleep.html"))
	tmpl.Execute(w, sleeps)
}

func (h *SleepHandler) CreateFromHTML(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}
	
	wokeUpStr := r.FormValue("woke_up")
	sleptStr := r.FormValue("slept")
	
	wokeUp, err := time.Parse("2006-01-02T15:04", wokeUpStr)
	if err != nil {
		http.Error(w, "bad woke up time", http.StatusBadRequest)
		return
	}
	
	slept, err := time.Parse("2006-01-02T15:04", sleptStr)
	if err != nil {
		http.Error(w, "bad slept time", http.StatusBadRequest)
		return
	}
	
	sleep := models.Sleep{
		WokeUp:    wokeUp,
		Slept:     slept,
		UserID:    "user1",
		Timestamp: time.Now(),
	}
	
	_, err = h.service.Create(sleep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	http.Redirect(w, r, "/sleep", http.StatusSeeOther)
}

func (h *SleepHandler) GetAllJSON(w http.ResponseWriter, r *http.Request) {
	sleeps, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sleeps)
}
