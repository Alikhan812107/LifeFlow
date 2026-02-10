package handlers

import (
	"Assignment3/internal/models"
	"Assignment3/internal/service"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type HealthHandler struct {
	sleepService     *service.SleepService
	nutritionService *service.NutritionService
	activityService  *service.ActivityService
}

func NewHealthHandler(sleepService *service.SleepService, nutritionService *service.NutritionService, activityService *service.ActivityService) *HealthHandler {
	return &HealthHandler{
		sleepService:     sleepService,
		nutritionService: nutritionService,
		activityService:  activityService,
	}
}

func (h *HealthHandler) ViewHTML(w http.ResponseWriter, r *http.Request) {
	sleeps, _ := h.sleepService.GetAll()
	nutritions, _ := h.nutritionService.GetAll()
	activities, _ := h.activityService.GetAll()
	
	data := struct {
		Sleeps     []models.Sleep
		Nutritions []models.Nutrition
		Activities []models.Activity
	}{
		Sleeps:     sleeps,
		Nutritions: nutritions,
		Activities: activities,
	}
	
	tmpl := template.Must(template.ParseFiles("templates/health.html"))
	tmpl.Execute(w, data)
}

func (h *HealthHandler) CreateSleep(w http.ResponseWriter, r *http.Request) {
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
	
	_, err = h.sleepService.Create(sleep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	http.Redirect(w, r, "/health", http.StatusSeeOther)
}

func (h *HealthHandler) CreateNutrition(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}
	
	caloriesStr := r.FormValue("calories")
	waterStr := r.FormValue("water")
	healthyStr := r.FormValue("healthy")
	
	calories, err := strconv.Atoi(caloriesStr)
	if err != nil {
		http.Error(w, "bad calories", http.StatusBadRequest)
		return
	}
	
	water, err := strconv.ParseFloat(waterStr, 64)
	if err != nil {
		http.Error(w, "bad water", http.StatusBadRequest)
		return
	}
	
	healthy := healthyStr == "yes"
	
	nutrition := models.Nutrition{
		Calories:  calories,
		Water:     water,
		Healthy:   healthy,
		UserID:    "user1",
		Timestamp: time.Now(),
	}
	
	_, err = h.nutritionService.Create(nutrition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	http.Redirect(w, r, "/health", http.StatusSeeOther)
}

func (h *HealthHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
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
	
	_, err := h.activityService.Create(activity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	http.Redirect(w, r, "/health", http.StatusSeeOther)
}
