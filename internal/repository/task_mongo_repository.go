package repository

import (
	"Assignment3/internal/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTaskRepository struct {
	collection *mongo.Collection
}

func NewMongoTaskRepository(col *mongo.Collection) TaskRepository {
	return &MongoTaskRepository{collection: col}
}

func (r *MongoTaskRepository) Create(task models.Task) models.Task {
	r.collection.InsertOne(context.Background(), task)
	return task
}

func (r *MongoTaskRepository) GetAll() []models.Task {
	var tasks []models.Task
	cursor, _ := r.collection.Find(context.Background(), map[string]interface{}{})
	cursor.All(context.Background(), &tasks)
	return tasks
}

// TEMP stubs (so code compiles)

func (r *MongoTaskRepository) GetByID(id int) (models.Task, bool) {
	return models.Task{}, false
}

func (r *MongoTaskRepository) Update(id int, task models.Task) (models.Task, bool) {
	return models.Task{}, false
}

func (r *MongoTaskRepository) Delete(id int) bool {
	return false
}
