package repository

import (
	"Assignment3/internal/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTaskRepository struct {
	collection *mongo.Collection
}

func NewMongoTaskRepository(col *mongo.Collection) TaskRepository {
	return &MongoTaskRepository{collection: col}
}

func (r *MongoTaskRepository) Create(task models.Task) (models.Task, error) {
	task.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(context.Background(), task)
	return task, err
}

func (r *MongoTaskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &tasks)
	return tasks, err
}

func (r *MongoTaskRepository) GetByID(id primitive.ObjectID) (models.Task, error) {
	var task models.Task
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	return task, err
}

func (r *MongoTaskRepository) Update(id primitive.ObjectID, task models.Task) (models.Task, error) {
	task.ID = id
	_, err := r.collection.ReplaceOne(context.Background(), bson.M{"_id": id}, task)
	return task, err
}

func (r *MongoTaskRepository) Delete(id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("not found")
	}
	return nil
}
