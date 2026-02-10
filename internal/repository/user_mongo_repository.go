package repository

import (
	"Assignment3/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(collection *mongo.Collection) UserRepository {
	return &MongoUserRepository{collection: collection}
}

func (r *MongoUserRepository) Create(user models.User) (models.User, error) {
	result, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return models.User{}, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (r *MongoUserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return user, err
}

func (r *MongoUserRepository) GetByID(id primitive.ObjectID) (models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	return user, err
}
