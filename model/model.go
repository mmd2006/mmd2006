package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description,omitempty"`
	Completed   bool               `json:"completed" bson:"completed"`
}
