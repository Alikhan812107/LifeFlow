package app

import (
	"Assignment3/internal/handlers"
	"net/http"
)

func RegisterRoutes(handler *handlers.TaskHandler) {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.Create(w, r)
		}
		if r.Method == http.MethodGet {
			handler.GetAll(w, r)
		}
	})

	http.HandleFunc("/tasks/item", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetByID(w, r)
		}
		if r.Method == http.MethodPut {
			handler.Update(w, r)
		}
		if r.Method == http.MethodDelete {
			handler.Delete(w, r)
		}
	})

	http.HandleFunc("/", handler.ViewHTML)
	http.HandleFunc("/tasks/html", handler.CreateFromHTML)

}
