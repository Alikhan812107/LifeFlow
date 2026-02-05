package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title"`
	Body   string             `json:"body" bson:"body"`
	Done   bool               `json:"done" bson:"done"`
	Folder string             `json:"folder" bson:"folder"`
	UserID string             `json:"user_id" bson:"user_id"`
}
