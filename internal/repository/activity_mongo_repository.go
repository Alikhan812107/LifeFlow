package repository

import (
	"Assignment3/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoActivityRepository struct {
	collection *mongo.Collection
}

func NewMongoActivityRepository(col *mongo.Collection) ActivityRepository {
	return &MongoActivityRepository{collection: col}
}

func (r *MongoActivityRepository) Create(activity models.Activity) (models.Activity, error) {
	activity.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(context.Background(), activity)
	return activity, err
}

func (r *MongoActivityRepository) GetAll() ([]models.Activity, error) {
	var activities []models.Activity
	opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}})
	cursor, err := r.collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &activities)
	return activities, err
}
