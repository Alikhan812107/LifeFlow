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

	collection := client.Database("lifeflow").Collection("tasks")
	repo := repository.NewMongoTaskRepository(collection)
	service := service.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	app.RegisterRoutes(handler)
	log.Println("server starting on :8080")
	app.Start()

}
