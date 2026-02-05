package handlers

import (
	"Assignment3/internal/models"
	"Assignment3/internal/service"
	"html/template"
	"net/http"
)

type UserHandler struct {
	taskService *service.TaskService
	noteService *service.NoteService
}

func NewUserHandler(taskService *service.TaskService, noteService *service.NoteService) *UserHandler {
	return &UserHandler{
		taskService: taskService,
		noteService: noteService,
	}
}

func (h *UserHandler) ViewProfile(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.taskService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	notes, err := h.noteService.GetAll()
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
	
	user := models.User{
		ID:       "user1",
		Name:     "John Student",
		Email:    "john@student.com",
		TasksNum: len(tasks),
		NotesNum: len(notes),
	}
	
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