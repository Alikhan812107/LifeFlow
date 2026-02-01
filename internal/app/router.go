package app

import (
	"net/http"

	"Assignment3/internal/handlers"
)

func RegisterRoutes(
	taskHandler *handlers.TaskHandler,
	authHandler *handlers.AuthHandler,
) {
	// Task API
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			taskHandler.Create(w, r)
			return
		}
		if r.Method == http.MethodGet {
			taskHandler.GetAll(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/tasks/item", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			taskHandler.GetByID(w, r)
			return
		}
		if r.Method == http.MethodPut {
			taskHandler.Update(w, r)
			return
		}
		if r.Method == http.MethodDelete {
			taskHandler.Delete(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	// Auth API
	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)

	// HTML routes
	http.HandleFunc("/", taskHandler.ViewHTML)
	http.HandleFunc("/tasks/html", taskHandler.CreateFromHTML)
	http.HandleFunc("/tasks/toggle", taskHandler.ToggleTask)
	http.HandleFunc("/tasks/delete", taskHandler.DeleteFromHTML)
	http.HandleFunc("/tasks/update", taskHandler.UpdateFromHTML)
}
