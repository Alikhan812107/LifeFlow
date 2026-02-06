package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Activity struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Description string             `json:"description" bson:"description"`
	UserID      string             `json:"user_id" bson:"user_id"`
	Timestamp   time.Time          `json:"timestamp" bson:"timestamp"`
}
