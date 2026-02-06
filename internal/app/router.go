package app

import (
	"Assignment3/internal/handlers"
	"net/http"
)

func RegisterRoutes(taskHandler *handlers.TaskHandler, noteHandler *handlers.NoteHandler, userHandler *handlers.UserHandler, sleepHandler *handlers.SleepHandler, nutritionHandler *handlers.NutritionHandler, activityHandler *handlers.ActivityHandler) {
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
	http.HandleFunc("/tasks/html", taskHandler.CreateFromHTML)
	http.HandleFunc("/tasks/toggle", taskHandler.ToggleTask)
	http.HandleFunc("/tasks/delete", taskHandler.DeleteFromHTML)
	http.HandleFunc("/tasks/update", taskHandler.UpdateFromHTML)

	http.HandleFunc("/notes", noteHandler.ViewHTML)
	http.HandleFunc("/notes/html", noteHandler.CreateFromHTML)
	http.HandleFunc("/notes/update", noteHandler.UpdateFromHTML)
	http.HandleFunc("/notes/delete", noteHandler.DeleteFromHTML)

	http.HandleFunc("/sleep", sleepHandler.ViewHTML)
	http.HandleFunc("/sleep/html", sleepHandler.CreateFromHTML)
	http.HandleFunc("/sleep/json", sleepHandler.GetAllJSON)

	http.HandleFunc("/nutrition", nutritionHandler.ViewHTML)
	http.HandleFunc("/nutrition/html", nutritionHandler.CreateFromHTML)
	http.HandleFunc("/nutrition/json", nutritionHandler.GetAllJSON)

	http.HandleFunc("/activity", activityHandler.ViewHTML)
	http.HandleFunc("/activity/html", activityHandler.CreateFromHTML)
	http.HandleFunc("/activity/json", activityHandler.GetAllJSON)

	http.HandleFunc("/profile", userHandler.ViewProfile)
}
