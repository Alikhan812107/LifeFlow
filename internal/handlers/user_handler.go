package handlers

import (
	"Assignment3/internal/middleware"
	"Assignment3/internal/models"
	"Assignment3/internal/service"
	"html/template"
	"net/http"
	
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	taskService *service.TaskService
	noteService *service.NoteService
	userService *service.UserService
}

func NewUserHandler(taskService *service.TaskService, noteService *service.NoteService, userService *service.UserService) *UserHandler {
	return &UserHandler{
		taskService: taskService,
		noteService: noteService,
		userService: userService,
	}
}

func (h *UserHandler) ViewProfile(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}
	
	user, err := h.userService.GetByID(objectID)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	
	tasks, err := h.taskService.GetAll(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	notes, err := h.noteService.GetAll(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	completedTasks := 0
	for _, task := range tasks {
		if task.Done {
			completedTasks++
		}
	}
	
	user.TasksNum = len(tasks)
	user.NotesNum = len(notes)
	
	data := struct {
		User           models.User
		CompletedTasks int
		ActiveTasks    int
	}{
		User:           user,
		CompletedTasks: completedTasks,
		ActiveTasks:    len(tasks) - completedTasks,
	}
	
	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"mul": func(a, b int) int { return a * b },
		"div": func(a, b int) int { 
			if b == 0 { return 0 }
			return a / b 
		},
		"gt": func(a, b int) bool { return a > b },
	}
	
	tmpl := template.Must(template.New("profile.html").Funcs(funcMap).ParseFiles("templates/profile.html"))
	tmpl.Execute(w, data)
}