package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sleep struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	WokeUp    time.Time          `json:"woke_up" bson:"woke_up"`
	Slept     time.Time          `json:"slept" bson:"slept"`
	UserID    string             `json:"user_id" bson:"user_id"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}
