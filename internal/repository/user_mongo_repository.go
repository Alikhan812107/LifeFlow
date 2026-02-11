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

func (r *UserMongoRepository) UpdateAvatar(id string, avatar string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"avatar": avatar}}

	res, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		_, err = r.collection.InsertOne(context.Background(), bson.M{
			"_id":    id,
			"name":   "John Student",
			"email":  "john@student.com",
			"avatar": avatar,
		})
	}

	return err
}
