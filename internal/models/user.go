package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"-" bson:"password"`
	TasksNum int                `json:"tasks_num" bson:"tasks_num"`
	NotesNum int                `json:"notes_num" bson:"notes_num"`
}
