package repository

import (
	"Assignment3/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoSleepRepository struct {
	collection *mongo.Collection
}

func NewMongoSleepRepository(col *mongo.Collection) SleepRepository {
	return &MongoSleepRepository{collection: col}
}

func (r *MongoSleepRepository) Create(sleep models.Sleep) (models.Sleep, error) {
	sleep.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(context.Background(), sleep)
	return sleep, err
}

func (r *MongoSleepRepository) GetAll() ([]models.Sleep, error) {
	var sleeps []models.Sleep
	opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}})
	cursor, err := r.collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &sleeps)
	return sleeps, err
}
