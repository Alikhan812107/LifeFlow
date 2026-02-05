package repository

import (
	"Assignment3/internal/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoNoteRepository struct {
	collection *mongo.Collection
}

func NewMongoNoteRepository(col *mongo.Collection) NoteRepository {
	return &MongoNoteRepository{collection: col}
}

func (r *MongoNoteRepository) Create(note models.Note) (models.Note, error) {
	note.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(context.Background(), note)
	return note, err
}

func (r *MongoNoteRepository) GetAll() ([]models.Note, error) {
	var notes []models.Note
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &notes)
	return notes, err
}

func (r *MongoNoteRepository) GetByID(id primitive.ObjectID) (models.Note, error) {
	var note models.Note
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&note)
	return note, err
}

func (r *MongoNoteRepository) Update(id primitive.ObjectID, note models.Note) (models.Note, error) {
	note.ID = id
	_, err := r.collection.ReplaceOne(context.Background(), bson.M{"_id": id}, note)
	return note, err
}

func (r *MongoNoteRepository) Delete(id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("not found")
	}
	return nil
}