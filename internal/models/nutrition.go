package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Nutrition struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Calories    int                `json:"calories" bson:"calories"`
	Water       float64            `json:"water" bson:"water"`
	Healthy     bool               `json:"healthy" bson:"healthy"`
	UserID      string             `json:"user_id" bson:"user_id"`
	Timestamp   time.Time          `json:"timestamp" bson:"timestamp"`
}
