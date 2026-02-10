package app

import (
	"Assignment3/internal/handlers"
	"Assignment3/internal/middleware"
	"net/http"
)

func RegisterRoutes(taskHandler *handlers.TaskHandler, noteHandler *handlers.NoteHandler, userHandler *handlers.UserHandler, healthHandler *handlers.HealthHandler, authHandler *handlers.AuthHandler) {
	http.HandleFunc("/register", authHandler.ViewRegister)
	http.HandleFunc("/login", authHandler.ViewLogin)
	http.HandleFunc("/logout", authHandler.Logout)
	
	http.HandleFunc("/api/register", authHandler.RegisterJSON)
	http.HandleFunc("/api/login", authHandler.LoginJSON)
	
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

	http.HandleFunc("/", middleware.AuthMiddleware(taskHandler.ViewHTML))
	http.HandleFunc("/development/html", middleware.AuthMiddleware(taskHandler.CreateFromHTML))
	http.HandleFunc("/development/toggle", middleware.AuthMiddleware(taskHandler.ToggleTask))
	http.HandleFunc("/development/delete", middleware.AuthMiddleware(taskHandler.DeleteFromHTML))
	http.HandleFunc("/development/update", middleware.AuthMiddleware(taskHandler.UpdateFromHTML))

	http.HandleFunc("/notes", middleware.AuthMiddleware(noteHandler.ViewHTML))
	http.HandleFunc("/notes/html", middleware.AuthMiddleware(noteHandler.CreateFromHTML))
	http.HandleFunc("/notes/update", middleware.AuthMiddleware(noteHandler.UpdateFromHTML))
	http.HandleFunc("/notes/delete", middleware.AuthMiddleware(noteHandler.DeleteFromHTML))

	http.HandleFunc("/health", middleware.AuthMiddleware(healthHandler.ViewHTML))
	http.HandleFunc("/health/sleep", middleware.AuthMiddleware(healthHandler.CreateSleep))
	http.HandleFunc("/health/nutrition", middleware.AuthMiddleware(healthHandler.CreateNutrition))
	http.HandleFunc("/health/activity", middleware.AuthMiddleware(healthHandler.CreateActivity))

	http.HandleFunc("/profile", middleware.AuthMiddleware(userHandler.ViewProfile))
}
