package repository

import (
	"Assignment3/internal/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository struct {
	collection *mongo.Collection
}

func NewUserMongoRepository(db *mongo.Database) *UserMongoRepository {
	return &UserMongoRepository{
		collection: db.Collection("users"),
	}
}

func (r *UserMongoRepository) GetByID(id string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserMongoRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserMongoRepository) Create(user models.User) (*models.User, error) {
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserMongoRepository) UpdateAvatar(id string, avatar string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"avatar": avatar}}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return err
}
