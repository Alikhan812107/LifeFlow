package repository

import (
	"Assignment3/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoNutritionRepository struct {
	collection *mongo.Collection
}

func NewMongoNutritionRepository(col *mongo.Collection) NutritionRepository {
	return &MongoNutritionRepository{collection: col}
}

func (r *MongoNutritionRepository) Create(nutrition models.Nutrition) (models.Nutrition, error) {
	nutrition.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(context.Background(), nutrition)
	return nutrition, err
}

func (r *MongoNutritionRepository) GetAll() ([]models.Nutrition, error) {
	var nutritions []models.Nutrition
	opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}})
	cursor, err := r.collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &nutritions)
	return nutritions, err
}
