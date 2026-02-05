package main

import (
	"Assignment3/internal/app"
	"Assignment3/internal/handlers"
	"Assignment3/internal/repository"
	"Assignment3/internal/service"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("need MONGO_URI")
	}

	client, err := repository.NewMongoClient(mongoURI)
	if err != nil {
		log.Fatal("cant connect to mongo:", err)
	}

	taskCollection := client.Database("lifeflow").Collection("tasks")
	taskRepo := repository.NewMongoTaskRepository(taskCollection)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	noteCollection := client.Database("lifeflow").Collection("notes")
	noteRepo := repository.NewMongoNoteRepository(noteCollection)
	noteService := service.NewNoteService(noteRepo)
	noteHandler := handlers.NewNoteHandler(noteService)

	userHandler := handlers.NewUserHandler(taskService, noteService)

	app.RegisterRoutes(taskHandler, noteHandler, userHandler)
	log.Println("server starting on :8080")
	app.Start()
}
