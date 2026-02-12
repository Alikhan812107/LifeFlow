package handlers

import (
	"Assignment3/internal/middleware"
	"Assignment3/internal/models"
	"Assignment3/internal/service"
	"encoding/base64"
	"html/template"
	"io"
	"net/http"
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
	userID, ok := middleware.GetUserID(r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	
	tasks, err := h.taskService.GetAllByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	notes, err := h.noteService.GetAllByUserID(userID)
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

	user, err := h.userService.GetByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.TasksNum = len(tasks)
	user.NotesNum = len(notes)

	data := struct {
		User           models.User
		CompletedTasks int
		ActiveTasks    int
	}{
		User:           *user,
		CompletedTasks: completedTasks,
		ActiveTasks:    len(tasks) - completedTasks,
	}

	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"mul": func(a, b int) int { return a * b },
		"div": func(a, b int) int {
			if b == 0 {
				return 0
			}
			return a / b
		},
		"gt": func(a, b int) bool { return a > b },
	}

	tmpl := template.Must(template.New("profile.html").Funcs(funcMap).ParseFiles("templates/profile.html"))
	_ = tmpl.Execute(w, data)
}
func (h *UserHandler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	userID, ok := middleware.GetUserID(r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	base64Image := base64.StdEncoding.EncodeToString(bytes)

	err = h.userService.UpdateAvatar(userID, base64Image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}
