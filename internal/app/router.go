package app

import (
	"Assignment3/internal/handlers"
	"net/http"
)

func RegisterRoutes(taskHandler *handlers.TaskHandler, noteHandler *handlers.NoteHandler, userHandler *handlers.UserHandler, healthHandler *handlers.HealthHandler) {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			taskHandler.Create(w, r)
		}
		if r.Method == http.MethodGet {
			taskHandler.GetAll(w, r)
		}
	})

	http.HandleFunc("/tasks/item", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			taskHandler.GetByID(w, r)
		}
		if r.Method == http.MethodPut {
			taskHandler.Update(w, r)
		}
		if r.Method == http.MethodDelete {
			taskHandler.Delete(w, r)
		}
	})

	http.HandleFunc("/", taskHandler.ViewHTML)
	http.HandleFunc("/development/html", taskHandler.CreateFromHTML)
	http.HandleFunc("/development/toggle", taskHandler.ToggleTask)
	http.HandleFunc("/development/delete", taskHandler.DeleteFromHTML)
	http.HandleFunc("/development/update", taskHandler.UpdateFromHTML)

	http.HandleFunc("/notes", noteHandler.ViewHTML)
	http.HandleFunc("/notes/html", noteHandler.CreateFromHTML)
	http.HandleFunc("/notes/update", noteHandler.UpdateFromHTML)
	http.HandleFunc("/notes/delete", noteHandler.DeleteFromHTML)

	http.HandleFunc("/health", healthHandler.ViewHTML)
	http.HandleFunc("/health/sleep", healthHandler.CreateSleep)
	http.HandleFunc("/health/nutrition", healthHandler.CreateNutrition)
	http.HandleFunc("/health/activity", healthHandler.CreateActivity)

	http.HandleFunc("/profile", userHandler.ViewProfile)
	http.HandleFunc("/profile/avatar", userHandler.UploadAvatar)
}
