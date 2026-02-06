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

	sleepCollection := client.Database("lifeflow").Collection("sleep")
	sleepRepo := repository.NewMongoSleepRepository(sleepCollection)
	sleepService := service.NewSleepService(sleepRepo)
	sleepHandler := handlers.NewSleepHandler(sleepService)

	nutritionCollection := client.Database("lifeflow").Collection("nutrition")
	nutritionRepo := repository.NewMongoNutritionRepository(nutritionCollection)
	nutritionService := service.NewNutritionService(nutritionRepo)
	nutritionHandler := handlers.NewNutritionHandler(nutritionService)

	activityCollection := client.Database("lifeflow").Collection("activity")
	activityRepo := repository.NewMongoActivityRepository(activityCollection)
	activityService := service.NewActivityService(activityRepo)
	activityHandler := handlers.NewActivityHandler(activityService)

	userHandler := handlers.NewUserHandler(taskService, noteService)

	app.RegisterRoutes(taskHandler, noteHandler, userHandler, sleepHandler, nutritionHandler, activityHandler)
	log.Println("server starting on :8080")
	app.Start()
}
