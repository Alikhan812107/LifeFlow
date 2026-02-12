package handlers

import (
	"Assignment3/internal/middleware"
	"Assignment3/internal/models"
	"Assignment3/internal/service"
	"encoding/json"
	"html/template"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoteHandler struct {
	service     *service.NoteService
	userService *service.UserService
}

func NewNoteHandler(service *service.NoteService, userService *service.UserService) *NoteHandler {
	return &NoteHandler{
		service:     service,
		userService: userService,
	}
}

func (h *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	result, err := h.service.Create(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *NoteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	notes, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func (h *NoteHandler) ViewHTML(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	
	user, err := h.userService.GetByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	userRole := user.Role
	if userRole == "" {
		userRole = "free"
	}
	
	notes, err := h.service.GetAllByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	data := struct {
		Notes []models.Note
		Role  string
		Count int
	}{
		Notes: notes,
		Role:  userRole,
		Count: len(notes),
	}
	
	tmpl := template.Must(template.ParseFiles("templates/notes.html"))
	tmpl.Execute(w, data)
}

func (h *NoteHandler) CreateFromHTML(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}
	
	userID, ok := middleware.GetUserID(r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	
	user, err := h.userService.GetByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	userRole := user.Role
	if userRole == "" {
		userRole = "free"
	}
	
	if userRole == "free" {
		notes, err := h.service.GetAllByUserID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(notes) >= 10 {
			http.Error(w, "Free users can only create 10 notes. Upgrade to premium for unlimited notes.", http.StatusForbidden)
			return
		}
	}
	
	title := r.FormValue("title")
	description := r.FormValue("description")
	if title == "" {
		http.Error(w, "need title", http.StatusBadRequest)
		return
	}
	note := models.Note{
		Title:       title,
		Description: description,
		UserID:      userID,
	}
	_, err = h.service.Create(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/notes", http.StatusSeeOther)
}

func (h *NoteHandler) UpdateFromHTML(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}
	
	userID, ok := middleware.GetUserID(r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	
	idStr := r.FormValue("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}
	title := r.FormValue("title")
	description := r.FormValue("description")
	
	if title == "" {
		http.Error(w, "need title", http.StatusBadRequest)
		return
	}
	
	note := models.Note{
		Title:       title,
		Description: description,
		UserID:      userID,
	}
	
	_, err = h.service.Update(id, note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/notes", http.StatusSeeOther)
}

func (h *NoteHandler) DeleteFromHTML(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}
	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/notes", http.StatusSeeOther)
}