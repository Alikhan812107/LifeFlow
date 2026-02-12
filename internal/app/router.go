package app

import (
	"Assignment3/internal/handlers"
	"Assignment3/internal/middleware"
	"net/http"
)

func RegisterRoutes(taskHandler *handlers.TaskHandler, noteHandler *handlers.NoteHandler, userHandler *handlers.UserHandler, healthHandler *handlers.HealthHandler, authHandler *handlers.AuthHandler) {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			authHandler.ShowRegister(w, r)
		} else if r.Method == http.MethodPost {
			authHandler.Register(w, r)
		}
	})
	
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			authHandler.ShowLogin(w, r)
		} else if r.Method == http.MethodPost {
			authHandler.Login(w, r)
		}
	})
	
	http.HandleFunc("/logout", authHandler.Logout)
	
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

	http.HandleFunc("/", middleware.RequireAuth(taskHandler.ViewHTML))
	http.HandleFunc("/development/html", middleware.RequireAuth(taskHandler.CreateFromHTML))
	http.HandleFunc("/development/toggle", middleware.RequireAuth(taskHandler.ToggleTask))
	http.HandleFunc("/development/delete", middleware.RequireAuth(taskHandler.DeleteFromHTML))
	http.HandleFunc("/development/update", middleware.RequireAuth(taskHandler.UpdateFromHTML))

	http.HandleFunc("/notes", middleware.RequireAuth(noteHandler.ViewHTML))
	http.HandleFunc("/notes/html", middleware.RequireAuth(noteHandler.CreateFromHTML))
	http.HandleFunc("/notes/update", middleware.RequireAuth(noteHandler.UpdateFromHTML))
	http.HandleFunc("/notes/delete", middleware.RequireAuth(noteHandler.DeleteFromHTML))

	http.HandleFunc("/health", middleware.RequireAuth(healthHandler.ViewHTML))
	http.HandleFunc("/health/sleep", middleware.RequireAuth(healthHandler.CreateSleep))
	http.HandleFunc("/health/nutrition", middleware.RequireAuth(healthHandler.CreateNutrition))
	http.HandleFunc("/health/activity", middleware.RequireAuth(healthHandler.CreateActivity))

	http.HandleFunc("/profile", middleware.RequireAuth(userHandler.ViewProfile))
	http.HandleFunc("/profile/avatar", middleware.RequireAuth(userHandler.UploadAvatar))
}
