package main

import (
	"Assignment3/internal/app"
	"Assignment3/internal/handlers"
	"Assignment3/internal/repository"
	"Assignment3/internal/service"
)

func main() {
	repo := repository.NewTaskRepository()
	service := service.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	app.RegisterRoutes(handler)
	app.Start()
}
