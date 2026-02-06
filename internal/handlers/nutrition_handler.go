package handlers

import (
	"Assignment3/internal/models"
	"Assignment3/internal/service"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type NutritionHandler struct {
	service *service.NutritionService
}

func NewNutritionHandler(service *service.NutritionService) *NutritionHandler {
	return &NutritionHandler{service: service}
}

func (h *NutritionHandler) ViewHTML(w http.ResponseWriter, r *http.Request) {
	nutritions, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/nutrition.html"))
	tmpl.Execute(w, nutritions)
}

func (h *NutritionHandler) CreateFromHTML(w http.ResponseWriter, r *http.Request) {
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
	
	_, err = h.service.Create(nutrition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	http.Redirect(w, r, "/nutrition", http.StatusSeeOther)
}

func (h *NutritionHandler) GetAllJSON(w http.ResponseWriter, r *http.Request) {
	nutritions, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nutritions)
}
