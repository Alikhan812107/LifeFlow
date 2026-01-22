package app

import (
	"Assignment3/internal/config"
	"Assignment3/internal/repository"
	"database/sql"
)

func NewServer(db *sql.DB, cfg *config.Config) *Server {
	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	authService := service.NewAuthService(userRepo)
	taskService := service.NewTaskService(taskRepo)

	authHandler := handlers.NewAuthHandler(authService)
	taskHandler := handlers.NewTaskHandler(taskService)

	router := NewRouter(authHandler, taskHandler)

	return &Server{
		router: router,
	}
}
