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
		log.Println("No .env file found")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set")
	}

	client, err := repository.NewMongoClient(mongoURI)
	if err != nil {
		panic(err)
	}

	collection := client.Database("lifeflow").Collection("tasks")
	repo := repository.NewMongoTaskRepository(collection)

	service := service.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	app.RegisterRoutes(handler)
	app.Start()
}

// LifeFlow v1.1 - Initial working backend with Task CRUD
