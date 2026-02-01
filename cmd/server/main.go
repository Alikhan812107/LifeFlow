package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"Assignment3/internal/app"
	"Assignment3/internal/handlers"
	"Assignment3/internal/repository"
	"Assignment3/internal/service"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	_ = godotenv.Load()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("lifeflow")

	taskRepo := repository.NewMongoTaskRepository(db.Collection("tasks"))
	userRepo := repository.NewUserMongoRepository(db)

	taskService := service.NewTaskService(taskRepo)
	authService := service.NewAuthService(userRepo, "super-secret-key")

	taskHandler := handlers.NewTaskHandler(taskService)
	authHandler := handlers.NewAuthHandler(authService)

	app.RegisterRoutes(taskHandler, authHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
